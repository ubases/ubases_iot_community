package cache

import (
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_weather_service/config"
)

var RepoCache CacheStore

func Init() error {
	if err := InitCache(); err != nil {
		return err
	}
	return nil
}

func InitCache() error {
	cli, err := iotredis.NewClient(config.Global.Redis)
	if err != nil {
		return err
	}
	cache, err := NewMemStore(NewRedisStore(cli))
	if err != nil {
		return err
	}
	RepoCache = cache
	return nil
}
