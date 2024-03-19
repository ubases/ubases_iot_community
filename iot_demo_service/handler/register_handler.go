package handler

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_proto/protos/protosService"

	"go-micro.dev/v4"
)

// 新增加handler后，请在此注册
func RegisterHandler(s micro.Service) error {
	err := protosService.RegisterConfigDictTypeServiceHandler(s.Server(), new(ConfigDictTypeHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterConfigDictTypeServiceHandler 错误:%s", err.Error())
		return err
	}
	//参考前面内容，挨个注册其它Handler
	return nil
}
