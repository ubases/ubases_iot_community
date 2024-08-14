package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"github.com/gin-gonic/gin"

	"cloud_platform/iot_cloud_api_service/controls/product/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/product/services"
	"cloud_platform/iot_common/iotgin"
)

var Modulecontroller PmModuleController

type PmModuleController struct{} //部门操作控制器

var moduleServices = apiservice.PmModuleService{}

func (PmModuleController) QueryList(c *gin.Context) {
	var filter entitys.PmModuleQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, total, err := moduleServices.QueryPmModuleList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (PmModuleController) QueryDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	res, err := moduleServices.GetPmModuleDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (PmModuleController) Edit(c *gin.Context) {
	var req entitys.PmModuleEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.UpdatedBy = controls.GetUserId(c)
	id, err := moduleServices.UpdatePmModule(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 模组芯片绑定固件版本
func (PmModuleController) SelectFirmwareVersions(c *gin.Context) {
	var req entitys.PmModuleEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.UpdatedBy = controls.GetUserId(c)
	id, err := moduleServices.UpdatePartPmModule(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (PmModuleController) Add(c *gin.Context) {
	var req entitys.PmModuleEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	//req.CreatedBy = controls.GetUserId(c)
	id, err := moduleServices.AddPmModule(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (PmModuleController) Delete(c *gin.Context) {
	var req entitys.PmModuleFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == "" {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = moduleServices.DeletePmModule(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// SetStatus 设置状态
func (PmModuleController) SetStatus(c *gin.Context) {
	var req entitys.PmModuleFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == "" || req.Status == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = moduleServices.SetStatusPmModule(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}
