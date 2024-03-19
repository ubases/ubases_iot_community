package xiaoai

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotgincache/persist"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotprotocol"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"cloud_platform/iot_smart_speaker_service/cached"
	"cloud_platform/iot_smart_speaker_service/entitys"
	"cloud_platform/iot_smart_speaker_service/service/common"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-oauth2/oauth2/v4/manage"

	"github.com/tidwall/gjson"
)

type XiaoaiIoTGateway struct {
}

var xiaoaiVoiceCode = "xiaoai"

func (s XiaoaiIoTGateway) Run(req entitys.XiaoaiRequest, manager *manage.Manager, testUserId string) (string, error) {
	iotlogger.LogHelper.Helper.Debug("session: ", req)
	var userId string
	if testUserId != "" {
		userId = testUserId
	} else {
		token := req.Session.User.AccessToken //  gjson.GetBytes(data, "session.user.access_token").String()
		ti, err := manager.LoadAccessToken(context.Background(), token)
		if err != nil {
			iotlogger.LogHelper.Helper.Error("加载访问密钥错误: ", err)
			return "加载访问密钥错误", nil
		}

		userId = ti.GetUserID()
	}
	iotlogger.LogHelper.Helper.Debug("userId: ", userId)

	var msg string = "暂不支持该操作" //返回默认
	slotInfo := req.Request.SlotInfo

	//欢迎（打开xxx）
	//if req.Request.SlotInfo.IntentName == "Mi_Welcome" {
	//	return "已打开", nil
	//}

	//通过前缀判断是否为获取属性
	isRead := strings.HasPrefix(req.Request.SlotInfo.IntentName, "get_")
	//slotInfo.IntentName 意图
	//slotInfo.Slots 插槽
	skillId := req.Request.Intent.AppID //技能Id
	//通过意图获取查询产品信息和其语控配置信息
	voiceInfo, err := common.GetDetail(protosService.OpmVoiceProductDetailReq{SkillId: skillId, VoiceNo: xiaoaiVoiceCode})
	if err != nil {
		iotlogger.LogHelper.Helper.Error("未获取到语控配置: ", err)
		return "未获取到语控配置", nil
	}
	iotlogger.LogHelper.Info("Voice ProductKey:", voiceInfo.ProductKey)
	//attrList := voiceInfo.AttrList
	productKey := voiceInfo.ProductKey
	//找到设备Id
	devices, err := common.DiscoveryDevices(userId, xiaoaiVoiceCode)
	if err != nil {
		return err.Error(), nil
	}
	//runDevice := make([]entitys.TmDevice, 0)
	runDeviceMap := make(map[string]entitys.TmDevice)
	//runDeviceIds := make([]string, 0)
	funcs := make(map[string]gjson.Result, 0)
	for _, device := range devices {
		//产品Key
		if device.ProductKey == productKey {
			hasName := false
			for _, slot := range slotInfo.Slots {
				switch strings.ToLower(slot.Name) {
				case "devicename":
					//设备名称
					if device.DeviceName == slot.RawValue {
						runDeviceMap[device.DeviceId] = device
					} else {
						delete(runDeviceMap, device.DeviceId)
					}
					hasName = true
				case "roomname":
					//房间名称
					if device.Zone == slot.RawValue {
						runDeviceMap[device.DeviceId] = device
					} else {
						delete(runDeviceMap, device.DeviceId)
					}
					hasName = true
				default:
					//从用户的意图配置中获取
					for _, item := range voiceInfo.AttrList {
						//将{亮度}设置为{50%}
						if slot.Name == item.VoiceCode {
							funcs[iotutil.ToString(item.VoiceCode)] = gjson.Parse("{\"value\": \"" + slot.RawValue + "\"}").Get("value")
							//item.AttrDpid  slot.RawValue
							//将50%转换为 max value * 50%
							// gjson.Parse(slot.RawValue) 转换为我司物模型
							//gjsonRes = gjson.Parse("{\"value\": \"" + slot.RawValue + "\"}").Get("value")
						}
					}
				}
			}
			//是否指定设备名称
			if !hasName {
				runDeviceMap[device.DeviceId] = device
			}
		}
	}
	if len(runDeviceMap) == 0 {
		iotlogger.LogHelper.Helper.Error("未找到设备")
		return "未找到设备", nil
	}
	//拼接指令，从意图分析用户操作

	var (
		runRes []entitys.DeviceResponse //运行结果
		runErr error                    //控制设备过程产生的错误信息
	)
	//数据格式转换
	runDeviceIds := make([]string, 0)
	runDevice := make([]entitys.TmDevice, 0)
	for k, v := range runDeviceMap {
		runDeviceIds = append(runDeviceIds, k)
		runDevice = append(runDevice, v)
	}

	if isRead {
		//runRes
		//问题：pm2.5是多少？
		//{ pm25: 36 }   ==> 转换为   pm2.5 38
		//通过get_pm获取到功能名称pm25
		//pm2.5=38 转换为 小米的值pm2.5=380
		//输出：pm2.5为xxxx
		runRes, runErr = s.GetDevicePropertyCached(runDeviceIds, funcs, userId, false, xiaoaiVoiceCode)
		if runErr != nil {
			return "未能成功获取属性", nil
		}
		iotlogger.LogHelper.Helper.Info("获取属性-返回内容：", iotutil.ToString(runRes))
		msgs := []string{}
		for devId, d := range runDeviceMap {
			for _, r := range runRes {
				if devId == r.DeviceId {
					//d.DeviceName 获取的 xxx 为 xxx
					//toVoiceVal, err := ConvertVoiceToAxy(d, r.Data, xiaoaiVoiceCode)
					for k, v := range r.Data {
						msgs = append(msgs, fmt.Sprintf("%s获取%v为%v", d.DeviceName, k, v))
					}
				}
			}
		}
		if len(msgs) > 0 {
			msg = strings.Join(msgs, "、")
			return msg, nil
		} else {
			msg = "未获取到对应属性"
		}
	} else {
		runRes, runErr = s.SetDeviceProperty(runDevice, funcs, userId, false, xiaoaiVoiceCode)
		if runErr != nil {
			return "未能成功设置属性", nil
		}
		msg = "好的"
		hasErr := false
		hasSuccess := false
		for _, r := range runRes {
			if r.ErrorCode == "SUCCESS" {
				hasSuccess = true
			} else if r.ErrorCode == "DEVICE_OFFLINE" {
				hasErr = true
				msg = "设备离线，无法执行"
			} else {
				hasErr = true
			}
		}
		if hasSuccess && hasErr {
			msg = "部分设备执行成功"
		} else if hasSuccess && !hasErr {
			msg = "已执行"
		}
		iotlogger.LogHelper.Helper.Info("设备属性-返回内容：", iotutil.ToString(runRes))
	}
	return msg, nil
}

var (
	MSG_GET_SUCCESS   = "%s获取%v为%v"
	MSG_ERROR         = "异常"
	MSG_NO_DEVICE     = "未找到设备"
	MSG_NO_PROPERTIES = "未获取到对应属性"
	MSG_OFFLIINE      = "设备离线，无法执行"
	MSG_PART_SUCCESS  = "部分设备执行成功"
	MSG_SUCCESS       = "已执行"
	MSG_NO_SUCCESS    = "执行失败"
)

type MapMsg map[string]string

var mapMsg MapMsg

func (s MapMsg) set(devName, msg string) MapMsg {
	return map[string]string{devName: msg}
}

func (s XiaoaiIoTGateway) RunV2(req entitys.XiaoaiRequest, manager *manage.Manager, testUserId string) (string, error) {
	iotlogger.LogHelper.Helper.Debug("XiaoaiIoTGateway RunV2 session: ", req)
	var userId string
	if testUserId != "" {
		userId = testUserId
	} else {
		token := req.Session.User.AccessToken //  gjson.GetBytes(data, "session.user.access_token").String()
		ti, err := manager.LoadAccessToken(context.Background(), token)
		if err != nil {
			iotlogger.LogHelper.Helper.Error("XiaoaiIoTGateway RunV2 加载访问密钥错误: ", err)
			return "加载访问密钥错误", nil
		}

		userId = ti.GetUserID()
	}
	iotlogger.LogHelper.Helper.Debug("XiaoaiIoTGateway RunV2 userId: ", userId)

	var msg string = "暂不支持该操作" //返回默认
	slotInfo := req.Request.SlotInfo

	//欢迎（打开xxx）
	//if req.Request.SlotInfo.IntentName == "Mi_Welcome" {
	//    return "已打开", nil
	//}

	//通过前缀判断是否为获取属性
	isRead := strings.HasPrefix(req.Request.SlotInfo.IntentName, "get_")
	//slotInfo.IntentName 意图
	//slotInfo.Slots 插槽
	skillId := req.Request.Intent.AppID //技能Id
	//通过意图获取查询产品信息和其语控配置信息
	voiceInfos, err := common.GetDetailList(protosService.OpmVoiceProductDetailReq{SkillId: skillId, VoiceNo: xiaoaiVoiceCode})
	if err != nil {
		iotlogger.LogHelper.Helper.Error("XiaoaiIoTGateway RunV2 未获取到语控配置: ", err)
		return "未获取到语控配置", nil
	}
	msgs := make([]MapMsg, 0)
	for _, voiceInfo := range voiceInfos {
		iotlogger.LogHelper.Info("XiaoaiIoTGateway RunV2 Voice ProductKey:", voiceInfo.ProductKey)
		productKey := voiceInfo.ProductKey
		//找到设备Id
		devices, err := common.DiscoveryDevices(userId, xiaoaiVoiceCode)
		if err != nil {
			//msgs = append(msgs, MSG_ERROR)
			iotlogger.LogHelper.Info("XiaoaiIoTGateway RunV2 Voice ProductKey:", voiceInfo.ProductKey)
			continue
		}
		runDeviceMap := make(map[string]entitys.TmDevice)
		funcs := make(map[string]gjson.Result, 0)
		for _, device := range devices {
			//产品Key
			if device.ProductKey == productKey {
				hasName := false
				for _, slot := range slotInfo.Slots {
					switch slot.Name {
					case "deviceName":
						//设备名称
						if device.DeviceName == slot.RawValue {
							runDeviceMap[device.DeviceId] = device
						} else {
							delete(runDeviceMap, device.DeviceId)
						}
						hasName = true
					case "roomName":
						//房间名称
						if device.Zone == slot.RawValue {
							runDeviceMap[device.DeviceId] = device
						} else {
							delete(runDeviceMap, device.DeviceId)
						}
						hasName = true
					default:
						//从用户的意图配置中获取
						for _, item := range voiceInfo.AttrList {
							//将{亮度}设置为{50%}
							if slot.Name == item.VoiceCode {
								funcs[iotutil.ToString(item.VoiceCode)] = gjson.Parse("{\"value\": \"" + slot.RawValue + "\"}").Get("value")
								//item.AttrDpid  slot.RawValue
								//将50%转换为 max value * 50%
								// gjson.Parse(slot.RawValue) 转换为我司物模型
								//gjsonRes = gjson.Parse("{\"value\": \"" + slot.RawValue + "\"}").Get("value")
							}
						}
					}
				}
				//是否指定设备名称
				if !hasName {
					runDeviceMap[device.DeviceId] = device
				}
			}
		}
		if len(runDeviceMap) == 0 {
			continue
		}
		var (
			runRes []entitys.DeviceResponse //运行结果
			runErr error                    //控制设备过程产生的错误信息
		)
		//数据格式转换
		runDeviceIds := make([]string, 0)
		runDevice := make([]entitys.TmDevice, 0)
		for k, v := range runDeviceMap {
			runDeviceIds = append(runDeviceIds, k)
			runDevice = append(runDevice, v)
		}

		if isRead {
			//runRes
			//问题：pm2.5是多少？
			//{ pm25: 36 }   ==> 转换为   pm2.5 38
			//通过get_pm获取到功能名称pm25
			//pm2.5=38 转换为 小米的值pm2.5=380
			//输出：pm2.5为xxxx
			runRes, runErr = s.GetDevicePropertyCached(runDeviceIds, funcs, userId, false, xiaoaiVoiceCode)
			if runErr != nil {
				return "未能成功获取属性", nil
			}
			iotlogger.LogHelper.Helper.Info("XiaoaiIoTGateway RunV2 获取属性-返回内容：", iotutil.ToString(runRes))
			theMsgs := []string{}
			var devName string
			for devId, d := range runDeviceMap {
				for _, r := range runRes {
					if devId == r.DeviceId {
						//d.DeviceName 获取的 xxx 为 xxx
						//toVoiceVal, err := ConvertVoiceToAxy(d, r.Data, xiaoaiVoiceCode)
						for k, v := range r.Data {
							theMsgs = append(theMsgs, fmt.Sprintf(MSG_GET_SUCCESS, d.DeviceName, k, v))
						}
						devName = d.DeviceName
					}
				}
			}
			if len(theMsgs) > 0 {
				msgs = append(msgs, mapMsg.set(devName, strings.Join(theMsgs, ",")))
			} else {
				msgs = append(msgs, mapMsg.set(devName, MSG_NO_PROPERTIES))
			}
		} else {
			runRes, runErr = s.SetDeviceProperty(runDevice, funcs, userId, false, xiaoaiVoiceCode)
			if runErr != nil {
				return "未能成功设置属性", nil
			}
			msg = "好的"
			for _, r := range runRes {
				if r.ErrorCode == "SUCCESS" {
					msgs = append(msgs, mapMsg.set(r.DeviceName, MSG_SUCCESS))
				} else if r.ErrorCode == "DEVICE_OFFLINE" {
					msgs = append(msgs, mapMsg.set(r.DeviceName, MSG_OFFLIINE))
				}
			}
			iotlogger.LogHelper.Helper.Info("XiaoaiIoTGateway RunV2 设备属性-返回内容：", iotutil.ToString(runRes))
		}
	}
	if len(msgs) == 0 {
		msg = MSG_NO_DEVICE
	} else {
		for _, s := range msgs {
			var dn, dm string
			for k, v := range s {
				dn, dm = k, v
				break
			}
			msg += fmt.Sprintf("%s%s", dn, dm) + ";"
		}
		msg = strings.Trim(msg, ";")
	}

	return msg, nil
}

// 天猫精灵设置设备属性
func (s XiaoaiIoTGateway) SetDeviceProperty(devices []entitys.TmDevice, params map[string]gjson.Result, userId string, adjust bool, voiceNo string) ([]entitys.DeviceResponse, error) {
	var (
		channels    []string
		controlData []*entitys.ControlData
		devRespList []entitys.DeviceResponse
		devRespMap  = make(map[string]entitys.DeviceResponse)
		checked     = map[string]struct{}{}
		ctx         = context.Background()
	)
	// 将设备列表语控功能点数据转换为爱星云功能点数据
	for _, dev := range devices {
		_, ctl, err := common.VoiceToAxyFuncData(ctx, dev.DeviceId, params, adjust, voiceNo)
		if err != nil {
			// 需将设备错误信息放入响应列表
			iotlogger.LogHelper.Helper.Error("convert voice funcs to axy funs error: ", err)
			continue
		}
		if !dev.IsOnline {
			devRespMap[dev.DeviceId] = entitys.DeviceResponse{
				DeviceId:   dev.DeviceId,
				DeviceName: dev.DeviceName,
				ErrorCode:  "DEVICE_OFFLINE",
				Message:    "offline",
				Data:       ctl.Data,
			}
			devRespList = append(devRespList, devRespMap[dev.DeviceId])
			continue
		}
		controlData = append(controlData, ctl)
		channels = append(channels, persist.GetRedisKey(iotconst.HKEY_ACK_DATA_PUB_PREFIX+".%s.%s", ctl.ProductKey, ctl.DeviceId))
	}
	pubsub := cached.RedisStore.GetClient().Subscribe(ctx, channels...)
	defer pubsub.Close()
	// 开协程下发数据到设备
	for i := range controlData {
		devRespMap[controlData[i].DeviceId] = entitys.DeviceResponse{
			DeviceId: controlData[i].DeviceId,
			Data:     controlData[i].DataVoice,
		}
		iotlogger.LogHelper.Helper.Debug("control data[%v]: ", controlData[i].DeviceId, iotutil.ToString(controlData[i]))
		if _, err := common.PubControl(controlData[i].ProductKey, controlData[i].DeviceId, controlData[i].Data); err != nil {
			iotlogger.LogHelper.Helper.Errorf("pub control %s %s error: %v", controlData[i].ProductKey, controlData[i].DeviceId, err)
		}
	}
	iotlogger.LogHelper.Helper.Debug("control data: ", iotutil.ToString(controlData))
	iotlogger.LogHelper.Helper.Debug("channel data: ", iotutil.ToString(channels))
	// 订阅redis，延时等待一秒，若没有收到设备Ack消息，则直接返回错误信息给语控平台
	ctxNew, cancel := context.WithTimeout(ctx, 1500*time.Millisecond)
	defer cancel()
	for {
		select {
		case <-ctxNew.Done():
			// 超过一秒钟，需要将没有响应的设备补充错误信息
			devRespList = make([]entitys.DeviceResponse, 0)
			for _, dev := range devRespMap {
				if _, ok := checked[dev.DeviceId]; !ok {
					errorCode, message := "DEVICE_NOT_SUPPORT_FUNCTION", "device not support"
					if dev.ErrorCode != "" {
						errorCode, message = dev.ErrorCode, dev.Message
					}
					devRespList = append(devRespList, entitys.DeviceResponse{
						DeviceId:   dev.DeviceId,
						DeviceName: dev.DeviceName,
						ErrorCode:  errorCode,
						Message:    message,
						Data:       dev.Data,
					})
				} else {
					devRespList = append(devRespList, dev)
				}
			}
			iotlogger.LogHelper.Helper.Debug("control device data: ", devRespList)
			return devRespList, nil
		case msg := <-pubsub.Channel():
			devId := strings.Split(msg.Channel, ".")[2]
			devRespMap[devId] = entitys.DeviceResponse{
				DeviceId:   devId,
				DeviceName: devRespMap[devId].DeviceName,
				ErrorCode:  "SUCCESS",
				Message:    "SUCCESS",
				Data:       devRespMap[devId].Data,
			}
			devRespList = append(devRespList, devRespMap[devId])
			checked[devId] = struct{}{}
			iotlogger.LogHelper.Helper.Debug("control deivce success: ", devRespMap)
			if len(devices) == len(checked) {
				return devRespList, nil
			}
		}
	}
}

// 天猫精灵设置设备属性
func (s XiaoaiIoTGateway) GetDeviceProperty(devList []string, params map[string]gjson.Result, userId string, adjust bool, voiceNo string) ([]entitys.DeviceResponse, error) {
	var channels []string
	var controlData []*entitys.ControlData
	var devRespList []entitys.DeviceResponse
	checked := map[string]struct{}{}
	ctx := context.Background()
	// 将设备列表语控功能点数据转换为爱星云功能点数据
	for i := range devList {
		_, ctl, err := common.VoiceToAxyFuncData(ctx, devList[i], params, adjust, voiceNo)
		if err != nil {
			// 需将设备错误信息放入响应列表
			iotlogger.LogHelper.Helper.Error("convert voice funcs to axy funs error: ", err)
			continue
		}
		controlData = append(controlData, ctl)
		channels = append(channels, persist.GetRedisKey(iotconst.HKEY_QUERY_ACK_DATA_PUB_PREFIX+".%s.%s", ctl.ProductKey, ctl.DeviceId))
	}
	pubsub := cached.RedisStore.GetClient().Subscribe(ctx, channels...)
	defer pubsub.Close()
	// 开协程下发数据到设备
	for i := range controlData {
		properties := []string{}
		for k, _ := range controlData[i].Data {
			properties = append(properties, k)
		}
		iotlogger.LogHelper.Helper.Debug("control data[%v]: ", controlData[i].DeviceId, iotutil.ToString(properties))
		if _, err := common.PubQuery(controlData[i].ProductKey, controlData[i].DeviceId, properties); err != nil {
			iotlogger.LogHelper.Helper.Errorf("pub control %s %s error: %v", controlData[i].ProductKey, controlData[i].DeviceId, err)
		}
	}
	iotlogger.LogHelper.Helper.Debug("channel data: ", iotutil.ToString(channels))
	// 订阅redis，延时等待一秒，若没有收到设备Ack消息，则直接返回错误信息给语控平台
	ctxNew, cancel := context.WithTimeout(ctx, 1500*time.Millisecond)
	defer cancel()
	for {
		select {
		case <-ctxNew.Done():
			// 超过一秒钟，需要将没有响应的设备补充错误信息
			for i := range controlData {
				if _, ok := checked[controlData[i].DeviceId]; !ok {
					devRespList = append(devRespList, entitys.DeviceResponse{
						DeviceId:  controlData[i].DeviceId,
						ErrorCode: "DEVICE_NOT_SUPPORT_FUNCTION",
						Message:   "device not support",
					})
				}
			}
			iotlogger.LogHelper.Helper.Debug("control device data: ", devRespList)
			return devRespList, nil
		case msg := <-pubsub.Channel():
			devId := strings.Split(msg.Channel, ".")[2]

			info := iotstruct.DeviceRedisData{}
			err := json.Unmarshal([]byte(msg.Payload), &info)
			if err != nil {
				iotlogger.LogHelper.Errorf("解析推送信息内容失败,内容[%s],错误:%s", string(msg.Payload), err.Error())
				continue
			}

			data := info.Data.(iotprotocol.QueryAck)
			if data.Code != 0 {
				//属性获取失败
				devRespList = append(devRespList, entitys.DeviceResponse{
					DeviceId:  devId,
					ErrorCode: "DEVICE_NOT_SUPPORT_FUNCTION",
					Message:   "device not support",
				})
			} else {
				//属性获取成功
				devRespList = append(devRespList, entitys.DeviceResponse{
					DeviceId:  devId,
					ErrorCode: "SUCCESS",
					Message:   "SUCCESS",
					Data:      data.DeviceData,
				})
			}
			checked[devId] = struct{}{}
			iotlogger.LogHelper.Helper.Debug("control deivce success: ", devRespList)
			if len(devList) == len(checked) {
				return devRespList, nil
			}
		}
	}
}

func (s XiaoaiIoTGateway) GetDevicePropertyCached(devList []string, params map[string]gjson.Result, userId string, adjust bool, voiceNo string) ([]entitys.DeviceResponse, error) {
	ctx := context.Background()
	var devRespList []entitys.DeviceResponse
	// 将设备列表语控功能点数据转换为爱星云功能点数据
	for i := range devList {
		ctl, err := common.VoiceGetFuncData(ctx, devList[i], params, nil, adjust, voiceNo)
		if err != nil {
			// 需将设备错误信息放入响应列表
			iotlogger.LogHelper.Helper.Error("convert voice funcs to axy funs error: ", err)
			continue
		}
		devRespList = append(devRespList, entitys.DeviceResponse{
			DeviceId:  devList[i],
			ErrorCode: "SUCCESS",
			Message:   "SUCCESS",
			Data:      ctl.Data,
		})
	}
	return devRespList, nil
}
