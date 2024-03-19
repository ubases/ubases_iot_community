package apis

import (
	"cloud_platform/iot_app_api_service/controls"

	"github.com/gin-gonic/gin"

	apiservice "cloud_platform/iot_app_api_service/controls/document/services"
	"cloud_platform/iot_common/iotgin"
)

var QuestionTypecontroller QuestionTypeController

type QuestionTypeController struct{} //部门操作控制器

var questionTypeServices = apiservice.QuestionTypeService{}

// QueryList 获取问题类型列表
// @Summary 获取问题类型列表
// @Description
// @Tags Document
// @Accept application/json
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /questionType [get]
func (QuestionTypeController) QueryList(c *gin.Context) {
	res, err := questionTypeServices.SetContext(controls.WithUserContext(c)).QueryQuestionTypeList()
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}
