package entitys

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
}

type DevicesRespPayload struct {
	Devices []TmDevice `json:"devices"`
}

type TmDeviceDiscoveryResp struct {
	Header  Header             `json:"header"`
	Payload DevicesRespPayload `json:"payload"`
}

func NewTmDeviceDiscoveryResp(messageId string, devices []TmDevice) Common {
	return Common{
		Header: Header{
			Namespace:      TmIotDeviceDiscovery,
			Name:           TmIotDeviceDiscoveryResp,
			MessageId:      messageId,
			PayLoadVersion: 2,
		},
		Payload: DevicesRespPayload{
			Devices: devices,
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
