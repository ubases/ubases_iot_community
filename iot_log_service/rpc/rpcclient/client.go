package rpcclient

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_log_service/rpc"
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
	ClientUcUserService protosService.UcUserService
	ClientOemAppService protosService.OemAppService
)

// 初始化其它服务的客户端
func InitServiceClient() {
	RegisterGrpcServiceClient(iotconst.IOT_APP_USER_SERVICE, func(name string, client client.Client) {
		ClientUcUserService = protosService.NewUcUserService(name, client)
	})
	RegisterGrpcServiceClient(iotconst.IOT_APP_OEM_SERVICE, func(name string, client client.Client) {
		ClientOemAppService = protosService.NewOemAppService(name, client)
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
	fun(name, cli)
}
