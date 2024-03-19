package rpcClient

import (
	"cloud_platform/iot_proto/protos/protosService"
)

// 基础管理
var (
	ClientConfigDictDataServerService protosService.ConfigDictDataService
	ClientConfigDictTypeServerService protosService.ConfigDictTypeService
)

// 语言包
var (
	ClientLangCustomResourceService protosService.LangCustomResourcesService
	ClientLangResourcesService      protosService.LangResourcesService
)
