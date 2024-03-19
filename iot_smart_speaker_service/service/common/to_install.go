package common

import (
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"github.com/tidwall/gjson"
	"strings"
)

var (
	DT_ENUM   = "ENUM"
	DT_BOOL   = "BOOL"
	DT_INT    = "INT"
	DT_FLOAT  = "FLOAT"
	DT_DOUBLE = "DOUBLE"
)

type ITo interface {
	VoiceToAxyConvert(devId string, adjust bool, dataType string, vpmap *protosService.OpmVoiceProductMap, inval gjson.Result) (interface{}, interface{}, error)
	AxyToVoiceConvert(devId string, adjust bool, dataType string, vpmap *protosService.OpmVoiceProductMap, axyVal interface{}) (interface{}, interface{}, error)
}

// VoiceToAxySetup 安装服务
func VoiceToAxySetup(vDataType string) ITo {
	switch vDataType {
	case DT_ENUM, DT_BOOL:
		return new(ToEnum)
	case DT_INT, DT_FLOAT, DT_DOUBLE:
		return new(ToFloat)
	}
	return nil
}

func getVoiceVal(m map[string]interface{}) interface{} {
	if v, ok := m["voiceValSynonym"]; ok {
		return v
	}
	if v, ok := m["voiceVal"]; ok {
		return v
	}
	return nil
}

func GetVoiceVal(m map[string]interface{}) interface{} {
	return getVoiceVal(m)
}

func getFirstVoiceVal(m map[string]interface{}) interface{} {
	//如果是别名模式，则需要返回“；”分隔中的第一项
	if v, ok := m["voiceValSynonym"]; ok {
		tempV := iotutil.ToString(v)
		return strings.Split(tempV, ";")[0]
	}
	if v, ok := m["voiceVal"]; ok {
		return v
	}
	return nil
}

func GetFirstVoiceVal(m map[string]interface{}) interface{} {
	return getVoiceVal(m)
}
