package execute

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_intelligence_service/cached"
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
			onlineStatus := deviceInfo["onlineStatus"] //onlineStatusCms.Val()
			if err != nil {
				return saveTaskResult(resultObj, 2, "异常", db)
			} else if onlineStatus != "online" {
				return saveTaskResult(resultObj, 3, "设备离线", db)
			} else {
				var functions []*Function
				err := json.Unmarshal([]byte(taskItem.Functions), &functions)
				if err == nil && len(functions) > 0 {
					//tlsMap := make(map[string]string)
					//proTslCmd := iotredis.GetClient().HGetAll(context2.Background(), iotconst.HKEY_PRODUCT_DATA+productKey)
					//if proTslCmd.Err() == nil {
					//	tlsMap = proTslCmd.Val()
					//}
					//{
					//	"custom": 0,
					//	"dataSpecs": "{\"custom\":0,\"dataType\":\"FLOAT\",\"min\":0.1,\"max\":3,\"step\":0.1,\"multiple\":10,\"unit\":\"\"}",
					//	"dataSpecsList": "",
					//	"dataType": "FLOAT",
					//	"dpid": 3,
					//	"identifier": "spray_mode_stepless",
					//	"name": "档位无极调节",
					//	"rwFlag": "READ_WRITE"
					//}
					pubControls := make(map[string]string)
					for _, function := range functions {
						val := function.FuncValue
						//设置倍数值
						//val = setMultiple(tlsMap, function.FuncValue, function.FuncKey)
						pubControls[function.FuncKey] = val
					}
					//isOnline := onlineStatus == "online"
					////调用MQTT服务，推送消息
					//iotlogger.LogHelper.WithTag("id", iotutil.ToString(intelligenceId)).WithTag("method", "DeviceExecute").Info("设备在线状态为", isOnline)

					msgId, _, err := PubControl(productKey, devId, iotutil.MapStringToInterface(pubControls))
					if err != nil {
						iotlogger.LogHelper.WithTag("id", iotutil.ToString(intelligenceId)).WithTag("method", "DeviceExcute").Error(err.Error())
						return saveTaskResult(resultObj, 2, "执行失败", db)
					} else {
						iotlogger.LogHelper.WithTag("id", iotutil.ToString(intelligenceId)).WithTag("method", "DeviceExcute").Error("开始等待控制结果上报")
						controlAckChan := make(chan bool, 0)
						go CheckControlResult(msgId, productKey, devId, controlAckChan)
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

func CheckControlResult(messageId, productKey, deviceId string, result chan bool) {
	ctx := context.Background()
	//TestDeviceChan()
	ackCh := strings.Join([]string{iotconst.HKEY_ACK_DATA_PUB_PREFIX, productKey, deviceId}, ".")
	ackSub := cached.RedisStore.GetClient().PSubscribe(ctx, ackCh)
	defer ackSub.Close()

	ackChannel := ackSub.Channel()
	for {
		select {
		case msg := <-ackChannel:
			data := iotstruct.DeviceRedisData{}
			if err := json.Unmarshal([]byte(msg.Payload), &data); err != nil {
				iotlogger.LogHelper.Helper.Error("json unmarshal online error: ", err)
				continue
			}
			iotlogger.LogHelper.Infof("redis control ack sub data[%s,%s] ", data.MessageId, messageId, iotutil.ToString(data))
			//if data.MessageId == messageId {
			result <- true
			break
			//}
		case <-time.After(2 * time.Second): //超时3s
			iotlogger.LogHelper.Info("redis control ack sub timeout")
			result <- false
			break
		}
	}
}
