package alexa

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_smart_speaker_service/entitys"
	"cloud_platform/iot_smart_speaker_service/service/common"
	"strings"
)

// 发现设备
func DiscoverDevices(res *entitys.DirectiveRequet, userId string, token string) error {
	//发现设备
	devices, err := common.DiscoveryDevices(userId, "alexa")
	if err != nil {
		return err
	}
	iotlogger.LogHelper.Helper.Debug("voice device list: ", iotutil.ToString(devices))
	discoverResponse(res, devices)
	return nil
}

// 发现设备响应
func discoverResponse(res *entitys.DirectiveRequet, devices []entitys.TmDevice) {
	res.Directive.Header.Name = "Discover.Response"
	endPoints := make([]entitys.AlexaEndpoints, 0)
	for _, device := range devices {
		capabilities := []entitys.AlexaCapabilities{}
		capabilities = append(capabilities, entitys.AlexaCapabilities{
			Type:      "AlexaInterface",
			Interface: "Alexa",
			Version:   "3",
		})
		capabilities = append(capabilities, entitys.AlexaCapabilities{
			Type:      "AlexaInterface",
			Interface: "Alexa.BrightnessController",
			Version:   "3",
			Properties: &entitys.AlexaProperties{
				Supported:           []entitys.AlexaSupported{{Name: "brightness"}},
				ProactivelyReported: false,
				Retrievable:         false,
				//NonControllable:     false,
			},
		})
		for _, rule := range device.VoiceProduct.FunctionMap {
			interfaceStr, instanceStr := "", ""
			arr := strings.Split(rule.VoiceCode, "#")
			interfaceStr = arr[0]
			if len(arr) > 1 {
				instanceStr = arr[1]
			}
			//如果是范围
			var configuration *map[string]interface{}
			switch rule.VDataType {
			case "ENUM":
				configuration = enumConfiguration(rule.ValueMap)
			case "DOUBLE", "FLOAT", "INT":
				configuration = rangeConfiguration(rule.ValueMap)
			case "BOOL":
				if interfaceStr == "Alexa.ModeController" {
					configuration = enumConfiguration(rule.ValueMap)
				}
			case "TEXT":
			}
			capabilities = append(capabilities, entitys.AlexaCapabilities{
				Type:      "AlexaInterface",
				Interface: interfaceStr,
				Instance:  instanceStr,
				Version:   "3",
				Properties: &entitys.AlexaProperties{
					Supported:           []entitys.AlexaSupported{{Name: instanceStr}},
					ProactivelyReported: false,
					Retrievable:         false,
					//NonControllable:     false,
				},
				CapabilityResources: &entitys.AlexaCapabilityResources{
					FriendlyNames: []map[string]interface{}{
						{
							"@type": "text",
							"value": map[string]interface{}{
								"text":   rule.FunName,
								"locale": "en-US",
							},
						},
					},
				},
				Configuration: configuration,
			})
		}
		endPoints = append(endPoints, entitys.AlexaEndpoints{
			EndpointID:       device.DeviceId,
			ManufacturerName: device.VoiceProduct.VoiceBrand,
			Description:      device.Brand,
			FriendlyName:     device.DeviceName, // device.DeviceName, //device.VoiceProduct.VoiceBrand,
			AdditionalAttributes: &entitys.AdditionalAttributes{
				Manufacturer:     device.Brand,
				Model:            device.ProductKey,
				SerialNumber:     device.ProductKey,
				FirmwareVersion:  "1.0.0",
				SoftwareVersion:  "1.0.0",
				CustomIdentifier: device.DeviceId,
			},
			DisplayCategories: []string{"OTHER"},
			Cookie:            &entitys.AlexaCookie{},
			Capabilities:      capabilities,
			Connections:       []entitys.AlexaConnections{},
		})
	}
	res.Directive.Payload = entitys.AlexaPayload{
		Endpoints: endPoints,
		Scope:     nil,
	}
}

// 范围类的配置内容转换
func rangeConfiguration(valueMapStr string) *map[string]interface{} {
	//{"min":0,"max":10000000,"step":1,"unit":"mg/m³","vUnit":"mg/m³","vMin":0,"vMax":1000000,"vStep":1}
	valueMap := iotutil.JsonToMap(valueMapStr)
	minVal, maxVal, stepVal := valueMap["vMin"], valueMap["vMax"], valueMap["vStep"]
	configuration := &map[string]interface{}{
		"supportedRange": map[string]interface{}{
			"minimumValue": minVal,
			"maximumValue": maxVal,
			"precision":    stepVal,
		},
		"presets": []map[string]interface{}{
			{
				"rangeValue": maxVal,
				"presetResources": map[string]interface{}{
					"friendlyNames": []map[string]interface{}{
						{
							"@type": "asset",
							"value": map[string]interface{}{
								"assetId": "Alexa.Value.Maximum",
							},
						},
						{
							"@type": "asset",
							"value": map[string]interface{}{
								"assetId": "Alexa.Value.High",
							},
						},
						{
							"@type": "text",
							"value": map[string]interface{}{
								"text":   "Highest",
								"locale": "en-US",
							},
						},
						{
							"@type": "text",
							"value": map[string]interface{}{
								"text":   "Fast",
								"locale": "en-US",
							},
						},
						{
							"@type": "text",
							"value": map[string]interface{}{
								"text":   "Alta",
								"locale": "es-MX",
							},
						},
						{
							"@type": "text",
							"value": map[string]interface{}{
								"text":   "Élevée",
								"locale": "fr-CA",
							},
						},
					},
				},
			},
			{
				"rangeValue": minVal,
				"presetResources": map[string]interface{}{
					"friendlyNames": []map[string]interface{}{
						{
							"@type": "asset",
							"value": map[string]interface{}{
								"assetId": "Alexa.Value.Minimum",
							},
						},
						{
							"@type": "asset",
							"value": map[string]interface{}{
								"assetId": "Alexa.Value.Low",
							},
						},
						{
							"@type": "text",
							"value": map[string]interface{}{
								"text":   "Lowest",
								"locale": "en-US",
							},
						},
						{
							"@type": "text",
							"value": map[string]interface{}{
								"text":   "Slow",
								"locale": "en-US",
							},
						},
						{
							"@type": "text",
							"value": map[string]interface{}{
								"text":   "Baja",
								"locale": "es-MX",
							},
						},
						{
							"@type": "text",
							"value": map[string]interface{}{
								"text":   "Faible",
								"locale": "fr-CA",
							},
						},
					},
				},
			},
		},
	}
	return configuration
}

// 枚举类的配置内容转换
func enumConfiguration(valueMapStr string) *map[string]interface{} {
	var configuration *map[string]interface{}
	//[{"val":0,"voiceVal":"sleep"},{"val":1,"voiceVal":"quick"},{"val":2,"voiceVal":"auto"},{"val":3,"voiceVal":"low"},{"val":4,"voiceVal":"middle"},{"val":5,"voiceVal":"high"}]
	enumList, err := iotutil.JsonToMapArrayErr(valueMapStr)
	if err == nil {
		supportedModes := []map[string]interface{}{}
		for _, v := range enumList {
			voiceVal := common.GetFirstVoiceVal(v) // v["voiceVal"]
			supportedModes = append(supportedModes, map[string]interface{}{
				"value": iotutil.ToString(voiceVal),
				"modeResources": map[string]interface{}{
					"friendlyNames": []map[string]interface{}{
						{
							"@type": "text",
							"value": map[string]interface{}{
								"text":   iotutil.ToString(voiceVal),
								"locale": "en-US",
							},
						},
					},
				},
			})
		}
		configuration = &map[string]interface{}{
			"ordered":        false,
			"supportedModes": supportedModes,
		}
	}
	return configuration
}
