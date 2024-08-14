package router

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/community/apis"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotnatsjs"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(e *gin.Engine) {
	webApiPrefix := "/v1/platform/app"
	app := e.Group(webApiPrefix)
	app.Use(controls.AuthCheck)
	app.Use(controls.SetParams)
	app.Use(iotgin.AppLogger(iotnatsjs.GetJsClientPub()))

	//获取社区产品列表
	app.POST("/community/product/list", apis.CommunityProductcontroller.GetCommunityProductList)

	//获取社区产品详情
	app.GET("/community/product/detail", apis.CommunityProductcontroller.GetCommunityProductDetail)
}
