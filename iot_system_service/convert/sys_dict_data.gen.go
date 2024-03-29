// Code generated by sgen.exe,2022-04-18 19:12:08. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package convert

import (
	"cloud_platform/iot_model/db_system/model"
	proto "cloud_platform/iot_proto/protos/protosService"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func SysDictData_pb2db(src *proto.SysDictData) *model.TSysDictData {
	if src == nil {
		return nil
	}
	dbObj := model.TSysDictData{
		DictCode:  src.DictCode,
		DictSort:  src.DictSort,
		DictLabel: src.DictLabel,
		DictValue: src.DictValue,
		DictType:  src.DictType,
		CssClass:  src.CssClass,
		ListClass: src.ListClass,
		IsDefault: src.IsDefault,
		Status:    src.Status,
		CreateBy:  src.CreateBy,
		UpdateBy:  src.UpdateBy,
		Remark:    src.Remark,
		CreatedAt: src.CreatedAt.AsTime(),
		UpdatedAt: src.UpdatedAt.AsTime(),
	}
	return &dbObj
}

func SysDictData_db2pb(src *model.TSysDictData) *proto.SysDictData {
	if src == nil {
		return nil
	}
	pbObj := proto.SysDictData{
		DictCode:  src.DictCode,
		DictSort:  src.DictSort,
		DictLabel: src.DictLabel,
		DictValue: src.DictValue,
		DictType:  src.DictType,
		CssClass:  src.CssClass,
		ListClass: src.ListClass,
		IsDefault: src.IsDefault,
		Status:    src.Status,
		CreateBy:  src.CreateBy,
		UpdateBy:  src.UpdateBy,
		Remark:    src.Remark,
		CreatedAt: timestamppb.New(src.CreatedAt),
		UpdatedAt: timestamppb.New(src.UpdatedAt),
	}
	return &pbObj
}
