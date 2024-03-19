package routers

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_smart_speaker_service/entitys"
	"cloud_platform/iot_smart_speaker_service/service/xiaomi"
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

var xiaomiRouter XiaomiIotGateWayRouter

type XiaomiIotGateWayRouter struct {
}

func (s *XiaomiIotGateWayRouter) GateWay(c *gin.Context) {
	// 定义设备列表变量
	var err error
	token := c.Request.Header.Get("User-Token")
	iotlogger.LogHelper.Helper.Debug("进入小米Iot, token:", token)
	// 获取设备列表查询接口请求原始数据
	data, err := c.GetRawData()
	if err != nil {
		c.String(http.StatusOK, "获取原始数据错误")
		return
	}
	iotlogger.LogHelper.Helper.Debug("进入小米Iot, data:", string(data))
	common := entitys.XiaomiRequest{}
	if err := json.Unmarshal(data, &common); err != nil {
		c.String(http.StatusOK, err.Error())
	}

	ti, err := manager.LoadAccessToken(context.Background(), token)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("加载访问密钥错误: ", err)
		c.String(http.StatusOK, "加载访问密钥错误")
		return
	}
	userId := ti.GetUserID()
	//userId := "783946466434580480" //测试环境用户
	iotlogger.LogHelper.Helper.Debug("userId: ", userId)
	xiaomi.RunIntent(&common, data, userId, token)
	s.response(c, common)
}

// 返回错误信息
func (s *XiaomiIotGateWayRouter) response(c *gin.Context, res entitys.XiaomiRequest) {
	c.Writer.Header().Set("Content-Type", "application/json")
	iotlogger.LogHelper.Helper.Debug("小米Iot, request:", iotutil.ToString(res))
	c.JSON(http.StatusOK, res)
}
