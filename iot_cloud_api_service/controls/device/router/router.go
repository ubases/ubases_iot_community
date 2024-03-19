package router

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/device/apis"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(e *gin.Engine) {
	webApiPrefix := "/v1/platform/web/iot"
	admin := e.Group(webApiPrefix)
	admin.Use(controls.AuthCheck)

	// 设备三元组
	admin.POST("/device/list", apis.DeviceTriadcontroller.QueryList)
	admin.GET("/device/detail/:id", apis.DeviceTriadcontroller.QueryDetail)
	admin.POST("/device/add", apis.DeviceTriadcontroller.Add)
	admin.PUT("/device/edit", apis.DeviceTriadcontroller.Edit)
	admin.POST("/device/edit", apis.DeviceTriadcontroller.Edit)
	admin.DELETE("/device/delete", apis.DeviceTriadcontroller.Delete)
	admin.POST("/device/delete", apis.DeviceTriadcontroller.Delete)
	admin.POST("/deviceTriad/generator", apis.DeviceTriadcontroller.GeneratorDeviceTriad)
	admin.POST("/deviceTriad/import", apis.DeviceTriadcontroller.ImportDeviceTriad)

	//生成虚拟设备
	//个人账号最多生成10个，
	//企业账号30个，企业认证账号100个
	admin.POST("/device/generatorVirtualDevice", apis.DeviceTriadcontroller.CreateDeviceTriad)
	//添加虚拟设备
	admin.POST("/device/addVirtualDevice", apis.DeviceTriadcontroller.CreateVirtualDeviceTriad)
	//虚拟设备列表
	admin.POST("/device/virtualDeviceList", apis.DeviceTriadcontroller.QueryVirtualDeviceList)
	//删除虚拟设备
	admin.DELETE("/device/deleteVirtualDevice", apis.DeviceTriadcontroller.Delete)
	admin.POST("/device/deleteVirtualDevice", apis.DeviceTriadcontroller.Delete)
	//添加APP账号
	admin.POST("/device/addAppAccount", apis.DeviceTriadcontroller.AddAppAccount)
	admin.GET("/app/build/qrCodeUrl", apis.DeviceTriadcontroller.GetDefaultApp)

	// 设备管理
	admin.POST("/activeDevice/produce/list", apis.DeviceInfocontroller.QueryProduceList)
	admin.POST("/activeDevice/list", apis.DeviceInfocontroller.QueryList)
	admin.POST("/activeDevice/platformList", apis.DeviceInfocontroller.PlatformQueryList)
	admin.POST("/activeDevice/count", apis.DeviceInfocontroller.Count)
	admin.GET("/activeDevice/detail/:did", apis.DeviceInfocontroller.QueryDetail)
	admin.POST("/activeDevice/add", apis.DeviceInfocontroller.Add)
	admin.PUT("/activeDevice/edit", apis.DeviceInfocontroller.Edit)
	admin.POST("/activeDevice/edit", apis.DeviceInfocontroller.Edit)
	admin.DELETE("/activeDevice/delete", apis.DeviceInfocontroller.Delete)
	admin.POST("/activeDevice/delete", apis.DeviceInfocontroller.Delete)

	//平台导出方法
	admin.GET("/activeDevice/platformExport", func(c *gin.Context) {
		apis.DeviceInfocontroller.GetExport(c, 1)
	})
	admin.POST("/activeDevice/platformExport", func(c *gin.Context) {
		apis.DeviceInfocontroller.Export(c, 1)
	})
	admin.GET("/activeDevice/export", func(c *gin.Context) {
		apis.DeviceInfocontroller.GetExport(c, 0)
	})
	admin.POST("/activeDevice/export", func(c *gin.Context) {
		apis.DeviceInfocontroller.Export(c, 0)
	})
	admin.POST("/activeDevice/produce/export", func(c *gin.Context) {
		apis.DeviceInfocontroller.Export(c, 0)
	})
	//======================================
	admin.GET("/activeDevice/triadExport", apis.DeviceInfocontroller.GetExportTriad)
	admin.GET("/activeDevice/triadExportCount", apis.DeviceInfocontroller.GetExportTriadCount)

	// 设备日志
	admin.POST("/activeDevice/logList", apis.DeviceLogcontroller.QueryList)
	admin.POST("/activeDevice/logCount", apis.DeviceLogcontroller.QueryCount)
	admin.GET("/activeDevice/logExport", apis.DeviceLogcontroller.Export)
	admin.POST("/activeDevice/logExport", apis.DeviceLogcontroller.ExportPostMethod)

	//
}
