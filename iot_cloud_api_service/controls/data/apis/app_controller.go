package apis

import (
	"cloud_platform/iot_cloud_api_service/controls/data/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"

	"github.com/gin-gonic/gin"
)

var AppdataController AppDataController

type AppDataController struct{} //部门操作控制器

func (AppDataController) getAppList(c *gin.Context) {
	var req entitys.AppQueryEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	rpcReq := protosService.PmAppDataListRequest{
		Page:      int64(req.Page),
		PageSize:  int64(req.Limit),
		OrderKey:  req.SortField,
		OrderDesc: req.Sort,
		SearchKey: req.SearchKey,
	}
	rpcReq.Query = &protosService.PmAppData{AppName: req.Query.AppName, DevAccount: req.Query.UserName}

	grpcRes, err := rpc.ClientAppDataServiceService.Lists(c, &rpcReq)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if grpcRes.Code != 200 {
		iotgin.ResErrCli(c, errors.New(grpcRes.Message))
		return
	}
	list := make([]*entitys.AppEntitys, 0, len(grpcRes.Data))
	for _, v := range grpcRes.Data {
		list = append(list, entitys.PmAppData_pb2e(v))
	}
	iotgin.ResPageSuccess(c, list, grpcRes.Total, int(req.Page))
}

func (AppDataController) getAppDetail(c *gin.Context) {
	userId := c.Query("appId")
	//userId := "8045219064500551680"
	req := protosService.AppDataDetailFilter{AppId: userId}
	grpcRes, err := rpc.ClientStatisticsService.GetAppDataDetail(context.Background(), &req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	resp := entitys.AppDetailEntitys{
		Account:           grpcRes.Account,
		AppName:           grpcRes.AppName,
		AppType:           grpcRes.AppType,
		RegisterUserTotal: grpcRes.RegisterUserTotal,
		AcitveUserTotal:   grpcRes.AcitveUserTotal,
	}
	for _, v := range grpcRes.VersionList {
		ver := entitys.VersionList{
			AppVersion:  v.AppVersion,
			DevStatus:   v.DevStatus,
			BuildNumber: v.BuildNumber,
			LastOptTime: v.LastOptTime,
			LastOptUser: v.LastOptUser,
		}
		resp.AppVersionList = append(resp.AppVersionList, ver)
	}
	iotgin.ResSuccess(c, resp)
}
