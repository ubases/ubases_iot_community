// Code generated by sgen.exe,2022-04-17 14:07:19. DO NOT EDIT.
// versions: v1.0.0
//  API结构体封装，请自行根据需要删减字段修改

package entitys

import (
	"cloud_platform/iot_proto/protos/protosService"
)

// 增、删、改及查询返回
type ToolsGenTableColumnEntitys struct {
	ColumnId         int64  `json:"columnId,omitempty"`
	TableId          int64  `json:"tableId,omitempty"`
	ColumnName       string `json:"columnName,omitempty"`
	ColumnComment    string `json:"columnComment,omitempty"`
	ColumnType       string `json:"columnType,omitempty"`
	GoType           string `json:"goType,omitempty"`
	GoField          string `json:"goField,omitempty"`
	HtmlField        string `json:"htmlField,omitempty"`
	IsPk             string `json:"isPk,omitempty"`
	IsIncrement      string `json:"isIncrement,omitempty"`
	IsRequired       string `json:"isRequired,omitempty"`
	IsInsert         string `json:"isInsert,omitempty"`
	IsEdit           string `json:"isEdit,omitempty"`
	IsList           string `json:"isList,omitempty"`
	IsQuery          string `json:"isQuery,omitempty"`
	QueryType        string `json:"queryType,omitempty"`
	HtmlType         string `json:"htmlType,omitempty"`
	DictType         string `json:"dictType,omitempty"`
	Sort             int32  `json:"sort,omitempty"`
	LinkTableName    string `json:"linkTableName,omitempty"`
	LinkTableClass   string `json:"linkTableClass,omitempty"`
	LinkTablePackage string `json:"linkTablePackage,omitempty"`
	LinkLabelId      string `json:"linkLabelId,omitempty"`
	LinkLabelName    string `json:"linkLabelName,omitempty"`
}

// 查询条件
type ToolsGenTableColumnQuery struct {
	Page      uint64                    `json:"page,omitempty"`
	Limit     uint64                    `json:"limit,omitempty"`
	Sort      string                    `json:"sort,omitempty"`
	SortField string                    `json:"sortField,omitempty"`
	Query     ToolsGenTableColumnFilter `json:"query,omitempty"`
}
type ToolsGenTableColumnFilter struct {
	ColumnId         int64  `json:"columnId,omitempty"`
	TableId          int64  `json:"tableId,omitempty"`
	ColumnName       string `json:"columnName,omitempty"`
	ColumnComment    string `json:"columnComment,omitempty"`
	ColumnType       string `json:"columnType,omitempty"`
	GoType           string `json:"goType,omitempty"`
	GoField          string `json:"goField,omitempty"`
	HtmlField        string `json:"htmlField,omitempty"`
	IsPk             string `json:"isPk,omitempty"`
	IsIncrement      string `json:"isIncrement,omitempty"`
	IsRequired       string `json:"isRequired,omitempty"`
	IsInsert         string `json:"isInsert,omitempty"`
	IsEdit           string `json:"isEdit,omitempty"`
	IsList           string `json:"isList,omitempty"`
	IsQuery          string `json:"isQuery,omitempty"`
	QueryType        string `json:"queryType,omitempty"`
	HtmlType         string `json:"htmlType,omitempty"`
	DictType         string `json:"dictType,omitempty"`
	Sort             int32  `json:"sort,omitempty"`
	LinkTableName    string `json:"linkTableName,omitempty"`
	LinkTableClass   string `json:"linkTableClass,omitempty"`
	LinkTablePackage string `json:"linkTablePackage,omitempty"`
	LinkLabelId      string `json:"linkLabelId,omitempty"`
	LinkLabelName    string `json:"linkLabelName,omitempty"`
}

// 实体转pb对象
func ToolsGenTableColumn_e2pb(src *ToolsGenTableColumnEntitys) *protosService.ToolsGenTableColumn {
	if src == nil {
		return nil
	}
	pbObj := protosService.ToolsGenTableColumn{
		ColumnId:         src.ColumnId,
		TableId:          src.TableId,
		ColumnName:       src.ColumnName,
		ColumnComment:    src.ColumnComment,
		ColumnType:       src.ColumnType,
		GoType:           src.GoType,
		GoField:          src.GoField,
		HtmlField:        src.HtmlField,
		IsPk:             src.IsPk,
		IsIncrement:      src.IsIncrement,
		IsRequired:       src.IsRequired,
		IsInsert:         src.IsInsert,
		IsEdit:           src.IsEdit,
		IsList:           src.IsList,
		IsQuery:          src.IsQuery,
		QueryType:        src.QueryType,
		HtmlType:         src.HtmlType,
		DictType:         src.DictType,
		Sort:             src.Sort,
		LinkTableName:    src.LinkTableName,
		LinkTableClass:   src.LinkTableClass,
		LinkTablePackage: src.LinkTablePackage,
		LinkLabelId:      src.LinkLabelId,
		LinkLabelName:    src.LinkLabelName,
	}
	return &pbObj
}

// pb对象转实体
func ToolsGenTableColumn_pb2e(src *protosService.ToolsGenTableColumn) *ToolsGenTableColumnEntitys {
	if src == nil {
		return nil
	}
	entitysObj := ToolsGenTableColumnEntitys{
		ColumnId:         src.ColumnId,
		TableId:          src.TableId,
		ColumnName:       src.ColumnName,
		ColumnComment:    src.ColumnComment,
		ColumnType:       src.ColumnType,
		GoType:           src.GoType,
		GoField:          src.GoField,
		HtmlField:        src.HtmlField,
		IsPk:             src.IsPk,
		IsIncrement:      src.IsIncrement,
		IsRequired:       src.IsRequired,
		IsInsert:         src.IsInsert,
		IsEdit:           src.IsEdit,
		IsList:           src.IsList,
		IsQuery:          src.IsQuery,
		QueryType:        src.QueryType,
		HtmlType:         src.HtmlType,
		DictType:         src.DictType,
		Sort:             src.Sort,
		LinkTableName:    src.LinkTableName,
		LinkTableClass:   src.LinkTableClass,
		LinkTablePackage: src.LinkTablePackage,
		LinkLabelId:      src.LinkLabelId,
		LinkLabelName:    src.LinkLabelName,
	}
	return &entitysObj
}
