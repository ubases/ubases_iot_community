// Code generated by sgen.exe,2022-05-06 14:01:20. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package convert

import (
	"cloud_platform/iot_model/db_product/model"
	proto "cloud_platform/iot_proto/protos/protosService"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func OpmControlPanels_pb2db(src *proto.OpmControlPanels) *model.TOpmControlPanels {
	if src == nil {
		return nil
	}
	dbObj := model.TOpmControlPanels{
		Id:            src.Id,
		Name:          src.Name,
		NameEn:        src.NameEn,
		Lang:          src.Lang,
		Desc:          src.Desc,
		Url:           src.Url,
		UrlName:       src.UrlName,
		PanelSize:     src.PanelSize,
		PreviewName:   src.PreviewName,
		PreviewUrl:    src.PreviewUrl,
		PreviewSize:   src.PreviewSize,
		ProductTypeId: src.ProductTypeId,
		ProductId:     src.ProductId,
		CreatedBy:     src.CreatedBy,
		UpdatedBy:     src.UpdatedBy,
		CreatedAt:     src.CreatedAt.AsTime(),
		UpdatedAt:     src.UpdatedAt.AsTime(),
		TenantId:      src.TenantId,
	}
	return &dbObj
}

func OpmControlPanels_db2pb(src *model.TOpmControlPanels) *proto.OpmControlPanels {
	if src == nil {
		return nil
	}
	pbObj := proto.OpmControlPanels{
		Id:            src.Id,
		Name:          src.Name,
		NameEn:        src.NameEn,
		Lang:          src.Lang,
		Desc:          src.Desc,
		Url:           src.Url,
		UrlName:       src.UrlName,
		PanelSize:     src.PanelSize,
		PreviewName:   src.PreviewName,
		PreviewUrl:    src.PreviewUrl,
		PreviewSize:   src.PreviewSize,
		ProductTypeId: src.ProductTypeId,
		ProductId:     src.ProductId,
		CreatedBy:     src.CreatedBy,
		UpdatedBy:     src.UpdatedBy,
		CreatedAt:     timestamppb.New(src.CreatedAt),
		UpdatedAt:     timestamppb.New(src.UpdatedAt),
		TenantId:      src.TenantId,
	}
	return &pbObj
}
