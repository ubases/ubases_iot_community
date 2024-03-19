package apis

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/dev/entitys"
	"cloud_platform/iot_app_api_service/controls/dev/services"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"

	"github.com/gin-gonic/gin"
)

var DeviceGroupcontroller DeviceGroupController

type DeviceGroupController struct {
}

var appDeviceGroupService = services.AppDeviceGroupService{}

// DevGroupInfo 设备群组信息
// @Summary 设备群组信息
// @Description
// @Tags 设备
// @Accept application/json
// @Param groupId path string true "请求参数"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/dev/group/info/{groupId} [post]
func (DeviceGroupController) DevGroupInfo(c *gin.Context) {
	groupId := c.Param("groupId")
	if groupId == "" {
		iotgin.ResBadRequest(c, "groupId")
		return
	}
	data, err := appDeviceGroupService.SetContext(controls.WithUserContext(c)).DevGroupInfo(iotutil.ToInt64(groupId))
	if err != nil {
		iotgin.ResBusinessP(c, err.Error())
		return
	}
	iotgin.ResSuccess(c, data)
}

// DevListByProductKey 可以添加到一个群组的设备列表
// @Summary 可以添加到一个群组的设备列表
// @Description
// @Tags 设备
// @Accept application/json
// @Param productKey path string true "请求参数"
// @Param homeId path string true "请求参数"
// @Param groupId path string true "请求参数"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/dev/group/devListByProductKey [get]
func (DeviceGroupController) DevListByProductKey(c *gin.Context) {
	productKey := c.Query("productKey")
	if productKey == "" {
		iotgin.ResBadRequest(c, "productKey")
		return
	}
	homeId := c.Query("homeId")
	if homeId == "" {
		iotgin.ResBadRequest(c, "homeId")
		return
	}
	groupId := c.Query("groupId")
	data, err := appDeviceGroupService.SetContext(controls.WithUserContext(c)).DevListByProductKey(productKey, homeId, groupId)
	if err != nil {
		iotgin.ResBusinessP(c, err.Error())
		return
	}
	iotgin.ResSuccess(c, data)
}

// DevGroupDevList 群组详细信息
// @Summary 群组详细信息
// @Description
// @Tags 设备
// @Accept application/json
// @Param groupId path string true "请求参数"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/dev/group/devList/{groupId} [get]
func (DeviceGroupController) DevGroupDevList(c *gin.Context) {
	groupId := c.Param("groupId")
	if groupId == "" {
		iotgin.ResBadRequest(c, "groupId")
		return
	}
	data, err := appDeviceGroupService.SetContext(controls.WithUserContext(c)).DevGroupDevList(iotutil.ToInt64(groupId))
	if err != nil {
		iotgin.ResBusinessP(c, err.Error())
		return
	}
	iotgin.ResSuccess(c, data)
}

// UpsertGroup 创建修改群组
// @Summary 创建修改群组
// @Description
// @Tags 设备
// @Accept application/json
// @Param data body entitys.UpsertGroup true "请求参数"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/dev/group/upsert [post]
func (DeviceGroupController) UpsertGroup(c *gin.Context) {
	req := entitys.UpsertGroup{}
	if err := c.BindJSON(&req); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	userId := controls.GetUserId(c)
	result, err := appDeviceGroupService.SetContext(controls.WithUserContext(c)).UpsertGroup(req, iotutil.ToInt64(userId))
	if err != nil {
		iotgin.ResBusinessP(c, err.Error())
		return
	}
	iotgin.ResSuccess(c, result)
}

// RemoveGroup 移除群组
// @Summary 移除群组
// @Description
// @Tags 设备
// @Accept application/json
// @Param groupId path string true "请求参数"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/dev/group/remove/{groupId} [post]
func (DeviceGroupController) RemoveGroup(c *gin.Context) {
	groupId := c.Param("groupId")
	if groupId == "" {
		iotgin.ResBadRequest(c, "groupId")
		return
	}
	userId := controls.GetUserId(c)
	err := appDeviceGroupService.SetContext(controls.WithUserContext(c)).RemoveGroup(iotutil.ToInt64(groupId), iotutil.ToInt64(userId))
	if err != nil {
		iotgin.ResBusinessP(c, err.Error())
		return
	}
	iotgin.ResSuccessMsg(c)
}

// Execute 群组控制
// @Summary 群组控制
// @Description
// @Tags 设备
// @Accept application/json
// @Param data body entitys.GroupExecute true "请求参数"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/dev/group/execute [post]
func (DeviceGroupController) Execute(c *gin.Context) {
	req := entitys.GroupExecute{}
	if err := c.BindJSON(&req); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	groupIdInt64, _ := iotutil.ToInt64AndErr(req.GroupId)

	if groupIdInt64 == 0 {
		iotgin.ResBadRequest(c, "groupId")
		return
	}
	err := appDeviceGroupService.SetContext(controls.WithUserContext(c)).Execute(req)
	if err != nil {
		iotgin.ResBusinessP(c, err.Error())
		return
	}
	iotgin.ResSuccessMsg(c)
}

// DevGroupTsl 群组设备物模型
// @Summary 群组设备物模型
// @Description
// @Tags 设备
// @Accept application/json
// @Param groupId path string true "请求参数"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/dev/group/tsl/{groupId} [get]
func (DeviceGroupController) DevGroupTsl(c *gin.Context) {
	groupId := c.Param("groupId")
	if groupId == "" {
		iotgin.ResBadRequest(c, "groupId")
		return
	}
	language := c.GetHeader("lang")
	tenantId := c.GetHeader("tenantId")
	data, err := appDeviceGroupService.SetContext(controls.WithUserContext(c)).DevGroupTsl(iotutil.ToInt64(groupId), language, tenantId)
	if err != nil {
		iotgin.ResBusinessP(c, err.Error())
		return
	}
	iotgin.ResSuccess(c, data)
}

// ExecuteSwitch 群组开关控制
// @Summary 群组开关控制
// @Description
// @Tags 设备
// @Accept application/json
// @Param data body entitys.GroupExecute true "请求参数"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/dev/group/executeSwitch [post]
func (DeviceGroupController) ExecuteSwitch(c *gin.Context) {
	req := entitys.GroupExecute{}
	if err := c.BindJSON(&req); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	groupIdInt64, _ := iotutil.ToInt64AndErr(req.GroupId)

	if groupIdInt64 == 0 {
		iotgin.ResBadRequest(c, "groupId")
		return
	}

	userId := controls.GetUserId(c)
	err := appDeviceGroupService.SetContext(controls.WithUserContext(c)).ExecuteSwitch(iotutil.ToString(groupIdInt64), iotutil.ToInt64(userId))
	if err != nil {
		iotgin.ResBusinessP(c, err.Error())
		return
	}
	iotgin.ResSuccessMsg(c)
}
