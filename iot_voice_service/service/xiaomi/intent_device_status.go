package xiaomi

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_voice_service/entitys"
	"cloud_platform/iot_voice_service/service/common"
)

// 设备状态
func StatusDevices(res *entitys.XiaomiRequest, userId string, token string) error {
	//发现设备
	devices, err := common.DiscoveryDevices(userId, "xiaomi")
	if err != nil {
		return err
	}
	iotlogger.LogHelper.Helper.Debug("voice device list: ", devices)
	res.Devices = statusResponse(res, devices)
	return nil
}

// 发现设备响应
func statusResponse(res *entitys.XiaomiRequest, devices []entitys.TmDevice) []interface{} {
	devList := make([]interface{}, 0)
	deviceMap := map[string]entitys.TmDevice{}
	for _, device := range devices {
		deviceMap[device.DeviceId] = device
	}

	for _, d := range res.Devices {
		switch d.(type) {
		case string:
			if v, ok := deviceMap[iotutil.ToString(d)]; ok {
				devList = append(devList, entitys.XiaomiDeviceStatus{
					Did:    v.DeviceId,
					Online: v.IsOnline,
					Name:   v.DeviceName,
				})
			} else {
				devList = append(devList, entitys.XiaomiDeviceStatus{
					Did:         v.DeviceId,
					Status:      -1,
					Description: "not found device",
				})
			}
		}
	}
	return devList
}
