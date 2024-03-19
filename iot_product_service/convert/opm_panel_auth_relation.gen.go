// Code generated by sgen,2023-06-02 13:48:11. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package convert

import (
	"cloud_platform/iot_model/db_product/model"
	proto "cloud_platform/iot_proto/protos/protosService"
)

func OpmPanelAuthRelation_pb2db(src *proto.OpmPanelAuthRelation) *model.TOpmPanelAuthRelation {
	if src == nil {
		return nil
	}
	dbObj := model.TOpmPanelAuthRelation{
		Id:            src.Id,
		ProductId:     src.ProductId,
		PanelId:       src.PanelId,
		ProductTypeId: src.ProductTypeId,
	}
	return &dbObj
}

func OpmPanelAuthRelation_db2pb(src *model.TOpmPanelAuthRelation) *proto.OpmPanelAuthRelation {
	if src == nil {
		return nil
	}
	pbObj := proto.OpmPanelAuthRelation{
		Id:            src.Id,
		ProductId:     src.ProductId,
		PanelId:       src.PanelId,
		ProductTypeId: src.ProductTypeId,
	}
	return &pbObj
}
