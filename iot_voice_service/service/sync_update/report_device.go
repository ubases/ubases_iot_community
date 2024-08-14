package sync_update

import (
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"cloud_platform/iot_voice_service/rpc/rpcclient"
	"cloud_platform/iot_voice_service/service/google"
	"context"
	"errors"
)

type ReportDeviceSvc struct {
	Data  *iotstruct.DeviceRedisData
	DevId string
}

// OnlineDevice 在线设备处理
func (s *ReportDeviceSvc) ReportDevice() error {
	// defer iotutil.PanicHandler()
	if err := s.check(); err != nil {
		return err
	}
	s.DevId = s.Data.DeviceId
	payloadMap := s.Data.Data.(map[string]interface{})
	ctx := context.Background()
	devTriad, err := rpcclient.ClientIotDeviceTriadService.Find(ctx, &protosService.IotDeviceTriadFilter{
		Did: s.Data.DeviceId,
	})
	if err != nil {
		return err
	}
	if len(devTriad.Data) == 0 {
		return errors.New("三元组数据为空")
	}

	//同步设备状态
	go google.GetSmartHome().ReportState(context.Background(), s.DevId)

	// 通过产品Key获取产品语控配置信息
	opmVoice, err := rpcclient.ClienOpmVoiceProductService.Find(ctx, &protosService.OpmVoiceProductFilter{
		ProductKey: s.Data.ProductKey,
	})
	if err != nil {
		return err
	}
	if len(opmVoice.Data) == 0 {
		return errors.New("通过产品key获取语控配置信息为空")
	}

	userId := iotutil.ToString(devTriad.Data[0].UserId)
	RunSync(userId, s.DevId, payloadMap, *opmVoice.Data[0])
	return nil
}

func (s *ReportDeviceSvc) check() error {
	if s.Data == nil {
		return errors.New("必须初始化data")
	}
	if s.Data.Data == nil {
		return errors.New("参数异常，Payload")
	}
	return nil
}
