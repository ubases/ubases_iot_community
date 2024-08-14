package rpc

import (
	"cloud_platform/iot_proto/protos/protosService"
)

// 全局服务客户端变量
// 系统管理
var (
	ClientSysRoleService       protosService.SysRoleService
	ClientSysRoleDeptService   protosService.SysRoleDeptService
	ClientSysUserPostService   protosService.SysUserPostService
	ClientSysCasbinRuleService protosService.SysCasbinRuleService
	ClientCasbinExtService     protosService.CasbinExtService

	ClientSysUserService     protosService.SysUserService
	ClientSysAuthRuleService protosService.SysAuthRuleService
	// 部门管理
	ClientSysDeptService protosService.SysDeptService
	ClientSysPostService protosService.SysPostService
	//ClientCasbinRuleService  proto.CasbinRuleService
	ClientSysDictDataService       protosService.SysDictDataService
	ClientSysDictTypeService       protosService.SysDictTypeService
	ClientSysJobService            protosService.SysJobService
	ClientSysLoginLogService       protosService.SysLoginLogService
	ClientSysAppDocDirService      protosService.SysAppDocDirService
	ClientSysAppEntryService       protosService.SysAppEntryService
	ClientSysAppEntrySetingService protosService.SysAppEntrySetingService
	ClientSysAppHelpCenterService  protosService.SysAppHelpCenterService

	IPService protosService.IPService
	// 开放平台企业信息
	//ClientCompanyService protosService.OpenCompanyService

	ClientSysAttachmentService   protosService.SysAttachmentService
	ClientSysRegionServerService protosService.SysRegionServerService
)

// 物联管理
var (
	ClientIotDeviceServer     protosService.IotDeviceTriadService
	ClientIotDeviceInfoServer protosService.IotDeviceInfoService
	ClientIotDeviceLogServer  protosService.IotDeviceLogService
	ClientIotDeviceFault      protosService.IotDeviceFaultService
	ClientIotDeviceHome       protosService.IotDeviceHomeService

	IotDeviceGroupService     protosService.IotDeviceGroupService
	IotDeviceGroupListService protosService.IotDeviceGroupListService

	IotDeviceSharedService       protosService.IotDeviceSharedService
	IotDeviceShareReceiveService protosService.IotDeviceShareReceiveService
)

// 基础管理
var (
	TConfigDictDataServerService          protosService.ConfigDictDataService
	TConfigTranslateServerService         protosService.ConfigTranslateService
	TConfigDictTypeServerService          protosService.ConfigDictTypeService
	TConfigTranslateLanguageServerService protosService.ConfigTranslateLanguageService
	ClientAreaService                     protosService.SysAreaService
	ClientConfigPlatformService           protosService.ConfigPlatformService
	SysRegionServerService                protosService.SysRegionServerService
)

// 认证管理
var (
	ClientCloudAuthService      protosService.CloudAuthService
	ClientOpenAuthService       protosService.OpenAuthService
	ClientSysUserOnlineService  protosService.SysUserOnlineService
	ClientOpenUserOnlineService protosService.OpenUserOnlineService
)

// 产品管理
var (
	ClientProductService              protosService.TPmProductService
	ClientProductTypeService          protosService.TPmProductTypeService
	ClientThingModelService           protosService.TPmThingModelService
	ClientThingModelPropertiesService protosService.TPmThingModelPropertiesService
	ClientThingModelServicesService   protosService.TPmThingModelServicesService
	ClientThingModelEventsService     protosService.TPmThingModelEventsService
	ClientModuleService               protosService.PmModuleService // 芯片模块
	ClientPmFirmwareService           protosService.PmFirmwareService
	ClientPmFirmwareVersionService    protosService.PmFirmwareVersionService
	ClientFirmwareSettingService      protosService.PmFirmwareSettingService
	ClientNetworkGuideService         protosService.PmNetworkGuideService
	ClientNetworkGuideStepService     protosService.PmNetworkGuideStepService
	//ClientControlPanelService               protosService.PmControlPanelService
	ClientControlPanelsService              protosService.PmControlPanelsService
	ClientProductModuleRelationService      protosService.PmProductModuleRelationService
	ClientProductPanelRelationService       protosService.PmProductPanelRelationService
	ClientOpmProductService                 protosService.OpmProductService
	ClientOpmThingModelService              protosService.OpmThingModelService
	ClientOpmProductModuleRelationService   protosService.OpmProductModuleRelationService
	ClientOpmProductFirmwareRelationService protosService.OpmProductFirmwareRelationService
	ClientOpmProductPanelRelationService    protosService.OpmProductPanelRelationService
	ClientOpmNetworkGuideService            protosService.OpmNetworkGuideService
	ClientFirmwareService                   protosService.OpmFirmwareService
	ClientFirmwareVersionService            protosService.OpmFirmwareVersionService
	ClientOpmProductTestReportService       protosService.OpmProductTestReportService

	//测试用例模板
	ClientTestcaseTemplateService protosService.TplTestcaseTemplateService
	ClientNoticeTemplateService   protosService.MsNoticeTemplateService
	ClientMessageTemplateService  protosService.MpMessageTemplateService

	//消息服务
	ClientMessageService        protosService.MpMessageService
	ClientMsNoticeRecordService protosService.MsNoticeRecordService

	//OTA
	ClientOtaPkgService           protosService.OpmOtaPkgService
	ClientOtaPublishService       protosService.OpmOtaPublishService
	ClientOtaPublishLogService    protosService.OpmOtaPublishLogService
	ClientOtaUpgradeRecordService protosService.IotOtaUpgradeRecordService

	ClientPmModuleFirmwareVersionService protosService.PmModuleFirmwareVersionService

	ClientOpmProductVoiceService       protosService.OpmVoiceProductService
	ClientOpmVoiceProductMapService    protosService.OpmVoiceProductMapService
	ClientOpmVoiceService              protosService.OpmVoiceService
	ClientOpmVoicePublishRecordService protosService.OpmVoicePublishRecordService

	ClientOpmProductMaterialsService            protosService.OpmProductMaterialsService
	ClientOpmProductMaterialRelationService     protosService.OpmProductMaterialRelationService
	ClientOpmProductMaterialTypeService         protosService.OpmProductMaterialTypeService
	ClientOpmProductMaterialLanguageService     protosService.OpmProductMaterialLanguageService
	ClientOpmProductMaterialTypeLanguageService protosService.OpmProductMaterialTypeLanguageService

	// ?
	ClientOpmProductManualService protosService.OpmProductManualService

	//面板设计器
	ClientOpmPanelService             protosService.OpmPanelService
	ClientOpmPanelStudioService       protosService.OpmPanelStudioService
	ClientOpmPanelStudioBuildServicie protosService.OpmPanelStudioBuildRecordService
	//社区产品
	ClientOpmCommunityProductService protosService.OpmCommunityProductService
	//产品物模型规则设置
	ClientOpmThingModelRuleService protosService.OpmThingModelRuleService

	//图片、图标、字体资源管理
	ClientOpmPanelImageAssetService protosService.OpmPanelImageAssetService
	ClientOpmPanelFontAssetService  protosService.OpmPanelFontAssetService
	ClientOpmPanelFontConfigService protosService.OpmPanelFontConfigService
	ClientOpmDocumentsService       protosService.OpmDocumentsService

	ClientProductTestAccountService protosService.OpmProductTestAccountService

	ClientProductAppRelationService protosService.OpmProductAppRelationService
)

// 文档模板
var (
	ClientDocumentTemplateService protosService.TplDocumentTemplateService
	//ClientIntroduceService        protosService.CmsIntroduceService
	ClientProductHelpConfService protosService.ProductHelpConfService
	ClientProductHelpDocService  protosService.ProductHelpDocService
)

// 开放平台
var (
	ClientOpenUserService            protosService.OpenUserService
	ClientOpenUserCompanyService     protosService.OpenUserCompanyService
	ClientOpenCompanyService         protosService.OpenCompanyService
	ClientOpenCompanyConnectService  protosService.OpenCompanyConnectService
	ClientOpenRoleService            protosService.OpenRoleService
	ClientOpenAuthRuleService        protosService.OpenAuthRuleService
	UcUserService                    protosService.UcUserService
	UcHomeService                    protosService.UcHomeService
	UcHomeUserService                protosService.UcHomeUserService
	IotDeviceHomeService             protosService.IotDeviceHomeService
	ClientOpenCompanyAuthLogsService protosService.OpenCompanyAuthLogsService
	ClientDeveloperService           protosService.DeveloperService
	ClientAuthQuantityService        protosService.OpenAuthQuantityService
	UcFeedbackService                protosService.UcUserFeedbackService
	UcFeedbackResultService          protosService.UcUserFeedbackResultService
	ClientUcUserPrizeCollectService  protosService.UcUserPrizeCollectService

	ClientUcUserThirdService protosService.UcUserThirdService
)

// oem app
var (
	ClientOemAppDefMenuService            protosService.OemAppDefMenuService
	ClientOemAppService                   protosService.OemAppService
	ClientOemAppUiConfigService           protosService.OemAppUiConfigService
	ClientOemAppIntroduceService          protosService.OemAppIntroduceService
	ClientOemAppFunctionConfigService     protosService.OemAppFunctionConfigService
	ClientOemAppIosCertService            protosService.OemAppIosCertService
	ClientOemAppAndroidCertService        protosService.OemAppAndroidCertService
	ClientOemAppPushCertService           protosService.OemAppPushCertService
	ClientOemAppBuildRecordService        protosService.OemAppBuildRecordService
	ClientOemAppDocService                protosService.OemAppDocService
	ClientOemAppDocDirService             protosService.OemAppDocDirService
	ClientOemAppDocRelationService        protosService.OemAppDocRelationService
	ClientOemAppEntryService              protosService.OemAppEntryService
	ClientOemAppEntrySetingService        protosService.OemAppEntrySetingService
	ClientOemAppFlashScreenService        protosService.OemAppFlashScreenService
	ClientOemAppFlashScreenUserService    protosService.OemAppFlashScreenUserService
	ClientOemAppVersionRecordService      protosService.OemAppVersionRecordService
	ClientOemAppBasicUiSettingService     protosService.OemAppBasicUiSettingService
	ClientOemAppAssistReleaseService      protosService.OemAppAssistReleaseService
	ClientOemAppTemplateService           protosService.OemAppTemplateService
	ClientOemAppTemplateFunctionService   protosService.OemAppTemplateFunctionService
	ClientOemAppTemplateMenuService       protosService.OemAppTemplateMenuService
	ClientOemAppTemplateSkinService       protosService.OemAppTemplateSkinService
	ClientOemAppTemplateThirdPartyService protosService.OemAppTemplateThirdPartyService
	ClientOemAppTemplateUiService         protosService.OemAppTemplateUiService
	ClientPublicAppVersionService         protosService.PublicAppVersionService
	ClientFeedbackTypeService             protosService.OemFeedbackTypeService
	ClientOemAppCustomRecordService       protosService.OemAppCustomRecordService
	ClientOemAppDebuggerService           protosService.OemAppDebuggerService
	ClientCloudAppBuildAuthService        protosService.CloudAppBuildAuthService
)

// 联系我们
var (
	ClientContractUsService protosService.ContractUsService
	ClientEmailService      protosService.EmailService
	ClientSmsService        protosService.SmsService
)

// 翻译
var (
	ClientLangTranslateService        protosService.LangTranslateService
	ClientLangResourcesService        protosService.LangResourcesService
	ClientLangCustomResourceService   protosService.LangCustomResourcesService
	ClientLangResourcesPackageService protosService.LangResourcePackageService
	ClientLangTranslateTypeService    protosService.LangTranslateTypeService
)

// 数据统计
var (
	ClientDataOverviewHourService  protosService.DataOverviewHourService
	ClientDataOverviewMonthService protosService.DataOverviewMonthService
	ClientProductFaultMonthService protosService.ProductFaultMonthService
	ClientProductFaultTypeService  protosService.ProductFaultTypeService
	ClientDeviceActiveService      protosService.DeviceActiveService
	ClientStatisticsService        protosService.StatisticsService
	ClientAppUserStatisticsService protosService.AppUserStatisticsService
	ClientAppDataServiceService    protosService.PmAppDataService
)

// 日志管理
var (
	ClientAppLogService protosService.AppLogService
)

// 语控管理
var (
	ClientVoiceService protosService.VoiceService
)

var (
	ClientSceneTemplateService protosService.SceneTemplateService
)

var (
	ClientPanelGenerateService protosService.PanelGenerateService
)
