package google

import (
	"cloud_platform/iot_smart_speaker_service/service/google/proto"
)

// 开关调节
// 参考链接 https://developers.google.com/assistant/smarthome/traits/onoff.html

//示例话语:
//are the lights off
//turn off the AC
//turn on my lights
//what is on in the kitchen ?

type OnOffCommand func(ctx Context, state bool) proto.DeviceError

func (t OnOffCommand) Execute(ctx Context, args map[string]interface{}) proto.CommandResponse {
	res := proto.CommandResponse{}
	if val, ok := args["on"]; ok {
		if state, ok := val.(bool); ok {
			res.ErrorCode = t(ctx, state)
		} else {
			res.ErrorCode = proto.ErrorCodeNotSupported
		}
	}
	return res
}

func (t OnOffCommand) Name() string {
	return proto.ACTION_DEVICES_COMMANDS_ONOFF
}
