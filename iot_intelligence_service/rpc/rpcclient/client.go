package rpcclient

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_intelligence_service/rpc"
	"cloud_platform/iot_proto/protos/protosService"
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
	ClientAppMessage     protosService.MpMessageService
	ClientAppHomeUser    protosService.UcHomeUserService
	ClientWeatherService protosService.WeatherService
	ClientMqttMessage    protosService.MqttService
)

// 初始化其它服务的客户端
func InitServiceClient() {
	RegisterGrpcServiceClient(iotconst.IOT_MESSAGE_SERVICE, func(name string, client client.Client) {
		ClientAppMessage = protosService.NewMpMessageService(name, client)
	})
	RegisterGrpcServiceClient(iotconst.IOT_APP_USER_SERVICE, func(name string, client client.Client) {
		ClientAppHomeUser = protosService.NewUcHomeUserService(name, client)
	})
	RegisterGrpcServiceClient(iotconst.IOT_WEATHER_SERVICE, func(name string, client client.Client) {
		ClientWeatherService = protosService.NewWeatherService(name, client)
	})
	RegisterGrpcServiceClient(iotconst.IOT_MQTT_SERVICE, func(name string, client client.Client) {
		ClientMqttMessage = protosService.NewMqttService(name, client)
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
		client.RequestTimeout(120*time.Second), //for debug
		client.DialTimeout(6*time.Second),
	)
	_ = cli.Init(grpc.MaxSendMsgSize(20 * 1024 * 1024))
	_ = cli.Init(grpc.MaxRecvMsgSize(20 * 1024 * 1024))
	fun(name, cli)
}
