package router

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/product/apis"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(e *gin.Engine) {
	webApiPrefix := "/v1/platform/web"
	admin := e.Group(webApiPrefix)
	admin.Use(controls.AuthCheck)

	// 产品类型
	admin.GET("/pm/product/get", apis.Productcontroller.Get)
	//检查产品名称是否重复
	admin.GET("/pm/product/nameExists", apis.Productcontroller.CheckExists)
	admin.POST("/pm/product/get", apis.Productcontroller.GetProductList)
	//
	admin.POST("/pm/product/save", apis.Productcontroller.Create)
	admin.POST("/pm/product/update", apis.Productcontroller.Update)
	admin.POST("/pm/product/delete", apis.Productcontroller.Delete)
	admin.POST("/pm/product/status", apis.Productcontroller.Status)
	admin.GET("/pm/thingModel/getStandard", apis.Productcontroller.GetStandardThingModelDetail)
	admin.GET("/pm/networkGuide/GetDefaultNetworkGuides", apis.Productcontroller.GetDefaultNetworkGuides)
	admin.GET("/pm/product/resetProductThingModels", apis.ProductThingsModelcontroller.ResetThingsModel)

	//产品类型按照步骤编辑 (考虑是否拆分接口保存）
	//基础信息
	//admin.POST("/pm/product/createProduct", apis.Productcontroller.CreateProduct)
	//指令控制
	//场景条件
	//控制面板
	//模组选择
	//关联固件
	//配网方式

	// 产品分类
	admin.GET("/pm/productType/get", apis.ProductTypecontroller.Get)
	admin.POST("/pm/productType/get", apis.ProductTypecontroller.GetProductTypeList)
	admin.POST("/pm/productType/save", apis.ProductTypecontroller.Create)
	admin.POST("/pm/productType/update", apis.ProductTypecontroller.Update)
	admin.POST("/pm/productType/delete", apis.ProductTypecontroller.Delete)
	admin.POST("/pm/productType/getTypeAndProducts", apis.ProductTypecontroller.GetTypeAndProductList)
	admin.GET("/pm/productType/getModelTemplate", apis.ProductTypecontroller.GetModelTemplate)

	//芯片模组管理
	admin.POST("/product/module/list", apis.Modulecontroller.QueryList)
	admin.POST("/product/module/selectList", apis.Modulecontroller.QueryList)
	admin.POST("/product/module/add", apis.Modulecontroller.Add)
	admin.PUT("/product/module/edit", apis.Modulecontroller.Edit)
	admin.POST("/product/module/edit", apis.Modulecontroller.Edit)
	admin.DELETE("/product/module/delete", apis.Modulecontroller.Delete)
	admin.POST("/product/module/delete", apis.Modulecontroller.Delete)
	admin.GET("/product/module/detail/:id", apis.Modulecontroller.QueryDetail)
	admin.POST("/product/module/setStatus", apis.Modulecontroller.SetStatus)

	//固件管理
	admin.POST("/product/firmware/list", apis.Firmwarecontroller.QueryList)
	admin.POST("/product/firmware/add", apis.Firmwarecontroller.Add)
	admin.PUT("/product/firmware/edit", apis.Firmwarecontroller.Edit)
	admin.POST("/product/firmware/edit", apis.Firmwarecontroller.Edit)
	admin.POST("/product/firmware/setStatus", apis.Firmwarecontroller.SetStatus)
	admin.DELETE("/product/firmware/delete", apis.Firmwarecontroller.Delete)
	admin.POST("/product/firmware/delete", apis.Firmwarecontroller.Delete)
	admin.GET("/product/firmware/detail/:id", apis.Firmwarecontroller.QueryDetail)

	//固件版本管理
	admin.POST("/product/firmwareVersion/list", apis.FirmwareVersioncontroller.QueryList)
	admin.POST("/product/firmwareVersion/add", apis.FirmwareVersioncontroller.Add)
	admin.PUT("/product/firmwareVersion/edit", apis.FirmwareVersioncontroller.Edit)
	admin.POST("/product/firmwareVersion/edit", apis.FirmwareVersioncontroller.Edit)
	admin.PUT("/product/firmwareVersion/setStatus", func(context *gin.Context) {
		apis.FirmwareVersioncontroller.SetStatus(context, -1)
	})
	admin.POST("/product/firmwareVersion/setStatus", func(context *gin.Context) {
		apis.FirmwareVersioncontroller.SetStatus(context, -1)
	})
	admin.PUT("/product/firmwareVersion/onShelf", func(context *gin.Context) {
		apis.FirmwareVersioncontroller.SetStatus(context, 1)
	})
	admin.POST("/product/firmwareVersion/onShelf", func(context *gin.Context) {
		apis.FirmwareVersioncontroller.SetStatus(context, 1)
	})
	admin.PUT("/product/firmwareVersion/unShelf", func(context *gin.Context) {
		apis.FirmwareVersioncontroller.SetStatus(context, 2)
	})
	admin.POST("/product/firmwareVersion/unShelf", func(context *gin.Context) {
		apis.FirmwareVersioncontroller.SetStatus(context, 2)
	})
	admin.DELETE("/product/firmwareVersion/delete", apis.FirmwareVersioncontroller.Delete)
	admin.POST("/product/firmwareVersion/delete", apis.FirmwareVersioncontroller.Delete)
	admin.GET("/product/firmwareVersion/detail/:id", apis.FirmwareVersioncontroller.QueryDetail)

	//固件设置管理
	admin.POST("/product/firmwareSetting/list", apis.FirmwareSettingcontroller.QueryList)
	admin.POST("/product/firmwareSetting/add", apis.FirmwareSettingcontroller.Add)
	admin.PUT("/product/firmwareSetting/edit", apis.FirmwareSettingcontroller.Edit)
	admin.POST("/product/firmwareSetting/edit", apis.FirmwareSettingcontroller.Edit)
	admin.DELETE("/product/firmwareSetting/delete", apis.FirmwareSettingcontroller.Delete)
	admin.POST("/product/firmwareSetting/delete", apis.FirmwareSettingcontroller.Delete)
	admin.GET("/product/firmwareSetting/detail/:id", apis.FirmwareSettingcontroller.QueryDetail)

	//控制面板管理
	//admin.GET("/pm/controlPanel/get", apis.ControlPanelcontroller.Get)
	//admin.POST("/pm/controlPanel/get", apis.ControlPanelcontroller.GetControlPanelList)
	//admin.POST("/pm/controlPanel/save", apis.ControlPanelcontroller.Create)
	//admin.POST("/pm/controlPanel/update", apis.ControlPanelcontroller.Update)
	//admin.POST("/pm/controlPanel/delete", apis.ControlPanelcontroller.Delete)

	admin.POST("/pm/controlPanel/add", apis.ControlpanelsController.Add)
	admin.PUT("/pm/controlPanel/edit", apis.ControlpanelsController.Update)
	admin.POST("/pm/controlPanel/edit", apis.ControlpanelsController.Update)
	admin.DELETE("/pm/controlPanel/delete/:id", apis.ControlpanelsController.Delete)
	admin.POST("/pm/controlPanel/delete/:id", apis.ControlpanelsController.Delete)
	admin.GET("/pm/controlPanel/detail/:id", apis.ControlpanelsController.Get)
	admin.POST("/pm/controlPanel/list", apis.ControlpanelsController.GetList)
	admin.PUT("/pm/controlPanel/setStatus", apis.ControlpanelsController.SetStatus)
	admin.POST("/pm/controlPanel/setStatus", apis.ControlpanelsController.SetStatus)
	//多语言资源下载（兼容post和get）
	//admin.GET("/pm/controlPanel/langDownload", apis2.Resourcescontroller.GetExport)
	//admin.POST("/pm/controlPanel/langDownload", apis2.Resourcescontroller.GetExport)

	//配网引导管理
	admin.GET("/pm/networkGuide/get", apis.NetworkGuidecontroller.Get)
	admin.POST("/pm/networkGuide/get", apis.NetworkGuidecontroller.GetNetworkGuideList)
	admin.POST("/pm/networkGuide/save", apis.NetworkGuidecontroller.Create)
	admin.POST("/pm/networkGuide/update", apis.NetworkGuidecontroller.Update)
	admin.POST("/pm/networkGuide/delete", apis.NetworkGuidecontroller.Delete)

	//配网引导步骤管理
	admin.GET("/pm/networkGuideStep/get", apis.NetworkGuideStepcontroller.Get)
	admin.POST("/pm/networkGuideStep/get", apis.NetworkGuideStepcontroller.GetNetworkGuideStepList)
	admin.POST("/pm/networkGuideStep/save", apis.NetworkGuideStepcontroller.Create)
	admin.POST("/pm/networkGuideStep/update", apis.NetworkGuideStepcontroller.Update)
	admin.POST("/pm/networkGuideStep/delete", apis.NetworkGuideStepcontroller.Delete)
}
