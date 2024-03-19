// Code generated by sgen.exe,2022-12-23 15:26:00. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package convert

import (
	"cloud_platform/iot_model/db_device/model"
	proto "cloud_platform/iot_proto/protos/protosService"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func IotJob_pb2db(src *proto.IotJob) *model.TIotJob {
	if src == nil {
		return nil
	}
	dbObj := model.TIotJob{
		Id:             src.Id,
		ProductKey:     src.ProductKey,
		DeviceId:       src.DeviceId,
		TaskId:         src.TaskId,
		TaskType:       src.TaskType,
		Enabled:        src.Enabled,
		Cron:           src.Cron,
		Data:           src.Data,
		CreatedBy:      src.CreatedBy,
		UpdatedBy:      src.UpdatedBy,
		EndData:        src.EndData,
		EndCron:        src.EndCron,
		RegionServerId: src.RegionServerId,
		Timezone:       src.Timezone,
	}
	return &dbObj
}

func IotJob_db2pb(src *model.TIotJob) *proto.IotJob {
	if src == nil {
		return nil
	}
	pbObj := proto.IotJob{
		Id:             src.Id,
		ProductKey:     src.ProductKey,
		DeviceId:       src.DeviceId,
		TaskId:         src.TaskId,
		TaskType:       src.TaskType,
		Enabled:        src.Enabled,
		Cron:           src.Cron,
		Data:           src.Data,
		CreatedBy:      src.CreatedBy,
		UpdatedBy:      src.UpdatedBy,
		CreatedAt:      timestamppb.New(src.CreatedAt),
		UpdatedAt:      timestamppb.New(src.UpdatedAt),
		EndData:        src.EndData,
		EndCron:        src.EndCron,
		RegionServerId: src.RegionServerId,
		Timezone:       src.Timezone,
	}
	return &pbObj
}
