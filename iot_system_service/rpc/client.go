package rpc

import (
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
