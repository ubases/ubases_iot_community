package apis

import (
	"cloud_platform/iot_cloud_api_service/cached"
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	"cloud_platform/iot_cloud_api_service/controls/open/services"
	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_common/iotlogger"

	"github.com/gin-gonic/gin"
	goerrors "go-micro.dev/v4/errors"
)

type ProductManualController struct{}

var ProductManualControl ProductManualController

var productManual services.ProductManualService

// 创建产品说明书
func (ProductManualController) CreateProductManual(c *gin.Context) {
	var req entitys.OpmProductManualEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("bind create product manual param error: ", err)
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrShouldBindJSON, nil)
		return
	}
	if err := req.AddCheck(); err != nil {
		iotlogger.LogHelper.Helper.Error("check create product manual param error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	err = productManual.SetContext(controls.WithOpenUserContext(c)).CreateProductManual(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("create product manual error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, nil)
}

// 更新产品说明书
func (ProductManualController) UpdateProductManual(c *gin.Context) {
	var req entitys.OpmProductManualEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("bind update product manual param error: ", err)
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrShouldBindJSON, nil)
		return
	}
	if err := req.UpdateCheck(); err != nil {
		iotlogger.LogHelper.Helper.Error("check update product manual param error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	err = productManual.SetContext(controls.WithOpenUserContext(c)).UpdateProductManual(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("update product manual error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, nil)
}

// 删除产品说明书
func (ProductManualController) DeleteProductManual(c *gin.Context) {
	var req entitys.OpmProductManualEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("bind set product manual param error: ", err)
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrShouldBindJSON, nil)
		return
	}
	err = productManual.SetContext(controls.WithOpenUserContext(c)).DeleteProductManual(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("set product manual error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, nil)
}

// 获取产品说明书
func (ProductManualController) GetProductManual(c *gin.Context) {
	productKey := c.Query("productKey")
	if len(productKey) == 0 {
		iotlogger.LogHelper.Helper.Error("product manual productKey is empty")
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrCloudRequestParamIsEmpty, nil)
		return
	}
	item, err := productManual.SetContext(controls.WithOpenUserContext(c)).GetProductManual(productKey)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("get product manual error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, item)
}

// 获取产品说明书列表
func (ProductManualController) GetProductManualList(c *gin.Context) {
	var req entitys.OpmProductManualQuery
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("bind get product manual list param error: ", err)
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrShouldBindJSON, nil)
		return
	}
	items, total, err := productManual.SetContext(controls.WithOpenUserContext(c)).GetProductManualList(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("get product manual list error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	ioterrs.ResponsePage(c, cached.RedisStore, ioterrs.Success, items, total, int(req.Page))
}
