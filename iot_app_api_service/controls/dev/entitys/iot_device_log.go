package entitys

type DeviceOperationFailLogRequest struct {
	DeviceId string `json:"deviceId"` //设备编号
	Type     int32  `json:"type"`     // 1-配网 2-OTA升级
	Content  string `json:"content"`  //上报内容json串
}

type AppFailLog struct {
	ProductKey string      `json:"productKey"` //可选，为空时是APP日志
	DeviceId   string      `json:"deviceId"`   //可选，为空时是APP日志
	Time       int64       `json:"time"`       //APP时间戳
	Type       int         `json:"type"`       //1-配网 2-OTA升级 3-局域网控制
	Code       int         `json:"code"`       //错误码
	Content    interface{} `json:"content"`    //错误内容，json串
}
