package apis

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_proto/protos/protosService"
	"errors"

	"github.com/gin-gonic/gin"
)

var Appcontroller AppController

type AppController struct {
}

// GetAppDetailByApp 获取APP信息
// @Summary 获取APP信息
// @Description 获取APP信息
// @Tags 通用
// @Param appKey header string true "APP Key"
// @Success 200 {object} iotgin.ResponseModel "{"code": 0, "data": [...]}"
// @Router /v1/platform/app/common/appInfo [get]
func (AppController) GetAppDetailByApp(c *gin.Context) {
	appKey := controls.GetAppKey(c)
	if appKey == "" {
		iotgin.ResBadRequest(c, "header.appKey")
		return
	}
	res, err := rpc.ClientOemAppService.Find(controls.WithUserContext(c), &protosService.OemAppFilter{AppKey: appKey})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if res.Code != 200 {
		iotgin.ResErrCli(c, errors.New(res.Message))
		return
	}
	resMap := map[string]interface{}{
		"appName":   res.Data[0].Name,
		"appEnName": res.Data[0].NameEn,
		"version":   res.Data[0].Version,
		"channel":   res.Data[0].Channel,
	}
	iotgin.ResSuccess(c, resMap)
}
