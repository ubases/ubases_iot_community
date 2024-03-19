package convert

import (
	"cloud_platform/iot_model/db_product/model"
	proto "cloud_platform/iot_proto/protos/protosService"
)

func PmProduct_db2pb(src *model.TPmProduct) *proto.BaseProductDetail {
	if src == nil {
		return nil
	}
	pbObj := proto.BaseProductDetail{
		Id:            src.Id,
		ProductTypeId: src.ProductTypeId,
		ProductKey:    src.ProductKey,
		Name:          src.Name,
		NameEn:        src.NameEn,
		Identifier:    src.Identifier,
		Model:         src.Model,
		ImageUrl:      src.ImageUrl,
		WifiFlag:      src.WifiFlag,
		NetworkType:   src.NetworkType,
		AttributeType: src.AttributeType,
		Status:        src.Status,
		IsVirtualTest: src.IsVirtualTest,
	}
	return &pbObj
}
