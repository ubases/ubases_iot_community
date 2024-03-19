package entitys

//增、删、改及查询返回
type TIotDeviceFunctionSetEntity struct {
	DeviceId       string `json:"deviceId"`       // 设备Id
	ProductKey     string `json:"productKey"`     // 产品Key
	FuncKey        string `json:"funcKey"`        // 功能Dpid
	FuncIdentifier string `json:"funcIdentifier"` // 功能标识符
	FuncValue      string `json:"funcValue"`      // 功能值
	CustomType     int32  `json:"customType"`     // 自定义的类型（1 = 物模型的属性  =2 物模型的值）
	CustomDesc     string `json:"customDesc"`     // 功能值描述
}
