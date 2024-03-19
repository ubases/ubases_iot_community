package google

import (
	"cloud_platform/iot_smart_speaker_service/service/google/proto"
	"errors"
)

//模式，适用于任何枚举的功能点

type NameValues struct {
	NameSynonym []string `json:"name_synonym"`
	Lang        string   `json:"lang"`
}
type SettingValues struct {
	SettingSynonym []string `json:"setting_synonym"`
	Lang           string   `json:"lang"`
}
type Settings struct {
	SettingName   string          `json:"setting_name"`
	SettingValues []SettingValues `json:"setting_values"`
}
type Modes struct {
	Name       string       `json:"name"`
	NameValues []NameValues `json:"name_values"`
	Settings   []Settings   `json:"settings"`
	Ordered    bool         `json:"ordered"`
}

type ModesTrait struct {
	//可用模式的列表。
	AvailableModes []Modes
	//指示设备是否支持单向 (true) 或双向 (false) 通信。如果设备无法响应此特征的 QUERY intent 或报告状态，请将此属性设为 true。
	CommandOnlyModes bool
	//如果设备支持仅执行查询，则必须提供。此属性指示设备是否只能查询状态信息，且无法控制。
	QueryOnlyModes  bool
	OnExecuteChange ModesCommand
	OnStateHandler  func(Context) (map[string]interface{}, proto.ErrorCode)
}

func (t ModesTrait) ValidateTrait() error {
	if t.OnExecuteChange == nil {
		return errors.New("OnExecuteChange cannot be nil")
	}
	if t.OnStateHandler == nil {
		return errors.New("OnStateHandler cannot be nil")
	}
	return nil
}
func (t ModesTrait) TraitName() string {
	return proto.ACTION_DEVICES_TRAITS_MODES
}

func (t ModesTrait) TraitStates(ctx Context) []State {
	var state State
	state.Name = "currentModeSettings"
	state.Value, state.Error = t.OnStateHandler(ctx)
	return []State{state}
}

func (t ModesTrait) TraitCommands() []Command {
	return []Command{t.OnExecuteChange}
}

func (t ModesTrait) TraitAttributes() []Attribute {
	return []Attribute{
		{
			Name:  "availableModes",
			Value: t.AvailableModes,
		},
	}
}
