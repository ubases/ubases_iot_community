package rpc

import (
	"cloud_platform/iot_proto/protos/protosService"
)

var (
	AppAuthService                     protosService.AppAuthService
	TUcUserService                     protosService.UcUserService
	ClientUcUserThirdService           protosService.UcUserThirdService
	EmailService                       protosService.EmailService
	SmsService                         protosService.SmsService
	UcHomeService                      protosService.UcHomeService
	UcHomeUserService                  protosService.UcHomeUserService
	UcHomeRoomService                  protosService.UcHomeRoomService
	UcFeedbackService                  protosService.UcUserFeedbackService
	IotDeviceHomeService               protosService.IotDeviceHomeService
	IotDeviceInfoService               protosService.IotDeviceInfoService
	IotDeviceTimerService              protosService.IotDeviceTimerService
	IotDeviceCountdownService          protosService.IotDeviceCountdownService
	IotDeviceSharedService             protosService.IotDeviceSharedService
	IotDeviceShareReceiveService       protosService.IotDeviceShareReceiveService
	MessageService                     protosService.MpMessageService
	SceneIntelligenceService           protosService.SceneIntelligenceService
	SceneIntelligenceConditionService  protosService.SceneIntelligenceConditionService
	SceneIntelligenceTaskService       protosService.SceneIntelligenceTaskService
	SceneIntelligenceResultTaskService protosService.SceneIntelligenceResultTaskService
	SceneIntelligenceResultService     protosService.SceneIntelligenceResultService
	SceneIntelligenceLogService        protosService.SceneIntelligenceLogService
	ConfigDictDataService              protosService.ConfigDictDataService
	LangCustomResourcesService         protosService.LangCustomResourcesService
	SysRegionServerService             protosService.SysRegionServerService

	ClientLangCustomResourceService protosService.LangCustomResourcesService
	ClientLangResourcesService      protosService.LangResourcesService

	IotDeviceGroupService     protosService.IotDeviceGroupService
	IotDeviceGroupListService protosService.IotDeviceGroupListService

	IotDeviceLogService             protosService.IotDeviceLogService
	UcAppleidInfoService            protosService.UcAppleidInfoService
	ClientUcUserPrizeCollectService protosService.UcUserPrizeCollectService
	ClientAreaService               protosService.SysAreaService

	ClientSceneTemplateService protosService.SceneTemplateService
)

var (
	MessageRedDotService    protosService.MpMessageRedDotService
	MessageUserOutService   protosService.MpMessageUserOutService
	MessageUserInService    protosService.MpMessageUserInService
	AppPushTokenService     protosService.AppPushTokenService
	AppPushTokenUserService protosService.AppPushTokenUserService
)

// 产品信息相关服务处理器
var (
	ProductTypeService                      protosService.TPmProductTypeService
	ProductBaseService                      protosService.TPmProductService
	ProductService                          protosService.OpmProductService
	ProductThingsModelService               protosService.OpmThingModelService
	ClientOpmNetworkGuideService            protosService.OpmNetworkGuideService
	ClientOtaPublishService                 protosService.OpmOtaPublishService
	ClientOpmThingModelService              protosService.OpmThingModelService
	ClientOemAppIntroduce                   protosService.OemAppIntroduceService
	ClientOemAppService                     protosService.OemAppService
	ClientOemAppCustomRecordService         protosService.OemAppCustomRecordService
	ClientOemAppFunctionConfigService       protosService.OemAppFunctionConfigService
	ClientOemAppUiConfigService             protosService.OemAppUiConfigService
	ClientOpmVoiceService                   protosService.OpmVoiceService
	ClientOemAppFlashScreenService          protosService.OemAppFlashScreenService
	ClientOemAppFlashScreenUserService      protosService.OemAppFlashScreenUserService
	ClientOpmProductMaterialsService        protosService.OpmProductMaterialsService
	ClientOpmProductMaterialRelationService protosService.OpmProductMaterialRelationService

	ClientOemAppAssistReleaseService protosService.OemAppAssistReleaseService

	ClientOemFeedbackTypeService protosService.OemFeedbackTypeService

	ClientOpmProductManualService    protosService.OpmProductManualService
	ClientOpmCommunityProductService protosService.OpmCommunityProductService
	ClientOpmThingModelRuleService   protosService.OpmThingModelRuleService
	ClientOpmDocumentsService        protosService.OpmDocumentsService
)

// 天气
var (
	IPService protosService.IPService

	WeatherService protosService.WeatherService

	ClientMqttService protosService.MqttService
)

var (
	ClientOemAppIntroduceService protosService.OemAppIntroduceService
	ProductHelpDocService        protosService.ProductHelpDocService
)

var (
	ClientOemAppDocDirService      protosService.OemAppDocDirService
	ClientOemAppDocRelationService protosService.OemAppDocRelationService
	ClientOemAppEntryService       protosService.OemAppEntryService
	ClientLangTranslateService     protosService.LangTranslateService
	ClientSysAppDocDirService      protosService.SysAppDocDirService
	ClientSysAppEntryService       protosService.SysAppEntryService

	ClientOemAppDebuggerService protosService.OemAppDebuggerService
	ClientSysAttachmentService  protosService.SysAttachmentService
)

var (
	ClientOpenCompany protosService.OpenCompanyService //开发者公司信息
)
