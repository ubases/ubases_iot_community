package rpcClient

import "cloud_platform/iot_proto/protos/protosService"

// 全局服务客户端变量
var (
	ClientAppMessage  protosService.MpMessageService
	ClientMqttService protosService.MqttService
)
