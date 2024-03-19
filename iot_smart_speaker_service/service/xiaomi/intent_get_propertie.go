package xiaomi

import (
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_smart_speaker_service/entitys"
	"fmt"
)

// 发现设备
func GetProperties(res *entitys.XiaomiRequest, data []byte, userId string, token string) error {
	svr := XiaomiService{}
	//发现设备
	devResp, err := svr.GetDeviceProperty(*res, userId, false, "xiaomi")
	if err != nil {
		return err
	}
	getPropertiesResponse(res, devResp)
	return nil
}

// 发现设备响应
func getPropertiesResponse(res *entitys.XiaomiRequest, deviceResponse []entitys.DeviceResponse) error {
	deivceProps := make(map[string]interface{})
	if len(deviceResponse) > 0 {
		for _, dev := range deviceResponse {
			for k, v := range dev.DataVoice {
				deivceProps[fmt.Sprintf("%v_%v", dev.DeviceId, k)] = v
			}
		}
	}
	for i, properties := range res.Properties {
		var item entitys.XiaomiProperties
		err := iotutil.StructToStructErr(properties, &item)
		if err != nil {
			continue
		}
		if item.Siid == float64(1) {
			//小米基础属性数值绑定
			name := XiaomiBaseProperties[iotutil.ToString(item.Piid)]
			if v, ok := deivceProps[fmt.Sprintf("%s_baseParams_%s", item.Did, name)]; ok {
				item.Value = v
				item.Status = 0
			} else {
				item.Status = -1
				item.Description = "base value not fond"
			}
		} else {
			//设备自定义功能属性数值绑定
			if v, ok := deivceProps[fmt.Sprintf("%v_%v", item.Did, item.Piid)]; ok {
				item.Value = v
				item.Status = 0
			} else {
				item.Status = -1
				item.Description = "value not fond"
			}
		}
		res.Properties[i] = item
	}
	return nil
}
