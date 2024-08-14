package rpcclient

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_proto/protos/protosService"
	"cloud_platform/iot_voice_service/rpc"
	"time"

	"github.com/asim/go-micro/plugins/client/grpc/v4"

	wrapper_uber_ratelimit "github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v4"
	"github.com/asim/go-micro/plugins/wrapper/select/roundrobin/v4"
	wrapper_opentracing "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v4"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4/client"
)

// 全局服务客户端变量
var (
	ClientMqttMessage                 protosService.MqttService
	ClientUcUserService               protosService.UcUserService
	ClientUcHomeRoomService           protosService.UcHomeRoomService
	ClientUcHomeUserService           protosService.UcHomeUserService
	ClientUcUserThirdService          protosService.UcUserThirdService
	AppAuthService                    protosService.AppAuthService
	ClientOemAppService               protosService.OemAppService
	ClientOemAppFunctionConfigService protosService.OemAppFunctionConfigService
	ClientIotDeviceHomeService        protosService.IotDeviceHomeService
	ClientIotDeviceInfoService        protosService.IotDeviceInfoService
	ClientIotDeviceTriadService       protosService.IotDeviceTriadService
	ClienOpmVoiceProductService       protosService.OpmVoiceProductService
	ClienOpmVoiceProductMapService    protosService.OpmVoiceProductMapService
	ClientOpmThingModelService        protosService.OpmThingModelService
	ClientOpmProductService           protosService.OpmProductService
	//ClientOpmVoiceProductMapGoogleService protosService.OpmVoiceProductMapGoogleService

	SysRegionServerService protosService.SysRegionServerService
	ConfigDictDataService  protosService.ConfigDictDataService
	IPService              protosService.IPService
)

// 初始化其它服务的客户端
func InitServiceClient() {
	RegisterGrpcServiceClient(iotconst.IOT_APP_USER_SERVICE, func(name string, client client.Client) {
		ClientUcUserService = protosService.NewUcUserService(name, client)
		ClientUcUserThirdService = protosService.NewUcUserThirdService(name, client)
		ClientUcHomeRoomService = protosService.NewUcHomeRoomService(name, client)
		ClientUcHomeUserService = protosService.NewUcHomeUserService(name, client)

		AppAuthService = protosService.NewAppAuthService(name, client)
	})

	//OEM APP
	RegisterGrpcServiceClient(iotconst.IOT_APP_OEM_SERVICE, func(name string, client client.Client) {
		ClientOemAppService = protosService.NewOemAppService(name, client)
		ClientOemAppFunctionConfigService = protosService.NewOemAppFunctionConfigService(name, client)
	})

	//Device
	RegisterGrpcServiceClient(iotconst.IOT_DEVICE_SERVICE, func(name string, client client.Client) {
		ClientIotDeviceHomeService = protosService.NewIotDeviceHomeService(name, client)
		ClientIotDeviceInfoService = protosService.NewIotDeviceInfoService(name, client)
		ClientIotDeviceTriadService = protosService.NewIotDeviceTriadService(name, client)
	})

	//Product
	RegisterGrpcServiceClient(iotconst.IOT_PRODUCT_SERVICE, func(name string, client client.Client) {
		ClienOpmVoiceProductService = protosService.NewOpmVoiceProductService(name, client)
		ClienOpmVoiceProductMapService = protosService.NewOpmVoiceProductMapService(name, client)
		ClientOpmThingModelService = protosService.NewOpmThingModelService(name, client)
		ClientOpmProductService = protosService.NewOpmProductService(name, client)
		//ClientOpmVoiceProductMapGoogleService = protosService.NewOpmVoiceProductMapGoogleService(name, client)
	})

	RegisterGrpcServiceClient(iotconst.IOT_MQTT_SERVICE, func(name string, client client.Client) {
		ClientMqttMessage = protosService.NewMqttService(name, client)
	})

	RegisterGrpcServiceClient(iotconst.IOT_BASIC_SERVICE, func(name string, client client.Client) {
		SysRegionServerService = protosService.NewSysRegionServerService(name, client)
		ConfigDictDataService = protosService.NewConfigDictDataService(name, client)
	})

	//天气微服务
	RegisterGrpcServiceClient(iotconst.IOT_WEATHER_SERVICE, func(name string, client client.Client) {
		IPService = protosService.NewIPService(name, client)
	})
}

func RegisterGrpcServiceClient(name string, fun func(name string, client client.Client)) {
	cli := grpc.NewClient(
		client.Registry(rpc.GetRegistry()),
		client.Wrap(roundrobin.NewClientWrapper()),
		//client.Wrap(wrapper_gobreaker.NewCustomClientWrapper(gobreaker.Settings{Name: name, MaxRequests: 10, Interval: 30 * time.Second, Timeout: 30 * time.Second}, wrapper_gobreaker.BreakServiceEndpoint)),
		client.Wrap(wrapper_uber_ratelimit.NewClientWrapper(1000)),
		client.Wrap(wrapper_opentracing.NewClientWrapper(opentracing.GlobalTracer())),
		client.Retries(0),
		client.RequestTimeout(120*time.Second),
		client.DialTimeout(6*time.Second),
	)
	_ = cli.Init(grpc.MaxSendMsgSize(20 * 1024 * 1024))
	_ = cli.Init(grpc.MaxRecvMsgSize(20 * 1024 * 1024))
	fun(name, cli)
}
