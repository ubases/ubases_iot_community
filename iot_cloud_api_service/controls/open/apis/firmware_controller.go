package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/open/services"
	"cloud_platform/iot_common/iotutil"

	"github.com/gin-gonic/gin"

	"cloud_platform/iot_common/iotgin"
)

var Firmwarecontroller OpmFirmwareController

type OpmFirmwareController struct{} //部门操作控制器

var firmwareServices = apiservice.OpmFirmwareService{}

func (OpmFirmwareController) QueryCustomFirmwareList(c *gin.Context) {
	page, err := iotutil.ToInt64AndErr(c.Query("page"))
	if err != nil {
		page = 1
	}
	limit, err := iotutil.ToInt64AndErr(c.Query("limit"))
	if err != nil {
		limit = 10
	}
	filter := entitys.OpmFirmwareQuery{
		Page:  page,
		Limit: limit,
		Query: &entitys.OpmFirmwareFilter{TenantId: controls.GetTenantId(c)},
	}
	res, total, err := firmwareServices.SetContext(controls.WithOpenUserContext(c)).QueryOpmFirmwareList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (OpmFirmwareController) QueryList(c *gin.Context) {
	var filter entitys.OpmFirmwareQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if filter.Query == nil {
		filter.Query = new(entitys.OpmFirmwareFilter)
	}
	filter.Query.TenantId = controls.GetTenantId(c)
	res, total, err := firmwareServices.SetContext(controls.WithOpenUserContext(c)).QueryOpmFirmwareList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (OpmFirmwareController) QueryDropDownList(c *gin.Context) {
	var filter entitys.OpmFirmwareQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, total, err := firmwareServices.SetContext(controls.WithOpenUserContext(c)).QueryOpmFirmwareList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (OpmFirmwareController) QueryDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	res, err := firmwareServices.SetContext(controls.WithOpenUserContext(c)).GetOpmFirmwareDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (OpmFirmwareController) Edit(c *gin.Context) {
	var req entitys.OpmFirmwareEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.TenantId = controls.GetTenantId(c)
	req.UpdatedBy = controls.GetUserId(c)
	id, err := firmwareServices.SetContext(controls.WithOpenUserContext(c)).UpdateOpmFirmware(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (OpmFirmwareController) Add(c *gin.Context) {
	var req entitys.OpmFirmwareEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.TenantId = controls.GetTenantId(c)
	id, err := firmwareServices.SetContext(controls.WithOpenUserContext(c)).AddOpmFirmware(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (OpmFirmwareController) Delete(c *gin.Context) {
	var req entitys.OpmFirmwareFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = firmwareServices.SetContext(controls.WithOpenUserContext(c)).DeleteOpmFirmware(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// SetStatus 设置状态
func (OpmFirmwareController) SetStatus(c *gin.Context) {
	var req entitys.OpmFirmwareFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == 0 || req.Status == nil {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	firmwareServices.Ctx = controls.WithOpenUserContext(c)
	err = firmwareServices.SetStatusOpmFirmware(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// SetStatus 设置状态
func (OpmFirmwareController) SetShelf(c *gin.Context, status int32) {
	var req entitys.OpmFirmwareFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	req.Status = &status // 已上架
	firmwareServices.Ctx = controls.WithOpenUserContext(c)
	err = firmwareServices.SetStatusOpmFirmware(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}
