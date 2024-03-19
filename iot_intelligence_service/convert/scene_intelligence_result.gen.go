// Code generated by sgen.exe,2022-05-20 13:36:03. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package convert

import (
	"cloud_platform/iot_model/db_device/model"
	proto "cloud_platform/iot_proto/protos/protosService"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func SceneIntelligenceResult_pb2db(src *proto.SceneIntelligenceResult) *model.TSceneIntelligenceResult {
	if src == nil {
		return nil
	}
	dbObj := model.TSceneIntelligenceResult{
		Id:             src.Id,
		RunTime:        src.RunTime.AsTime(),
		IntelligenceId: src.IntelligenceId,
		RunStatus:      src.RunStatus,
	}
	return &dbObj
}

func SceneIntelligenceResult_db2pb(src *model.TSceneIntelligenceResult) *proto.SceneIntelligenceResult {
	if src == nil {
		return nil
	}
	pbObj := proto.SceneIntelligenceResult{
		Id:             src.Id,
		RunTime:        timestamppb.New(src.RunTime),
		IntelligenceId: src.IntelligenceId,
		RunStatus:      src.RunStatus,
	}
	return &pbObj
}
