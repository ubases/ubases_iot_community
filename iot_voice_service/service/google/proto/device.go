package proto

const (
	CUSTOMDATA_PRODUCTKEY = "productKey"
)

type OtherDeviceIds struct {
	DeviceId string `json:"deviceId"`
}

//Device  Google智能家居设备
type Device struct {
	Id              string                 `json:"id"`
	Type            DeviceType             `json:"type"`
	Traits          []string               `json:"traits"`
	Name            DeviceName             `json:"name"`
	WillReportState bool                   `json:"willReportState"`
	Attributes      map[string]interface{} `json:"attributes,omitempty"`
	RoomHint        string                 `json:"roomHint,omitempty"`
	DeviceInfo      DeviceInfo             `json:"deviceInfo,omitempty"`
	CustomData      map[string]interface{} `json:"customData,omitempty"`
	OtherDeviceIds  []OtherDeviceIds       `json:"otherDeviceIds,omitempty"`
}

// DeviceName 设备名、默认名称、昵称
type DeviceName struct {
	Name         string   `json:"name"`
	DefaultNames []string `json:"defaultNames,omitempty"`
	Nicknames    []string `json:"nicknames,omitempty"`
}

// DeviceInfo 设备元数据
type DeviceInfo struct {
	Manufacturer string `json:"manufacturer,omitempty"`
	Model        string `json:"model,omitempty"`
	HwVersion    string `json:"hwVersion,omitempty"`
	SwVersion    string `json:"swVersion,omitempty"`
}
