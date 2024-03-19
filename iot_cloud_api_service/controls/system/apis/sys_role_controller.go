package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"

	"github.com/gin-gonic/gin"

	"cloud_platform/iot_cloud_api_service/controls/system/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/system/services"
	"cloud_platform/iot_common/iotgin"
)

var Rolecontroller RoleController

type RoleController struct{} //部门操作控制器

var roleServices = apiservice.SysRoleService{}

func (RoleController) QueryList(c *gin.Context) {
	var filter entitys.SysRoleQuery
	err := c.ShouldBindQuery(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, total, err := roleServices.QuerySysRoleList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (RoleController) QueryDetail(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	res, err := roleServices.GetSysRoleDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (RoleController) RoleDeptTreeSelect(c *gin.Context) {
	id := c.Query("roleId")
	if id == "" {
		iotgin.ResBadRequest(c, "roleId")
		return
	}
	res, err := roleServices.RoleDeptTreeSelect(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (RoleController) RoleDataScope(c *gin.Context) {
	var req entitys.SysRoleAddEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := roleServices.RoleDataScope(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 角色状态
func (RoleController) StatusSetRole(c *gin.Context) {
	var req entitys.SysRoleStatusReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := roleServices.RoleStatusUpdate(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)

}

// 角色新增的时候获取菜单列表
func (RoleController) AddRoleByMenuList(c *gin.Context) {
	res, err := authruleServices.QueryRoleAuthRuleList(entitys.SysAuthRuleQuery{
		Title:  "",
		Status: 1,
	})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (RoleController) Edit(c *gin.Context) {
	var req entitys.SysRoleAddEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := roleServices.UpdateSysRole(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (RoleController) Add(c *gin.Context) {
	var req entitys.SysRoleAddEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := roleServices.AddSysRole(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (RoleController) Delete(c *gin.Context) {
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
	err = roleServices.DeleteSysRole(controls.WithUserContext(c), req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}
