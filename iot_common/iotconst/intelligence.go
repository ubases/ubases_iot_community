package iotconst

type TaskType int
type ConditionType int
type ConditionModel int
type WeatherType int

func (c ConditionType) Is(val int32) bool {
	return c == ConditionType(val)
}
func (c ConditionType) ToInt() int32 {
	return int32(c)
}
func (c WeatherType) ToInt() int32 {
	return int32(c)
}
func (c ConditionModel) ToInt() int32 {
	return int32(c)
}
func (c TaskType) ToInt() int32 {
	return int32(c)
}

var (
	//执行状态
	RUN_SUCESS    string = "执行成功"
	RUN_CONTINUTE string = "执行中"
	RUN_ERORR     string = "异常"
	RUN_FAIL      string = "执行失败"

	//任务类型 1设备/2智能/3延时/4发送通知提醒/5群组
	TASK_TYPE_DEVICE  TaskType = 1 //1设备
	TASK_TYPE_INTELL  TaskType = 2 //2智能
	TASK_TYPE_DELAYED TaskType = 3 //3延时
	TASK_TYPE_SENDMSG TaskType = 4 //4发送通知提醒
	TASK_TYPE_GROUP   TaskType = 5 //5群组

	//条件类型
	CONDITION_TYPE_CLICK      ConditionType = 1  //一件执行
	CONDITION_TYPE_WEATHER    ConditionType = 2  //气候变化
	CONDITION_TYPE_TIMER      ConditionType = 3  //定时
	CONDITION_TYPE_SATACHANGE ConditionType = 4  //设备状态变化
	CONDITION_TYPE_POSITION   ConditionType = 5  //位置变化
	CONDITION_TYPE_CLOSE      ConditionType = -1 //关闭

	//条件模式
	CONDITION_MODEL_1 ConditionModel = 1 //满足所有
	CONDITION_MODEL_2 ConditionModel = 2 //满足任意一条

	//天气类型
	WEATHER_TYPE_TEMPERATURE WeatherType = 1 //温度
	WEATHER_TYPE_HUMIDITY    WeatherType = 2 //湿度
	WEATHER_TYPE_WEATHER     WeatherType = 3 //天气
	WEATHER_TYPE_PM25        WeatherType = 4 //PM2.5
	WEATHER_TYPE_QUALITY     WeatherType = 5 //空气质量
	WEATHER_TYPE_SUN         WeatherType = 6 //日落日出
	WEATHER_TYPE_WINDSPEED   WeatherType = 7 //风速

	WEATHER_VALUE_0 int32 = 0 // "other"
	WEATHER_VALUE_1 int32 = 1 // "晴天"
	WEATHER_VALUE_2 int32 = 2 // "阴天"
	WEATHER_VALUE_3 int32 = 3 // "雨天"
	WEATHER_VALUE_4 int32 = 4 // "雪天"
	WEATHER_VALUE_5 int32 = 5 // "霾天"
)
