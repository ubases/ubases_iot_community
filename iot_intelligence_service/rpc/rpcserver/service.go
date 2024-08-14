package rpcserver

import (
	"cloud_platform/iot_intelligence_service/handler"
	"cloud_platform/iot_intelligence_service/rpc"
	"cloud_platform/iot_intelligence_service/rpc/rpcclient"
	"time"

	"go-micro.dev/v4/server"

	"github.com/asim/go-micro/plugins/server/grpc/v4"
	"github.com/asim/go-micro/plugins/wrapper/monitoring/prometheus/v4"
	wrapper_uber_ratelimit "github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v4"
	wrapper_opentracing "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v4"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4"

	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_common/iotlogger"
)

type GrpcService struct {
	micro.Service
}

func NewGrpcService(name, version string, qps int) *GrpcService {
	service := micro.NewService(
		micro.Server(grpc.NewServer(server.Wait(nil))),
		micro.Name(name),
		micro.Version(version),
		//micro.Address(setting.Global.Server.Addr),
		micro.Registry(rpc.GetRegistry()),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.WrapHandler(wrapper_uber_ratelimit.NewHandlerWrapper(qps)),                                                  //服务限流
		micro.WrapHandler(wrapper_opentracing.NewHandlerWrapper(opentracing.GlobalTracer())),                              //链路追踪
		micro.WrapHandler(prometheus.NewHandlerWrapper(prometheus.ServiceName(name), prometheus.ServiceVersion(version))), //服务监控
		micro.WrapHandler(ioterrs.BatPanicHandler()),
	)
	_ = service.Server().Init(grpc.MaxMsgSize(20 * 1024 * 1024))
	s := &GrpcService{service}
	service.Init(micro.BeforeStart(s.BeforeStart), micro.AfterStart(s.AfterStart))
	return s
}

func (s *GrpcService) BeforeStart() error {
	rpcclient.InitServiceClient()
	iotlogger.LogHelper.Infof("service %s BeforeStart", s.Name())

	return nil
}

func (s *GrpcService) AfterStart() error {
	iotlogger.LogHelper.Infof("service %s AfterStart address %s", s.Name(), s.Server().Options().Address)
	return nil
}

func (s *GrpcService) RegisterHandler() error {
	handler.RegisterHandler(s)
	return nil
}
