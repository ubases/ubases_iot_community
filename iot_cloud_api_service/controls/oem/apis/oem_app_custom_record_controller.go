package apis

import (
	"cloud_platform/iot_cloud_api_service/cached"
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/oem/entitys"
	"cloud_platform/iot_cloud_api_service/controls/oem/services"
	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotlogger"

	"github.com/gin-gonic/gin"
	goerrors "go-micro.dev/v4/errors"
)

type OemAppCustomRecordController struct{}

var OemAppCustomRecordControl OemAppCustomRecordController

var oemAppCustomRecord services.OemAppCustomRecordService

// 创建自定义app
func (OemAppCustomRecordController) CreateOemAppCustomRecord(c *gin.Context) {
	var req entitys.OemAppCustomRecordEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("bind create custom app param error: ", err)
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrShouldBindJSON, nil)
		return
	}
	if err := req.AddCheck(); err != nil {
		iotlogger.LogHelper.Helper.Error("check create custom app param error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	err = oemAppCustomRecord.SetContext(controls.WithOpenUserContext(c)).CreateOemAppCustomRecord(&req, true)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("create custom app error: ", err)
		iotgin.ResErrCli(c, err)
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, nil)
}

// 更新自定义app
func (OemAppCustomRecordController) UpdateOemAppCustomRecord(c *gin.Context) {
	var req entitys.OemAppCustomRecordEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("bind update custom app param error: ", err)
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrShouldBindJSON, nil)
		return
	}
	if err := req.UpdateCheck(); err != nil {
		iotlogger.LogHelper.Helper.Error("check update custom app param error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	err = oemAppCustomRecord.SetContext(controls.WithOpenUserContext(c)).UpdateOemAppCustomRecord(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("update custom app error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, nil)
}

// 更新自定义app
func (OemAppCustomRecordController) SetOemAppCustomRecord(c *gin.Context) {
	var req entitys.OemAppCustomRecordEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("bind update custom app param error: ", err)
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrShouldBindJSON, nil)
		return
	}
	if err := req.UpdateCheck(); err != nil {
		iotlogger.LogHelper.Helper.Error("check update custom app param error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	err = oemAppCustomRecord.SetContext(controls.WithOpenUserContext(c)).SetOemAppCustomRecord(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("update custom app error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, nil)
}

// 更新上架记录
func (OemAppCustomRecordController) SetLaunchMarkets(c *gin.Context) {
	var req entitys.OemAppCustomRecordEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("bind update custom app param error: ", err)
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrShouldBindJSON, nil)
		return
	}
	if req.LaunchMarkets == nil {
		iotgin.ResBadRequest(c, "LaunchMarkets")
		return
	}
	err = oemAppCustomRecord.SetContext(controls.WithOpenUserContext(c)).SetLaunchMarkets(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("update custom app error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, nil)
}

// 更新上架记录
func (OemAppCustomRecordController) SetRemark(c *gin.Context) {
	var req entitys.OemAppCustomRecordEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("bind update custom app param error: ", err)
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrShouldBindJSON, nil)
		return
	}
	if req.Description == "" {
		iotgin.ResBadRequest(c, "description")
		return
	}
	err = oemAppCustomRecord.SetContext(controls.WithOpenUserContext(c)).SetRemark(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("update custom app error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, nil)
}

// 更新提醒方式
func (OemAppCustomRecordController) SetRemindMode(c *gin.Context) {
	var req entitys.OemAppCustomRecordEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("bind update custom app param error: ", err)
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrShouldBindJSON, nil)
		return
	}
	//if req.RemindDescEn == "" {
	//	iotgin.ResBadRequest(c, "descEn")
	//	return
	//}
	//if req.RemindDesc == "" {
	//	iotgin.ResBadRequest(c, "desc")
	//	return
	//}
	if req.RemindMode == 0 {
		iotgin.ResBadRequest(c, "remindMode")
		return
	}
	err = oemAppCustomRecord.SetContext(controls.WithOpenUserContext(c)).SetRemindMode(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("update custom app error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, nil)
}

// 删除自定义app
func (OemAppCustomRecordController) DeleteOemAppCustomRecord(c *gin.Context) {
	var req entitys.OemAppCustomRecordEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("bind set custom app param error: ", err)
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrShouldBindJSON, nil)
		return
	}
	err = oemAppCustomRecord.SetContext(controls.WithOpenUserContext(c)).DeleteOemAppCustomRecord(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("set custom app error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, nil)
}

// 获取自定义app
func (OemAppCustomRecordController) GetOemAppCustomRecord(c *gin.Context) {
	id := c.Query("id")
	if len(id) == 0 {
		iotlogger.LogHelper.Helper.Error("custom app id is empty")
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrCloudRequestParamIsEmpty, nil)
		return
	}
	item, err := oemAppCustomRecord.SetContext(controls.WithOpenUserContext(c)).GetOemAppCustomRecord(id)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("get custom app error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, item)
}

// 获取自定义app列表
func (OemAppCustomRecordController) GetOemAppCustomRecordList(c *gin.Context) {
	var req entitys.OemAppCustomRecordQuery
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("bind get custom app list param error: ", err)
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrShouldBindJSON, nil)
		return
	}
	if req.Query == nil {
		req.Query = &entitys.OemAppCustomRecordFilter{}
	}
	//req.Query.Status = 1
	items, total, err := oemAppCustomRecord.SetContext(controls.WithOpenUserContext(c)).GetOemAppCustomRecordList(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("get custom app list error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	ioterrs.ResponsePage(c, cached.RedisStore, ioterrs.Success, items, total, int(req.Page))
}

// 打开自定义app二维码链接内容
func (OemAppCustomRecordController) CustomPackageQrCode(c *gin.Context) {
	var req entitys.OemAppCommonReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.Writer.WriteString(err.Error())
		return
	}
	ret, errBr := oemAppCustomRecord.SetContext(controls.WithUserContext(c)).CustomPackageQrCode(req)
	if errBr != nil {
		c.Writer.WriteString(errBr.Error())
		return
	}
	c.Writer.WriteString(ret)
}
