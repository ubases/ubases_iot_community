package handler

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_proto/protos/protosService"

	"go-micro.dev/v4"
)

// 新增加handler后，请在此注册
func RegisterHandler(s micro.Service) error {

	err := protosService.RegisterOemAppIntroduceServiceHandler(s.Server(), new(OemAppIntroduceHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterAppIntroduceServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOemAppAndroidCertServiceHandler(s.Server(), new(OemAppAndroidCertHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOemAppAndroidCertServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOemAppBuildRecordServiceHandler(s.Server(), new(OemAppBuildRecordHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOemAppBuildRecordServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOemAppFunctionConfigServiceHandler(s.Server(), new(OemAppFunctionConfigHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOemAppFunctionConfigServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOemAppServiceHandler(s.Server(), new(OemAppHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOemAppServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOemAppIosCertServiceHandler(s.Server(), new(OemAppIosCertHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOemAppIosCertServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOemAppPushCertServiceHandler(s.Server(), new(OemAppPushCertHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOemAppPushCertServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOemAppUiConfigServiceHandler(s.Server(), new(OemAppUiConfigHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOemAppUiConfigServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOemAppDefMenuServiceHandler(s.Server(), new(OemAppDefMenuHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOemAppDefMenuServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOemAppDocServiceHandler(s.Server(), new(OemAppDocHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOemAppDocServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOemAppDocRelationServiceHandler(s.Server(), new(OemAppDocRelationHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOemAppDocRelationServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOemAppDocDirServiceHandler(s.Server(), new(OemAppDocDirHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOemAppDocDirServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOemAppEntryServiceHandler(s.Server(), new(OemAppEntryHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOemAppEntryServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOemAppEntrySetingServiceHandler(s.Server(), new(OemAppEntrySetingHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOemAppEntrySetingServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOemAppFlashScreenServiceHandler(s.Server(), new(OemAppFlashScreenHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOemAppFlashScreenServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOemAppFlashScreenUserServiceHandler(s.Server(), new(OemAppFlashScreenUserHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOemAppFlashScreenUserServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOemAppVersionRecordServiceHandler(s.Server(), new(OemAppVersionRecordHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOemAppVersionRecordServiceHandler 错误:%s", err.Error())
	}

	err = protosService.RegisterOemAppBasicUiSettingServiceHandler(s.Server(), new(OemAppBasicUiSettingHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOemAppBasicUiSettingServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOemAppAssistReleaseServiceHandler(s.Server(), new(OemAppAssistReleaseHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOemAppAssistReleaseServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOemAppTemplateServiceHandler(s.Server(), new(OemAppTemplateHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOemAppTemplateServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOemAppTemplateFunctionServiceHandler(s.Server(), new(OemAppTemplateFunctionHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOemAppTemplateFunctionServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOemAppTemplateMenuServiceHandler(s.Server(), new(OemAppTemplateMenuHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOemAppTemplateMenuServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOemAppTemplateSkinServiceHandler(s.Server(), new(OemAppTemplateSkinHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOemAppTemplateSkinServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOemAppTemplateThirdPartyServiceHandler(s.Server(), new(OemAppTemplateThirdPartyHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOemAppTemplateThirdPartyServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOemAppTemplateUiServiceHandler(s.Server(), new(OemAppTemplateUiHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOemAppTemplateUiServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterPublicAppVersionServiceHandler(s.Server(), new(PublicAppVersionHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterPublicAppVersionServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOemFeedbackTypeServiceHandler(s.Server(), new(OemFeedbackTypeHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOemFeedbackTypeServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOemAppCustomRecordServiceHandler(s.Server(), new(OemAppCustomRecordHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOemAppCustomRecordServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOemAppDebuggerServiceHandler(s.Server(), new(OemAppDebuggerHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOemAppDebuggerServiceHandler 错误:%s", err.Error())
		return err
	}

	return nil
}
