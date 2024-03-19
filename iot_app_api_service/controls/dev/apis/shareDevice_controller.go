package apis

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/dev/entitys"
	"cloud_platform/iot_app_api_service/controls/dev/services"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"

	"github.com/gin-gonic/gin"
)

var ShareDevicecontroller ShareDeviceController

type ShareDeviceController struct {
}

var appShareDeviceService = services.AppShareDeviceService{}

// ShareDeviceList 分享设备列表
// @Summary 分享设备列表
// @Description
// @Tags 设备
// @Accept application/json
// @Param homeId path string true "家庭Id"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/dev/shareDeviceList/{homeId} [get]
func (ShareDeviceController) ShareDeviceList(c *gin.Context) {
	homeId := c.Param("homeId")
	if homeId == "" {
		iotgin.ResBadRequest(c, "homeId is empty")
		return
	}
	data, err := appShareDeviceService.ShareDeviceList(homeId)
	if err != nil {
		iotgin.ResBusinessP(c, err.Error())
		return
	}
	iotgin.ResSuccess(c, data)
}

// ShareUserList 分享用户列表
// @Summary 分享用户列表
// @Description
// @Tags 设备
// @Accept application/json
// @Param homeId query string true "家庭Id"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/dev/shareUserList/{devId} [get]
func (ShareDeviceController) ShareUserList(c *gin.Context) {
	devId := c.Param("devId")
	if devId == "" {
		iotgin.ResBadRequest(c, "devId")
		return
	}

	homeId := c.Query("homeId")
	if homeId == "" {
		iotgin.ResBadRequest(c, "homeId")
		return
	}
	data, err := appShareDeviceService.ShareUserList(devId, homeId)
	if err != nil {
		iotgin.ResBusinessP(c, err.Error())
		return
	}
	iotgin.ResSuccess(c, data)
}

// AddShared 创建定时器
// @Summary 创建定时器
// @Description
// @Tags 设备
// @Accept application/json
// @Param data body entitys.Addshared true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/dev/addShared [post]
func (ShareDeviceController) AddShared(c *gin.Context) {
	req := entitys.Addshared{}
	if err := c.BindJSON(&req); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	if req.ProductKey == "" {
		iotgin.ResBadRequest(c, "产品Key不能为空")
	}
	var (
		userId         = controls.GetUserId(c)
		appKey         = controls.GetAppKey(c)
		tenantId       = controls.GetTenantId(c)
		regionServerId = controls.GetRegionInt(c)
	)
	code, msg := appShareDeviceService.SetContext(controls.WithUserContext(c)).AddShared(iotutil.ToString(userId), appKey, tenantId, regionServerId, req)
	if code != 0 {
		iotlogger.LogHelper.Errorf(msg)
		iotgin.ResFailCode(c, msg, code)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// ReceiveSharedList 接受共享设备列表
// @Summary 接受共享设备列表
// @Description
// @Tags 设备
// @Accept application/json
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/dev/receiveSharedList [get]
func (ShareDeviceController) ReceiveSharedList(c *gin.Context) {
	userId := controls.GetUserId(c)
	data, err := appShareDeviceService.ReceiveSharedList(iotutil.ToInt64(userId))
	if err != nil {
		iotgin.ResBusinessP(c, err.Error())
		return
	}
	iotgin.ResSuccess(c, data)
}

// CancelShare 取消共享
// @Summary 取消共享
// @Description
// @Tags 设备
// @Accept application/json
// @Param data body entitys.CancelShare true "请求参数"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/dev/cancelShare [post]
func (ShareDeviceController) CancelShare(c *gin.Context) {
	req := entitys.CancelShare{}
	if err := c.BindJSON(&req); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	var (
		userId   = controls.GetUserId(c)
		appKey   = controls.GetAppKey(c)
		tenantId = controls.GetTenantId(c)
	)
	err := appShareDeviceService.CancelShare(req, iotutil.ToString(userId), appKey, tenantId)
	if err != nil {
		iotgin.ResBusinessP(c, err.Error())
		return
	}
	iotgin.ResSuccessMsg(c)
}

// ReceiveShare 接受共享
// @Summary 接受共享
// @Description
// @Tags 设备
// @Accept application/json
// @Param id path string true "共享Id"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/dev/receiveShare/{id} [post]
func (ShareDeviceController) ReceiveShare(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id is empty")
		return
	}
	var (
		userId   = controls.GetUserId(c)
		appKey   = controls.GetAppKey(c)
		tenantId = controls.GetTenantId(c)
	)
	err := appShareDeviceService.ReceiveShare(iotutil.ToInt64(userId), iotutil.ToInt64(id), appKey, tenantId)
	if err != nil {
		iotgin.ResBusinessP(c, err.Error())
		return
	}
	iotgin.ResSuccessMsg(c)
}

// CancelReceiveShared 取消接受共享
// @Summary 取消接受共享
// @Description
// @Tags 设备
// @Accept application/json
// @Param data body entitys.CancelReceiveShared true "请求参数"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/dev/cancelReceiveShared [post]
func (ShareDeviceController) CancelReceiveShared(c *gin.Context) {
	req := entitys.CancelReceiveShared{}
	if err := c.BindJSON(&req); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	var (
		userId   = controls.GetUserId(c)
		appKey   = controls.GetAppKey(c)
		tenantId = controls.GetTenantId(c)
	)
	err := appShareDeviceService.CancelReceiveShared(req, iotutil.ToString(userId), appKey, tenantId)
	if err != nil {
		iotgin.ResBusinessP(c, err.Error())
		return
	}
	iotgin.ResSuccessMsg(c)
}
