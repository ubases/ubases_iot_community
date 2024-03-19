package main

import (
	"cloud_platform/iot_basic_service/cached"
	"cloud_platform/iot_basic_service/config"
	"cloud_platform/iot_basic_service/rpc"
	"cloud_platform/iot_basic_service/service"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iottrace"
	model "cloud_platform/iot_model"
	"log"

	"github.com/opentracing/opentracing-go"

	"cloud_platform/iot_common/iotlogger"
)

var (
	version string = "2.0.0"
	name           = "iot_basic_service"
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

	err = cached.InitCache()
	if err != nil {
		iotlogger.LogHelper.Errorf("Server start failed:%s", err.Error())
		return
	}

	//链路追踪
	iottrace.SetSamplingFrequency(50)
	t, io, err := iottrace.NewZipkinTracer(iotconst.IOT_BASIC_SERVICE, "", config.Global.Zipkin.Url)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//grpc服务
	grpcservice := rpc.NewGrpcService(iotconst.IOT_BASIC_SERVICE, version, serviceCfg.Grpcqps)

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

	//推送翻译修改
	err = service.GetJsPublisherMgr().AddPublisher(iotconst.NATS_SUBJECT_LANGUAGE_UPDATE, config.Global.Nats.Addrs)
	if err != nil {
		iotlogger.LogHelper.Errorf("初始化发布器失败:%s", err.Error())
		return
	}
	go service.GetJsPublisherMgr().Run()

	//初始化缓存
	service.LangCached()

	bs, err := service.NewBuildSubscriber()
	if err != nil {
		iotlogger.LogHelper.Error("创建翻译新增订阅服务错误:", err.Error())
		return
	}
	go bs.Run()
	defer bs.Close()

	//运行服务
	if err = grpcservice.Run(); err != nil {
		iotlogger.LogHelper.Error("grpcservice.Run failed:", err)
	}
	iotlogger.LogHelper.Warn("Server exiting")
}
