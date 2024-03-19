package iotstruct

import (
	"cloud_platform/iot_common/iotutil"
)

type TranslatePush struct {
	Id            string              `json:"id"`
	SourceTable   string              `json:"sourceTable,omitempty"`
	PlatformType  int32               `json:"platformType"`
	SourceRowId   string              `json:"sourceRowId,omitempty"`
	Lang          string              `json:"lang"`
	FieldName     string              `json:"fieldName"`
	FieldType     int32               `json:"fieldType"`
	FieldValue    string              `json:"fieldValue"`
	TranslateList []TranslatePushItem `json:"translateList,omitempty"`
}

type TranslatePushItem struct {
	Id          string `json:"id,omitempty"`
	Lang        string `json:"lang,omitempty"`
	SourceRowId string `json:"sourceRowId,omitempty"`
	FieldName   string `json:"fieldName,omitempty"`
	FieldType   int32  `json:"fieldType,omitempty"`
	FieldValue  string `json:"fieldValue,omitempty"`
}

// SetContent TODO 设置翻译
func (s TranslatePush) SetContent(sourceTable string, sourceRowId interface{}, fieldName string, name string, nameEn string) string {
	defer iotutil.PanicHandler()
	if name == "" {
		return ""
	}
	s.SourceTable = sourceTable
	s.SourceRowId = iotutil.ToString(sourceRowId)
	s.TranslateList = []TranslatePushItem{
		{
			Lang:       "zh",
			FieldName:  fieldName,
			FieldValue: name,
		},
	}
	if nameEn != "" {
		s.TranslateList = append(s.TranslateList, TranslatePushItem{
			Lang:       "en",
			FieldName:  fieldName,
			FieldValue: nameEn,
		})
	}
	return iotutil.ToString(s)
}

func (s *TranslatePush) AppendContent(sourceRowId string, fieldName string, name string, nameEn string) {
	if name == "" {
		return
	}
	s.TranslateList = append(s.TranslateList, TranslatePushItem{
		SourceRowId: sourceRowId,
		Lang:        "zh",
		FieldName:   fieldName,
		FieldValue:  name,
	})
	if nameEn != "" {
		s.TranslateList = append(s.TranslateList, TranslatePushItem{
			SourceRowId: sourceRowId,
			Lang:        "en",
			FieldName:   fieldName,
			FieldValue:  nameEn,
		})
	}
}

func (s *TranslatePush) AppendContentAndSourceTable(sourceTable string, sourceRowId interface{}, fieldName string, name string, nameEn string) {
	if name == "" {
		return
	}
	s.SourceTable = sourceTable
	s.SourceRowId = iotutil.ToString(sourceRowId)
	s.TranslateList = append(s.TranslateList, TranslatePushItem{
		Lang:       "zh",
		FieldName:  fieldName,
		FieldValue: name,
	})
	if nameEn != "" {
		s.TranslateList = append(s.TranslateList, TranslatePushItem{
			Lang:       "en",
			FieldName:  fieldName,
			FieldValue: nameEn,
		})
	}
}
