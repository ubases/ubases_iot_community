// convert包，主要存放数据库结构与protobuf结构之间的转换函数
// 以下是示例代码
package convert

import (
	"cloud_platform/iot_model/db_config/model"
	proto "cloud_platform/iot_proto/protos/protosService"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// protobuf结构转数据库结构
func ConfigDictType_pb2db(src *proto.ConfigDictType) *model.TConfigDictType {
	if src == nil {
		return nil
	}
	dbObj := model.TConfigDictType{
		DictId:    src.DictId,
		DictName:  src.DictName,
		DictType:  src.DictType,
		Status:    src.Status,
		ValueType: src.ValueType,
		Remark:    src.Remark,
		IsSystem:  src.IsSystem,
		CreatedBy: src.CreatedBy,
		UpdatedBy: src.UpdatedBy,
	}
	return &dbObj
}

// 数据库结构转protobuf结构
func ConfigDictType_db2pb(src *model.TConfigDictType) *proto.ConfigDictType {
	if src == nil {
		return nil
	}
	pbObj := proto.ConfigDictType{
		DictId:    src.DictId,
		DictName:  src.DictName,
		DictType:  src.DictType,
		Status:    src.Status,
		ValueType: src.ValueType,
		Remark:    src.Remark,
		IsSystem:  src.IsSystem,
		CreatedBy: src.CreatedBy,
		UpdatedBy: src.UpdatedBy,
		CreatedAt: timestamppb.New(src.CreatedAt),
		UpdatedAt: timestamppb.New(src.UpdatedAt),
	}
	return &pbObj
}
