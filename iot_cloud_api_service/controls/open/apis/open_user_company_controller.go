package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/open/services"
	"cloud_platform/iot_common/iotgin"
	"errors"

	"github.com/gin-gonic/gin"
)

var OpenUserCompanycontroller OpenUserCompanyController

var OpenUserCompanyService apiservice.OpenUserCompanyService

type OpenUserCompanyController struct {
} //用户操作控制器

// 修改基础信息
func (OpenUserCompanyController) UserCompanyAuth(c *gin.Context) {

	var req entitys.OpenUserCompanyAuthReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := OpenUserCompanyService.SetContext(controls.WithOpenUserContext(c)).UserCompanyAuth(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 修改授权备注
func (OpenUserCompanyController) UserCompanyUpdateReamk(c *gin.Context) {

	var req entitys.OpenUserCompanyAuthRemarkReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := OpenUserCompanyService.SetContext(controls.WithOpenUserContext(c)).UserCompanyUpdateReamk(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 修改授权备注
func (OpenUserCompanyController) UserCompanyDelete(c *gin.Context) {
	var req entitys.OpenDeleteReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, errors.New("id 参数错误."))
		return
	}
	res, err := OpenUserCompanyService.SetContext(controls.WithOpenUserContext(c)).UserCompanyDelete(req.Id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 授权列表
func (OpenUserCompanyController) UserCompanyAuthList(c *gin.Context) {

	var req entitys.OpenUserCompanyAuthListReq
	err := c.ShouldBindQuery(&req)

	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := OpenUserCompanyService.SetContext(controls.WithOpenUserContext(c)).UserCompanyAuthList(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}
