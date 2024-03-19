package cached

import (
	"cloud_platform/iot_demo_api_service/config"
	"cloud_platform/iot_common/iotgincache/persist"
	"cloud_platform/iot_common/iotredis"
	"context"
	"strings"
)

var RedisStore *persist.RedisStoreEx

var (
	APP = "APP"
)

func InitCache() error {
	addrs := strings.Join(config.Global.Redis.Addrs, ",")
	cnf := iotredis.Config{
		Cluster:      config.Global.Redis.Cluster,
		Addrs:        addrs,
		Username:     config.Global.Redis.Username,
		Password:     config.Global.Redis.Password,
		Database:     config.Global.Redis.Database,
		MinIdleConns: config.Global.Redis.MinIdleConns,
		IdleTimeout:  config.Global.Redis.IdleTimeout,
		PoolSize:     config.Global.Redis.PoolSize,
		MaxConnAge:   config.Global.Redis.MaxConnAge,
	}
	cli, err := iotredis.NewClient(cnf)
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
