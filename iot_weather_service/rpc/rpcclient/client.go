package rpcclient

import (
	"cloud_platform/iot_weather_service/rpc"
	"time"

	"github.com/asim/go-micro/plugins/client/grpc/v4"

	wrapper_uber_ratelimit "github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v4"
	"github.com/asim/go-micro/plugins/wrapper/select/roundrobin/v4"
	wrapper_opentracing "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v4"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4/client"
)

// 初始化其它服务的客户端
func InitServiceClient() {

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
