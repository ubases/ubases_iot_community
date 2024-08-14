package google

import "cloud_platform/iot_voice_service/service/google/proto"

// 湿度设置
type SetHumidityCommand func(ctx Context, value int) proto.DeviceError

func (t SetHumidityCommand) Execute(ctx Context, args map[string]interface{}) proto.CommandResponse {
	res := proto.CommandResponse{}
	if val, ok := args["humidity"]; ok {
		if state, ok := val.(float64); ok {
			res.ErrorCode = t(ctx, int(state))
		} else {
			res.ErrorCode = proto.ErrorCodeNotSupported
		}
	}
	return res
}

func (t SetHumidityCommand) Name() string {
	return proto.ACTION_DEVICES_COMMANDS_SETHUMIDITY
}

type HumidityRelativeCommand func(ctx Context, value int) proto.DeviceError

func (t HumidityRelativeCommand) Execute(ctx Context, args map[string]interface{}) proto.CommandResponse {
	res := proto.CommandResponse{}
	if val, ok := args["humidityRelativePercent"]; ok {
		//要更改的亮度的确切百分比。
		if state, ok := val.(int); ok {
			res.ErrorCode = t(ctx, state)
		} else {
			res.ErrorCode = proto.ErrorCodeNotSupported
		}
	} else if val, ok := args["humidityRelativeWeight"]; ok {
		//这表示明暗的亮度变化量。从小容量到大容量，此参数会缩放到整数 0 到 5，其中符号表示方向。
		if state, ok := val.(int); ok {
			//[0,5]转成百分比，统一按照百分比处理
			state = state * (100 / 5)
			res.ErrorCode = t(ctx, state)
		} else {
			res.ErrorCode = proto.ErrorCodeNotSupported
		}
	}
	return res
}

func (t HumidityRelativeCommand) Name() string {
	return proto.ACTION_DEVICES_COMMANDS_HUMIDITYRELATIVE
}
