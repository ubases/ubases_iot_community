package main

import (
	"cloud_platform/iot_app_user_service/cached"
	"cloud_platform/iot_app_user_service/config"
	"cloud_platform/iot_app_user_service/rpc"
	"cloud_platform/iot_app_user_service/service"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iottrace"
	"cloud_platform/iot_common/iotutil"
	model "cloud_platform/iot_model"
	"context"
	"log"

	"github.com/opentracing/opentracing-go"

	"go-micro.dev/v4/broker"

	"cloud_platform/iot_common/iotlogger"
)

var (
	version string = "2.1.0"
	name           = "iot_app_user_service"
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

	//链路追踪
	iottrace.SetSamplingFrequency(50)
	t, io, err := iottrace.NewZipkinTracer(iotconst.IOT_APP_USER_SERVICE, "", config.Global.Zipkin.Url)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//grpc服务
	grpcservice := rpc.NewGrpcService(iotconst.IOT_APP_USER_SERVICE, version, serviceCfg.Grpcqps)

	err = grpcservice.RegisterHandler()
	if err != nil {
		iotlogger.LogHelper.Error("grpcservice.RegisterHandler failed:%s", err)
		return
	}

	log.Println("load db config!" + iotutil.ToString(config.Global))
	dbcnf := model.DBConfig{
		DBType:     config.Global.Database.Driver,
		ConnectStr: config.Global.Database.Connstr,
	}
	err = model.InitDB(dbcnf)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return
	}
	//初始化微信小程序
	service.InitWechat()

	// 通知清理缓存
	ctx, cancel := context.WithCancel(context.Background())
	go service.InitClearCachedSub(ctx)

	//运行服务
	if err = grpcservice.Run(); err != nil {
		iotlogger.LogHelper.Error("grpcservice.Run failed:", err)
	}
	cancel()
	iotlogger.LogHelper.Warn("Server exiting")
}

func SubHandler(m broker.Event) error {
	iotlogger.LogHelper.Info("[sub] received message:", string(m.Message().Body), "header", m.Message().Header)
	return nil
}
