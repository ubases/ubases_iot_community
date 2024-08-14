package cached

import (
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_open_system_service/config"
)

func InitCache() error {
	_, err := iotredis.NewClient(config.Global.Redis)
	if err != nil {
		return err
	}
	return nil
}
