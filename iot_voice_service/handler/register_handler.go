package handler

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_proto/protos/protosService"

	"go-micro.dev/v4"
)

// 新增加handler后，请在此注册
func RegisterHandler(s micro.Service) error {
	err := protosService.RegisterVoiceServiceHandler(s.Server(), new(ClientInfo))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterVoiceServiceHandler 错误:%s", err.Error())
		return err
	}
	return nil
}
