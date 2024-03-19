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
type IotOtaVersionVerifyEntitys struct {
	Id            int64     `json:"id,omitempty"`
	VersionId     int64     `json:"versionId,omitempty"`
	DeviceVersion string    `json:"deviceVersion,omitempty"`
	Did           string    `json:"did,omitempty"`
	DeviceId      int64     `json:"deviceId,omitempty"`
	Status        int32     `json:"status,omitempty"`
	DeviceLog     string    `json:"deviceLog,omitempty"`
	CreatedBy     int64     `json:"createdBy,omitempty"`
	UpdatedBy     int64     `json:"updatedBy,omitempty"`
	CreatedAt     time.Time `json:"createdAt,omitempty"`
	UpdatedAt     time.Time `json:"updatedAt,omitempty"`
	DeletedAt     time.Time `json:"deletedAt,omitempty"`
}

// 查询条件
type IotOtaVersionVerifyQuery struct {
	Page      uint64                    `json:"page,omitempty"`
	Limit     uint64                    `json:"limit,omitempty"`
	Sort      string                    `json:"sort,omitempty"`
	SortField string                    `json:"sortField,omitempty"`
	Query     IotOtaVersionVerifyFilter `json:"query,omitempty"`
}

// IotOtaVersionVerifyFilter，查询条件，字段请根据需要自行增减
type IotOtaVersionVerifyFilter struct {
	Id            int64     `json:"id,omitempty"`
	VersionId     int64     `json:"versionId,omitempty"`
	DeviceVersion string    `json:"deviceVersion,omitempty"`
	Did           string    `json:"did,omitempty"`
	DeviceId      int64     `json:"deviceId,omitempty"`
	Status        int32     `json:"status,omitempty"`
	DeviceLog     string    `json:"deviceLog,omitempty"`
	CreatedBy     int64     `json:"createdBy,omitempty"`
	UpdatedBy     int64     `json:"updatedBy,omitempty"`
	CreatedAt     time.Time `json:"createdAt,omitempty"`
	UpdatedAt     time.Time `json:"updatedAt,omitempty"`
	DeletedAt     time.Time `json:"deletedAt,omitempty"`
}

// 实体转pb对象
func IotOtaVersionVerify_e2pb(src *IotOtaVersionVerifyEntitys) *proto.IotOtaVersionVerify {
	if src == nil {
		return nil
	}
	pbObj := proto.IotOtaVersionVerify{
		Id:            src.Id,
		VersionId:     src.VersionId,
		DeviceVersion: src.DeviceVersion,
		Did:           src.Did,
		DeviceId:      src.DeviceId,
		Status:        src.Status,
		DeviceLog:     src.DeviceLog,
		CreatedBy:     src.CreatedBy,
		UpdatedBy:     src.UpdatedBy,
		CreatedAt:     timestamppb.New(src.CreatedAt),
		UpdatedAt:     timestamppb.New(src.UpdatedAt),
	}
	return &pbObj
}

// pb对象转实体
func IotOtaVersionVerify_pb2e(src *proto.IotOtaVersionVerify) *IotOtaVersionVerifyEntitys {
	if src == nil {
		return nil
	}
	entitysObj := IotOtaVersionVerifyEntitys{
		Id:            src.Id,
		VersionId:     src.VersionId,
		DeviceVersion: src.DeviceVersion,
		Did:           src.Did,
		DeviceId:      src.DeviceId,
		Status:        src.Status,
		DeviceLog:     src.DeviceLog,
		CreatedBy:     src.CreatedBy,
		UpdatedBy:     src.UpdatedBy,
		CreatedAt:     src.CreatedAt.AsTime(),
		UpdatedAt:     src.UpdatedAt.AsTime(),
	}
	return &entitysObj
}
