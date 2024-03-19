package convert

import (
	models "cloud_platform/iot_model/ch_log/model"
	proto "cloud_platform/iot_proto/protos/protosService"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func AppLogRecordsPb2Db(pb *proto.AppLogRecords) *models.AppLogRecords {
	return &models.AppLogRecords{
		Id:             pb.Id,
		Account:        pb.Account,
		AppKey:         pb.AppKey,
		TenantId:       pb.TenantId,
		RegionServerId: pb.RegionServerId,
		LogType:        pb.LogType,
		EventName:      pb.EventName,
		Details:        pb.Details,
		CreatedAt:      pb.CreatedAt.AsTime(),
	}
}

func AppLogRecordsDb2Pb(db *models.AppLogRecords) *proto.AppLogRecords {
	return &proto.AppLogRecords{
		Id:             db.Id,
		Account:        db.Account,
		AppKey:         db.AppKey,
		TenantId:       db.TenantId,
		RegionServerId: db.RegionServerId,
		LogType:        db.LogType,
		EventName:      db.EventName,
		Details:        db.Details,
		CreatedAt:      timestamppb.New(db.CreatedAt),
	}
}
