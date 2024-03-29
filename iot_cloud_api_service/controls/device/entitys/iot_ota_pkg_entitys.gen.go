// Code generated by sgen.exe,2022-04-21 14:24:41. DO NOT EDIT.
// versions: v1.0.0
//  API结构体封装，请自行根据需要删减字段修改

package entitys

import (
	proto "cloud_platform/iot_proto/protos/protosService"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

// 增、删、改及查询返回
type IotOtaPkgEntitys struct {
	Id                 int64     `json:"id,omitempty"`
	Type               int32     `json:"type,omitempty"`
	Name               string    `json:"name,omitempty"`
	Version            string    `json:"version,omitempty"`
	UpgradeMode        int32     `json:"upgradeMode,omitempty"`
	Url                string    `json:"url,omitempty"`
	KeyVersionFlag     int32     `json:"keyVersionFlag,omitempty"`
	SystemType         int32     `json:"systemType,omitempty"`
	MinimumEcuRequired string    `json:"minimumEcuRequired,omitempty"`
	MinimumMcuRequired string    `json:"minimumMcuRequired,omitempty"`
	Status             int32     `json:"status,omitempty"`
	UploadTime         time.Time `json:"uploadTime,omitempty"`
	CreatedBy          int64     `json:"createdBy,omitempty"`
	UpdatedBy          int64     `json:"updatedBy,omitempty"`
	CreatedAt          time.Time `json:"createdAt,omitempty"`
	UpdatedAt          time.Time `json:"updatedAt,omitempty"`
	DeletedAt          time.Time `json:"deletedAt,omitempty"`
}

// 查询条件
type IotOtaPkgQuery struct {
	Page      uint64          `json:"page,omitempty"`
	Limit     uint64          `json:"limit,omitempty"`
	Sort      string          `json:"sort,omitempty"`
	SortField string          `json:"sortField,omitempty"`
	Query     IotOtaPkgFilter `json:"query,omitempty"`
}

// IotOtaPkgFilter，查询条件，字段请根据需要自行增减
type IotOtaPkgFilter struct {
	Id                 int64     `json:"id,omitempty"`
	Type               int32     `json:"type,omitempty"`
	Name               string    `json:"name,omitempty"`
	Version            string    `json:"version,omitempty"`
	UpgradeMode        int32     `json:"upgradeMode,omitempty"`
	Url                string    `json:"url,omitempty"`
	KeyVersionFlag     int32     `json:"keyVersionFlag,omitempty"`
	SystemType         int32     `json:"systemType,omitempty"`
	MinimumEcuRequired string    `json:"minimumEcuRequired,omitempty"`
	MinimumMcuRequired string    `json:"minimumMcuRequired,omitempty"`
	Status             int32     `json:"status,omitempty"`
	UploadTime         time.Time `json:"uploadTime,omitempty"`
	CreatedBy          int64     `json:"createdBy,omitempty"`
	UpdatedBy          int64     `json:"updatedBy,omitempty"`
	CreatedAt          time.Time `json:"createdAt,omitempty"`
	UpdatedAt          time.Time `json:"updatedAt,omitempty"`
	DeletedAt          time.Time `json:"deletedAt,omitempty"`
}

// 实体转pb对象
func IotOtaPkg_e2pb(src *IotOtaPkgEntitys) *proto.IotOtaPkg {
	if src == nil {
		return nil
	}
	pbObj := proto.IotOtaPkg{
		Id:                 src.Id,
		Type:               src.Type,
		Name:               src.Name,
		Version:            src.Version,
		UpgradeMode:        src.UpgradeMode,
		Url:                src.Url,
		KeyVersionFlag:     src.KeyVersionFlag,
		SystemType:         src.SystemType,
		MinimumEcuRequired: src.MinimumEcuRequired,
		MinimumMcuRequired: src.MinimumMcuRequired,
		Status:             src.Status,
		UploadTime:         timestamppb.New(src.UploadTime),
		CreatedBy:          src.CreatedBy,
		UpdatedBy:          src.UpdatedBy,
		CreatedAt:          timestamppb.New(src.CreatedAt),
		UpdatedAt:          timestamppb.New(src.UpdatedAt),
	}
	return &pbObj
}

// pb对象转实体
func IotOtaPkg_pb2e(src *proto.IotOtaPkg) *IotOtaPkgEntitys {
	if src == nil {
		return nil
	}
	entitysObj := IotOtaPkgEntitys{
		Id:                 src.Id,
		Type:               src.Type,
		Name:               src.Name,
		Version:            src.Version,
		UpgradeMode:        src.UpgradeMode,
		Url:                src.Url,
		KeyVersionFlag:     src.KeyVersionFlag,
		SystemType:         src.SystemType,
		MinimumEcuRequired: src.MinimumEcuRequired,
		MinimumMcuRequired: src.MinimumMcuRequired,
		Status:             src.Status,
		UploadTime:         src.UploadTime.AsTime(),
		CreatedBy:          src.CreatedBy,
		UpdatedBy:          src.UpdatedBy,
		CreatedAt:          src.CreatedAt.AsTime(),
		UpdatedAt:          src.UpdatedAt.AsTime(),
	}
	return &entitysObj
}
