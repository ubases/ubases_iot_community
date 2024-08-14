package entitys

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"fmt"
)

const (
	TmIotDeviceDiscovery     = "AliGenie.Iot.Device.Discovery"
	TmIotDeviceDiscoveryResp = "DiscoveryDevicesResponse"
	TmIotDeviceControl       = "AliGenie.Iot.Device.Control"
	TmIotDeviceControlResp   = "CorrectResponse"
)

type Header struct {
	Namespace      string `json:"namespace"`
	Name           string `json:"name"`
	MessageId      string `json:"messageId"`
	PayLoadVersion int    `json:"payLoadVersion"`
}

type Common struct {
	Header  Header      `json:"header"`
	Payload interface{} `json:"payload"`
}

type DevicesReqPayload struct {
	AccessToken string `json:"accessToken"`
}

type TmDevice struct {
	ProductKey     string                 `json:"productKey"`
	DeviceId       string                 `json:"deviceId"`
	DeviceName     string                 `json:"deviceName"`
	DeviceType     string                 `json:"deviceType"`
	Brand          string                 `json:"brand"`
	Model          string                 `json:"model"`
	Zone           string                 `json:"zone"`
	Status         map[string]interface{} `json:"status"`
	Extensions     map[string]interface{} `json:"extensions"`
	VoiceProduct   *VoiceProductCached    `json:"convertRule"` //功能转换规则
	IsOnline       bool                   `json:"isOnline"`
	SubscriptionId string                 `json:"subscriptionId"`
	Icon           string                 `json:"icon"`
}

type TmDeviceV2 struct {
	DeviceId   string                   `json:"deviceId"`
	DeviceName string                   `json:"deviceName"`
	DeviceType string                   `json:"deviceType"`
	Brand      string                   `json:"brand"`
	Model      string                   `json:"model"`
	Zone       string                   `json:"zone"`
	Icon       string                   `json:"icon"`
	Properties []map[string]interface{} `json:"properties"`
	Actions    []string                 `json:"actions"`
	Extensions map[string]interface{}   `json:"extensions,omitempty"`
}

type DevicesRespPayload struct {
	Devices []TmDeviceV2 `json:"devices"`
}

type TmDeviceDiscoveryResp struct {
	Header  Header             `json:"header"`
	Payload DevicesRespPayload `json:"payload"`
}

func NewTmDeviceDiscoveryResp(messageId string, devices []TmDevice) Common {
	var newDevs = make([]TmDeviceV2, 0)

	for _, device := range devices {
		var properties = make([]map[string]interface{}, 0)
		var actions = make([]string, 0)
		for _, productMap := range device.VoiceProduct.FunctionMap {
			p := map[string]interface{}{
				"name":  productMap.VoiceCode,
				"value": device.Status[productMap.VoiceCode],
			}
			//INT、DOUBLE、TEXT、ENUM、BOOL
			switch productMap.VDataType {
			case "INT", "DOUBLE":
				actions = append(actions, fmt.Sprintf("Set%v", productMap.VoiceCode))
				actions = append(actions, fmt.Sprintf("Adjust%v", productMap.VoiceCode))
			case "TEXT", "ENUM":
				actions = append(actions, fmt.Sprintf("Set%v", productMap.VoiceCode))
			case "BOOL":
				actions = append(actions, fmt.Sprintf("TurnOn"))
				actions = append(actions, fmt.Sprintf("TurnOff"))
				if iotutil.ToString(p["value"]) == "true" {
					p["value"] = "on"
				} else {
					p["value"] = "off"
				}
			}
			//productMap.VoiceCode
			properties = append(properties, p)
		}
		actions = append(actions, fmt.Sprintf("%v", "Query"))
		actions = iotutil.RemoveRepeatElement(actions)

		newDevs = append(newDevs, TmDeviceV2{
			DeviceId:   device.DeviceId,
			DeviceName: device.DeviceName,
			DeviceType: device.DeviceType,
			Brand:      device.Brand,
			Model:      device.Model,
			Zone:       device.Zone,
			Icon:       device.Icon,
			Properties: properties,
			Actions:    actions,
		})
	}
	iotlogger.LogHelper.Errorf("NewTmDeviceDiscoveryResp: %v", iotutil.ToString(newDevs))
	return Common{
		Header: Header{
			Namespace:      TmIotDeviceDiscovery,
			Name:           TmIotDeviceDiscoveryResp,
			MessageId:      messageId,
			PayLoadVersion: 2,
		},
		Payload: DevicesRespPayload{
			Devices: newDevs,
		},
	}
}

type DevicePropertyPayload struct {
	AccessToken string                 `json:"accessToken"`
	DeviceIds   []string               `json:"deviceIds"`
	Params      map[string]interface{} `json:"params"`
	Extensions  map[string]interface{} `json:"extensions"`
}

type DeviceResponse struct {
	DeviceId   string                 `json:"deviceId"`
	DeviceName string                 `json:"deviceName"`
	ErrorCode  string                 `json:"errorCode"`
	Message    string                 `json:"message"`
	Data       map[string]interface{} `json:"data"`
	DataDpid   map[int32]interface{}  //Key为Dpid
	DataVoice  map[string]interface{}
}

type DeviceResponsePayload struct {
	DeviceResponseList []DeviceResponse `json:"deviceResponseList"`
}

func NewDeviceResponse(messageId string, devRespList []DeviceResponse) Common {
	return Common{
		Header: Header{
			Namespace:      TmIotDeviceControl,
			Name:           TmIotDeviceControlResp,
			MessageId:      messageId,
			PayLoadVersion: 2,
		},
		Payload: DeviceResponsePayload{
			DeviceResponseList: devRespList,
		},
	}
}
