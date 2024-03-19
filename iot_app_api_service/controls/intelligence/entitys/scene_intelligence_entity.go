package entitys

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	proto "cloud_platform/iot_proto/protos/protosService"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

// 智能场景Vo
type SceneIntelligenceVo struct {
	Id               int64        `json:"id,string"` // 唯一主键
	Type             int32        `json:"type"`      // =1 一键执行任务 =2 自动化场景任务
	Title            string       `json:"title"`     // 标题
	SortNo           int32        `json:"sortNo"`    // 序号
	EnableDisplay    int32        `json:"enabled"`   // 启用首页显示
	Status           bool         `json:"status"`    // 状态（=1 启用 =2 禁用）
	RunStatus        int32        `json:"runStatus"` // 运行状态 （ 1 待运行 2 运行中 2运行结束）
	UserId           int64        `json:"userId,string"`
	HomeId           int64        `json:"homeId,string"` // 家庭编号
	ConditionMode    int32        `json:"conditionMode"` // 条件模式 =1 所有条件满足 =2 任意条件满足
	StyleIcon        string       `json:"styleIcon"`
	StyleColor       string       `json:"styleColor"`
	StyleImg         string       `json:"styleImg"`         // 背景图片
	EffectTimeSwitch bool         `json:"effectTimeSwitch"` // 有效时间段开关
	EffectTimeDesc   string       `json:"effectTimeDesc"`   // 时间描述
	EffectTimeWeeks  string       `json:"effectTimeWeeks"`  // 周设置
	EffectTimeStart  string       `json:"effectTimeStart"`  // 有效开始时间（例如：01:12）
	EffectTimeEnd    string       `json:"effectTimeEnd"`    // 有效结束时间（例如：01:12）
	Condition        []*Condition `json:"wheres"`
	Task             []*Task      `json:"taskList"`
	FailureFlag      int32        `json:"failureFlag"` //=1 相关设备已离线 =2 相关设备已移除 =3 全部设备已离线 =4 全部设备已移除
}

// AppQueryProductForm query OpmProduct  form ;  if some field is required, create binding:"required" to tag by self
type SceneIntelligenceQueryForm struct {
	Id             int64  `json:"id,string"`       // 唯一主键
	Type           int32  `json:"type"`            // 一键执行任务 =2 自动化场景任务
	ResultId       int64  `json:"resultId,string"` // 执行任务结果
	UserId         int64  `json:"userId,string"`
	HomeId         int64  `json:"homeId,string"`         // 家庭编号
	IntelligenceId int64  `json:"intelligenceId,string"` // 家庭编号
	ProductKey     string `json:"productKey"`            //产品Key
	DeviceId       string `json:"devId"`                 //设备Id
	Order          int    `json:"order" form:"order"`    // example: orderMap[column]=desc
	Page           int    `json:"page" form:"page"`      //
	Limit          int    `json:"limit" form:"limit"`    //
	EnableDisplay  int32  `json:"enabled"`               // 启用首页显示
}

// 智能场景提交
type SceneIntelligenceForm struct {
	Id               int64  `json:"id,string"` // 唯一主键
	Type             int32  `json:"type"`      // =1 一键执行任务 =2 自动化场景任务
	Title            string `json:"title"`     // 标题
	SortNo           int32  `json:"sortNo"`    // 序号
	EnableDisplay    int32  `json:"enabled"`   // 启用首页显示
	Status           bool   `json:"status"`    // 状态（=1 启用 =2 禁用）
	RunStatus        int32  `json:"runStatus"` // 运行状态 （ 1 待运行 2 运行中 2运行结束）
	UserId           int64  `json:"userId,string"`
	HomeId           int64  `json:"homeId,string"` // 家庭编号
	ConditionMode    int32  `json:"conditionMode"` // 条件模式 =1 所有条件满足 =2 任意条件满足
	StyleIcon        string `json:"styleIcon"`
	StyleColor       string `json:"styleColor"`
	StyleImg         string `json:"styleImg"`         // 背景图片
	EffectTimeSwitch bool   `json:"effectTimeSwitch"` // 有效时间段开关
	EffectTimeDesc   string `json:"effectTimeDesc"`   // 时间描述
	EffectTimeWeeks  string `json:"effectTimeWeeks"`  // 周设置
	EffectTimeStart  string `json:"effectTimeStart"`  // 有效开始时间（例如：01:12）
	EffectTimeEnd    string `json:"effectTimeEnd"`    // 有效结束时间（例如：01:12）
	Condition        []SceneIntelligenceConditionForm
	Task             []SceneIntelligenceTaskForm
}

// 智能场景条件
type SceneIntelligenceConditionForm struct {
	Id                   int64       `json:"id,string"`                    // 唯一主键
	IntelligenceId       int64       `json:"intelligenceId,string"`        // 场景ID
	ConditionType        int32       `json:"conditionType"`                // 条件类型
	Desc                 string      `json:"desc"`                         // 天气描述
	WeatherCountry       string      `json:"weatherCountry"`               // 国家
	WeatherCity          string      `json:"weatherCity"`                  // 城市
	WeatherArea          string      `json:"weatherArea"`                  // 区域
	WeatherType          int32       `json:"weatherType"`                  // 天气类型
	WeatherValue         string      `json:"weatherValue"`                 // 天气值
	WeatherCompare       int32       `json:"weatherCompare"`               // 条件比较（1为等于 2 大于 ....)
	TimerWeeks           string      `json:"timerWeeks"`                   // 定时周设置，逗号分隔(例如：0,1,2,3,4)
	TimerValue           string      `json:"timerValue"`                   // 定时的值设置（01:33)
	DeviceDid            string      `json:"deviceDid"`                    // 设备ID
	DevicePropKey        string      `json:"devicePropKey"`                // 设备功能的Key
	DevicePropCompare    int32       `json:"devicePropCompare"`            // 条件比较（1为等于 2 大于 ....)
	DevicePropValue      interface{} `json:"devicePropValue"`              // 设备属性条件值
	DevicePropIdentifier string      `json:"devicePropIdentifier"`         // 设备属性条件值
	DevicePropMultiple   interface{} `json:"devicePropMultiple,omitempty"` // 设备属性倍数
	DevicePropDataType   interface{} `json:"devicePropDataType"`           // 设备属性类型
	DevicePropDesc       string      `json:"devicePropDesc"`               // 设备属描述
	ProductKey           string      `json:"productKey"`
	ProductId            int64       `json:"productId"`
	FailureFlag          int32       `json:"failureFlag"` //=1 相关设备已离线 =2 相关设备已移除 =3 全部设备已离线 =4 全部设备已移除
}

// 智能场景任务
type SceneIntelligenceTaskForm struct {
	Id             int64          `json:"id,string"`             // 唯一主键
	IntelligenceId int64          `json:"intelligenceId,string"` // 智能场景Id
	TaskImg        string         `json:"taskImg"`               // 任务图片(产品图片、智能图片、功能图标）
	TaskDesc       string         `json:"taskDesc"`              // 任务描述
	TaskType       int32          `json:"taskType"`              // 任务类型（=1 延时 =2 设备执行 =3 场景开关）
	ObjectId       string         `json:"objectId"`              // 对象ID（设备Id、场景Id）
	ObjectDesc     string         `json:"objectDesc"`            // 对象的标题或者描述（设备名称、场景名称）
	FuncKey        string         `json:"funcKey"`               // 执行功能Key
	FuncDesc       string         `json:"funcDesc"`              // 冗余：功能描述
	FuncIdentifier string         `json:"funcIdentifier"`        // 冗余：功能标识符
	FuncValue      string         `json:"funcValue"`
	ProductKey     string         `json:"productKey"`
	ProductId      int64          `json:"productId"`
	Functions      []TaskFunction `json:"functions"`
	FailureFlag    int32          `json:"failureFlag"` //=1 相关设备已离线 =2 相关设备已移除 =3 全部设备已离线 =4 全部设备已移除
}

// 执行任务功能
type TaskFunction struct {
	FuncCompare    int32       `json:"funcCompare"`
	FuncKey        string      `json:"funcKey"`
	FuncDesc       string      `json:"funcDesc"`
	ShowDesc       string      `json:"showDesc"` //显示描述，不参与保存
	FuncIdentifier string      `json:"funcIdentifier"`
	FuncValue      string      `json:"funcValue"`
	FuncValueDesc  string      `json:"funcValueDesc"`
	Multiple       interface{} `json:"multiple,omitempty"`
	DataType       string      `json:"dataType"`
}

// 智能场景日志
type SceneIntelligenceLogVo struct {
	Id        int64     `json:"id,string"` // 唯一主键
	ObjectId  int64     `json:"objectId,string"`
	HomeId    int64     `json:"homeId,string"`
	UserId    int64     `json:"userId,string"`
	Content   string    `json:"content"`
	ResultId  int32     `json:"resultId"`  // 执行结果编号（t_scene_intelligence_result.id）
	IsSuccess int32     `json:"isSuccess"` // 是否执行成功
	SceneType int32     `json:"sceneType"` // =1 一键执行 =2 自动场景
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// 智能场景结果
type SceneIntelligenceResultVo struct {
	Id             int64     `json:"id,string"`             // 唯一主键
	RunTime        time.Time `json:"runTime"`               // 运行时间
	IntelligenceId int64     `json:"intelligenceId,string"` // 智能场景编号
	RunStatus      int32     `json:"runStatus"`             // 运行状态 =1 执行中 = 2 执行成功 =3 执行失败
}

// 智能场景任务结果
type SceneIntelligenceResultTaskVo struct {
	Id             int64     `json:"id,string"` // 唯一主键
	StartTime      time.Time `json:"startTime"` // 运行时间
	EndTime        time.Time `json:"endTime"`
	IntelligenceId int64     `json:"intelligenceId,string"` // 智能场景编号
	IsSuccess      int32     `json:"isSuccess"`             // 是否成功
	ResultMsg      string    `json:"resultMsg"`             // 结果描述
	TaskId         int64     `json:"taskId,string"`         // 任务编号
	TaskImg        string    `json:"taskImg"`               // 任务图片(产品图片、智能图片、功能图标）
	TaskDesc       string    `json:"taskDesc"`              // 任务描述
	TaskType       int32     `json:"taskType"`              // 任务类型（=1 延时 =2 设备执行 =3 场景开关）
	ObjectId       string    `json:"objectId"`              // 对象ID（设备Id、场景Id）
	ObjectDesc     string    `json:"objectDesc"`            // 对象的标题或者描述（设备名称、场景名称）
	FuncKey        string    `json:"funcKey"`               // 执行功能Key
	FuncDesc       string    `json:"funcDesc"`              // 冗余：功能描述
	FuncResult     string    `json:"funcResult"`            // 功能结果
	FuncValue      string    `json:"funcValue"`
	ProductKey     string    `json:"productKey"`
	ProductId      int64     `json:"productId"`
}

// 兼容旧的智能场景提交
type OldSceneIntelligenceForm struct {
	Id             int64                  `json:"id,string"` // 唯一主键
	Type           int32                  `json:"type"`      // =1 一键执行任务 =2 自动化场景任务
	Title          string                 `json:"title"`     // 标题
	SortNo         int32                  `json:"sortno"`    // 序号
	EnableDisplay  int32                  `json:"enabled"`   // 启用首页显示
	Status         int32                  `json:"status"`    // 状态（=1 启用 =2 禁用）
	RunStatus      int32                  `json:"runStatus"` // 运行状态 （ 1 待运行 2 运行中 2运行结束）
	UserId         int64                  `json:"userId,string"`
	HomeId         int64                  `json:"homeId,string"` // 家庭编号
	ConditionMode  int32                  `json:"whereMode"`     // 条件模式 =1 所有条件满足 =2 任意条件满足
	Style          SceneIntelligenceStyle `json:"style"`
	TakeEffectTime TakeEffectTime         `json:"takeEffectTime"`
	Condition      []Condition            `json:"wheres"`
	Task           []Task                 `json:"taskList"`
	EntryId        int64                  `json:"entryId"`   // 任务cron编号 新结构不存在
	HomeName       string                 `json:"homeName"`  // 家庭名称 新结构不存在
	RunTime        int64                  `json:"runTime"`   // 运行时间 新结构不存在
	Showindex      bool                   `json:"showIndex"` // 显示在首页 true/false 新结构不存在
}

type SceneIntelligenceStyle struct {
	StyleIcon  string `json:"icon"`
	StyleColor string `json:"color"`
	StyleImg   string `json:"img"` // 背景图片
}

// 旧任务执行时间
type TakeEffectTime struct {
	EffectTimeSwitch bool   `json:"switch"`    // 有效时间段开关
	EffectTimeDesc   string `json:"desc"`      // 时间描述
	EffectTimeWeeks  string `json:"weeks"`     // 周设置
	EffectTimeStart  string `json:"startTime"` // 有效开始时间（例如：01:12）
	EffectTimeEnd    string `json:"endTime"`   // 有效结束时间（例如：01:12）
}

// 旧条件对象
type Condition struct {
	Id             int64       `json:"id,string"`             // 条件ID
	IntelligenceId int64       `json:"intelligenceId,string"` // 场景id
	Weather        Weather     `json:"weather"`               // 天气
	Timer          Timer       `json:"timer"`                 // 定时器
	Statechange    Statechange `json:"statechange"`           // 设备属性过滤
	ConditionType  int32       `json:"type"`                  // 条件类型

}

// 旧天气对象
type Weather struct {
	Desc           string      `json:"desc"`           // 天气描述
	WeatherCountry string      `json:"country"`        // 国家
	WeatherCity    string      `json:"city"`           // 城市
	WeatherArea    string      `json:"area"`           // 区域
	WeatherType    int32       `json:"type"`           // 天气类型
	WeatherValue   interface{} `json:"value"`          // 天气值
	WeatherCompare int32       `json:"weatherCompare"` // 条件比较（1为等于 2 大于 ....) 新结构才存在
}

// 旧定时器对象
type Timer struct {
	TimerWeeks string `json:"weeks"`    // 定时周设置，逗号分隔(例如：0,1,2,3,4)
	TimerValue string `json:"value"`    // 定时的值设置（01:33)
	CronSpec   string `json:"cronSpec"` // cron表达式 新结构不存在
	Daysmode   string `json:"daysMode"` // 0不重复 1指定天 2每天 新结构不存在
	Desc       string `json:"desc"`     // 描述 新结构不存在
	EntryId    int32  `json:"entryId"`  //  新结构不存在
}

// 旧设备条件对象
type Statechange struct {
	DeviceDid            string      `json:"objectId"`           // 设备ID
	DeviceDevId          string      `json:"devId"`              // 设备ID（兼容处理）
	DevicePropKey        string      `json:"pid"`                // 设备功能的Key
	DevicePropKeyDpId    string      `json:"dpId"`               // 设备功能的Key（兼容处理）
	DevicePropKeyDpid    string      `json:"dpid"`               // 设备功能的Key（兼容处理 dpid 小写）
	DevicePropCompare    int32       `json:"devicePropCompare"`  // 条件比较（1为等于 2 大于 ....)
	DevicePropValue      string      `json:"value"`              // 设备属性条件值
	DevicePropIdentifier string      `json:"identifier"`         // 设备属性条件值
	DevicePropDesc       string      `json:"funcDesc"`           // 设备属描述
	DevicePropMultiple   interface{} `json:"multiple,omitempty"` // 设备属倍数（int、double、float时候生效）
	DevicePropDataType   interface{} `json:"dataType"`           // 设备属数据类型
	Desc                 string      `json:"desc"`               // 描述 新结构不存在
	ProductKey           string      `json:"productKey"`
	ProductId            int64       `json:"productId,string"`
	ProductImg           string      `json:"proImg"`  // 任务图片(产品图片、智能图片、功能图标）
	ProductName          string      `json:"proName"` // 任务描述
	RoomName             string      `json:"roomName"`
	FailureFlag          int32       `json:"failureFlag"` //=1 设备已离线 =2 设备已移除
}

// 旧任务对象
type Task struct {
	Id             int64       `json:"taskId,string"`         // 唯一主键
	IntelligenceId int64       `json:"intelligenceId,string"` // 智能场景Id
	TaskImg        string      `json:"proImg"`                // 任务图片(产品图片、智能图片、功能图标）
	TaskDesc       string      `json:"proName"`               // 任务描述
	TaskType       int32       `json:"taskType"`              // 任务类型（=1 延时 =2 设备执行 =3 场景开关）
	ObjectId       string      `json:"objectId"`              // 对象ID（设备Id、场景Id）
	ObjectDesc     string      `json:"devDesc"`               // 对象的标题或者描述（设备名称、场景名称）
	FuncKey        string      `json:"funcCommon"`            // 执行功能Key
	FuncDesc       string      `json:"funcDesc"`              // 冗余：功能描述
	FuncValue      interface{} `json:"funcValue"`             // 功能值
	FuncResult     string      `json:"funcResult"`            // 功能结果
	FuncIdentifier string      `json:"funcIdentifier"`        // 冗余：功能描述
	Delayed        string      `json:"delayed"`               // 延时时间设置 新结构不存在
	RoomName       string      `json:"roomName"`              // 房间名称 新结构不存在
	Secretkey      string      `json:"secretKey"`             // 描述 设备秘钥 新结构不存在

	ProductKey  string         `json:"productKey"`
	ProductId   int64          `json:"productId,string"`
	Functions   []TaskFunction `json:"functions"`
	FailureFlag int32          `json:"failureFlag"` //=1 设备已离线 =2 设备已移除
}

// 新旧场景对象转换
func Intelligence_old2new(src *OldSceneIntelligenceForm) SceneIntelligenceForm {
	if src == nil {
		return SceneIntelligenceForm{}
	}
	//转换智能条件对象
	condition := Condition_old2new(src.Condition)
	//转换智能任务对象
	task := Task_old2new(src.Task)

	entitysObj := SceneIntelligenceForm{
		Id:               src.Id,
		Type:             src.Type,
		Title:            src.Title,
		SortNo:           src.SortNo,
		EnableDisplay:    src.EnableDisplay,
		Status:           src.Status == 1,
		RunStatus:        src.RunStatus,
		UserId:           src.UserId,
		HomeId:           src.HomeId,
		ConditionMode:    src.ConditionMode,
		StyleIcon:        src.Style.StyleIcon,
		StyleColor:       src.Style.StyleColor,
		StyleImg:         src.Style.StyleImg,
		EffectTimeSwitch: src.TakeEffectTime.EffectTimeSwitch,
		EffectTimeDesc:   src.TakeEffectTime.EffectTimeDesc,
		EffectTimeWeeks:  src.TakeEffectTime.EffectTimeWeeks,
		EffectTimeStart:  src.TakeEffectTime.EffectTimeStart,
		EffectTimeEnd:    src.TakeEffectTime.EffectTimeEnd,
		Condition:        condition,
		Task:             task,
	}
	return entitysObj
}

// 转换新旧智能条件对象
func Condition_old2new(odlConditions []Condition) []SceneIntelligenceConditionForm {
	if odlConditions == nil || len(odlConditions) == 0 {
		return nil
	}
	var conditions []SceneIntelligenceConditionForm
	for _, oldCondition := range odlConditions {
		condition := SceneIntelligenceConditionForm{
			Id:            oldCondition.Id,
			ConditionType: oldCondition.ConditionType,
		}
		switch iotconst.ConditionType(oldCondition.ConditionType) {
		case iotconst.CONDITION_TYPE_WEATHER:
			condition.Desc = oldCondition.Weather.Desc
			condition.WeatherCountry = oldCondition.Weather.WeatherCountry
			condition.WeatherCity = oldCondition.Weather.WeatherCity
			condition.WeatherArea = oldCondition.Weather.WeatherArea
			condition.WeatherType = oldCondition.Weather.WeatherType
			condition.WeatherValue = iotutil.ToString(oldCondition.Weather.WeatherValue)
			condition.WeatherCompare = oldCondition.Weather.WeatherCompare
		case iotconst.CONDITION_TYPE_SATACHANGE:
			condition.Desc = oldCondition.Statechange.Desc
			condition.DeviceDid = oldCondition.Statechange.DeviceDid
			if condition.DeviceDid == "" {
				condition.DeviceDid = oldCondition.Statechange.DeviceDevId
			}
			condition.DevicePropKey = oldCondition.Statechange.DevicePropKey
			//dpid兼容处理
			if oldCondition.Statechange.DevicePropKey == "" {
				condition.DevicePropKey = oldCondition.Statechange.DevicePropKeyDpId
				if oldCondition.Statechange.DevicePropKey == "" {
					condition.DevicePropKey = oldCondition.Statechange.DevicePropKeyDpid
				}
			}
			condition.DevicePropCompare = oldCondition.Statechange.DevicePropCompare
			condition.DevicePropValue = oldCondition.Statechange.DevicePropValue
			condition.DevicePropIdentifier = oldCondition.Statechange.DevicePropIdentifier
			condition.DevicePropDesc = oldCondition.Statechange.DevicePropDesc
			condition.ProductKey = oldCondition.Statechange.ProductKey
			condition.ProductId = oldCondition.Statechange.ProductId
		case iotconst.CONDITION_TYPE_TIMER:
			condition.Desc = oldCondition.Timer.Desc
			condition.TimerWeeks = oldCondition.Timer.TimerWeeks
			condition.TimerValue = oldCondition.Timer.TimerValue
		}
		conditions = append(conditions, condition)
	}
	return conditions
}

// 转换新旧智能任务对象
func Task_old2new(oldTasks []Task) []SceneIntelligenceTaskForm {
	if oldTasks == nil || len(oldTasks) == 0 {
		return nil
	}
	var tasks []SceneIntelligenceTaskForm
	for _, oldTask := range oldTasks {
		task := SceneIntelligenceTaskForm{
			Id:             oldTask.Id,
			TaskImg:        oldTask.TaskImg,
			TaskDesc:       oldTask.TaskDesc,
			TaskType:       oldTask.TaskType,
			ObjectId:       oldTask.ObjectId,
			ObjectDesc:     oldTask.ObjectDesc,
			FuncKey:        oldTask.FuncKey,
			FuncDesc:       oldTask.FuncDesc,
			FuncValue:      iotutil.ToString(oldTask.FuncValue),
			ProductKey:     oldTask.ProductKey,
			ProductId:      oldTask.ProductId,
			Functions:      oldTask.Functions,
			FuncIdentifier: oldTask.FuncIdentifier,
		}
		tasks = append(tasks, task)
	}
	return tasks
}

// 转换新旧智能任务对象
func Task_proto2old(lang, tenantId string, langMap map[string]string, oldTasks []*proto.SceneIntelligenceTask) ([]*Task, bool, int32) {
	var failureFlag int32 = 0
	if oldTasks == nil || len(oldTasks) == 0 {
		return nil, false, failureFlag
	}
	var tasks []*Task
	var hasOffline, hasOnline, hasRemove bool
	var hasDeviceTask bool
	for _, oldTask := range oldTasks {
		task := &Task{
			Id:         oldTask.Id,
			TaskImg:    oldTask.TaskImg,
			TaskDesc:   oldTask.TaskDesc,
			TaskType:   oldTask.TaskType,
			ObjectId:   oldTask.ObjectId,
			ObjectDesc: oldTask.ObjectDesc,
			FuncKey:    oldTask.FuncKey,
			FuncDesc:   oldTask.FuncDesc,
			FuncValue:  oldTask.FuncValue,
			ProductKey: oldTask.ProductKey,
			ProductId:  oldTask.ProductId,
			FuncResult: oldTask.FuncDesc,
			//Functions:  oldTask.Functions,
		}
		if task.TaskType == int32(iotconst.TASK_TYPE_DEVICE) {
			hasDeviceTask = true
			if oldTask.Functions != "" {
				json.Unmarshal([]byte(oldTask.Functions), &task.Functions)
				if task.Functions != nil {
					var funcKey, funcDesc, funcValue, funcResult []string
					for _, function := range task.Functions {
						funcKey = append(funcKey, function.FuncIdentifier)

						//功能键翻译
						langKey := fmt.Sprintf("%s_%s_%s_name", lang, task.ProductKey, function.FuncIdentifier)
						function.FuncDesc = iotutil.MapGetStringVal(langMap[langKey], function.FuncDesc)
						funcDesc = append(funcDesc, function.FuncDesc)

						//功能值翻译
						langKey = fmt.Sprintf("%s_%s_%s_%v_name", lang, task.ProductKey, function.FuncIdentifier, function.FuncValue)
						function.FuncValue = iotutil.MapGetStringVal(langMap[langKey], function.FuncValue)
						funcValue = append(funcValue, function.FuncValue)

						funcValueDesc := function.FuncValueDesc
						if !iotutil.IsNumeric(function.FuncValueDesc) {
							funcValueDesc = iotutil.MapGetStringVal(langMap[langKey], function.FuncValue)
						}
						funcResult = append(funcResult, fmt.Sprintf("%s：%s", function.FuncDesc, funcValueDesc))
					}
					task.FuncKey = strings.Join(funcKey, "、")
					task.FuncValue = strings.Join(funcValue, "、")
					task.FuncDesc = strings.Join(funcDesc, "、")
					task.FuncResult = strings.Join(funcResult, "、")
				}
			} else {
				//历史兼容
				if task.FuncKey != "" {
					langKey := fmt.Sprintf("%s_%s_%s_%s_name", lang, task.ProductKey, task.FuncKey, task.FuncValue)
					task.FuncDesc = iotutil.MapGetStringVal(langMap[langKey], task.FuncDesc)
				}
			}
			deviceInfos := iotredis.GetClient().HMGet(context.Background(), iotconst.HKEY_DEV_DATA_PREFIX+task.ObjectId, iotconst.FIELD_ONLINE)
			if deviceInfos.Err() == nil {
				list := deviceInfos.Val()
				onlineStatus := ""
				if len(list) > 0 {
					onlineStatus = iotutil.ToString(list[0])
				}
				if onlineStatus == "" {
					//设备已移除
					task.FailureFlag = 2
					hasRemove = true
				} else if onlineStatus == "offline" {
					//设备已离线
					task.FailureFlag = 1
					hasOffline = true
				} else {
					hasOnline = true
				}
			}
		}
		tasks = append(tasks, task)
	}
	if hasDeviceTask {
		//=1 相关设备已离线 =2 相关设备已移除 =3 全部设备已离线 =4 全部设备已移除
		if hasOnline && hasOffline && hasRemove {
			failureFlag = 1
		} else if hasOnline && hasOffline && !hasRemove {
			failureFlag = 1 //有在线 有离线
		} else if hasOnline && !hasOffline && hasRemove {
			failureFlag = 2 //有在线 有移除
		} else if !hasOnline && hasOffline && hasRemove {
			failureFlag = 3 //无在线 有离线 有移除
		} else if !hasOnline && hasOffline && !hasRemove {
			failureFlag = 3 //无在线 有离线 无移除
		} else if !hasOnline && !hasOffline && hasRemove {
			failureFlag = 4 //无在线 无离线 有移除
		} else if !hasOnline && !hasOffline && !hasRemove {
			failureFlag = 4 //无在线 无离线 无移除
		}
	}
	return tasks, hasDeviceTask, failureFlag
}

// 检查条件中的设备数据是否有效
func Condition_proto2old(lang, tenantId string, langMap map[string]string, oldConditions []*proto.SceneIntelligenceCondition) (bool, int32) {
	var failureFlag int32 = -1
	if oldConditions == nil || len(oldConditions) == 0 {
		return false, failureFlag
	}
	var hasOffline, hasOnline, hasRemove bool
	var hasDeviceC = false
	for _, c := range oldConditions {
		if c.ConditionType == int32(iotconst.CONDITION_TYPE_SATACHANGE) {
			onlineStatus, _ := iotredis.GetClient().HGet(context.Background(), iotconst.HKEY_DEV_DATA_PREFIX+c.DeviceDid, "onlineStatus").Result()
			if onlineStatus == "" {
				hasRemove = true
			} else if onlineStatus == "offline" {
				hasOffline = true
			} else {
				hasOnline = true
			}
			hasDeviceC = true
		}
	}
	//有设备条件才需要判断，无设备直接返回-1
	if hasDeviceC {
		//=1 相关设备已离线 =2 相关设备已移除 =3 全部设备已离线 =4 全部设备已移除
		if hasOnline && hasOffline && hasRemove {
			failureFlag = 1
		} else if hasOnline && hasOffline && !hasRemove {
			failureFlag = 1 //有在线 有离线
		} else if hasOnline && !hasOffline && hasRemove {
			failureFlag = 2 //有在线 有移除
		} else if !hasOnline && hasOffline && hasRemove {
			failureFlag = 3 //无在线 有离线 有移除
		} else if !hasOnline && hasOffline && !hasRemove {
			failureFlag = 3 //无在线 有离线 无移除
		} else if !hasOnline && !hasOffline && hasRemove {
			failureFlag = 4 //无在线 无离线 有移除
		} else if !hasOnline && !hasOffline && !hasRemove {
			failureFlag = 4 //无在线 无离线 无移除
		}
	}
	return hasDeviceC, failureFlag
}

// 智能场景from对象转poto对象
func Intelligence_e2pb(src *SceneIntelligenceForm) (*proto.SceneIntelligence, error) {
	if src == nil {
		return nil, errors.New("参数异常")
	}
	//条件对象转换
	condition := Condition_e2pb(src.Condition)
	//任务对象转换
	task, err := Task_e2pb(src.Task)
	if err != nil {
		return nil, err
	}
	res := proto.SceneIntelligence{
		Id:                         src.Id,
		Type:                       src.Type,
		Title:                      src.Title,
		SortNo:                     src.SortNo,
		EnableDisplay:              src.EnableDisplay,
		Status:                     iotutil.IfInt32(src.Status, 1, 2),
		RunStatus:                  src.RunStatus,
		UserId:                     src.UserId,
		HomeId:                     src.HomeId,
		ConditionMode:              src.ConditionMode,
		StyleIcon:                  src.StyleIcon,
		StyleColor:                 src.StyleColor,
		StyleImg:                   src.StyleImg,
		EffectTimeSwitch:           iotutil.IfInt32(src.EffectTimeSwitch, 1, 2),
		EffectTimeDesc:             src.EffectTimeDesc,
		EffectTimeWeeks:            src.EffectTimeWeeks,
		EffectTimeStart:            src.EffectTimeStart,
		EffectTimeEnd:              src.EffectTimeEnd,
		SceneIntelligenceCondition: condition,
		SceneIntelligenceTask:      task,
	}
	return &res, nil
}

// 智能场景条件对象转poto对象
func Condition_e2pb(srcs []SceneIntelligenceConditionForm) []*proto.SceneIntelligenceCondition {
	if srcs == nil || len(srcs) == 0 {
		return nil
	}
	var entitysObj []*proto.SceneIntelligenceCondition
	for _, src := range srcs {
		entityObj := proto.SceneIntelligenceCondition{
			Id:                   src.Id,
			IntelligenceId:       src.IntelligenceId,
			ConditionType:        src.ConditionType,
			Desc:                 src.Desc,
			WeatherCountry:       src.WeatherCountry,
			WeatherCity:          src.WeatherCity,
			WeatherArea:          src.WeatherArea,
			WeatherType:          src.WeatherType,
			WeatherValue:         src.WeatherValue,
			WeatherCompare:       src.WeatherCompare,
			TimerWeeks:           src.TimerWeeks,
			TimerValue:           src.TimerValue,
			DeviceDid:            src.DeviceDid,
			DevicePropKey:        src.DevicePropKey,
			DevicePropCompare:    src.DevicePropCompare,
			DevicePropValue:      iotutil.ToString(src.DevicePropValue),
			DevicePropIdentifier: src.DevicePropIdentifier,
			DevicePropDesc:       src.DevicePropDesc,
			ProductKey:           src.ProductKey,
			ProductId:            src.ProductId,
		}
		entitysObj = append(entitysObj, &entityObj)
	}
	return entitysObj
}

// 智能场景任务对象转poto对象
func Task_e2pb(srcs []SceneIntelligenceTaskForm) (entitysObj []*proto.SceneIntelligenceTask, err error) {
	if srcs == nil || len(srcs) == 0 {
		return nil, errors.New("请设置场景任务")
	}
	for _, src := range srcs {
		r := proto.SceneIntelligenceTask{
			Id:             src.Id,
			IntelligenceId: src.IntelligenceId,
			TaskImg:        src.TaskImg,
			TaskDesc:       src.TaskDesc,
			TaskType:       src.TaskType,
			ObjectId:       src.ObjectId,
			ObjectDesc:     src.ObjectDesc,
			FuncKey:        src.FuncKey,
			FuncDesc:       src.FuncDesc,
			FuncValue:      src.FuncValue,
			ProductKey:     src.ProductKey,
			ProductId:      src.ProductId,
			FuncIdentifier: src.FuncIdentifier,
		}
		if src.Functions != nil && len(src.Functions) > 0 {
			r.Functions = iotutil.ToString(src.Functions)
		}
		entitysObj = append(entitysObj, &r)
	}
	return entitysObj, nil
}

// 智能场景poto对象转form对象
func Intelligence_pb2e(src *proto.SceneIntelligence) *SceneIntelligenceForm {
	if src == nil {
		return nil
	}
	//条件对象转换
	condition := Condition_pb2e(src.SceneIntelligenceCondition)
	//任务对象转换
	task := Task_pb2e(src.SceneIntelligenceTask)
	entitysObj := SceneIntelligenceForm{
		Id:               src.Id,
		Type:             src.Type,
		Title:            src.Title,
		SortNo:           src.SortNo,
		EnableDisplay:    src.EnableDisplay,
		Status:           src.Status == 1,
		RunStatus:        src.RunStatus,
		UserId:           src.UserId,
		HomeId:           src.HomeId,
		ConditionMode:    src.ConditionMode,
		StyleIcon:        src.StyleIcon,
		StyleColor:       src.StyleColor,
		StyleImg:         src.StyleImg,
		EffectTimeSwitch: src.EffectTimeSwitch == 1,
		EffectTimeDesc:   src.EffectTimeDesc,
		EffectTimeWeeks:  src.EffectTimeWeeks,
		EffectTimeStart:  src.EffectTimeStart,
		EffectTimeEnd:    src.EffectTimeEnd,
		Condition:        condition,
		Task:             task,
	}
	return &entitysObj
}

// 智能场景poto对象转form对象
func Intelligence_pb2eVo(lang, tenantId string, langMap map[string]string, src *proto.SceneIntelligence) *SceneIntelligenceVo {
	if src == nil {
		return nil
	}
	entitysObj := SceneIntelligenceVo{
		Id:               src.Id,
		Type:             src.Type,
		Title:            src.Title,
		SortNo:           src.SortNo,
		EnableDisplay:    src.EnableDisplay,
		Status:           src.Status == 1,
		RunStatus:        src.RunStatus,
		UserId:           src.UserId,
		HomeId:           src.HomeId,
		ConditionMode:    src.ConditionMode,
		StyleIcon:        src.StyleIcon,
		StyleColor:       src.StyleColor,
		StyleImg:         src.StyleImg,
		EffectTimeSwitch: src.EffectTimeSwitch == 1,
		EffectTimeDesc:   src.EffectTimeDesc,
		EffectTimeWeeks:  src.EffectTimeWeeks,
		EffectTimeStart:  src.EffectTimeStart,
		EffectTimeEnd:    src.EffectTimeEnd,
	}
	var failureFlag int32 = 0
	var t_failureFlag int32 = 0
	var c_failureFlag int32 = -1
	var hasDevice = false
	if src.SceneIntelligenceTask != nil {
		var hasTaskDev bool
		entitysObj.Task, hasTaskDev, t_failureFlag = Task_proto2old(lang, tenantId, langMap, src.SceneIntelligenceTask)
		if hasTaskDev {
			hasDevice = true
		}
	}
	if src.SceneIntelligenceCondition != nil {
		var hasConnDev bool
		hasConnDev, c_failureFlag = Condition_proto2old(lang, tenantId, langMap, src.SceneIntelligenceCondition)
		if hasConnDev {
			hasDevice = true
		}
	}
	//如果有设备才需要返回failureFlag
	if hasDevice {
		if c_failureFlag == -1 {
			failureFlag = t_failureFlag
		} else {
			//=1 相关设备已离线 =2 相关设备已移除 =3 全部设备已离线 =4 全部设备已移除
			if t_failureFlag == 1 || c_failureFlag == 1 {
				failureFlag = 1
			} else if t_failureFlag == 2 || c_failureFlag == 2 {
				failureFlag = 2
			} else if t_failureFlag == 3 && c_failureFlag == 3 {
				failureFlag = 3
			} else if t_failureFlag == 4 && c_failureFlag == 4 {
				failureFlag = 4
			} else if t_failureFlag == 4 && c_failureFlag == 4 {
				failureFlag = 4
			}
		}
	}
	entitysObj.FailureFlag = failureFlag

	return &entitysObj
}

// 智能场景条件对象转poto对象
func Condition_pb2e(srcs []*proto.SceneIntelligenceCondition) []SceneIntelligenceConditionForm {
	if srcs == nil || len(srcs) == 0 {
		return nil
	}
	var entitysObj []SceneIntelligenceConditionForm
	for _, src := range srcs {
		entityObj := SceneIntelligenceConditionForm{
			Id:                   src.Id,
			IntelligenceId:       src.IntelligenceId,
			ConditionType:        src.ConditionType,
			Desc:                 src.Desc,
			WeatherCountry:       src.WeatherCountry,
			WeatherCity:          src.WeatherCity,
			WeatherArea:          src.WeatherArea,
			WeatherType:          src.WeatherType,
			WeatherValue:         src.WeatherValue,
			WeatherCompare:       src.WeatherCompare,
			TimerWeeks:           src.TimerWeeks,
			TimerValue:           src.TimerValue,
			DeviceDid:            src.DeviceDid,
			DevicePropKey:        src.DevicePropKey,
			DevicePropCompare:    src.DevicePropCompare,
			DevicePropValue:      src.DevicePropValue,
			DevicePropIdentifier: src.DevicePropIdentifier,
			DevicePropDesc:       src.DevicePropDesc,
			ProductKey:           src.ProductKey,
			ProductId:            src.ProductId,
		}
		entitysObj = append(entitysObj, entityObj)
	}
	return entitysObj
}

// 智能场景任务对象转poto对象
func Task_pb2e(srcs []*proto.SceneIntelligenceTask) []SceneIntelligenceTaskForm {
	if srcs == nil || len(srcs) == 0 {
		return nil
	}
	var entitysObj []SceneIntelligenceTaskForm
	for _, src := range srcs {
		entityObj := SceneIntelligenceTaskForm{
			Id:             src.Id,
			IntelligenceId: src.IntelligenceId,
			TaskImg:        src.TaskImg,
			TaskDesc:       src.TaskDesc,
			TaskType:       src.TaskType,
			ObjectId:       src.ObjectId,
			ObjectDesc:     src.ObjectDesc,
			FuncKey:        src.FuncKey,
			FuncDesc:       src.FuncDesc,
			FuncValue:      src.FuncValue,
			ProductKey:     src.ProductKey,
			ProductId:      src.ProductId,
			FuncIdentifier: src.FuncIdentifier,
		}
		if src.Functions != "" {
			json.Unmarshal([]byte(src.Functions), &entityObj.Functions)
		}
		entitysObj = append(entitysObj, entityObj)
	}
	return entitysObj
}

// 新旧场景对象转换
func Intelligence_new2old(src *SceneIntelligenceForm, tenantId, lang string, deviceRoomMap map[string]string) OldSceneIntelligenceForm {
	if src == nil {
		return OldSceneIntelligenceForm{}
	}
	//转换智能条件对象
	condition := Condition_new2old(src.Condition, tenantId, lang, deviceRoomMap)
	//转换智能任务对象
	task := Task_new2old(src.Task, deviceRoomMap)

	entitysObj := OldSceneIntelligenceForm{
		Id:            src.Id,
		Type:          src.Type,
		Title:         src.Title,
		SortNo:        src.SortNo,
		EnableDisplay: src.EnableDisplay,
		Status:        iotutil.IfInt32(src.Status, 1, 2),
		RunStatus:     src.RunStatus,
		UserId:        src.UserId,
		HomeId:        src.HomeId,
		ConditionMode: src.ConditionMode,
		Style: SceneIntelligenceStyle{
			StyleIcon:  src.StyleIcon,
			StyleColor: src.StyleColor,
			StyleImg:   src.StyleImg,
		},
		TakeEffectTime: TakeEffectTime{
			EffectTimeSwitch: src.EffectTimeSwitch,
			EffectTimeDesc:   src.EffectTimeDesc,
			EffectTimeWeeks:  src.EffectTimeWeeks,
			EffectTimeStart:  src.EffectTimeStart,
			EffectTimeEnd:    src.EffectTimeEnd,
		},
		Condition: condition,
		Task:      task,
	}
	return entitysObj
}

// 转换新旧智能条件对象
func Condition_new2old(newConditions []SceneIntelligenceConditionForm, tenantId, lang string, deviceRoomMap map[string]string) []Condition {
	if newConditions == nil || len(newConditions) == 0 {
		return nil
	}
	var conditions []Condition
	for _, new := range newConditions {
		if new.WeatherValue == "" {
			new.WeatherValue = "0"
		}
		condition := Condition{
			Id:             new.Id,
			IntelligenceId: new.IntelligenceId,
			ConditionType:  new.ConditionType,
		}
		switch iotconst.ConditionType(new.ConditionType) {
		case iotconst.CONDITION_TYPE_WEATHER:
			condition.Weather = Weather{
				Desc:           new.Desc,
				WeatherCountry: new.WeatherCountry,
				WeatherCity:    new.WeatherCity,
				WeatherArea:    new.WeatherArea,
				WeatherType:    new.WeatherType,
				WeatherValue:   iotutil.ToInt32(new.WeatherValue),
				WeatherCompare: new.WeatherCompare,
			}
			condition.Weather.Desc = new.Desc
		case iotconst.CONDITION_TYPE_SATACHANGE:
			condition.Statechange = Statechange{
				DeviceDid:            new.DeviceDid,
				DeviceDevId:          new.DeviceDid,
				DevicePropKey:        new.DevicePropKey,
				DevicePropKeyDpId:    new.DevicePropKey, //兼容处理
				DevicePropKeyDpid:    new.DevicePropKey, //兼容处理
				DevicePropIdentifier: new.DevicePropIdentifier,
				DevicePropCompare:    new.DevicePropCompare,
				DevicePropValue:      iotutil.ToString(new.DevicePropValue),
				ProductKey:           new.ProductKey,
				ProductId:            new.ProductId,
				DevicePropDesc:       new.DevicePropDesc,
				DevicePropMultiple:   new.DevicePropMultiple,
				DevicePropDataType:   new.DevicePropDataType,
				RoomName:             deviceRoomMap[new.DeviceDid],
				FailureFlag:          new.FailureFlag,
			}

			condition.Statechange.ProductName = controls.GetDeviceName(new.DeviceDid)

			proCached := controls.ProductCachedData{}
			productInfo, _ := proCached.GetProduct(new.ProductKey)
			if productInfo != nil {
				condition.Statechange.ProductImg = productInfo.ImageUrl
				//如果名称一样，则读取产品的名称，支持翻译显示，否则直接已用户的名称显示
				if condition.Statechange.ProductName == "" {
					condition.Statechange.ProductName = productInfo.Name
					//处理名称翻译
					cacheKey := fmt.Sprintf("%s_%s", tenantId, iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_PRODUCT_NAME)
					langMap, err := iotredis.GetClient().HGetAll(context.Background(), cacheKey).Result()
					if err == nil {
						condition.Statechange.ProductName = iotutil.MapGetStringVal(langMap[fmt.Sprintf("%s_%s_name", lang, new.ProductKey)], productInfo.Name)
					}
				}
			}
			condition.Statechange.Desc = new.Desc

		case iotconst.CONDITION_TYPE_TIMER:
			condition.Timer = Timer{
				TimerWeeks: new.TimerWeeks,
				TimerValue: new.TimerValue,
			}
			condition.Timer.Desc = new.Desc
		}
		conditions = append(conditions, condition)
	}
	return conditions
}

// 转换新旧智能任务对象
func Task_new2old(newTasks []SceneIntelligenceTaskForm, deviceRoomMap map[string]string) []Task {
	if newTasks == nil || len(newTasks) == 0 {
		return nil
	}
	var tasks []Task
	for _, new := range newTasks {
		task := Task{
			Id:             new.Id,
			IntelligenceId: new.IntelligenceId,
			TaskImg:        new.TaskImg,
			TaskDesc:       new.TaskDesc,
			TaskType:       new.TaskType,
			ObjectId:       new.ObjectId,
			ObjectDesc:     new.ObjectDesc,
			FuncKey:        new.FuncKey,
			FuncDesc:       new.FuncDesc,
			FuncValue:      new.FuncValue,
			ProductKey:     new.ProductKey,
			ProductId:      new.ProductId,
			Functions:      new.Functions,
			FailureFlag:    new.FailureFlag,
		}
		if task.ObjectId != "" && iotconst.TaskType(task.TaskType) == iotconst.TASK_TYPE_DEVICE {
			task.RoomName = deviceRoomMap[task.ObjectId]
		}
		tasks = append(tasks, task)
	}
	return tasks
}

// 任务日志proto转实体
func ResultTask_pb2e(src *proto.SceneIntelligenceResultTask) *SceneIntelligenceResultTaskVo {
	if src == nil {
		return nil
	}
	entity := SceneIntelligenceResultTaskVo{
		Id:             src.Id,
		StartTime:      src.StartTime.AsTime(),
		EndTime:        src.EndTime.AsTime(),
		IntelligenceId: src.IntelligenceId,
		IsSuccess:      src.IsSuccess,
		ResultMsg:      src.ResultMsg,
		TaskId:         src.TaskId,
		TaskImg:        src.TaskImg,
		TaskDesc:       src.TaskDesc,
		TaskType:       src.TaskType,
		ObjectId:       src.ObjectId,
		ObjectDesc:     src.ObjectDesc,
		ProductKey:     src.ProductKey,
		FuncKey:        src.FuncKey,
		FuncDesc:       src.FuncDesc,
		FuncValue:      src.FuncValue,
		FuncResult:     src.FuncDesc,
	}

	return &entity
}

type TlsInfo struct {
	All map[string]map[string]string
}

// 获取倍数和数据类型
func (s *TlsInfo) GetTlsInfo(productKey string) (*TlsInfo, error) {
	if s.All == nil {
		s.All = make(map[string]map[string]string)
	}
	var err error
	s.All[productKey], err = iotredis.GetClient().HGetAll(context.Background(), fmt.Sprintf("product_data_%v", productKey)).Result()
	return s, err
}

// 获取倍数和数据类型
func (s *TlsInfo) GetMultipleAndDataType(productKey string, dpid interface{}) (multiple interface{}, dataType string) {
	if s.All == nil {
		s.All = make(map[string]map[string]string)
	}
	if _, ok := s.All[productKey]; !ok {
		var err error
		s.All[productKey], err = iotredis.GetClient().HGetAll(context.Background(), fmt.Sprintf("product_data_%v", productKey)).Result()
		if err != nil {
			return
		}
	}
	funcs, ok := s.All[productKey][fmt.Sprintf("tls_%v", dpid)]
	if !ok {
		return
	}
	funcsMap, err := iotutil.JsonToMapErr(funcs)
	if err != nil {
		return
	}
	//有倍数的需要获取倍数获取倍数
	//{"custom":0,"dataSpecs":"{\"dataType\":\"INT\",\"max\":\"720\",\"min\":\"0\",\"step\":\"1\",\"unit\":\"min\"}","dataSpecsList":"","dataType":"INT","dpid":8,"identifier":"mist_countdown_remain","name":"喷雾倒计时剩余","rwFlag":"READ"}
	if v, ok := funcsMap["dataSpecs"]; ok && v != "" {
		dataSpecs, err := iotutil.JsonToMapErr(iotutil.ToString(v))
		if err != nil {
			return
		}
		if m, ok := dataSpecs["multiple"]; ok && m != "" {
			multiple, _ = iotutil.ToInt32Err(m)
		}
	}
	dataType = iotutil.ToString(funcsMap["dataType"])
	return
}
