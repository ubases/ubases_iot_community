package convert

import (
	models "cloud_platform/iot_model/ch_log/model"
	proto "cloud_platform/iot_proto/protos/protosService"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func AppLogUserPb2Db(pb *proto.AppLogUser) *models.AppLogUser {
	return &models.AppLogUser{
		Id:             pb.Id,
		Account:        pb.Account,
		AppKey:         pb.AppKey,
		TenantId:       pb.TenantId,
		AppName:        pb.AppName,
		Region:         pb.Region,
		RegionServerId: pb.RegionServerId,
		LoginTime:      pb.LoginTime.AsTime(),
		CreatedAt:      pb.CreatedAt.AsTime(),
	}
}

func AppLogUserDb2Pb(db *models.AppLogUser) *proto.AppLogUser {
	return &proto.AppLogUser{
		Id:             db.Id,
		Account:        db.Account,
		AppKey:         db.AppKey,
		TenantId:       db.TenantId,
		AppName:        db.AppName,
		Region:         db.Region,
		RegionServerId: db.RegionServerId,
		LoginTime:      timestamppb.New(db.LoginTime),
		CreatedAt:      timestamppb.New(db.CreatedAt),
	}
}
