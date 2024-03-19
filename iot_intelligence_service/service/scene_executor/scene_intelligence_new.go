package scene_executor

import (
	rules2 "cloud_platform/iot_intelligence_service/service/scene_executor/rules"
	"cloud_platform/iot_intelligence_service/service/scene_executor/valscene"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"fmt"
)

var weatherManagers *ObserverWeatherItems     //天气管理器
var deviceManagers *ObserverDeviceStatusItems //设备状态管理器
var timeManagers *ObserverTimerItems
var rules IntelligenceRule

// 场景规则对象
type IntelligenceRule struct {
}

func (s IntelligenceRule) InitIntelligence() {
	weatherManagers = new(ObserverWeatherItems)
	deviceManagers = new(ObserverDeviceStatusItems)
	timeManagers = new(ObserverTimerItems)

	//初始化天气和设备状态编号订阅
	weatherManagers.initSub()
	deviceManagers.initSub()

	s.initCreateSub()
}

//通过DB初始化场景规则

// CreateRule 创建场景规则
func (s *IntelligenceRule) CreateRule(id, desc string) {
	iotlogger.LogHelper.WithTag("Intelligence", "CreateRule").Infof("智能规则创建，id:%s, desc:%s", id, desc)
	clearRedis(id)
	//缓存获取场景数据
	intelligenceCmd := iotredis.GetClient().Get(context.Background(), iotconst.HKEY_INTELLIGENCE_DATA+id)
	if intelligenceCmd.Val() == "" {
		iotlogger.LogHelper.Error("智能场景数据未加入到缓存中")
		return
	}
	var intellObj *protosService.SceneIntelligence
	err := iotutil.JsonToStruct(intelligenceCmd.Val(), &intellObj)
	if err != nil {
		iotlogger.LogHelper.Error("抛出了异常，是否需要关闭当前场景?")
		return
	}
	if intellObj == nil {
		iotlogger.LogHelper.Error("未获取到场景数据", intellObj.Id)
		return
	}
	for _, condition := range intellObj.SceneIntelligenceCondition {
		conditionId := iotutil.ToString(condition.Id)
		switch iotconst.ConditionType(condition.ConditionType) {
		case iotconst.CONDITION_TYPE_WEATHER:
			iotlogger.LogHelper.WithTag("Intelligence", "Condition").Infof("创建天气条件，city:%s, type: %v, compare:%v, value:%v", condition.WeatherCity, condition.WeatherType, condition.WeatherCompare, condition.WeatherValue)
			obs := WeatherObserver{id: id, city: condition.WeatherCity, weatherType: condition.WeatherType, weatherCompare: condition.WeatherCompare, weatherValue: condition.WeatherValue}
			weatherManagers.register(obs)
			valscene.WeatherRuleBuilder.BuildRuleWithIncremental(genRuleAuto(id, desc, 10, conditionId))
			//天气创建规则后自动执行
			c := SceneIntelligenceCondition{}
			c.ConditionExecutor(id, "")
		case iotconst.CONDITION_TYPE_SATACHANGE:
			iotlogger.LogHelper.WithTag("Intelligence", "Condition").Infof("创建设备状态变化条件，did:%s, key: %v, compare:%v, value:%v", condition.DeviceDid, condition.DevicePropKey, condition.DevicePropCompare, condition.DevicePropValue)
			obs := DeviceObserver{id: id, did: condition.DeviceDid, funcKey: condition.DevicePropKey, funcValue: condition.DevicePropValue}
			deviceManagers.register(obs)

			valscene.DeviceRuleBuilder.BuildRuleWithIncremental(genRuleAuto(id, desc, 10, conditionId))
		case iotconst.CONDITION_TYPE_TIMER:
			iotlogger.LogHelper.WithTag("Intelligence", "Condition").Infof("创建定时任务，weeks:%v, value: %v", condition.TimerWeeks, condition.TimerValue)
			obs := TimerObserver{id: id, weekVal: condition.TimerWeeks, timeVal: condition.TimerValue, timezone: intellObj.Timezone, regionServerId: intellObj.RegionServerId}
			timeManagers.register(obs)

			valscene.TimerRuleBuilder.BuildRuleWithIncremental(genRuleAuto(id, desc, 10, conditionId))

		}
	}
}

func (s *IntelligenceRule) initCreateSub() {
	rules2.CreateRuleChan = make(chan rules2.CreateRuleChanData, 0)
	go func() {
		for {
			select {
			case job := <-rules2.CreateRuleChan:
				iotlogger.LogHelper.Info("chan sub data " + iotutil.ToString(job))
				if job.Status == 1 {
					s.CreateRule(job.Id, job.Desc)
				} else {
					s.DeleteRule(job.Id)
				}
			}
		}
	}()
}

// DeleteRule 删除场景规则
func (s *IntelligenceRule) DeleteRule(id ...interface{}) {
	iotlogger.LogHelper.Infof("智能规则删除，id:%v", id)
	idStr := []string{}
	for _, i := range id {
		idStr = append(idStr, iotutil.ToString(i))
	}
	clearRedis(idStr...)
	valscene.WeatherRuleBuilder.RemoveRules(idStr)
	valscene.DeviceRuleBuilder.RemoveRules(idStr)
	valscene.TimerRuleBuilder.RemoveRules(idStr)
}

//TODO 合并天气管理和设备状态管理

// genRuleAuto 通过智能场景生成规则
func genRuleAuto(id string, desc string, salience int32, conditionId string) string {
	rule1 := `
	rule "{{.id}}" "{{.desc}}" salience {{.salience}}
	begin
	executor.ConditionExecutor("{{.id}}", "{{.conditionId}}")
	end
	`
	ruleStr, err := setRuleParams(rule1, map[string]interface{}{
		"id":          id,
		"desc":        desc,
		"salience":    salience,
		"conditionId": conditionId,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	iotlogger.LogHelper.Info(ruleStr)
	return ruleStr
}
