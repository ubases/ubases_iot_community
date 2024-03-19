package models

//定义想要注入的结构体
type SceneIntelligence struct {
	Id               int64  `json:"id"`            // 唯一主键
	Type             int32  `json:"type"`          // =1 一键执行任务 =2 自动化场景任务
	Title            string `json:"title"`         // 标题
	SortNo           int32  `json:"sortNo"`        // 序号
	EnableDisplay    int32  `json:"enableDisplay"` // 启用首页显示
	Status           int32  `json:"status"`        // 状态（=1 启用 =2 禁用）
	RunTime          int64  `json:"runTime"`       // 运行时间
	RunStatus        int32  `json:"runStatus"`     // 运行状态 （ 1 待运行 2 运行中 2运行结束）
	UserId           int64  `json:"userId"`
	HomeId           int64  `json:"homeId"`        // 家庭编号
	ConditionMode    int32  `json:"conditionMode"` // 条件模式 =1 所有条件满足 =2 任意条件满足
	StyleIcon        string `json:"styleIcon"`
	StyleColor       string `json:"styleColor"`
	StyleImg         string `json:"styleImg"`         // 背景图片
	EffectTimeSwitch int32  `json:"effectTimeSwitch"` // 有效时间段开关
	EffectTimeDesc   string `json:"effectTimeDesc"`   // 时间描述
	EffectTimeWeeks  string `json:"effectTimeWeeks"`  // 周设置
	EffectTimeStart  string `json:"effectTimeStart"`  // 有效开始时间（例如：01:12）
	EffectTimeEnd    string `json:"effectTimeEnd"`    // 有效结束时间（例如：01:12）
	Condition        []SceneIntelligenceConditionForm
	Task             []SceneIntelligenceTaskForm
	AppKey           string `json:"appKey"`   // APP Key
	TenantId         string `json:"tenantId"` // 开发者租户编号
}

//智能场景条件
type SceneIntelligenceConditionForm struct {
	Id                int64  `json:"id"`                // 唯一主键
	IntelligenceId    int64  `json:"intelligenceId"`    // 场景ID
	ConditionType     int32  `json:"conditionType"`     // 条件类型
	Desc              string `json:"desc"`              // 天气描述
	WeatherCountry    string `json:"weatherCountry"`    // 国家
	WeatherCity       string `json:"weatherCity"`       // 城市
	WeatherArea       string `json:"weatherArea"`       // 区域
	WeatherType       int32  `json:"weatherType"`       // 天气类型
	WeatherValue      string `json:"weatherValue"`      // 天气值
	WeatherCompare    int32  `json:"weatherCompare"`    // 条件比较（1为等于 2 大于 ....)
	TimerWeeks        string `json:"timerWeeks"`        // 定时周设置，逗号分隔(例如：0,1,2,3,4)
	TimerValue        string `json:"timerValue"`        // 定时的值设置（01:33)
	DeviceDid         string `json:"deviceDid"`         // 设备ID
	DevicePropKey     string `json:"devicePropKey"`     // 设备功能的Key
	DevicePropCompare int32  `json:"devicePropCompare"` // 条件比较（1为等于 2 大于 ....)
	DevicePropValue   string `json:"devicePropValue"`   // 设备属性条件值
}

//智能场景任务
type SceneIntelligenceTaskForm struct {
	Id             int64  `json:"id"`             // 唯一主键
	IntelligenceId int64  `json:"intelligenceId"` // 智能场景Id
	TaskImg        string `json:"taskImg"`        // 任务图片(产品图片、智能图片、功能图标）
	TaskDesc       string `json:"taskDesc"`       // 任务描述
	TaskType       int32  `json:"taskType"`       // 任务类型（=1 延时 =2 设备执行 =3 场景开关）
	ObjectId       string `json:"objectId"`       // 对象ID（设备Id、场景Id）
	ObjectDesc     string `json:"objectDesc"`     // 对象的标题或者描述（设备名称、场景名称）
	FuncKey        string `json:"funcKey"`        // 执行功能Key
	FuncDesc       string `json:"funcDesc"`       // 冗余：功能描述
	FuncValue      string `json:"funcValue"`
}
