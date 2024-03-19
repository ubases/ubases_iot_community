package google

import (
	"cloud_platform/iot_smart_speaker_service/service/google/proto"
	"errors"
)

// 湿度
type HumiditySetpointRange struct {
	MinPercent float32
	MaxPercent float32
}

type HumidityData struct {
	HumiditySetpointPercent int
	HumidityAmbientPercent  int
}

type HumidityTrait struct {
	HumiditySetpointRange      HumiditySetpointRange
	CommandOnlyHumiditySetting bool
	QueryOnlyHumiditySetting   bool
	OnHumidityChange           SetHumidityCommand
	OnHumidityRelativeCommand  HumidityRelativeCommand
	OnStateHandler             func(Context) (HumidityData, proto.ErrorCode)
}

func (t HumidityTrait) ValidateTrait() error {
	if t.OnHumidityChange == nil {
		return errors.New("OnHumidityChange cannot be nil")
	}
	if t.OnHumidityRelativeCommand == nil {
		return errors.New("OnHumidityRelativeCommand cannot be nil")
	}
	if t.OnStateHandler == nil {
		return errors.New("OnStateHandler cannot be nil")
	}
	return nil
}
func (t HumidityTrait) TraitName() string {
	return proto.ACTION_DEVICES_TRAITS_HUMIDITYSETTING
}

func (t HumidityTrait) TraitStates(ctx Context) []State {
	var humidityState State
	var humidityAmbientState State
	humidityState.Name = "humiditySetpointPercent"
	humidityAmbientState.Name = "humidityAmbientPercent"
	data, err := t.OnStateHandler(ctx)
	if err != nil {
		humidityState.Error = err
		humidityAmbientState.Error = err
	}
	humidityState.Value = data.HumiditySetpointPercent
	humidityAmbientState.Value = data.HumidityAmbientPercent
	return []State{humidityState, humidityAmbientState}
}

func (t HumidityTrait) TraitCommands() []Command {
	return []Command{t.OnHumidityChange, t.OnHumidityRelativeCommand}
}

func (t HumidityTrait) TraitAttributes() []Attribute {
	return []Attribute{
		{
			Name:  "humiditySetpointRange",
			Value: t.HumiditySetpointRange,
		},
		{
			Name:  "commandOnlyHumiditySetting",
			Value: t.CommandOnlyHumiditySetting,
		},
		{
			Name:  "queryOnlyHumiditySetting",
			Value: t.QueryOnlyHumiditySetting,
		},
	}
}
