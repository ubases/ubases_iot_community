package main

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iottrace"
	model "cloud_platform/iot_model"
	"cloud_platform/iot_system_service/cached"
	"cloud_platform/iot_system_service/config"
	"cloud_platform/iot_system_service/rpc"
	"cloud_platform/iot_system_service/service"
	"log"

	"github.com/opentracing/opentracing-go"

	"cloud_platform/iot_common/iotlogger"
)

var (
	version string = "2.0.0"
	name           = "iot_system_service"
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
		log.Println("初始化日志发生错误:", err)
		return
	}
	iotlogger.LogHelper.Info("Server start running, current version:", version)

	//链路追踪
	iottrace.SetSamplingFrequency(50)
	t, io, err := iottrace.NewZipkinTracer(iotconst.IOT_SYSTEM_SERVICE, "", config.Global.Zipkin.Url)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	err = cached.InitCache()
	if err != nil {
		iotlogger.LogHelper.Errorf("Server InitCache failed:%s", err.Error())
		return
	}

	grpcservice := rpc.NewGrpcService(iotconst.IOT_SYSTEM_SERVICE, version, serviceCfg.Grpcqps)

	err = grpcservice.RegisterHandler()
	if err != nil {
		iotlogger.LogHelper.Error("grpcservice.RegisterHandler failed:%s", err)
		return
	}

	//log.Println("load db config!" + iotutil.ToString(config.Global))
	dbcnf := model.DBConfig{
		DBType:     config.Global.Database.Driver,
		ConnectStr: config.Global.Database.Connstr,
	}
	err = model.InitDB(dbcnf)
	///log.Println("load db config post!" + iotutil.ToString(dbcnf))
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return
	}

	err = service.InitCasbin(config.Global.Database.Connstr, "./conf/rbac_model.conf")
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return
	}

	if err = grpcservice.Run(); err != nil {
		iotlogger.LogHelper.Error("grpcservice.Run failed:", err)
	}

	iotlogger.LogHelper.Warn("Server exiting")
}
