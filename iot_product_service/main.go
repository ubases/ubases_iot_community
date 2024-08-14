package main

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotnatsjs"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_common/iottrace"
	model "cloud_platform/iot_model"
	"cloud_platform/iot_product_service/cached"
	"cloud_platform/iot_product_service/config"
	"cloud_platform/iot_product_service/rpc"
	"cloud_platform/iot_product_service/service"
	"log"

	"github.com/opentracing/opentracing-go"

	"cloud_platform/iot_common/iotlogger"
)

var (
	version string = "2.1.0"
	name           = "iot_product_service"
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
	t, io, err := iottrace.NewZipkinTracer(iotconst.IOT_PRODUCT_SERVICE, "", config.Global.Zipkin.Url)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//grpc服务
	grpcservice := rpc.NewGrpcService(iotconst.IOT_PRODUCT_SERVICE, version, serviceCfg.Grpcqps)

	err = grpcservice.RegisterHandler()
	if err != nil {
		iotlogger.LogHelper.Error("grpcservice.RegisterHandler failed:%s", err)
		return
	}

	//fixme 数据库配置请在合适的地方加载
	dbcnf := model.DBConfig{
		DBType:     config.Global.Database.Driver,
		ConnectStr: config.Global.Database.Connstr,
	}
	err = model.InitDB(dbcnf)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return
	}

	//初始化产品缓存
	//service.ProductCached()

	//推送翻译修改
	err = iotnatsjs.GetJsClientPub().InitJsClient(config.Global.Nats.Addrs)
	if err != nil {
		iotlogger.LogHelper.Errorf("初始化发布器失败:%s", err.Error())
		return
	}
	go iotnatsjs.GetJsClientPub().Run()

	iotnatsjs.GetJsClientPub().PushData(&iotnatsjs.NatsPubData{
		Subject: iotconst.NATS_SUBJECT_LANGUAGE_UPDATE,
		Data:    iotstruct.TranslatePush{}.SetContent(iotconst.LANG_PRODUCT_NAME, "test", "name", "test", "test2"),
	})

	bs, err := service.InitSubscriber()
	if err != nil {
		iotlogger.LogHelper.Error("创建设备激活订阅服务错误:", err.Error())
		return
	}
	go bs.RunSub()
	defer bs.CloseSub()

	//运行服务
	if err = grpcservice.Run(); err != nil {
		iotlogger.LogHelper.Error("grpcservice.Run failed:", err)
	}
	iotlogger.LogHelper.Infof("服务%s开已停止", "iot_product_service")
}
