package entitys

import (
	proto "cloud_platform/iot_proto/protos/protosService"
)

type FailLogQueryEntitys struct {
	Page      uint64          `json:"page,omitempty"`
	Limit     uint64          `json:"limit,omitempty"`
	Sort      string          `json:"sort,omitempty"`
	SortField string          `json:"sortField,omitempty"`
	SearchKey string          `json:"searchKey"`
	Query     FailLogQueryObj `json:"query,omitempty"`
}

type FailLogQueryObj struct {
	Did          string `json:"did"`                 //
	LastDay      int32  `json:"lastDay,omitempty"`   //0,今天，7近7天，30近30天，-1所有
	StartTime    int64  `json:"startTime,omitempty"` //预留
	EndTime      int64  `json:"endTime,omitempty"`   //预留
	ProductKey   string `json:"productKey,omitempty"`
	Code         string `json:"faultCode,omitempty"`    //错误码
	UploadFrom   string `json:"uploadFrom,omitempty"`   //上报端
	UploadMethod string `json:"uploadMethod,omitempty"` //上报方法
}

type DeviceOperationFailLogListResponseObj struct {
	Id           string `json:"id"`
	DeviceId     string `json:"deviceId"`     //设备Id
	Type         int32  `json:"type"`         //类型 1-配网 2-OTA升级
	Content      string `json:"content"`      //上报内容
	UserId       string `json:"userId"`       //用户Id
	UserAccount  string `json:"userAccount"`  //用户账号
	TenantId     string `json:"tenantId"`     //租户编号
	AppKey       string `json:"appKey"`       //APP Key
	ProductKey   string `json:"productKey"`   //产品key
	FailTime     int64  `json:"failTime"`     //时间,时间戳
	Code         int32  `json:"code"`         //错误码
	Timezone     string `json:"timezone"`     //用户app 时区
	Region       string `json:"region"`       //用户app 登录的区域，区域id
	Lang         string `json:"lang"`         //用户app app语言
	Os           string `json:"os"`           //操作系统
	Model        string `json:"model"`        //手机型号
	AppVersion   string `json:"appVersion"`   //app版本
	Desc         string `json:"desc"`         //错误描述
	UploadFrom   string `json:"uploadFrom"`   //上报端 app\device\broker
	UploadMethod string `json:"uploadMethod"` //上报方式 http\mqtt
}

func DeviceOperationFailLogListResponseObj_pb2e(src *proto.DeviceOperationFailLogListResponseObj) *DeviceOperationFailLogListResponseObj {
	if src == nil {
		return nil
	}
	entitysObj := DeviceOperationFailLogListResponseObj{
		Id:           src.Id,
		DeviceId:     src.DeviceId,
		Type:         src.Type,
		Content:      src.Content,
		UserId:       src.UserId,
		UserAccount:  src.UserAccount,
		TenantId:     src.TenantId,
		AppKey:       src.AppKey,
		ProductKey:   src.ProductKey,
		FailTime:     src.FailTime,
		Code:         src.Code,
		Timezone:     src.Timezone,
		Region:       src.Region,
		Lang:         src.Lang,
		Os:           src.Os,
		Model:        src.Model,
		AppVersion:   src.AppVersion,
		Desc:         src.Desc,
		UploadFrom:   src.UploadFrom,
		UploadMethod: src.UploadMethod,
	}
	return &entitysObj
}
