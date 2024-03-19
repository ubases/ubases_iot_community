package routers

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_smart_speaker_service/cached"
	"cloud_platform/iot_smart_speaker_service/entitys"
	"cloud_platform/iot_smart_speaker_service/service/common"
	"cloud_platform/iot_smart_speaker_service/service/tianmao"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

const (
	DiscoveryDevices     = "DiscoveryDevices"
	SetDeviceProperty    = "thing.attribute.set"
	AdjustDeviceProperty = "thing.attribute.adjust"
)

func TMiotGateWay(c *gin.Context) {
	// 定义设备列表变量
	var err error
	devices := []entitys.TmDevice{}

	// 获取设备列表查询接口请求原始数据
	data, err := c.GetRawData()
	if err != nil {
		c.String(http.StatusOK, "获取原始数据错误")
		return
	}
	commonRes := entitys.Common{}
	if err := json.Unmarshal(data, &commonRes); err != nil {
		c.String(http.StatusOK, err.Error())
	}
	iotlogger.LogHelper.Helper.Debug("discovery devices: ", commonRes)
	token := gjson.GetBytes(data, "payload.accessToken").String()

	ti, err := manager.LoadAccessToken(context.Background(), token)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("加载访问密钥错误: ", err)
		c.String(http.StatusOK, "加载访问密钥错误")
		return
	}
	userId := ti.GetUserID()
	iotlogger.LogHelper.Helper.Debug("userId: ", userId)

	switch commonRes.Header.Name {
	case DiscoveryDevices:
		devices, err = common.DiscoveryDevices(userId, "tianmao")
		if err != nil {
			c.String(http.StatusOK, err.Error())
			return
		}
		iotlogger.LogHelper.Helper.Debug("voice device list: ", devices)
		responseDeviceDiscovery(c, commonRes.Header.MessageId, devices)
		err = cached.RedisStore.GetClient().Set(context.Background(), fmt.Sprintf(iotconst.VoiceUserTokenKey, userId), token, time.Hour*24*2).Err()
		if err != nil {
			iotlogger.LogHelper.Helper.Error("set voice user token error: ", err)
		}
		return
	case SetDeviceProperty:
		devResp, err := tianmao.SetDeviceProperty(data, userId, false, "tianmao")
		if err != nil {
			c.String(http.StatusOK, err.Error())
			return
		}
		responseSetDeviceProperty(c, commonRes.Header.MessageId, devResp)
		return
	case AdjustDeviceProperty:
		devResp, err := tianmao.SetDeviceProperty(data, userId, true, "tianmao")
		if err != nil {
			c.String(http.StatusOK, err.Error())
			return
		}
		responseSetDeviceProperty(c, commonRes.Header.MessageId, devResp)
		return
	default:
		iotlogger.LogHelper.Helper.Error("tian mao unexpected control type: ", commonRes.Header.Name)
	}
}

func responseDeviceDiscovery(c *gin.Context, messageId string, devices []entitys.TmDevice) {
	c.JSON(http.StatusOK, entitys.NewTmDeviceDiscoveryResp(messageId, devices))
}

func responseSetDeviceProperty(c *gin.Context, messageId string, deviceResponse []entitys.DeviceResponse) {
	c.JSON(http.StatusOK, entitys.NewDeviceResponse(messageId, deviceResponse))
}
