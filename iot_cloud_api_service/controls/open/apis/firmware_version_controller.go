package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/open/services"
	"cloud_platform/iot_common/iotutil"

	"github.com/gin-gonic/gin"

	"cloud_platform/iot_common/iotgin"
)

var FirmwareVersioncontroller OpmFirmwareVersionController

type OpmFirmwareVersionController struct{} //部门操作控制器

var firmwareVersionServices = apiservice.OpmFirmwareVersionService{}

func (OpmFirmwareVersionController) QueryList(c *gin.Context) {
	var filter entitys.OpmFirmwareVersionQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if filter.Query == nil {
		filter.Query = new(entitys.OpmFirmwareVersionFilter)
	}
	filter.Query.TenantId = controls.GetTenantId(c)
	var status int32 = -1
	if filter.Query.Status != nil {
		status = *filter.Query.Status
	}
	filter.Query.Status = &status
	res, total, err := firmwareVersionServices.SetContext(controls.WithOpenUserContext(c)).QueryOpmFirmwareVersionList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (OpmFirmwareVersionController) QueryEnableList(c *gin.Context) {
	var filter entitys.OpmFirmwareVersionQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if filter.Query == nil {
		filter.Query = new(entitys.OpmFirmwareVersionFilter)
	}
	filter.Query.TenantId = controls.GetTenantId(c)
	var status int32 = -1
	filter.Query.Status = &status
	res, total, err := firmwareVersionServices.SetContext(controls.WithOpenUserContext(c)).QueryOpmFirmwareVersionList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (OpmFirmwareVersionController) QueryCustomEnableList(c *gin.Context) {
	firmwareId := c.Query("firmwareId")
	if firmwareId == "" {
		iotgin.ResBadRequest(c, "firmwareId")
		return
	}
	//productId := c.Query("productId")
	//if productId == "" {
	//	iotgin.ResBadRequest(c, "productId")
	//	return
	//}
	firmwareIdInt, err := iotutil.ToInt64AndErr(firmwareId)
	if err != nil {
		iotgin.ResBadRequest(c, "firmwareId error ")
		return
	}
	filter := entitys.OpmFirmwareVersionQuery{
		Query: &entitys.OpmFirmwareVersionFilter{
			FirmwareId: firmwareIdInt,
			TenantId:   controls.GetTenantId(c),
		},
	}
	var status int32 = 1
	filter.Query.Status = &status
	res, total, err := firmwareVersionServices.SetContext(controls.WithOpenUserContext(c)).QueryOpmFirmwareVersionList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (OpmFirmwareVersionController) QueryDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	res, err := firmwareVersionServices.SetContext(controls.WithOpenUserContext(c)).GetOpmFirmwareVersionDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (OpmFirmwareVersionController) Edit(c *gin.Context) {
	var req entitys.OpmFirmwareVersionEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.TenantId = controls.GetTenantId(c)
	req.UpdatedBy = controls.GetUserId(c)
	id, err := firmwareVersionServices.SetContext(controls.WithOpenUserContext(c)).UpdateOpmFirmwareVersion(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (OpmFirmwareVersionController) Add(c *gin.Context) {
	var req entitys.OpmFirmwareVersionEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.TenantId = controls.GetTenantId(c)
	id, err := firmwareVersionServices.SetContext(controls.WithOpenUserContext(c)).AddOpmFirmwareVersion(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (OpmFirmwareVersionController) Delete(c *gin.Context) {
	var req entitys.OpmFirmwareVersionFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = firmwareVersionServices.SetContext(controls.WithOpenUserContext(c)).DeleteOpmFirmwareVersion(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

func (OpmFirmwareVersionController) SetStatus(c *gin.Context, status int32) {
	var req entitys.OpmFirmwareVersionFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	if status != -1 {
		req.Status = &status
	}
	err = firmwareVersionServices.SetContext(controls.WithOpenUserContext(c)).SetStatusOpmFirmwareVersion(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

func (OpmFirmwareVersionController) OnShelf(c *gin.Context) {
	var req entitys.OpmFirmwareVersionFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = firmwareVersionServices.SetContext(controls.WithOpenUserContext(c)).OnShelfOpmFirmwareVersion(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}
