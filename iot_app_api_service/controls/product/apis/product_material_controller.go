package apis

import (
	"cloud_platform/iot_app_api_service/cached"
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/product/services"
	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"

	"github.com/gin-gonic/gin"
	goerrors "go-micro.dev/v4/errors"
)

type ProductMaterialController struct{}

var ProductMaterialControl ProductMaterialController

var productMaterial services.ProductMaterialService

// GetProductMaterial 获取耗材详情APP
// @Summary获取耗材详情APP
// @Description 获取耗材详情APP
// @Tags APP
// @Accept application/json
// @Param uids query []int false "耗材id"
// @Param productKey query string true "产品key"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @router /v1/platform/web/open/material/detail [get]
func (ProductMaterialController) GetProductMaterial(c *gin.Context) {
	id := c.QueryArray("uids")
	if len(id) == 0 {
		iotlogger.LogHelper.Helper.Error("product material id is empty")
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrCloudRequestParamIsEmpty, nil)
		return
	}
	productKey := c.Query("productKey")
	if len(id) == 0 {
		iotlogger.LogHelper.Helper.Error("product key is empty")
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrCloudRequestParamIsEmpty, nil)
		return
	}
	tenantId, _ := c.Get("tenantId")
	lang := c.GetHeader("lang")
	items, err := productMaterial.SetContext(controls.WithUserContext(c)).GetProductMaterial(id, tenantId.(string), productKey, lang)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("get product material error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	ioterrs.ResponsePage(c, cached.RedisStore, ioterrs.Success, items, int64(len(items)), 1)
}

// @Summary 统计耗材点击量
// @Description 统计耗材点击量
// @Tags APP
// @Accept application/json
// @Param id query string true "耗材id"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/material/click [get]
func (ProductMaterialController) ClickProductMaterial(c *gin.Context) {
	id := c.Query("id")
	if len(id) == 0 {
		iotlogger.LogHelper.Helper.Error("product material id is empty")
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrCloudRequestParamIsEmpty, nil)
		return
	}
	nID, err := iotutil.ToInt64AndErr(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	err = productMaterial.SetContext(controls.WithUserContext(c)).ClickProductMaterial(nID)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("click product material url error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, nil)
}
