package apis

import (
	"cloud_platform/iot_cloud_api_service/controls/system/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/system/services"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"

	"github.com/gin-gonic/gin"
)

var Authrulecontroller AuthruleController

type AuthruleController struct{} //菜单操作控制器

var authruleServices = apiservice.SysAuthRuleService{}

func (AuthruleController) QueryList(c *gin.Context) {
	var filter entitys.SysAuthRuleQuery
	err := c.ShouldBindQuery(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := authruleServices.QueryAuthRuleList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (AuthruleController) QueryDetail(c *gin.Context) {
	id := c.Query("menuId")
	if id == "" {
		iotgin.ResBadRequest(c, "menuId")
		return
	}
	res, err := authruleServices.GetAuthRuleDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (AuthruleController) Edit(c *gin.Context) {
	var req entitys.SysAuthRuleEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	//赋值公共字段
	userid, _ := c.Get("userId")
	entitys.SysAuthRule_SetCommonFiled(&req, iotutil.ToInt64(userid), 2)

	id, err := authruleServices.UpdateAuthRule(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (AuthruleController) Add(c *gin.Context) {
	var req entitys.SysAuthRuleEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	//赋值公共字段
	userid, _ := c.Get("userId")
	entitys.SysAuthRule_SetCommonFiled(&req, iotutil.ToInt64(userid), 1)
	id, err := authruleServices.AddAuthRule(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (AuthruleController) Delete(c *gin.Context) {
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
	err = authruleServices.DeleteAuthRule(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}
