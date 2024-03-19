// Code generated by sgen.exe,2022-05-06 14:01:21. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package convert

import (
	"cloud_platform/iot_model/db_product/model"
	proto "cloud_platform/iot_proto/protos/protosService"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func OpmProduct_pb2db(src *proto.OpmProduct) *model.TOpmProduct {
	if src == nil {
		return nil
	}
	dbObj := model.TOpmProduct{
		Id:                src.Id,
		ProductTypeId:     src.ProductTypeId,
		ProductKey:        src.ProductKey,
		Name:              src.Name,
		NameEn:            src.NameEn,
		Identifier:        src.Identifier,
		Model:             src.Model,
		ImageUrl:          src.ImageUrl,
		WifiFlag:          src.WifiFlag,
		NetworkType:       src.NetworkType,
		AttributeType:     src.AttributeType,
		PowerConsumeType:  src.PowerConsumeType,
		Status:            src.Status,
		IsVirtualTest:     src.IsVirtualTest,
		IsScheme:          src.IsScheme,
		Desc:              src.Desc,
		CreatedBy:         src.CreatedBy,
		UpdatedBy:         src.UpdatedBy,
		ProductTypeName:   src.ProductTypeName,
		TenantId:          src.TenantId,
		ControlPanelId:    src.ControlPanelId,
		BaseProductId:     src.ProductId,
		ModuleId:          src.ModuleId,
		FirmwareVersion:   src.FirmwareVersion,
		FirmwareVersionId: src.FirmwareVersionId,
		FirmwareId:        src.FirmwareId,
		PanelProImg:       src.PanelProImg,
		IsShowImg:         src.IsShowImg,
		StyleLinkage:      src.StyleLinkage,
	}
	if src.TslUpdatedAt != nil {
		dbObj.TslUpdatedAt = src.TslUpdatedAt.AsTime()
	}
	return &dbObj
}

func OpmProduct_db2pb(src *model.TOpmProduct) *proto.OpmProduct {
	if src == nil {
		return nil
	}
	pbObj := proto.OpmProduct{
		Id:                src.Id,
		ProductTypeId:     src.ProductTypeId,
		ProductKey:        src.ProductKey,
		Name:              src.Name,
		NameEn:            src.NameEn,
		Identifier:        src.Identifier,
		Model:             src.Model,
		ImageUrl:          src.ImageUrl,
		WifiFlag:          src.WifiFlag,
		NetworkType:       src.NetworkType,
		AttributeType:     src.AttributeType,
		PowerConsumeType:  src.PowerConsumeType,
		Status:            src.Status,
		IsVirtualTest:     src.IsVirtualTest,
		IsScheme:          src.IsScheme,
		Desc:              src.Desc,
		CreatedBy:         src.CreatedBy,
		CreatedAt:         timestamppb.New(src.CreatedAt),
		UpdatedBy:         src.UpdatedBy,
		UpdatedAt:         timestamppb.New(src.UpdatedAt),
		ProductTypeName:   src.ProductTypeName,
		TenantId:          src.TenantId,
		ProductId:         src.BaseProductId,
		ControlPanelId:    src.ControlPanelId,
		ModuleId:          src.ModuleId,
		DeviceNatureKey:   src.DeviceNatureKey,
		FirmwareVersion:   src.FirmwareVersion,
		FirmwareVersionId: src.FirmwareVersionId,
		FirmwareId:        src.FirmwareId,
		PanelProImg:       src.PanelProImg,
		IsShowImg:         src.IsShowImg,
		StyleLinkage:      src.StyleLinkage,
		TslUpdatedAt:      timestamppb.New(src.TslUpdatedAt),
	}
	return &pbObj
}
