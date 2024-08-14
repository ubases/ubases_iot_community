package cached

import (
	"cloud_platform/iot_common/iotgincache/persist"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_intelligence_service/config"
)

var (
	DEVICE_STATUS = "DS_"
)

var RedisStore *persist.RedisStoreEx

func InitCache() error {
	cli, err := iotredis.NewClient(config.Global.Redis)
	if err != nil {
		return err
	}

	RedisStore = persist.NewBatRedisStore(cli)

	return nil
}
