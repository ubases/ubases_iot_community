package cached

import (
	"cloud_platform/iot_device_service/config"
	"cloud_platform/iot_common/iotgincache/persist"
	"cloud_platform/iot_common/iotredis"
	"strings"
)

var RedisStore *persist.RedisStoreEx

var (
	DEVICE_STATUS = "DS_"
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

	//err = RedisStore.Set("test0001", cnf, time.Hour*100)
	//fmt.Println(err)
	//cnf2 := iotredis.Config{}
	//err = RedisStore.Get("test0001", &cnf2)
	//fmt.Println(err)
	return nil
}
