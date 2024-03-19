package router

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/lang/apis"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(e *gin.Engine) {
	webApiPrefix := "/v1/platform/web/lang"
	admin := e.Group(webApiPrefix)
	//支持语言列表查询（字典获取，不需要在翻译类型表配置了）
	admin.GET("/langTypeList", apis.Translatecontroller.LangTypeList)
	admin.Use(controls.AuthCheck)
	//获取语言翻译
	admin.GET("/translation/get", apis.Translatecontroller.TranslateGet)
	//获取语言包括默认语言
	admin.GET("/translation/getV2", apis.Translatecontroller.TranslateGetV2)
	//提交语言翻译
	admin.POST("/translation/save", func(c *gin.Context) {
		apis.Translatecontroller.TranslateSave(1, c)
	})
	//语言列表
	admin.POST("/translation/list", apis.Translatecontroller.QueryList)
	//提交语言翻译
	admin.POST("/translation/customSave", func(c *gin.Context) {
		apis.Translatecontroller.TranslateSave(2, c)
	})

	//新增语言翻译
	//删除语言翻译
	//修改语言翻译
	//查询语言翻译列表
	//获取语言翻译详情
	//获取语言翻译模板
	//获取指定语言翻译包

	//多语言包详情
	admin.GET("/appResources/detail", apis.Resourcescontroller.ResourcePackageDetail)

	admin.GET("/appResourcesPackage/get/:id", apis.ResourcesPackagecontroller.Get)
	admin.POST("/appResourcesPackage/list", apis.ResourcesPackagecontroller.List)
	admin.POST("/appResourcesPackage/add", apis.ResourcesPackagecontroller.Add)
	admin.PUT("/appResourcesPackage/edit", apis.ResourcesPackagecontroller.Update)
	admin.POST("/appResourcesPackage/edit", apis.ResourcesPackagecontroller.Update)
	admin.DELETE("/appResourcesPackage/delete/:id", apis.ResourcesPackagecontroller.Delete)
	admin.POST("/appResourcesPackage/delete/:id", apis.ResourcesPackagecontroller.Delete)

	//多语言资源上传
	admin.POST("/appResources/import", apis.Resourcescontroller.ResourceImport)
	//多语言资源下载（兼容post和get）
	admin.GET("/appResources/download", apis.Resourcescontroller.GetExport)
	admin.POST("/appResources/download", apis.Resourcescontroller.GetExport)
	admin.GET("/appResources/templateDownload", apis.Resourcescontroller.ResourceImportExcelTemplate)

	//支持的翻译类型目前配置在字典中
	//translate_type

	open := e.Group("/v1/platform/open/lang")

	open.Use(controls.AuthCheck)
	//自定义资源
	open.GET("/customResources/detail", apis.Resourcescontroller.CustomResourceDetail)
	//自定义多语言资源上传
	open.POST("/customResources/import", apis.Resourcescontroller.CustomResourceImport)
	//自定义多语言资源下载（兼容post和get）
	open.GET("/customResources/download", apis.Resourcescontroller.CustomResourceExport)
	open.POST("/customResources/download", apis.Resourcescontroller.CustomResourceExport)
	open.GET("/customResources/jsonData", apis.Resourcescontroller.CustomResourceJsonData)

	//自定义资源编辑
	open.POST("/customResources/save", apis.Resourcescontroller.CustomResourcesSave)
	open.GET("/customResources/get", apis.Resourcescontroller.CustomResourcesGet)
	open.GET("/customResources/list", apis.Resourcescontroller.QueryCustomResourceV2)
}
