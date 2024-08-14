package service

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_voice_service/service/common"
	"cloud_platform/iot_voice_service/service/google"
	"cloud_platform/iot_voice_service/service/google/proto"
	"strconv"
)

const DT_EMUN = "ENUM"
const DT_BOOL = "BOOL"
const DT_INT = "INT"
const DT_DOUBLE = "DOUBLE"
const DT_TEXT = "TEXT"
const DT_FAULT = "FAULT"
const DT_FLOAT = "FLOAT"

func GetInfo(ctx google.Context, code string) (*google.CommandInfo, string, error) {
	mapCmd := ctx.Target.DeviceCommand()
	if mapCmd == nil {
		return nil, "", proto.DeviceErrorNotSupported
	}
	cmdinfo, ok := mapCmd[code]
	if !ok {
		return nil, "", proto.DeviceErrorNotSupported
	}
	mapCustom := ctx.Target.DeviceCustomData()
	if mapCustom == nil {
		return nil, "", proto.DeviceErrorNotSupported
	}
	productKey, ok := mapCustom[proto.CUSTOMDATA_PRODUCTKEY]
	if !ok {
		return nil, "", proto.DeviceErrorNotSupported
	}
	return cmdinfo, productKey.(string), nil
}

func GetDataInfo(ctx google.Context, code string, dataType string) (interface{}, error) {
	mapCmd := ctx.Target.DeviceCommand()
	if mapCmd == nil {
		return "", proto.DeviceErrorNotSupported
	}
	cmdinfo, ok := mapCmd[code]
	if !ok {
		return "", proto.DeviceErrorNotSupported
	}
	mapData, err := common.GetDeviceInfo(ctx.Target.DeviceId())
	if err != nil {
		return "", err
	}
	if mapData == nil {
		return "", proto.DeviceErrorNotSupported
	}
	value, ok := mapData[strconv.Itoa(int(cmdinfo.Dpid))]
	if !ok {
		return "", proto.DeviceErrorNotSupported
	}
	ret := common.ConvertValueType(dataType, value)
	if ret == nil {
		return "", proto.DeviceErrorNotSupported
	}
	return ret, nil
}

func SendOneCommand(ctx google.Context, key string, state interface{}) proto.DeviceError {
	cmd, pk, err := GetInfo(ctx, key)
	if err != nil {
		return err
	}
	data := common.ConvertValueType(cmd.DataType, state)
	if data == nil {
		return nil
	}
	cmdParas := make(map[string]interface{})
	cmdParas[strconv.Itoa(int(cmd.Dpid))] = data
	_, err = common.PubControl(pk, ctx.Target.DeviceId(), cmdParas)
	return err
}

func OnOffTrait_Command(ctx google.Context, state bool) proto.DeviceError {
	return SendOneCommand(ctx, "on", state)
}

func OnOffTrait_State(ctx google.Context) (bool, proto.ErrorCode) {
	ret, err := GetDataInfo(ctx, "on", DT_BOOL)
	if err != nil {
		iotlogger.LogHelper.Errorf("OnOffTrait_State id = %s,error = %s", ctx.Target.DeviceId(), err.Error())
		return false, nil
	}
	return ret.(bool), nil
}

func ModesTrait_Command(ctx google.Context, state map[string]interface{}) proto.DeviceError {
	cmdParas := make(map[string]interface{})
	var productKey string
	for k, v := range state {
		cmd, pk, err := GetInfo(ctx, k)
		if err != nil {
			continue
		}

		var data interface{}
		if cmd.MapValue != nil { //语音平台的枚举->爱星云平台的枚举
			val, ok := cmd.MapValue[v.(string)]
			if ok {
				data = common.ConvertValueType(cmd.DataType, val)
			}
		} else if cmd.MapNumberRange != nil { //语音平台的枚举->爱星云平台的数值范围
			valRange, ok := cmd.MapNumberRange[v.(string)]
			if ok {
				//约定下发该范围内的最大值
				data = common.ConvertValueType(cmd.DataType, valRange.Max)
			}
		}
		if data == nil {
			continue
		}
		productKey = pk
		cmdParas[strconv.Itoa(int(cmd.Dpid))] = data
	}
	_, err := common.PubControl(productKey, ctx.Target.DeviceId(), cmdParas)
	return err
}

func ModesTrait_State(ctx google.Context) (map[string]interface{}, proto.ErrorCode) {
	//先找到设备的所有属性
	mapData, err := common.GetDeviceInfo(ctx.Target.DeviceId())
	if err != nil {
		iotlogger.LogHelper.Errorf("ModesTrait_State id = %s,error = %s", ctx.Target.DeviceId(), err.Error())
	}
	//再过滤mode模式
	ret := make(map[string]interface{})
	for k, v := range ctx.Target.DeviceCommand() {
		if v.TraitName == proto.ACTION_DEVICES_TRAITS_MODES {
			find := false
			if mapData != nil {
				if val, ok := mapData[strconv.Itoa(int(v.Dpid))]; ok {
					nVal, err := strconv.Atoi(val)
					if err == nil {
						if v.MapValue != nil {
							for kk, vv := range v.MapValue {
								if vv == nVal {
									ret[k] = kk
									find = true
									break
								}
							}
						} else if v.MapNumberRange != nil {
							for kk, vv := range v.MapNumberRange {
								if int(vv.Min) <= nVal && nVal <= int(vv.Max) {
									ret[k] = kk
									find = true
									break
								}
							}
						}
					}
				}
			}
			if !find {
				ret[k] = v.Default
			}
		}
	}
	return ret, nil
}

// start、adjust
func TimerTrait_CommandValue(ctx google.Context, state float64) proto.DeviceError {
	minites := state / 60
	return SendOneCommand(ctx, "timerTimeSec", minites)
}

// Pause、Resume、Cancel
func TimerTrait_Command(ctx google.Context, state string) proto.DeviceError {
	//fixme 目前没有暂停、重启，只有取消
	return SendOneCommand(ctx, "timerTimeSec", 0)
}

func TimerTrait_State(ctx google.Context) (google.TimerData, proto.ErrorCode) {
	ret, err := GetDataInfo(ctx, "timerTimeSec", DT_INT)
	if err != nil {
		iotlogger.LogHelper.Errorf("TimerTrait_State id = %s,error = %s", ctx.Target.DeviceId(), err.Error())
		return google.TimerData{TimerRemainingSec: -1, TimerPaused: false}, nil
	}
	sec := ret.(int64)
	if sec > 0 {
		return google.TimerData{TimerRemainingSec: sec, TimerPaused: false}, nil
	}
	return google.TimerData{TimerRemainingSec: -1, TimerPaused: false}, nil
}

func OpenClose_Command(ctx google.Context, openPercent float64) proto.DeviceError {
	//fixme 只实现必填部分，其余用2个到再补充。
	return SendOneCommand(ctx, "openPercent", openPercent)
}

func OpenClose_State(ctx google.Context) (float64, proto.ErrorCode) {
	ret, err := GetDataInfo(ctx, "openPercent", DT_DOUBLE)
	if err != nil {
		return 0.0, nil
	}
	return ret.(float64), nil
}

func VolumeTrait_CommandMute(ctx google.Context, state bool) proto.DeviceError {
	return SendOneCommand(ctx, "mute", state)
}

func VolumeTrait_CommandSetVolume(ctx google.Context, state int) proto.DeviceError {
	return SendOneCommand(ctx, "volumeLevel", state)
}

func VolumeTrait_CommandVolumeRelative(ctx google.Context, state int) proto.DeviceError {
	return SendOneCommand(ctx, "relativeSteps", state)
}

func VolumeTrait_State(ctx google.Context) (int, proto.ErrorCode) {
	ret, err := GetDataInfo(ctx, "currentVolume", DT_INT)
	if err != nil {
		return 0, nil
	}
	return ret.(int), nil
}

func BrightnessTrait_BrightnessAbsoluteCommand(ctx google.Context, value int) proto.DeviceError {
	return SendOneCommand(ctx, "brightness", value)
}
func BrightnessTrait_BrightnessRelativeCommand(ctx google.Context, value int) proto.DeviceError {
	return SendOneCommand(ctx, "brightness", value)
}
func BrightnessTrait_State(ctx google.Context) (int, proto.ErrorCode) {
	ret, err := GetDataInfo(ctx, "brightness", DT_INT)
	if err != nil {
		return 0, nil
	}
	return ret.(int), nil
}

func HumidityTrait_Command(ctx google.Context, value int) proto.DeviceError {
	return SendOneCommand(ctx, "humidity", value)
}

func HumidityTrait_HumidityRelativeCommand(ctx google.Context, value int) proto.DeviceError {
	//fixme 百分比，待处理
	return SendOneCommand(ctx, "humidity", value)
}

func HumidityTrait_State(ctx google.Context) (google.HumidityData, proto.ErrorCode) {
	ret, err := GetDataInfo(ctx, "humidityRelativePercent", DT_INT)
	if err != nil {
		iotlogger.LogHelper.Errorf("HumidityTrait_State id = %s,error = %s", ctx.Target.DeviceId(), err.Error())
		ret = 0
	}
	ret1, err1 := GetDataInfo(ctx, "humidityRelativeWeight", DT_INT)
	if err1 != nil {
		iotlogger.LogHelper.Errorf("HumidityTrait_State id = %s,error = %s", ctx.Target.DeviceId(), err.Error())
		ret1 = 0
	}
	return google.HumidityData{HumiditySetpointPercent: ret.(int), HumidityAmbientPercent: ret1.(int)}, nil
}

func TogglesTrait_Command(ctx google.Context, state map[string]interface{}) proto.DeviceError {
	cmdParas := make(map[string]interface{})
	var productKey string
	for k, v := range state {
		cmd, pk, err := GetInfo(ctx, k)
		if err != nil {
			continue
		}
		data := common.ConvertValueType(cmd.DataType, v)
		if data == nil {
			continue
		}
		productKey = pk
		cmdParas[strconv.Itoa(int(cmd.Dpid))] = data
	}
	_, err := common.PubControl(productKey, ctx.Target.DeviceId(), cmdParas)
	return err
}

func TogglesTrait_State(ctx google.Context) (map[string]bool, proto.ErrorCode) {
	//先找到设备的所有属性
	mapData, err := common.GetDeviceInfo(ctx.Target.DeviceId())
	if err != nil {
		iotlogger.LogHelper.Errorf("TogglesTrait_State id = %s,error = %s", ctx.Target.DeviceId(), err.Error())
	}
	//再过滤mode模式
	ret := make(map[string]bool)
	for k, v := range ctx.Target.DeviceCommand() {
		if v.TraitName == proto.ACTION_DEVICES_TRAITS_TOGGLES {
			find := false
			if mapData != nil {
				if val, ok := mapData[strconv.Itoa(int(v.Dpid))]; ok {
					bVal := common.ConvertValueType(DT_BOOL, val)
					if bVal != nil {
						ret[k] = bVal.(bool)
						find = true
					}
				}
			}
			if !find {
				ret[k] = false
			}
		}
	}
	return ret, nil
}
