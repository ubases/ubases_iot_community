package google

import (
	"cloud_platform/iot_voice_service/service/google/proto"
)

// 亮度调节
// 参考链接 https://developers.google.com/assistant/smarthome/traits/brightness.html

//示例话语:
//brighten light at 10%
//brighten lights by 20%
//brighten the kitchen slightly
//brighten the kitchen to 50
//brighten the lights up
//dim down the lights
//dim lights a little more
//dim lights by 20%
//how bright are my lights
//increase brightness in living room by 100%
//make it darker in the living room
//make the bedroom brighter
//make the lights a little bit brighter
//make the living room a little bit dimmer
//reduce brightness in the living room by 50%
//set the light to maximum
//turn brightness on light to maximum

// 亮度绝对值，百分比
type BrightnessAbsoluteCommand func(ctx Context, value int) proto.DeviceError

func (t BrightnessAbsoluteCommand) Execute(ctx Context, args map[string]interface{}) proto.CommandResponse {
	res := proto.CommandResponse{}
	if val, ok := args["brightness"]; ok {
		if state, ok := val.(float64); ok {
			res.ErrorCode = t(ctx, int(state))
		} else {
			res.ErrorCode = proto.ErrorCodeNotSupported
		}
	}
	return res
}

func (t BrightnessAbsoluteCommand) Name() string {
	return proto.ACTION_DEVICES_COMMANDS_BRIGHTNESSABSOLUTE
}

type BrightnessRelativeCommand func(ctx Context, value int) proto.DeviceError

func (t BrightnessRelativeCommand) Execute(ctx Context, args map[string]interface{}) proto.CommandResponse {
	res := proto.CommandResponse{}
	if val, ok := args["brightnessRelativePercent"]; ok {
		//要更改的亮度的确切百分比。
		if state, ok := val.(float64); ok {
			res.ErrorCode = t(ctx, int(state))
		} else {
			res.ErrorCode = proto.ErrorCodeNotSupported
		}
	} else if val, ok := args["brightnessRelativeWeight"]; ok {
		//这表示明暗的亮度变化量。从小容量到大容量，此参数会缩放到整数 0 到 5，其中符号表示方向。
		if state, ok := val.(float64); ok {
			//[0,5]转成百分比，统一按照百分比处理
			state = state * (100 / 5)
			res.ErrorCode = t(ctx, int(state))
		} else {
			res.ErrorCode = proto.ErrorCodeNotSupported
		}
	}
	return res
}

func (t BrightnessRelativeCommand) Name() string {
	return proto.ACTION_DEVICES_COMMANDS_BRIGHTNESSRELATIVE
}
