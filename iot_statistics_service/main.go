package main

import (
	model "cloud_platform/iot_model"
	"cloud_platform/iot_statistics_service/etl"
	"cloud_platform/iot_statistics_service/rpc"
	"cloud_platform/iot_statistics_service/task"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iottrace"
	"log"
	"strings"

	"github.com/opentracing/opentracing-go"

	"github.com/xxl-job/xxl-job-executor-go"

	"cloud_platform/iot_common/iotlogger"

	"cloud_platform/iot_statistics_service/config"
)

var (
	version string = "2.0.0"
	name           = "iot_statistics_service"
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

	for _, v := range config.Global.Database {
		dbcnf := model.DBConfig{DBType: v.Driver, ConnectStr: v.Connstr}
		if db, err := model.InitDBEx(dbcnf); err != nil {
			iotlogger.LogHelper.Errorf("初始化数据库连接失败:%s", err.Error())
			return
		} else {
			config.DBMap[v.Database] = db
			if v.Database == "iot_statistics" {
				model.SetDB(db)
			}
		}
	}
	cnf := iotredis.Config{
		Cluster:      config.Global.Redis.Cluster,
		Addrs:        strings.Join(config.Global.Redis.Addrs, ","),
		Username:     config.Global.Redis.Username,
		Password:     config.Global.Redis.Password,
		Database:     config.Global.Redis.Database,
		MinIdleConns: config.Global.Redis.MinIdleConns,
		IdleTimeout:  config.Global.Redis.IdleTimeout,
		PoolSize:     config.Global.Redis.PoolSize,
		MaxConnAge:   config.Global.Redis.MaxConnAge,
	}
	_, err = iotredis.NewClient(cnf)
	if err != nil {
		iotlogger.LogHelper.Errorf("初始化redis连接:%s", err.Error())
		return
	}

	if config.Global.XxlJob.Enable {
		go XxlService()
	}

	//链路追踪
	iottrace.SetSamplingFrequency(50)
	t, io, err := iottrace.NewZipkinTracer(iotconst.IOT_STATISTICS_SERVICE, "", config.Global.Zipkin.Url)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//grpc服务
	grpcservice := rpc.NewGrpcService(iotconst.IOT_STATISTICS_SERVICE, version, serviceCfg.Grpcqps)
	err = grpcservice.RegisterHandler()
	if err != nil {
		iotlogger.LogHelper.Error("grpcservice.RegisterHandler failed:%s", err)
		return
	}

	var uopt etl.UcUserOperateETL
	go uopt.Handler()

	//运行服务
	if err = grpcservice.Run(); err != nil {
		iotlogger.LogHelper.Error("grpcservice.Run failed:", err)
	}

	iotlogger.LogHelper.Info("Server stoped")
}

func XxlService() error {
	xxlJobCnf := config.Global.XxlJob
	exec := xxl.NewExecutor(
		xxl.ServerAddr(xxlJobCnf.ServerAddr),
		xxl.AccessToken(xxlJobCnf.AccessToken),
		xxl.ExecutorIp(xxlJobCnf.ExecutorIp),
		xxl.ExecutorPort(xxlJobCnf.ExecutorPort),
		xxl.RegistryKey(xxlJobCnf.RegistryKey),
		xxl.SetLogger(&task.XxlLogger{}),
	)
	exec.Init()
	exec.LogHandler(task.XxlLogHandler)
	//注册任务handler
	exec.RegTask("task.hour.dataoverview", task.HourDataOveriewStatistics)              //每小时1次
	exec.RegTask("task.month.dataoverview", task.MonthDataOveriewStatistics)            //每天1次
	exec.RegTask("task.hour.device.active", task.HourActive)                            //每小时1次
	exec.RegTask("task.month.device.fault", task.MonthFault)                            //每天1次
	exec.RegTask("task.developer.statistics", task.DeveloperStatistics)                 //全量，每天1次
	exec.RegTask("task.app.statistics", task.AppListStatistics)                         //全量，每天1次
	exec.RegTask("task.month.app.user", task.MonthAppUser)                              //每天1次
	exec.RegTask("task.day.app.user", task.DayAppActiveUser)                            //每天1次
	exec.RegTask("task.hour.history.device.active", task.HistoryActive)                 //手动执行
	exec.RegTask("task.month.history.device.fault", task.HistoryFault)                  //手动执行
	exec.RegTask("task.month.history.app.user", task.MonthHistoryAppUser)               //手动执行
	exec.RegTask("task.hour.history.dataoverview", task.HistoryHourDataOveriewActive)   //手动执行
	exec.RegTask("task.month.history.dataoverview", task.HistoryMonthDataOveriewActive) //手动执行
	if err := exec.Run(); err != nil {
		iotlogger.LogHelper.Errorf("exec.Run error:%s", err.Error())
		return err
	}
	return nil
}
