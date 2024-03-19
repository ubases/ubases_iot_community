package iotredis

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

var _rclient Client

type Cmdable interface {
	Subscribe(ctx context.Context, channels ...string) *redis.PubSub
	PSubscribe(ctx context.Context, channels ...string) *redis.PubSub
	redis.Cmdable
}

type Client interface {
	Cmdable
	Close() error
}

type Config struct {
	Cluster      bool
	Addrs        string //多个地址，逗号分割
	Username     string
	Password     string
	Database     int
	MinIdleConns int // 最小空闲连接
	IdleTimeout  int // 空闲时间
	PoolSize     int // 连接池大小
	MaxConnAge   int // 连接最大可用时间
}

func NewClient(conf Config) (Client, error) {
	config := conf
	ctx := context.Background()
	hostMembers := strings.Split(config.Addrs, ",")

	if conf.MinIdleConns == 0 {
		conf.MinIdleConns = 2
	}
	if conf.IdleTimeout == 0 || conf.IdleTimeout > 1800 {
		conf.IdleTimeout = 1800
	}
	if conf.PoolSize == 0 {
		conf.PoolSize = 10
	}
	if conf.MaxConnAge == 0 || conf.MaxConnAge > 3600 {
		conf.MaxConnAge = 3600
	}
	if len(hostMembers) <= 1 && !config.Cluster { //单机模式
		rdb := redis.NewClient(&redis.Options{
			Addr:         config.Addrs,
			Username:     config.Username,
			Password:     config.Password,
			DB:           config.Database,
			MinIdleConns: config.MinIdleConns,
			IdleTimeout:  time.Second * time.Duration(config.IdleTimeout),
			PoolSize:     config.PoolSize,
			MaxConnAge:   time.Second * time.Duration(config.MaxConnAge),
		})
		_, err := rdb.Ping(ctx).Result()
		if err != nil {
			return nil, err
		}
		_rclient = rdb
		return rdb, nil
	}
	rdb := redis.NewClusterClient(&redis.ClusterOptions{ //集群模式
		Addrs:        hostMembers,
		Username:     config.Username,
		Password:     config.Password,
		MinIdleConns: config.MinIdleConns,
		IdleTimeout:  time.Second * time.Duration(config.IdleTimeout),
		PoolSize:     config.PoolSize,
		MaxConnAge:   time.Second * time.Duration(config.MaxConnAge),
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	_rclient = rdb
	return rdb, nil
}

func GetClient() Client {
	if _rclient == nil {
		panic(errors.New("请先调用iotredis.NewClient创建Client"))
	}
	return _rclient
}

func Ping() error {
	return GetClient().Ping(context.Background()).Err()
}
