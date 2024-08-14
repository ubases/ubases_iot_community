package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"

	"github.com/gin-gonic/gin"
)

// 禁止移出到router目录
// 后续所有开发均方apis目录
func RegisterRouter(e *gin.Engine) {
	webApiPrefix := "/v1/platform/web/data"
	admin := e.Group(webApiPrefix)
	admin.POST("/open/overview/failLoglist", devicedatacontroller.getFailLogList)
	admin.Use(controls.AuthCheck)
	admin.GET("/pm/overview/accumulate", Overviewcontroller.getAccumulateData)
	admin.GET("/pm/overview/today", Overviewcontroller.getTodayData)
	admin.GET("/pm/overview/city", Overviewcontroller.getCityDeviceData)

	admin.GET("/open/overview/accumulate", Openoverviewcontroller.getAccumulateData)
	admin.GET("/open/overview/today", Openoverviewcontroller.getTodayData)
	admin.GET("/open/overview/appUser", Openappusercontroller.getUserAppStatistics)
	admin.GET("/open/overview/deviceActive", Opendevicecontroller.getActiveStatistics)
	admin.GET("/open/overview/deviceFault", Opendevicecontroller.getFaultStatistics)

	admin.GET("/open/overview/appUser/export", Openappusercontroller.ExportUserAppStatistics)
	admin.GET("/open/overview/deviceActive/export", Opendevicecontroller.ExportActiveStatistics)
	admin.GET("/open/overview/deviceFault/export", Opendevicecontroller.ExportFaultStatistics)

	admin.POST("/pm/deviceFault/list", devicedatacontroller.getFaultList)

	admin.GET("/pm/device/total", devicedatacontroller.GetDeviceTotalStatistics)
	admin.POST("/pm/developer/list", DeveloperdataController.getDeveloperList)
	admin.GET("/pm/developer/detail", DeveloperdataController.getDeveloperDetail)
	admin.GET("/pm/developer/total", DeveloperdataController.getDeveloperTotal)

	admin.POST("/pm/app/list", AppdataController.getAppList)
	admin.GET("/pm/app/detail", AppdataController.getAppDetail)
}

//https://www.bejson.com/transfor/json2go/
//https://www.json.cn/#
