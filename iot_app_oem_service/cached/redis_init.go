package cached

import (
	"cloud_platform/iot_app_oem_service/config"
	"cloud_platform/iot_common/iotredis"
)

func InitCache() error {
	_, err := iotredis.NewClient(config.Global.Redis)
	if err != nil {
		return err
	}

	return nil
}
