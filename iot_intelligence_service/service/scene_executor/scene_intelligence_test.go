package scene_executor

import (
	"cloud_platform/iot_intelligence_service/cached"
	"cloud_platform/iot_intelligence_service/config"
	"cloud_platform/iot_intelligence_service/rpc/rpcclient"
	cron2 "cloud_platform/iot_intelligence_service/service/scene_executor/cron"
	"cloud_platform/iot_intelligence_service/service/scene_executor/execute"
	model "cloud_platform/iot_model"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"
)

// 条件测试
func TestCondition(t *testing.T) {
	config.InitTest("../../")
	cached.InitCache()
	serviceCfg := config.Global.Service
	//统一日志到服务的日志
	err := iotlogger.InitLog(serviceCfg.Logfile, "iot_intelligence_service", serviceCfg.Loglevel)
	if err != nil {
		panic(err)
	}
	iotlogger.LogHelper.Info("Server start running")
	dbcnf := model.DBConfig{
		DBType:     config.Global.Database.Driver,
		ConnectStr: config.Global.Database.Connstr,
	}
	err = model.InitDB(dbcnf)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return
	}
	cron2.InitTimerCron()
	rpcclient.InitServiceClient()

	NewBuilder()
	//executer := SceneIntelligenceCondition{}
	rule := IntelligenceRule{}
	//rules2.CreateRuleChan
	rule.initCreateSub()
	rule.CreateRule("953350369763557376", "test")
	//t.Log(executer.ConditionExecutor("7127477083537571840", ""))
	//t.Log(executer.ConditionExecutor("5924289872947740672"))

	select {}
}

func TestTimer(t *testing.T) {
	config.InitTest("../../")
	cached.InitCache()
	serviceCfg := config.Global.Service
	//统一日志到服务的日志
	err := iotlogger.InitLog(serviceCfg.Logfile, "iot_intelligence_service", serviceCfg.Loglevel)
	if err != nil {
		panic(err)
	}
	iotlogger.LogHelper.Info("Server start running")
	dbcnf := model.DBConfig{
		DBType:     config.Global.Database.Driver,
		ConnectStr: config.Global.Database.Connstr,
	}
	err = model.InitDB(dbcnf)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return
	}
	cron2.InitTimerCron()
	rpcclient.InitServiceClient()

	_, err = cron2.CronCtx.AddFunc("20 33 10 * * 1,2,3,4,5,6,0", func() {
		t.Log("success")
	})
	if err != nil {
		t.Error(err)
	}

	select {}
}

// 任务测试
func TestTask(t *testing.T) {
	config.InitTest("../../")
	cached.InitCache()
	serviceCfg := config.Global.Service
	//统一日志到服务的日志
	err := iotlogger.InitLog(serviceCfg.Logfile, "iot_device_service", serviceCfg.Loglevel)
	if err != nil {
		panic(err)
	}
	iotlogger.LogHelper.Info("Server start running")
	////初始化grpc客户端
	rpcclient.InitServiceClient()

	//初始化数据库
	dbcnf := model.DBConfig{
		DBType:     config.Global.Database.Driver,
		ConnectStr: config.Global.Database.Connstr,
	}
	err = model.InitDB(dbcnf)
	if err != nil {
		t.Error(err)
		return
	}
	//devIds := new([]string)
	go func() {
		//推送通知
		res, err := execute.NoticeExecute(execute.NoticeParams{
			HomeId:           3492204888857673728,
			UserId:           []int64{5286625992243249152},
			IntelligenceName: "",
		}, 123, time.Now().Unix(), &protosService.SceneIntelligenceTask{
			Id:             8651918578347507712,
			IntelligenceId: 5924289872947740672,
			TaskImg:        "https://xxxx.jpg",
			TaskDesc:       "测试",
			TaskType:       iotconst.TASK_TYPE_DELAYED.ToInt(),
			FuncDesc:       "测试消息推送",
		})
		if err != nil {
			t.Error("消息通知失败", err)
		} else {
			t.Log("消息通知成功", iotutil.ToString(res))
		}
	}()
	select {}
	//
	////设备执行
	//res, err = execute.DeviceExecute(1, time.Now().Unix(), devIds, &protosService.SceneIntelligenceTask{
	//	Id:             8651918578347507712,
	//	IntelligenceId: 4511178014456709120,
	//	TaskImg:        "https://xxxx.jpg",
	//	TaskDesc:       "测试",
	//	TaskType:       iotconst.TASK_TYPE_DEVICE.ToInt(),
	//	ObjectId:       "dIIlDhrWFsyOkR",
	//	ObjectDesc:     "测试",
	//	FuncKey:        "switch",
	//	FuncDesc:       "开关",
	//	FuncValue:      "1",
	//})
	//if err != nil {
	//	t.Error("设备失败", err)
	//} else {
	//	t.Log("设备成功", iotutil.ToString(res))
	//}
	//
	////延时执行
	//res, err = execute.DelayedExecute(1, time.Now().Unix(), &protosService.SceneIntelligenceTask{
	//	Id:             8651918578347507712,
	//	IntelligenceId: 4511178014456709120,
	//	TaskImg:        "https://xxxx.jpg",
	//	TaskDesc:       "测试",
	//	TaskType:       iotconst.TASK_TYPE_DELAYED.ToInt(),
	//	FuncDesc:       "延时2分钟执行",
	//	FuncValue:      "00:02",
	//})
	//if err != nil {
	//	t.Error("延时失败", err)
	//} else {
	//	t.Log("延时成功", iotutil.ToString(res))
	//}
	//
	////关闭智能执行
	//res, err = execute.SceneIntelligenceExecute(1, time.Now().Unix(), &protosService.SceneIntelligenceTask{
	//	Id:             8651918578347507712,
	//	IntelligenceId: 4511178014456709120,
	//	ObjectId:       "3630329497569361920",
	//	ObjectDesc:     "智能",
	//	TaskImg:        "https://xxxx.jpg",
	//	TaskDesc:       "测试",
	//	TaskType:       iotconst.TASK_TYPE_INTELL.ToInt(),
	//	FuncDesc:       "关闭智能",
	//	FuncValue:      "1",
	//})
	//if err != nil {
	//	t.Error("场景开关", err)
	//} else {
	//	t.Log("场景开关", iotutil.ToString(res))
	//}

}

//模拟初始化场景数据
//obj := &models.SceneIntelligence{
//	Id:               1,
//	Type:             2,
//	Title:            "测试智能场景",
//	SortNo:           1,
//	EnableDisplay:    1,
//	Status:           1,
//	RunStatus:        1,
//	UserId:           1,
//	HomeId:           1,
//	ConditionMode:    1,
//	EffectTimeSwitch: 0,
//	EffectTimeDesc:   "",
//	EffectTimeWeeks:  "",
//	EffectTimeStart:  "",
//	EffectTimeEnd:    "",
//	Condition: []models.SceneIntelligenceConditionForm{
//		{
//			Id:                1,
//			IntelligenceId:    1,
//			ConditionType:     iotconst.CONDITION_TYPE_SATACHANGE.ToInt(),
//			Desc:              "设备状态变化测试",
//			DeviceDid:         "123123",
//			DevicePropKey:     "switch",
//			DevicePropCompare: 1,
//			DevicePropValue:   "1",
//		},
//		{
//			Id:             2,
//			IntelligenceId: 1,
//			ConditionType:  iotconst.CONDITION_TYPE_WEATHER.ToInt(),
//			Desc:           "天气变化测试",
//			WeatherCountry: "中国",
//			WeatherCity:    "长沙",
//			WeatherArea:    "湖南",
//			WeatherType:    iotconst.WEATHER_TYPE_TEMPERATURE.ToInt(),
//			WeatherValue:   "20",
//			WeatherCompare: 1,
//		},
//		{
//			Id:             3,
//			IntelligenceId: 1,
//			ConditionType:  iotconst.CONDITION_TYPE_TIMER.ToInt(),
//			Desc:           "定时执行测试",
//			TimerWeeks:     "1,2,3,4,5",
//			TimerValue:     "09:00",
//		},
//	},
//	//任务类型 1设备/2智能/3延时/4发送通知提醒/5群组
//	Task: []models.SceneIntelligenceTaskForm{
//		{
//			Id:             1,
//			IntelligenceId: 1,
//			TaskImg:        "#112322",
//			TaskDesc:       "测试",
//			TaskType:       iotconst.TASK_TYPE_DEVICE.ToInt(),
//			ObjectId:       "123123",
//			ObjectDesc:     "测试设备",
//			FuncKey:        "switch",
//			FuncDesc:       "开关",
//			FuncValue:      "0",
//		},
//	},
//}
//redisCmd := iotredis.GetClient().Set(context2.Background(), fmt.Sprintf("%d", obj.Id), iotutil.ToString(obj), 0)
//if redisCmd.Err() != nil {
//	fmt.Println("redis存储失败", redisCmd.Err().Error())
//}

func TestSyncMaps(t *testing.T) {
	var observerHasMap sync.Map = sync.Map{}
	observerHasMap.Store("123", DeviceObserver{id: "1", did: "1231231"})

	o, ok := observerHasMap.Load("123")
	if ok {
		t.Log("有值的")
		d := o.(DeviceObserver)
		t.Log(d)
	} else {
		t.Log(o)
	}

}

func TestChannel(t *testing.T) {
	config.InitTest("../../")
	cached.InitCache()
	serviceCfg := config.Global.Service
	//统一日志到服务的日志
	err := iotlogger.InitLog(serviceCfg.Logfile, "iot_device_service", serviceCfg.Loglevel)
	if err != nil {
		panic(err)
	}
	iotlogger.LogHelper.Info("Server start running")
	dbcnf := model.DBConfig{
		DBType:     config.Global.Database.Driver,
		ConnectStr: config.Global.Database.Connstr,
	}
	err = model.InitDB(dbcnf)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return
	}

	controlAckChan := make(chan bool, 0)
	go CheckControlResult2("test", "rxxxx", "test", controlAckChan)
	select {
	case msg := <-controlAckChan:
		if msg {
			t.Log("test")
		} else {
			t.Log("error")
		}
	}
}

func CheckControlResult2(messageId, productKey, deviceId string, result chan bool) {
	ctx := context.Background()
	//TestDeviceChan()
	ackCh := strings.Join([]string{iotconst.HKEY_ACK_DATA_PUB_PREFIX, productKey, deviceId}, ".")
	ackSub := cached.RedisStore.GetClient().PSubscribe(ctx, ackCh)
	defer ackSub.Close()

	ackChannel := ackSub.Channel()
	for {
		select {
		case msg := <-ackChannel:
			data := iotstruct.DeviceRedisData{}
			if err := json.Unmarshal([]byte(msg.Payload), &data); err != nil {
				iotlogger.LogHelper.Helper.Error("json unmarshal online error: ", err)
				continue
			}
			iotlogger.LogHelper.Infof("redis control ack sub data[%s,%s] ", data.MessageId, messageId, iotutil.ToString(data))
			if data.MessageId == messageId {
				result <- true
				break
			}
		case <-time.After(3 * time.Second): //超时3s
			iotlogger.LogHelper.Info("redis control ack sub timeout")
			result <- false
			break
		}
	}

	fmt.Println("testetestset")
}

func TestConvertSpec(t *testing.T) {
	convertSpec("09:09", "1,2,3,4,5,6,7", "Asia/Shanghai")
	t.Log("ok")
}
