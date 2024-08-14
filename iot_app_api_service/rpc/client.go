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
	RegisterGrpcServiceClient(iotconst.IOT_APP_USER_SERVICE, func(name string, client client.Client) {
		//BaseDataServerService = protosService.NewBaseDataService(name, client)
		TUcUserService = protosService.NewUcUserService(name, client)
		UcHomeService = protosService.NewUcHomeService(name, client)
		UcHomeUserService = protosService.NewUcHomeUserService(name, client)
		UcHomeRoomService = protosService.NewUcHomeRoomService(name, client)
		UcFeedbackService = protosService.NewUcUserFeedbackService(name, client)
		ClientUcUserThirdService = protosService.NewUcUserThirdService(name, client)
		UcAppleidInfoService = protosService.NewUcAppleidInfoService(name, client)
		ClientUcUserPrizeCollectService = protosService.NewUcUserPrizeCollectService(name, client)
		AppAuthService = protosService.NewAppAuthService(name, client)
		//AppUpgradeService = protosService.NewAppUpgradeService(name, client)
	})
	RegisterGrpcServiceClient(iotconst.IOT_DEVICE_SERVICE, func(name string, client client.Client) {
		IotDeviceHomeService = protosService.NewIotDeviceHomeService(name, client)
		IotDeviceInfoService = protosService.NewIotDeviceInfoService(name, client)
		IotDeviceTimerService = protosService.NewIotDeviceTimerService(name, client)
		IotDeviceCountdownService = protosService.NewIotDeviceCountdownService(name, client)
		IotDeviceSharedService = protosService.NewIotDeviceSharedService(name, client)
		IotDeviceShareReceiveService = protosService.NewIotDeviceShareReceiveService(name, client)
		IotDeviceGroupService = protosService.NewIotDeviceGroupService(name, client)
		IotDeviceGroupListService = protosService.NewIotDeviceGroupListService(name, client)
		IotDeviceLogService = protosService.NewIotDeviceLogService(name, client)
	})

	//APP产品相关微服务
	RegisterGrpcServiceClient(iotconst.IOT_PRODUCT_SERVICE, func(name string, client client.Client) {
		ProductTypeService = protosService.NewTPmProductTypeService(name, client)
		ProductService = protosService.NewOpmProductService(name, client)
		ProductThingsModelService = protosService.NewOpmThingModelService(name, client)
		ClientOpmNetworkGuideService = protosService.NewOpmNetworkGuideService(name, client)
		ClientOtaPublishService = protosService.NewOpmOtaPublishService(name, client)
		ClientOpmThingModelService = protosService.NewOpmThingModelService(name, client)
		ProductBaseService = protosService.NewTPmProductService(name, client)
		ClientOpmVoiceService = protosService.NewOpmVoiceService(name, client)
		ClientOpmProductMaterialsService = protosService.NewOpmProductMaterialsService(name, client)
		ClientOpmProductMaterialRelationService = protosService.NewOpmProductMaterialRelationService(name, client)
		ClientOpmProductManualService = protosService.NewOpmProductManualService(name, client)
		ClientOpmCommunityProductService = protosService.NewOpmCommunityProductService(name, client)
		ClientOpmThingModelRuleService = protosService.NewOpmThingModelRuleService(name, client)
		ClientOpmDocumentsService = protosService.NewOpmDocumentsService(name, client)
		ProductHelpDocService = protosService.NewProductHelpDocService(name, client)
	})

	RegisterGrpcServiceClient(iotconst.IOT_MESSAGE_SERVICE, func(name string, client client.Client) {
		MessageUserOutService = protosService.NewMpMessageUserOutService(name, client)
		MessageUserInService = protosService.NewMpMessageUserInService(name, client)
		MessageRedDotService = protosService.NewMpMessageRedDotService(name, client)
		MessageService = protosService.NewMpMessageService(name, client)
		AppPushTokenService = protosService.NewAppPushTokenService(name, client)
		AppPushTokenUserService = protosService.NewAppPushTokenUserService(name, client)
		EmailService = protosService.NewEmailService(name, client)
		SmsService = protosService.NewSmsService(name, client)
	})

	//智能场景微服务
	RegisterGrpcServiceClient(iotconst.IOT_INTELLIGENCE_SERVICE, func(name string, client client.Client) {
		SceneIntelligenceService = protosService.NewSceneIntelligenceService(name, client)
		SceneIntelligenceConditionService = protosService.NewSceneIntelligenceConditionService(name, client)
		SceneIntelligenceTaskService = protosService.NewSceneIntelligenceTaskService(name, client)
		SceneIntelligenceResultService = protosService.NewSceneIntelligenceResultService(name, client)
		SceneIntelligenceResultTaskService = protosService.NewSceneIntelligenceResultTaskService(name, client)
		SceneIntelligenceLogService = protosService.NewSceneIntelligenceLogService(name, client)
		ClientSceneTemplateService = protosService.NewSceneTemplateService(name, client)
	})
	//字典微服务
	RegisterGrpcServiceClient(iotconst.IOT_BASIC_SERVICE, func(name string, client client.Client) {
		ConfigDictDataService = protosService.NewConfigDictDataService(name, client)
		SysRegionServerService = protosService.NewSysRegionServerService(name, client)
		ClientAreaService = protosService.NewSysAreaService(name, client)
		LangCustomResourcesService = protosService.NewLangCustomResourcesService(name, client)
		ClientLangCustomResourceService = protosService.NewLangCustomResourcesService(name, client)
		ClientLangResourcesService = protosService.NewLangResourcesService(name, client)
		ClientLangTranslateService = protosService.NewLangTranslateService(name, client)
	})
	//天气微服务
	RegisterGrpcServiceClient(iotconst.IOT_WEATHER_SERVICE, func(name string, client client.Client) {
		IPService = protosService.NewIPService(name, client)
		WeatherService = protosService.NewWeatherService(name, client)
	})

	//OEM APP
	RegisterGrpcServiceClient(iotconst.IOT_APP_OEM_SERVICE, func(name string, client client.Client) {
		ClientOemAppIntroduceService = protosService.NewOemAppIntroduceService(name, client)
		ClientOemAppService = protosService.NewOemAppService(name, client)
		ClientOemAppCustomRecordService = protosService.NewOemAppCustomRecordService(name, client)
		ClientOemAppDocDirService = protosService.NewOemAppDocDirService(name, client)
		ClientOemAppDocRelationService = protosService.NewOemAppDocRelationService(name, client)
		ClientOemAppEntryService = protosService.NewOemAppEntryService(name, client)
		ClientOemAppFunctionConfigService = protosService.NewOemAppFunctionConfigService(name, client)
		ClientOemAppUiConfigService = protosService.NewOemAppUiConfigService(name, client)
		ClientOemAppFlashScreenService = protosService.NewOemAppFlashScreenService(name, client)
		ClientOemAppFlashScreenUserService = protosService.NewOemAppFlashScreenUserService(name, client)
		//辅助上架
		ClientOemAppAssistReleaseService = protosService.NewOemAppAssistReleaseService(name, client)

		ClientOemFeedbackTypeService = protosService.NewOemFeedbackTypeService(name, client)
		ClientOemAppDebuggerService = protosService.NewOemAppDebuggerService(name, client)
	})

	//系统管理服务
	RegisterGrpcServiceClient(iotconst.IOT_SYSTEM_SERVICE, func(name string, client client.Client) {
		ClientSysAppDocDirService = protosService.NewSysAppDocDirService(name, client)
		ClientSysAppEntryService = protosService.NewSysAppEntryService(name, client)
		ClientSysAttachmentService = protosService.NewSysAttachmentService(name, client)
	})
	//开发平台系统服务
	RegisterGrpcServiceClient(iotconst.IOT_OPEN_SYSTEM_SERVICE, func(name string, client client.Client) {
		ClientOpenCompany = protosService.NewOpenCompanyService(name, client)
	})

	RegisterGrpcServiceClient(iotconst.IOT_MQTT_SERVICE, func(name string, client client.Client) {
		ClientMqttService = protosService.NewMqttService(name, client)
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
	_ = cli.Init(grpc.MaxSendMsgSize(20 * 1024 * 1024))
	_ = cli.Init(grpc.MaxRecvMsgSize(20 * 1024 * 1024))
	fun(name, cli)
}
