package apis

import (
	"cloud_platform/iot_app_api_service/cached"
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_proto/protos/protosService"

	"github.com/gin-gonic/gin"
)

var DeviceFunctionSetcontroller DeviceFunctionSetController

type DeviceFunctionSetController struct {
}

// UpdateFunction 修改设备的功能描述信息
// @Summary 修改设备的功能描述信息
// @Description
// @Tags room
// @Accept application/json
// @Param data body protosService.IotDeviceFunctionSet true "请求参数结构体"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /dev/functionDesc/update [post]
func (DeviceFunctionSetController) UpdateFunction(c *gin.Context) {
	req := protosService.IotDeviceFunctionSet{}
	if err := c.BindJSON(&req); err != nil {
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrAppRequestParam, nil)
		return
	}
	res, err := rpc.IotDeviceInfoService.SaveDeviceFunctionSet(controls.WithUserContext(c), &req)
	if err != nil {
		iotgin.ResBusinessP(c, err.Error())
		return
	}
	if res.Code != 200 {
		iotgin.ResBusinessP(c, res.Message)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// BatchUpdateFunction 修改设备的功能描述信息（批量操作）
// @Summary 修改设备的功能描述信息（批量操作）
// @Description
// @Tags room
// @Accept application/json
// @Param data body protosService.IotDeviceFunctionBatchSet true "请求参数结构体"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /dev/functionDesc/batchUpdate [post]
func (DeviceFunctionSetController) BatchUpdateFunction(c *gin.Context) {
	req := protosService.IotDeviceFunctionBatchSet{}
	if err := c.BindJSON(&req); err != nil {
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrAppRequestParam, nil)
		return
	}
	res, err := rpc.IotDeviceInfoService.SaveDeviceFunctionBatchSet(controls.WithUserContext(c), &req)
	if err != nil {
		iotgin.ResBusinessP(c, err.Error())
		return
	}
	if res.Code != 200 {
		iotgin.ResBusinessP(c, res.Message)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// GetDeviceFunctionSetList 获取设备功能设置列表
// @Summary 获取设备功能设置列表
// @Description
// @Tags 设备控制面板
// @Accept application/json
// @Param deviceId query string true "设备Id"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /dev/functionDesc/list [get]
func (DeviceFunctionSetController) GetDeviceFunctionSetList(c *gin.Context) {
	deviceId := c.Query("deviceId")
	res, err := rpc.IotDeviceInfoService.GetDeviceFunctionSetList(controls.WithUserContext(c), &protosService.IotDeviceFunctionSet{
		DeviceId: deviceId,
	})
	if err != nil {
		iotgin.ResBusinessP(c, err.Error())
		return
	}
	if res.Code != 200 {
		iotgin.ResBusinessP(c, res.Message)
		return
	}
	iotgin.ResSuccess(c, res.Data)
}
