package handler

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_proto/protos/protosService"
	"go-micro.dev/v4"
)

// 新增加handler后，请在此注册
func RegisterHandler(s micro.Service) error {
	err := protosService.RegisterWeatherServiceHandler(s.Server(), new(WeatherServiceHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterWeatherServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterIPServiceHandler(s.Server(), new(IPServiceHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterIPServiceHandler 错误:%s", err.Error())
		return err
	}
	return nil
}
