package sync_update

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_proto/protos/protosService"
	"cloud_platform/iot_voice_service/rpc/rpcclient"
	"cloud_platform/iot_voice_service/service/common"
	"cloud_platform/iot_voice_service/service/tianmao"
	"context"
	"errors"
)

type OnlineDeviceSvc struct {
	Data  *iotstruct.DeviceRedisData
	DevId string
}

// OnlineDevice 在线设备处理
func (s *OnlineDeviceSvc) OnlineDevice() error {
	// defer iotutil.PanicHandler()
	if err := s.check(); err != nil {
		return err
	}
	s.DevId = s.Data.DeviceId
	payloadMap := s.Data.Data.(map[string]interface{})
	onlineStatus := payloadMap[iotconst.FIELD_ONLINE]
	// 通过产品Key获取产品语控配置信息
	ctx := context.Background()
	opmVoice, err := rpcclient.ClienOpmVoiceProductService.Find(ctx, &protosService.OpmVoiceProductFilter{
		ProductKey: s.Data.ProductKey,
	})
	if err != nil {
		return err
	}
	if len(opmVoice.Data) == 0 {
		return errors.New("通过产品key获取语控配置信息为空")
	}

	devStatus, err := common.GetDeviceInfo(s.Data.DeviceId)
	userId := devStatus["userId"]
	if userId == "" {
		return errors.New("未获取到用户Id")
	}
	//同步在线状态
	//go google.GetSmartHome().RequestSync(context.Background(), userId)
	//天猫的在线数据同步
	tianmao.RequestSync(s.DevId, userId, onlineStatus, *opmVoice.Data[0])
	//小米同步用户的设备数据
	//xiaomi.RequestSync(userId)
	//Alexa的在线设备同步
	//alexa.RequestSync(userId)

	return nil
}

func (s *OnlineDeviceSvc) check() error {
	if s.Data == nil {
		return errors.New("必须初始化data")
	}
	if s.Data.Data == nil {
		return errors.New("参数异常，Payload")
	}
	return nil
}
