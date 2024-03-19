// Code generated by sgen.exe,2022-05-04 16:39:29. DO NOT EDIT.
// versions: v1.0.0
//  API结构体封装，请自行根据需要删减字段修改

package entitys

import (
	proto "cloud_platform/iot_proto/protos/protosService"
	"time"
)

// 增、删、改及查询返回
type UcUserEntitys struct {
	Id               int64   `json:"id,omitempty"`
	Uid              string  `json:"uid,omitempty"`
	NickName         string  `json:"nickName,omitempty"`
	Phone            string  `json:"phone,omitempty"`
	Password         string  `json:"password,omitempty"`
	DeviceSecretHttp string  `json:"deviceSecretHttp,omitempty"`
	DeviceSecretMqtt string  `json:"deviceSecretMqtt,omitempty"`
	Photo            string  `json:"photo,omitempty"`
	Status           int32   `json:"status,omitempty"`
	Lang             string  `json:"lang,omitempty"`
	Lat              float64 `json:"lat,omitempty"`
	Lng              float64 `json:"lng,omitempty"`
	Country          string  `json:"country,omitempty"`
	Province         string  `json:"province,omitempty"`
	City             string  `json:"city,omitempty"`
	District         string  `json:"district,omitempty"`
	Email            string  `json:"email,omitempty"`
	DefaultHomeId    string  `json:"defaultHomeId,omitempty"`
	Gender           int32   `json:"gender,omitempty"`
	RegisterRegion   string  `json:"registerRegion,omitempty"`
	CreatedBy        int64   `json:"createdBy,omitempty"`
	UpdatedBy        int64   `json:"updatedBy,omitempty"`
	CancelTime       int64   `json:"cancelTime,omitempty"`
}

// 查询条件
type UcUserQuery struct {
	Page      uint64        `json:"page,omitempty"`
	Limit     uint64        `json:"limit,omitempty"`
	Sort      string        `json:"sort,omitempty"`
	SortField string        `json:"sortField,omitempty"`
	SearchKey string        `json:"searchKey,omitempty"`
	Query     *UcUserFilter `json:"query,omitempty"`
}

// UcUserFilter，查询条件，字段请根据需要自行增减
type UcUserFilter struct {
	Id               int64     `json:"id,omitempty"`
	Uid              string    `json:"uid,omitempty"`
	NickName         string    `json:"nickName,omitempty"`
	Phone            string    `json:"phone,omitempty"`
	Password         string    `json:"password,omitempty"`
	DeviceSecretHttp string    `json:"deviceSecretHttp,omitempty"`
	DeviceSecretMqtt string    `json:"deviceSecretMqtt,omitempty"`
	Photo            string    `json:"photo,omitempty"`
	Status           int32     `json:"status,omitempty"`
	Lang             string    `json:"lang,omitempty"`
	Lat              float64   `json:"lat,omitempty"`
	Lng              float64   `json:"lng,omitempty"`
	Country          string    `json:"country,omitempty"`
	Province         string    `json:"province,omitempty"`
	City             string    `json:"city,omitempty"`
	District         string    `json:"district,omitempty"`
	Email            string    `json:"email,omitempty"`
	DefaultHomeId    string    `json:"defaultHomeId,omitempty"`
	Gender           int32     `json:"gender,omitempty"`
	RegisterRegion   string    `json:"registerRegion,omitempty"`
	CreatedBy        int64     `json:"createdBy,omitempty"`
	UpdatedBy        int64     `json:"updatedBy,omitempty"`
	CancelTime       int64     `json:"cancelTime,omitempty"`
	CreatedAt        time.Time `json:"createdAt,omitempty"`
	UpdatedAt        time.Time `json:"updatedAt,omitempty"`
}

// 实体转pb对象
func UcUser_e2pb(src *UcUserEntitys) *proto.UcUser {
	if src == nil {
		return nil
	}
	pbObj := proto.UcUser{
		Id:               src.Id,
		Uid:              src.Uid,
		NickName:         src.NickName,
		Phone:            src.Phone,
		Password:         src.Password,
		DeviceSecretHttp: src.DeviceSecretHttp,
		DeviceSecretMqtt: src.DeviceSecretMqtt,
		Photo:            src.Photo,
		Status:           src.Status,
		Lang:             src.Lang,
		Lat:              src.Lat,
		Lng:              src.Lng,
		Country:          src.Country,
		Province:         src.Province,
		City:             src.City,
		District:         src.District,
		Email:            src.Email,
		DefaultHomeId:    src.DefaultHomeId,
		Gender:           src.Gender,
		RegisterRegion:   src.RegisterRegion,
		CreatedBy:        src.CreatedBy,
		UpdatedBy:        src.UpdatedBy,
		CancelTime:       src.CancelTime,
	}
	return &pbObj
}

// pb对象转实体
func UcUser_pb2e(src *proto.UcUser) *UcUserEntitys {
	if src == nil {
		return nil
	}
	entitysObj := UcUserEntitys{
		Id:               src.Id,
		Uid:              src.Uid,
		NickName:         src.NickName,
		Phone:            src.Phone,
		Password:         src.Password,
		DeviceSecretHttp: src.DeviceSecretHttp,
		DeviceSecretMqtt: src.DeviceSecretMqtt,
		Photo:            src.Photo,
		Status:           src.Status,
		Lang:             src.Lang,
		Lat:              src.Lat,
		Lng:              src.Lng,
		Country:          src.Country,
		Province:         src.Province,
		City:             src.City,
		District:         src.District,
		Email:            src.Email,
		DefaultHomeId:    src.DefaultHomeId,
		Gender:           src.Gender,
		RegisterRegion:   src.RegisterRegion,
		CreatedBy:        src.CreatedBy,
		UpdatedBy:        src.UpdatedBy,
		CancelTime:       src.CancelTime,
	}
	return &entitysObj
}
