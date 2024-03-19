package entitys

// 1.0.3激活数据
type OpenActiveEntitys struct {
	DeviceTodayActive int32 `json:"deviceTodayActive"`
	Device7DayActive  int32 `json:"device7DayActive"`
	DeviceActiveAll   int32 `json:"deviceActiveAll"`
	DeviceMonActive   Data  `json:"deviceMonActive"`
	DeviceDayActive   Data  `json:"deviceDayActive"`
}

// 1.0.3故障数据
type OpenFaultEntitys struct {
	DeviceMonFault  Data      `json:"deviceMonFault"`
	DeviceFaultType FaultData `json:"deviceFaultType"`
}

type FaultData struct {
	Total int64           `json:"total"`
	Data  []FaultTypeData `json:"data"`
}

type FaultTypeData struct {
	FaultType string `json:"faultType"`
	Total     int64  `json:"total"`
}
