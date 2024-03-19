// Code generated by sgen.exe,2022-04-27 10:55:25. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package convert

import (
	"cloud_platform/iot_model/db_open_system/model"
	proto "cloud_platform/iot_proto/protos/protosService"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func OpenConfig_pb2db(src *proto.OpenConfig) *model.TOpenConfig {
	if src == nil {
		return nil
	}
	dbObj := model.TOpenConfig{
		ConfigId:    src.ConfigId,
		ConfigName:  src.ConfigName,
		ConfigKey:   src.ConfigKey,
		ConfigValue: src.ConfigValue,
		ConfigType:  src.ConfigType,
		CreateBy:    src.CreateBy,
		UpdateBy:    src.UpdateBy,
		Remark:      src.Remark,
		CreatedAt:   src.CreatedAt.AsTime(),
		UpdatedAt:   src.UpdatedAt.AsTime(),
		CreatedBy:   src.CreatedBy,
		UpdatedBy:   src.UpdatedBy,
	}
	return &dbObj
}

func OpenConfig_db2pb(src *model.TOpenConfig) *proto.OpenConfig {
	if src == nil {
		return nil
	}
	pbObj := proto.OpenConfig{
		ConfigId:    src.ConfigId,
		ConfigName:  src.ConfigName,
		ConfigKey:   src.ConfigKey,
		ConfigValue: src.ConfigValue,
		ConfigType:  src.ConfigType,
		CreateBy:    src.CreateBy,
		UpdateBy:    src.UpdateBy,
		Remark:      src.Remark,
		CreatedAt:   timestamppb.New(src.CreatedAt),
		UpdatedAt:   timestamppb.New(src.UpdatedAt),
		CreatedBy:   src.CreatedBy,
		UpdatedBy:   src.UpdatedBy,
	}
	return &pbObj
}
