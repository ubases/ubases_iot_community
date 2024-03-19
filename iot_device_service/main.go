package main

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iottrace"
	"cloud_platform/iot_device_service/cached"
	"cloud_platform/iot_device_service/config"
	"cloud_platform/iot_device_service/rpc"
	"cloud_platform/iot_device_service/service"
	"cloud_platform/iot_device_service/service/job"
	model "cloud_platform/iot_model"
	"log"

	"github.com/opentracing/opentracing-go"

	"cloud_platform/iot_common/iotlogger"
)

var (
	version string = "2.0.0"
	name           = "iot_device_service"
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
		iotlogger.LogHelper.Errorf("cached.InitCache failed:%s", err.Error())
		return
	}

	err = config.FaultConfig.Init()
	if err != nil {
		iotlogger.LogHelper.Errorf("FaultConfig.Init failed:%s", err.Error())
		return
	}

	//链路追踪
	iottrace.SetSamplingFrequency(50)
	t, io, err := iottrace.NewZipkinTracer(iotconst.IOT_DEVICE_SERVICE, "", config.Global.Zipkin.Url)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//grpc服务
	grpcservice := rpc.NewGrpcService(iotconst.IOT_DEVICE_SERVICE, version, serviceCfg.Grpcqps)

	err = grpcservice.RegisterHandler()
	if err != nil {
		iotlogger.LogHelper.Error("grpcservice.RegisterHandler failed:%s", err)
		return
	}

	for _, v := range config.Global.Database {
		dbcnf := model.DBConfig{DBType: v.Driver, ConnectStr: v.Connstr}
		if db, err := model.InitDBEx(dbcnf); err != nil {
			iotlogger.LogHelper.Errorf("初始化数据库连接失败:%s", err.Error())
			return
		} else {
			if v.Database == "iot_device" && v.Driver != "clickhouse" {
				model.SetDB(db)
			} else {
				config.DBMap[v.Database] = db
			}
		}
	}

	//发布job消息
	go job.RunToMQTT()

	//如果当前是第一个进程，则运行job，否则不运行
	if !service.StartAuto() {
		service.RunJob()
	}

	err = service.GetJsPublisherMgr().AddPublisher(iotconst.NATS_STREAM_DEVICE, iotconst.NATS_SUBJECT_AUTH, config.Global.Nats.Addrs)
	if err != nil {
		iotlogger.LogHelper.Errorf("初始化发布器失败:%s", err.Error())
		return
	}
	err = service.GetJsPublisherMgr().AddPublisher(iotconst.NATS_STREAM_APP, iotconst.NATS_SUBJECT_RECORDS, config.Global.Nats.Addrs)
	if err != nil {
		iotlogger.LogHelper.Errorf("日志初始化发布器失败:%s", err.Error())
		return
	}
	err = service.GetJsPublisherMgr().AddPublisherEx(iotconst.NATS_APPNAME_PRODUCT, iotconst.NATS_PRODUCT_PUBLISH, iotconst.NATS_SUBJECT_PRODUCT_PUBLISH, config.Global.Nats.Addrs)
	if err != nil {
		iotlogger.LogHelper.Errorf("日志初始化发布器失败:%s", err.Error())
		return
	}
	go service.GetJsPublisherMgr().Run()
	//数据表更新
	//db_device.Migrate(model.GetDB())
	//mqtt订阅初始化
	//go subdata.NewSubstrbuteDeviceAttr().SubTopic()
	//开启设备激活订阅服务
	bs, err := service.InitSubscriber()
	if err != nil {
		iotlogger.LogHelper.Error("创建设备激活订阅服务错误:", err.Error())
		return
	}
	go bs.RunSub()
	defer bs.CloseSub()

	iotlogger.LogHelper.Info("starting")
	//运行服务
	if err = grpcservice.Run(); err != nil {
		iotlogger.LogHelper.Error("grpcservice.RunToMQTT failed:", err)
	}
	iotlogger.LogHelper.Warn("Server exiting")
}
