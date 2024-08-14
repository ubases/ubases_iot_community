package service

import (
	"cloud_platform/iot_app_user_service/config"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotnatsjs"
	"cloud_platform/iot_common/iotstruct"
	"context"
	"encoding/json"
	"strings"
	"github.com/nats-io/nats.go/jetstream"
)

func InitClearCachedSub(ctx context.Context) error {
	cachedCh := strings.Join([]string{iotconst.HKEY_CACHED_CLEAR_PUB_PREFIX, "*"}, ".")
	client, err := iotnatsjs.NewJsClient(config.Global.Nats.Addrs...)
	if err != nil {
		return err
	}

	err = client.CreateOrUpdateConsumer(ctx, iotconst.HKEY_CACHED_CLEAR_PUB_PREFIX, []string{cachedCh}, "clearCached")
	if err != nil {
		return err
	}

	jsctx, err := client.Consume(func(msg jetstream.Msg) {
		data := iotstruct.DeviceRedisUpdate{}
		if err := json.Unmarshal([]byte(msg.Data()), &data); err != nil {
			iotlogger.LogHelper.Helper.Error("json unmarshal cached clear error: ", err)
			return
		}
		if err := clearHomeCached(data); err != nil {
			iotlogger.LogHelper.Helper.Error("向天猫推送在线离线信息错误: ", err)
		}
	}, func(consumeCtx jetstream.ConsumeContext, err error) {})

	if err != nil {
		return err
	}
	defer jsctx.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil
		}
	}

}
