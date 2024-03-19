package handler

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_proto/protos/protosService"

	"go-micro.dev/v4"
)

// 新增加handler后，请在此注册
func RegisterHandler(s micro.Service) error {

	err := protosService.RegisterOpenUserServiceHandler(s.Server(), new(OpenUserHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpenUserServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOpenAuthRuleServiceHandler(s.Server(), new(OpenAuthRuleHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpenAuthRuleServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOpenCasbinRuleServiceHandler(s.Server(), new(OpenCasbinRuleHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpenCasbinRuleServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOpenCompanyConnectServiceHandler(s.Server(), new(OpenCompanyConnectHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpenCompanyConnectServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOpenCompanyServiceHandler(s.Server(), new(OpenCompanyHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpenCompanyServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOpenConfigServiceHandler(s.Server(), new(OpenConfigHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpenConfigServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOpenLoginLogServiceHandler(s.Server(), new(OpenLoginLogHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpenLoginLogServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOpenModelInfoServiceHandler(s.Server(), new(OpenModelInfoHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpenModelInfoServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOpenOperLogServiceHandler(s.Server(), new(OpenOperLogHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpenOperLogServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOpenRoleServiceHandler(s.Server(), new(OpenRoleHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpenRoleServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOpenUserCompanyServiceHandler(s.Server(), new(OpenUserCompanyHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpenUserCompanyServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOpenCompanyAuthLogsServiceHandler(s.Server(), new(OpenCompanyAuthLogsHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpenCompanyAuthLogsServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOpenUserOnlineServiceHandler(s.Server(), new(OpenUserOnlineHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpenUserOnlineServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterOpenAuthQuantityServiceHandler(s.Server(), new(OpenAuthQuantityHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpenAuthQuantityServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterDeveloperServiceHandler(s.Server(), new(DeveloperHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterDeveloperServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOpenAuthHandler(s.Server(), new(OpenAuthHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpenAuthHandler:%s", err.Error())
		return err
	}

	return nil
}
