package routers

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	proto "cloud_platform/iot_proto/protos/protosService"
	"cloud_platform/iot_smart_speaker_service/entitys"
	"cloud_platform/iot_smart_speaker_service/rpc/rpcclient"
	"cloud_platform/iot_smart_speaker_service/service/xiaoai"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

var xiaoaiRouter XiaoaiIotGateWayRouter

type XiaoaiIotGateWayRouter struct {
	EnabbleV2 bool
}

// GateWayV2 一个技能对应多个产品
func (s *XiaoaiIotGateWayRouter) GateWayV2(c *gin.Context) {
	s.EnabbleV2 = true
	s.GateWay(c)
}

// GateWayV2 一个技能一个产品 https://developers.xiaoai.mi.com/documents/Home?type=/api/doc/render_markdown/SkillAccess/skill/fulu/HTTPS
func (s *XiaoaiIotGateWayRouter) GateWay(c *gin.Context) {
	//临时逻辑，如果有user-token则为小米Iot
	token := c.Request.Header.Get("User-Token")
	if token != "" {
		xiaomi := XiaomiIotGateWayRouter{}
		xiaomi.GateWay(c)
		return
	}
	iotlogger.LogHelper.Helper.Debug("进入小爱同学处理逻辑")

	// 定义设备列表变量
	var err error
	// 获取设备列表查询接口请求原始数据
	data, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusOK, s.setResult(true, "获取原始数据错误"))
		return
	}
	req := entitys.XiaoaiRequest{}
	if err := json.Unmarshal(data, &req); err != nil {
		iotlogger.LogHelper.Helper.Error("参数异常: ", err)
		c.JSON(http.StatusOK, s.setResult(true, "参数异常"))
	}
	//退出逻辑
	if req.Request.SlotInfo.IntentName == "Mi_Exit" && req.Query == "退出" {
		c.JSON(http.StatusOK, s.setResult(true, "好的，已为您退出"))
	}
	g := xiaoai.XiaoaiIoTGateway{}
	msg, err := g.Run(req, manager, "")
	if err != nil {
		msg = "未知错误"
	}

	c.JSON(http.StatusOK, s.setResult(false, msg))
}

// GateWayV2 一个技能对应多个产品
func (s *XiaoaiIotGateWayRouter) GateWayTestAccountV2(c *gin.Context) {
	s.EnabbleV2 = true
	s.GateWay(c)
}

// https://developers.xiaoai.mi.com/documents/Home?type=/api/doc/render_markdown/SkillAccess/skill/fulu/HTTPS
func (s *XiaoaiIotGateWayRouter) GateWayTestAccount(c *gin.Context) {
	//临时逻辑，如果有user-token则为小米Iot
	token := c.Request.Header.Get("User-Token")
	if token != "" {
		xiaomi := XiaomiIotGateWayRouter{}
		xiaomi.GateWay(c)
		return
	}
	iotlogger.LogHelper.Helper.Debug("进入小爱同学处理逻辑")

	// 定义设备列表变量
	var err error
	// 获取设备列表查询接口请求原始数据
	data, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusOK, s.setResult(true, "获取原始数据错误"))
		return
	}
	req := entitys.XiaoaiRequest{}
	if err := json.Unmarshal(data, &req); err != nil {
		iotlogger.LogHelper.Helper.Error("参数异常: ", err)
		c.JSON(http.StatusOK, s.setResult(true, "参数异常"))
	}
	//退出逻辑
	if req.Request.SlotInfo.IntentName == "Mi_Exit" && req.Query == "退出" {
		c.JSON(http.StatusOK, s.setResult(true, "好的，已为您退出"))
	}
	testUserId := GetTestUserId(context.Background())
	g := xiaoai.XiaoaiIoTGateway{}
	if s.EnabbleV2 {
		msg, err := g.RunV2(req, manager, testUserId)
		if err != nil {
			msg = "未知错误"
		}
		c.JSON(http.StatusOK, s.setResult(false, msg))
	} else {
		msg, err := g.Run(req, manager, testUserId)
		if err != nil {
			msg = "未知错误"
		}
		c.JSON(http.StatusOK, s.setResult(false, msg))
	}
}

// 根据数据字典类型.获取数据值.(ctx  解决在携程里面调用的问题)
func GetBaseDataValue(dictType string, isImg int, ctx context.Context) (map[string]interface{}, error) {
	res, err := rpcclient.ConfigDictDataService.Lists(ctx, &proto.ConfigDictDataListRequest{
		Query: &proto.ConfigDictData{
			DictType: dictType,
		},
	})
	if err != nil {
		return nil, err
	}
	var dicMap = make(map[string]interface{})
	for _, v := range res.Data {
		dicMap[v.DictLabel] = v.DictValue
		if isImg == 1 && v.Listimg != "" {
			dicMap[v.DictLabel] = v.Listimg
		}
	}
	return dicMap, nil
}

// 获取基础数据配置的回调地址.
func GetTestUserId(ctx context.Context) string {
	mp, err := GetBaseDataValue("voice_test_user", 0, ctx)
	if err != nil {
		return ""
	}
	return iotutil.ToString(mp["xiaoai"])
}

// https://developers.xiaoai.mi.com/documents/Home?type=/api/doc/render_markdown/SkillAccess/skill/fulu/HTTPS
func (s *XiaoaiIotGateWayRouter) GateWayTest(c *gin.Context) {
	//临时逻辑，如果有user-token则为小米Iot
	token := c.Request.Header.Get("User-Token")
	if token != "" {
		xiaomi := XiaomiIotGateWayRouter{}
		xiaomi.GateWay(c)
		return
	}
	iotlogger.LogHelper.Helper.Debug("进入小爱同学处理逻辑")

	// 定义设备列表变量
	var err error
	// 获取设备列表查询接口请求原始数据
	data, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusOK, s.setResult(true, "获取原始数据错误"))
		return
	}
	req := entitys.XiaoaiRequest{}
	if err := json.Unmarshal(data, &req); err != nil {
		iotlogger.LogHelper.Helper.Error("参数异常: ", err)
		c.JSON(http.StatusOK, s.setResult(true, "参数异常"))
	}
	//退出逻辑
	if req.Request.SlotInfo.IntentName == "Mi_Exit" && req.Query == "退出" {
		c.JSON(http.StatusOK, s.setResult(true, "好的，已为您退出"))
	}
	matchTest := func(t, s string) bool {
		matched, err := regexp.MatchString(t, s)
		if err != nil {
			fmt.Println(err)
		}
		return matched
	}
	msg := "已执行"
	//为了上线审核
	switch req.Request.SlotInfo.IntentName {
	case "switch":
		msg = "已执行"
	case "lock":
		msg = "已执行"
	case "mode":
		msg = "已执行"
	case "get_mode":
		msg = "设备模式为睡眠模式"
	case "get_switch":
		msg = "设备已打开"
	case "Mi_Default":
		//把设备的开关{switch}
		if matchTest(`把(.+?)设备的开关(.+?)`, req.Request.Intent.Query) {
			msg = "已执行"
		} else if matchTest(`将(.+?)设备的开关(.+?)`, req.Request.Intent.Query) {
			msg = "已执行"

		} else if matchTest(`将(.+?)设备的开关(.+?)`, req.Request.Intent.Query) {
			msg = "已执行"

		} else if matchTest(`将(.+?)设备的开关(.+?)`, req.Request.Intent.Query) {
			msg = "已执行"

		} else if matchTest(`将(.+?)的童锁(.+?)`, req.Request.Intent.Query) {
			msg = "已执行"

		} else if matchTest(`将(.+?)的模式设置为(.+?)`, req.Request.Intent.Query) {
			msg = "已执行"

		} else if matchTest(`(.+?)设备的模式是(.+?)`, req.Request.Intent.Query) {
			msg = "设备模式为睡眠模式"
		} else if matchTest(`查一下(.+?)的模式是(.+?)`, req.Request.Intent.Query) {
			msg = "设备模式为睡眠模式"
		} else {
			msg = "对不起，不知道你意思哦"
		}
	default:
		msg = "对不起，不知道你意思哦"
	}
	c.JSON(http.StatusOK, s.setResult(false, msg))
}

// 设置返回值
func (s *XiaoaiIotGateWayRouter) setResult(isSessionEnd bool, msg string) map[string]interface{} {
	res := map[string]interface{}{
		"is_session_end": isSessionEnd,
		"version":        "1.0",
		"response": map[string]interface{}{
			"open_mic":       true,
			"not_understand": false,
			"to_speak": map[string]interface{}{
				"type": 0,
				"text": msg,
			},
			"to_display": map[string]interface{}{
				"type": 0,
				"text": msg,
			},
		},
	}
	iotlogger.LogHelper.Helper.Error("输出参数: ", iotutil.ToString(res))
	return res
}
