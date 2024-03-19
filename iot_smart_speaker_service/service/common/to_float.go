package common

import (
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"encoding/json"
	"github.com/tidwall/gjson"
)

type ToFloat struct {
}

// VoiceToAxyConvert dataType 自有平台的数据类型
// vpmap 语音的配置
// inval 音响传入的数据
// 返回值：设备的值、音响的值、错误
func (s *ToFloat) VoiceToAxyConvert(devId string, adjust bool, dataType string, vpmap *protosService.OpmVoiceProductMap, inval gjson.Result) (interface{}, interface{}, error) {
	op := FloatToAxy{}
	switch dataType {
	case DT_ENUM, DT_BOOL:
		return op.float2enum(vpmap, inval)
	case DT_INT, DT_FLOAT, DT_DOUBLE:
		return op.float2float(devId, adjust, vpmap, inval)
	}
	return nil, nil, nil
}

// AxyToVoiceConvert dataType 自有平台的数据类型
// vpmap 语音的配置
// axyVal 音响传入的数据
// 返回值：设备的值、音响的值、错误
func (s *ToFloat) AxyToVoiceConvert(devId string, adjust bool, dataType string, vpmap *protosService.OpmVoiceProductMap, axyVal interface{}) (interface{}, interface{}, error) {
	op := EnumToVoice{}
	switch dataType {
	case DT_ENUM:
		return op.enum2enum(vpmap, axyVal)
	case DT_BOOL:
		return op.enum2enum(vpmap, axyVal)
	case DT_INT, DT_FLOAT, DT_DOUBLE:
		return op.enum2float(devId, adjust, vpmap, axyVal)
	}
	return nil, nil, nil
}

type FloatToAxy struct {
}

// float2enum 浮点型转换为枚举
// [{"val":1,"voiceValSynonym":"","vMin":1,"vMax":100,"vStep":1},{"val":4,"voiceValSynonym":"","vMin":201,"vMax":300,"vStep":1}]
func (FloatToAxy) float2enum(vpmap *protosService.OpmVoiceProductMap, inval gjson.Result) (interface{}, interface{}, error) {
	var resData, resVData interface{}
	var voiceMap []map[string]interface{}
	err := json.Unmarshal([]byte(vpmap.ValueMap), &voiceMap)
	if err != nil {
		return nil, nil, err
	}
	vVal, _ := iotutil.ToFloat64Err(inval.Value())
	for _, m := range voiceMap {
		vMin, _ := iotutil.ToFloat64Err(m["vMin"])
		vMax, _ := iotutil.ToFloat64Err(m["vMax"])
		if vVal >= vMin && vVal <= vMax {
			resData, _ = iotutil.ToFloat64Err(m["val"])
			resVData = inval.Value()
			break
		}
	}
	return resData, resVData, nil
}

// float2float 浮点型转语控浮点型
// {"min":1,"max":4320,"step":30,"unit":"min","multiple":null,"vMin":100,"vMax":43200,"vStep":300,"vUnit":"min"}
func (FloatToAxy) float2float(devId string, adjust bool, vpmap *protosService.OpmVoiceProductMap, inval gjson.Result) (interface{}, interface{}, error) {
	var voiceMap map[string]interface{}
	err := json.Unmarshal([]byte(vpmap.ValueMap), &voiceMap)
	if err != nil {
		return nil, nil, err
	}
	var resData, resVData interface{}
	resData = ConvertToAxyNumberRange(devId, adjust, vpmap.AttrDpid, voiceMap, inval.Value())
	resVData = inval.Value()
	return resData, resVData, nil
}

// FloatToVoice Axy转语音平台
type FloatToVoice struct {
}

// float2enum 浮点型转换为枚举
// [{"val":1,"voiceValSynonym":"","vMin":1,"vMax":100,"vStep":1},{"val":4,"voiceValSynonym":"","vMin":201,"vMax":300,"vStep":1}]
func (FloatToVoice) float2enum(vpmap *protosService.OpmVoiceProductMap, axyVal interface{}) (interface{}, interface{}, error) {
	var resData, resVData interface{}
	var voiceMap []map[string]interface{}
	err := json.Unmarshal([]byte(vpmap.ValueMap), &voiceMap)
	if err != nil {
		return nil, nil, err
	}
	vVal, _ := iotutil.ToFloat64Err(axyVal)
	for _, m := range voiceMap {
		vMin, _ := iotutil.ToFloat64Err(m["min"])
		vMax, _ := iotutil.ToFloat64Err(m["max"])
		if vVal >= vMin && vVal <= vMax {
			//resData = getVoiceVal(m)
			resData = getFirstVoiceVal(m)
			break
		}
	}
	resVData = axyVal
	return resData, resVData, nil
}

// float2float 浮点型转语控浮点型
// {"min":1,"max":4320,"step":30,"unit":"min","multiple":null,"vMin":100,"vMax":43200,"vStep":300,"vUnit":"min"}
func (FloatToVoice) float2float(devId string, adjust bool, vpmap *protosService.OpmVoiceProductMap, axyVal interface{}) (interface{}, interface{}, error) {
	var voiceMap map[string]interface{}
	err := json.Unmarshal([]byte(vpmap.ValueMap), &voiceMap)
	if err != nil {
		return nil, nil, err
	}
	var resData, resVData interface{}
	resData = ConvertToVoiceNumberRange(devId, adjust, vpmap.AttrDpid, voiceMap, axyVal)
	resVData = axyVal
	return resData, resVData, nil
}
