package entitys

//云管平台数据概览响应实体
type OverviewEntitys struct {
	ActiveDevice Data `json:"activeDevice"`
	Developer    Data `json:"developer"`
	AppUser      Data `json:"appUser"`
	DeviceFault  Data `json:"deviceFault"`
}

type TimeData struct {
	Time  string `json:"time"`
	Total int64  `json:"total"`
}
type Data struct {
	Total int64      `json:"total"`
	Data  []TimeData `json:"data,omitempty"`
}

//设备区域分布
type DeviceCityEntitys struct {
	Country string     `json:"country"`
	Data    []CityData `json:"data,omitempty"`
}
type CityData struct {
	CityName     string  `json:"cityName"`
	CityCode     string  `json:"cityCode"`
	Longitude    float64 `json:"longitude"`
	Latitude     float64 `json:"latitude"`
	ActiveDevice int64   `json:"activeDevice"`
	OnlineDevice int64   `json:"onlineDevice"`
}
