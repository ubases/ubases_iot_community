package google

import (
	"cloud_platform/iot_smart_speaker_service/service/common"
	"cloud_platform/iot_smart_speaker_service/service/google/proto"
)

// Device
// 参考链接 https://developers.google.com/assistant/smarthome/concepts/devices-traits
type Device interface {
	DeviceId() string
	DeviceName() proto.DeviceName
	DeviceType() proto.DeviceType
	DeviceTraits() []Trait
	DeviceCustomData() map[string]interface{}
	DeviceCommand() map[string]*CommandInfo
	GetOtherDeviceIds() []proto.OtherDeviceIds
	GetDeviceOnlineStatus() bool
}

type DeviceInfoProvider interface {
	DeviceInfo() proto.DeviceInfo
}

type DeviceRoomHintProvider interface {
	DeviceRoomHint() string
}

type NumberRange struct {
	Min  float64
	Max  float64
	Step float64
}

type CommandInfo struct {
	Dpid           uint8
	Name           string
	DataType       string
	TraitName      string
	Default        string                 //从0或1开始
	MapValue       map[string]int         //用于枚举，语控的枚举值->功能点的枚举值
	MapNumberRange map[string]NumberRange //用于枚举，语控的枚举值->功能点的数值范围
}

type BasicDevice struct {
	Id             string
	Name           proto.DeviceName
	Type           proto.DeviceType
	Traits         []Trait
	Info           proto.DeviceInfo
	RoomHint       string
	CustomData     map[string]interface{}
	CommandInfo    map[string]*CommandInfo
	OtherDeviceIds []proto.OtherDeviceIds
}

func (d BasicDevice) DeviceRoomHint() string {
	return d.RoomHint
}

func (d BasicDevice) DeviceType() proto.DeviceType {
	return d.Type
}

func (d BasicDevice) DeviceInfo() proto.DeviceInfo {
	return d.Info
}

func (d BasicDevice) DeviceId() string {
	return d.Id
}

func (d BasicDevice) DeviceName() proto.DeviceName {
	return d.Name
}

func (d BasicDevice) DeviceTraits() []Trait {
	return d.Traits
}

func (d BasicDevice) DeviceCustomData() map[string]interface{} {
	return d.CustomData
}

func (d BasicDevice) DeviceCommand() map[string]*CommandInfo {
	return d.CommandInfo
}
func (d BasicDevice) GetOtherDeviceIds() []proto.OtherDeviceIds {
	return d.OtherDeviceIds
}
func (d BasicDevice) GetDeviceOnlineStatus() bool {
	return common.GetDeviceOnline(d.Id)
}
