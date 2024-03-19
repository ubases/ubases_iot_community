package rpcClient

import "cloud_platform/iot_proto/protos/protosService"

// 全局服务客户端变量
var (
	ClientUpgradeRecordService        protosService.IotOtaUpgradeRecordService
	ClientDeviceLogService            protosService.IotDeviceLogService
	ClientLvglService                 protosService.LvglService
	ClientLangResourcesPackageService protosService.LangResourcePackageService
)
