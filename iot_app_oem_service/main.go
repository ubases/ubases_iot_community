package main

import (
	"cloud_platform/iot_app_oem_service/config"
	"cloud_platform/iot_app_oem_service/rpc/rpcServer"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iottrace"
	model "cloud_platform/iot_model"
	"log"

	"github.com/opentracing/opentracing-go"

	"cloud_platform/iot_common/iotlogger"
)

var (
	version string = "2.1.0"
	name           = "iot_app_oem_service"
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
		log.Println("初始化日志发生错误:", err)
		return
	}
	iotlogger.LogHelper.Info("Server start running, current version:", version)

	//链路追踪
	iottrace.SetSamplingFrequency(50)
	t, io, err := iottrace.NewZipkinTracer(iotconst.IOT_APP_OEM_SERVICE, "", config.Global.Zipkin.Url)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	grpcservice := rpcServer.NewGrpcService(iotconst.IOT_APP_OEM_SERVICE, version, serviceCfg.Grpcqps)

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

	iotlogger.LogHelper.Info("数据库成功" + config.Global.Database.Connstr)

	if err = grpcservice.Run(); err != nil {
		iotlogger.LogHelper.Error("grpcservice.Run failed:", err)
	}

	iotlogger.LogHelper.Warn("Server exiting")
}
