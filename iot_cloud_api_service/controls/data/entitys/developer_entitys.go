package entitys

import "cloud_platform/iot_proto/protos/protosService"

// 查询条件
type DeveloperQueryEntitys struct {
	Page      uint64            `json:"page,omitempty"`
	Limit     uint64            `json:"limit,omitempty"`
	Sort      string            `json:"sort,omitempty"`
	SortField string            `json:"sortField,omitempty"`
	SearchKey string            `json:"searchKey"`
	Query     DeveloperQueryObj `json:"query,omitempty"`
}

// IotDeviceInfoFilter，查询条件，字段请根据需要自行增减
type DeveloperQueryObj struct {
	LastDay   int32  `json:"lastDay,omitempty"`   //0,今天，7近7天，30近30天，-1所有
	StartTime int64  `json:"startTime,omitempty"` //预留
	EndTime   int64  `json:"endTime,omitempty"`   //预留
	UserName  string `json:"userName,omitempty"`
}

type DeveloperEntitys struct {
	UserID            string `json:"userId"`
	UserName          string `json:"userName"`
	RegisterTime      int64  `json:"registerTime"`
	AppTotal          int64  `json:"appTotal"`
	ActiveDeviceTotal int64  `json:"activeDeviceTotal"`
	LoginAddr         string `json:"loginAddr"`
	Quantity          int64  `json:"quantity"`
	Online            int32  `json:"online"`
}

type DeveloperDetailEntitys struct {
	Account           string    `json:"account"`
	CompanyName       string    `json:"companyName"`
	RoleName          string    `json:"roleName"`
	ActiveDeviceTotal int64     `json:"activeDeviceTotal"`
	AppTotal          int64     `json:"appTotal"`
	AppList           []AppList `json:"appList"`
}
type AppList struct {
	AppID     string `json:"appId"`
	AppName   string `json:"appName"`
	DevStatus string `json:"devStatus"`
	Version   string `json:"version"`
	VerTotal  int64  `json:"verTotal"`
}

type DeveloperTotalEntitys struct {
	DevUserTotal       int64 `json:"devUserTotal"`
	DevUserOnlineTotal int64 `json:"devUserOnlineTotal"`
}

func DeveloperStat_pb2e(src *protosService.DeveloperStat) *DeveloperEntitys {
	if src == nil {
		return nil
	}
	obj := DeveloperEntitys{
		UserID:            src.UserId,
		UserName:          src.UserName,
		RegisterTime:      src.RegisterTime.GetSeconds(),
		AppTotal:          src.AppTotal,
		ActiveDeviceTotal: src.ActiveDeviceTotal,
		LoginAddr:         src.LoginAddr,
		Quantity:          src.Quantity,
		Online:            src.Online,
	}
	return &obj
}
