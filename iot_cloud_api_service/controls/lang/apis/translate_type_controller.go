package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/lang/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/lang/services"

	"github.com/gin-gonic/gin"

	"cloud_platform/iot_common/iotgin"
)

var TranslateTypecontroller LangTranslateTypeController

type LangTranslateTypeController struct{} //部门操作控制器

var translateTypeServices = apiservice.LangTranslateTypeService{}

func (LangTranslateTypeController) QueryList(c *gin.Context) {
	var filter entitys.LangTranslateTypeQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if filter.Query == nil {
		filter.Query = new(entitys.LangTranslateTypeFilter)
	}
	res, total, err := translateTypeServices.SetContext(controls.WithOpenUserContext(c)).QueryLangTranslateTypeList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (LangTranslateTypeController) QueryDropDownList(c *gin.Context) {
	var filter entitys.LangTranslateTypeQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, total, err := translateTypeServices.SetContext(controls.WithOpenUserContext(c)).QueryLangTranslateTypeList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (LangTranslateTypeController) QueryDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	res, err := translateTypeServices.SetContext(controls.WithOpenUserContext(c)).GetLangTranslateTypeDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (LangTranslateTypeController) Edit(c *gin.Context) {
	var req entitys.LangTranslateTypeEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.UpdatedBy = controls.GetUserId(c)
	id, err := translateTypeServices.SetContext(controls.WithOpenUserContext(c)).UpdateLangTranslateType(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (LangTranslateTypeController) Add(c *gin.Context) {
	var req entitys.LangTranslateTypeEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := translateTypeServices.SetContext(controls.WithOpenUserContext(c)).AddLangTranslateType(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (LangTranslateTypeController) Delete(c *gin.Context) {
	var req entitys.LangTranslateTypeFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = translateTypeServices.SetContext(controls.WithOpenUserContext(c)).DeleteLangTranslateType(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}
