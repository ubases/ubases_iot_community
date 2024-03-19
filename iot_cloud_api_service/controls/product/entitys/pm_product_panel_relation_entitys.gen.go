// Code generated by sgen.exe,2022-04-21 12:44:21. DO NOT EDIT.
// versions: v1.0.0
//  API结构体封装，请自行根据需要删减字段修改

package entitys

import (
	proto "cloud_platform/iot_proto/protos/protosService"
)

// 增、删、改及查询返回
type PmProductPanelRelationEntitys struct {
	Id             int64 `json:"id,omitempty"`
	ProductId      int64 `json:"productId,omitempty"`
	ControlPanelId int64 `json:"controlPanelId,omitempty"`
}

// 查询条件
type PmProductPanelRelationQuery struct {
	Page      uint64                       `json:"page,omitempty"`
	Limit     uint64                       `json:"limit,omitempty"`
	Sort      string                       `json:"sort,omitempty"`
	SortField string                       `json:"sortField,omitempty"`
	Query     PmProductPanelRelationFilter `json:"query,omitempty"`
}

// PmProductPanelRelationFilter，查询条件，字段请根据需要自行增减
type PmProductPanelRelationFilter struct {
	Id             int64 `json:"id,omitempty"`
	ProductId      int64 `json:"productId,omitempty"`
	ControlPanelId int64 `json:"controlPanelId,omitempty"`
}

// 实体转pb对象
func PmProductPanelRelation_e2pb(src *PmProductPanelRelationEntitys) *proto.PmProductPanelRelation {
	if src == nil {
		return nil
	}
	pbObj := proto.PmProductPanelRelation{
		Id:             src.Id,
		ProductId:      src.ProductId,
		ControlPanelId: src.ControlPanelId,
	}
	return &pbObj
}

// pb对象转实体
func PmProductPanelRelation_pb2e(src *proto.PmProductPanelRelation) *PmProductPanelRelationEntitys {
	if src == nil {
		return nil
	}
	entitysObj := PmProductPanelRelationEntitys{
		Id:             src.Id,
		ProductId:      src.ProductId,
		ControlPanelId: src.ControlPanelId,
	}
	return &entitysObj
}
