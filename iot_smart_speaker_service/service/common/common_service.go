package common

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"cloud_platform/iot_smart_speaker_service/entitys"
	"cloud_platform/iot_smart_speaker_service/rpc/rpcclient"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
)

// 天猫精灵设备发现
func DiscoveryDevices(userId string, voiceNo string) ([]entitys.TmDevice, error) {
	ctx := context.Background()
	// 查询用户信息，并通过默认家庭id获取家庭设备列表
	respUser, err := rpcclient.ClientUcUserService.FindById(ctx, &protosService.UcUserFilter{
		Id:     iotutil.ToInt64(userId),
		Status: 1,
	})
	if err != nil {
		return nil, err
	}
	if respUser.Code != 200 {
		return nil, errors.New("用户不存在")
	}
	homeDevList, err := rpcclient.ClientIotDeviceHomeService.Lists(ctx, &protosService.IotDeviceHomeListRequest{
		Query: &protosService.IotDeviceHome{
			HomeId: iotutil.ToInt64(respUser.Data[0].DefaultHomeId),
		},
	})
	if err != nil {
		return nil, err
	}
	// 构建天猫所需厂商设备列表
	tmDevice := []entitys.TmDevice{}
	for i := range homeDevList.Data {
		device, err := GetTmDeviceInfo(ctx, homeDevList.Data[i].DeviceId, homeDevList.Data[i].RoomId, voiceNo)
		if err != nil {
			iotlogger.LogHelper.Helper.Error("get tm device info error: ", err)
			continue
		}
		tmDevice = append(tmDevice, device)
	}
	//iotlogger.LogHelper.Helper.Debug("get tm device list: ",	 iotutil.ToString(tmDevice))
	return tmDevice, nil
}

// 通过设备id获取设备信息
func GetTmDeviceInfo(ctx context.Context, devId string, roomId int64, voiceNo string) (entitys.TmDevice, error) {
	var (
		status = make(map[string]interface{})
		device entitys.TmDevice
	)
	deviceStatus, err := GetDeviceInfo(devId)
	if err != nil {
		return device, err
	}
	isOnline := deviceStatus["onlineStatus"] == "online"
	productKey := deviceStatus["productKey"]
	subscriptionId := deviceStatus["subscriptionId"]
	deviceName := deviceStatus["deviceName"]
	if deviceName == "" {
		deviceName = deviceStatus["productName"]
	}
	// TODO 通过房间id获取房间信息，待优化
	devRoom, err := rpcclient.ClientUcHomeRoomService.FindById(ctx, &protosService.UcHomeRoomFilter{
		Id: roomId,
	})
	if err != nil {
		return device, err
	}
	// 通过产品Key获取产品语控配置信息
	opmVoice, err := GetVoiceProductDetails(ctx, productKey, voiceNo)
	if err != nil {
		return device, err
	}
	for _, productMap := range opmVoice.FunctionMap {
		dpid := iotutil.ToString(productMap.AttrDpid)

		v, ok := deviceStatus[dpid]
		if !ok {
			continue
		}
		switch productMap.ValueType {
		case iotconst.VoiceNumberRange:
			var voiceMap map[string]interface{}
			err = json.Unmarshal([]byte(productMap.ValueMap), &voiceMap)
			if err != nil {
				return device, err
			}
			status[productMap.VoiceCode] = ConvertToVoiceNumberRange("", false, productMap.AttrDpid, voiceMap, v)
		case iotconst.VoiceNumberList:
			var voiceMap []map[string]interface{}
			err = json.Unmarshal([]byte(productMap.ValueMap), &voiceMap)
			if err != nil {
				return device, err
			}
			axyValue := ConvertValueType(productMap.DataType, v)
			for j := range voiceMap {
				voiceValue := ConvertValueType(productMap.DataType, voiceMap[j]["val"])
				if axyValue == voiceValue {
					// 此处需按照dpid的值类型，就行值的响应转换
					status[productMap.VoiceCode] = voiceValue // ConvertValueType(productMap.VDataType, voiceMap[j]["voiceVal"])
				}
			}
		case iotconst.VoiceString:
		default:
			iotlogger.LogHelper.Helper.Error("unexpected voice type: ", productMap.ValueType)
			return device, fmt.Errorf("unexpected voice type: %v", productMap.ValueType)
		}
	}
	// 构建天猫所需厂商设备信息
	device = entitys.TmDevice{
		DeviceId:       devId,
		DeviceType:     opmVoice.VoiceProductInfo.VoiceProductType,
		Brand:          gjson.Get(opmVoice.VoiceProductInfo.VoiceOther, "voiceBrand").String(),
		Model:          gjson.Get(opmVoice.VoiceProductInfo.VoiceOther, "voiceModel").String(),
		Zone:           devRoom.Data[0].RoomName,
		Status:         status,
		ProductKey:     productKey,
		VoiceProduct:   opmVoice,
		Extensions:     map[string]interface{}{},
		DeviceName:     deviceName,
		IsOnline:       isOnline,
		SubscriptionId: subscriptionId,
	}
	return device, nil
}

// VoiceToAxyFuncData 语音的属性值转换为爱星云属性值
func VoiceToAxyFuncData(ctx context.Context, devId string, params map[string]gjson.Result, adjust bool, voiceNo string) (bool, *entitys.ControlData, error) {
	devInfo, err := GetDeviceInfo(devId)
	if err != nil {
		return false, nil, err
	}
	//是否在线
	isOnline := devInfo["onlineStatus"] == "online"
	productKey := devInfo["productKey"]
	voiceProduct, err := GetVoiceProductDetails(ctx, productKey, voiceNo)
	if err != nil {
		return isOnline, nil, err
	}
	data := map[string]interface{}{}
	dataVoice := map[string]interface{}{}
	for k, v := range params {
		for _, productMap := range voiceProduct.FunctionMap {
			if k != productMap.VoiceCode {
				continue
			}
			// 属性值需要做转换, 需要判断值类型，数值范围，数值列表或者字符串
			dpId := iotutil.ToString(productMap.AttrDpid)
			sc := VoiceToAxySetup(productMap.VDataType)
			data[dpId], dataVoice[k], err = sc.VoiceToAxyConvert(devId, adjust, productMap.DataType, productMap, v)
			if err != nil {
				return isOnline, nil, err
			}
		}
	}

	//if len(data) == 0 {
	//未控制到属性
	//}

	ctl := &entitys.ControlData{
		ProductKey: productKey,
		DeviceId:   devId,
		Data:       data,
		DataVoice:  dataVoice,
	}
	return isOnline, ctl, nil
}

// VoiceGetFuncData 获取属性并设置为语控返回值  爱星云属性值转换为语音的属性值
func VoiceGetFuncData(ctx context.Context, devId string, params map[string]gjson.Result, baseParams map[string]string, adjust bool, voiceNo string) (*entitys.ControlData, error) {
	devInfo, err := GetDeviceInfo(devId)
	if err != nil {
		return nil, err
	}
	productKey := devInfo["productKey"]
	subscriptionId := devInfo["subscriptionId"]
	voiceProduct, err := GetVoiceProductDetails(ctx, productKey, voiceNo)
	if err != nil {
		return nil, err
	}
	data := map[string]interface{}{}
	dataVoiceCode := map[string]interface{}{}
	dataDpid := map[int32]interface{}{}
	//基础属性获取（兼容小米Iot逻辑）
	if baseParams != nil {
		for _, v := range baseParams {
			key := fmt.Sprintf("baseParams_%v", v)
			data[key] = devInfo[v]
			dataVoiceCode[key] = data[key]
		}
	}
	for k, _ := range params {
		for _, productMap := range voiceProduct.FunctionMap {
			if k != productMap.VoiceCode {
				continue
			}
			// 属性值需要做转换, 需要判断值类型，数值范围，数值列表或者字符串
			dpId := iotutil.ToString(productMap.AttrDpid)
			sc := VoiceToAxySetup(productMap.VDataType)
			if _, ok := devInfo[dpId]; ok {
				val, err := iotutil.ToFloat64Err(devInfo[dpId])
				if err == nil {
					data[dpId], _, err = sc.AxyToVoiceConvert(devId, adjust, productMap.DataType, productMap, val)
					if err != nil {
						return nil, err
					}
					dataVoiceCode[productMap.VoiceCode] = data[productMap.FunName]
					dataDpid[productMap.AttrDpid] = data[productMap.FunName]
				}
			}
			//
			//switch productMap.ValueType {
			//case iotconst.VoiceNumberRange:
			//	var voiceMap map[string]interface{}
			//	err = json.Unmarshal([]byte(productMap.ValueMap), &voiceMap)
			//	if err != nil {
			//		return nil, err
			//	}
			//	if _, ok := devInfo[dpid]; ok {
			//		val, err := iotutil.ToFloat64Err(devInfo[dpid])
			//		if err == nil {
			//			data[productMap.FunName] = ConvertToVoiceNumberRange(devId, adjust, productMap.AttrDpid, voiceMap, val)
			//			dataVoiceCode[productMap.VoiceCode] = data[productMap.FunName]
			//			dataDpid[productMap.AttrDpid] = data[productMap.FunName]
			//		}
			//	}
			//case iotconst.VoiceNumberList:
			//	var voiceMap []map[string]interface{}
			//	err = json.Unmarshal([]byte(productMap.ValueMap), &voiceMap)
			//	if err != nil {
			//		return nil, err
			//	}
			//	if _, ok := devInfo[dpid]; ok {
			//		val := iotutil.ToString(devInfo[dpid])
			//		//data[dpid] = val
			//		for j := range voiceMap {
			//			tempV := iotutil.ToString(voiceMap[j]["val"])
			//			if tempV == val {
			//				data[productMap.FunName] = iotutil.ToString(voiceMap[j]["voiceVal"])
			//				dataVoiceCode[productMap.VoiceCode] = data[productMap.FunName]
			//				dataDpid[productMap.AttrDpid] = data[productMap.FunName]
			//			}
			//		}
			//	}
			//case iotconst.VoiceString:
			//	if _, ok := devInfo[dpid]; ok {
			//		val, err := iotutil.ToFloat64Err(devInfo[dpid])
			//		if err == nil {
			//			data[productMap.FunName] = val
			//			dataVoiceCode[productMap.VoiceCode] = data[productMap.FunName]
			//			dataDpid[productMap.AttrDpid] = data[productMap.FunName]
			//		}
			//	}
			//default:
			//	if _, ok := devInfo[dpid]; ok {
			//		val, err := iotutil.ToFloat64Err(devInfo[dpid])
			//		if err == nil {
			//			data[productMap.FunName] = val
			//			dataVoiceCode[productMap.VoiceCode] = data[productMap.FunName]
			//			dataDpid[productMap.AttrDpid] = data[productMap.FunName]
			//		}
			//	}
			//}
		}
	}
	ctl := &entitys.ControlData{
		ProductKey:     productKey,
		DeviceId:       devId,
		Data:           data,
		DataDpid:       dataDpid,
		DataVoiceCode:  dataVoiceCode,
		SubscriptionId: subscriptionId,
	}
	return ctl, nil
}
