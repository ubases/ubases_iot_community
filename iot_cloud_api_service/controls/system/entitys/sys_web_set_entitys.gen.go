// Code generated by sgen.exe,2022-04-17 14:07:19. DO NOT EDIT.
// versions: v1.0.0
//  API结构体封装，请自行根据需要删减字段修改

package entitys

import (
	"cloud_platform/iot_proto/protos/protosService"
)

// 增、删、改及查询返回
type SysWebSetEntitys struct {
	WebId      int64  `json:"webId,omitempty"`
	WebContent string `json:"webContent,omitempty"`
}

// 查询条件
type SysWebSetQuery struct {
	Page      uint64          `json:"page,omitempty"`
	Limit     uint64          `json:"limit,omitempty"`
	Sort      string          `json:"sort,omitempty"`
	SortField string          `json:"sortField,omitempty"`
	Query     SysWebSetFilter `json:"query,omitempty"`
}
type SysWebSetFilter struct {
	WebId      int64  `json:"webId,omitempty"`
	WebContent string `json:"webContent,omitempty"`
}

// 实体转pb对象
func SysWebSet_e2pb(src *SysWebSetEntitys) *protosService.SysWebSet {
	if src == nil {
		return nil
	}
	pbObj := protosService.SysWebSet{
		WebId:      src.WebId,
		WebContent: src.WebContent,
	}
	return &pbObj
}

// pb对象转实体
func SysWebSet_pb2e(src *protosService.SysWebSet) *SysWebSetEntitys {
	if src == nil {
		return nil
	}
	entitysObj := SysWebSetEntitys{
		WebId:      src.WebId,
		WebContent: src.WebContent,
	}
	return &entitysObj
}
