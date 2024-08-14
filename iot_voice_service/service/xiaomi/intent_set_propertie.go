package xiaomi

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_voice_service/entitys"
	"fmt"
)

// 发现设备
func SetProperties(res *entitys.XiaomiRequest, data []byte, userId string, token string) error {
	svc := XiaomiService{}
	//发现设备
	devResp, err := svc.SetDeviceProperty(data, *res, userId, false, "xiaomi")
	if err != nil {
		return err
	}
	setPropertiesResponse(res, devResp)
	return nil
}

// 发现设备响应
func setPropertiesResponse(res *entitys.XiaomiRequest, deviceResponse []entitys.DeviceResponse) error {
	iotlogger.LogHelper.Helper.Debug("小米Iot, deviceResponse:", iotutil.ToString(deviceResponse))
	deivceProps := make(map[string]entitys.XiaomiProperties)
	if len(deviceResponse) > 0 {
		for _, dev := range deviceResponse {
			for k, _ := range dev.Data {
				status := -1
				description := ""
				if dev.ErrorCode == "SUCCESS" {
					status = 0
				} else {
					description = dev.Message
				}
				pid, _ := iotutil.ToIntErr(k)
				deivceProps[fmt.Sprintf("%v_%v", dev.DeviceId, k)] = entitys.XiaomiProperties{
					Did:         dev.DeviceId,
					Siid:        2,
					Piid:        pid,
					Status:      status,
					Description: description,
				}
			}
		}
	}
	iotlogger.LogHelper.Helper.Debug("小米Iot, deivceProps:", iotutil.ToString(deivceProps))
	for i, properties := range res.Properties {
		var item entitys.XiaomiProperties
		err := iotutil.StructToStructErr(properties, &item)
		if err != nil {
			item.Status = -1
			item.Description = "error"
		} else {
			if v, ok := deivceProps[fmt.Sprintf("%v_%v", item.Did, item.Piid)]; ok {
				item.Value = v.Value
				item.Status = v.Status
				item.Description = v.Description
			} else {
				item.Value = v.Value
				item.Status = -1
				item.Description = "value not fond"
			}
		}
		res.Properties[i] = item
	}
	return nil
}
