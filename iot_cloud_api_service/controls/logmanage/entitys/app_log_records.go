package entitys

import (
	"cloud_platform/iot_common/iotutil"
	proto "cloud_platform/iot_proto/protos/protosService"
	"time"
)

type AppLogRecords struct {
	Id             string            `json:"id"`
	Account        string            `json:"account"`
	AppKey         string            `json:"appKey"`
	TenantId       string            `json:"tenantId"`
	RegionServerId int64             `json:"regionServerId,string"`
	LogType        string            `json:"logType"`
	EventName      string            `json:"eventName"`
	Details        map[string]string `json:"details"`
	CreatedAt      int64             `json:"createdAt"`
}

func AppLogRecordsPb2Db(pb *proto.AppLogRecords) *AppLogRecords {
	return &AppLogRecords{
		Id:             iotutil.ToString(pb.Id),
		Account:        pb.Account,
		AppKey:         pb.AppKey,
		TenantId:       pb.TenantId,
		RegionServerId: pb.RegionServerId,
		LogType:        pb.LogType,
		EventName:      pb.EventName,
		Details:        pb.Details,
		CreatedAt:      pb.CreatedAt.Seconds,
	}
}

type AppLogRecordsListQuery struct {
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
	EventName string `json:"eventName"`
	LogType   string `json:"logType"`
}

type AppLogRecordsListReq struct {
	Account        string                  `json:"account"`
	AppKey         string                  `json:"appKey"`
	TenantId       string                  `json:"tenantId"`
	RegionServerId int64                   `json:"regionServerId,string"`
	Page           int64                   `json:"page"`
	Limit          int64                   `json:"limit"`
	Query          *AppLogRecordsListQuery `json:"query"`
}

func AppLogRecordsListReqDB2Pb(al *AppLogRecordsListReq) *proto.AppLogRecordsListReq {
	var startTime, endTime time.Time
	if al.Query.StartTime == al.Query.EndTime {
		startTime = iotutil.GetTodaySartTime(time.Unix(al.Query.StartTime, 0))
		endTime = iotutil.GetTodayLastTime(time.Unix(al.Query.EndTime, 0))
	} else {
		startTime = time.Unix(al.Query.StartTime, 0)
		endTime = time.Unix(al.Query.EndTime, 0)
	}
	return &proto.AppLogRecordsListReq{
		Account:        al.Account,
		AppKey:         al.AppKey,
		TenantId:       al.TenantId,
		RegionServerId: al.RegionServerId,
		Page:           al.Page,
		Limit:          al.Limit,
		Query: &proto.AppLogRecordsQuery{
			StartTime: iotutil.GetLocalTimeStr(startTime),
			EndTime:   iotutil.GetLocalTimeStr(endTime),
			EventName: al.Query.EventName,
			LogType:   al.Query.LogType,
		},
	}
}
