package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_common/iotutil"

	"github.com/gin-gonic/gin"

	"cloud_platform/iot_cloud_api_service/controls/product/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/product/services"
	"cloud_platform/iot_common/iotgin"
)

var FirmwareVersioncontroller PmFirmwareVersionController

type PmFirmwareVersionController struct{} //部门操作控制器

var firmwareVersionServices = apiservice.PmFirmwareVersionService{}

func (PmFirmwareVersionController) QueryList(c *gin.Context) {
	var filter entitys.PmFirmwareVersionQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, total, err := firmwareVersionServices.QueryPmFirmwareVersionList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (PmFirmwareVersionController) QueryDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	res, err := firmwareVersionServices.GetPmFirmwareVersionDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (PmFirmwareVersionController) Edit(c *gin.Context) {
	var req entitys.PmFirmwareVersionEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.UpdatedBy = controls.GetUserId(c)
	id, err := firmwareVersionServices.UpdatePmFirmwareVersion(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (PmFirmwareVersionController) Add(c *gin.Context) {
	var req entitys.PmFirmwareVersionEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	//req.CreatedBy = controls.GetUserId(c)
	id, err := firmwareVersionServices.AddPmFirmwareVersion(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (PmFirmwareVersionController) Delete(c *gin.Context) {
	var req entitys.PmFirmwareVersionFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = firmwareVersionServices.DeletePmFirmwareVersion(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

func (PmFirmwareVersionController) SetStatus(c *gin.Context, status int32) {
	var req entitys.PmFirmwareVersionFilter
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
		req.Status = status
	}
	err = firmwareVersionServices.SetStatusPmFirmwareVersion(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

func (PmFirmwareVersionController) QueryChangeVersionList(c *gin.Context) {
	moduleId := c.Query("moduleId")
	if moduleId == "" {
		iotgin.ResBadRequest(c, "moduleId")
		return
	}
	moduleIdInt64, err := iotutil.ToInt64AndErr(moduleId)
	if err != nil {
		iotgin.ResBadRequest(c, "moduleId")
		return
	}
	var (
		page  int64 = 1
		limit int64 = 10
	)
	if val := c.Query("page"); val != "" {
		page, _ = iotutil.ToInt64AndErr(val)
		if err != nil {
			page = 0
		}
	}
	if val := c.Query("limit"); val != "" {
		limit, _ = iotutil.ToInt64AndErr(val)
		if err != nil {
			limit = 0
		}
	}
	res, total, err := firmwareVersionServices.QueryPmFirmwareChangeVersionList(moduleIdInt64, page, limit)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(page))
}
