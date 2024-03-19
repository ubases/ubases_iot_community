package apis

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/dev/entitys"
	"cloud_platform/iot_app_api_service/controls/dev/services"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"

	"github.com/gin-gonic/gin"
)

var Devcontroller DevController

type DevController struct {
} //家庭房间操作控制器

var devService = services.AppDevService{}

// DeviceInfo 设备信息
// @Summary 通用设备信息
// @Description
// @Tags room
// @Accept application/json
// @Param devId query string true "设备Id"
// @success 200 {object} iotgin.ResponseModel "执行结果"
// @Router /dev/deviceInfo/{devId} [get]
func (DevController) DeviceInfo(c *gin.Context) {
	devId := c.Param("devId")
	if devId == "" {
		iotgin.ResBusinessP(c, "devId is empty")
		return
	}
	userId := controls.GetUserId(c)
	devSecret, _ := c.Get("DevSecrt")
	language := c.GetHeader("lang")
	data, _ := devService.SetContext(controls.WithUserContext(c)).DeviceInfo(devId, iotutil.ToInt64(userId), iotutil.ToString(devSecret), language)
	iotgin.ResSuccess(c, data)
}

// RemoveDev 移除设备
// @Summary 移除设备
// @Description
// @Tags room
// @Accept application/json
// @Param data body entitys.RemoveDevFilter true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /dev/removeDev [post]
func (DevController) RemoveDev(c *gin.Context) {
	req := entitys.RemoveDevFilter{}
	if err := c.BindJSON(&req); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	userId := controls.GetUserId(c)
	if len(req.DevIdList) == 0 {
		iotgin.ResBadRequest(c, "devIdList")
		return
	}
	err := devService.SetContext(controls.WithUserContext(c)).RemoveDev(req, iotutil.ToInt64(userId))
	if err != nil {
		iotgin.ResBusinessP(c, err.Error())
		return
	}
	iotgin.ResSuccessMsg(c)
}

// RemoveRoomDev 移除房间设备
// @Summary 移除房间设备
// @Description
// @Tags room
// @Accept application/json
// @Param data body entitys.RemoveDevFilter true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /dev/removeRoomDev [post]
func (DevController) RemoveRoomDev(c *gin.Context) {
	req := entitys.RemoveDevFilter{}
	if err := c.BindJSON(&req); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	if req.RoomId == "" || req.RoomId == "0" {
		iotgin.ResBadRequest(c, "roomId")
		return
	}
	if req.DevId == "" {
		iotgin.ResBadRequest(c, "devId")
		return
	}
	userId := controls.GetUserId(c)
	err := devService.SetContext(controls.WithUserContext(c)).RemoveDev(req, iotutil.ToInt64(userId))
	if err != nil {
		iotgin.ResBusinessP(c, err.Error())
		return
	}
	iotgin.ResSuccessMsg(c)
}

// UpdateDev 修改设备
// @Summary 修改设备
// @Description
// @Tags room
// @Accept application/json
// @Param devId path string true "设备Id"
// @Param data body entitys.UpdateDevFilter true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /dev/updateDev/{devId} [post]
func (DevController) UpdateDev(c *gin.Context) {
	devId := c.Param("devId")
	if devId == "" {
		iotgin.ResBadRequest(c, "devId")
		return
	}
	req := entitys.UpdateDevFilter{}
	if err := c.BindJSON(&req); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	homeId, _ := iotutil.ToInt64AndErr(req.HomeId)
	if homeId == 0 {
		iotgin.ResBadRequest(c, "homeId")
		return
	}
	req.DevId = devId
	userId := controls.GetUserId(c)
	err := devService.UpdateDev(req, iotutil.ToInt64(userId))
	if err != nil {
		iotgin.ResBusinessP(c, err.Error())
		return
	}
	iotgin.ResSuccessMsg(c)
}

// AddDev 新增设备
// @Summary 新增设备
// @Description
// @Tags room
// @Accept application/json
// @Param data body entitys.AddDevFilter true "新增参数"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /dev/addDev [post]
func (DevController) AddDev(c *gin.Context) {
	req := entitys.AddDevFilter{}
	if err := c.BindJSON(&req); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	homeId, _ := iotutil.ToInt64AndErr(req.HomeId)
	if homeId == 0 {
		iotgin.ResBadRequest(c, "homeId")
		return
	}
	roomId, _ := iotutil.ToInt64AndErr(req.RoomId)
	if roomId == 0 {
		iotgin.ResBadRequest(c, "roomId")
		return
	}
	devId := req.DevId
	if devId == "" {
		iotgin.ResBadRequest(c, "devId")
		return
	}
	sort, _ := iotutil.ToInt64AndErr(req.Sort)
	if sort == 0 {
		iotgin.ResBadRequest(c, "sort")
		return
	}
	err := devService.AddDev(req)
	if err != nil {
		iotgin.ResBusinessP(c, err.Error())
		return
	}
	iotgin.ResSuccessMsg(c)
}
