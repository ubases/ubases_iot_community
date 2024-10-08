// Code generated by sgen.exe,2022-05-17 13:13:12. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package convert

import (
	"cloud_platform/iot_model/db_config/model"
	proto "cloud_platform/iot_proto/protos/protosService"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func LangTranslate_pb2db(src *proto.LangTranslate) *model.TLangTranslate {
	if src == nil {
		return nil
	}
	dbObj := model.TLangTranslate{
		Id:           src.Id,
		SourceTable:  src.SourceTable,
		SourceRowId:  src.SourceRowId,
		Lang:         src.Lang,
		FieldName:    src.FieldName,
		FieldType:    src.FieldType,
		FieldValue:   src.FieldValue,
		CreatedBy:    src.CreatedBy,
		UpdatedBy:    src.UpdatedBy,
		PlatformType: src.PlatformType,
	}
	return &dbObj
}

func LangTranslate_db2pb(src *model.TLangTranslate) *proto.LangTranslate {
	if src == nil {
		return nil
	}
	pbObj := proto.LangTranslate{
		Id:           src.Id,
		SourceTable:  src.SourceTable,
		SourceRowId:  src.SourceRowId,
		Lang:         src.Lang,
		FieldName:    src.FieldName,
		FieldType:    src.FieldType,
		FieldValue:   src.FieldValue,
		CreatedBy:    src.CreatedBy,
		UpdatedBy:    src.UpdatedBy,
		CreatedAt:    timestamppb.New(src.CreatedAt),
		UpdatedAt:    timestamppb.New(src.UpdatedAt),
		PlatformType: src.PlatformType,
	}
	return &pbObj
}
