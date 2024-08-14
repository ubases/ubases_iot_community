package main

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iottrace"
	"cloud_platform/iot_message_service/cached"
	"cloud_platform/iot_message_service/config"
	"cloud_platform/iot_message_service/rpc/rpcserver"
	"cloud_platform/iot_message_service/service"
	"cloud_platform/iot_message_service/service/push"
	model "cloud_platform/iot_model"
	"log"

	"github.com/opentracing/opentracing-go"

	"cloud_platform/iot_common/iotlogger"
)

var (
	version string = "2.1.0"
	name           = "iot_message_service"
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

	err = cached.InitCache()
	if err != nil {
		iotlogger.LogHelper.Errorf("Server start failed:%s", err.Error())
		return
	}

	service.EmailMgr.Init()
	service.EmailMgr.StartQueueHandle()
	defer service.EmailMgr.Close()

	if err = service.SmdMgr.Init(); err != nil {
		iotlogger.LogHelper.Error("SmdMgr.Init() failed:%s", err)
		return
	}

	go service.SmdMgr.QueueHandle()
	//测试发短信
	//err = service.SmdMgr.SendSMS(service.CodeInput{Code: "123456"}, "SMS_182405394", "15013739503")
	//if err != nil {
	//	fmt.Println(err)
	//}
	defer service.SmdMgr.Close()

	//链路追踪
	iottrace.SetSamplingFrequency(50)
	t, io, err := iottrace.NewZipkinTracer(iotconst.IOT_MESSAGE_SERVICE, "", config.Global.Zipkin.Url)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//grpc服务
	grpcservice := rpcserver.NewGrpcService(iotconst.IOT_MESSAGE_SERVICE, version, serviceCfg.Grpcqps)

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

	bs, err := service.NewBuildSubscriber()
	if err != nil {
		iotlogger.LogHelper.Error("推送消息订阅服务错误:", err.Error())
		return
	}
	go bs.Run()
	defer bs.Close()

	//初始化推送客户端队列
	if err = push.PushMgr.Init(); err != nil {
		iotlogger.LogHelper.Error("PushClientMgr.Init() failed:%s", err)
		return
	}
	go push.PushMgr.QueueHandle()

	//运行服务
	if err = grpcservice.Run(); err != nil {
		iotlogger.LogHelper.Error("grpcservice.Run failed:", err)
	}
	iotlogger.LogHelper.Warn("Server exiting")
}
