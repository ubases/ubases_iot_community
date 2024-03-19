package service

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_smart_speaker_service/cached"
	"cloud_platform/iot_smart_speaker_service/service/sync_update"
	"context"
	"encoding/json"
	"strings"
)

func InitRedisSub() {
	ctx := context.Background()
	onlineCh := strings.Join([]string{iotconst.HKEY_ONLINE_DATA_PUB_PREFIX, "*"}, ".")
	onlineSub := cached.RedisStore.GetClient().PSubscribe(ctx, onlineCh)
	defer onlineSub.Close()

	reportCh := strings.Join([]string{iotconst.HKEY_REPORT_DATA_PUB_PREFIX, "*"}, ".")
	reportSub := cached.RedisStore.GetClient().PSubscribe(ctx, reportCh)
	defer reportSub.Close()

	updateCh := strings.Join([]string{iotconst.HKEY_UPDATE_DATA_PUB_PREFIX, "*"}, ".")
	updateSub := cached.RedisStore.GetClient().PSubscribe(ctx, updateCh)
	defer updateSub.Close()

	onlineChannel := onlineSub.Channel()
	reportChannel := reportSub.Channel()
	updateChannel := updateSub.Channel()
	for {
		select {
		case msg := <-onlineChannel:
			data := iotstruct.DeviceRedisData{}
			if err := json.Unmarshal([]byte(msg.Payload), &data); err != nil {
				iotlogger.LogHelper.Helper.Error("json unmarshal online error: ", err)
				continue
			}
			svc := sync_update.OnlineDeviceSvc{Data: &data}
			if err := svc.OnlineDevice(); err != nil {
				iotlogger.LogHelper.Helper.Error("向天猫推送在线离线信息错误: ", err)
			}
		case msg := <-reportChannel:
			data := iotstruct.DeviceRedisData{}
			if err := json.Unmarshal([]byte(msg.Payload), &data); err != nil {
				iotlogger.LogHelper.Helper.Error("json unmarshal report error: ", err)
				continue
			}
			svc := sync_update.ReportDeviceSvc{Data: &data}
			if err := svc.ReportDevice(); err != nil {
				//iotlogger.LogHelper.Helper.Error("向天猫推送report报文错误: ", err)
			}
		case msg := <-updateChannel:
			data := iotstruct.DeviceRedisUpdate{}
			if err := json.Unmarshal([]byte(msg.Payload), &data); err != nil {
				iotlogger.LogHelper.Helper.Error("json unmarshal report error: ", err)
				continue
			}
			svc := sync_update.UpdateDeviceSvc{Data: &data}
			if err := svc.UpdateDevice(); err != nil {
				iotlogger.LogHelper.Helper.Error("更新天猫设备列表错误: ", err)
			}
		}
	}
}
