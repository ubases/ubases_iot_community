package google

import (
	"errors"

	"cloud_platform/iot_voice_service/service/google/proto"
)

//Toggles

type Toggles struct {
	Name       string       `json:"name"`
	NameValues []NameValues `json:"name_values"`
}

type TogglesTrait struct {
	AvailableToggles []Toggles
	//指示设备是否支持单向 (true) 或双向 (false) 通信。如果设备无法响应此特征的 QUERY intent 或报告状态，请将此属性设为 true。
	CommandOnlyToggles bool
	//如果设备支持仅执行查询，则必须提供。此属性指示设备是否只能查询状态信息，且无法控制。
	QueryOnlyToggles bool
	OnExecuteChange  TogglesCommand
	OnStateHandler   func(Context) (map[string]bool, proto.ErrorCode)
}

func (t TogglesTrait) ValidateTrait() error {
	if t.OnExecuteChange == nil {
		return errors.New("OnExecuteChange cannot be nil")
	}
	if t.OnStateHandler == nil {
		return errors.New("OnStateHandler cannot be nil")
	}

	return nil
}
func (t TogglesTrait) TraitName() string {
	return proto.ACTION_DEVICES_TRAITS_TOGGLES
}

func (t TogglesTrait) TraitStates(ctx Context) []State {
	var state State
	state.Name = "currentToggleSettings"
	state.Value, state.Error = t.OnStateHandler(ctx)
	return []State{state}
}

func (t TogglesTrait) TraitCommands() []Command {
	return []Command{t.OnExecuteChange}
}

func (t TogglesTrait) TraitAttributes() []Attribute {
	return []Attribute{
		{
			Name:  "availableToggles",
			Value: t.AvailableToggles,
		},
		{
			Name:  "commandOnlyToggles",
			Value: t.QueryOnlyToggles,
		},
		{
			Name:  "queryOnlyToggles",
			Value: t.CommandOnlyToggles,
		},
	}
}
