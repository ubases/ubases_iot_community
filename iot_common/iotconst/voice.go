package iotconst

import "cloud_platform/iot_common/iotstruct"

const (
	VoiceNumberRange = 1
	VoiceNumberList  = 2
	VoiceString      = 3
)

var (
	VoiceUserTokenKey       = "voice_user_token_%s"
	XiaomiVoiceUserTokenKey = "xiaomi_voice_user_token_%s"
	AlexaVoiceUserTokenKey  = "alexa_voice_user_token_%s"
	AlexaVoiceAllUserIdKey  = "alexa_voice_all_user_id"
)

var (
	VoiceDataUnitList = map[string][]iotstruct.VoiceDataUnit{
		"℃": {
			{Label: "无", Unit: ""},
			{Label: "摄氏度", Unit: "℃"},
		},
		"%": {
			{Label: "无", Unit: ""},
			{Label: "百分比", Unit: "%"},
		},
		"ppm": {
			{Label: "无", Unit: ""},
			{Label: "百万分率", Unit: "ppm"},
		},
		"day": {
			{Label: "无", Unit: ""},
			{Label: "天", Unit: "day"},
		},
		"min": {
			{Label: "无", Unit: ""},
			{Label: "分钟", Unit: "min"},
		},
		"ug/m³": {
			{Label: "无", Unit: ""},
			{Label: "微克每立方米", Unit: "ug/m³"},
			{Label: "毫克每立方米", Unit: "mg/m³"},
			{Label: "克每立方米", Unit: "g/m³"},
			{Label: "千克每立方米", Unit: "kg/m³"},
		},
		"mg/m³": {
			{Label: "无", Unit: ""},
			{Label: "微克每立方米", Unit: "ug/m³"},
			{Label: "毫克每立方米", Unit: "mg/m³"},
			{Label: "克每立方米", Unit: "g/m³"},
			{Label: "千克每立方米", Unit: "kg/m³"},
		},
		"g/m³": {
			{Label: "无", Unit: ""},
			{Label: "微克每立方米", Unit: "ug/m³"},
			{Label: "毫克每立方米", Unit: "mg/m³"},
			{Label: "克每立方米", Unit: "g/m³"},
			{Label: "千克每立方米", Unit: "kg/m³"},
		},
		"kg/m³": {
			{Label: "无", Unit: ""},
			{Label: "微克每立方米", Unit: "ug/m³"},
			{Label: "毫克每立方米", Unit: "mg/m³"},
			{Label: "克每立方米", Unit: "g/m³"},
			{Label: "千克每立方米", Unit: "kg/m³"},
		},
	}

	VoiceUnitConvert = map[string]float64{
		"kg/m³": 1,
		"g/m³":  1000,
		"mg/m³": 1000000,
		"ug/m³": 1000000000,
		"℃":     1,
		"℉":     33.8,
		"%":     1,
	}
)
