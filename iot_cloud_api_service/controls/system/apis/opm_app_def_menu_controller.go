package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/system/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/system/services"
	"cloud_platform/iot_common/iotgin"

	"github.com/gin-gonic/gin"
)

var OemAppDefMenucontroller OemAppDefMenuController

var menuService apiservice.OemAppDefMenuService

type OemAppDefMenuController struct {
} //用户操作控制器

func (OemAppDefMenuController) QueryList(c *gin.Context) {
	var filter entitys.OemAppDefMenuQuery
	err := c.BindQuery(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, total, err := menuService.SetContext(controls.WithUserContext(c)).QueryOemAppDefMenuList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (OemAppDefMenuController) Add(c *gin.Context) {
	var req entitys.OemAppDefMenuEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := menuService.SetContext(controls.WithUserContext(c)).AddOemAppDefMenu(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (OemAppDefMenuController) Edit(c *gin.Context) {
	var req entitys.OemAppDefMenuEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := menuService.SetContext(controls.WithUserContext(c)).EditOemAppDefMenu(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}
