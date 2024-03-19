package entitys

import (
	"cloud_platform/iot_common/iotutil"
	proto "cloud_platform/iot_proto/protos/protosService"
)

// 查询条件
type AppQueryEntitys struct {
	Page      uint64      `json:"page,omitempty"`
	Limit     uint64      `json:"limit,omitempty"`
	Sort      string      `json:"sort,omitempty"`
	SortField string      `json:"sortField,omitempty"`
	SearchKey string      `json:"searchKey"`
	Query     AppQueryObj `json:"query,omitempty"`
}

type AppQueryObj struct {
	AppName  string `json:"appName,omitempty"`  //app名称
	UserName string `json:"userName,omitempty"` //用户账号
}

type AppEntitys struct {
	AppID             string `json:"appId"`
	AppName           string `json:"appName"`
	DeveloperID       string `json:"developerId"`
	RegisterUserTotal int64  `json:"registerUserTotal"`
	AcitveUserTotal   int64  `json:"acitveUserTotal"`
	Version           string `json:"version"`
	VerTotal          int64  `json:"verTotal"`
	FeedbackQuantity  int64  `json:"feedbackQuantity"`
}

type AppDetailEntitys struct {
	Account           string        `json:"account"`
	AppName           string        `json:"appName"`
	AppType           string        `json:"appType"`
	RegisterUserTotal int64         `json:"registerUserTotal"`
	AcitveUserTotal   int64         `json:"acitveUserTotal"`
	AppVersionList    []VersionList `json:"versionList,omitempty"`
}

type VersionList struct {
	AppVersion  string `json:"appVersion"`
	DevStatus   int64  `json:"devStatus"`
	BuildNumber int64  `json:"buildNumber"`
	LastOptTime int64  `json:"lastOptTime"`
	LastOptUser string `json:"lastOptUser"`
}

func PmAppData_pb2e(src *proto.PmAppData) *AppEntitys {
	if src == nil {
		return nil
	}
	obj := AppEntitys{
		AppID:             iotutil.ToString(src.AppId),
		AppName:           src.AppName,
		DeveloperID:       src.DevAccount,
		RegisterUserTotal: src.RegisterUserSum,
		AcitveUserTotal:   src.ActiveUserSum,
		Version:           src.LastVersion,
		VerTotal:          src.VersionSum,
		FeedbackQuantity:  src.FeedbackSum,
	}
	return &obj
}
