package sync_update

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"cloud_platform/iot_voice_service/service/alexa"
	"cloud_platform/iot_voice_service/service/common"
	"cloud_platform/iot_voice_service/service/tianmao"
	"cloud_platform/iot_voice_service/service/xiaomi"
	"context"
	"encoding/json"
	"fmt"
)

func RunSync(userId string, deviceId string, payloadMap map[string]interface{}, product protosService.OpmVoiceProduct) {
	defer iotutil.PanicHandler(userId, deviceId)
	//天猫精灵同步用户的设备数据
	payload, err := QuerySyncVoice(product, payloadMap, "tianmao")
	if err == nil {
		tianmao.RequestSyncDeviceData(userId, deviceId, payload, product)
	}
	//小米同步用户的设备数据
	payload, err = QuerySyncVoice(product, payloadMap, "xiaomi")
	if err == nil {
		xiaomi.RequestSync(userId, deviceId, payload)
	}
	//Alexa的在线设备同步
	payload, err = QuerySyncVoice(product, payloadMap, "alexa")
	if err == nil {
		alexa.RequestSync(userId, deviceId, payload)
	}
}

func RunDeviceListSync(res UpdateDeviceSvc) {
	//天猫精灵同步用户的设备数据
	tianmao.RequestSyncDeviceList(*res.Data)
	//小米同步用户的设备数据
	xiaomi.RequestDeviceListSync(*res.Data)
	//Alexa的在线设备同步
	switch res.Data.UpdateType {
	case iotstruct.UPDATE_TYPE_REMOVE_DEVICE:
		alexa.RequestDeviceSync(*res.Data)
	case iotstruct.UPDATE_TYPE_CHANGE_FAMILY:
		alexa.RequestDeviceListSync(*res.Data)
	}
}

func QuerySyncVoice(product protosService.OpmVoiceProduct, payloadMap map[string]interface{}, voiceNo string) (map[string]interface{}, error) {
	res, err := common.GetVoiceProductDetails(context.Background(), product.ProductKey, voiceNo)
	if err != nil {
		return nil, err
	}
	payload := map[string]interface{}{}
	for _, productMap := range res.FunctionMap {
		dpid := iotutil.ToString(productMap.AttrDpid)
		v, ok := payloadMap[dpid]
		if !ok {
			continue
		}
		switch productMap.ValueType {
		case iotconst.VoiceNumberRange:
			var voiceMap map[string]interface{}
			err = json.Unmarshal([]byte(productMap.ValueMap), &voiceMap)
			if err != nil {
				return nil, err
			}
			payload[productMap.VoiceCode] = common.ConvertToVoiceNumberRange("", false, productMap.AttrDpid, voiceMap, v)
		case iotconst.VoiceNumberList:
			var voiceMap []map[string]interface{}
			err = json.Unmarshal([]byte(productMap.ValueMap), &voiceMap)
			if err != nil {
				return nil, err
			}
			axyValue := common.ConvertValueType(productMap.DataType, v)
			for j := range voiceMap {
				voiceValue := common.ConvertValueType(productMap.DataType, voiceMap[j]["val"])
				if axyValue == voiceValue {
					// 此处需按照dpid的值类型，就行值的响应转换
					voiceVal := common.GetFirstVoiceVal(voiceMap[j])
					payload[productMap.VoiceCode] = common.ConvertValueType(productMap.VDataType, voiceVal)
				}
			}
		case iotconst.VoiceString:
		default:
			iotlogger.LogHelper.Helper.Error("unexpected voice type: ", productMap.ValueType)
			return nil, fmt.Errorf("unexpected voice type: %v", productMap.ValueType)
		}
	}
	return payload, nil
}
