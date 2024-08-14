package cached

import (
	"cloud_platform/iot_common/iotgincache/persist"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_voice_service/config"
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
