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

// 要调用的外部服务声明
var (
// ClientXXXXXXXXXService protosService.XXXXXXXXXService
)

// 外部服务客户端初始化
func initServiceClient() {
	//打开下列注释，其中iotconst.XXXXXXXXXXX为服务名，参考官方服务在iotconst下定义即可
	//RegisterGrpcServiceClient(iotconst.XXXXXXXXXXX, func(name string, client client.Client) {
	//  初始化NewXXXXXXXXXService客户端
	//	ClientXXXXXXXXXService = protosService.NewXXXXXXXXXService(name, client)
	//})
	//其它微服务的客户端初始化往后面加RegisterGrpcServiceClient方法
}

func RegisterGrpcServiceClient(name string, fun func(name string, client client.Client)) {
	cli := grpc.NewClient(
		client.Registry(getRegistry()),
		client.Wrap(roundrobin.NewClientWrapper()),
		client.Wrap(wrapper_uber_ratelimit.NewClientWrapper(1000)),
		client.Wrap(wrapper_opentracing.NewClientWrapper(opentracing.GlobalTracer())),
		client.Retries(0),
		client.RequestTimeout(60*time.Second),
		client.DialTimeout(6*time.Second),
	)
	_ = cli.Init(grpc.MaxSendMsgSize(20 * 1024 * 1024))
	_ = cli.Init(grpc.MaxRecvMsgSize(20 * 1024 * 1024))
	fun(name, cli)
}
