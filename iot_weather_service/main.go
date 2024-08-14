package main

import (
	"cloud_platform/iot_common/iottrace"
	"cloud_platform/iot_weather_service/service"
	"cloud_platform/iot_weather_service/service/cache"
	"cloud_platform/iot_weather_service/service/geoip"
	"log"

	"github.com/opentracing/opentracing-go"

	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_weather_service/config"
	"cloud_platform/iot_weather_service/rpc/rpcserver"
)

var (
	version string = "2.1.0"
	name           = "iot_weather_service"
)

func main() {
	log.Println(version)
	if err := config.Init2(); err != nil {
		log.Println("加载配置文件发生错误:", err)
		return
	}
	serviceCfg := config.Global.Service
	//统一日志到服务的日志
	err := iotlogger.InitLog(serviceCfg.Logfile, name, serviceCfg.Loglevel)
	if err != nil {
		panic(err)
	}
	iotlogger.LogHelper.Info("Server start running, current version:", version)

	if err := cache.InitCache(); err != nil {
		iotlogger.LogHelper.Error("cache.InitCache failed:%s", err)
		return
	}

	if err = geoip.InitGeoMgr(); err != nil {
		iotlogger.LogHelper.Error("service.InitGeoMgr failed:%s", err)
		return
	}

	go service.GetSubscriber().WatchWeather()

	//链路追踪
	iottrace.SetSamplingFrequency(50)
	t, io, err := iottrace.NewZipkinTracer(iotconst.IOT_WEATHER_SERVICE, "", config.Global.Zipkin.Url)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//grpc服务
	grpcservice := rpcserver.NewGrpcService(iotconst.IOT_WEATHER_SERVICE, version, serviceCfg.Grpcqps)
	err = grpcservice.RegisterHandler()
	if err != nil {
		iotlogger.LogHelper.Error("grpcservice.RegisterHandler failed:%s", err)
		return
	}

	//运行服务
	if err = grpcservice.Run(); err != nil {
		iotlogger.LogHelper.Error("grpcservice.Run failed:", err)
	}
	service.GetSubscriber().Close()

	iotlogger.LogHelper.Warn("Server exiting")
}
