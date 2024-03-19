package google

import "cloud_platform/iot_smart_speaker_service/service/google/proto"

// 逻辑
// 参考链接 https://developers.home.google.com/cloud-to-cloud/traits/toggles

type TogglesCommand func(ctx Context, state map[string]interface{}) proto.DeviceError

func (t TogglesCommand) Execute(ctx Context, args map[string]interface{}) proto.CommandResponse {
	res := proto.CommandResponse{}
	if val, ok := args["updateToggleSettings"]; ok {
		if state, ok := val.(map[string]interface{}); ok {
			res.ErrorCode = t(ctx, state)
		} else {
			res.ErrorCode = proto.ErrorCodeNotSupported
		}
	}
	return res
}

func (t TogglesCommand) Name() string {
	return proto.ACTION_DEVICES_COMMANDS_SETTOGGLES
}
