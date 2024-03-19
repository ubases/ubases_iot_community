package rpc

import (
	"cloud_platform/iot_app_user_service/rpc/rpcClient"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_proto/protos/protosService"
	"time"

	"github.com/asim/go-micro/plugins/client/grpc/v4"

	wrapper_uber_ratelimit "github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v4"
	"github.com/asim/go-micro/plugins/wrapper/select/roundrobin/v4"
	wrapper_opentracing "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v4"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4/client"
)

// 初始化其它服务的客户端
func initServiceClient() {
	RegisterGrpcServiceClient(iotconst.IOT_DEVICE_SERVICE, func(name string, client client.Client) {
		rpcClient.IotDeviceHomeService = protosService.NewIotDeviceHomeService(name, client)
	})
	RegisterGrpcServiceClient(iotconst.IOT_PRODUCT_SERVICE, func(name string, client client.Client) {
		rpcClient.OpmProductService = protosService.NewOpmProductService(name, client)
	})
	RegisterGrpcServiceClient(iotconst.IOT_BASIC_SERVICE, func(name string, client client.Client) {
		rpcClient.DictDataService = protosService.NewConfigDictDataService(name, client)
		rpcClient.ClientSysAreaService = protosService.NewSysAreaService(name, client)
	})
	RegisterGrpcServiceClient(iotconst.IOT_APP_OEM_SERVICE, func(name string, client client.Client) {
		rpcClient.ClientOemAppDocDirService = protosService.NewOemAppDocDirService(name, client)
		rpcClient.ClientOemAppService = protosService.NewOemAppService(name, client)
		rpcClient.ClientOemAppUiConfigService = protosService.NewOemAppUiConfigService(name, client)
	})
	//天气微服务
	RegisterGrpcServiceClient(iotconst.IOT_WEATHER_SERVICE, func(name string, client client.Client) {
		rpcClient.IPService = protosService.NewIPService(name, client)
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
