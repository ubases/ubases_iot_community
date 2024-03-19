package apis

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/document/entitys"
	apiservice "cloud_platform/iot_app_api_service/controls/document/services"
	"cloud_platform/iot_common/iotutil"

	"github.com/gin-gonic/gin"

	"cloud_platform/iot_common/iotgin"
)

var FeedBackcontroller UcUserFeedbackController

type UcUserFeedbackController struct{} //部门操作控制器

var feedbackServices = apiservice.UcFeedBackService{}

// QueryList 获取反馈列表
// @Summary 获取反馈列表
// @Description
// @Tags Document
// @Accept application/json
// @Param data body entitys.UcUserFeedbackQuery true "查询参数"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /feedback [post]
func (UcUserFeedbackController) QueryList(c *gin.Context) {
	var filter entitys.UcUserFeedbackQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	var (
		userId   = controls.GetUserId(c)
		appKey   = controls.GetAppKey(c)
		tenantId = controls.GetTenantId(c)
		lang     = controls.GetLang(c)
	)
	filter.UserId = iotutil.ToInt64(userId)
	res, _, err := feedbackServices.SetContext(controls.WithUserContext(c)).QueryUcFeedBackList(lang, appKey, tenantId, filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// QueryDetail 获取反馈详情
// @Summary QueryDetail
// @Description
// @Tags Document
// @Accept application/json
// @Param id path string true "文档编码"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /feedback/details/{id} [get]
func (UcUserFeedbackController) QueryDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	var (
		appKey   = controls.GetAppKey(c)
		tenantId = controls.GetTenantId(c)
		lang     = controls.GetLang(c)
	)
	res, err := feedbackServices.SetContext(controls.WithUserContext(c)).GetUcFeedBackDetail(lang, appKey, tenantId, id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// Edit 编辑反馈
// @Summary 辑反馈
// @Description
// @Tags Document
// @Accept application/json
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /feedback/edit [post]
func (UcUserFeedbackController) Edit(c *gin.Context) {
	var req entitys.UcUserFeedbackEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := feedbackServices.SetContext(controls.WithUserContext(c)).UpdateUcFeedBack(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// Add 新增反馈
// @Summary 新增反馈
// @Description
// @Tags Document
// @Accept application/json
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /feedback/add [post]
func (UcUserFeedbackController) Add(c *gin.Context) {
	var req entitys.UcUserFeedbackEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	userId := controls.GetUserId(c)
	req.UserId = iotutil.ToString(userId)
	req.MobileLang = c.GetHeader("lang")
	sysMap := controls.GetSystemInfo(c)
	req.MobileModel = sysMap.Model
	req.AppVersion = sysMap.Version
	req.MobileSystem = sysMap.Os
	id, err := feedbackServices.SetContext(controls.WithUserContext(c)).AddUcFeedBack(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// Delete 删除反馈
// @Summary 删除反馈
// @Description
// @Tags Document
// @Accept application/json
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /feedback/delete [post]
func (UcUserFeedbackController) Delete(c *gin.Context) {
	var req entitys.UcUserFeedbackFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = feedbackServices.SetContext(controls.WithUserContext(c)).DeleteUcFeedBack(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// GetFeedBackQuestionType 获取问题类型
// @Summary 获取问题类型
// @Description
// @Tags Document
// @Accept application/json
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /feedback/questionType [get]
func (UcUserFeedbackController) GetFeedBackQuestionType(c *gin.Context) {
	appKey := controls.GetAppKey(c)
	tenantId := controls.GetTenantId(c)
	lang := controls.GetLang(c)
	res, err := feedbackServices.SetContext(controls.WithUserContext(c)).GetFeedBackQuestionType(lang, appKey, tenantId, 0)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}
