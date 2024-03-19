package main

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iottrace"
	"cloud_platform/iot_demo_service/config"
	"cloud_platform/iot_demo_service/rpc"
	model "cloud_platform/iot_model"
	"log"

	"github.com/opentracing/opentracing-go"

	"cloud_platform/iot_common/iotlogger"
)

var (
	version string = "2.0.0"
	name           = "iot_demo_service"
)

func main() {
	log.Println(version)
	//初始化配置
	if err := config.Init(); err != nil {
		log.Println("加载配置文件发生错误:", err)
		return
	}
	serviceCfg := config.Global.Service
	//初始化日志
	err := iotlogger.InitLog(serviceCfg.Logfile, name, serviceCfg.Loglevel)
	if err != nil {
		panic(err)
	}
	iotlogger.LogHelper.Info("Server start running, current version:", version)

	//链路追踪配置
	iottrace.SetSamplingFrequency(50)
	t, io, err := iottrace.NewZipkinTracer(iotconst.IOT_DEMO_SERVICE, "", config.Global.Zipkin.Url)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//GRPC初始化及服务注册
	grpcservice := rpc.NewGrpcService(iotconst.IOT_DEMO_SERVICE, version, serviceCfg.Grpcqps)
	err = grpcservice.RegisterHandler()
	if err != nil {
		iotlogger.LogHelper.Error("grpcservice.RegisterHandler failed:%s", err)
		return
	}

	//初始化数据库连接
	dbcnf := model.DBConfig{
		DBType:     config.Global.Database.Driver,
		ConnectStr: config.Global.Database.Connstr,
	}
	err = model.InitDB(dbcnf)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return
	}
	//运行服务
	if err = grpcservice.Run(); err != nil {
		iotlogger.LogHelper.Error("grpcservice.Run failed:", err)
	}
	iotlogger.LogHelper.Warn("Server exiting")
}
