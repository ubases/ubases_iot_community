package cache

import (
	"cloud_platform/iot_weather_service/config"
	"cloud_platform/iot_common/iotredis"
	"strings"
)

var RepoCache CacheStore

func Init() error {
	if err := InitCache(); err != nil {
		return err
	}
	return nil
}

func InitCache() error {
	rdscnf := config.Global.Redis
	cnf := iotredis.Config{
		Cluster:      rdscnf.Cluster,
		Addrs:        strings.Join(rdscnf.Addrs, ","),
		Username:     rdscnf.Username,
		Password:     rdscnf.Password,
		Database:     rdscnf.Database,
		MinIdleConns: rdscnf.MinIdleConns,
		IdleTimeout:  rdscnf.IdleTimeout,
		PoolSize:     rdscnf.PoolSize,
		MaxConnAge:   rdscnf.MaxConnAge,
	}
	cli, err := iotredis.NewClient(cnf)
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
