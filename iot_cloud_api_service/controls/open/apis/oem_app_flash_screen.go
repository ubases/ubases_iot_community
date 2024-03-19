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

type OemAppFlashScreenController struct{}

var OemAppFlashScreenControl OemAppFlashScreenController

var flashScreen services.OemAppFlashScreenService

// 创建闪屏
func (OemAppFlashScreenController) CreateFlashScreen(c *gin.Context) {
	var req entitys.OemAppFlashScreenEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("bind create flash screen param error: ", err)
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrShouldBindJSON, nil)
		return
	}
	if err := req.AddCheck(); err != nil {
		iotlogger.LogHelper.Helper.Error("check create flash screen param error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	err = flashScreen.SetContext(controls.WithOpenUserContext(c)).CreateFlashScreen(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("create flash screen error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, nil)
}

// 更新闪屏
func (OemAppFlashScreenController) UpdateFlashScreen(c *gin.Context) {
	var req entitys.OemAppFlashScreenEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("bind update flash screen param error: ", err)
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrShouldBindJSON, nil)
		return
	}
	if err := req.UpdateCheck(); err != nil {
		iotlogger.LogHelper.Helper.Error("check update flash screen param error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	err = flashScreen.SetContext(controls.WithOpenUserContext(c)).UpdateFlashScreen(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("update flash screen error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, nil)
}

// 设置闪屏
func (OemAppFlashScreenController) SetFlashScreen(c *gin.Context) {
	var req entitys.OemAppFlashScreenEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("bind set flash screen param error: ", err)
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrShouldBindJSON, nil)
		return
	}
	err = flashScreen.SetContext(controls.WithOpenUserContext(c)).SetFlashScreen(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("set flash screen error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, nil)
}

// 获取闪屏
func (OemAppFlashScreenController) GetFlashScreen(c *gin.Context) {
	id := c.Query("id")
	if len(id) == 0 {
		iotlogger.LogHelper.Helper.Error("flash screen id is empty")
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrCloudRequestParamIsEmpty, nil)
		return
	}
	item, err := flashScreen.SetContext(controls.WithOpenUserContext(c)).GetFlashScreen(id)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("get flash screen error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, item)
}

// 获取闪屏列表
func (OemAppFlashScreenController) GetFlashScreenList(c *gin.Context) {
	var req entitys.OemAppFlashScreenQuery
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("bind get flash screen list param error: ", err)
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrShouldBindJSON, nil)
		return
	}
	items, total, err := flashScreen.SetContext(controls.WithOpenUserContext(c)).GetFlashScreenList(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("get flash screen list error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	ioterrs.ResponsePage(c, cached.RedisStore, ioterrs.Success, items, total, int(req.Page))
}
