// Code generated by sgen.exe,2022-10-24 08:40:56. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package convert

import (
	"cloud_platform/iot_model/db_app_oem/model"
	proto "cloud_platform/iot_proto/protos/protosService"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func OemAppTemplateUi_pb2db(src *proto.OemAppTemplateUi) *model.TOemAppTemplateUi {
	if src == nil {
		return nil
	}
	dbObj := model.TOemAppTemplateUi{
		Id:            src.Id,
		Type:          src.Type,
		Name:          src.Name,
		NameEn:        src.NameEn,
		Code:          src.Code,
		Sort:          src.Sort,
		PageJson:      src.PageJson,
		AppTemplateId: src.AppTemplateId,
		CreatedBy:     src.CreatedBy,
		CreatedAt:     src.CreatedAt.AsTime(),
		UpdatedBy:     src.UpdatedBy,
		UpdatedAt:     src.UpdatedAt.AsTime(),
	}
	return &dbObj
}

func OemAppTemplateUi_db2pb(src *model.TOemAppTemplateUi) *proto.OemAppTemplateUi {
	if src == nil {
		return nil
	}
	pbObj := proto.OemAppTemplateUi{
		Id:            src.Id,
		Type:          src.Type,
		Name:          src.Name,
		NameEn:        src.NameEn,
		Code:          src.Code,
		Sort:          src.Sort,
		PageJson:      src.PageJson,
		AppTemplateId: src.AppTemplateId,
		CreatedBy:     src.CreatedBy,
		CreatedAt:     timestamppb.New(src.CreatedAt),
		UpdatedBy:     src.UpdatedBy,
		UpdatedAt:     timestamppb.New(src.UpdatedAt),
	}
	return &pbObj
}
