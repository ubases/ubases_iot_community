package xiaomi

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_smart_speaker_service/entitys"
	"cloud_platform/iot_smart_speaker_service/service/common"
	"context"
)

// 发现设备
func SubscribeDevices(res *entitys.XiaomiRequest, userId string, token string) error {
	//发现设备
	devices, err := common.DiscoveryDevices(userId, "xiaomi")
	if err != nil {
		return err
	}
	iotlogger.LogHelper.Helper.Debug("voice device list: ", devices)
	subscribeResponse(res, devices)
	return nil
}

func UnSubscribeDevices(res *entitys.XiaomiRequest, userId string, token string) error {
	unSubscribeResponse(res)
	return nil
}

// 订阅
func subscribeResponse(res *entitys.XiaomiRequest, devices []entitys.TmDevice) {
	deviceMap := map[string]entitys.TmDevice{}
	for _, device := range devices {
		deviceMap[device.DeviceId] = device
	}
	for i, device := range res.Devices {
		var item entitys.XiaomiSubscribe
		err := iotutil.StructToStructErr(device, &item)
		if err != nil {
			continue
		}
		if item.Did != "" {
			iotredis.GetClient().HMSet(context.Background(), iotconst.HKEY_DEV_DATA_PREFIX+item.Did, "subscriptionId", item.SubscriptionId)
		}
		if _, ok := deviceMap[item.Did]; ok {
			item.Status = 0
		} else {
			item.Status = -1
			item.Description = "invalid device id"
		}
		res.Devices[i] = item
	}
}

// 取消订阅
func unSubscribeResponse(res *entitys.XiaomiRequest) {
	for i, device := range res.Devices {
		var item entitys.XiaomiSubscribe
		err := iotutil.StructToStructErr(device, &item)
		if err != nil {
			item.Status = 0
		} else {
			item.Status = -1
			item.Description = "err"
		}
		res.Devices[i] = item
	}
}
