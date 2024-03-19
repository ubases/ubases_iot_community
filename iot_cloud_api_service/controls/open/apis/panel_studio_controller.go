package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/open/services"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"
	"github.com/gin-gonic/gin"
)

var Panelcontroller OpmPanelController

type OpmPanelController struct{} //部门操作控制器

var panelServices = apiservice.OpmPanelService{}

func (OpmPanelController) QueryList(c *gin.Context) {
	var filter entitys.OpmPanelQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if filter.Query == nil {
		filter.Query = &entitys.OpmPanelFilter{}
	}
	filter.Query.TenantId = controls.GetTenantId(c)

	res, total, err := panelServices.SetContext(controls.WithUserContext(c)).QueryOpmPanelList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (OpmPanelController) QueryDropDownList(c *gin.Context) {
	var filter entitys.OpmPanelQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, total, err := panelServices.SetContext(controls.WithUserContext(c)).QueryOpmPanelList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (OpmPanelController) QueryDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	res, err := panelServices.SetContext(controls.WithUserContext(c)).GetOpmPanelDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (OpmPanelController) Edit(c *gin.Context) {
	var req entitys.OpmPanelEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.UpdatedBy = controls.GetUserId(c)
	req.TenantId = controls.GetTenantId(c)
	req.UpdatedBy = controls.GetUserId(c)
	id, err := panelServices.SetContext(controls.WithUserContext(c)).UpdateOpmPanel(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (OpmPanelController) EditStudio(c *gin.Context) {
	var req entitys.OpmPanelEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.UpdatedBy = controls.GetUserId(c)
	req.TenantId = controls.GetTenantId(c)
	req.UpdatedBy = controls.GetUserId(c)
	id, err := panelServices.SetContext(controls.WithUserContext(c)).UpdateOpmPanelStudio(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (OpmPanelController) Add(c *gin.Context) {
	var req entitys.OpmPanelEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	//req.CreatedBy = controls.GetUserId(c)
	req.TenantId = controls.GetTenantId(c)
	req.CreatedBy = controls.GetUserId(c)
	id, err := panelServices.SetContext(controls.WithUserContext(c)).AddOpmPanel(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (OpmPanelController) Delete(c *gin.Context) {
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
	err = panelServices.SetContext(controls.WithUserContext(c)).DeleteOpmPanel(entitys.OpmPanelFilter{Id: idInt})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// SetStatus 设置状态
func (OpmPanelController) SetStatus(c *gin.Context) {
	var req entitys.OpmPanelFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == 0 || req.Status == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = panelServices.SetContext(controls.WithUserContext(c)).SetStatusOpmPanel(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}
