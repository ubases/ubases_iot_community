package xiaomi

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_voice_service/entitys"
	"cloud_platform/iot_voice_service/service/common"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

type XiaomiService struct {
}

var (
	//小米定义设备基础信息
	XiaomiBaseProperties map[string]string = map[string]string{
		"1": "productName",
		"2": "productKey",
		"3": "did",
		"4": "fwVer",
		"5": "serialNumber",
	}
)

func (s *XiaomiService) SetDeviceProperty(reqBody []byte, header entitys.XiaomiRequest, userId string, adjust bool, voiceNo string) ([]entitys.DeviceResponse, error) {
	var (
		channels    []string
		controlData []*entitys.ControlData
		devRespList []entitys.DeviceResponse
		devRespMap  = make(map[string]entitys.DeviceResponse)
		checked     = map[string]struct{}{}
		ctx         = context.Background()
	)
	params := map[string]gjson.Result{}
	devMap := map[string]string{}
	for _, p := range header.Properties {
		var property entitys.XiaomiProperties
		err := iotutil.StructToStructErr(p, &property)
		if err != nil {
			iotlogger.LogHelper.Helper.Error("属性对象转换失败: ", iotutil.ToString(p))
			continue
		}
		devMap[property.Did] = property.Did
		jsonStr := iotutil.ToString(map[string]interface{}{
			"value": property.Value,
		})
		params[iotutil.ToString(property.Piid)] = gjson.Parse(jsonStr).Get("value")
	}
	devList := iotutil.Keys(devMap)
	// 将设备列表语控功能点数据转换为爱星云功能点数据
	Subjects := make(map[string]bool)
	for i := range devList {
		isOnline, ctl, err := common.VoiceToAxyFuncData(ctx, devList[i], params, adjust, voiceNo)
		if err != nil {
			// 需将设备错误信息放入响应列表
			iotlogger.LogHelper.Helper.Error("convert voice funcs to axy funs error: ", err)
			continue
		}
		if !isOnline {
			devRespMap[devList[i]] = entitys.DeviceResponse{
				DeviceId:  devList[i],
				ErrorCode: "DEVICE_OFFLINE",
				Message:   "offline",
				Data:      ctl.Data,
			}
			devRespList = append(devRespList, devRespMap[devList[i]])
			continue
		}
		controlData = append(controlData, ctl)
		Subjects[fmt.Sprintf(iotconst.HKEY_ACK_DATA_PUB_PREFIX+".%s.%s", ctl.ProductKey, ctl.DeviceId)] = true
		//channels = append(channels, persist.GetRedisKey(iotconst.HKEY_ACK_DATA_PUB_PREFIX+".%s.%s", ctl.ProductKey, ctl.DeviceId))
	}
	//pubsub := cached.RedisStore.GetClient().Subscribe(ctx, channels...)
	//defer pubsub.Close()
	suber := common.Subscriber{ID: iotutil.GetSecret(8), Subjects: Subjects, Channel: make(chan common.EventData, 64)}
	common.GetNatsSubscriber().GetPublisher().AddSubscriber(&suber)
	defer func() {
		common.GetNatsSubscriber().GetPublisher().DelSubscriber(&suber)
		close(suber.Channel)
	}()
	// 开协程下发数据到设备

	for i := range controlData {
		devRespMap[controlData[i].DeviceId] = entitys.DeviceResponse{
			DeviceId: controlData[i].DeviceId,
			Data:     controlData[i].DataVoice,
		}
		if _, err := common.PubControl(controlData[i].ProductKey, controlData[i].DeviceId, controlData[i].Data); err != nil {
			iotlogger.LogHelper.Helper.Errorf("pub control %s %s error: %v", controlData[i].ProductKey, controlData[i].DeviceId, err)
		}
	}
	iotlogger.LogHelper.Helper.Debug("control data: ", iotutil.ToString(controlData))
	iotlogger.LogHelper.Helper.Debug("channel data: ", iotutil.ToString(channels))
	// 订阅redis，延时等待一秒，若没有收到设备Ack消息，则直接返回错误信息给语控平台
	ctxNew, cancel := context.WithTimeout(ctx, 1500*time.Millisecond)
	defer cancel()
	for {
		select {
		case <-ctxNew.Done():
			// 超过一秒钟，需要将没有响应的设备补充错误信息
			devRespList = make([]entitys.DeviceResponse, 0)
			for _, dev := range devRespMap {
				if _, ok := checked[dev.DeviceId]; !ok {
					devRespList = append(devRespList, entitys.DeviceResponse{
						DeviceId:  dev.DeviceId,
						ErrorCode: "DEVICE_NOT_SUPPORT_FUNCTION",
						Message:   "device not support",
						Data:      dev.Data,
					})
				} else {
					devRespList = append(devRespList, dev)
				}
			}
			iotlogger.LogHelper.Helper.Debug("control device data: ", devRespList)
			return devRespList, nil
		//case msg := <-pubsub.Channel():
		//	devId := strings.Split(msg.Channel, ".")[2]
		case msg := <-suber.Channel:
			devId := strings.Split(msg.Subject, ".")[3]
			devRespMap[devId] = entitys.DeviceResponse{
				DeviceId:  devId,
				ErrorCode: "SUCCESS",
				Message:   "SUCCESS",
				Data:      devRespMap[devId].Data,
			}
			devRespList = append(devRespList, devRespMap[devId])
			checked[devId] = struct{}{}
			iotlogger.LogHelper.Helper.Debug("control deivce success: ", devRespMap)
			if len(devList) == len(checked) {
				return devRespList, nil
			}
		}
	}
}

func (s XiaomiService) GetDeviceProperty(req entitys.XiaomiRequest, userId string, adjust bool, voiceNo string) ([]entitys.DeviceResponse, error) {
	ctx := context.Background()
	var devRespList []entitys.DeviceResponse
	// 将设备列表语控功能点数据转换为爱星云功能点数据
	deviceProps := make(map[string][]entitys.XiaomiProperties)
	for _, p := range req.Properties {
		var item entitys.XiaomiProperties
		err := iotutil.StructToStructErr(p, &item)
		if err != nil {
			continue
		}
		if _, ok := deviceProps[item.Did]; !ok {
			deviceProps[item.Did] = make([]entitys.XiaomiProperties, 0)
		}
		deviceProps[item.Did] = append(deviceProps[item.Did], item)
	}

	for did, properties := range deviceProps {
		params := map[string]gjson.Result{}
		baseParams := map[string]string{}
		for _, property := range properties {
			dpId := iotutil.ToString(property.Piid)
			//小米基础属性获取
			if property.Siid == 1 {
				baseParams[dpId] = XiaomiBaseProperties[dpId]
			} else {
				params[dpId] = gjson.Parse("{\"value\": \"\"}").Get("value")
			}
		}
		ctl, err := common.VoiceGetFuncData(ctx, did, params, baseParams, adjust, voiceNo)
		if err != nil {
			// 需将设备错误信息放入响应列表
			iotlogger.LogHelper.Helper.Error("convert voice funcs to axy funs error: ", err)
			continue
		}
		devRespList = append(devRespList, entitys.DeviceResponse{
			DeviceId:  did,
			ErrorCode: "SUCCESS",
			Message:   "SUCCESS",
			Data:      ctl.Data,
			DataVoice: ctl.DataVoiceCode,
			DataDpid:  ctl.DataDpid,
		})
	}

	return devRespList, nil
}
