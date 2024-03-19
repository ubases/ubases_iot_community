package apis

import (
	"cloud_platform/iot_cloud_api_service/controls/config/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/config/services"

	"github.com/gin-gonic/gin"

	"cloud_platform/iot_common/iotgin"
)

var ConfigPlatformcontroller ConfigPlatformController

type ConfigPlatformController struct{} //部门操作控制器

var configPlatformServices = apiservice.ConfigPlatformService{}

func (ConfigPlatformController) QueryList(c *gin.Context) {
	var filter entitys.ConfigPlatformQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, total, err := configPlatformServices.QueryConfigPlatformList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (ConfigPlatformController) QueryDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	res, err := configPlatformServices.GetConfigPlatformDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (ConfigPlatformController) Edit(c *gin.Context) {
	var req entitys.ConfigPlatformEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := configPlatformServices.UpdateConfigPlatform(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (ConfigPlatformController) Add(c *gin.Context) {
	var req entitys.ConfigPlatformEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := configPlatformServices.AddConfigPlatform(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (ConfigPlatformController) Delete(c *gin.Context) {
	var req entitys.ConfigPlatformFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = configPlatformServices.DeleteConfigPlatform(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}
