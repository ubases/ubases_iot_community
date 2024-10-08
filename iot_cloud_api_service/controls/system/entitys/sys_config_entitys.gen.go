// Code generated by sgen.exe,2022-04-17 14:07:16. DO NOT EDIT.
// versions: v1.0.0
//  API结构体封装，请自行根据需要删减字段修改

package entitys

import (
	"cloud_platform/iot_proto/protos/protosService"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// 增、删、改及查询返回
type SysConfigEntitys struct {
	ConfigId    int64     `json:"configId,omitempty"`
	ConfigName  string    `json:"configName,omitempty"`
	ConfigKey   string    `json:"configKey,omitempty"`
	ConfigValue string    `json:"configValue,omitempty"`
	ConfigType  int32     `json:"configType,omitempty"`
	CreateBy    int32     `json:"createBy,omitempty"`
	UpdateBy    int32     `json:"updateBy,omitempty"`
	Remark      string    `json:"remark,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
	DeletedAt   time.Time `json:"deletedAt,omitempty"`
}

// 查询条件
type SysConfigQuery struct {
	Page      uint64          `json:"page,omitempty"`
	Limit     uint64          `json:"limit,omitempty"`
	Sort      string          `json:"sort,omitempty"`
	SortField string          `json:"sortField,omitempty"`
	Query     SysConfigFilter `json:"query,omitempty"`
}
type SysConfigFilter struct {
	ConfigId    int64     `json:"configId,omitempty"`
	ConfigName  string    `json:"configName,omitempty"`
	ConfigKey   string    `json:"configKey,omitempty"`
	ConfigValue string    `json:"configValue,omitempty"`
	ConfigType  int32     `json:"configType,omitempty"`
	CreateBy    int32     `json:"createBy,omitempty"`
	UpdateBy    int32     `json:"updateBy,omitempty"`
	Remark      string    `json:"remark,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
	DeletedAt   time.Time `json:"deletedAt,omitempty"`
}

// 实体转pb对象
func SysConfig_e2pb(src *SysConfigEntitys) *protosService.SysConfig {
	if src == nil {
		return nil
	}
	pbObj := protosService.SysConfig{
		ConfigId:    src.ConfigId,
		ConfigName:  src.ConfigName,
		ConfigKey:   src.ConfigKey,
		ConfigValue: src.ConfigValue,
		ConfigType:  src.ConfigType,
		CreateBy:    src.CreateBy,
		UpdateBy:    src.UpdateBy,
		Remark:      src.Remark,
		CreatedAt:   timestamppb.New(src.CreatedAt),
		UpdatedAt:   timestamppb.New(src.UpdatedAt),
	}
	return &pbObj
}

// pb对象转实体
func SysConfig_pb2e(src *protosService.SysConfig) *SysConfigEntitys {
	if src == nil {
		return nil
	}
	entitysObj := SysConfigEntitys{
		ConfigId:    src.ConfigId,
		ConfigName:  src.ConfigName,
		ConfigKey:   src.ConfigKey,
		ConfigValue: src.ConfigValue,
		ConfigType:  src.ConfigType,
		CreateBy:    src.CreateBy,
		UpdateBy:    src.UpdateBy,
		Remark:      src.Remark,
		CreatedAt:   src.CreatedAt.AsTime(),
		UpdatedAt:   src.UpdatedAt.AsTime(),
	}
	return &entitysObj
}
