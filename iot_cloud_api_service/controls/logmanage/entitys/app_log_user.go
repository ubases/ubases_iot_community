package entitys

import (
	"cloud_platform/iot_common/iotutil"
	proto "cloud_platform/iot_proto/protos/protosService"
)

type AppLogUser struct {
	Id             string `json:"id"`
	Account        string `json:"account"`
	AppKey         string `json:"appKey"`
	TenantId       string `json:"tenantId"`
	AppName        string `json:"appName"`
	Region         string `json:"region"`
	RegionServerId int64  `json:"regionServerId,string"`
	LoginTime      int64  `json:"loginTime"`
	CreatedAt      int64  `json:"createdAt"`
}

func AppLogUserPb2Db(pb *proto.AppLogUser) *AppLogUser {
	return &AppLogUser{
		Id:             iotutil.ToString(pb.Id),
		Account:        pb.Account,
		AppKey:         pb.AppKey,
		TenantId:       pb.TenantId,
		AppName:        pb.AppName,
		Region:         pb.Region,
		RegionServerId: pb.RegionServerId,
		LoginTime:      pb.LoginTime.Seconds,
		CreatedAt:      pb.CreatedAt.Seconds,
	}
}
