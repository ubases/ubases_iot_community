package main

import (
	"cloud_platform/iot_app_api_service/cached"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotnats/jetstream"
	"cloud_platform/iot_common/iottrace"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/opentracing/opentracing-go"

	"cloud_platform/iot_app_api_service/config"
	routers "cloud_platform/iot_app_api_service/router"
	"cloud_platform/iot_common/iotlogger"
)

var (
	version string = "2.0.0"
	name           = "iot_app_api_service"
)

// @title api title
// @version 1.0
// @description api description
// @termsOfService 服务条款

// @contact.name 联系人姓名
// @contact.url 联系地址
// @contact.email 联系邮箱

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:1880
// @BasePath /v1/platform
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

	err = cached.InitMsg()
	if err != nil {
		iotlogger.LogHelper.Errorf("Server init msg failed:%s", err.Error())
		return
	}

	iottrace.SetSamplingFrequency(50)
	t, io, err := iottrace.NewZipkinTracer(iotconst.IOT_APP_API_SERVICE, serviceCfg.HttpAddr, config.Global.Zipkin.Url)
	if err != nil {
		iotlogger.LogHelper.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// 初始化nats消息队列，新增发布者
	if err := jetstream.GetJsPublisherMgr().AddPublisher("iot_app_api_service", iotconst.NATS_STREAM_APP, iotconst.NATS_SUBJECT_RECORDS, config.Global.Nats.Addrs); err != nil {
		iotlogger.LogHelper.Error("nats.AddPublisher failed:%s", err)
		return
	}

	go jetstream.GetJsPublisherMgr().Run()

	//grpc客户端
	rpc.InitServiceClient()

	srv := &http.Server{Addr: config.Global.Service.HttpAddr, Handler: routers.Init()}
	go listenAndServe(srv)
	gracefulShutdown(srv)
}

func listenAndServe(srv *http.Server) {
	iotlogger.LogHelper.Info(config.Global.Service.HttpAddr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		iotlogger.LogHelper.Fatalf("listen: %s\n", err)
	}
}

func gracefulShutdown(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	iotlogger.LogHelper.Warn("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		iotlogger.LogHelper.Error("Server forced to shutdown: ", err)
	}
	iotlogger.LogHelper.Warn("Server exiting")
}
