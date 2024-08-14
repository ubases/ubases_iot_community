package main

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iottrace"
	"cloud_platform/iot_demo_api_service/cached"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/opentracing/opentracing-go"

	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_demo_api_service/config"
	routers "cloud_platform/iot_demo_api_service/router"

	"cloud_platform/iot_demo_api_service/rpc"
)

var (
	version string = "2.1.0"
	name           = "iot_demo_api_service"
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
	t, io, err := iottrace.NewZipkinTracer(iotconst.IOT_DEMO_API_SERVICE, serviceCfg.HttpAddr, config.Global.Zipkin.Url)
	if err != nil {
		iotlogger.LogHelper.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//grpc客户端
	rpc.InitServiceClient()

	//注册 gin  routers
	srv := &http.Server{Addr: config.Global.Service.HttpAddr, Handler: routers.Init()}
	//设置超时时间
	if serviceCfg.ReadTimeout > 0 {
		srv.ReadTimeout = time.Duration(serviceCfg.ReadTimeout) * time.Second
	}
	if serviceCfg.WriteTimeout > 0 {
		srv.WriteTimeout = time.Duration(serviceCfg.WriteTimeout) * time.Second
	}
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
