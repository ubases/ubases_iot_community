// Code generated by sgen.exe,2022-04-19 20:40:04. DO NOT EDIT.
// versions: v1.0.0
//  API结构体封装，请自行根据需要删减字段修改

package entitys

import (
	proto "cloud_platform/iot_proto/protos/protosService"
)

// 增、删、改及查询返回
type ConfigTranslateLanguageEntitys struct {
	Id    int64  `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
}

// 查询条件
type ConfigTranslateLanguageQuery struct {
	Page      uint64                        `json:"page,omitempty"`
	Limit     uint64                        `json:"limit,omitempty"`
	Sort      string                        `json:"sort,omitempty"`
	SortField string                        `json:"sortField,omitempty"`
	Query     ConfigTranslateLanguageFilter `json:"query,omitempty"`
}

// ConfigTranslateLanguageFilter，查询条件，字段请根据需要自行增减
type ConfigTranslateLanguageFilter struct {
	Id    int64  `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
}

// 实体转pb对象
func ConfigTranslateLanguage_e2pb(src *ConfigTranslateLanguageEntitys) *proto.ConfigTranslateLanguage {
	if src == nil {
		return nil
	}
	pbObj := proto.ConfigTranslateLanguage{
		Id:    src.Id,
		Title: src.Title,
		Value: src.Value,
	}
	return &pbObj
}

// pb对象转实体
func ConfigTranslateLanguage_pb2e(src *proto.ConfigTranslateLanguage) *ConfigTranslateLanguageEntitys {
	if src == nil {
		return nil
	}
	entitysObj := ConfigTranslateLanguageEntitys{
		Id:    src.Id,
		Title: src.Title,
		Value: src.Value,
	}
	return &entitysObj
}
