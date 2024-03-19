package rpc

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_proto/protos/protosService"

	"go-micro.dev/v4/client"

	"time"

	"github.com/asim/go-micro/plugins/client/grpc/v4"
	wrapper_uber_ratelimit "github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v4"
	"github.com/asim/go-micro/plugins/wrapper/select/roundrobin/v4"
	wrapper_opentracing "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v4"
	"github.com/opentracing/opentracing-go"
)

// 初始化其它服务的客户端
func InitServiceClient() {
	RegisterGrpcServiceClient(iotconst.IOT_SYSTEM_SERVICE, func(name string, client client.Client) {
		ClientSysRoleService = protosService.NewSysRoleService(name, client)
		ClientSysRoleDeptService = protosService.NewSysRoleDeptService(name, client)
		ClientSysUserService = protosService.NewSysUserService(name, client)
		ClientSysUserOnlineService = protosService.NewSysUserOnlineService(name, client)
		ClientSysUserPostService = protosService.NewSysUserPostService(name, client)
		ClientSysCasbinRuleService = protosService.NewSysCasbinRuleService(name, client)
		ClientCasbinExtService = protosService.NewCasbinExtService(name, client)
		ClientSysDeptService = protosService.NewSysDeptService(name, client)
		ClientSysDictDataService = protosService.NewSysDictDataService(name, client)
		ClientSysDictTypeService = protosService.NewSysDictTypeService(name, client)
		ClientSysJobService = protosService.NewSysJobService(name, client)
		ClientSysAuthRuleService = protosService.NewSysAuthRuleService(name, client)
		ClientSysLoginLogService = protosService.NewSysLoginLogService(name, client)
		ClientSysPostService = protosService.NewSysPostService(name, client)
		ClientSysAppDocDirService = protosService.NewSysAppDocDirService(name, client)
		ClientSysAppEntryService = protosService.NewSysAppEntryService(name, client)
		ClientSysAppEntrySetingService = protosService.NewSysAppEntrySetingService(name, client)
		ClientSysAppHelpCenterService = protosService.NewSysAppHelpCenterService(name, client)

		ClientCloudAuthService = protosService.NewCloudAuthService(name, client)
	})

	RegisterGrpcServiceClient(iotconst.IOT_DEVICE_SERVICE, func(name string, client client.Client) {
		ClientIotDeviceServer = protosService.NewIotDeviceTriadService(name, client)
		ClientIotDeviceInfoServer = protosService.NewIotDeviceInfoService(name, client)
		ClientIotDeviceLogServer = protosService.NewIotDeviceLogService(name, client)
		IotDeviceHomeService = protosService.NewIotDeviceHomeService(name, client)
		ClientIotDeviceFault = protosService.NewIotDeviceFaultService(name, client)
		ClientOtaUpgradeRecordService = protosService.NewIotOtaUpgradeRecordService(name, client)
		ClientIotDeviceHome = protosService.NewIotDeviceHomeService(name, client)

		IotDeviceSharedService = protosService.NewIotDeviceSharedService(name, client)
		IotDeviceShareReceiveService = protosService.NewIotDeviceShareReceiveService(name, client)
		IotDeviceGroupService = protosService.NewIotDeviceGroupService(name, client)
		IotDeviceGroupListService = protosService.NewIotDeviceGroupListService(name, client)
	})

	RegisterGrpcServiceClient(iotconst.IOT_BASIC_SERVICE, func(name string, client client.Client) {
		TConfigDictDataServerService = protosService.NewConfigDictDataService(name, client)
		TConfigDictTypeServerService = protosService.NewConfigDictTypeService(name, client)
		TConfigTranslateServerService = protosService.NewConfigTranslateService(name, client)
		TConfigTranslateLanguageServerService = protosService.NewConfigTranslateLanguageService(name, client)

		ClientAreaService = protosService.NewSysAreaService(name, client)
		ClientConfigPlatformService = protosService.NewConfigPlatformService(name, client)
		SysRegionServerService = protosService.NewSysRegionServerService(name, client)

		ClientLangTranslateService = protosService.NewLangTranslateService(name, client)
		ClientLangTranslateTypeService = protosService.NewLangTranslateTypeService(name, client)
		ClientLangResourcesPackageService = protosService.NewLangResourcePackageService(name, client)
		ClientLangResourcesService = protosService.NewLangResourcesService(name, client)
		ClientLangCustomResourceService = protosService.NewLangCustomResourcesService(name, client)
	})
	//产品管理
	RegisterGrpcServiceClient(iotconst.IOT_PRODUCT_SERVICE, func(name string, client client.Client) {
		ClientProductService = protosService.NewTPmProductService(name, client)
		ClientProductTypeService = protosService.NewTPmProductTypeService(name, client)
		ClientThingModelService = protosService.NewTPmThingModelService(name, client)
		ClientThingModelPropertiesService = protosService.NewTPmThingModelPropertiesService(name, client)
		ClientThingModelServicesService = protosService.NewTPmThingModelServicesService(name, client)
		ClientThingModelEventsService = protosService.NewTPmThingModelEventsService(name, client)
		ClientModuleService = protosService.NewPmModuleService(name, client)
		ClientFirmwareSettingService = protosService.NewPmFirmwareSettingService(name, client)
		ClientPmFirmwareService = protosService.NewPmFirmwareService(name, client)
		ClientPmFirmwareVersionService = protosService.NewPmFirmwareVersionService(name, client)
		ClientNetworkGuideService = protosService.NewPmNetworkGuideService(name, client)
		ClientNetworkGuideStepService = protosService.NewPmNetworkGuideStepService(name, client)
		ClientControlPanelsService = protosService.NewPmControlPanelsService(name, client)
		ClientTestcaseTemplateService = protosService.NewTplTestcaseTemplateService(name, client)

		//开放平台固件管理
		ClientFirmwareService = protosService.NewOpmFirmwareService(name, client)
		ClientFirmwareVersionService = protosService.NewOpmFirmwareVersionService(name, client)

		//开放平台产品信息
		ClientOpmProductService = protosService.NewOpmProductService(name, client)
		ClientOpmThingModelService = protosService.NewOpmThingModelService(name, client)
		ClientOpmProductModuleRelationService = protosService.NewOpmProductModuleRelationService(name, client)
		ClientOpmProductFirmwareRelationService = protosService.NewOpmProductFirmwareRelationService(name, client)
		ClientOpmProductPanelRelationService = protosService.NewOpmProductPanelRelationService(name, client)
		ClientOpmNetworkGuideService = protosService.NewOpmNetworkGuideService(name, client)
		ClientOpmProductTestReportService = protosService.NewOpmProductTestReportService(name, client)

		//OTA Pkg
		ClientOtaPkgService = protosService.NewOpmOtaPkgService(name, client)
		ClientOtaPublishService = protosService.NewOpmOtaPublishService(name, client)
		ClientOtaPublishLogService = protosService.NewOpmOtaPublishLogService(name, client)

		ClientPmModuleFirmwareVersionService = protosService.NewPmModuleFirmwareVersionService(name, client)

		ClientOpmProductVoiceService = protosService.NewOpmVoiceProductService(name, client)
		ClientOpmVoiceProductMapService = protosService.NewOpmVoiceProductMapService(name, client)
		ClientOpmVoiceService = protosService.NewOpmVoiceService(name, client)
		ClientOpmVoicePublishRecordService = protosService.NewOpmVoicePublishRecordService(name, client)

		ClientOpmProductMaterialsService = protosService.NewOpmProductMaterialsService(name, client)
		ClientOpmProductMaterialRelationService = protosService.NewOpmProductMaterialRelationService(name, client)
		ClientOpmProductMaterialTypeService = protosService.NewOpmProductMaterialTypeService(name, client)
		ClientOpmProductManualService = protosService.NewOpmProductManualService(name, client)
		ClientOpmProductMaterialLanguageService = protosService.NewOpmProductMaterialLanguageService(name, client)
		ClientOpmProductMaterialTypeLanguageService = protosService.NewOpmProductMaterialTypeLanguageService(name, client)

		ClientOpmPanelService = protosService.NewOpmPanelService(name, client)
		ClientOpmPanelStudioService = protosService.NewOpmPanelStudioService(name, client)
		ClientOpmPanelStudioBuildServicie = protosService.NewOpmPanelStudioBuildRecordService(name, client)

		ClientOpmCommunityProductService = protosService.NewOpmCommunityProductService(name, client)
		ClientOpmThingModelRuleService = protosService.NewOpmThingModelRuleService(name, client)

		ClientOpmPanelImageAssetService = protosService.NewOpmPanelImageAssetService(name, client)
		ClientOpmPanelFontAssetService = protosService.NewOpmPanelFontAssetService(name, client)
		ClientOpmPanelFontConfigService = protosService.NewOpmPanelFontConfigService(name, client)

		ClientOpmDocumentsService = protosService.NewOpmDocumentsService(name, client)

		ClientDocumentTemplateService = protosService.NewTplDocumentTemplateService(name, client)
		ClientProductHelpConfService = protosService.NewProductHelpConfService(name, client)
		ClientProductHelpDocService = protosService.NewProductHelpDocService(name, client)

		ClientProductTestAccountService = protosService.NewOpmProductTestAccountService(name, client)
	})
	//开放平台
	RegisterGrpcServiceClient(iotconst.IOT_OPEN_SYSTEM_SERVICE, func(name string, client client.Client) {
		ClientOpenUserService = protosService.NewOpenUserService(name, client)
		ClientOpenUserCompanyService = protosService.NewOpenUserCompanyService(name, client)
		ClientOpenCompanyService = protosService.NewOpenCompanyService(name, client)
		ClientOpenCompanyConnectService = protosService.NewOpenCompanyConnectService(name, client)
		ClientOpenRoleService = protosService.NewOpenRoleService(name, client)
		ClientOpenAuthRuleService = protosService.NewOpenAuthRuleService(name, client)
		ClientOpenCompanyAuthLogsService = protosService.NewOpenCompanyAuthLogsService(name, client)
		ClientDeveloperService = protosService.NewDeveloperService(name, client)
		ClientAuthQuantityService = protosService.NewOpenAuthQuantityService(name, client)

		ClientOpenUserOnlineService = protosService.NewOpenUserOnlineService(name, client)
		ClientOpenAuthService = protosService.NewOpenAuthService(name, client)
	})

	//APP用户
	RegisterGrpcServiceClient(iotconst.IOT_APP_USER_SERVICE, func(name string, client client.Client) {
		UcUserService = protosService.NewUcUserService(name, client)
		UcHomeService = protosService.NewUcHomeService(name, client)
		UcHomeUserService = protosService.NewUcHomeUserService(name, client)
		UcFeedbackService = protosService.NewUcUserFeedbackService(name, client)
		UcFeedbackResultService = protosService.NewUcUserFeedbackResultService(name, client)
		ClientUcUserPrizeCollectService = protosService.NewUcUserPrizeCollectService(name, client)
		ClientUcUserThirdService = protosService.NewUcUserThirdService(name, client)
	})

	//OEM APP
	RegisterGrpcServiceClient(iotconst.IOT_APP_OEM_SERVICE, func(name string, client client.Client) {
		ClientOemAppDefMenuService = protosService.NewOemAppDefMenuService(name, client)
		ClientOemAppService = protosService.NewOemAppService(name, client)
		ClientOemAppUiConfigService = protosService.NewOemAppUiConfigService(name, client)
		ClientOemAppFunctionConfigService = protosService.NewOemAppFunctionConfigService(name, client)

		ClientOemAppIntroduceService = protosService.NewOemAppIntroduceService(name, client)
		ClientOemAppIosCertService = protosService.NewOemAppIosCertService(name, client)
		ClientOemAppAndroidCertService = protosService.NewOemAppAndroidCertService(name, client)
		ClientOemAppPushCertService = protosService.NewOemAppPushCertService(name, client)
		ClientOemAppBuildRecordService = protosService.NewOemAppBuildRecordService(name, client)

		ClientOemAppDocService = protosService.NewOemAppDocService(name, client)
		ClientOemAppDocDirService = protosService.NewOemAppDocDirService(name, client)
		ClientOemAppDocRelationService = protosService.NewOemAppDocRelationService(name, client)
		ClientOemAppEntryService = protosService.NewOemAppEntryService(name, client)
		ClientOemAppEntrySetingService = protosService.NewOemAppEntrySetingService(name, client)
		ClientOemAppFlashScreenService = protosService.NewOemAppFlashScreenService(name, client)
		ClientOemAppFlashScreenUserService = protosService.NewOemAppFlashScreenUserService(name, client)
		ClientOemAppVersionRecordService = protosService.NewOemAppVersionRecordService(name, client)

		ClientOemAppBasicUiSettingService = protosService.NewOemAppBasicUiSettingService(name, client)
		ClientOemAppAssistReleaseService = protosService.NewOemAppAssistReleaseService(name, client)
		ClientOemAppTemplateService = protosService.NewOemAppTemplateService(name, client)
		ClientOemAppTemplateFunctionService = protosService.NewOemAppTemplateFunctionService(name, client)
		ClientOemAppTemplateMenuService = protosService.NewOemAppTemplateMenuService(name, client)
		ClientOemAppTemplateSkinService = protosService.NewOemAppTemplateSkinService(name, client)
		ClientOemAppTemplateThirdPartyService = protosService.NewOemAppTemplateThirdPartyService(name, client)
		ClientOemAppTemplateUiService = protosService.NewOemAppTemplateUiService(name, client)
		ClientPublicAppVersionService = protosService.NewPublicAppVersionService(name, client)
		ClientFeedbackTypeService = protosService.NewOemFeedbackTypeService(name, client)
		ClientOemAppCustomRecordService = protosService.NewOemAppCustomRecordService(name, client)
		ClientOemAppDebuggerService = protosService.NewOemAppDebuggerService(name, client)
	})

	//OEM APP
	RegisterGrpcServiceClient(iotconst.IOT_MESSAGE_SERVICE, func(name string, client client.Client) {
		ClientMessageTemplateService = protosService.NewMpMessageTemplateService(name, client)
		ClientMessageService = protosService.NewMpMessageService(name, client)

		ClientContractUsService = protosService.NewContractUsService(name, client)
		ClientNoticeTemplateService = protosService.NewMsNoticeTemplateService(name, client)
		ClientEmailService = protosService.NewEmailService(name, client)
		ClientSmsService = protosService.NewSmsService(name, client)
	})

	RegisterGrpcServiceClient(iotconst.IOT_STATISTICS_SERVICE, func(name string, client client.Client) {
		ClientDataOverviewHourService = protosService.NewDataOverviewHourService(name, client)
		ClientDataOverviewMonthService = protosService.NewDataOverviewMonthService(name, client)
		ClientProductFaultMonthService = protosService.NewProductFaultMonthService(name, client)
		ClientProductFaultTypeService = protosService.NewProductFaultTypeService(name, client)
		ClientDeviceActiveService = protosService.NewDeviceActiveService(name, client)
		ClientStatisticsService = protosService.NewStatisticsService(name, client)
		ClientAppUserStatisticsService = protosService.NewAppUserStatisticsService(name, client)
		ClientAppDataServiceService = protosService.NewPmAppDataService(name, client)
	})

	// 日志管理
	RegisterGrpcServiceClient(iotconst.IOT_LOG_SERVICE, func(name string, client client.Client) {
		ClientAppLogService = protosService.NewAppLogService(name, client)
	})

	// 语控管理
	RegisterGrpcServiceClient(iotconst.IOT_SMART_SPEAKER_SERVICE, func(name string, client client.Client) {
		ClientVoiceService = protosService.NewVoiceService(name, client)
	})

	// 智能场景
	RegisterGrpcServiceClient(iotconst.IOT_INTELLIGENCE_SERVICE, func(name string, client client.Client) {
		ClientSceneTemplateService = protosService.NewSceneTemplateService(name, client)
	})

	//天气微服务
	RegisterGrpcServiceClient(iotconst.IOT_WEATHER_SERVICE, func(name string, client client.Client) {
		IPService = protosService.NewIPService(name, client)
	})

	RegisterGrpcServiceClient(iotconst.IOT_PANEL_DESIGN_SERVICE, func(name string, client client.Client) {
		ClientPanelGenerateService = protosService.NewPanelGenerateService(name, client)
	})

}

func RegisterGrpcServiceClient(name string, fun func(name string, client client.Client)) {
	cli := grpc.NewClient(
		client.Registry(getRegistry()),
		client.Wrap(roundrobin.NewClientWrapper()),
		//client.Wrap(wrapper_gobreaker.NewCustomClientWrapper(gobreaker.Settings{Name: name, MaxRequests: 10, Interval: 30 * time.Second, Timeout: 30 * time.Second}, wrapper_gobreaker.BreakServiceEndpoint)),
		client.Wrap(wrapper_uber_ratelimit.NewClientWrapper(1000)),
		client.Wrap(wrapper_opentracing.NewClientWrapper(opentracing.GlobalTracer())),
		client.Retries(0),
		client.RequestTimeout(120*time.Second), //for debug
		client.DialTimeout(6*time.Second),
	)
	fun(name, cli)
}
