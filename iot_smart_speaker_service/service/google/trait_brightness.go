package google

import (
	"errors"

	"cloud_platform/iot_smart_speaker_service/service/google/proto"
)

//亮度

type BrightnessTrait struct {
	CommandOnlyBrightness      bool
	OnBrightnessChange         BrightnessAbsoluteCommand
	OnBrightnessRelativeChange BrightnessRelativeCommand
	OnStateHandler             func(Context) (int, proto.ErrorCode)
}

func (t BrightnessTrait) ValidateTrait() error {
	if t.OnBrightnessChange == nil {
		return errors.New("OnBrightnessChange cannot be nil")
	}
	if t.OnBrightnessRelativeChange == nil {
		return errors.New("OnBrightnessRelativeChange cannot be nil")
	}
	if t.OnStateHandler == nil {
		return errors.New("OnStateHandler cannot be nil")
	}

	return nil
}
func (t BrightnessTrait) TraitName() string {
	return proto.ACTION_DEVICES_TRAITS_BRIGHTNESS
}

func (t BrightnessTrait) TraitStates(ctx Context) []State {
	var state State
	state.Name = "brightness"
	state.Value, state.Error = t.OnStateHandler(ctx)
	return []State{state}
}

func (t BrightnessTrait) TraitCommands() []Command {
	return []Command{t.OnBrightnessChange, t.OnBrightnessRelativeChange}
}

func (t BrightnessTrait) TraitAttributes() []Attribute {
	return []Attribute{
		{
			Name:  "commandOnlyBrightness",
			Value: t.CommandOnlyBrightness,
		},
	}
}
