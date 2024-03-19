package handler

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_proto/protos/protosService"

	"go-micro.dev/v4"
)

// 新增加handler后，请在此注册
func RegisterHandler(s micro.Service) error {
	err := protosService.RegisterAppLogServiceHandler(s.Server(), new(AppLogHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterAppLogServiceHandler 错误:%s", err.Error())
		return err
	}
	return nil
}
