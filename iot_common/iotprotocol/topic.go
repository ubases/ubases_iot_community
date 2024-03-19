package iotprotocol

import (
	"fmt"
	"strings"
)

// 云端服务请订阅TP_E_开头的主题,即设备发布的主题
// 第1节%s为产品key，云端服务订阅填写通配符“+”，只关注指定产品的消息则可以填产品ID
// 第2节%s为设备ID，在云平台某些地方描述为是DeviceKey，DeviceId非设备表id，请注意这一点。
const (
	TP_C_CONTROL        = "control"        //控制指令
	TP_E_CONTROL_ACK    = "control/ack"    //控制指令响应
	TP_C_QUERY          = "query"          //获取设备状态信息
	TP_E_QUERY_ACK      = "query/ack"      //获取设备状态信息响应
	TP_C_INFO           = "info"           //获取设备配置信息
	TP_E_INFO_REPORT    = "info/report"    //获取设备配置信息响应
	TP_E_REPORT         = "report"         //设备主动上报信息，云端不响应
	TP_E_LOG            = "log"            //设备主动上报日志，云端不响应
	TP_C_UPGRADE        = "upgrade"        //OTA升级信息
	TP_E_UPGRADE_ACK    = "upgrade/ack"    //OTA升级信息响应
	TP_E_UPGRADE_REPORT = "upgrade/report" //OTA升级进度主动上报
	TP_E_ONLINE         = "online"         //在线状态，遗嘱消息、上线消息
	TP_C_CONFIG         = "config"         //下发设备配置
	TP_E_CONFIG_ACK     = "config/ack"     //下发设备配置响应
	TP_E_UTC            = "utc"            //设备请求更新时间
	TP_C_UTC_ACK        = "utc/ack"        //云端下发最新时间
)

var (
	//APP通知topic
	TP_E_NOTICE = "upgrade/notice" //app通知topic
)

func GetVagueTopic(topic string) string {
	return "+/+/" + topic
}

// 去掉产品key和设备Id部分的topic，即返回上边宏定义的内容
func GetTopicSuffix(topic string) string {
	slist := strings.SplitAfterN(topic, "/", 3)
	if len(slist) > 0 {
		return slist[len(slist)-1]
	}
	return topic
}

func ParseTopic(topic string) (string, string, string, error) {
	slist := strings.SplitAfterN(topic, "/", 3)
	if len(slist) >= 3 {
		return strings.ReplaceAll(slist[0], "/", ""), strings.ReplaceAll(slist[1], "/", ""), slist[2], nil
	}
	return "", "", "", fmt.Errorf("topic不规范")
}

func GetTopic(topic, productKey, deviceId string) string {
	return productKey + "/" + deviceId + "/" + topic
}
