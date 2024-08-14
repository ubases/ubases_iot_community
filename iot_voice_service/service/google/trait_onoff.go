package google

import (
	"errors"

	"cloud_platform/iot_voice_service/service/google/proto"
)

//开关

type OnOffTrait struct {
	//指示设备是否只能通过命令控制，且无法查询状态信息。
	CommandOnlyOnOff bool
	//指示设备是否只能查询状态信息，且无法通过命令控制。
	QueryOnlyOnOff  bool
	OnExecuteChange OnOffCommand
	OnStateHandler  func(Context) (bool, proto.ErrorCode)
}

func (t OnOffTrait) ValidateTrait() error {
	if t.OnExecuteChange == nil {
		return errors.New("OnExecuteChange cannot be nil")
	}
	if t.OnStateHandler == nil {
		return errors.New("OnStateHandler cannot be nil")
	}

	return nil
}
func (t OnOffTrait) TraitName() string {
	return proto.ACTION_DEVICES_TRAITS_ONOFF
}

func (t OnOffTrait) TraitStates(ctx Context) []State {
	var onOffState State
	onOffState.Name = "on"
	onOffState.Value, onOffState.Error = t.OnStateHandler(ctx)
	return []State{onOffState}
}

func (t OnOffTrait) TraitCommands() []Command {
	return []Command{t.OnExecuteChange}
}

func (t OnOffTrait) TraitAttributes() []Attribute {
	return []Attribute{
		{
			Name:  "commandOnlyOnOff",
			Value: t.CommandOnlyOnOff,
		},
		{
			Name:  "queryOnlyOnOff",
			Value: t.QueryOnlyOnOff,
		}}
}
