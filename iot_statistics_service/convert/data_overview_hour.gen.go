// Code generated by sgen.exe,2022-06-17 09:58:13. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package convert

import (
	"cloud_platform/iot_model/db_statistics/model"
	proto "cloud_platform/iot_proto/protos/protosService"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func DataOverviewHour_pb2db(src *proto.DataOverviewHour) *model.TDataOverviewHour {
	if src == nil {
		return nil
	}
	dbObj := model.TDataOverviewHour{
		DataTime:             src.DataTime.AsTime(),
		TenantId:             src.TenantId,
		DeviceActiveSum:      src.DeviceActiveSum,
		DeviceFaultSum:       src.DeviceFaultSum,
		DeveloperRegisterSum: src.DeveloperRegisterSum,
		UserRegisterSum:      src.UserRegisterSum,
		UpdatedAt:            src.UpdatedAt.AsTime(),
	}
	return &dbObj
}

func DataOverviewHour_db2pb(src *model.TDataOverviewHour) *proto.DataOverviewHour {
	if src == nil {
		return nil
	}
	pbObj := proto.DataOverviewHour{
		DataTime:             timestamppb.New(src.DataTime),
		TenantId:             src.TenantId,
		DeviceActiveSum:      src.DeviceActiveSum,
		DeviceFaultSum:       src.DeviceFaultSum,
		DeveloperRegisterSum: src.DeveloperRegisterSum,
		UserRegisterSum:      src.UserRegisterSum,
		UpdatedAt:            timestamppb.New(src.UpdatedAt),
	}
	return &pbObj
}
