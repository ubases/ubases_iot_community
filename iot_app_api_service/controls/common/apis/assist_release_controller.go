package apis

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotgin"
	proto "cloud_platform/iot_proto/protos/protosService"
	"context"

	"github.com/gin-gonic/gin"
)

//辅助上架

type AssistReleaseController struct {
}

var AssistReleasecontroller AssistReleaseController

// CheckSkin APP临时皮肤
// @Summary APP临时皮肤
// @Description APP临时皮肤
// @Tags 通用
// @Param version query string true "APP版本"
// @Param appKey header string true "APP Key"
// @Param tenantId header string true "租户Id"
// @Success 200 {object} iotgin.ResponseModel "{"code": 0, "data": [...]}"
// @Router /v1/platform/app/common/checkSkin [get]
func (s AssistReleaseController) CheckSkin(c *gin.Context) {
	version := c.Query("version")
	if version == "" {
		iotgin.ResBadRequest(c, "version")
		return
	}
	appKey := controls.GetAppKey(c)
	tenantId := controls.GetTenantId(c)

	req, err := rpc.ClientOemAppAssistReleaseService.Find(context.Background(), &proto.OemAppAssistReleaseFilter{
		TenantId:    tenantId,
		AppKey:      appKey,
		AppVersion:  version,
		IsCheckSkin: true,
		//Status:      1,
	})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Code != 200 || len(req.Data) == 0 {
		iotgin.ResSuccess(c, false)
		return
	}
	iotgin.ResSuccess(c, map[string]interface{}{
		"skinKey":   req.Data[0].SkinKey,
		"skinName":  req.Data[0].SkinName,
		"startTime": req.Data[0].StartTime.AsTime().Unix(),
		"endTime":   req.Data[0].EndTime.AsTime().Unix(),
	})
}
