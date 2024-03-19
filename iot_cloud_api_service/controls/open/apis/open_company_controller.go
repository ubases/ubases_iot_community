package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/open/services"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"

	"github.com/gin-gonic/gin"
)

var OpenCompanycontroller OpenUserController

var OpenCompanyService apiservice.OpenCompanyService

type OpenCompanyController struct {
} //用户操作控制器

// 获取企业认证消息
func (OpenUserController) GetCompanyAuth(c *gin.Context) {
	res, err := OpenCompanyService.SetContext(controls.WithOpenUserContext(c)).GetCompanyAuth()
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 获取基础信息
func (OpenUserController) GetBaseInfo(c *gin.Context) {
	userId, _ := c.Get("userId")
	userName, _ := c.Get("nickName")
	res, err := OpenCompanyService.SetContext(controls.WithUserContext(c)).GetBaseInfo(iotutil.ToInt64(userId))
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res.UserName = iotutil.ToString(userName)
	iotgin.ResSuccess(c, res)
}

// 企业认证提交审核
func (OpenUserController) CompanyAuth(c *gin.Context) {

	var req entitys.OpenCompanyAuthReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := OpenCompanyService.SetContext(controls.WithUserContext(c)).CompanyAuth(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 修改公司名称
func (OpenUserController) CompanyChangeName(c *gin.Context) {
	var req entitys.OpenCompanyChangeNameReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := OpenCompanyService.SetContext(controls.WithUserContext(c)).CompanyChangeName(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 修改基础信息
func (OpenUserController) UpdateBaseInfo(c *gin.Context) {

	var req entitys.OpenCompanyBaseReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	userId := c.GetInt64("UserId")
	res, err := OpenCompanyService.SetContext(controls.WithUserContext(c)).UpdateBaseInfo(req, userId)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 获取基础信息
func (OpenUserController) GetCompanyInfo(c *gin.Context) {
	tenantId, _ := c.Get("tenantId")
	res, err := OpenCompanyService.SetContext(controls.WithUserContext(c)).GetCompanyInfo(iotutil.ToString(tenantId))
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}
