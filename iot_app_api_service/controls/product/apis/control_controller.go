package apis

import (
	"cloud_platform/iot_app_api_service/controls/product/services"
	"cloud_platform/iot_common/iotgin"

	"github.com/gin-gonic/gin"
)

var Controlcontroller ControlController

type ControlController struct {
}

var controlServices = services.AppControlService{}

// @Tags 控制页面
// @Summary 检查控制页面是否需要升级
// @Description 检查控制页面是否需要升级
// @Param authorization header  string true "token"
// @Param productId query string true "产品Id"
// @Param panelId query string true "面板Id"
// @Param appPanelType query string true "APP面板类型"
// @success 200 {object} iotgin.ResponseModel "执行结果"
// @Router /v1/platform/app/controlPage/isUpdate [GET]
func (this *ControlController) CheckControlPageVersion(c *gin.Context) {
	productId := c.Query("productId")
	if productId == "" {
		iotgin.ResBadRequest(c, "productId")
		return
	}
	panelId := c.Query("panelId")
	appPanelType := c.Query("appPanelType")
	res, err := controlServices.CheckControlPageVersion(productId, panelId, appPanelType)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}
