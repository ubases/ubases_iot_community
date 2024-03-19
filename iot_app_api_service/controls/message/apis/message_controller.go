package apis

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/message/entitys"
	"cloud_platform/iot_app_api_service/controls/message/services"
	"cloud_platform/iot_common/iotgin"

	"github.com/gin-gonic/gin"
)

var Messagecontroller MessageController

type MessageController struct{} //部门操作控制器

var messageServices = services.MessageService{}

// @Summary 获取消息统计数据（红点数据）
// @Description
// @Tags 设备
// @Accept application/json
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/message/reddot [get]
func (s MessageController) GetMessageRedDot(c *gin.Context) {
	userId := controls.GetUserId(c)
	lang := controls.GetLang(c)
	res, err := messageServices.SetContext(controls.WithUserContext(c)).GetMessageRedDot(lang, userId)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 获取设备消息统计列表
func (s MessageController) GetMessageReportCount(c *gin.Context) {
	iotgin.ResSuccess(c, nil)
}

// 获取家庭消息、系统消息列表
func (s MessageController) GetHomeAndSystemMessage(c *gin.Context) {
	iotgin.ResSuccess(c, nil)
}

// ClearMessage 清空家庭消息、系统消息
// @Summary 清空家庭消息、系统消息
// @Description
// @Tags 设备
// @Accept application/json
// @Param type path string true "消息类型"
// @Param devId path string true "设备Id"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/message/allDel/{type} [post]
func (s MessageController) ClearMessage(c *gin.Context) {
	typeStr := c.Param("type")
	if typeStr == "" {
		iotgin.ResBadRequest(c, "type")
		return
	}
	deviceId := c.Query("devId")
	userId := controls.GetUserId(c)
	err := messageServices.SetContext(controls.WithUserContext(c)).ClearMessage(userId, typeStr, deviceId)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, nil)
}

// DeleteMessage 删除指定的消息
// @Summary 删除指定的消息
// @Description
// @Tags 设备
// @Accept application/json
// @Param id path string true "消息Id"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/message/del/{id} [post]
func (s MessageController) DeleteMessage(c *gin.Context) {
	userId := c.Param("id")
	err := messageServices.SetContext(controls.WithUserContext(c)).DeleteByIdMessage(userId)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, nil)
}

// GetMessageGroupList 获取场景日志列表数据
// @Summary 获取场景日志列表数据
// @Description
// @Tags 设备
// @Accept application/json
// @Param type path string true "消息类型"
// @Param devId query string true "devId"
// @Param messageId query string true "messageId"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/message/grouplist/{type} [get]
func (s MessageController) GetMessageGroupList(c *gin.Context) {
	typeStr := c.Param("type")
	if typeStr == "" {
		iotgin.ResBadRequest(c, "type")
		return
	}
	devId := c.Query("devId")
	messageId := c.Query("messageId")
	userId := controls.GetUserId(c)
	res, err := messageServices.SetContext(controls.WithUserContext(c)).QueryMessageGroupList(userId, typeStr, devId, messageId)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// GetMessageList 消息查询
// @Summary 消息查询
// @Description
// @Tags 设备
// @Accept application/json
// @Param type path string true "消息类型"
// @Param data body entitys.MpMessageUserOutQuery true "devId"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/message/list/{type} [get]
func (s MessageController) GetMessageList(c *gin.Context) {
	typeStr := c.Param("type")
	if typeStr == "" {
		iotgin.ResBadRequest(c, "type")
		return
	}
	var filter entitys.MpMessageUserOutQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if filter.Query == nil {
		filter.Query = new(entitys.MpMessageUserOutFilter)
	}
	userId := controls.GetUserId(c)
	res, total, err := messageServices.SetContext(controls.WithUserContext(c)).QueryMessageList(userId, typeStr, filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

// 获取场景日志列表数据
func (s MessageController) GetSceneLogList(c *gin.Context) {
	iotgin.ResSuccess(c, nil)
}

// 清空场景日志
func (s MessageController) DeleteSceneLog(c *gin.Context) {
	iotgin.ResSuccess(c, nil)
}

// 获取设备消息列表
func (s MessageController) GetDeviceMessage(c *gin.Context) {
	iotgin.ResSuccess(c, nil)
}

// 清空设备消息
func (s MessageController) DeleteDeviceMessage(c *gin.Context) {
	iotgin.ResSuccess(c, nil)
}
