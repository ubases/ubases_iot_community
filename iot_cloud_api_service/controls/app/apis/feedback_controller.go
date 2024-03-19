package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/app/entitys"
	"cloud_platform/iot_cloud_api_service/controls/app/services"
	"cloud_platform/iot_common/iotgin"

	"github.com/gin-gonic/gin"
)

var Feedbackcontroller FeedbackController

type FeedbackController struct{}

var feedbackService = services.FeedbackService{}

// FeedbackList 反馈列表
// @Summary 反馈列表
// @Description
// @Tags 反馈
// @Accept application/json
// @Param data body string true "请求参数结构体 json UcUserFeedbackQuery"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/feedback/list [post]
func (s FeedbackController) FeedbackList(c *gin.Context) {
	var filter entitys.UcUserFeedbackQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	lang := controls.GetLang(c)
	tenantId := controls.GetTenantId(c)
	res, total, err := feedbackService.SetContext(controls.WithUserContext(c)).QueryUcFeedbackList(filter, tenantId, lang)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

// FeedbackDetail 反馈详情
// @Summary 反馈列表
// @Description
// @Tags 反馈
// @Accept application/json
// @Param id path string true "反馈Id"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/feedback/detail/{id} [get]
func (s FeedbackController) FeedbackDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	tenantId := controls.GetTenantId(c)
	res, err := feedbackService.SetContext(controls.WithUserContext(c)).GetUcFeedbackDetail(tenantId, id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// FeedbackReplySubmit 提交反馈
// @Summary 反馈列表
// @Description
// @Tags 反馈
// @Accept application/json
// @Param data body entitys.FeedbackReplySubmit true "请求参数结构体"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/feedback/reply [post]
func (s FeedbackController) FeedbackReplySubmit(c *gin.Context) {
	var req entitys.FeedbackReplySubmit
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.Operator = c.GetInt64("userId")
	err = feedbackService.SetContext(controls.WithUserContext(c)).FeedbackReplySubmit(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}
