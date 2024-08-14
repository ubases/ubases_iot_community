package main

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iottrace"
	model "cloud_platform/iot_model"
	"cloud_platform/iot_voice_service/cached"
	"cloud_platform/iot_voice_service/config"
	routers "cloud_platform/iot_voice_service/router"
	"cloud_platform/iot_voice_service/rpc/rpcclient"
	"cloud_platform/iot_voice_service/rpc/rpcserver"
	"cloud_platform/iot_voice_service/service"
	"cloud_platform/iot_voice_service/service/alexa"
	"cloud_platform/iot_voice_service/service/common"
	"context"
	"log"
	"net/http"

	"github.com/opentracing/opentracing-go"

	"cloud_platform/iot_common/iotlogger"
)

var (
	version string = "2.1.0"
	name           = "iot_voice_service"
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

	iottrace.SetSamplingFrequency(50)
	t, io, err := iottrace.NewZipkinTracer("bat.iot.voice.service", serviceCfg.HttpAddr, config.Global.Zipkin.Url)
	if err != nil {
		iotlogger.LogHelper.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//grpc服务
	grpcservice := rpcserver.NewGrpcService(iotconst.IOT_SMART_SPEAKER_SERVICE, version, serviceCfg.Grpcqps)

	err = grpcservice.RegisterHandler()
	if err != nil {
		iotlogger.LogHelper.Error("grpcservice.RegisterHandler failed:%s", err)
		return
	}

	//grpc客户端
	rpcclient.InitServiceClient()

	dbcnf := model.DBConfig{
		DBType:     config.Global.Database.Driver,
		ConnectStr: config.Global.Database.Connstr,
	}
	err = model.InitDB(dbcnf)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return
	}

	// 初始化oauth2.0服务器
	if err := routers.InitOAuth2(); err != nil {
		iotlogger.LogHelper.Helper.Error("init oauth2 server error: ", err)
		return
	}

	go common.GetNatsSubscriber().Run()

	ctx, cancel := context.WithCancel(context.Background())
	go service.InitReportSub(ctx)

	//注册 gin  routers
	ginHandler := routers.Init()

	srv := &http.Server{
		Addr:    config.Global.Service.HttpAddr,
		Handler: ginHandler,
	}
	go func() {
		iotlogger.LogHelper.Info(config.Global.Service.HttpAddr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			iotlogger.LogHelper.Fatalf("listen: %s\n", err)
		}
	}()

	alexa.InitTimerCron()

	//运行服务
	if err = grpcservice.Run(); err != nil {
		iotlogger.LogHelper.Error("grpcservice.Run failed:", err)
	}
	cancel()
	iotlogger.LogHelper.Warn("Server exiting")
}
