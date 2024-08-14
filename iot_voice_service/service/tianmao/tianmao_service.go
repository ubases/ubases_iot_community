package tianmao

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

// 天猫精灵设置设备属性
func SetDeviceProperty(reqBody []byte, userId string, adjust bool, voiceNo string) ([]entitys.DeviceResponse, error) {
	var channels []string
	var controlData []*entitys.ControlData
	var devRespList []entitys.DeviceResponse
	checked := map[string]struct{}{}
	ctx := context.Background()
	devList := gjson.GetBytes(reqBody, "payload.deviceIds").Array()
	params := gjson.GetBytes(reqBody, "payload.params").Map()
	iotlogger.LogHelper.Helper.Debug("test1")
	// 将设备列表语控功能点数据转换为爱星云功能点数据
	Subjects := make(map[string]bool)
	for i := range devList {
		_, ctl, err := common.VoiceToAxyFuncData(ctx, devList[i].String(), params, adjust, voiceNo)
		if err != nil {
			// 需将设备错误信息放入响应列表
			iotlogger.LogHelper.Helper.Error("convert voice funcs to axy funs error: ", err)
			continue
		}
		controlData = append(controlData, ctl)
		Subjects[fmt.Sprintf(iotconst.HKEY_ACK_DATA_PUB_PREFIX+".%s.%s", ctl.ProductKey, ctl.DeviceId)] = true
		//channels = append(channels, persist.GetRedisKey(iotconst.HKEY_ACK_DATA_PUB_PREFIX+".%s.%s", ctl.ProductKey, ctl.DeviceId))
	}
	suber := common.Subscriber{ID: iotutil.GetSecret(8), Subjects: Subjects, Channel: make(chan common.EventData, 64)}
	common.GetNatsSubscriber().GetPublisher().AddSubscriber(&suber)
	defer func() {
		common.GetNatsSubscriber().GetPublisher().DelSubscriber(&suber)
		close(suber.Channel)
	}() //pubsub := cached.RedisStore.GetClient().Subscribe(ctx, channels...)
	//defer pubsub.Close()
	// 开协程下发数据到设备
	for i := range controlData {
		if _, err := common.PubControl(controlData[i].ProductKey, controlData[i].DeviceId, controlData[i].Data); err != nil {
			iotlogger.LogHelper.Helper.Errorf("pub control %s %s error: %v", controlData[i].ProductKey, controlData[i].DeviceId, err)
		}
	}
	iotlogger.LogHelper.Helper.Debug("control data: ", controlData)
	iotlogger.LogHelper.Helper.Debug("channel data: ", channels)
	// 订阅redis，延时等待一秒，若没有收到设备Ack消息，则直接返回错误信息给语控平台
	ctxNew, cancel := context.WithTimeout(ctx, 1500*time.Millisecond)
	defer cancel()
	for {
		select {
		case <-ctxNew.Done():
			// 超过一秒钟，需要将没有响应的设备补充错误信息
			for i := range controlData {
				if _, ok := checked[controlData[i].DeviceId]; !ok {
					devRespList = append(devRespList, entitys.DeviceResponse{
						DeviceId:  controlData[i].DeviceId,
						ErrorCode: "DEVICE_NOT_SUPPORT_FUNCTION",
						Message:   "device not support",
					})
				}
			}
			iotlogger.LogHelper.Helper.Debug("control device data: ", devRespList)
			return devRespList, nil
		//case msg := <-pubsub.Channel():
		//	devId := strings.Split(msg.Channel, ".")[2]
		case msg := <-suber.Channel:
			devId := strings.Split(msg.Subject, ".")[3]
			devRespList = append(devRespList, entitys.DeviceResponse{
				DeviceId:  devId,
				ErrorCode: "SUCCESS",
				Message:   "SUCCESS",
			})
			checked[devId] = struct{}{}
			iotlogger.LogHelper.Helper.Debug("control deivce success: ", devRespList)
			if len(devList) == len(checked) {
				return devRespList, nil
			}
		}
	}
}
