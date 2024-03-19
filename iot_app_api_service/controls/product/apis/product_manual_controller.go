package apis

import (
	"cloud_platform/iot_app_api_service/cached"
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/product/services"
	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_common/iotlogger"

	"github.com/gin-gonic/gin"
	goerrors "go-micro.dev/v4/errors"
)

type ProductManualController struct{}

var ProductManualControl ProductManualController

var productManual services.ProductManualService

// GetProductManual 获取产品说明书
// @Summary 获取产品说明书
// @Description
// @Tags APP
// @Accept application/json
// @Param productKey query string true "设备Id"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/product/manual/detail [get]
func (ProductManualController) GetProductManual(c *gin.Context) {
	productKey := c.Query("productKey")
	if len(productKey) == 0 {
		iotlogger.LogHelper.Helper.Error("product manual productKey is empty")
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrCloudRequestParamIsEmpty, nil)
		return
	}
	item, err := productManual.SetContext(controls.WithUserContext(c)).GetProductManual(productKey)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("get product manual error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, item)
}

// GetProductManualOld 获取产品说明书
func (ProductManualController) GetProductManualOld(c *gin.Context) {
	productKey := c.Query("productKey")
	if len(productKey) == 0 {
		iotlogger.LogHelper.Helper.Error("product manual productKey is empty")
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrCloudRequestParamIsEmpty, nil)
		return
	}
	item, err := productManual.SetContext(controls.WithUserContext(c)).GetProductManualOld(productKey)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("get product manual error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, item)
}
