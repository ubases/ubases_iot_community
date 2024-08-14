package execute

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_common/iotutil"
	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_device/model"
	"cloud_platform/iot_proto/protos/protosService"
	context2 "context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"

	"golang.org/x/net/context"
)

import (
	"cloud_platform/iot_common/iotnatsjs"
	"cloud_platform/iot_intelligence_service/config"
)

// 执行任务功能
type Function struct {
	FuncCompare    int32  `json:"funcCompare"`
	FuncKey        string `json:"funcKey"`
	FuncIdentifier string `json:"funcIdentifier"`
	FuncValue      string `json:"funcValue"`
	FuncDesc       string `json:"funcDesc"`
}

func saveTaskResult(resultObj *model.TSceneIntelligenceResultTask, status int32, errMsg string, res *gorm.DB) (*model.TSceneIntelligenceResultTask, error) {
	resultObj.ResultMsg = errMsg
	resultObj.IsSuccess = status
	resultObj.EndTime = time.Now()
	return resultObj, res.Save(resultObj).Error
}

// DeviceExecute 设备执行
func DeviceExecute(userId, resultId int64, runtime int64, devIds *[]string, taskItem *protosService.SceneIntelligenceTask) (*model.TSceneIntelligenceResultTask, error) {
	var (
		devId                 = taskItem.ObjectId //设备Id
		intelligenceId        = taskItem.IntelligenceId
		resMsg                = iotconst.RUN_CONTINUTE
		runStatus      int32  = 0  //状态 =1 执行中
		productKey     string = "" //产品Key
	)
	*devIds = append(*devIds, devId) //记录设备
	//写入进行中
	resultObj := &model.TSceneIntelligenceResultTask{
		Id:             iotutil.GetNextSeqInt64(),
		IntelligenceId: taskItem.IntelligenceId,
		TaskId:         taskItem.Id,
		IsSuccess:      0,
		ResultId:       resultId,
		ResultMsg:      resMsg,
		Functions:      taskItem.Functions,
		TaskImg:        taskItem.TaskImg,
		TaskDesc:       taskItem.TaskDesc,
		TaskType:       taskItem.TaskType,
		FuncDesc:       taskItem.FuncDesc,
		FuncKey:        taskItem.FuncKey,
		FuncValue:      taskItem.FuncValue,
		ObjectId:       taskItem.ObjectId,
		ProductKey:     taskItem.ProductKey,
		StartTime:      time.Unix(runtime, 0),
	}
	db := iotmodel.GetDB()
	err := db.Save(resultObj).Error
	if err != nil {
		return saveTaskResult(resultObj, 2, err.Error(), db)
	}

	//获取设备信息
	deviceCmd := iotredis.GetClient().HGetAll(context2.Background(), iotconst.HKEY_DEV_DATA_PREFIX+devId) //, "productKey"
	if deviceCmd.Err() != nil {
		return saveTaskResult(resultObj, 2, "异常 device cached err", db)
	} else {
		deviceInfo := deviceCmd.Val()
		productKey = deviceInfo["productKey"]
		devUserId := deviceInfo["userId"]
		if devUserId != iotutil.ToString(userId) {
			return saveTaskResult(resultObj, 2, "设备已经不是你的了", db)
		}
		if productKey == "" {
			iotlogger.LogHelper.WithTag("id", iotutil.ToString(intelligenceId)).WithTag("method", "DeviceExcute").Errorf("产品缓存信息获取异常， 设备Id：%s", taskItem.ObjectId)
			return saveTaskResult(resultObj, 2, "异常 ProductKey is empty", db)
		} else {
			onlineStatus := deviceInfo["onlineStatus"]
			if err != nil {
				return saveTaskResult(resultObj, 2, "异常", db)
			} else if onlineStatus != "online" {
				return saveTaskResult(resultObj, 3, "设备离线", db)
			} else {
				var functions []*Function
				err := json.Unmarshal([]byte(taskItem.Functions), &functions)
				if err == nil && len(functions) > 0 {
					pubControls := make(map[string]string)
					for _, function := range functions {
						val := function.FuncValue
						pubControls[function.FuncKey] = val
					}
					msgId, _, err := PubControl(productKey, devId, iotutil.MapStringToInterface(pubControls))
					if err != nil {
						iotlogger.LogHelper.WithTag("id", iotutil.ToString(intelligenceId)).WithTag("method", "DeviceExcute").Error(err.Error())
						return saveTaskResult(resultObj, 2, "执行失败", db)
					} else {
						iotlogger.LogHelper.WithTag("id", iotutil.ToString(intelligenceId)).WithTag("method", "DeviceExcute").Error("开始等待控制结果上报")
						controlAckChan := make(chan bool, 0)
						go CheckControlResult(msgId, productKey, devId, pubControls, controlAckChan)
						select {
						case msg := <-controlAckChan:
							if msg {
								iotlogger.LogHelper.WithTag("id", iotutil.ToString(intelligenceId)).WithTag("method", "DeviceExcute").Error("执行成功")
								runStatus, resMsg = 1, "执行成功"
							} else {
								iotlogger.LogHelper.WithTag("id", iotutil.ToString(intelligenceId)).WithTag("method", "DeviceExcute").Error("执行失败")
								runStatus, resMsg = 2, "执行失败"
							}
						}
						return saveTaskResult(resultObj, runStatus, resMsg, db)
					}
				} else {
					iotlogger.LogHelper.WithTag("id", iotutil.ToString(intelligenceId)).WithTag("method", "DeviceExcute").Error("设备执行异常，属性参数未设置")
					return saveTaskResult(resultObj, 2, "设备指令为空", db)
				}
			}
		}
	}
	return saveTaskResult(resultObj, 2, "异常", db)
}

func setMultiple(tlsMap map[string]string, funcVal, funcKey string) string {
	//设置倍数值
	if tlsVal, tlsOk := tlsMap[fmt.Sprintf("tls_%v", funcKey)]; tlsOk {
		tlsInfo, err := iotutil.JsonToMapErr(tlsVal)
		if err == nil {
			dataType := iotutil.ToString(tlsInfo["dataType"])
			if tlsInfo != nil && dataType == "FLOAT" {
				dataSpecs, err := iotutil.JsonToMapErr(tlsVal)
				if err == nil {
					multiple, err := iotutil.ToFloat64Err(dataSpecs["multiple"])
					if err == nil {
						funcVal = iotutil.ToString(iotutil.ToInt64(iotutil.ToFloat64(funcVal) * multiple))
					}
				}
			}
		}
	}
	return funcVal
}

// CheckControlResult 检查设备控制结果
//func CheckControlResult(messageId, productKey, deviceId string, pubControls map[string]string, result chan bool) {
//	ackCh := strings.Join([]string{iotconst.HKEY_ACK_DATA_PUB_PREFIX, productKey, deviceId}, ".")
//	client, err := iotredis.NewRedisPubSubClient(config.Global.Redis)
//	if err != nil {
//		panic(err)
//		return
//	}
//	client.PSubscribeSync(ackCh)
//	ackSub := client.PSubscribeSync(ackCh)
//	defer func() {
//		ackSub.Close()
//		client.Close()
//	}()
//	ackChannel := ackSub.Channel()
//	for {
//		select {
//		case msg := <-ackChannel:
//			data := iotstruct.DeviceRedisData{}
//			if err := json.Unmarshal([]byte(msg.Payload), &data); err != nil {
//				iotlogger.LogHelper.Helper.Error("json unmarshal online error: ", err)
//				continue
//			}
//			iotlogger.LogHelper.Infof("redis control ack sub data[%s,%s] ", data.MessageId, messageId, iotutil.ToString(data))
//			result <- true
//			return
//		case <-time.After(3 * time.Second): //超时3s
//			deviceCmd := iotredis.GetClient().HGetAll(context2.Background(), iotconst.HKEY_DEV_DATA_PREFIX+deviceId)
//			isallcharge := false
//			if deviceCmd.Err() == nil {
//				deviceInfo := deviceCmd.Val()
//				for k, v := range pubControls {
//					if dv, ok := deviceInfo[k]; ok && dv == v {
//						isallcharge = true
//					} else {
//						isallcharge = false
//					}
//				}
//			}
//			if isallcharge {
//				iotlogger.LogHelper.Info("redis control ack sub timeout, but it was judged to be successful.")
//				result <- false
//			} else {
//				iotlogger.LogHelper.Info("redis control ack sub timeout")
//				result <- false
//			}
//			return
//		}
//	}
//}

func CheckControlResult(messageId, productKey, deviceId string, pubControls map[string]string, result chan bool) {
	ackCh := strings.Join([]string{iotconst.HKEY_ACK_DATA_PUB_PREFIX, productKey, deviceId}, ".")
	client, err := iotnatsjs.NewJsClient(config.Global.Nats.Addrs...)
	if err != nil {
		return
	}
	err = client.CreateOrUpdateConsumer(context.Background(), iotconst.NATS_STREAM_ORIGINAL_REDIS, []string{ackCh}, "CheckControlResult"+deviceId)
	if err != nil {
		return
	}

	defer client.Close()

	ackChannel, err := client.Fetch(1)
	if err != nil {
		return
	}

	if ackChannel.Error() != nil {
		return
	}

	for {
		select {
		case msg := <-ackChannel.Messages():
			data := iotstruct.DeviceRedisData{}
			if err := json.Unmarshal([]byte(msg.Data()), &data); err != nil {
				iotlogger.LogHelper.Helper.Error("json unmarshal online error: ", err)
				continue
			}
			iotlogger.LogHelper.Infof("redis control ack sub data[%s,%s] ", data.MessageId, messageId, iotutil.ToString(data))
			result <- true
			return
		case <-time.After(3 * time.Second): //超时3s
			deviceCmd := iotredis.GetClient().HGetAll(context2.Background(), iotconst.HKEY_DEV_DATA_PREFIX+deviceId)
			isallcharge := false
			if deviceCmd.Err() == nil {
				deviceInfo := deviceCmd.Val()
				for k, v := range pubControls {
					if dv, ok := deviceInfo[k]; ok && dv == v {
						isallcharge = true
					} else {
						isallcharge = false
					}
				}
			}
			if isallcharge {
				iotlogger.LogHelper.Info("redis control ack sub timeout, but it was judged to be successful.")
				result <- false
			} else {
				iotlogger.LogHelper.Info("redis control ack sub timeout")
				result <- false
			}
			return
		}
	}
}
