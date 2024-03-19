package apis

import (
	"cloud_platform/iot_cloud_api_service/controls/template/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/template/services"

	"github.com/gin-gonic/gin"

	"cloud_platform/iot_common/iotgin"
)

var MessageTemplatecontroller MpMessageTemplateController

type MpMessageTemplateController struct{} //部门操作控制器

var messageTemplateServices = apiservice.MpMessageTemplateService{}

func (MpMessageTemplateController) QueryList(c *gin.Context) {
	var filter entitys.MpMessageTemplateQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, total, err := messageTemplateServices.QueryMpMessageTemplateList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (MpMessageTemplateController) QueryDropDownList(c *gin.Context) {
	var filter entitys.MpMessageTemplateQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, total, err := messageTemplateServices.QueryMpMessageTemplateList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (MpMessageTemplateController) QueryDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	res, err := messageTemplateServices.GetMpMessageTemplateDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (MpMessageTemplateController) Edit(c *gin.Context) {
	var req entitys.MpMessageTemplateEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := messageTemplateServices.UpdateMpMessageTemplate(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (MpMessageTemplateController) Add(c *gin.Context) {
	var req entitys.MpMessageTemplateEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := messageTemplateServices.AddMpMessageTemplate(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (MpMessageTemplateController) Delete(c *gin.Context) {
	var req entitys.MpMessageTemplateFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = messageTemplateServices.DeleteMpMessageTemplate(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}
