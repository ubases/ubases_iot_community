package apis

import (
	"cloud_platform/iot_app_api_service/controls/document/entitys"
	"cloud_platform/iot_common/iotutil"

	"github.com/gin-gonic/gin"

	apiservice "cloud_platform/iot_app_api_service/controls/document/services"
	"cloud_platform/iot_common/iotgin"
)

var Questioncontroller QuestionController

type QuestionController struct{} //部门操作控制器

var questionServices = apiservice.QuestionService{}

// QueryTop5 获取常见问题Top5
// @Summary 获取常见问题Top5
// @Description
// @Tags Document
// @Accept application/json
// @Param tenantId header string true "租户Id"
// @Param lang header string true "所属语言"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /questiontop5 [get]
func (QuestionController) QueryTop5(c *gin.Context) {
	appKey := c.Request.Header.Get("appKey")
	lang := c.GetHeader("lang")
	res, err := questionServices.QueryQuestionTop5(appKey, lang)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	var resultList = []*entitys.QuestionItem{}
	for _, item := range res {
		resultList = append(resultList, &entitys.QuestionItem{
			Id:    iotutil.ToInt64(item.SetingId),
			Title: item.Title,
		})
	}
	iotgin.ResSuccess(c, resultList)
}

// QueryList 获取常见问题列表
// @Summary 获取常见问题列表
// @Description
// @Tags Document
// @Accept application/json
// @Param page query string true "分页页码"
// @Param limit query string true "分页记录条数"
// @Param title query string true "标题"
// @Param content query string true "内容"
// @Param isTop query string true "是否置顶"
// @Param typeId query string true "类型Id"
// @Param productKey query string true "产品Key"
// @Param tenantId header string true "租户Id"
// @Param lang header string true "所属语言"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /question [get]
func (QuestionController) QueryList(c *gin.Context) {
	var filter = entitys.QuestionQuery{
		Page:     iotutil.ToInt64(c.DefaultQuery("page", "0")),
		Limit:    iotutil.ToInt64(c.DefaultQuery("limit", "0")),
		Title:    c.Query("title"),
		Abstract: c.Query("content"),
		IsTop:    c.Query("isTop"),
		TypeId:   c.Query("typeId"),
		Model:    c.Query("productKey"),
	}
	appKey := c.Request.Header.Get("appKey")
	tenantId := c.Request.Header.Get("tenantId")
	lang := c.GetHeader("lang")
	res, err := questionServices.QueryQuestionList(filter, appKey, lang, tenantId)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	var resultList = []*entitys.QuestionItem{}
	for _, item := range res {
		resultList = append(resultList, &entitys.QuestionItem{
			Id:    iotutil.ToInt64(item.SetingId),
			Title: item.Title,
		})
	}
	iotgin.ResPageSuccess(c, resultList, 0, int(filter.Page))
}

// QueryDetail 获取常见问题详情
// @Summary 常见问题
// @Description
// @Tags Document
// @Accept application/json
// @Param id path string true "问题Id"
// @Param type query string true "类型"
// @Param tenantId header string true "租户Id"
// @Param lang header string true "所属语言"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /question/{id} [get]
func (QuestionController) QueryDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	queryType, _ := iotutil.ToInt32ErrNew(c.Query("type"))
	if queryType == 0 {
		iotgin.ResBadRequest(c, "type")
		return
	}
	lang := c.GetHeader("lang")
	appKey := c.GetHeader("appKey")
	res, err := questionServices.GetQuestionDetail(id, queryType, appKey, lang)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}
