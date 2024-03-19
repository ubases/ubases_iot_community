package apis

import (
	"cloud_platform/iot_cloud_api_service/controls/system/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/system/services"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"

	"github.com/gin-gonic/gin"
)

var OpenAuthrulecontroller OpenAuthruleController

type OpenAuthruleController struct{} //菜单操作控制器

var openAuthruleServices = apiservice.OpenAuthRuleService{}

func (OpenAuthruleController) QueryList(c *gin.Context) {
	var filter entitys.OpenAuthRuleQuery
	err := c.ShouldBindQuery(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := openAuthruleServices.QueryAuthRuleList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (OpenAuthruleController) QueryDetail(c *gin.Context) {
	id := c.Query("menuId")
	if id == "" {
		iotgin.ResBadRequest(c, "menuId")
		return
	}
	res, err := openAuthruleServices.GetAuthRuleDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (OpenAuthruleController) Edit(c *gin.Context) {
	var req entitys.OpenAuthRuleEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	//赋值公共字段
	userid, _ := c.Get("userId")
	entitys.OpenAuthRule_SetCommonFiled(&req, iotutil.ToInt64(userid), 2)

	id, err := openAuthruleServices.UpdateAuthRule(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (OpenAuthruleController) Add(c *gin.Context) {
	var req entitys.OpenAuthRuleEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	//赋值公共字段
	userid, _ := c.Get("userId")
	entitys.OpenAuthRule_SetCommonFiled(&req, iotutil.ToInt64(userid), 1)
	id, err := openAuthruleServices.AddAuthRule(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (OpenAuthruleController) Delete(c *gin.Context) {
	var req entitys.DeleteCommonQuery
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if len(req.Ids) == 0 {
		iotgin.ResBadRequest(c, "id")
		return
	}
	err = openAuthruleServices.DeleteAuthRule(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}
