package router

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/upgrade/apis"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(e *gin.Engine) {
	webApiPrefix := "/v1/platform/app"
	admin := e.Group(webApiPrefix)
	admin.Use(controls.AuthCheck)
	// 获取APP版本信息
	admin.POST("/appVersion/get", apis.AppUpgradecontroller.GetLatestApp)
	admin.GET("/functionConfig/autoUpgrade", apis.AppUpgradecontroller.GetFunctionConfigAutoUpgrade)
}
