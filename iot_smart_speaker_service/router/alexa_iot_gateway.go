package routers

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_smart_speaker_service/entitys"
	"cloud_platform/iot_smart_speaker_service/service/alexa"
	"cloud_platform/iot_smart_speaker_service/service/common"
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/tidwall/gjson"

	"github.com/gin-gonic/gin"
)

var alexaRouter AlexaIotGateWayRouter

type AlexaIotGateWayRouter struct {
	//
}

const (
	AlexaDiscoverDevicesResponse = "Alexa.Discovery"
	AlexaDeleteReport            = "DeleteReport"
	AlexaAddOrUpdateReport       = "AddOrUpdateReport"
	AcceptGrant                  = "Alexa.Authorization"
)

func (s *AlexaIotGateWayRouter) GateWay(c *gin.Context) {
	// 定义设备列表变量
	var err error
	// 获取设备列表查询接口请求原始数据
	data, err := c.GetRawData()
	if err != nil {
		c.String(http.StatusOK, "获取原始数据错误")
		return
	}
	common := entitys.DirectiveRequet{}
	if err := json.Unmarshal(data, &common); err != nil {
		c.String(http.StatusOK, err.Error())
	}
	token := gjson.GetBytes(data, "directive.payload.scope.token").String()
	if token == "" {
		token = gjson.GetBytes(data, "directive.endpoint.scope.token").String()
		if token == "" {
			token = gjson.GetBytes(data, "directive.payload.grantee.token").String()
		}
	}
	if token == "" {
		s.responseError(c, common, "Token为空", nil)
		return
	}
	//
	ti, err := manager.LoadAccessToken(context.Background(), token)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("加载访问密钥错误: ", err)
		s.responseError(c, common, "加载访问密钥错误", nil)
		return
	}
	userId := ti.GetUserID()
	iotlogger.LogHelper.Helper.Debug("userId: ", userId, "request:"+iotutil.ToString(common))

	err = alexa.RunController(&common, data, userId, token)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("set device discovery token error: ", err)
		s.responseError(c, common, err.Error(), nil)
		return
	}
	c.JSON(http.StatusOK, entitys.EventRequet{
		Event: common.Directive,
	})
}

// 返回错误信息
func (s *AlexaIotGateWayRouter) responseError(c *gin.Context, res entitys.DirectiveRequet, errMsg string, deviceResponse []entitys.DeviceResponse) {
	newRes := entitys.EventRequet{
		Event: entitys.CommonVoiceRequest{
			Header:   res.Directive.Header,
			Endpoint: res.Directive.Endpoint,
			Payload: map[string]interface{}{
				"type":    "INTERNAL_SERVICE_EXCEPTION",
				"message": errMsg,
			},
		},
	}
	newRes.Event.Header.Name = "ErrorResponse"
	newRes.Event.Header.Namespace = "Alexa"

	c.JSON(http.StatusOK, entitys.EventRequet{
		Event: res.Directive,
	})
}

// -------------
func (s *AlexaIotGateWayRouter) responseDeviceDiscovery(c *gin.Context, res entitys.DirectiveRequet, messageId string, devices []entitys.TmDevice) {
	endPoints := make([]entitys.AlexaEndpoints, 0)
	for _, device := range devices {
		capabilities := []entitys.AlexaCapabilities{}
		capabilities = append(capabilities, entitys.AlexaCapabilities{
			Type:      "AlexaInterface",
			Interface: "Alexa",
			Version:   "3",
			//Properties: entitys.AlexaProperties{},
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
			if rule.ValueType == 1 {
				//{"min":0,"max":10000000,"step":1,"unit":"mg/m³","vUnit":"mg/m³","vMin":0,"vMax":1000000,"vStep":1}
				valueMap := iotutil.JsonToMap(rule.ValueMap)
				minVal, maxVal, stepVal := valueMap["vMin"], valueMap["vMax"], valueMap["vStep"]
				configuration = &map[string]interface{}{
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
			} else if rule.ValueType == 2 {
				//[{"val":0,"voiceVal":"sleep"},{"val":1,"voiceVal":"quick"},{"val":2,"voiceVal":"auto"},{"val":3,"voiceVal":"low"},{"val":4,"voiceVal":"middle"},{"val":5,"voiceVal":"high"}]
				enumList, err := iotutil.JsonToMapArrayErr(rule.ValueMap)
				if err == nil {
					supportedModes := []map[string]interface{}{}
					for _, v := range enumList {
						voiceVal := common.GetFirstVoiceVal(v) // v["voiceVal"]
						supportedModes = append(supportedModes, map[string]interface{}{
							"ordered": false,
							"supportedModes": []map[string]interface{}{
								{
									"value": voiceVal,

									"modeResources": map[string]interface{}{
										"friendlyNames": []map[string]interface{}{
											{
												"@type": "text",
												"value": map[string]interface{}{
													"text":   voiceVal,
													"locale": "en-US",
												},
											},
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
			}
			capabilities = append(capabilities, entitys.AlexaCapabilities{
				Type:      "AlexaInterface",
				Interface: interfaceStr,
				Instance:  instanceStr,
				Version:   "3",
				Properties: &entitys.AlexaProperties{
					Supported:           []entitys.AlexaSupported{{Name: rule.AttrCode}},
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
			ManufacturerName: device.DeviceName,
			Description:      device.Brand,
			FriendlyName:     device.VoiceProduct.VoiceBrand, // device.DeviceName, //device.Zone
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
	}
	c.JSON(http.StatusOK, entitys.EventRequet{
		Event: res.Directive,
	})
}

func (s *AlexaIotGateWayRouter) responseSetDeviceProperty(c *gin.Context, res entitys.DirectiveRequet, messageId string, deviceResponse []entitys.DeviceResponse) {
	properties := []entitys.AlexaControlProperties{}
	if len(deviceResponse) > 0 {
		p := deviceResponse[0]
		for k, v := range p.Data {
			properties = append(properties, entitys.AlexaControlProperties{
				Namespace:                 res.Directive.Header.Namespace,
				Name:                      k,
				Value:                     iotutil.ToString(v),
				TimeOfSample:              time.Now().String(),
				UncertaintyInMilliseconds: 500,
			})
		}
	}
	res.Directive.Payload = entitys.AlexaControlPropertiesContext{Properties: properties}
	res.Directive.Header.Name = "Response"
	res.Directive.Header.Namespace = "Alexa"

	c.JSON(http.StatusOK, entitys.EventRequet{
		Event: res.Directive,
	})
}

func (s *AlexaIotGateWayRouter) responseAcceptGrant(c *gin.Context, res entitys.DirectiveRequet, messageId string) {
	res.Directive.Payload = entitys.AlexaControlPropertiesContext{}
	res.Directive.Header.Namespace = "Alexa.Authorization"
	res.Directive.Header.Name = "AcceptGrant.Response"
	c.JSON(http.StatusOK, entitys.EventRequet{
		Event: res.Directive,
	})
}
