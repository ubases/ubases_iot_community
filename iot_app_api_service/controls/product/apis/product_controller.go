package apis

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/product/entitys"
	"cloud_platform/iot_app_api_service/controls/product/services"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"

	"github.com/gin-gonic/gin"
)

var Productcontroller ProductController

type ProductController struct {
} //用户操作控制器

var Productservices = services.ProductService{}

// @Summary 根据wifiName获取产品列表
// @Description
// @Tags home
// @Accept application/json
// @Param data body entitys.AppQueryProductForm true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Failure 400 object iotgin.ResponseModel 失败返回
// @Router /product/info [post]
func (ProductController) GetProductByWifiName(c *gin.Context) {
	var filter entitys.AppQueryProductForm
	err := c.ShouldBind(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	//if filter.WifiFlags == nil || len(filter.WifiFlags) == 0 {
	//	iotgin.ResBadRequest(c, "wifiFlags不能为空")
	//	return
	//}
	userId := controls.GetUserId(c)
	filter.TenantId = c.Request.Header.Get("tenantId")
	filter.AppKey = c.Request.Header.Get("appKey")
	language := c.GetHeader("lang")

	res, _, err := Productservices.SetContext(controls.WithUserContext(c)).GetProductByApp(filter, iotutil.ToString(userId), language)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// @Summary 通过产品类型获取产品信息
// @Description
// @Tags home
// @Accept application/json
// @Param data body entitys.AppQueryProductForm true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/product/getByProductTypeId [post]
func (ProductController) GetProductByProductTypeId(c *gin.Context) {
	var filter entitys.AppQueryProductForm
	err := c.ShouldBind(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	productTypeIdInt, _ := iotutil.ToInt64AndErr(filter.ProductTypeId)
	if productTypeIdInt == 0 {
		iotgin.ResBadRequest(c, "productTypeId不能为空")
		return
	}
	userId, _ := c.Get("UserId")
	//tenantId := c.Request.Header.Get("tenantId")
	res, total, err := Productservices.SetContext(controls.WithUserContext(c)).GetProductByApp(filter, iotutil.ToString(userId), "")
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, int64(total), int(filter.Page))

}

// @Summary 获取产品配网引导信息
// @Description
// @Tags home
// @Accept application/json
// @Param productId query string true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/product/getNetworkGuide [post]
func (ProductController) QueryProductNetworkGuide(c *gin.Context) {
	id := c.Query("productId")
	if id == "" {
		iotgin.ResBadRequest(c, "productId")
		return
	}
	//networkGuideType := c.Query("type")
	//if networkGuideType == "" {
	//	iotgin.ResBadRequest(c, "type")
	//	return
	//}
	res, err := Productservices.SetContext(controls.WithUserContext(c)).QueryProductNetworkGuide(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// @Summary 根据产品类型Id获取产品列表
// @Description
// @Tags home
// @Accept application/json
// @Param data body entitys.AppQueryProductForm true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Failure 400 object iotgin.ResponseModel 失败返回
// @Router /product/getByProductBaseId [get]
func (ProductController) GetProductByProductBaseId(c *gin.Context) {
	var filter entitys.AppQueryProductForm
	err := c.ShouldBind(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if filter.BaseProductId == 0 {
		iotgin.ResBadRequest(c, "baseProductId不能为空")
		return
	}
	userId := controls.GetUserId(c)
	res, total, err := Productservices.SetContext(controls.WithUserContext(c)).GetProductByApp(filter, iotutil.ToString(userId), "")
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, int64(total), int(filter.Page))

}

// @Summary 我的产品列表
// @Description
// @Tags home
// @Accept application/json
// @Param data body entitys.OpmProductQuery true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/web/open/product/list [post]
func (ProductController) GetOpmProductList(c *gin.Context) {
	var filter entitys.OpmProductQuery
	err := c.ShouldBind(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}

	res, total, err := Productservices.SetContext(controls.WithUserContext(c)).GetOpmProductList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

// QueryFunctionRules 产品功能规则
// @Summary 产品功能规则
// @Description
// @Tags APP
// @Accept application/json
// @Param productKey query string true "设备Id"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/product/function/rules [get]
func (ProductController) QueryFunctionRules(c *gin.Context) {
	productKey := c.Query("productKey")
	if productKey == "" {
		iotgin.ResBadRequest(c, "productKey")
		return
	}
	var dataOriginInt int32 = 0
	res, err := Productservices.SetContext(controls.WithUserContext(c)).GetFunctionRules(productKey, dataOriginInt, nil)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// QueryDocuments 获取产品文档清单
// @Summary 获取产品文档清单
// @Description
// @Tags APP
// @Accept application/json
// @Param productKey query string true "设备Id"
// @Param docCodes query string true "文档编码参数"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/product/docs [get]
func (ProductController) QueryDocuments(c *gin.Context) {
	productKey := c.Query("productKey")
	if productKey == "" {
		iotgin.ResBadRequest(c, "productKey")
		return
	}
	docCodes := c.Query("docCodes")
	if docCodes == "" {
		docCodes = "product_manual_doc"
	}
	res, err := Productservices.SetContext(controls.WithUserContext(c)).QueryOpmDocumentsList(productKey, docCodes)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}

	iotgin.ResSuccess(c, res)
}
