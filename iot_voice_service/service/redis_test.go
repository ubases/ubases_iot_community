package service

//
//import (
//	"cloud_platform/iot_common/iotlogger"
//	"cloud_platform/iot_common/iotstruct"
//	"cloud_platform/iot_voice_service/service/sync_update"
//	"context"
//	"fmt"
//	"testing"
//	"time"
//
//	"github.com/go-redis/redis/v8"
//	"github.com/google/uuid"
//)
//
//func Test_pubsub(t *testing.T) {
//	client := redis.NewClient(&redis.Options{
//		Addr:     "120.77.64.252:6579",
//		Password: "test!@#%2022%",
//		DB:       0,
//	})
//
//	pubsub := client.PSubscribe(context.Background(), "report.*")
//	defer pubsub.Close()
//	for msg := range pubsub.Channel() {
//		fmt.Printf("channel=%s message=%s\n", msg.Channel, msg.Payload)
//	}
//}
//
//func Test_Online(t *testing.T) {
//	data := iotstruct.DeviceRedisData{
//		ProductKey: "l0U5SXfL",
//		DeviceId:   "GOjsLWaaOwOchD",
//		MessageId:  uuid.New().String(),
//		Time:       time.Now().Unix(),
//		Data: map[string]interface{}{
//			"onlineStatus": "online",
//		},
//	}
//	// svc := OnlineDeviceSvc{Data: &data}
//	// if err := svc.OnlineDevice(); err != nil {
//	// 	iotlogger.LogHelper.Helper.Error("向天猫推送在线离线信息错误: ", err)
//	// }
//
//	client := redis.NewClient(&redis.Options{
//		Addr:         "120.77.64.252:6579",
//		Password:     "test!@#%2022%",
//		DB:           0,
//		MinIdleConns: 5,
//		IdleTimeout:  600 * time.Second,
//		PoolSize:     50,
//		MaxConnAge:   3600 * time.Second,
//	})
//	if err := client.Publish(context.Background(), "online.l0U5SXfL.GOjsLWaaOwOchD", data).Err(); err != nil {
//		fmt.Println("发布在线信息错误: ", err)
//	}
//}
//
//func Test_Report(t *testing.T) {
//	data := iotstruct.DeviceRedisData{
//		ProductKey: "l0U5SXfL",
//		DeviceId:   "GOjsLWaaOwOchD",
//		MessageId:  uuid.New().String(),
//		Time:       time.Now().Unix(),
//		Data: map[string]interface{}{
//			"1":   false,
//			"101": 50,
//		},
//	}
//	svc := sync_update.ReportDeviceSvc{Data: &data}
//	if err := svc.ReportDevice(); err != nil {
//		iotlogger.LogHelper.Helper.Error("向天猫推送report报文错误: ", err)
//	}
//}
