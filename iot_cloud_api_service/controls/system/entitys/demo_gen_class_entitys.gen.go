// Code generated by sgen.exe,2022-04-17 14:07:15. DO NOT EDIT.
// versions: v1.0.0
//  API结构体封装，请自行根据需要删减字段修改

package entitys

import (
	"cloud_platform/iot_proto/protos/protosService"
)

// 增、删、改及查询返回
type DemoGenClassEntitys struct {
	Id        int32  `json:"id,omitempty"`
	ClassName string `json:"className,omitempty"`
}

// 查询条件
type DemoGenClassQuery struct {
	Page      uint64             `json:"page,omitempty"`
	Limit     uint64             `json:"limit,omitempty"`
	Sort      string             `json:"sort,omitempty"`
	SortField string             `json:"sortField,omitempty"`
	Query     DemoGenClassFilter `json:"query,omitempty"`
}
type DemoGenClassFilter struct {
	Id        int32  `json:"id,omitempty"`
	ClassName string `json:"className,omitempty"`
}

// 实体转pb对象
func DemoGenClass_e2pb(src *DemoGenClassEntitys) *protosService.DemoGenClass {
	if src == nil {
		return nil
	}
	pbObj := protosService.DemoGenClass{
		Id:        src.Id,
		ClassName: src.ClassName,
	}
	return &pbObj
}

// pb对象转实体
func DemoGenClass_pb2e(src *protosService.DemoGenClass) *DemoGenClassEntitys {
	if src == nil {
		return nil
	}
	entitysObj := DemoGenClassEntitys{
		Id:        src.Id,
		ClassName: src.ClassName,
	}
	return &entitysObj
}
