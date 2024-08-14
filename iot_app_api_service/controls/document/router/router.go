package router

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/document/apis"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotnatsjs"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(e *gin.Engine) {
	webApiPrefix := "/v1/platform/app"
	admin := e.Group(webApiPrefix)
	admin.GET("/introduce/:code", apis.Introducecontroller.GetIntroduceDetailByCode)

	admin.Use(controls.AuthCheck)
	admin.Use(controls.SetParams)
	admin.Use(iotgin.AppLogger(iotnatsjs.GetJsClientPub()))
	// 产品类型
	admin.GET("/questionType", apis.QuestionTypecontroller.QueryList)
	admin.GET("/questiontop5", apis.Questioncontroller.QueryTop5)
	admin.GET("/question", apis.Questioncontroller.QueryList)
	admin.GET("/question/:id", apis.Questioncontroller.QueryDetail)

	//admin.GET("/introduce/:id", apis.Introducecontroller.GetIntroduceDetailByApp)
	//admin.GET("/introduce", apis.Introducecontroller.GetIntroduceByApp)

	admin.POST("/feedback/add", apis.FeedBackcontroller.Add)
	admin.POST("/feedback", apis.FeedBackcontroller.QueryList)
	admin.GET("/feedback/details/:id", apis.FeedBackcontroller.QueryDetail)
	admin.GET("/feedback/questionType", apis.FeedBackcontroller.GetFeedBackQuestionType)

	admin.GET("/productHelpDoc/list/:productKey", apis.ProductHelpDoccontroller.GetProductHelpDoc)
}
