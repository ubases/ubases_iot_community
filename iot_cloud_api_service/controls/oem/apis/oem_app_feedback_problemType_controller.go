package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/oem/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/oem/services"
	"cloud_platform/iot_common/iotgin"
	"github.com/gin-gonic/gin"
	"strconv"
)

var OemAppFeedbackProblemTypecontroller OemAppFeedbackProblemTypeController

var serviceFeedbackProblemType apiservice.OemAppFeedbackProblemTypeService

type OemAppFeedbackProblemTypeController struct {
}

// 反馈问题类型,新增
func (OemAppFeedbackProblemTypeController) Add(c *gin.Context) {
	var req entitys.OemFeedbackTypeSaveReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := serviceFeedbackProblemType.SetContext(controls.WithOpenUserContext(c)).OemAppFeedbackProblemTypeAdd(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 反馈问题类型,修改
func (OemAppFeedbackProblemTypeController) Update(c *gin.Context) {
	var req entitys.OemFeedbackTypeSaveReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	res, err := serviceFeedbackProblemType.SetContext(controls.WithOpenUserContext(c)).OemAppFeedbackProblemTypeUpdate(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 反馈问题类型,获取详细
func (OemAppFeedbackProblemTypeController) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if id == 0 || err != nil {
		iotgin.ResBadRequest(c, "id")
		return
	}

	res, err := serviceFeedbackProblemType.SetContext(controls.WithOpenUserContext(c)).OemAppFeedbackProblemTypeDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 反馈问题类型,获取列表
func (OemAppFeedbackProblemTypeController) List(c *gin.Context) {
	var req entitys.OemFeedbackTypeEntitys
	//err := c.ShouldBindQuery(&req)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, _, err := serviceFeedbackProblemType.SetContext(controls.WithOpenUserContext(c)).OemAppFeedbackProblemTypeList(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 删除反馈问题类型
func (OemAppFeedbackProblemTypeController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if id == 0 || err != nil {
		iotgin.ResBadRequest(c, "id")
		return
	}

	res, err := serviceFeedbackProblemType.SetContext(controls.WithOpenUserContext(c)).DeleteOemAppFeedbackProblemType(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}
