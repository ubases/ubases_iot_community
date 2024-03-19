// Code generated by sgen.exe,2022-04-29 15:04:31. DO NOT EDIT.
// versions: v1.0.0
//  API结构体封装，请自行根据需要删减字段修改

package entitys

import (
	proto "cloud_platform/iot_proto/protos/protosService"
)

// 增、删、改及查询返回
type OpmProductFirmwareRelationEntitys struct {
	Id         int64 `json:"id,omitempty"`
	ProductId  int64 `json:"productId,string,omitempty"`
	FirmwareId int64 `json:"firmwareId,string,omitempty"`
	IsCustom   int32 `json:"isCustom,omitempty"`
}

// 查询条件
type OpmProductFirmwareRelationQuery struct {
	Page      uint64                            `json:"page,omitempty"`
	Limit     uint64                            `json:"limit,omitempty"`
	Sort      string                            `json:"sort,omitempty"`
	SortField string                            `json:"sortField,omitempty"`
	SearchKey string                            `json:"searchKey,omitempty"`
	Query     *OpmProductFirmwareRelationFilter `json:"query,omitempty"`
}

// OpmProductFirmwareRelationFilter，查询条件，字段请根据需要自行增减
type OpmProductFirmwareRelationFilter struct {
	Id         int64 `json:"id,omitempty"`
	ProductId  int64 `json:"productId,omitempty"`
	FirmwareId int64 `json:"firmwareId,omitempty"`
}

// 实体转pb对象
func OpmProductFirmwareRelation_e2pb(src *OpmProductFirmwareRelationEntitys) *proto.OpmProductFirmwareRelation {
	if src == nil {
		return nil
	}
	pbObj := proto.OpmProductFirmwareRelation{
		Id:         src.Id,
		ProductId:  src.ProductId,
		FirmwareId: src.FirmwareId,
	}
	return &pbObj
}

// pb对象转实体
func OpmProductFirmwareRelation_pb2e(src *proto.OpmProductFirmwareRelation) *OpmProductFirmwareRelationEntitys {
	if src == nil {
		return nil
	}
	entitysObj := OpmProductFirmwareRelationEntitys{
		Id:         src.Id,
		ProductId:  src.ProductId,
		FirmwareId: src.FirmwareId,
	}
	return &entitysObj
}

type ProductFirmwareItemRes struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	IsCurrent int32  `json:"isCurrent,omitempty"` //是否当前
	IsCustom  int32  `json:"isCustom"`            //是否自定义
	ModuleId  string `json:"moduleId,omitempty"`  //模组编号
	IsMust    int32  `json:"isMust"`              //是否必须
}
