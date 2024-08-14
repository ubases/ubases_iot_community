package alexa

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

type AlexaService struct {
}

// 天猫精灵设置设备属性
func (s *AlexaService) SetDeviceProperty(reqBody []byte, header entitys.CommonHeader, userId string, adjust bool, voiceNo string) ([]entitys.DeviceResponse, error) {
	var (
		channels    []string
		controlData []*entitys.ControlData
		devRespList []entitys.DeviceResponse
		devRespMap  = make(map[string]entitys.DeviceResponse)
		checked     = map[string]struct{}{}
		ctx         = context.Background()
	)
	devList := []string{gjson.GetBytes(reqBody, "directive.endpoint.endpointId").String()}
	params := map[string]gjson.Result{}
	voiceCode := header.Namespace
	if header.Instance != "" {
		voiceCode = fmt.Sprintf("%s#%s", header.Namespace, header.Instance)
	}
	playload := gjson.GetBytes(reqBody, "directive.payload").Map()
	if header.Namespace == "Alexa.RangeController" {
		params[voiceCode] = playload["rangeValue"]
	} else if header.Namespace == "Alexa.ModeController" {
		params[voiceCode] = playload["mode"]
	} else {
		params[voiceCode] = gjson.Parse("{\"value\": \"" + header.Name + "\"}").Get("value")
	}
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
		channels = append(channels)
	}

	suber := common.Subscriber{ID: iotutil.GetSecret(8), Subjects: Subjects, Channel: make(chan common.EventData, 64)}
	common.GetNatsSubscriber().GetPublisher().AddSubscriber(&suber)
	defer func() {
		common.GetNatsSubscriber().GetPublisher().DelSubscriber(&suber)
		close(suber.Channel)
	}()

	//pubsub := cached.RedisStore.GetClient().Subscribe(ctx, channels...)
	//defer pubsub.Close()
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
	if controlData == nil && len(controlData) == 0 {
		iotlogger.LogHelper.Helper.Debug("未获取到有效的设备（考虑设备离线）: ", iotutil.ToString(devRespList))
		return devRespList, nil
	}
	iotlogger.LogHelper.Helper.Debug("control data: ", iotutil.ToString(controlData))
	iotlogger.LogHelper.Helper.Debug("channel data: ", iotutil.ToString(channels))
	// 订阅redis，延时等待一秒，若没有收到设备Ack消息，则直接返回错误信息给语控平台
	ctxNew, cancel := context.WithTimeout(ctx, 1500*time.Millisecond)
	defer cancel()
	for {
		select {
		case <-ctxNew.Done():
			devRespList = make([]entitys.DeviceResponse, 0)
			for _, dev := range devRespMap {
				if _, ok := checked[dev.DeviceId]; !ok {
					devRespList = append(devRespList, entitys.DeviceResponse{
						DeviceId:  dev.DeviceId,
						ErrorCode: iotutil.GetStringAndDef(dev.ErrorCode, "DEVICE_NOT_SUPPORT_FUNCTION"),
						Message:   iotutil.GetStringAndDef(dev.Message, "device not support"),
						Data:      dev.Data,
					})
				} else {
					devRespList = append(devRespList, dev)
				}
			}
			// 超过一秒钟，需要将没有响应的设备补充错误信息

			iotlogger.LogHelper.Helper.Debug("control device data: ", devRespList)
			return devRespList, nil
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
			iotlogger.LogHelper.Helper.Debug("control deivce success: ", iotutil.ToString(devRespMap))
			if len(devList) == len(checked) {
				return devRespList, nil
			}
		}
	}
}
