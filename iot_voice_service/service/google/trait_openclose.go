package google

import (
	"errors"

	"cloud_platform/iot_voice_service/service/google/proto"
)

type MultiDirectionOpenCloseTrait struct {
	////当设置为 true 时，表示设备必须完全打开或完全关闭（也就是说，不支持介于 0% 到 100% 之间的值）。
	DiscreteOnlyOpenClose bool
	//设备可打开或关闭的受支持路线列表。如果设备支持多个方向打开和关闭，请添加此属性。
	//支持的方向: UP DOWN LEFT RIGHT IN OUT
	OpenDirection []OpenCloseTraitDirection
	//指示设备是否支持单向 (true) 或双向 (false) 通信。如果设备无法响应此特征的 QUERY intent 或报告状态，请将此属性设为 true。
	CommandOnlyOpenClose bool
	//指示设备是否支持单向 (true) 或双向 (false) 通信。如果设备无法响应此特征的 QUERY intent 或报告状态，请将此属性设为 true。
	QueryOnlyOpenClose bool
	OnExecuteChange    DirectionalOpenCloseCommand
	OnStateHandler     func(Context) ([]OpenState, proto.ErrorCode)
}

type OpenCloseTraitDirection string

const OpenCloseTraitDirectionNone OpenCloseTraitDirection = ""
const OpenCloseTraitDirectionUp OpenCloseTraitDirection = "UP"
const OpenCloseTraitDirectionDown OpenCloseTraitDirection = "DOWN"
const OpenCloseTraitDirectionLeft OpenCloseTraitDirection = "LEFT"
const OpenCloseTraitDirectionRight OpenCloseTraitDirection = "RIGHT"

func (t MultiDirectionOpenCloseTrait) ValidateTrait() error {
	if t.OnExecuteChange == nil {
		return errors.New("OnExecuteChange cannot be nil")
	}
	if t.OnStateHandler == nil {
		return errors.New("OnStateHandlers cannot be nil")
	}
	return nil
}

func (t MultiDirectionOpenCloseTrait) TraitName() string {
	return proto.ACTION_DEVICES_TRAITS_OPENCLOSE
}

type OpenState struct {
	OpenPercent   float64
	OpenDirection OpenCloseTraitDirection
}

func (t MultiDirectionOpenCloseTrait) TraitStates(ctx Context) []State {
	onOffState := State{
		Name:  "on",
		Value: true,
	}
	handlerOpenState, err := t.OnStateHandler(ctx)
	curOpenState := State{
		Name:  "openState",
		Value: handlerOpenState,
		Error: err,
	}
	return []State{onOffState, curOpenState}
}

func (t MultiDirectionOpenCloseTrait) TraitCommands() []Command {
	return []Command{t.OnExecuteChange}
}

func (t MultiDirectionOpenCloseTrait) TraitAttributes() []Attribute {
	atr := []Attribute{
		{
			Name:  "discreteOnlyOpenClose",
			Value: t.DiscreteOnlyOpenClose,
		},
		{
			Name:  "queryOnlyOpenClose",
			Value: t.QueryOnlyOpenClose,
		},
	}

	if len(t.OpenDirection) > 0 {
		openDirectionArg := Attribute{
			Name:  "openDirection",
			Value: t.OpenDirection,
		}
		atr = append(atr, openDirectionArg)
	}

	return atr
}

type OpenCloseTrait struct {
	DiscreteOnlyOpenClose bool
	QueryOnlyOpenClose    bool
	OnExecuteChange       OpenCloseCommand
	OnStateHandler        func(Context) (float64, proto.ErrorCode)
}

func (t OpenCloseTrait) ValidateTrait() error {
	if t.OnExecuteChange == nil {
		return errors.New("OnExecuteChange cannot be nil")
	}
	if t.OnStateHandler == nil {
		return errors.New("OnStateHandlers cannot be nil")
	}
	return nil
}

func (t OpenCloseTrait) TraitName() string {
	return proto.ACTION_DEVICES_TRAITS_OPENCLOSE
}

func (t OpenCloseTrait) TraitStates(ctx Context) []State {
	onOffState := State{
		Name:  "on",
		Value: true,
	}
	handlerOpenState, err := t.OnStateHandler(ctx)
	curOpenState := State{
		Name:  "openState",
		Value: handlerOpenState,
		Error: err,
	}
	return []State{onOffState, curOpenState}
}

func (t OpenCloseTrait) TraitCommands() []Command {
	return []Command{t.OnExecuteChange}
}

func (t OpenCloseTrait) TraitAttributes() []Attribute {
	atr := []Attribute{
		{
			Name:  "discreteOnlyOpenClose",
			Value: t.DiscreteOnlyOpenClose,
		},
		{
			Name:  "queryOnlyOpenClose",
			Value: t.QueryOnlyOpenClose,
		},
	}

	return atr
}
