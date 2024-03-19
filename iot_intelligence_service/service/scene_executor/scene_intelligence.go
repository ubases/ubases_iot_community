package scene_executor

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	cron2 "cloud_platform/iot_intelligence_service/service/scene_executor/cron"
	"cloud_platform/iot_intelligence_service/service/scene_executor/execute"
	"cloud_platform/iot_intelligence_service/service/scene_executor/models"
	"cloud_platform/iot_intelligence_service/service/scene_executor/valscene"
	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_device/model"
	"cloud_platform/iot_proto/protos/protosService"
	context2 "context"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/bilibili/gengine/builder"
	"github.com/bilibili/gengine/context"
	"github.com/bilibili/gengine/engine"
)

// SceneIntelligenceExecutor 智能场景执行器
type SceneIntelligenceExecutor struct {
}

// GetWeekIndex 星期索引值转换
func GetWeekIndex(weekStr string) string {
	newWeeks := []string{}
	weeks := strings.Split(weekStr, ",")
	for _, w := range weeks {
		strW := iotutil.ToString(w)
		if w == "7" {
			strW = "0"
		}
		newWeeks = append(newWeeks, strW)
	}
	return strings.Join(newWeeks, ",")
}

func (u *SceneIntelligenceExecutor) ConditionByTask(intelligenceId string, taskId string) bool {
	//如果taskId，可作为
	return false
}

// ConditionExecutor 条件执行器
func (u *SceneIntelligenceExecutor) ConditionExecutor(intelligenceId string) bool {

	//Redis读取智能任务详情
	ctx := context2.Background()
	intelligenceCmd := iotredis.GetClient().Get(ctx, iotconst.HKEY_INTELLIGENCE_DATA+intelligenceId)
	if intelligenceCmd.Val() == "" {
		iotlogger.LogHelper.Error("智能场景数据未加入到缓存中")
		RemoveRule(intelligenceId)
		return false
	}
	var intellObj *protosService.SceneIntelligence
	err := iotutil.JsonToStruct(intelligenceCmd.Val(), &intellObj)
	if err != nil {
		iotlogger.LogHelper.Error("抛出了异常，是否需要关闭当前场景?")
		RemoveRule(intellObj.Id)
		return false
	}
	if intellObj == nil {
		iotlogger.LogHelper.Error("未获取到场景数据", intellObj.Id)
		RemoveRule(intellObj.Id)
		return false
	}

	if intellObj.Status != 1 {
		iotlogger.LogHelper.Info("场景已关闭")
		RemoveRule(intellObj.Id)
		return false
	}

	//时间验证
	if intellObj.EffectTimeSwitch != 1 {
		currDt := time.Now()
		weekStr := iotutil.ToString(int(currDt.Weekday())) //当前星期序号
		startDtStr := fmt.Sprintf("%s %s", iotutil.DateFormat(currDt), intellObj.EffectTimeStart)
		startDt, _ := iotutil.GetTimeByStrToMintue(startDtStr)
		endDtStr := fmt.Sprintf("%s %s", iotutil.DateFormat(currDt), intellObj.EffectTimeEnd)
		endDt, _ := iotutil.GetTimeByStrToMintue(endDtStr)
		weeks := GetWeekIndex(intellObj.EffectTimeWeeks)
		inWeek := strings.Index(weeks, weekStr)
		if currDt.After(endDt) || currDt.Before(startDt) || inWeek == -1 {
			//iotlogger.LogHelper.Infof("不在有效执行时间段，%s  %s", startDtStr, endDtStr)
			return false
		}
	}

	//读取设备详情数据
	resultFlag := make(models.ResultBools, 0)
	redisSaveQueue := new([]*RedisTempSet)
	//isCloseIntell := false
	//resultFlag := ResultCondition{}.Def(true, types)
	for _, cond := range intellObj.SceneIntelligenceCondition {
		switch iotconst.ConditionType(cond.ConditionType) {
		case iotconst.CONDITION_TYPE_SATACHANGE:
			resultFlag = append(resultFlag, u.DeviceStatusCheck(intellObj, cond, redisSaveQueue))
		case iotconst.CONDITION_TYPE_TIMER:
			isSuc := false
			isSuc, _ = u.DelayedCheck(intellObj, cond, redisSaveQueue)
			resultFlag = append(resultFlag, isSuc)
		case iotconst.CONDITION_TYPE_WEATHER:
			resultFlag = append(resultFlag, u.WeatherCheck(intellObj, cond, redisSaveQueue))
		}
	}
	isExecutor := false
	if len(resultFlag) > 0 {
		switch iotconst.ConditionModel(intellObj.ConditionMode) {
		case iotconst.CONDITION_MODEL_1:
			isExecutor = resultFlag.AllTrue()
		case iotconst.CONDITION_MODEL_2:
			isExecutor = resultFlag.EveryOneTrue()
		}
	}
	if isExecutor {
		iotlogger.LogHelper.Infof("智能任务开始执行 %d", intellObj.Id)
		//缓存存储
		setRedisBeforeData(redisSaveQueue)
		u.ExecuteTask(intellObj, 0)
		//如果是执行一次的场景任务，直接关闭智能场景。（TODO 已经通过crond实现）
		//if isCloseIntell {
		//	execute.SceneIntelligenceClose(intellObj.Id)
		//}
	}
	return false
}

// WeatherCheck 天气检查
func (u *SceneIntelligenceExecutor) WeatherCheck(intellObj *protosService.SceneIntelligence,
	condition *protosService.SceneIntelligenceCondition, redisSaveQueue *[]*RedisTempSet) bool {
	//只允许设备状态
	if !iotconst.CONDITION_TYPE_WEATHER.Is(condition.ConditionType) {
		return false
	}
	if condition.WeatherValue == "" {
		return false
	}
	if condition.WeatherCity == "" {
		RemoveRule(intellObj.Id)
		iotlogger.LogHelper.Error("城市名称为空，任务停止")
		return false
	}
	//通过redis缓存数据获取
	ctx := context2.Background()
	//iotconst.HKEY_WEATHER_DATA+
	weatherCmd := iotredis.GetClient().Get(ctx, iotconst.HKEY_WEATHER_DATA+ConvertCityName(condition.WeatherCity))
	weatherInfo := make(map[string]interface{})
	if weatherCmd.Val() != "" {
		err := iotutil.JsonToStruct(weatherCmd.Val(), &weatherInfo)
		if err != nil {
			return false
		}
	}
	weatherKey := valscene.WeatherType[condition.WeatherType]
	hasChange := false
	isSetCache := false
	redisKey := fmt.Sprintf("scene_%d_weather_%d_before", intellObj.Id, condition.WeatherType)
	weatherType := iotconst.WeatherType(condition.WeatherType)
	switch weatherType {
	case iotconst.WEATHER_TYPE_WEATHER, iotconst.WEATHER_TYPE_SUN:
		redisValCmd := iotredis.GetClient().Get(context2.Background(), redisKey)
		var newVal float64 = 0
		var beforeVal float64 = 0
		if val, ok := weatherInfo[weatherKey]; ok {
			fVal, err := iotutil.ToFloat64Err(val)
			if err != nil {
				iotlogger.LogHelper.Infof("天气缓存数据异常, %v: %v", weatherType, val)
				return false
			}
			newVal = fVal
			if redisValCmd.Val() != "" {
				beforeVal = iotutil.ToFloat64(redisValCmd.Val())
			}
		} else {
			return false
		}

		//idArr := strings.Split(condition.WeatherValue, ",") //输入值：截取前端设置温度值 1|-39°C  运算符|温度
		//judge := iotutil.ToInt32(idArr[0])                  //输入值：类型判断 条件运算符 =1 大于 =2 等于 =3 小于
		inputValue := iotutil.ToFloat64(condition.WeatherValue) //iotutil.ToFloat64(strings.Split(idArr[1], "%")[0])

		//condition.WeatherCompare 1为大于 2 等于 3 小于 ....
		switch condition.WeatherCompare {
		case 2:
			hasChange = newVal == inputValue && (beforeVal != inputValue || beforeVal == 0)
			isSetCache = beforeVal != inputValue && newVal != beforeVal
		}
		if hasChange {
			isSetCache = true
			iotlogger.LogHelper.Infof("%d属性变化条件满足，当前天气类型%d，值为%v", intellObj.Id, condition.WeatherType, newVal)
		}
		//设置天气缓存
		if isSetCache {
			//iotlogger.LogHelper.Infof("%d重新设置天气缓存【%d]，值为%v", intellObj.Id, condition.WeatherType, newVal)
			//iotredis.GetClient().Set(context2.Background(), redisKey, newVal, 0)
			*redisSaveQueue = append(*redisSaveQueue, setKeyValue(redisKey, "", newVal, 0))
		}
	case iotconst.WEATHER_TYPE_TEMPERATURE, iotconst.WEATHER_TYPE_HUMIDITY, iotconst.WEATHER_TYPE_PM25, iotconst.WEATHER_TYPE_WINDSPEED:
		redisValCmd := iotredis.GetClient().Get(context2.Background(), redisKey)
		var newVal float64 = 0
		var beforeVal float64 = 0
		//当前的天气信息能获取到，并且before天气为空 ||
		if val, ok := weatherInfo[weatherKey]; ok {
			fVal, err := iotutil.ToFloat64Err(val)
			if err != nil {
				iotlogger.LogHelper.Infof("天气缓存数据异常, %v: %v", weatherType, val)
				return false
			}
			newVal = fVal
			if iotconst.WEATHER_TYPE_TEMPERATURE == weatherType {
				newVal = math.Round(iotutil.HToSTemperature(newVal))
			}
			if redisValCmd.Val() != "" {
				beforeVal = iotutil.ToFloat64(redisValCmd.Val())
			}
		} else {
			return false
		}
		inputValue := iotutil.ToFloat64(condition.WeatherValue) //iotutil.ToString(strings.Split(idArr[1], "%")[0])
		switch condition.WeatherCompare {
		case 1:
			hasChange = newVal >= inputValue && (beforeVal < inputValue || beforeVal == 0)
			isSetCache = newVal < inputValue
		case 2:
			hasChange = newVal == inputValue && (beforeVal != inputValue || beforeVal == 0)
			isSetCache = newVal != inputValue && newVal != beforeVal
		case 3:
			hasChange = newVal <= inputValue && (beforeVal > inputValue || beforeVal == 0)
			isSetCache = beforeVal > inputValue
		}
		if hasChange {
			isSetCache = true
			iotlogger.LogHelper.Infof("%d属性变化条件满足，当前天气类型【%d]，值为%v", intellObj.Id, condition.WeatherType, newVal)
		}
		//设置天气缓存
		if isSetCache {
			//iotlogger.LogHelper.Infof("%d重新设置天气缓存【%d]，值为%v", intellObj.Id, condition.WeatherType, newVal)
			//iotredis.GetClient().Set(context2.Background(), redisKey, newVal, 0)
			*redisSaveQueue = append(*redisSaveQueue, setKeyValue(redisKey, "", newVal, 0))
		}
	}
	//发送变化
	if hasChange {
		iotlogger.LogHelper.Infof("天气符合")
		return true
	}
	return false
}

// DeviceStatusCheck 设备状态检查
func (u *SceneIntelligenceExecutor) DeviceStatusCheck(intellObj *protosService.SceneIntelligence,
	condition *protosService.SceneIntelligenceCondition, redisSaveQueue *[]*RedisTempSet) bool {
	//只允许设备状态
	if !iotconst.CONDITION_TYPE_SATACHANGE.Is(condition.ConditionType) {
		return false
	}
	//通过redis缓存数据获取
	//redis存储设备属性
	//读取设备详情数据
	deviceCmd := iotredis.GetClient().HGetAll(context2.Background(), iotconst.HKEY_DEV_DATA_PREFIX+condition.DeviceDid)
	deviceInfo := deviceCmd.Val()
	if deviceCmd.Err() != nil {
		iotlogger.LogHelper.Error(deviceCmd.Err())
		return false
	}
	if currVal, ok := deviceInfo[condition.DevicePropKey]; ok {
		redisKey := fmt.Sprintf("scene_%d_status_before", intellObj.Id)
		//beforeVal := deviceInfo[condition.DevicePropKey+"_before"]
		beforeValCmd := iotredis.GetClient().HGet(context2.Background(), redisKey, condition.DevicePropKey)

		if GetTestIntelligenceId(iotutil.ToString(intellObj.Id)) {
			iotlogger.LogHelper.Infof("信息对比：id：%d, currVal: %s, before: %s, key: %s", intellObj.Id, currVal, beforeValCmd.Val(), condition.DevicePropValue)
			iotlogger.LogHelper.Infof("信息对比：属性是否变化[old-%v:curr-%v]：%v, 是否符合条件[cond-%v:curr-%v]：%v", beforeValCmd.Val(), currVal, beforeValCmd.Val() != currVal, condition.DevicePropValue, currVal, condition.DevicePropValue == currVal)
		}
		// 如何判断属性变化
		//1为等于 2 大于 ....  condition.DevicePropCompare 暂时不需要大于和小于
		//保证不会重复执行
		if beforeValCmd.Val() == currVal {
			//设备状态无变化
			return false
		}
		hasChange := condition.DevicePropValue == currVal
		if hasChange {
			//条件满足
			iotlogger.LogHelper.Infof("%d属性变化条件满足，设备当前状态为%s", intellObj.Id, currVal)
			*redisSaveQueue = append(*redisSaveQueue, setKeyValue(redisKey, condition.DevicePropKey, currVal, 0))
			return true
		} else {
			iotlogger.LogHelper.Debugf("属性变化条件未满足，清理缓存数据的值 %d", intellObj.Id)
			//如果出现其它值，则清除设备执行的缓存数据
			iotredis.GetClient().HSet(context2.Background(), redisKey, condition.DevicePropKey, "")
			return false
		}
	}
	return false
}

// DelayedCheck 定时检查
func (u *SceneIntelligenceExecutor) DelayedCheck(intellObj *protosService.SceneIntelligence,
	condition *protosService.SceneIntelligenceCondition, redisSaveQueue *[]*RedisTempSet) (runSuccess bool, isCloseIntell bool) {
	//只允许设备状态
	if !iotconst.CONDITION_TYPE_TIMER.Is(condition.ConditionType) {
		return false, false
	}
	//通过redis缓存数据获取
	weekStr := iotutil.ToString(time.Now().Weekday())
	timeStr := iotutil.DateTimeFormat(time.Now())
	weeks := GetWeekIndex(condition.TimerWeeks)
	isInWeek := weeks == "" || strings.Index(weeks, weekStr) != -1
	if isInWeek && timeStr == condition.TimerValue {
		redisKey := fmt.Sprintf("scene_%d_delayed_before", intellObj.Id)
		beforeTimeCmd := iotredis.GetClient().Get(context2.Background(), redisKey)
		if beforeTimeCmd.Val() == timeStr {
			runSuccess, isCloseIntell = false, false
			return
		}
		//记录redis缓存1分钟，确保当前分钟内不会执行第二次。
		//iotredis.GetClient().Set(context2.Background(), redisKey, timeStr, 1*time.Minute)
		*redisSaveQueue = append(*redisSaveQueue, setKeyValue(redisKey, "", timeStr, 1))

		iotlogger.LogHelper.Infof("%d定时条件满足，当前时间%s", intellObj.Id, timeStr)
		//weeks == ""  执行一次的逻辑，需要关闭场景任务。
		runSuccess, isCloseIntell = true, weeks == ""
		return
	}
	return false, false
}

// ExecuteTask 执行任务
func (u *SceneIntelligenceExecutor) ExecuteTask(intellObj *protosService.SceneIntelligence, resultId int64) (int64, bool) {
	devIds := new([]string)
	result := &model.TSceneIntelligenceResult{
		RunTime:        time.Now(),
		IntelligenceId: intellObj.Id,
		RunStatus:      1,
	}
	if resultId == 0 {
		resultId = iotutil.GetNextSeqInt64()
	}
	userId := intellObj.UserId
	result.Id = resultId
	//执行结果
	//记录执行时间
	err := execute.SetResult(result)
	if err != nil {
		iotlogger.LogHelper.Errorf("初始化写入执行结果错误，%s", err.Error())
		return 0, false
	}
	var (
		resInt     int32  = 0
		resMsg     string = ""
		hasSuccess bool   = false //存在成功步骤
		hasError   bool   = false //存在失败步骤
		setResult         = func(taskRes *model.TSceneIntelligenceResultTask, theErr error) {
			if theErr != nil {
				resInt = 2
				resMsg = theErr.Error()
				hasError = true
			} else {
				if taskRes.IsSuccess == 1 {
					hasSuccess = true
				} else {
					resInt = taskRes.IsSuccess
					resMsg = taskRes.ResultMsg
					hasError = true
				}
			}
		}
	)
	if intellObj.SceneIntelligenceTask == nil || len(intellObj.SceneIntelligenceTask) == 0 {
		resInt = 2
		resMsg = "智能场景的任务列表为空"
	} else {
		//设置结果
		for _, task := range intellObj.SceneIntelligenceTask {
			switch iotconst.TaskType(task.TaskType) {
			case iotconst.TASK_TYPE_DEVICE:
				iotlogger.LogHelper.Debug("控制")
				setResult(execute.DeviceExecute(userId, resultId, result.RunTime.Unix(), devIds, task))
			case iotconst.TASK_TYPE_INTELL:
				iotlogger.LogHelper.Debug("场景开关")
				setResult(execute.SceneIntelligenceExecute(resultId, result.RunTime.Unix(), task))
			case iotconst.TASK_TYPE_DELAYED:
				iotlogger.LogHelper.Debug("延时")
				result.RunStatus = 3 //延时执行
				execute.SetResult(result)
				setResult(execute.DelayedExecute(resultId, result.RunTime.Unix(), task))
			case iotconst.TASK_TYPE_SENDMSG:
				iotlogger.LogHelper.Debug("通知")
				setResult(execute.NoticeExecute(execute.NoticeParams{
					HomeId:           intellObj.HomeId,
					UserId:           []int64{},
					IntelligenceName: intellObj.Title,
					AppKey:           intellObj.AppKey,
					TenantId:         intellObj.TenantId,
				}, resultId, result.RunTime.Unix(), task))
			}
		}
	}
	result.RunStatus = resInt
	if hasSuccess && hasError {
		result.RunStatus = 4 //部分成功
	} else if hasSuccess && !hasError {
		result.RunStatus = 1 //完全执行成功
	} else {
		result.RunStatus = 2 //执行失败
	}
	err = execute.SetResult(result)
	if err != nil {
		iotlogger.LogHelper.Errorf("初始化写入执行结果错误，%s", err.Error())
		return 0, false
	}
	//记录日志
	execute.WriteLog(&model.TSceneIntelligenceLog{
		ObjectId:          intellObj.Id,
		HomeId:            intellObj.HomeId,
		UserId:            intellObj.UserId,
		Content:           resMsg,
		ResultId:          result.Id,
		IsSuccess:         result.RunStatus,
		SceneType:         intellObj.Type,
		IntelligenceId:    intellObj.Id,
		IntelligenceTitle: intellObj.Title,
	})
	return result.Id, true
}

// genRule 通过智能场景生成规则
func genRule(id string, desc string, salience int32, temp string) string {
	//rule1 := `
	//rule "{{.id}}" "{{.desc}}" salience {{.salience}}
	//begin
	//if executor.ConditionCheck("{{.id}}") {
	//	executor.ExecuteTask("{{.id}}", "执行了智能任务! {{.temp}} {{.id}}")
	//}
	//end
	//`
	rule1 := `
	rule "{{.id}}" "{{.desc}}" salience {{.salience}}
	begin
	executor.ConditionExecutor("{{.id}}")
	end
	`
	ruleStr, err := setRuleParams(rule1, map[string]interface{}{
		"id":       id,
		"desc":     desc,
		"salience": salience,
		"temp":     temp,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	iotlogger.LogHelper.Info(ruleStr)
	return ruleStr
}

func clearRedis(ids ...string) {
	defer iotutil.PanicHandler("清除redis异常", ids)
	redisKeys := []string{}
	for _, id := range ids {
		redisKeys = append(redisKeys, fmt.Sprintf("scene_%s_status_before", id))
		redisKeys = append(redisKeys, fmt.Sprintf("scene_%s_delayed_before", id))
		redisKeys = append(redisKeys, fmt.Sprintf("scene_%s_weather_1_before", id))
		redisKeys = append(redisKeys, fmt.Sprintf("scene_%s_weather_2_before", id))
		redisKeys = append(redisKeys, fmt.Sprintf("scene_%s_weather_3_before", id))
		redisKeys = append(redisKeys, fmt.Sprintf("scene_%s_weather_4_before", id))
		redisKeys = append(redisKeys, fmt.Sprintf("scene_%s_weather_5_before", id))
		redisKeys = append(redisKeys, fmt.Sprintf("scene_%s_weather_6_before", id))
		redisKeys = append(redisKeys, fmt.Sprintf("scene_%s_weather_7_before", id))
	}
	iotredis.GetClient().Del(context2.Background(), redisKeys...)
}

// CreateRule 创建规则
func CreateRule(id string, desc string, salience int32) {
	iotlogger.LogHelper.Infof("智能规则创建，id:%s, desc:%s", id, desc)
	clearRedis(id)
	valscene.RuleBuilder.BuildRuleWithIncremental(genRule(id, desc, 10, ""))
}

// RemoveRule 移除规则
func RemoveRule(id interface{}) {
	iotlogger.LogHelper.Infof("智能规则删除，id:%v", id)
	idStr := iotutil.ToString(id)
	clearRedis(idStr)
	valscene.RuleBuilder.RemoveRules([]string{idStr})
}

// RemoveRules 批量移除规则
func RemoveRules(id ...interface{}) {
	idStr := []string{}
	for _, i := range id {
		idStr = append(idStr, iotutil.ToString(i))
	}
	iotlogger.LogHelper.Infof("智能规则删除，ids:%v", idStr)
	clearRedis(idStr...)
	valscene.RuleBuilder.RemoveRules(idStr)
}

// InitBuilder 初始化生成器
func InitBuilder() {
	//初始化定时器
	cron2.InitTimerCron()
	ClearWeatherRedisKey()

	sceneExecutor := &SceneIntelligenceExecutor{}

	dataContext := context.NewDataContext()
	//注入初始化的结构体（执行器）
	dataContext.Add("executor", sceneExecutor)
	dataContext.Add("flog", iotlogger.LogHelper)

	//init rule engine
	valscene.RuleBuilder = builder.NewRuleBuilder(dataContext)
	valscene.Gengine = engine.NewGengine()

	//加载规则
	var intelligences []struct {
		Id    int64  `gorm:"column:id;primaryKey" json:"id"`
		Title string `gorm:"column:title" json:"title"`
	}

	//TODO 考虑分页读取初始化数据
	db := iotmodel.GetDB().Model(model.TSceneIntelligence{}).Where(model.TSceneIntelligence{
		Status: 1,
		Type:   2,
	}).Find(&intelligences)
	if db.Error != nil {
		iotlogger.LogHelper.Error(db.Error)
		return
	}
	//Subscribe
	rules = IntelligenceRule{}
	for _, item := range intelligences {
		valscene.RuleBuilder.BuildRuleWithIncremental(genRule(fmt.Sprintf("%d", item.Id), fmt.Sprintf("智能场景-%s", item.Title), 10, ""))
	}
	//从redis中找到城市key并创建监听任务
	InitMonitorWeatherData()
	//sceneExecutor.ConditionExecutor("4511178014456709120")
	monitorRules()
}

// monitorRules 监控规则
func monitorRules() {
	id, _ := cron2.CronCtx.AddFunc("*/1 * * * * *", func() {
		//start := time.Now().UnixNano()
		//执行规则
		err := valscene.Gengine.Execute(valscene.RuleBuilder, true)
		//end := time.Now().UnixNano()
		if err != nil {
			iotlogger.LogHelper.Errorf("execute rule error: %v", err)
		}
		//iotlogger.LogHelper.Infof("execute rule cost %d ns", end-start)
	})
	iotlogger.LogHelper.Infof("创建了定时任务：id:%d", id)

	//模拟修改场景
	//cron2.CronCtx.AddFunc("*/5 * * * * *", func() {
	//	for i := 1; i < 3; i++ {
	//		ruleBuilder.BuildRuleWithIncremental(genRule(fmt.Sprintf("%d", i), fmt.Sprintf("测试规则%d", i), 10, "修改了哦。"))
	//	}
	//})
}
