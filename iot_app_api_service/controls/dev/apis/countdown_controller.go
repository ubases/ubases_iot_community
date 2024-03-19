package apis

import (
	"cloud_platform/iot_app_api_service/cached"
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/dev/entitys"
	"cloud_platform/iot_app_api_service/controls/dev/services"
	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"time"

	"github.com/gin-gonic/gin"
	goerrors "go-micro.dev/v4/errors"
)

var Countdowncontroller CountdownController

type CountdownController struct {
} //倒计时操作控制器

var countdownService = services.AppCountdownService{}

// AddCountdown 创建倒计时
// @Summary 创建倒计时
// @Description
// @Tags 设备
// @Accept application/json
// @Param data body entitys.IotDeviceCountdownEntitys true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /dev/countdown/add [post]
func (CountdownController) AddCountdown(c *gin.Context) {
	req := entitys.IotDeviceCountdownEntitys{}
	if err := c.BindJSON(&req); err != nil {
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrAppRequestParam, nil)
		return
	}
	devId := req.DeviceId
	if devId == "" {
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrAppDeviceIdEmpty, nil)
		return
	}
	if req.FuncKey == "" {
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrAppFuncKeyEmpty, nil)
		return
	}

	userId := controls.GetUserId(c)
	req.UserId = iotutil.ToInt64(userId)
	req.CreatedBy = iotutil.ToInt64(userId)
	req.Timezone = controls.GetTimezone(c)
	req.RegionServerId = controls.GetRegionInt(c)
	err := countdownService.AddCountdown(req)
	if err != nil {
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		iotlogger.LogHelper.Helper.Error("add count down error: ", err.Error())
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, nil)
}

// CountdownInfo 倒计时信息查询
// @Summary 倒计时信息查询
// @Description
// @Tags 设备
// @Accept application/json
// @Param id path string true "定时Id"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /dev/countdown/info/{id} [get]
func (CountdownController) CountdownInfo(c *gin.Context) {
	id := c.Param("devId")
	if id == "" {
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrAppDeviceIdEmpty, nil)
		return
	}
	data, err := countdownService.CountdownInfo(id)
	if err != nil {
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		iotlogger.LogHelper.Helper.Error("get count down info error: ", err.Error())
		return
	}
	var surplusTime int64
	if data != nil {
		surplusTime = data.ExecutionTime.Unix() - time.Now().Unix()
		if surplusTime < 0 {
			surplusTime = 0
		} else {
			surplusTime = surplusTime * 1000
		}
		data.SurplusTime = surplusTime
	}

	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, data)
}

// EnabledCountdown 开启倒计时
// @Summary 开启倒计时
// @Description
// @Tags 设备
// @Accept application/json
// @Param id path string true "定时Id"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /dev/countdown/open/{id} [post]
func (CountdownController) EnabledCountdown(c *gin.Context) {
	id := c.Param("devId")
	if id == "" {
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrAppDeviceIdEmpty, nil)
		return
	}
	data, err := countdownService.CountdownInfo(id)
	if err != nil {
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		iotlogger.LogHelper.Helper.Error("get count down info error: ", err.Error())
		return
	}
	if data == nil {
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrAppCountDownNotExist, nil)
		return
	}
	err = countdownService.EnabledCountdown(data.Id)
	if err != nil {
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		iotlogger.LogHelper.Helper.Error("enable count down error: ", err.Error())
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, nil)
}

// DisabledCountdown 关闭倒计时
// @Summary 关闭倒计时
// @Description
// @Tags 设备
// @Accept application/json
// @Param id path string true "定时Id"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /dev/countdown/close/{id} [post]
func (CountdownController) DisabledCountdown(c *gin.Context) {
	id := c.Param("devId")
	if id == "" {
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrAppDeviceIdEmpty, nil)
		return
	}
	data, err := countdownService.CountdownInfo(id)
	if err != nil {
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		iotlogger.LogHelper.Helper.Error("get count down info error: ", err.Error())
		return
	}
	if data == nil {
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrAppCountDownNotExist, nil)
		return
	}
	err = countdownService.DisabledCountdown(data.Id)
	if err != nil {
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		iotlogger.LogHelper.Helper.Error("enable count down error: ", err.Error())
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, nil)
}

// RemoveCountdown 移除倒计时
// @Summary 移除倒计时
// @Description
// @Tags 设备
// @Accept application/json
// @Param id path string true "定时Id"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /dev/countdown/delete/{id} [post]
func (CountdownController) RemoveCountdown(c *gin.Context) {
	id := c.Param("devId")
	if id == "" {
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrAppDeviceIdEmpty, nil)
		return
	}
	err := countdownService.DeleteCountdown(id)
	if err != nil {
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		iotlogger.LogHelper.Helper.Error("remove count down error: ", err.Error())
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, nil)
}
