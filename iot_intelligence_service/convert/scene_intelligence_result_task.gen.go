// Code generated by sgen.exe,2022-05-20 13:36:04. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package convert

import (
	"cloud_platform/iot_model/db_device/model"
	proto "cloud_platform/iot_proto/protos/protosService"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func SceneIntelligenceResultTask_pb2db(src *proto.SceneIntelligenceResultTask) *model.TSceneIntelligenceResultTask {
	if src == nil {
		return nil
	}
	dbObj := model.TSceneIntelligenceResultTask{
		Id:             src.Id,
		StartTime:      src.StartTime.AsTime(),
		EndTime:        src.EndTime.AsTime(),
		IntelligenceId: src.IntelligenceId,
		ResultId:       src.ResultId,
		IsSuccess:      src.IsSuccess,
		ResultMsg:      src.ResultMsg,
		TaskId:         src.TaskId,
		TaskImg:        src.TaskImg,
		TaskDesc:       src.TaskDesc,
		TaskType:       src.TaskType,
		ObjectId:       src.ObjectId,
		ObjectDesc:     src.ObjectDesc,
		FuncKey:        src.FuncKey,
		FuncDesc:       src.FuncDesc,
		FuncValue:      src.FuncValue,
		ProductKey:     src.ProductKey,
		CreatedAt:      src.CreatedAt.AsTime(),
		UpdatedAt:      src.UpdatedAt.AsTime(),
	}
	return &dbObj
}

func SceneIntelligenceResultTask_db2pb(src *model.TSceneIntelligenceResultTask) *proto.SceneIntelligenceResultTask {
	if src == nil {
		return nil
	}
	pbObj := proto.SceneIntelligenceResultTask{
		Id:             src.Id,
		StartTime:      timestamppb.New(src.StartTime),
		EndTime:        timestamppb.New(src.EndTime),
		IntelligenceId: src.IntelligenceId,
		ResultId:       src.ResultId,
		IsSuccess:      src.IsSuccess,
		ResultMsg:      src.ResultMsg,
		TaskId:         src.TaskId,
		TaskImg:        src.TaskImg,
		TaskDesc:       src.TaskDesc,
		TaskType:       src.TaskType,
		ObjectId:       src.ObjectId,
		ObjectDesc:     src.ObjectDesc,
		FuncKey:        src.FuncKey,
		FuncDesc:       src.FuncDesc,
		FuncValue:      src.FuncValue,
		Functions:      src.Functions,
		ProductKey:     src.ProductKey,
		CreatedAt:      timestamppb.New(src.CreatedAt),
		UpdatedAt:      timestamppb.New(src.UpdatedAt),
	}
	return &pbObj
}
