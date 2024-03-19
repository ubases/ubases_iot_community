package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/open/services"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"
	"errors"

	"github.com/gin-gonic/gin"
)

var OpenCompanyConnectcontroller OpenCompanyConnectController

var OpenCompanyConnectService apiservice.OpenCompanyConnectService

// 企业联系人
type OpenCompanyConnectController struct {
}

// 修改基础信息
func (OpenCompanyConnectController) AddConnect(c *gin.Context) {

	var req entitys.OpenCompanyConnect
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}

	res, err := OpenCompanyConnectService.SetContext(controls.WithOpenUserContext(c)).AddConnect(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 修改基础信息
func (OpenCompanyConnectController) UpdateConnect(c *gin.Context) {

	var req entitys.OpenCompanyConnect
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}

	res, err := OpenCompanyConnectService.SetContext(controls.WithOpenUserContext(c)).UpdateConnect(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 修改基础信息
func (OpenCompanyConnectController) DeleteConnect(c *gin.Context) {

	id := c.Query("id")
	if id == "" {
		iotgin.ResErrCli(c, errors.New("id error"))
		return
	}
	res, err := OpenCompanyConnectService.SetContext(controls.WithOpenUserContext(c)).DeleteConnect(iotutil.ToInt64(id))
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}
