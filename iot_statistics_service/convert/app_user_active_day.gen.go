// Code generated by sgen.exe,2022-08-10 22:27:08. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package convert

import (
	"cloud_platform/iot_model/db_statistics/model"
	proto "cloud_platform/iot_proto/protos/protosService"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func AppUserActiveDay_pb2db(src *proto.AppUserActiveDay) *model.TAppUserActiveDay {
	if src == nil {
		return nil
	}
	dbObj := model.TAppUserActiveDay{
		DataTime:  src.DataTime.AsTime(),
		TenantId:  src.TenantId,
		AppKey:    src.AppKey,
		ActiveSum: src.ActiveSum,
		UpdatedAt: src.UpdatedAt.AsTime(),
	}
	return &dbObj
}

func AppUserActiveDay_db2pb(src *model.TAppUserActiveDay) *proto.AppUserActiveDay {
	if src == nil {
		return nil
	}
	pbObj := proto.AppUserActiveDay{
		DataTime:  timestamppb.New(src.DataTime),
		TenantId:  src.TenantId,
		AppKey:    src.AppKey,
		ActiveSum: src.ActiveSum,
		UpdatedAt: timestamppb.New(src.UpdatedAt),
	}
	return &pbObj
}
