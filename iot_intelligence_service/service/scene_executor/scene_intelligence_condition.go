package scene_executor

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_intelligence_service/config"
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

// SceneIntelligenceCondition 智能场景执行器
type SceneIntelligenceCondition struct {
}

func (u *SceneIntelligenceCondition) ConditionByTask(intelligenceId string, taskId string) bool {
	//如果taskId，可作为
	return false
}

// ConditionExecutor 条件执行器
func (u *SceneIntelligenceCondition) ConditionExecutor(intelligenceId string, conditionId string) bool {

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
		weeks := u.GetWeekIndex(intellObj.EffectTimeWeeks)
		inWeek := strings.Index(weeks, weekStr)
		if currDt.After(endDt) || currDt.Before(startDt) || inWeek == -1 {
			//iotlogger.LogHelper.Infof("不在有效执行时间段，%s  %s", startDtStr, endDtStr)
			return false
		}
	}
	//iotlogger.LogHelper.WithTag("Id", intelligenceId).WithTag("conditionId", conditionId).Infof("开始条件判断")
	//读取设备详情数据
	resultFlag := make(models.ResultBools, 0)
	redisSaveQueue := new([]*RedisTempSet)
	for _, cond := range intellObj.SceneIntelligenceCondition {
		//外部conditionId，代表该条件已经满足
		//iotlogger.LogHelper.WithTag("Id", intelligenceId).WithTag("conditionId", conditionId).Infof("条件判断 " + iotutil.ToString(cond.Id))
		ct := iotconst.ConditionType(cond.ConditionType)
		if conditionId == iotutil.ToString(cond.Id) {
			//设备需要检查属性状态变化
			if ct == iotconst.CONDITION_TYPE_SATACHANGE {
				iotlogger.LogHelper.Debugf("开始检查设备状态变化,%v", intellObj.Id)
				resultFlag = append(resultFlag, u.DeviceStatusCheck(intellObj, cond, redisSaveQueue))
			} else {
				iotlogger.LogHelper.Debugf("条件发生变更并直接成功,%v", intellObj.Id)
				resultFlag = append(resultFlag, true)
			}
			continue
		}
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
		//如果为任意条件满足，同时当前设置条件中已有条件满足，则直接执行
		if iotconst.ConditionModel(intellObj.ConditionMode) == iotconst.CONDITION_MODEL_2 &&
			resultFlag.EveryOneTrue() {
			break
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
	iotlogger.LogHelper.WithTag("Id", intelligenceId).Infof("条件判断结果：%v", isExecutor)
	if isExecutor {
		iotlogger.LogHelper.Infof("智能任务开始执行 %d", intellObj.Id)
		//缓存存储
		setRedisBeforeData(redisSaveQueue)
		u.ExecuteTask(intellObj, 0)
		//如果是执行一次的场景任务，直接关闭智能场景。 （TODO 已经通过crond实现）
		//if isCloseIntell {
		//	execute.SceneIntelligenceClose(intellObj.Id)
		//}
	}
	return false
}

// WeatherCheck 天气检查
func (u *SceneIntelligenceCondition) WeatherCheck(intellObj *protosService.SceneIntelligence,
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
			newVal = math.Floor(fVal)
			if redisValCmd.Val() != "" {
				beforeVal = iotutil.ToFloat64(redisValCmd.Val())
			}
		} else {
			return false
		}

		//idArr := strings.Split(condition.WeatherValue, ",") //输入值：截取前端设置温度值 1|-39°C  运算符|温度
		//judge := iotutil.ToInt32(idArr[0])                  //输入值：类型判断 条件运算符 =1 大于 =2 等于 =3 小于
		inputValue, err := iotutil.ToFloat64Err(condition.WeatherValue) //iotutil.ToFloat64(strings.Split(idArr[1], "%")[0])
		if err != nil {
			iotlogger.LogHelper.Infof("传入的天气值异常, %v: %v", weatherType, inputValue)
			return false
		}
		inputValue = math.Floor(inputValue)
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
			newVal = math.Round(fVal)
			//if iotconst.WEATHER_TYPE_TEMPERATURE == weatherType {
			//	newVal = math.Round(iotutil.HToSTemperature(newVal))
			//}
			if redisValCmd.Val() != "" {
				beforeVal = iotutil.ToFloat64(redisValCmd.Val())
			}
		} else {
			return false
		}
		//inputValue := iotutil.ToFloat64(condition.WeatherValue) //iotutil.ToString(strings.Split(idArr[1], "%")[0])
		inputValue, err := iotutil.ToFloat64Err(condition.WeatherValue) //iotutil.ToFloat64(strings.Split(idArr[1], "%")[0])
		if err != nil {
			iotlogger.LogHelper.Infof("传入的天气值异常, %v: %v", weatherType, inputValue)
			return false
		}
		inputValue = math.Round(inputValue)
		switch condition.WeatherCompare {
		case 1:
			hasChange = newVal > inputValue && (beforeVal < inputValue || beforeVal == 0)
			isSetCache = newVal < inputValue
		case 2:
			hasChange = newVal == inputValue && (beforeVal != inputValue || beforeVal == 0)
			isSetCache = newVal != inputValue && newVal != beforeVal
		case 3:
			hasChange = newVal < inputValue && (beforeVal > inputValue || beforeVal == 0)
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
func (u *SceneIntelligenceCondition) DeviceStatusCheck(intellObj *protosService.SceneIntelligence,
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
func (u *SceneIntelligenceCondition) DelayedCheck(intellObj *protosService.SceneIntelligence,
	condition *protosService.SceneIntelligenceCondition, redisSaveQueue *[]*RedisTempSet) (runSuccess bool, isCloseIntell bool) {
	//只允许设备状态
	if !iotconst.CONDITION_TYPE_TIMER.Is(condition.ConditionType) {
		return false, false
	}
	//通过redis缓存数据获取
	weekStr := iotutil.ToString(time.Now().Weekday())
	timeStr := iotutil.DateTimeFormat(time.Now())
	weeks := u.GetWeekIndex(condition.TimerWeeks)
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
func (u *SceneIntelligenceCondition) ExecuteTask(intellObj *protosService.SceneIntelligence, resultId int64) (int64, bool) {
	//TODO 如何防止重复执行
	userId := intellObj.UserId
	devIds := new([]string)
	result := &model.TSceneIntelligenceResult{
		RunTime:        time.Now(),
		IntelligenceId: intellObj.Id,
		RunStatus:      1,
	}
	if resultId == 0 {
		resultId = iotutil.GetNextSeqInt64()
	}
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
				iotlogger.LogHelper.Info("控制")
				setResult(execute.DeviceExecute(userId, resultId, result.RunTime.Unix(), devIds, task))
			case iotconst.TASK_TYPE_INTELL:
				iotlogger.LogHelper.Info("场景开关")
				setResult(execute.SceneIntelligenceExecute(resultId, result.RunTime.Unix(), task))
			case iotconst.TASK_TYPE_DELAYED:
				iotlogger.LogHelper.Info("延时")
				result.RunStatus = 3 //延时执行
				execute.SetResult(result)
				setResult(execute.DelayedExecute(resultId, result.RunTime.Unix(), task))
			case iotconst.TASK_TYPE_SENDMSG:
				iotlogger.LogHelper.Info("通知")
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

// NewBuilder 初始化生成器
func NewBuilder() {
	iotlogger.LogHelper.Info("NewBuilder 初始化启动")
	//初始化定时器
	cron2.InitTimerCron()
	//清理天气redisKey
	ClearWeatherRedisKey()

	sceneExecutor := &SceneIntelligenceCondition{}
	dataContext := context.NewDataContext()
	//注入初始化的结构体（执行器）
	dataContext.Add("executor", sceneExecutor)
	dataContext.Add("flog", iotlogger.LogHelper)

	//init rule engine
	valscene.RuleBuilder = builder.NewRuleBuilder(dataContext)
	valscene.WeatherRuleBuilder = builder.NewRuleBuilder(dataContext)
	valscene.DeviceRuleBuilder = builder.NewRuleBuilder(dataContext)
	valscene.TimerRuleBuilder = builder.NewRuleBuilder(dataContext)
	valscene.Gengine = engine.NewGengine()

	do := iotmodel.GetDB().Model(model.TSceneIntelligence{}).Where(model.TSceneIntelligence{
		Status:         1,
		Type:           2,
		RegionServerId: config.Global.Service.ServerId,
	})
	var totalCount int64
	countTx := do.Count(&totalCount)
	if countTx.Error != nil {
		iotlogger.LogHelper.Error(countTx.Error)
		return
	}
	limit := 1000

	pageCount := int(math.Ceil(float64(totalCount) / float64(limit)))
	//Subscribe
	rules = IntelligenceRule{}
	rules.InitIntelligence()
	for page := 0; page < pageCount; page++ {
		offset := limit * page
		//加载规则
		var intelligences []struct {
			Id    int64  `gorm:"column:id;primaryKey" json:"id"`
			Title string `gorm:"column:title" json:"title"`
		}
		res := do.Offset(offset).Limit(limit).Scan(&intelligences)
		if res.Error != nil {
			continue
		}
		if len(intelligences) == 0 {
			continue
		}
		for _, item := range intelligences {
			rules.CreateRule(fmt.Sprintf("%v", item.Id), fmt.Sprintf("智能场景-%s", item.Title))
		}
	}
	//从redis中找到城市key并创建监听任务
	InitMonitorWeatherData()
	//初始化redis订阅（设备状态）
	go InitRedisSub()
	iotlogger.LogHelper.Info("NewBuilder 初始化方法调用结束")
}

// GetWeekIndex 星期索引值转换
func (u *SceneIntelligenceCondition) GetWeekIndex(weekStr string) string {
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
