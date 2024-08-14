package google

import (
	"errors"

	"cloud_platform/iot_voice_service/service/google/proto"
)

//定时器

type TimerData struct {
	//当前剩余时间（以秒为单位），-1 或 [0, maxTimerLimitSec]。设置为 -1 表示没有任何计时器正在运行。
	TimerRemainingSec int64
	//如果存在有效的计时器，但当前已暂停，则为 true。
	TimerPaused bool
}

type TimerTrait struct {
	//指明设备上可用的最长计时器设置（以秒为单位）。
	MaxTimerLimitSec int64
	//指示设备是否支持单向 (true) 或双向 (false) 通信。如果设备无法响应此特征的 QUERY intent 或报告状态，请将此属性设为 true。
	CommandOnlyTimer      bool
	OnExecuteChangeStart  TimerStartCommand
	OnExecuteChangeAdjust TimerAdjustCommand
	OnExecuteChangePause  TimerPauseCommand
	OnExecuteChangeResume TimerResumeCommand
	OnExecuteChangeCancel TimerCancelCommand
	OnStateHandler        func(Context) (TimerData, proto.ErrorCode)
}

func (t TimerTrait) ValidateTrait() error {
	if t.OnExecuteChangeStart == nil {
		return errors.New("OnExecuteChangeStart cannot be nil")
	}
	if t.OnExecuteChangeAdjust == nil {
		return errors.New("OnExecuteChangeAdjust cannot be nil")
	}
	if t.OnExecuteChangePause == nil {
		return errors.New("OnExecuteChangePause cannot be nil")
	}
	if t.OnExecuteChangeResume == nil {
		return errors.New("OnExecuteChangeResume cannot be nil")
	}
	if t.OnExecuteChangeCancel == nil {
		return errors.New("OnExecuteChangeCancel cannot be nil")
	}
	if t.OnStateHandler == nil {
		return errors.New("OnStateHandler cannot be nil")
	}

	return nil
}
func (t TimerTrait) TraitName() string {
	return proto.ACTION_DEVICES_TRAITS_TIMER
}

func (t TimerTrait) TraitStates(ctx Context) []State {
	data, err := t.OnStateHandler(ctx)
	var timerState State
	var timerState2 State
	timerState.Name = "timerRemainingSec"
	timerState.Value = data.TimerRemainingSec
	timerState.Error = err
	timerState2.Name = "timerPaused"
	timerState2.Value = data.TimerPaused
	timerState.Error = err
	return []State{timerState, timerState2}
}

func (t TimerTrait) TraitCommands() []Command {
	return []Command{t.OnExecuteChangeStart, t.OnExecuteChangeAdjust, t.OnExecuteChangeResume, t.OnExecuteChangePause, t.OnExecuteChangeCancel}
}

func (t TimerTrait) TraitAttributes() []Attribute {
	return []Attribute{
		{
			Name:  "maxTimerLimitSec",
			Value: t.MaxTimerLimitSec,
		},
		{
			Name:  "commandOnlyTimer",
			Value: t.CommandOnlyTimer,
		}}
}
