package entitys

//设备配网Token缓存数据接口
type DeviceNetworkTokenCacheModel struct {
	UserId  string   `json:"user_id"`
	Devices []string `json:"devices"`
}
