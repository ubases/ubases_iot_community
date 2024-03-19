package router

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/product/apis"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(e *gin.Engine) {
	webApiPrefix := "/v1/platform/app"
	admin := e.Group(webApiPrefix)
	admin.Use(controls.AuthCheck)
	// 产品类型
	admin.GET("/product/productType", apis.ProductTypecontroller.GetProductTypeByApp)
	admin.GET("/product/productTypeV2", apis.ProductTypecontroller.GetProductTypeByAppV2)
	// 根据wifi名称查询产品列表
	admin.POST("/product/info", apis.Productcontroller.GetProductByWifiName)
	// 根据租户id查询此租户下的产品列表
	admin.POST("/product/list", apis.Productcontroller.GetOpmProductList)
	// 根据产品类型ID查询产品列表
	admin.POST("/product/getByProductTypeId", apis.Productcontroller.GetProductByProductTypeId)
	//查询配网引导详情 （GetDefaultNetworkGuides）
	admin.GET("/product/getNetworkGuide", apis.Productcontroller.QueryProductNetworkGuide)

	admin.GET("/controlPage/isUpdate", apis.Controlcontroller.CheckControlPageVersion)

	admin.GET("/material/click", apis.ProductMaterialControl.ClickProductMaterial)
	admin.GET("/material/detail", apis.ProductMaterialControl.GetProductMaterial)

	admin.GET("/product/manual/detail", apis.ProductManualControl.GetProductManualOld)
	admin.GET("/product/function/rules", apis.Productcontroller.QueryFunctionRules)

	//product/docs
	admin.GET("/product/docs", apis.Productcontroller.QueryDocuments)
}
