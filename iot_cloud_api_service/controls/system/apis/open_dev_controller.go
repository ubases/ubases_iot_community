package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/system/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/system/services"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"

	"github.com/gin-gonic/gin"
)

var OpenDevcontroller OpenDevController

// 开发者认证
type OpenDevController struct{} //菜单操作控制器

var openDevServices = apiservice.OpenDevService{}

func (OpenDevController) QueryList(c *gin.Context) {
	var filter entitys.OpenDevListReq
	err := c.BindQuery(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, total, err := openDevServices.GetOpenDevList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.PageNum))
}

func (OpenDevController) QueryDetail(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	res, err := openDevServices.GetOpenDevDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (OpenDevController) OpenDevAuth(c *gin.Context) {
	//OpenDevCompanyAuthReq

	var filter entitys.OpenDevCompanyAuthReq
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	userId, _ := c.Get("userId")
	res, err := openDevServices.SetContext(controls.WithOpenUserContext(c)).OpenDevAuth(filter, iotutil.ToString(userId))
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}
