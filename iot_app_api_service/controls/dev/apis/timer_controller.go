package apis

import (
	"cloud_platform/iot_app_api_service/cached"
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/dev/entitys"
	"cloud_platform/iot_app_api_service/controls/dev/services"
	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_common/iotutil"
	"github.com/gin-gonic/gin"
	goerrors "go-micro.dev/v4/errors"
)

var Timercontroller TimerController

type TimerController struct {
} //定时器操作控制器

var timerService = services.AppTimerService{}

// AddTimer 创建定时器
// @Summary 创建定时器
// @Description
// @Tags 设备
// @Accept application/json
// @Param data body entitys.IotDeviceTimerEntitys true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /dev/timer/add [post]
func (TimerController) AddTimer(c *gin.Context) {
	req := entitys.IotDeviceTimerEntitys{}
	if err := c.BindJSON(&req); err != nil {
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrAppRequestParam, nil)
		return
	}
	if req.DeviceId == "" {
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrAppDeviceIdEmpty, nil)
		return
	}
	//Functions 新定时器功能参数、 FuncKey历史定时器参数
	if req.Functions == nil && req.FuncKey == "" {
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrAppFuncKeyEmpty, nil)
		return
	}
	//兼容历史定时功能（当无functions数据，则以funcKey和funcValue为准）
	if req.Functions == nil && req.FuncKey != "" {
		//兼容逻辑，原定时器funcKey和funcValue为字符串，推送给设备的时候也会推送字符串，此处需要对推送内容进行转换
		funcMap := map[string]string{"funcKey": req.FuncKey, "funcValue": iotutil.ToString(req.FuncValue)}
		funcMapStr := iotutil.MapStringToInterface(funcMap)
		var funcObj iotstruct.TimerFunctions
		err := iotutil.JsonToStruct(iotutil.ToString(funcMapStr), &funcObj)
		if err != nil {
			ioterrs.Response(c, cached.RedisStore, ioterrs.ErrAppRequestParam, nil)
			return
		}
		req.Functions = []iotstruct.TimerFunctions{funcObj}
	}

	//如果Functions不存在或者数量为0，则返回异常
	if req.Functions == nil || len(req.Functions) == 0 {
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrAppFuncKeyEmpty, nil)
		return
	}

	userId := controls.GetUserId(c)
	req.UserId = iotutil.ToInt64(userId)
	req.CreatedBy = iotutil.ToInt64(userId)
	req.Timezone = controls.GetTimezone(c)
	req.RegionServerId = controls.GetRegionInt(c)
	timerId, err := timerService.SetContext(controls.WithUserContext(c)).AddTimer(req)
	if err != nil {
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		iotlogger.LogHelper.Helper.Error("add timer error: ", err.Error())
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, timerId)
}

// UpdateTimer 创建定时器
// @Summary 创建定时器
// @Description
// @Tags 设备
// @Accept application/json
// @Param id path string true "定时Id"
// @Param data body entitys.IotDeviceTimerUpdate true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /dev/timer/update/{id} [post]
func (TimerController) UpdateTimer(c *gin.Context) {
	var (
		id  = c.Param("id")
		err error
	)
	if id == "" {
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrAppRequestParam, nil)
		return
	}
	req := entitys.IotDeviceTimerUpdate{}
	if err := c.BindJSON(&req); err != nil {
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrAppRequestParam, nil)
		return
	}
	req.Id, err = iotutil.ToInt64AndErr(id)
	if err != nil {
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrAppRequestParam, nil)
		return
	}
	//Functions 新定时器功能参数、 FuncKey历史定时器参数
	if req.Functions == nil && req.FuncKey == "" {
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrAppFuncKeyEmpty, nil)
		return
	}
	iotlogger.LogHelper.Helper.Debug("fuctions: ", req.Functions)
	//兼容历史定时功能（当无functions数据，则以funcKey和funcValue为准）
	if req.Functions == nil && req.FuncKey != "" {
		//兼容逻辑，原定时器funcKey和funcValue为字符串，推送给设备的时候也会推送字符串，此处需要对推送内容进行转换
		funcMap := map[string]string{"funcKey": req.FuncKey, "funcValue": iotutil.ToString(req.FuncValue)}
		funcMapStr := iotutil.MapStringToInterface(funcMap)
		var funcObj iotstruct.TimerFunctions
		err := iotutil.JsonToStruct(iotutil.ToString(funcMapStr), &funcObj)
		if err != nil {
			ioterrs.Response(c, cached.RedisStore, ioterrs.ErrAppRequestParam, nil)
			return
		}
		req.Functions = []*iotstruct.TimerFunctions{&funcObj}
	}

	//如果Functions不存在或者数量为0，则返回异常
	if req.Functions == nil || len(req.Functions) == 0 {
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrAppFuncKeyEmpty, nil)
		return
	}

	req.Timezone = controls.GetTimezone(c)
	err = timerService.SetContext(controls.WithUserContext(c)).UpdateTimer(req)
	if err != nil {
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		iotlogger.LogHelper.Helper.Error("update timer error: ", err.Error())
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, nil)
}

// RemoveTimer 移除定时器
// @Summary 移除定时器
// @Description
// @Tags 设备
// @Accept application/json
// @Param id path string true "定时Id"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /dev/timer/delete/{id} [post]
func (TimerController) RemoveTimer(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrAppRequestParam, nil)
		return
	}
	err := timerService.SetContext(controls.WithUserContext(c)).RemoveTimer(iotutil.ToInt64(id))
	if err != nil {
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		iotlogger.LogHelper.Helper.Error("remove timer error: ", err.Error())
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, nil)
}

// TimerInfo 定时器信息查询
// @Summary 定时器信息查询
// @Description
// @Tags 设备
// @Accept application/json
// @Param id path string true "定时Id"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /dev/timer/info/{id} [get]
func (TimerController) TimerInfo(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrAppRequestParam, nil)
		return
	}
	data, err := timerService.SetContext(controls.WithUserContext(c)).TimerInfo(iotutil.ToInt64(id))
	if err != nil {
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		iotlogger.LogHelper.Helper.Error("get timer detail error: ", err.Error())
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, data)
}

// TimerList 定时器列表
// @Summary 定时器列表
// @Description
// @Tags 设备
// @Accept application/json
// @Param data body entitys.TimerListQuery true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /dev/timer/list [post]
func (TimerController) TimerList(c *gin.Context) {
	var filter entitys.TimerListQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrAppRequestParam, nil)
		return
	}
	if filter.DevId == "" {
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrAppDeviceIdEmpty, nil)
		return
	}
	//userId:= controls.GetUserId(c)
	data, count, err := timerService.SetContext(controls.WithUserContext(c)).TimerList(filter.DevId, filter)
	if err != nil {
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		iotlogger.LogHelper.Helper.Error("get timer list error: ", err.Error())
		return
	}
	ioterrs.ResponsePage(c, cached.RedisStore, ioterrs.Success, data, count, 1)
}

// EnabledTimer 开启定时器
// @Summary 开启定时器
// @Description
// @Tags 设备
// @Accept application/json
// @Param id path string true "定时Id"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /dev/timer/open/{id} [post]
func (TimerController) EnabledTimer(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrAppRequestParam, nil)
		return
	}
	tType := c.Query("type")
	if tType == "" {
		tType = "0"
	}
	err := timerService.SetContext(controls.WithUserContext(c)).EnabledTimer(iotutil.ToInt64(id), iotutil.ToInt(tType))
	if err != nil {
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		iotlogger.LogHelper.Helper.Error("enable timer error: ", err.Error())
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, nil)
}

// DisabledTimer 关闭定时器
// @Summary 关闭定时器
// @Description
// @Tags 设备
// @Accept application/json
// @Param id path string true "定时Id"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /dev/timer/close/{id} [post]
func (TimerController) DisabledTimer(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrAppRequestParam, nil)
		return
	}
	err := timerService.SetContext(controls.WithUserContext(c)).DisabledTimer(iotutil.ToInt64(id))
	if err != nil {
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		iotlogger.LogHelper.Helper.Error("disable timer error: ", err.Error())
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, nil)
}
