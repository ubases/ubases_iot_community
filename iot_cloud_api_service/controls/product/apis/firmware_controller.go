package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"

	"github.com/gin-gonic/gin"

	"cloud_platform/iot_cloud_api_service/controls/product/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/product/services"
	"cloud_platform/iot_common/iotgin"
)

var Firmwarecontroller PmFirmwareController

type PmFirmwareController struct{} //部门操作控制器

var firmwareServices = apiservice.PmFirmwareService{}

func (PmFirmwareController) QueryList(c *gin.Context) {
	var filter entitys.PmFirmwareQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, total, err := firmwareServices.QueryPmFirmwareList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (PmFirmwareController) QueryDropDownList(c *gin.Context) {
	var filter entitys.PmFirmwareQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, total, err := firmwareServices.QueryPmFirmwareList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (PmFirmwareController) QueryDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	res, err := firmwareServices.GetPmFirmwareDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (PmFirmwareController) Edit(c *gin.Context) {
	var req entitys.PmFirmwareEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.UpdatedBy = controls.GetUserId(c)
	id, err := firmwareServices.UpdatePmFirmware(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (PmFirmwareController) Add(c *gin.Context) {
	var req entitys.PmFirmwareEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	//req.CreatedBy = controls.GetUserId(c)
	id, err := firmwareServices.AddPmFirmware(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (PmFirmwareController) Delete(c *gin.Context) {
	var req entitys.PmFirmwareFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == "" {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = firmwareServices.DeletePmFirmware(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// SetStatus 设置状态
func (PmFirmwareController) SetStatus(c *gin.Context) {
	var req entitys.PmFirmwareFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == "" || req.Status == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = firmwareServices.SetStatusPmFirmware(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// SetStatus 设置状态
func (PmFirmwareController) SetShelf(c *gin.Context, status int32) {
	var req entitys.PmFirmwareFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == "" {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	req.Status = status // 已上架
	err = firmwareServices.SetStatusPmFirmware(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}
