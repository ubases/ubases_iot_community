package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/template/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/template/services"
	"github.com/gin-gonic/gin"

	"cloud_platform/iot_common/iotgin"
)

var DocumentTemplatecontroller TplDocumentTemplateController

type TplDocumentTemplateController struct{} //部门操作控制器

var documentTemplateServices = apiservice.TplDocumentTemplateService{}

func (TplDocumentTemplateController) QueryList(c *gin.Context) {
	var filter entitys.TplDocumentTemplateQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, total, err := documentTemplateServices.QueryTplDocumentTemplateList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (TplDocumentTemplateController) QueryDropDownList(c *gin.Context) {
	var filter entitys.TplDocumentTemplateQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, total, err := documentTemplateServices.QueryTplDocumentTemplateList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (TplDocumentTemplateController) QueryDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	res, err := documentTemplateServices.GetTplDocumentTemplateDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (TplDocumentTemplateController) Edit(c *gin.Context) {
	var req entitys.TplDocumentTemplateEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.UpdatedBy = controls.GetUserId(c)
	id, err := documentTemplateServices.UpdateTplDocumentTemplate(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (TplDocumentTemplateController) Add(c *gin.Context) {
	var req entitys.TplDocumentTemplateEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.CreatedBy = controls.GetUserId(c)
	id, err := documentTemplateServices.AddTplDocumentTemplate(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (TplDocumentTemplateController) Delete(c *gin.Context) {
	var req entitys.TplDocumentTemplateFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == "" {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = documentTemplateServices.DeleteTplDocumentTemplate(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// SetStatus 设置状态
func (TplDocumentTemplateController) SetStatus(c *gin.Context) {
	var req entitys.TplDocumentTemplateFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == "" || req.Status == nil {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = documentTemplateServices.SetStatusTplDocumentTemplate(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}
