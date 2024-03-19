package apis

import (
	"cloud_platform/iot_cloud_api_service/controls/data/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"strconv"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/gin-gonic/gin"
)

var DeveloperdataController DeveloperDataController

type DeveloperDataController struct{} //部门操作控制器

func (DeveloperDataController) getDeveloperList(c *gin.Context) {
	var req entitys.DeveloperQueryEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	rpcReq := protosService.DeveloperStatListRequest{
		Page:      int64(req.Page),
		PageSize:  int64(req.Limit),
		OrderKey:  req.SortField,
		OrderDesc: req.Sort,
		SearchKey: req.SearchKey,
	}
	rpcReq.Query = &protosService.DeveloperListFilter{
		UserName:  req.Query.UserName,
		StartTime: timestamppb.New(GetStartTime(req.Query.LastDay)),
		EndTime:   timestamppb.New(time.Now()),
	}

	grpcRes, err := rpc.ClientStatisticsService.GetDeveloperList(context.Background(), &rpcReq)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if grpcRes.Code != 200 {
		iotgin.ResErrCli(c, errors.New(grpcRes.Message))
		return
	}

	//填充在线状态或登录地区
	var userIds []string
	for _, v := range grpcRes.Data {
		userIds = append(userIds, v.UserId)
	}
	mapData := GetDeveloperLoginInfo(userIds)
	list := make([]*entitys.DeveloperEntitys, 0, len(grpcRes.Data))
	for _, v := range grpcRes.Data {
		//在线状态和登录地区，grpc返回的内容固定为online=2,loginaddr=-
		if info, ok := mapData[v.UserId]; ok {
			if time.Now().Before(time.Unix(info.ExpiresAt, 0)) { //未过期表示在线
				v.Online = 1
			}
			v.LoginAddr = info.Addr
		}
		list = append(list, entitys.DeveloperStat_pb2e(v))
	}
	iotgin.ResPageSuccess(c, list, grpcRes.Total, int(req.Page))
}

func (DeveloperDataController) getDeveloperDetail(c *gin.Context) {
	userId := c.Query("userId")
	req := protosService.DeveloperDetailFilter{UserId: userId}
	grpcRes, err := rpc.ClientStatisticsService.GetDeveloperDetail(context.Background(), &req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	resp := entitys.DeveloperDetailEntitys{
		Account:           grpcRes.Account,
		CompanyName:       grpcRes.CompanyName,
		RoleName:          grpcRes.RoleName,
		ActiveDeviceTotal: grpcRes.ActiveDeviceTotal,
		AppTotal:          grpcRes.AppTotal,
	}
	resp.AppList = make([]entitys.AppList, 0, len(grpcRes.AppList))
	for _, v := range grpcRes.AppList {
		app := entitys.AppList{
			AppID:     v.AppId,
			AppName:   v.AppName,
			DevStatus: v.DevStatus,
			Version:   v.Version,
			VerTotal:  v.VerTotal,
		}
		resp.AppList = append(resp.AppList, app)
	}
	iotgin.ResSuccess(c, resp)
}

func (DeveloperDataController) getDeveloperTotal(c *gin.Context) {
	grpcRes, err := rpc.ClientStatisticsService.GetDeveloperStatistics(context.Background(), nil)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	ret := entitys.DeveloperTotalEntitys{DevUserTotal: grpcRes.Total, DevUserOnlineTotal: grpcRes.OnlineTotal}
	iotgin.ResSuccess(c, ret)
}

func GetStartTime(flag int32) time.Time {
	t0 := iotutil.New(time.Now()).BeginningOfDay()
	t := time.Time{}
	switch flag {
	case 1: //今日
		t = t0
	case 2: //近7日
		t = t0.Add(-6 * 24 * time.Hour)
	case 3: //近30日
		t = t0.Add(-29 * 24 * time.Hour)
	case 4: //近60日
		//t = t0.Add(-59 * 24 * time.Hour)
	default:
	}
	return t
}

func GetDeveloperLoginInfo(ids []string) map[string]iotstruct.OpenUserLogin {
	//填充在线状态或登录地区
	mapData := make(map[string]iotstruct.OpenUserLogin)
	logins, err := iotredis.GetClient().HMGet(context.Background(), iotconst.USERLASTLOGIN, ids...).Result()
	if err != nil {
		return mapData
	}
	for _, v := range logins {
		if v == nil {
			continue
		}
		if str, ok := v.(string); ok && str != "" {
			var info iotstruct.OpenUserLogin
			if err := info.UnmarshalBinary([]byte(str)); err == nil && info.UserId > 0 {
				mapData[strconv.Itoa(int(info.UserId))] = info
			}
		}
	}
	return mapData
}
