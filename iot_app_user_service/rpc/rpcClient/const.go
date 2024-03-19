package rpcClient

import "cloud_platform/iot_proto/protos/protosService"

// 全局服务客户端变量
var (
	IotDeviceHomeService        protosService.IotDeviceHomeService
	OpmProductService           protosService.OpmProductService
	DictDataService             protosService.ConfigDictDataService
	ClientOemAppDocDirService   protosService.OemAppDocDirService
	ClientOemAppService         protosService.OemAppService
	ClientOemAppUiConfigService protosService.OemAppUiConfigService
	ClientSysAreaService        protosService.SysAreaService
	IPService                   protosService.IPService
)
