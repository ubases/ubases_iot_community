package cached

import (
	"cloud_platform/iot_system_service/config"
	"cloud_platform/iot_common/iotredis"
	"strings"
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
	_, err := iotredis.NewClient(cnf)
	if err != nil {
		return err
	}
	return nil
}
