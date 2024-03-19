package services

import (
	"cloud_platform/iot_cloud_api_service/controls/lang/entitys"
	entitys3 "cloud_platform/iot_cloud_api_service/controls/lang/entitys"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_common/iotutil"
	"context"
)

// SetDefaultTranslate 设置翻译
func SetDefaultTranslate(ctx context.Context, sourceTable string, sourceRowId interface{}, fieldName string, name string, nameEn string) {
	defer iotutil.PanicHandler()
	if name == "" {
		return
	}
	langSvr := LangTranslateService{Ctx: ctx}
	langReq := entitys.LangTranslateEntitys{
		SourceTable: sourceTable,
		SourceRowId: iotutil.ToString(sourceRowId),
		TranslateList: []entitys3.BatchSaveTranslateItem{
			{
				Lang:       "zh",
				FieldName:  fieldName,
				FieldValue: name,
			},
		},
	}
	if nameEn == "" {
		langReq.TranslateList = append(langReq.TranslateList, entitys3.BatchSaveTranslateItem{
			Lang:       "en",
			FieldName:  fieldName,
			FieldValue: nameEn,
		})
	}
	langSvr.BatchInsert(langReq)
}

// SetDefaultTranslateByNats TODO 设置翻译
func SetDefaultTranslateByNats(ctx context.Context, sourceTable string, sourceRowId interface{}, fieldName string, name string, nameEn string) {
	defer iotutil.PanicHandler()
	if name == "" {
		return
	}
	//langSvr := LangTranslateService{Ctx: ctx}
	langReq := iotstruct.TranslatePush{
		SourceTable: sourceTable,
		SourceRowId: iotutil.ToString(sourceRowId),
		TranslateList: []iotstruct.TranslatePushItem{
			{
				Lang:       "zh",
				FieldName:  fieldName,
				FieldValue: name,
			},
		},
	}
	if nameEn == "" {
		langReq.TranslateList = append(langReq.TranslateList, iotstruct.TranslatePushItem{
			Lang:       "en",
			FieldName:  fieldName,
			FieldValue: nameEn,
		})
	}
	//langSvr.BatchInsert(langReq)
	//TODO 实现推送逻辑
}
