package google

import (
	"cloud_platform/iot_voice_service/service/google/proto"
)

// 音量调节
// 参考链接 https://developers.home.google.com/cloud-to-cloud/traits/volume

type MuteCommand func(ctx Context, state bool) proto.DeviceError

func (t MuteCommand) Execute(ctx Context, args map[string]interface{}) proto.CommandResponse {
	res := proto.CommandResponse{
		//Results: map[string]interface{}{},
	}
	if val, ok := args["mute"]; ok {
		if state, ok := val.(bool); ok {
			res.ErrorCode = t(ctx, state)
		} else {
			res.ErrorCode = proto.ErrorCodeNotSupported
		}
	}
	return res
}

func (t MuteCommand) Name() string {
	return proto.ACTION_DEVICES_COMMANDS_MUTE
}

type SetVolumeCommand func(ctx Context, state int) proto.DeviceError

func (t SetVolumeCommand) Execute(ctx Context, args map[string]interface{}) proto.CommandResponse {
	res := proto.CommandResponse{
		//Results: map[string]interface{}{},
	}
	if val, ok := args["volumeLevel"]; ok {
		if state, ok := val.(int); ok {
			res.ErrorCode = t(ctx, state)
		} else {
			res.ErrorCode = proto.ErrorCodeNotSupported
		}
	}
	return res
}

func (t SetVolumeCommand) Name() string {
	return proto.ACTION_DEVICES_COMMANDS_SETVOLUME
}

type VolumeRelativeCommand func(ctx Context, state int) proto.DeviceError

func (t VolumeRelativeCommand) Execute(ctx Context, args map[string]interface{}) proto.CommandResponse {
	res := proto.CommandResponse{}
	//relativeSteps 为“递减”的负数。
	if val, ok := args["relativeSteps"]; ok {
		if state, ok := val.(int); ok {
			res.ErrorCode = t(ctx, state)
		} else {
			res.ErrorCode = proto.ErrorCodeNotSupported
		}
	}
	return res
}

func (t VolumeRelativeCommand) Name() string {
	return proto.ACTION_DEVICES_COMMANDS_VOLUMERELATIVE
}
