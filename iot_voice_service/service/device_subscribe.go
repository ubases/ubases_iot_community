package service

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotnatsjs"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_voice_service/config"
	"cloud_platform/iot_voice_service/service/sync_update"
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

//func InitReportSub() {
//	ctx := context.Background()
//	onlineCh := strings.Join([]string{iotconst.HKEY_ONLINE_DATA_PUB_PREFIX, "*"}, ".")
//	onlineSub := cached.RedisStore.GetClient().PSubscribe(ctx, onlineCh)
//	defer onlineSub.Close()
//
//	reportCh := strings.Join([]string{iotconst.HKEY_REPORT_DATA_PUB_PREFIX, "*"}, ".")
//	reportSub := cached.RedisStore.GetClient().PSubscribe(ctx, reportCh)
//	defer reportSub.Close()
//
//	updateCh := strings.Join([]string{iotconst.HKEY_UPDATE_DATA_PUB_PREFIX, "*"}, ".")
//	updateSub := cached.RedisStore.GetClient().PSubscribe(ctx, updateCh)
//	defer updateSub.Close()
//
//	onlineChannel := onlineSub.Channel()
//	reportChannel := reportSub.Channel()
//	updateChannel := updateSub.Channel()
//	for {
//		select {
//		case msg := <-onlineChannel:
//			data := iotstruct.DeviceRedisData{}
//			if err := json.Unmarshal([]byte(msg.Payload), &data); err != nil {
//				iotlogger.LogHelper.Helper.Error("json unmarshal online error: ", err)
//				continue
//			}
//			svc := sync_update.OnlineDeviceSvc{Data: &data}
//			if err := svc.OnlineDevice(); err != nil {
//				iotlogger.LogHelper.Helper.Error("向天猫推送在线离线信息错误: ", err)
//			}
//		case msg := <-reportChannel:
//			data := iotstruct.DeviceRedisData{}
//			if err := json.Unmarshal([]byte(msg.Payload), &data); err != nil {
//				iotlogger.LogHelper.Helper.Error("json unmarshal report error: ", err)
//				continue
//			}
//			svc := sync_update.ReportDeviceSvc{Data: &data}
//			if err := svc.ReportDevice(); err != nil {
//				//iotlogger.LogHelper.Helper.Error("向天猫推送report报文错误: ", err)
//			}
//		case msg := <-updateChannel:
//			data := iotstruct.DeviceRedisUpdate{}
//			if err := json.Unmarshal([]byte(msg.Payload), &data); err != nil {
//				iotlogger.LogHelper.Helper.Error("json unmarshal report error: ", err)
//				continue
//			}
//			svc := sync_update.UpdateDeviceSvc{Data: &data}
//			if err := svc.UpdateDevice(); err != nil {
//				iotlogger.LogHelper.Helper.Error("更新天猫设备列表错误: ", err)
//			}
//		}
//	}
//}

func InitReportSub(ctx context.Context) {
	client, err := iotnatsjs.NewJsClient(config.Global.Nats.Addrs...)
	if err != nil {
		panic(err)
		return
	}
	defer client.Close()

	//onlineCh := strings.Join([]string{iotconst.HKEY_ONLINE_DATA_PUB_PREFIX, "*"}, ".")
	//reportCh := strings.Join([]string{iotconst.HKEY_REPORT_DATA_PUB_PREFIX, "*"}, ".")
	//updateCh := strings.Join([]string{iotconst.HKEY_UPDATE_DATA_PUB_PREFIX, "*"}, ".")
	onlineCh := iotconst.HKEY_ONLINE_DATA_PUB_PREFIX + ".>"
	reportCh := iotconst.HKEY_REPORT_DATA_PUB_PREFIX + ".>"
	updateCh := iotconst.HKEY_UPDATE_DATA_PUB_PREFIX + ".>"
	err = client.CreateOrUpdateConsumer(ctx, iotconst.NATS_STREAM_ORIGINAL_REDIS, []string{onlineCh, reportCh, updateCh}, "iot_smart_speak_service")
	if err != nil {
		panic(err)
		return
	}

	jsctx, err := client.Consume(MessageHandler, ErrorHandler)
	if err != nil {
		return
	}
	defer jsctx.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}

func MessageHandler(msg jetstream.Msg) {
	info := iotstruct.MqttToNatsData{}
	err := json.Unmarshal(msg.Data(), &info)
	if err != nil {
		iotlogger.LogHelper.Errorf("解析激活失败,内容[%s],错误:%s", string(msg.Data()), err.Error())
		return
	}
	subject := msg.Subject()
	subprefix := strings.Split(msg.Subject(), ".")
	if len(subprefix) >= 3 {
		subject = subprefix[0] + "." + subprefix[1]
	}
	switch subject {
	case iotconst.HKEY_ONLINE_DATA_PUB_PREFIX:
		data := iotstruct.DeviceRedisData{}
		if err := json.Unmarshal(msg.Data(), &data); err != nil {
			iotlogger.LogHelper.Helper.Error("json unmarshal online error: ", err)
			return
		}
		svc := sync_update.OnlineDeviceSvc{Data: &data}
		if err := svc.OnlineDevice(); err != nil {
			iotlogger.LogHelper.Helper.Error("向天猫推送在线离线信息错误: ", err)
		}
	case iotconst.HKEY_REPORT_DATA_PUB_PREFIX:
		data := iotstruct.DeviceRedisData{}
		if err := json.Unmarshal(msg.Data(), &data); err != nil {
			iotlogger.LogHelper.Helper.Error("json unmarshal report error: ", err)
			return
		}
		svc := sync_update.ReportDeviceSvc{Data: &data}
		if err := svc.ReportDevice(); err != nil {
			//iotlogger.LogHelper.Helper.Error("向天猫推送report报文错误: ", err)
		}
	case iotconst.HKEY_UPDATE_DATA_PUB_PREFIX:
		data := iotstruct.DeviceRedisUpdate{}
		if err := json.Unmarshal(msg.Data(), &data); err != nil {
			iotlogger.LogHelper.Helper.Error("json unmarshal report error: ", err)
			return
		}
		svc := sync_update.UpdateDeviceSvc{Data: &data}
		if err := svc.UpdateDevice(); err != nil {
			iotlogger.LogHelper.Helper.Error("更新天猫设备列表错误: ", err)
		}
	}
}

func ErrorHandler(consumeCtx jetstream.ConsumeContext, err error) {
	if errors.Is(err, nats.ErrConnectionClosed) {
		iotlogger.LogHelper.Errorf("拉取消息失败,原因:%s", err.Error())
		time.Sleep(3 * time.Second)
	} else if !errors.Is(err, nats.ErrTimeout) {
		iotlogger.LogHelper.Errorf("拉取消息失败,原因:%s", err.Error())
	}
}
