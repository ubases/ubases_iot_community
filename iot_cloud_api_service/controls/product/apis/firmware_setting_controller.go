package apis

import (
	"github.com/gin-gonic/gin"

	"cloud_platform/iot_cloud_api_service/controls/product/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/product/services"
	"cloud_platform/iot_common/iotgin"
)

var FirmwareSettingcontroller PmFirmwareSettingController

type PmFirmwareSettingController struct{} //部门操作控制器

var firmwareSettingServices = apiservice.PmFirmwareSettingService{}

func (PmFirmwareSettingController) QueryList(c *gin.Context) {
	var filter entitys.PmFirmwareSettingQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, total, err := firmwareSettingServices.QueryPmFirmwareSettingList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (PmFirmwareSettingController) QueryDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	res, err := firmwareSettingServices.GetPmFirmwareSettingDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (PmFirmwareSettingController) Edit(c *gin.Context) {
	var req entitys.PmFirmwareSettingEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := firmwareSettingServices.UpdatePmFirmwareSetting(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (PmFirmwareSettingController) Add(c *gin.Context) {
	var req entitys.PmFirmwareSettingEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := firmwareSettingServices.AddPmFirmwareSetting(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (PmFirmwareSettingController) Delete(c *gin.Context) {
	var req entitys.PmFirmwareSettingFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == "" {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = firmwareSettingServices.DeletePmFirmwareSetting(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}
