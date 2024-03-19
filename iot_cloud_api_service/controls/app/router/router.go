package router

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/app/apis"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(e *gin.Engine) {
	webApiPrefix := "/v1/platform/web/open"
	appRouter := e.Group(webApiPrefix)

	appRouter.GET("/regionList", apis.Usercontroller.RegionList)
	appRouter.POST("/appCancelAccount", apis.Usercontroller.CancelAccount)
	appRouter.POST("/sendVerityCode", apis.Usercontroller.SendVerityCode)
	appRouter.GET("/getAppInfo", apis.Usercontroller.GetAppInfo)

	appRouter.Use(controls.AuthCheck)
	//反馈接口
	appRouter.POST("/feedback/list", apis.Feedbackcontroller.FeedbackList)
	appRouter.GET("/feedback/detail/:id", apis.Feedbackcontroller.FeedbackDetail)
	appRouter.POST("/feedback/reply", apis.Feedbackcontroller.FeedbackReplySubmit)

}
