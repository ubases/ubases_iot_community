// Code generated by sgen.exe,2022-04-27 10:55:25. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package convert

import (
	"cloud_platform/iot_model/db_open_system/model"
	proto "cloud_platform/iot_proto/protos/protosService"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func OpenModelInfo_pb2db(src *proto.OpenModelInfo) *model.TOpenModelInfo {
	if src == nil {
		return nil
	}
	dbObj := model.TOpenModelInfo{
		ModelId:         src.ModelId,
		ModelCategoryId: src.ModelCategoryId,
		ModelName:       src.ModelName,
		ModelTitle:      src.ModelTitle,
		ModelPk:         src.ModelPk,
		ModelOrder:      src.ModelOrder,
		ModelSort:       src.ModelSort,
		ModelList:       src.ModelList,
		ModelEdit:       src.ModelEdit,
		ModelIndexes:    src.ModelIndexes,
		SearchList:      src.SearchList,
		CreateTime:      src.CreateTime,
		UpdateTime:      src.UpdateTime,
		ModelStatus:     src.ModelStatus,
		ModelEngine:     src.ModelEngine,
		CreateBy:        src.CreateBy,
		UpdateBy:        src.UpdateBy,
		CreatedAt:       src.CreatedAt.AsTime(),
		UpdatedAt:       src.UpdatedAt.AsTime(),
	}
	return &dbObj
}

func OpenModelInfo_db2pb(src *model.TOpenModelInfo) *proto.OpenModelInfo {
	if src == nil {
		return nil
	}
	pbObj := proto.OpenModelInfo{
		ModelId:         src.ModelId,
		ModelCategoryId: src.ModelCategoryId,
		ModelName:       src.ModelName,
		ModelTitle:      src.ModelTitle,
		ModelPk:         src.ModelPk,
		ModelOrder:      src.ModelOrder,
		ModelSort:       src.ModelSort,
		ModelList:       src.ModelList,
		ModelEdit:       src.ModelEdit,
		ModelIndexes:    src.ModelIndexes,
		SearchList:      src.SearchList,
		CreateTime:      src.CreateTime,
		UpdateTime:      src.UpdateTime,
		ModelStatus:     src.ModelStatus,
		ModelEngine:     src.ModelEngine,
		CreateBy:        src.CreateBy,
		UpdateBy:        src.UpdateBy,
		CreatedAt:       timestamppb.New(src.CreatedAt),
		UpdatedAt:       timestamppb.New(src.UpdatedAt),
	}
	return &pbObj
}
