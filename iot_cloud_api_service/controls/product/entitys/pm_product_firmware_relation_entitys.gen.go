// Code generated by sgen.exe,2022-04-21 12:44:21. DO NOT EDIT.
// versions: v1.0.0
//  API结构体封装，请自行根据需要删减字段修改

package entitys

import (
	proto "cloud_platform/iot_proto/protos/protosService"
)

// 增、删、改及查询返回
type PmProductFirmwareRelationEntitys struct {
	Id         int64 `json:"id,omitempty"`
	ProductId  int64 `json:"productId,omitempty"`
	FirmwareId int64 `json:"firmwareId,omitempty"`
}

// 查询条件
type PmProductFirmwareRelationQuery struct {
	Page      uint64                          `json:"page,omitempty"`
	Limit     uint64                          `json:"limit,omitempty"`
	Sort      string                          `json:"sort,omitempty"`
	SortField string                          `json:"sortField,omitempty"`
	Query     PmProductFirmwareRelationFilter `json:"query,omitempty"`
}

// PmProductFirmwareRelationFilter，查询条件，字段请根据需要自行增减
type PmProductFirmwareRelationFilter struct {
	Id         int64 `json:"id,omitempty"`
	ProductId  int64 `json:"productId,omitempty"`
	FirmwareId int64 `json:"firmwareId,omitempty"`
}

// 实体转pb对象
func PmProductFirmwareRelation_e2pb(src *PmProductFirmwareRelationEntitys) *proto.PmProductFirmwareRelation {
	if src == nil {
		return nil
	}
	pbObj := proto.PmProductFirmwareRelation{
		Id:         src.Id,
		ProductId:  src.ProductId,
		FirmwareId: src.FirmwareId,
	}
	return &pbObj
}

// pb对象转实体
func PmProductFirmwareRelation_pb2e(src *proto.PmProductFirmwareRelation) *PmProductFirmwareRelationEntitys {
	if src == nil {
		return nil
	}
	entitysObj := PmProductFirmwareRelationEntitys{
		Id:         src.Id,
		ProductId:  src.ProductId,
		FirmwareId: src.FirmwareId,
	}
	return &entitysObj
}
