// Code generated by sgen.exe,2022-04-27 17:47:12. DO NOT EDIT.
// versions: v1.0.0
//  API结构体封装，请自行根据需要删减字段修改

package entitys

import (
	"cloud_platform/iot_common/iotutil"
	proto "cloud_platform/iot_proto/protos/protosService"
	"errors"
)

// 增、删、改及查询返回
type TplDocumentTemplateFieldEntitys struct {
	Id        string `json:"id,omitempty"`
	TplType   int32  `json:"tplType,omitempty"`
	TplId     string `json:"tplId,omitempty"`
	FieldName string `json:"fieldName,omitempty"`
	FieldCode string `json:"fieldCode,omitempty"`
	FieldDesc string `json:"fieldDesc,omitempty"`
}

func (s *TplDocumentTemplateFieldEntitys) AddCheck() error {
	if s.FieldCode == "" {
		return errors.New("字段名不能为空")
	}
	if s.FieldName == "" {
		return errors.New("字段中文名不能为空")
	}
	if s.TplId == "" {
		return errors.New("所属模板编号不能为空")
	}
	return nil
}

func (s *TplDocumentTemplateFieldEntitys) UpdateCheck() error {
	if s.FieldCode == "" {
		return errors.New("字段名不能为空")
	}
	if s.FieldName == "" {
		return errors.New("字段中文名不能为空")
	}
	if s.TplId == "" {
		return errors.New("所属模板编号不能为空")
	}
	return nil
}

func (s *TplDocumentTemplateFieldQuery) QueryCheck() error {
	return nil
}

// 查询条件
type TplDocumentTemplateFieldQuery struct {
	Page      uint64                          `json:"page,omitempty"`
	Limit     uint64                          `json:"limit,omitempty"`
	Sort      string                          `json:"sort,omitempty"`
	SortField string                          `json:"sortField,omitempty"`
	SearchKey string                          `json:"searchKey,omitempty"`
	Query     *TplDocumentTemplateFieldFilter `json:"query,omitempty"`
}

// TplDocumentTemplateFieldFilter，查询条件，字段请根据需要自行增减
type TplDocumentTemplateFieldFilter struct {
	Id        int64  `json:"id,omitempty"`
	TplType   int32  `json:"tplType,omitempty"`
	TplId     int64  `json:"tplId,omitempty"`
	FieldName string `json:"fieldName,omitempty"`
	FieldCode string `json:"fieldCode,omitempty"`
	FieldDesc string `json:"fieldDesc,omitempty"`
}

// 实体转pb对象
func TplDocumentTemplateField_e2pb(src *TplDocumentTemplateFieldEntitys) *proto.TplDocumentTemplateField {
	if src == nil {
		return nil
	}
	pbObj := proto.TplDocumentTemplateField{
		TplType:   src.TplType,
		TplId:     iotutil.ToInt64(src.TplId),
		FieldName: src.FieldName,
		FieldCode: src.FieldCode,
		FieldDesc: src.FieldDesc,
	}
	if src.Id != "" {
		pbObj.Id = iotutil.ToInt64(src.Id)
	}
	return &pbObj
}

// pb对象转实体
func TplDocumentTemplateField_pb2e(src *proto.TplDocumentTemplateField) *TplDocumentTemplateFieldEntitys {
	if src == nil {
		return nil
	}
	entitysObj := TplDocumentTemplateFieldEntitys{
		Id:        iotutil.ToString(src.Id),
		TplType:   src.TplType,
		TplId:     iotutil.ToString(src.TplId),
		FieldName: src.FieldName,
		FieldCode: src.FieldCode,
		FieldDesc: src.FieldDesc,
	}
	return &entitysObj
}
