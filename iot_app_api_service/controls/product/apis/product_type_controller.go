package apis

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/product/entitys"
	"cloud_platform/iot_app_api_service/controls/product/services"
	"cloud_platform/iot_common/iotgin"

	"github.com/gin-gonic/gin"
)

var ProductTypecontroller ProductTypeController

type ProductTypeController struct {
} //用户操作控制器

var ProductTypeservices = services.ProductTypeService{}

// GetProductTypeByApp
// @Summary 获取产品类型列表
// @Description
// @Tags home
// @Accept application/json
// @Param data body entitys.AppQueryProductTypeForm true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Failure 400 object iotgin.ResponseModel 失败返回
// @Router /product/productType [get]
func (ProductTypeController) GetProductTypeByApp(c *gin.Context) {
	var filter entitys.AppQueryProductTypeForm
	err := c.ShouldBind(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, _, err := ProductTypeservices.SetContext(controls.WithUserContext(c)).GetProductTypeByApp(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)

}

// GetProductTypeByAppV2
// @Summary 获取产品类型列表V2
// @Description
// @Tags home
// @Accept application/json
// @Param data body entitys.AppQueryProductTypeForm true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /product/productTypeV2 [get]
func (ProductTypeController) GetProductTypeByAppV2(c *gin.Context) {
	var filter entitys.AppQueryProductTypeForm
	err := c.ShouldBind(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, _, err := ProductTypeservices.SetContext(controls.WithUserContext(c)).GetProductTree(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}
