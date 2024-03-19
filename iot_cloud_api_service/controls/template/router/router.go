package router

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/template/apis"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(e *gin.Engine) {
	webApiPrefix := "/v1/platform/web"
	admin := e.Group(webApiPrefix)
	admin.Use(controls.AuthCheck)
	//测试用例模板管理 （测试用例与产品相关，其模板存放到产品库）
	admin.POST("/template/testCaseTpl/list", apis.TestcaseTemplatecontroller.QueryList)
	admin.POST("/template/testCaseTpl/add", apis.TestcaseTemplatecontroller.Add)
	admin.PUT("/template/testCaseTpl/edit", apis.TestcaseTemplatecontroller.Edit)
	admin.POST("/template/testCaseTpl/edit", apis.TestcaseTemplatecontroller.Edit)
	admin.POST("/template/testCaseTpl/setStatus", apis.TestcaseTemplatecontroller.SetStatus)
	admin.DELETE("/template/testCaseTpl/delete", apis.TestcaseTemplatecontroller.Delete)
	admin.POST("/template/testCaseTpl/delete", apis.TestcaseTemplatecontroller.Delete)
	admin.GET("/template/testCaseTpl/detail/:id", apis.TestcaseTemplatecontroller.QueryDetail)

	//TODO 下载测试用例模板 productTypeId
	admin.GET("/template/testCaseTpl/download", apis.TestcaseTemplatecontroller.GetTestReportTemplate)

	//通知消息模板
	admin.POST("/template/noticeTpl/list", apis.NoticeTemplatecontroller.QueryList)
	admin.POST("/template/noticeTpl/add", apis.NoticeTemplatecontroller.Add)
	admin.PUT("/template/noticeTpl/edit", apis.NoticeTemplatecontroller.Edit)
	admin.POST("/template/noticeTpl/edit", apis.NoticeTemplatecontroller.Edit)
	admin.DELETE("/template/noticeTpl/delete", apis.NoticeTemplatecontroller.Delete)
	admin.POST("/template/noticeTpl/delete", apis.NoticeTemplatecontroller.Delete)
	admin.GET("/template/noticeTpl/detail/:id", apis.NoticeTemplatecontroller.QueryDetail)

	//文档模板
	admin.POST("/template/documentTpl/list", apis.DocumentTemplatecontroller.QueryList)
	admin.POST("/template/documentTpl/add", apis.DocumentTemplatecontroller.Add)
	admin.PUT("/template/documentTpl/edit", apis.DocumentTemplatecontroller.Edit)
	admin.POST("/template/documentTpl/edit", apis.DocumentTemplatecontroller.Edit)
	admin.DELETE("/template/documentTpl/delete", apis.DocumentTemplatecontroller.Delete)
	admin.POST("/template/documentTpl/delete", apis.DocumentTemplatecontroller.Delete)
	admin.POST("/template/documentTpl/setStatus", apis.DocumentTemplatecontroller.SetStatus)
	admin.GET("/template/documentTpl/detail/:id", apis.DocumentTemplatecontroller.QueryDetail)

	admin.GET("/template/getTplFile/:code", apis.TestcaseTemplatecontroller.GetTplFile)

	//App消息模板
	admin.POST("/template/messageTpl/list", apis.MessageTemplatecontroller.QueryList)
	admin.POST("/template/messageTpl/add", apis.MessageTemplatecontroller.Add)
	admin.PUT("/template/messageTpl/edit", apis.MessageTemplatecontroller.Edit)
	admin.POST("/template/messageTpl/edit", apis.MessageTemplatecontroller.Edit)
	admin.DELETE("/template/messageTpl/delete", apis.MessageTemplatecontroller.Delete)
	admin.POST("/template/messageTpl/delete", apis.MessageTemplatecontroller.Delete)
	admin.GET("/template/messageTpl/detail/:id", apis.MessageTemplatecontroller.QueryDetail)

}
