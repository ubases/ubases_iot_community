package handler

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_proto/protos/protosService"

	"go-micro.dev/v4"
)

// 新增加handler后，请在此注册
func RegisterHandler(s micro.Service) error {
	err := protosService.RegisterConfigDictDataServiceHandler(s.Server(), new(ConfigDictDataHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterConfigDictDataServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterConfigDictTypeServiceHandler(s.Server(), new(ConfigDictTypeHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterConfigDictTypeServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterConfigLicenseServiceHandler(s.Server(), new(ConfigLicenseHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterConfigLicenseServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterConfigOssServiceHandler(s.Server(), new(ConfigOssHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterConfigOssServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterConfigPlatformServiceHandler(s.Server(), new(ConfigPlatformHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterConfigPlatformServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterConfigTranslateServiceHandler(s.Server(), new(ConfigTranslateHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterConfigTranslateServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterSysAreaServiceHandler(s.Server(), new(SysAreaHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterSysAreaServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterSysRegionServiceHandler(s.Server(), new(SysRegionHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterSysRegionServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterSysRegionServerServiceHandler(s.Server(), new(SysRegionServerHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterSysRegionServerServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterConfigTranslateLanguageServiceHandler(s.Server(), new(ConfigTranslateLanguageHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterConfigTranslateLanguageServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterLangTranslateServiceHandler(s.Server(), new(LangTranslateHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterLangTranslateServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterLangTranslateTypeServiceHandler(s.Server(), new(LangTranslateTypeHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterLangTranslateTypeServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterLangResourcesServiceHandler(s.Server(), new(LangResourcesHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterLangResourcesServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterLangCustomResourcesServiceHandler(s.Server(), new(LangCustomResourcesHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterLangCustomResourcesServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterLangResourcePackageServiceHandler(s.Server(), new(LangResourcePackageHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterLangResourcePackageServiceHandler 错误:%s", err.Error())
		return err
	}
	return nil
}
