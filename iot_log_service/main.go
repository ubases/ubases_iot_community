package main

import (
	"cloud_platform/iot_log_service/config"
	"cloud_platform/iot_log_service/rpc/rpcserver"
	"cloud_platform/iot_log_service/service"
	model "cloud_platform/iot_model"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iottrace"
	"log"

	"github.com/opentracing/opentracing-go"

	"cloud_platform/iot_common/iotlogger"
)

var (
	version string = "2.0.0"
	name           = "iot_log_service"
)

func main() {
	log.Println(version)
	if err := config.Init(); err != nil {
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

	//链路追踪
	iottrace.SetSamplingFrequency(50)
	t, io, err := iottrace.NewZipkinTracer(iotconst.IOT_LOG_SERVICE, "", config.Global.Zipkin.Url)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//grpc服务
	grpcservice := rpcserver.NewGrpcService(iotconst.IOT_LOG_SERVICE, version, serviceCfg.Grpcqps)

	err = grpcservice.RegisterHandler()
	if err != nil {
		iotlogger.LogHelper.Error("grpcservice.RegisterHandler failed:%s", err)
		return
	}

	dbcnf := model.DBConfig{
		DBType:     config.Global.Database.Driver,
		ConnectStr: config.Global.Database.Connstr,
	}
	err = model.InitDB(dbcnf)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return
	}

	// 定时清除clickhouse数据
	service.InitCron()
	if err := service.GetCronJob().StartCronJob(); err != nil {
		iotlogger.LogHelper.Error("启动定时任务错误:", err)
		return
	}
	defer service.GetCronJob().StopCronJob()

	bs, err := service.NewBuildSubscriber()
	if err != nil {
		iotlogger.LogHelper.Error("创建日志服务错误:", err.Error())
		return
	}
	defer bs.Close()
	go bs.Run()

	//运行服务
	if err = grpcservice.Run(); err != nil {
		iotlogger.LogHelper.Error("grpcservice.Run failed:", err)
	}
	iotlogger.LogHelper.Warn("Server exiting")
}
