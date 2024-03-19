package router

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/community/apis"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(e *gin.Engine) {
	webApiPrefix := "/v1/platform/app"
	app := e.Group(webApiPrefix)
	app.Use(controls.AuthCheck)

	//获取社区产品列表
	app.POST("/community/product/list", apis.CommunityProductcontroller.GetCommunityProductList)

	//获取社区产品详情
	app.GET("/community/product/detail", apis.CommunityProductcontroller.GetCommunityProductDetail)
}
