package common

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_smart_speaker_service/cached"
	"context"
	"fmt"
	"math"
)

func ConvertToAxyNumberRange(devId string, adjust bool, dpid int32, params map[string]interface{}, val interface{}) interface{} {
	var v float64
	switch val.(type) {
	case float64:
		v = val.(float64)
	default:
		v2, err := iotutil.ToFloat64Err(val)
		if err != nil {
			return nil
		}
		v = v2
	}
	var value interface{}
	//min := GetMaxValue(params["min"].(float64), params["vMin"].(float64))
	//max := GetMinValue(params["max"].(float64), params["vMax"].(float64))
	step := params["step"].(float64)

	var aMultiple float64 = 1
	if v, ok := params["multiple"]; ok {
		aMultiple, _ = iotutil.ToFloat64Err(v)
		//兼容逻辑，倍数不能为0
		if aMultiple == 0 {
			aMultiple = 1
		}
	} else {
		//TODO 兼容逻辑，后期根据需求去除
		if step < 1 && step != 0 {
			aMultiple = 1 / step
		}
	}

	//数值转换
	vMin := params["vMin"].(float64)
	vMax := params["vMax"].(float64)
	aMax := params["max"].(float64)
	aMin := params["min"].(float64)

	r := (v - vMin) / (vMax - vMin)
	aV := aMin + math.Round(((aMax-aMin)*r)*10)/10

	if adjust {
		vStr := cached.RedisStore.GetClient().HGet(context.Background(), fmt.Sprintf("dev_data_%s", devId), iotutil.ToString(dpid)).Val()
		iotlogger.LogHelper.Helper.Debug("vStr: ", vStr, v, dpid)
		if vStr == "" {
			vStr = "0"
		}
		aV += iotutil.ToFloat64(vStr)
	}

	//乘以倍数
	if aV > aMax {
		value = aMax * aMultiple
	} else if aV < aMin {
		value = aMin * aMultiple
	} else {
		value = aV * aMultiple
	}
	return value
}

func ConvertToVoiceNumberRange(devId string, adjust bool, dpid int32, params map[string]interface{}, val interface{}) interface{} {
	var v float64
	switch val.(type) {
	case float64:
		v = val.(float64)
	default:
		v2, err := iotutil.ToFloat64Err(val)
		if err != nil {
			return nil
		}
		v = v2
	}
	var value interface{}
	min := GetMaxValue(params["min"].(float64), params["vMin"].(float64))
	max := GetMinValue(params["max"].(float64), params["vMax"].(float64))

	// 判断爱星云单位和语控单位是否存在，如果存在且不为空，则需要进行单位换算
	unit, okUnit := params["unit"].(string)
	vUnit, okvUnit := params["vUnit"].(string)
	if (okUnit && okvUnit) && (len(unit) != 0 && len(vUnit) != 0) && (unit != vUnit) {
		base, okbase := iotconst.VoiceUnitConvert[unit]
		vBase, okvBase := iotconst.VoiceUnitConvert[vUnit]
		if okbase && okvBase {
			v = v * vBase / base
		}
	}

	if v > max {
		value = max
	} else if v < min {
		value = min
	} else {
		value = v
	}
	return value
}

func ConvertValueType(dataType string, value interface{}) interface{} {
	var v interface{}
	var err error
	switch dataType {
	case "ENUM":
		v, err = iotutil.ToIntErr(value)
	case "BOOL":
		v = iotutil.StringToBool(iotutil.ToString(value))
	case "INT":
		v, err = iotutil.ToInt64AndErr(value)
	case "DOUBLE", "FLOAT":
		v, err = iotutil.ToFloat64Err(value)
	case "TEXT":
		v = iotutil.ToString(value)
	case "FAULT":
		v = iotutil.ToInt(value)
	}
	if err != nil {
		v = value
	}
	return v
}

func GetMinValue(v1, v2 float64) float64 {
	if v1 > v2 {
		return v2
	}
	return v1
}

func GetMaxValue(v1, v2 float64) float64 {
	if v1 > v2 {
		return v1
	}
	return v2
}
