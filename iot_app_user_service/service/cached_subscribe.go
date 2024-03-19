package service

import (
	"cloud_platform/iot_app_user_service/cached"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotstruct"
	"context"
	"encoding/json"
	"strings"
)

func InitRedisSub() {
	ctx := context.Background()
	cachedCh := strings.Join([]string{iotconst.HKEY_CACHED_CLEAR_PUB_PREFIX, "*"}, ".")
	cachedSub := cached.RedisStore.GetClient().PSubscribe(ctx, cachedCh)
	defer cachedSub.Close()
	cachedChannel := cachedSub.Channel()
	for {
		select {
		case msg := <-cachedChannel:
			data := iotstruct.DeviceRedisUpdate{}
			if err := json.Unmarshal([]byte(msg.Payload), &data); err != nil {
				iotlogger.LogHelper.Helper.Error("json unmarshal cached clear error: ", err)
				continue
			}
			if err := clearHomeCached(data); err != nil {
				iotlogger.LogHelper.Helper.Error("向天猫推送在线离线信息错误: ", err)
			}
		}
	}
}
