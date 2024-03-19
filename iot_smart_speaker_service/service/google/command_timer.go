package google

import (
	"cloud_platform/iot_smart_speaker_service/service/google/proto"
)

// 定时器
// 参考链接 https://developers.home.google.com/cloud-to-cloud/traits/timer

type TimerStartCommand func(ctx Context, state float64) proto.DeviceError

func (t TimerStartCommand) Execute(ctx Context, args map[string]interface{}) proto.CommandResponse {
	res := proto.CommandResponse{}
	//计时器的时长（以秒为单位）；必须在 [1, maxTimerLimitSec] 之内。
	if val, ok := args["timerTimeSec"]; ok {
		if state, ok := val.(float64); ok {
			res.ErrorCode = t(ctx, state)
		} else {
			res.ErrorCode = proto.ErrorCodeNotSupported
		}
	}
	return res
}

func (t TimerStartCommand) Name() string {
	return proto.ACTION_DEVICES_COMMANDS_TIMERSTART
}

type TimerAdjustCommand func(ctx Context, state float64) proto.DeviceError

func (t TimerAdjustCommand) Execute(ctx Context, args map[string]interface{}) proto.CommandResponse {
	res := proto.CommandResponse{}
	//计时器的正负调整（以秒为单位）；必须在 [-maxTimerLimitSec、maxTimerLimitSec] 之内。
	if val, ok := args["timerTimeSec"]; ok {
		if state, ok := val.(float64); ok {
			res.ErrorCode = t(ctx, state)
		} else {
			res.ErrorCode = proto.ErrorCodeNotSupported
		}
	}
	return res
}

func (t TimerAdjustCommand) Name() string {
	return proto.ACTION_DEVICES_COMMANDS_TIMERADJUST
}

type TimerPauseCommand func(ctx Context, state string) proto.DeviceError

func (t TimerPauseCommand) Execute(ctx Context, args map[string]interface{}) proto.CommandResponse {
	res := proto.CommandResponse{}
	res.ErrorCode = t(ctx, "Pause")
	return res
}

func (t TimerPauseCommand) Name() string {
	return proto.ACTION_DEVICES_COMMANDS_TIMERPAUSE
}

type TimerResumeCommand func(ctx Context, state string) proto.DeviceError

func (t TimerResumeCommand) Execute(ctx Context, args map[string]interface{}) proto.CommandResponse {
	res := proto.CommandResponse{}
	res.ErrorCode = t(ctx, "Resume")
	return res
}

func (t TimerResumeCommand) Name() string {
	return proto.ACTION_DEVICES_COMMANDS_TIMERRESUME
}

type TimerCancelCommand func(ctx Context, state string) proto.DeviceError

func (t TimerCancelCommand) Execute(ctx Context, args map[string]interface{}) proto.CommandResponse {
	res := proto.CommandResponse{}
	res.ErrorCode = t(ctx, "Cancel")
	return res
}

func (t TimerCancelCommand) Name() string {
	return proto.ACTION_DEVICES_COMMANDS_TIMERCANCEL
}
