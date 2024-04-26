// Code generated by sgen,{{.CurrentTime}}. DO NOT EDIT.
// versions: v1.0.0
//  API结构体封装，请自行根据需要删减字段修改

package entitys

import (
	proto "{{.ProtoPackage}}"
    {{- $hasDateTimeA := true }}{{range $i, $v := .Fields}} {{if $hasDateTimeA}} {{if and (ne $v.ColumnName "deleted_at") ( eq $v.DataType "timestamp")}}
	"time"
	"google.golang.org/protobuf/types/known/timestamppb"
	{{- $hasDateTimeA = false }}{{end}}{{end}}{{end}}
)

//增、删、改及查询返回
type {{.ModelName}}Entitys struct { {{range $i, $v := .Fields}} {{if eq $v.DataType "bigint"}}
    {{case2CamelAndUcfirst $v.ColumnName}} int64  `json:"{{case2CamelAndLcfirst $v.ColumnName}},string,omitempty"` {{else if eq $v.DataType "tinyint"}}
    {{case2CamelAndUcfirst $v.ColumnName}} int32  `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"` {{else if eq $v.DataType "int"}}
    {{case2CamelAndUcfirst $v.ColumnName}} int32  `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"` {{else if eq $v.DataType "integer"}}
    {{case2CamelAndUcfirst $v.ColumnName}} int32  `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"` {{else if eq $v.DataType "smallint"}}
    {{case2CamelAndUcfirst $v.ColumnName}} int32  `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"` {{else if eq $v.DataType "mediumint"}}
    {{case2CamelAndUcfirst $v.ColumnName}} int32  `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"` {{else if eq $v.DataType "year"}}
    {{case2CamelAndUcfirst $v.ColumnName}} int32  `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"` {{else if eq $v.DataType "float"}}
    {{case2CamelAndUcfirst $v.ColumnName}} float32 `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"` {{else if eq $v.DataType "double"}}
    {{case2CamelAndUcfirst $v.ColumnName}} float64 `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"` {{else if eq $v.DataType "decimal"}}
    {{case2CamelAndUcfirst $v.ColumnName}} float64 `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"` {{else if eq $v.DataType "binary"}}
    {{case2CamelAndUcfirst $v.ColumnName}} string  `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"`  {{else if eq $v.DataType "varbinary"}}
    {{case2CamelAndUcfirst $v.ColumnName}} string  `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"`  {{else if eq $v.DataType "tinyblob"}}
    {{case2CamelAndUcfirst $v.ColumnName}} string  `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"`  {{else if eq $v.DataType "blob"}}
    {{case2CamelAndUcfirst $v.ColumnName}} string  `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"`  {{else if eq $v.DataType "mediumblob"}}
    {{case2CamelAndUcfirst $v.ColumnName}} string  `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"`  {{else if eq $v.DataType "longblob"}}
    {{case2CamelAndUcfirst $v.ColumnName}} string  `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"`  {{else if eq $v.DataType "time"}}
    {{case2CamelAndUcfirst $v.ColumnName}} time.Time `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"` {{else if eq $v.DataType "date"}}
    {{case2CamelAndUcfirst $v.ColumnName}} time.Time `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"` {{else if eq $v.DataType "datetime"}}
    {{case2CamelAndUcfirst $v.ColumnName}} time.Time `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"` {{else if eq $v.DataType "timestamp"}}
    {{case2CamelAndUcfirst $v.ColumnName}} time.Time `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"`  {{else if eq $v.DataType "boolean"}}
    {{case2CamelAndUcfirst $v.ColumnName}} bool `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"` {{else if eq $v.DataType "bit"}}
    {{case2CamelAndUcfirst $v.ColumnName}} string `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"` {{else}}
    {{case2CamelAndUcfirst $v.ColumnName}} string `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"` {{end}}{{end}}
}

//新增参数非空检查
func (s *{{.ModelName}}Entitys) AddCheck() error {
	return nil
}

//修改参数非空检查
func (s *{{.ModelName}}Entitys) UpdateCheck() error {
	return nil
}

//查询参数必填检查
func (*{{.ModelName}}Query) QueryCheck() error {
	return nil
}

//查询条件
type {{.ModelName}}Query struct {
	Page      uint64 `json:"page,omitempty"`
	Limit     uint64 `json:"limit,omitempty"`
	Sort      string `json:"sort,omitempty"`
	SortField string `json:"sortField,omitempty"`
	SearchKey string `json:"searchKey,omitempty"`
	Query   *{{.ModelName}}Filter `json:"query,omitempty"`
}
//{{.ModelName}}Filter，查询条件，字段请根据需要自行增减
type {{.ModelName}}Filter struct { {{range $i, $v := .Fields}} {{if eq $v.DataType "bigint"}}
    {{case2CamelAndUcfirst $v.ColumnName}} int64  `json:"{{case2CamelAndLcfirst $v.ColumnName}},string,omitempty"` {{else if eq $v.DataType "tinyint"}}
    {{case2CamelAndUcfirst $v.ColumnName}} int32  `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"` {{else if eq $v.DataType "int"}}
    {{case2CamelAndUcfirst $v.ColumnName}} int32  `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"` {{else if eq $v.DataType "integer"}}
    {{case2CamelAndUcfirst $v.ColumnName}} int32  `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"` {{else if eq $v.DataType "smallint"}}
    {{case2CamelAndUcfirst $v.ColumnName}} int32  `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"` {{else if eq $v.DataType "mediumint"}}
    {{case2CamelAndUcfirst $v.ColumnName}} int32  `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"` {{else if eq $v.DataType "year"}}
    {{case2CamelAndUcfirst $v.ColumnName}} int32  `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"` {{else if eq $v.DataType "float"}}
    {{case2CamelAndUcfirst $v.ColumnName}} float32 `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"` {{else if eq $v.DataType "double"}}
    {{case2CamelAndUcfirst $v.ColumnName}} float64 `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"` {{else if eq $v.DataType "decimal"}}
    {{case2CamelAndUcfirst $v.ColumnName}} float64 `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"` {{else if eq $v.DataType "binary"}}
    {{case2CamelAndUcfirst $v.ColumnName}} string  `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"`  {{else if eq $v.DataType "varbinary"}}
    {{case2CamelAndUcfirst $v.ColumnName}} string  `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"`  {{else if eq $v.DataType "tinyblob"}}
    {{case2CamelAndUcfirst $v.ColumnName}} string  `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"`  {{else if eq $v.DataType "blob"}}
    {{case2CamelAndUcfirst $v.ColumnName}} string  `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"`  {{else if eq $v.DataType "mediumblob"}}
    {{case2CamelAndUcfirst $v.ColumnName}} string  `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"`  {{else if eq $v.DataType "longblob"}}
    {{case2CamelAndUcfirst $v.ColumnName}} string  `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"`  {{else if eq $v.DataType "time"}}
    {{case2CamelAndUcfirst $v.ColumnName}} time.Time `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"` {{else if eq $v.DataType "date"}}
    {{case2CamelAndUcfirst $v.ColumnName}} time.Time `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"` {{else if eq $v.DataType "datetime"}}
    {{case2CamelAndUcfirst $v.ColumnName}} time.Time `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"` {{else if eq $v.DataType "timestamp"}}
    {{case2CamelAndUcfirst $v.ColumnName}} time.Time `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"`  {{else if eq $v.DataType "boolean"}}
    {{case2CamelAndUcfirst $v.ColumnName}} bool `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"` {{else if eq $v.DataType "bit"}}
    {{case2CamelAndUcfirst $v.ColumnName}} string `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"` {{else}}
    {{case2CamelAndUcfirst $v.ColumnName}} string `json:"{{case2CamelAndLcfirst $v.ColumnName}},omitempty"` {{end}}{{end}}
}

//实体转pb对象
func {{.ModelName}}_e2pb(src *{{.ModelName}}Entitys) *proto.{{.ModelName}} {
    if src == nil {
        return nil
    }
    pbObj := proto.{{.ModelName}}{ {{range $i, $v := .Fields}} {{if ne $v.ColumnName "deleted_at"}} {{if or (eq $v.DataType "time") (eq $v.DataType "datetime") (eq $v.DataType "date") (eq $v.DataType "timestamp") }}
        {{case2CamelAndUcfirst $v.ColumnName}}:timestamppb.New(src.{{case2CamelAndUcfirst $v.ColumnName}}),{{else}}
        {{case2CamelAndUcfirst $v.ColumnName}}:src.{{case2CamelAndUcfirst $v.ColumnName}},{{end}}{{end}}{{end}}
    }
    return &pbObj
}

//pb对象转实体
func {{.ModelName}}_pb2e(src *proto.{{.ModelName}}) *{{.ModelName}}Entitys {
	if src == nil {
		return nil
	}
	entitysObj := {{.ModelName}}Entitys{ {{range $i, $v := .Fields}} {{if ne $v.ColumnName "deleted_at"}} {{if or (eq $v.DataType "time") (eq $v.DataType "datetime") (eq $v.DataType "date") (eq $v.DataType "timestamp") }}
         {{case2CamelAndUcfirst $v.ColumnName}}:src.{{case2CamelAndUcfirst $v.ColumnName}}.AsTime(),{{else}}
         {{case2CamelAndUcfirst $v.ColumnName}}:src.{{case2CamelAndUcfirst $v.ColumnName}},{{end}}{{end}}{{end}}
	}
	return &entitysObj
}