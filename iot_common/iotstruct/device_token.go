package iotstruct

// DeviceNetworkTokenCacheModel 设备配网Token缓存数据接口
type DeviceNetworkTokenCacheModel struct {
	UserId       int64                   `json:"userId"`
	HomeId       int64                   `json:"homeId"`
	ProductId    int64                   `json:"productId"`
	ProductKey   string                  `json:"productKey"`
	ProductName  string                  `json:"productName"`
	UserName     string                  `json:"userName"`
	Account      string                  `json:"account"`
	AppName      string                  `json:"appName"`
	AppKey       string                  `json:"appKey"`
	HomeName     string                  `json:"homeName"`
	Devices      []string                `json:"devices"`
	DevicesMap   map[string]DeviceResult `json:"devicesMap"`
	Lat          float64                 `json:"lat"`
	Lng          float64                 `json:"lng"`
	Country      string                  `json:"country"`
	Province     string                  `json:"province"`
	City         string                  `json:"city"`
	District     string                  `json:"district"`
	TenantId     string                  `json:"tenantId"`
	RegionId       string                  `json:"regionId"`
	RegionServerId string                  `json:"regionServerId"`
	DeviceNature int32                   `json:"deviceNature"`
}

type DeviceResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type MqttToNatsData struct {
	ProductKey string      `json:"productKey"` //产品Key
	DeviceId   string      `json:"deviceId"`   //设备Key
	Gid        string      `json:"gid"`        //网关编号,预留
	Time       int64       `json:"time"`       //数据采集时间,时间戳，单位秒
	Topic      string      `json:"topic"`      //原topic
	Payload    interface{} `json:"payload"`    //发布的数据，可序列化为json串
	Ns         string      `json:"ns"`         //MQTT消息头中的ns
	Name       string      `json:"name"`       //MQTT消息头中的name
	MID        string      `json:"mid"`        //消息Id
	Ver        string      `json:"ver"`        //消息Id
	From       string      `json:"from"`       //网关编号,预留
}

//对应主题iotconst.HKEY_DATA_PUB_PREFIX
type DeviceRedisData struct {
	ProductKey string      `json:"productKey"` //产品Key
	DeviceId   string      `json:"deviceId"`   //设备Key
	MessageId  string      `json:"messageId"`  //消息id
	Ns         string      `json:"ns"`         //MQTT消息头中的ns
	Name       string      `json:"name"`       //MQTT消息头中的name
	From       string      `json:"from"`       //网关编号,预留
	Time       int64       `json:"time"`       //数据采集时间,时间戳，单位秒
	Data       interface{} `json:"data"`       //发布的数据，可序列化为json串
}

type UpdateType int

var (
	UPDATE_TYPE_CHANGE_FAMILY UpdateType = 0 //切换家庭
	UPDATE_TYPE_REMOVE_DEVICE UpdateType = 1 //移除设备
)

// 用户切换家庭，增删改设备，需要发布相关消息
type DeviceRedisUpdate struct {
	UserId     string     `json:"userId"`     //用户id
	UpdateType UpdateType `json:"updateType"` //数据更新的类型 =0 家庭切换
	DeviceIds  []string   `json:"deviceIds"`  //删除设备Id列表
	HomeId     string     `json:"homeId"`     //家庭Id
}
