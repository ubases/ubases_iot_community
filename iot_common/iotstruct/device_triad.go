package iotstruct

type MqttToNatsDeviceTriadData struct {
	ProductKey string `json:"productKey"` //产品Key
	DeviceId   string `json:"deviceId"`   //设备Key
	UserName   string `json:"userName"`   // 用户名
	Passward   string `json:"passward"`   // 设备密码
	Salt       string `json:"salt"`       // 盐值
}
