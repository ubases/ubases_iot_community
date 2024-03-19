package apis

import (
	"cloud_platform/iot_cloud_api_service/controls/data/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_proto/protos/protosService"
	"errors"

	"github.com/gin-gonic/gin"
)

var Openappusercontroller OpenAppUserController

type OpenAppUserController struct {
}

func (OpenAppUserController) getUserAppStatistics(c *gin.Context) {
	appKey := c.Query("appKey")
	req := protosService.AppUserStatisticsFilter{AppKey: appKey}
	resp, err := rpc.ClientAppUserStatisticsService.GetAppUserStatistics(c, &req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Data == nil || resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	ret := entitys.OpenAppUserEntitys{
		AppUserTodayActive: resp.Data.AppUserTodayActive,
		AppUserToday:       resp.Data.AppUserToday,
		AppUserAll:         resp.Data.AppUserAll,
		AppUser:            entitys.Data{Total: resp.Data.AppUser.Total},
		ActiveUser:         entitys.Data{Total: resp.Data.ActiveUser.Total},
	}
	for _, v := range resp.Data.AppUser.GetData() {
		ret.AppUser.Data = append(ret.AppUser.Data, entitys.TimeData{Time: v.Time, Total: v.Total})
	}
	for _, v := range resp.Data.ActiveUser.GetData() {
		ret.ActiveUser.Data = append(ret.ActiveUser.Data, entitys.TimeData{Time: v.Time, Total: v.Total})
	}
	iotgin.ResSuccess(c, ret)
}
