// Code generated by sgen.exe,2022-04-21 12:44:21. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package convert

import (
	"cloud_platform/iot_model/db_product/model"
	proto "cloud_platform/iot_proto/protos/protosService"
)

func PmProductModuleRelation_pb2db(src *proto.PmProductModuleRelation) *model.TPmProductModuleRelation {
	if src == nil {
		return nil
	}
	dbObj := model.TPmProductModuleRelation{
		Id:        src.Id,
		ProductId: src.ProductId,
		ModuleId:  src.ModuleId,
	}
	return &dbObj
}

func PmProductModuleRelation_db2pb(src *model.TPmProductModuleRelation) *proto.PmProductModuleRelation {
	if src == nil {
		return nil
	}
	pbObj := proto.PmProductModuleRelation{
		Id:        src.Id,
		ProductId: src.ProductId,
		ModuleId:  src.ModuleId,
	}
	return &pbObj
}
