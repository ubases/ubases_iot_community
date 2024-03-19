package handler

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_proto/protos/protosService"

	"go-micro.dev/v4"
)

func RegisterHandler(service micro.Service) error {
	var err error
	//产品、物模型
	RegisterTPmProductHandler(service)
	RegisterTPmProductTypeHandler(service)
	RegisterTPmThingModelHandler(service)
	RegisterTPmThingModelPropertiesHandler(service)
	RegisterTPmThingModelServicesHandler(service)
	RegisterTPmThingModelEventsHandler(service)

	//芯片模组、固件设置
	err = protosService.RegisterPmFirmwareServiceHandler(service.Server(), new(PmFirmwareHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterPmFirmwareServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterPmFirmwareSettingServiceHandler(service.Server(), new(PmFirmwareSettingHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterPmFirmwareSettingServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterPmFirmwareVersionServiceHandler(service.Server(), new(PmFirmwareVersionHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterPmFirmwareVersionServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterPmModuleServiceHandler(service.Server(), new(PmModuleHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterPmModuleServiceHandler 错误:%s", err.Error())
		return err
	}
	//add by chensg,at 2022-04-27
	err = protosService.RegisterPmControlPanelsServiceHandler(service.Server(), new(PmControlPanelsHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterPmModuleServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterTplTestcaseTemplateServiceHandler(service.Server(), new(TplTestcaseTemplateHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterTplTestcaseTemplateServiceHandler 错误:%s", err.Error())
		return err
	}

	//开放平台-固件
	err = protosService.RegisterOpmFirmwareServiceHandler(service.Server(), new(OpmFirmwareHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpmFirmwareServiceHandler 错误:%s", err.Error())
		return err
	}
	//开放平台-固件版本
	err = protosService.RegisterOpmFirmwareVersionServiceHandler(service.Server(), new(OpmFirmwareVersionHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpmFirmwareVersionServiceHandler 错误:%s", err.Error())
		return err
	}

	//开发平台产品
	err = protosService.RegisterOpmProductServiceHandler(service.Server(), new(OpmProductHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpmFirmwareVersionServiceHandler 错误:%s", err.Error())
		return err
	}

	//开发平台产品 - 配网引导
	err = protosService.RegisterOpmNetworkGuideServiceHandler(service.Server(), new(OpmNetworkGuideHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpmNetworkGuideServiceHandler 错误:%s", err.Error())
		return err
	}

	//开发平台产品 - 物模型
	err = protosService.RegisterOpmThingModelServiceHandler(service.Server(), new(OpmThingModelHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpmThingModelServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterOpmProductModuleRelationServiceHandler(service.Server(), new(OpmProductModuleRelationHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpmProductModuleRelationServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOpmProductPanelRelationServiceHandler(service.Server(), new(OpmProductPanelRelationHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpmProductPanelRelationServiceHandler 错误:%s", err.Error())
		return err
	}
	// OTA升级相关表
	err = protosService.RegisterOpmOtaPkgServiceHandler(service.Server(), new(OpmOtaPkgHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpmOtaPublishServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterOpmOtaPublishServiceHandler(service.Server(), new(OpmOtaPublishHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpmOtaPublishServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterOpmOtaPublishLogServiceHandler(service.Server(), new(OpmOtaPublishLogHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpmOtaPublishLogServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterOpmProductTestReportServiceHandler(service.Server(), new(OpmProductTestReportHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpmProductTestReportServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOpmVoiceServiceHandler(service.Server(), new(OpmVoiceHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpmVoiceServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterOpmVoiceProductServiceHandler(service.Server(), new(OpmVoiceProductHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpmVoiceProductServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterOpmVoiceProductMapServiceHandler(service.Server(), new(OpmVoiceProductMapHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpmVoiceProductMapServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterOpmVoicePublishRecordServiceHandler(service.Server(), new(OpmVoicePublishRecordHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpmVoicePublishRecordServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOpmProductFirmwareRelationServiceHandler(service.Server(), new(OpmProductFirmwareRelationHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpmProductFirmwareRelationServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOpmProductMaterialsServiceHandler(service.Server(), new(OpmProductMaterialsHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpmProductMaterialsServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOpmProductMaterialRelationServiceHandler(service.Server(), new(OpmProductMaterialRelationHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpmProductMaterialRelationServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterPmModuleFirmwareVersionServiceHandler(service.Server(), new(PmModuleFirmwareVersionHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterPmModuleFirmwareVersionServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOpmProductMaterialTypeServiceHandler(service.Server(), new(OpmProductMaterialTypeHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpmProductMaterialTypeServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOpmProductManualServiceHandler(service.Server(), new(OpmProductManualHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpmProductManualServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOpmProductMaterialLanguageServiceHandler(service.Server(), new(OpmProductMaterialLanguageHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpmProductMaterialLanguageServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOpmProductMaterialTypeLanguageServiceHandler(service.Server(), new(OpmProductMaterialTypeLanguageHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpmProductMaterialTypeLanguageServiceHandler 错误:%s", err.Error())
		return err
	}

	//err = protosService.RegisterOpmVoiceProductMapGoogleServiceHandler(service.Server(), new(OpmVoiceProductMapGoogleHandler))
	//if err != nil {
	//	iotlogger.LogHelper.Errorf("RegisterOpmVoiceProductMapGoogleServiceHandler 错误:%s", err.Error())
	//	return err
	//}

	err = protosService.RegisterOpmPanelServiceHandler(service.Server(), new(OpmPanelHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpmPanelServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOpmPanelStudioServiceHandler(service.Server(), new(OpmPanelStudioHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpmPanelStudioServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOpmPanelStudioBuildRecordServiceHandler(service.Server(), new(OpmPanelStudioBuildRecordHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpmPanelStudioBuildRecordServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOpmPanelAuthRelationServiceHandler(service.Server(), new(OpmPanelAuthRelationHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpmPanelAuthRelationServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOpmCommunityProductServiceHandler(service.Server(), new(OpmCommunityProductHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpmCommunityProductServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOpmThingModelRuleServiceHandler(service.Server(), new(OpmThingModelRuleHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpmThingModelRuleServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOpmPanelFontAssetServiceHandler(service.Server(), new(OpmPanelFontAssetHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpmPanelFontAssetServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterOpmPanelFontConfigServiceHandler(service.Server(), new(OpmPanelFontConfigHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpmPanelFontConfigServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterOpmPanelImageAssetServiceHandler(service.Server(), new(OpmPanelImageAssetHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpmPanelImageAssetServiceHandler 错误:%s", err.Error())
	}

	err = protosService.RegisterOpmDocumentsServiceHandler(service.Server(), new(OpmDocumentsHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpmDocumentsServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterProductHelpConfServiceHandler(service.Server(), new(ProductHelpConfHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterProductHelpConfServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterProductHelpDocServiceHandler(service.Server(), new(ProductHelpDocHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterProductHelpConfServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterTplDocumentTemplateServiceHandler(service.Server(), new(TplDocumentTemplateHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterTplDocumentTemplateServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterOpmProductTestAccountServiceHandler(service.Server(), new(OpmProductTestAccountHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterOpmProductTestAccountServiceHandler 错误:%s", err.Error())
		return err
	}
	return nil
}
