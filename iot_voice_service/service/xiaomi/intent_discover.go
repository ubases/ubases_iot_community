package xiaomi

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_voice_service/entitys"
	"cloud_platform/iot_voice_service/service/common"
)

// 发现设备
func DiscoverDevices(res *entitys.XiaomiRequest, userId string, token string) error {
	//发现设备
	devices, err := common.DiscoveryDevices(userId, "xiaomi")
	if err != nil {
		return err
	}
	iotlogger.LogHelper.Helper.Debug("voice device list: ", devices)
	res.Devices = discoverResponse(devices)
	return nil
}

// 发现设备响应
func discoverResponse(devices []entitys.TmDevice) []interface{} {

	devList := make([]interface{}, 0)
	for _, device := range devices {
		devList = append(devList, entitys.XiaomiDevices{
			Did:  device.DeviceId,
			Type: device.VoiceProduct.VoiceProductInfo.VoiceProductType, // "urn:miot-spec-v2:device:humidifier:0000A00E:180-xjy:1", //读取语控配置中的type
			Name: device.DeviceName,
		})
	}
	return devList
}
