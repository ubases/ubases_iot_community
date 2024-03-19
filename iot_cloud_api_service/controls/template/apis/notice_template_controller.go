package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/template/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/template/services"
	"github.com/gin-gonic/gin"

	"cloud_platform/iot_common/iotgin"
)

var NoticeTemplatecontroller MsNoticeTemplateController

type MsNoticeTemplateController struct{} //部门操作控制器

var noticeTemplateServices = apiservice.MsNoticeTemplateService{}

func (MsNoticeTemplateController) QueryList(c *gin.Context) {
	var filter entitys.MsNoticeTemplateQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, total, err := noticeTemplateServices.QueryMsNoticeTemplateList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (MsNoticeTemplateController) QueryDropDownList(c *gin.Context) {
	var filter entitys.MsNoticeTemplateQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, total, err := noticeTemplateServices.QueryMsNoticeTemplateList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (MsNoticeTemplateController) QueryDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	res, err := noticeTemplateServices.GetMsNoticeTemplateDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (MsNoticeTemplateController) Edit(c *gin.Context) {
	var req entitys.MsNoticeTemplateEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.UpdatedBy = controls.GetUserId(c)
	id, err := noticeTemplateServices.UpdateMsNoticeTemplate(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (MsNoticeTemplateController) Add(c *gin.Context) {
	var req entitys.MsNoticeTemplateEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.CreatedBy = controls.GetUserId(c)
	id, err := noticeTemplateServices.AddMsNoticeTemplate(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (MsNoticeTemplateController) Delete(c *gin.Context) {
	var req entitys.MsNoticeTemplateFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == "" {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = noticeTemplateServices.DeleteMsNoticeTemplate(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}
