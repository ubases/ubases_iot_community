package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/system/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/system/services"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"
	"errors"

	"github.com/gin-gonic/gin"
)

var SysAppHelpCentercontroller SysAppHelpCenterController

var helpCenterApp apiservice.SysAppHelpCenterService

type SysAppHelpCenterController struct {
} //用户操作控制器

// 创建帮助中心
func (SysAppHelpCenterController) CreateHelpCenter(c *gin.Context) {
	var req entitys.SysAppHelpCenterEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	err = helpCenterApp.SetContext(controls.WithOpenUserContext(c)).CreateHelpCenter(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, nil)
}

// 复制帮助中心
func (SysAppHelpCenterController) CopyHelpCenter(c *gin.Context) {
	var req entitys.SysAppHelpCenterEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	err = helpCenterApp.SetContext(controls.WithOpenUserContext(c)).CopyHelpCenter(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, nil)
}

// 更新帮助中心
func (SysAppHelpCenterController) UpdateHelpCenter(c *gin.Context) {
	var req entitys.SysAppHelpCenterEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	err = helpCenterApp.SetContext(controls.WithOpenUserContext(c)).UpdateHelpCenter(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, nil)
}

// 删除帮助中心
func (SysAppHelpCenterController) DeleteHelpCenter(c *gin.Context) {
	var req entitys.SysAppHelpCenterEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	err = helpCenterApp.SetContext(controls.WithOpenUserContext(c)).DeleteHelpCenter(iotutil.ToInt64(req.Id))
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, nil)
}

// 获取帮助中心详情
func (SysAppHelpCenterController) GetHelpCenter(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		iotgin.ResErrCli(c, errors.New("参数错误"))
		return
	}
	res, err := helpCenterApp.SetContext(controls.WithOpenUserContext(c)).GetHelpCenter(iotutil.ToInt64(id))
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 获取帮助中心列表
func (SysAppHelpCenterController) GetHelpCenterList(c *gin.Context) {
	var req entitys.SysAppHelpCenterQuery
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := helpCenterApp.SetContext(controls.WithOpenUserContext(c)).GetHelpCenterList(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, int64(len(res)), int(req.Page))
}

// 获取帮助中心列表(开放平台)
func (SysAppHelpCenterController) GetHelpCenterListForOpen(c *gin.Context) {
	var req entitys.SysAppHelpCenterQuery
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := helpCenterApp.SetContext(controls.WithOpenUserContext(c)).GetHelpCenterListForOpen(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, int64(len(res)), int(req.Page))
}
