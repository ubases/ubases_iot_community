package cached

import (
	"cloud_platform/iot_common/iotgincache/persist"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_demo_api_service/config"
	"context"
)

var RedisStore *persist.RedisStoreEx

var (
	APP = "APP"
)

func InitCache() error {
	cli, err := iotredis.NewClient(config.Global.Redis)
	if err != nil {
		return err
	}
	RedisStore = persist.NewBatRedisStore(cli)
	return nil
}

func ClearCachedByKeys(ctx context.Context, keys ...string) {
	for _, key := range keys {
		iotredis.GetClient().Del(context.Background(), key)
	}
}
