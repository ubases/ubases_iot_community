package google

import (
	"cloud_platform/iot_voice_service/service/google/proto"
)

// 打开\关闭
// 参考链接 https://developers.home.google.com/cloud-to-cloud/traits/openclose

//示例话语:
//Are the blinds in the kitchen open
//Close the blinds 25% in my room
//Close the door more
//Just close the front door
//open the blinds in my room
//open the blinds to 25%
//open the door by 25%

type OpenCloseCommand func(ctx Context, openPercent float64) proto.DeviceError

func (t OpenCloseCommand) Execute(ctx Context, args map[string]interface{}) proto.CommandResponse {
	res := proto.CommandResponse{
		ErrorCode: proto.ErrorCodeProtocolError,
	}
	if argOpenPercent, ok := args["openPercent"]; ok {
		if openPercent, ok := argOpenPercent.(float64); ok {
			res.ErrorCode = t(ctx, openPercent)
		} else {
			res.ErrorCode = proto.ErrorCodeNotSupported
		}
	}
	return res
}

func (t OpenCloseCommand) Name() string {
	return proto.ACTION_DEVICES_COMMANDS_OPENCLOSE
}

type DirectionalOpenCloseCommand func(ctx Context, openPercent float64, openDirection OpenCloseTraitDirection) proto.DeviceError

func (t DirectionalOpenCloseCommand) Execute(ctx Context, args map[string]interface{}) proto.CommandResponse {
	res := proto.CommandResponse{
		ErrorCode: proto.ErrorCodeProtocolError,
	}
	openDirection := OpenCloseTraitDirectionNone
	//openPercent 表示设备已打开的百分比，其中 0 表示关闭，100 则表示完全打开。
	if argOpenPercent, ok := args["openPercent"]; ok {
		if openPercent, ok := argOpenPercent.(float64); ok {
			//openDirection 设备的打开方向。
			if argDir, ok := args["openDirection"]; ok {
				if dir, ok := argDir.(string); ok {
					openDirection = OpenCloseTraitDirection(dir)
				}
			}
			res.ErrorCode = t(ctx, openPercent, openDirection)
		}
		res.ErrorCode = proto.ErrorCodeNotSupported
	}
	return res
}

func (t DirectionalOpenCloseCommand) Name() string {
	return proto.ACTION_DEVICES_COMMANDS_OPENCLOSE
}
