// Code generated by sgen.exe,2022-11-11 10:46:48. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package convert

import (
	"cloud_platform/iot_model/db_product/model"
	proto "cloud_platform/iot_proto/protos/protosService"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func OpmProductMaterials_pb2db(src *proto.OpmProductMaterials) *model.TOpmProductMaterials {
	if src == nil {
		return nil
	}
	dbObj := model.TOpmProductMaterials{
		Id:             src.Id,
		TenantId:       src.TenantId,
		ImageUrl:       src.ImageUrl,
		ProductPage:    src.ProductPage,
		Count:          src.Count,
		BrandCode:      src.BrandCode,
		FragranceCode:  src.FragranceCode,
		CreatedBy:      src.CreatedBy,
		CreatedAt:      src.CreatedAt.AsTime(),
		UpdatedBy:      src.UpdatedBy,
		UpdatedAt:      src.UpdatedAt.AsTime(),
		MaterialTypeId: src.MaterialTypeId,
	}
	return &dbObj
}

func OpmProductMaterials_db2pb(src *model.TOpmProductMaterials) *proto.OpmProductMaterials {
	if src == nil {
		return nil
	}
	pbObj := proto.OpmProductMaterials{
		Id:            src.Id,
		TenantId:      src.TenantId,
		ImageUrl:      src.ImageUrl,
		ProductPage:   src.ProductPage,
		Count:         src.Count,
		BrandCode:     src.BrandCode,
		FragranceCode: src.FragranceCode,
		CreatedBy:     src.CreatedBy,
		CreatedAt:     timestamppb.New(src.CreatedAt),
		UpdatedBy:     src.UpdatedBy,
		UpdatedAt:     timestamppb.New(src.UpdatedAt),
	}
	return &pbObj
}
