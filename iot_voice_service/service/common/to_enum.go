package common

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"encoding/json"
	"github.com/tidwall/gjson"
	"strings"
)

type ToEnum struct {
}

type VoiceToAxyResult struct {
	Data  interface{}
	VData interface{}
}

// VoiceToAxyConvert dataType 自有平台的数据类型
// vpmap 语音的配置
// inval 音响传入的数据
// 返回值：设备的值、音响的值、错误
func (s *ToEnum) VoiceToAxyConvert(devId string, adjust bool, dataType string, vpmap *protosService.OpmVoiceProductMap, inval gjson.Result) (interface{}, interface{}, error) {
	op := EnumToAxy{}
	switch dataType {
	case DT_ENUM:
		return op.enum2enum(vpmap, inval)
	case DT_BOOL:
		return op.enum2enum(vpmap, inval)
	case DT_INT:
		return op.enum2float(devId, adjust, vpmap, inval)
	case DT_FLOAT, DT_DOUBLE:
		return op.enum2float(devId, adjust, vpmap, inval)
	}
	return nil, nil, nil
}

// AxyToVoiceConvert dataType 自有平台的数据类型
// vpmap 语音的配置
// axyVal 音响传入的数据
// 返回值：设备的值、音响的值、错误
func (s *ToEnum) AxyToVoiceConvert(devId string, adjust bool, dataType string, vpmap *protosService.OpmVoiceProductMap, axyVal interface{}) (*VoiceToAxyResult, error) {
	op := EnumToVoice{}
	switch dataType {
	case DT_ENUM:
		return op.enum2enum(vpmap, axyVal)
	case DT_BOOL:
		return op.enum2enum(vpmap, axyVal)
	case DT_INT:
		return op.enum2float(devId, adjust, vpmap, axyVal)
	case DT_FLOAT, DT_DOUBLE:
		return op.enum2float(devId, adjust, vpmap, axyVal)
	}
	return nil, nil
}

type EnumToAxy struct {
}

// 枚举转枚举
// [{"val":0,"voiceValSynonym":"111"},{"val":1,"voiceValSynonym":"222"}]
func (EnumToAxy) enum2enum(vpmap *protosService.OpmVoiceProductMap, inval gjson.Result) (interface{}, interface{}, error) {
	var resData, resVData interface{}
	var voiceMap []map[string]interface{}
	err := json.Unmarshal([]byte(vpmap.ValueMap), &voiceMap)
	if err != nil {
		return nil, nil, err
	}
	vStr := strings.ToLower(iotutil.ToString(inval.Value()))
	for j := range voiceMap {
		voiceVal := iotutil.ToString(voiceMap[j]["voiceValSynonym"])
		voiceValArr := make([]string, 0)
		if voiceVal == "" {
			voiceVal = iotutil.ToString(voiceMap[j]["voiceVal"])
			voiceValArr = append(voiceValArr, voiceVal)
		} else {
			voiceValArr = strings.Split(voiceVal, ";")
		}
		iotlogger.LogHelper.Debugf("enum2enum voiceMap[%v]，%v", j, iotutil.ToString(voiceMap[j]))
		for _, s := range voiceValArr {
			tS := strings.ToLower(s)
			iotlogger.LogHelper.Debugf("enum2enum 数值比较：tempV:%v --- vStr: %v", vStr, tS)
			if vStr == tS {
				resData = ConvertValueType(vpmap.DataType, voiceMap[j]["val"])
				resVData = vStr
				return resData, resVData, nil
			}
		}
	}
	return resData, resVData, nil
}

// 枚举转数值
// [{"val":0,"voiceValSynonym":"","vMin":1,"vMax":8,"vStep":1},{"val":1,"voiceValSynonym":"","vMin":9,"vMax":100,"vStep":1}]
func (EnumToAxy) enum2float(devId string, adjust bool, vpmap *protosService.OpmVoiceProductMap, inval gjson.Result) (interface{}, interface{}, error) {
	var voiceMap []map[string]interface{}
	err := json.Unmarshal([]byte(vpmap.ValueMap), &voiceMap)
	if err != nil {
		return nil, nil, err
	}
	var resData, resVData interface{}
	vVal := iotutil.ToString(inval.Value())
	vVal = strings.ToLower(vVal)
	if err != nil {
		return nil, nil, err
	}
	for i, m := range voiceMap {
		//[{"min":0.1,"max":1,"step":0.1,"voiceValSynonym":"1"},{"val":"","voiceValSynonym":"2","min":1.1,"max":2,"step":0.1},{"val":"","voiceValSynonym":"3","min":2.1,"max":2,"step":0.1}]
		iotlogger.LogHelper.Debugf("enum2float voiceMap[%v], %v", i, iotutil.ToString(m))
		vMax, _ := iotutil.ToFloat64Err(m["max"])
		val := getVoiceVal(m)
		voiceValArr := strings.Split(iotutil.ToString(val), ";")
		for _, s := range voiceValArr {
			tS := strings.ToLower(s)
			iotlogger.LogHelper.Debugf("enum2float 数值比较：tempV:%v --- vStr: %v", vVal, tS)
			if vVal == tS {
				resData = vMax // * aMultiple
				break
			}
		}
	}
	resVData = inval.Value()
	return resData, resVData, nil
}

type EnumToVoice struct {
}

// 枚举转枚举
// [{"val":0,"voiceValSynonym":"111"},{"val":1,"voiceValSynonym":"222"}]
func (EnumToVoice) enum2enum(vpmap *protosService.OpmVoiceProductMap, axyVal interface{}) (*VoiceToAxyResult, error) {
	var (
		res      = &VoiceToAxyResult{}
		voiceMap []map[string]interface{}
	)
	err := json.Unmarshal([]byte(vpmap.ValueMap), &voiceMap)
	if err != nil {
		return res, err
	}
	vStr := strings.ToLower(iotutil.ToString(axyVal))
	for j := range voiceMap {
		//tempV := iotutil.ToString(voiceMap[j]["voiceValSynonym"])
		tempV := strings.ToLower(iotutil.ToString(getVoiceVal(voiceMap[j])))
		iotlogger.LogHelper.Debugf("数值比较：tempV:%v --- vStr: %v, voiceMap[%v]: %v", tempV, vStr, j, iotutil.ToString(voiceMap[j]))
		voiceValArr := strings.Split(iotutil.ToString(tempV), ";")
		for _, s := range voiceValArr {
			if vStr == s {
				tempVal := getFirstVoiceVal(voiceMap[j])
				res.VData = tempVal
				res.Data = iotutil.MapStringToInterface(map[string]string{"val": tempVal})["val"]
				break
			}
		}
		//if tempV == vStr {
		//	resData = getFirstVoiceVal(voiceMap[j])
		//}
	}
	return res, nil
}

// 枚举转数值
// [{"val":0,"voiceValSynonym":"","vMin":1,"vMax":8,"vStep":1},{"val":1,"voiceValSynonym":"","vMin":9,"vMax":100,"vStep":1}]
func (EnumToVoice) enum2float(devId string, adjust bool, vpmap *protosService.OpmVoiceProductMap, axyVal interface{}) (*VoiceToAxyResult, error) {
	var res = &VoiceToAxyResult{}
	var voiceMap []map[string]interface{}
	err := json.Unmarshal([]byte(vpmap.ValueMap), &voiceMap)
	if err != nil {
		return res, err
	}
	vVal, err := iotutil.ToFloat64Err(axyVal)
	if err != nil {
		return res, err
	}
	for _, m := range voiceMap {
		//{"val":0,"voiceValSynonym":"","vMin":1,"vMax":8,"vStep":1}
		//vMin, _ := iotutil.ToFloat64Err(m["vMin"])
		vMax, _ := iotutil.ToFloat64Err(m["vMax"])
		val, _ := iotutil.ToFloat64Err(m["val"])
		if vVal == val {
			res.Data = vMax
		}
	}
	res.VData = axyVal
	return res, nil
}
