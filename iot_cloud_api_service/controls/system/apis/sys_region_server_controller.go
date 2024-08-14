package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/system/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/system/services"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"

	"github.com/gin-gonic/gin"
)

var SysRegionServercontroller SysRegionServerController

type SysRegionServerController struct{} //部门操作控制器

var sysRegionServerServices = apiservice.SysRegionServerService{}

// @Summary 查询APP服务器管理信息
// @Description
// @Tags APP
// @Accept application/json
// @Param data body string true "请求参数结构体 {"code": "", "name":"", ...}"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/web/system/cloudPlatform/list [get]
func (SysRegionServerController) QueryList(c *gin.Context) {
	var filter entitys.SysRegionServerQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if filter.Query == nil {
		filter.Query = &entitys.SysRegionServerFilter{}
	}
	res, total, err := sysRegionServerServices.SetContext(controls.WithUserContext(c)).QuerySysRegionServerList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

// @Summary 查询APP服务器管理详细信息
// @Description
// @Tags APP
// @Accept application/json
// @Param id path string true "平台Id"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/web/system/cloudPlatform/detail/{id} [get]
func (SysRegionServerController) QueryDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	res, err := sysRegionServerServices.SetContext(controls.WithUserContext(c)).GetSysRegionServerDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// @Summary 编辑APP服务器管理信息
// @Description
// @Tags APP
// @Accept application/json
// @Param data body string true "请求参数结构体 {"code": "", "name":"", ...}"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/web/system/cloudPlatform/edit [post]
func (SysRegionServerController) Edit(c *gin.Context) {
	var req entitys.SysRegionServerEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := sysRegionServerServices.SetContext(controls.WithUserContext(c)).UpdateSysRegionServer(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// @Summary 新增APP服务器管理信息
// @Description
// @Tags APP
// @Accept application/json
// @Param data body string true "请求参数结构体 {"code": "", "name":"", ...}"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/web/system/cloudPlatform/add [post]
func (SysRegionServerController) Add(c *gin.Context) {
	var req entitys.SysRegionServerEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.CreatedBy = controls.GetUserId(c)
	id, err := sysRegionServerServices.SetContext(controls.WithUserContext(c)).AddSysRegionServer(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// @Summary 删除APP服务器管理信息
// @Description
// @Tags APP
// @Accept application/json
// @Param data body string true "请求参数结构体 {"code": "", "name":"", ...}"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/web/system/cloudPlatform/delete/{id} [post]
func (SysRegionServerController) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	idInt, err := iotutil.ToInt64AndErr(id)
	if err != nil {
		iotgin.ResBadRequest(c, "id format")
		return
	}
	err = sysRegionServerServices.SetContext(controls.WithUserContext(c)).DeleteSysRegionServer(entitys.SysRegionServerFilter{Id: idInt})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// @Summary 设置状态
// @Description
// @Tags APP
// @Accept application/json
// @Param data body string true "请求参数结构体 {"id": "", "status":"", ...}"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/web/system/cloudPlatform/setStatus [post]
func (SysRegionServerController) SetStatus(c *gin.Context) {
	var req entitys.SysRegionServerFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == 0 || req.Enabled == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = sysRegionServerServices.SetContext(controls.WithUserContext(c)).SetStatusSysRegionServer(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}
