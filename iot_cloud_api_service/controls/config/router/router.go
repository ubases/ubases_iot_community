package router

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/config/apis"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(e *gin.Engine) {
	webApiPrefix := "/v1/platform/web/config"
	admin := e.Group(webApiPrefix)
	admin.Use(controls.AuthCheck)

	admin.POST("/area/list", apis.Areacontroller.QueryList)
	admin.POST("/area/add", apis.Areacontroller.Add)
	admin.PUT("/area/edit", apis.Areacontroller.Edit)
	admin.POST("/area/edit", apis.Areacontroller.Edit)
	admin.DELETE("/area/delete", apis.Areacontroller.Delete)
	admin.POST("/area/delete", apis.Areacontroller.Delete)
	admin.GET("/area/detail/:id", apis.Areacontroller.QueryDetail)

	//获取区域的树型数据
	admin.GET("/area/treeData/:parentId/:showChild", apis.Areacontroller.GetAreas)

	//平台配置项
	admin.POST("/systemConfig/list", apis.ConfigPlatformcontroller.QueryList)
	admin.POST("/systemConfig/add", apis.ConfigPlatformcontroller.Add)
	admin.PUT("/systemConfig/edit", apis.ConfigPlatformcontroller.Edit)
	admin.POST("/systemConfig/edit", apis.ConfigPlatformcontroller.Edit)
	admin.DELETE("/systemConfig/delete", apis.ConfigPlatformcontroller.Delete)
	admin.POST("/systemConfig/delete", apis.ConfigPlatformcontroller.Delete)
	admin.GET("/systemConfig/detail/:id", apis.ConfigPlatformcontroller.QueryDetail)

}
