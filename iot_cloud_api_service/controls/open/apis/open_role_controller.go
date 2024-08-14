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

var OpenRolecontroller OpenRoleController

var OpenRoleService apiservice.OpenRoleService

type OpenRoleController struct {
} //用户操作控制器

// 新增角色
func (OpenRoleController) AddRole(c *gin.Context) {
	var req entitys.OpenRoleAddReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := OpenRoleService.SetContext(controls.WithOpenUserContext(c)).AddRole(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 修改角色
func (OpenRoleController) EditRole(c *gin.Context) {
	var req entitys.OpenRoleEditReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := OpenRoleService.SetContext(controls.WithOpenUserContext(c)).EditRole(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 删除角色
func (OpenRoleController) DeleteRole(c *gin.Context) {

	var req entitys.OpenDeleteReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := OpenRoleService.SetContext(controls.WithOpenUserContext(c)).DeleteRole(req.Id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 删除角色
func (OpenRoleController) RoleList(c *gin.Context) {
	tenantId, isExists := c.Get("tenantId")
	if !isExists {
		iotgin.ResErrCli(c, errors.New("用户空间数据错误."))
		return
	}
	res, err := OpenRoleService.SetContext(controls.WithOpenUserContext(c)).RoleList(iotutil.ToString(tenantId))
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 删除角色
func (OpenRoleController) RoleDetail(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		iotgin.ResErrCli(c, errors.New("id参数错误"))
		return
	}
	res, err := OpenRoleService.SetContext(controls.WithOpenUserContext(c)).RoleDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (OpenRoleController) RoleSetUser(c *gin.Context) {
	var req entitys.OpenRoleSetUserReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.RoleId == "" || req.RoleId == "0" {
		iotgin.ResBadRequest(c, "角色Id不能为空！")
		return
	}
	if req.UserName == "" {
		iotgin.ResBadRequest(c, "用户名不能为空！")
		return
	}
	res, err := OpenRoleService.SetContext(controls.WithOpenUserContext(c)).RoleSetUser(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)

}
