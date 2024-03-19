package handler

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_proto/protos/protosService"

	"go-micro.dev/v4"
)

// 新增加handler后，请在此注册
func RegisterHandler(s micro.Service) error {
	// err := protosService.RegisterOpenCompanyServiceHandler(s.Server(), new(OpenCompanyHandler))
	// if err != nil {
	// 	iotlogger.LogHelper.Errorf("RegisterOpenCompanyServiceHandler 错误:%s", err.Error())
	// 	return err
	// }
	// err = protosService.RegisterOpenUserServiceHandler(s.Server(), new(OpenUserHandler))
	// if err != nil {
	// 	iotlogger.LogHelper.Errorf("RegisterOpenUserServiceHandler 错误:%s", err.Error())
	// 	return err
	// }
	err := protosService.RegisterSysApisServiceHandler(s.Server(), new(SysApisHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterSysApiServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterSysAuthRuleServiceHandler(s.Server(), new(SysAuthRuleHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterSysCasbinRuleServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterSysCasbinRuleServiceHandler(s.Server(), new(SysCasbinRuleHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterSysCasbinRuleServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterSysConfigServiceHandler(s.Server(), new(SysConfigHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterSysConfigServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterSysDeptServiceHandler(s.Server(), new(SysDeptHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterSysDeptServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterSysDictDataServiceHandler(s.Server(), new(SysDictDataHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterSysDictDataServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterSysDictTypeServiceHandler(s.Server(), new(SysDictTypeHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterSysDictTypeServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterSysJobServiceHandler(s.Server(), new(SysJobHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterSysJobServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterSysLoginLogServiceHandler(s.Server(), new(SysLoginLogHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterSysLoginLogServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterSysModelInfoServiceHandler(s.Server(), new(SysModelInfoHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterSysModelInfoServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterSysOperLogServiceHandler(s.Server(), new(SysOperLogHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterSysOperLogServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterSysPostServiceHandler(s.Server(), new(SysPostHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterSysPostServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterSysRoleDeptServiceHandler(s.Server(), new(SysRoleDeptHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterSysRoleDeptServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterSysRoleServiceHandler(s.Server(), new(SysRoleHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterSysRoleServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterSysUserServiceHandler(s.Server(), new(SysUserHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterSysUserServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterSysUserOnlineServiceHandler(s.Server(), new(SysUserOnlineHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterSysUserOnlineServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterSysUserPostServiceHandler(s.Server(), new(SysUserPostHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterSysUserPostServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterSysWebSetServiceHandler(s.Server(), new(SysWebSetHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterSysWebSetServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterSysAppDocDirServiceHandler(s.Server(), new(SysAppDocDirHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterSysAppDocDirServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterSysAppEntrySetingServiceHandler(s.Server(), new(SysAppEntrySetingHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterSysAppEntrySetingServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterSysAppEntryServiceHandler(s.Server(), new(SysAppEntryHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterSysAppEntryServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterSysAppHelpCenterServiceHandler(s.Server(), new(SysAppHelpCenterHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterSysAppHelpCenterServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterCloudAuthHandler(s.Server(), new(CloudAuthHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterCloudAuthHandler:%s", err.Error())
		return err
	}

	RegisterCasbinRuleExtHandler(s)
	return nil
}
