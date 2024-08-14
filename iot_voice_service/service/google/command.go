package google

import "cloud_platform/iot_voice_service/service/google/proto"

// Command Google功能点，Device Traits
// 参考链接
type Command interface {
	//Google功能点名称
	Name() string
	//Google功能点执行接口
	Execute(ctx Context, args map[string]interface{}) proto.CommandResponse
}
