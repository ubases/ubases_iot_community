package iotredis

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var _rclient Client

type Cmdable interface {
	redis.Cmdable
}

type Client interface {
	Cmdable
	//Close() error
}

type Config struct {
	Cluster      bool
	Addrs        []string //多个地址，逗号分割
	Username     string
	Password     string
	Database     int
	MinIdleConns int // 最小空闲连接
	IdleTimeout  int // 空闲时间
	PoolSize     int // 连接池大小
	MaxConnAge   int // 连接最大可用时间
}

func NewClient(conf Config) (Client, error) {
	client, err := _NewClient(conf, false)
	if err != nil {
		return nil, err
	}
	_rclient = client
	return client, nil
}

//func NewPubSubClient(conf Config) (Client, error) {
//	client, err := _NewClient(conf, true)
//	if err != nil {
//		return nil, err
//	}
//	return client, nil
//}
func _NewClient(conf Config, pubsub bool) (Client, error) {
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
	if conf.Cluster {
		return NewRedisClusterClient(conf, pubsub)
	}
	return NewRedisClient(conf, pubsub)
}
func NewRedisClient(conf Config, pubsub bool) (Client, error) {
	opt := &redis.Options{
		Addr:         conf.Addrs[0],
		Username:     conf.Username,
		Password:     conf.Password,
		DB:           conf.Database,
		MinIdleConns: conf.MinIdleConns,
		IdleTimeout:  time.Second * time.Duration(conf.IdleTimeout),
		PoolSize:     conf.PoolSize,
		MaxConnAge:   time.Second * time.Duration(conf.MaxConnAge),
	}
	if pubsub {
		opt.MinIdleConns = 0
		opt.MaxConnAge = 0
		opt.MaxRetries = -1
		opt.DialTimeout = 10 * time.Second
		opt.ReadTimeout = 30 * time.Second
		opt.WriteTimeout = 30 * time.Second
		opt.OnConnect = func(ctx context.Context, cn *redis.Conn) error {
			clientId, err := cn.ClientID(ctx).Result()
			log.Println("clientId:", clientId)
			return err
		}
	}
	rdb := redis.NewClient(opt)
	_, err := rdb.Ping(context.Background()).Result()
		if err != nil {
			return nil, err
		}
		return rdb, nil
	}

func NewRedisClusterClient(conf Config, pubsub bool) (Client, error) {
	opt := &redis.ClusterOptions{
		Addrs:        conf.Addrs,
		Username:     conf.Username,
		Password:     conf.Password,
		MinIdleConns: conf.MinIdleConns,
		IdleTimeout:  time.Second * time.Duration(conf.IdleTimeout),
		PoolSize:     conf.PoolSize,
		MaxConnAge:   time.Second * time.Duration(conf.MaxConnAge),
	}
	if pubsub {
		opt.MinIdleConns = 0
		opt.MaxConnAge = 0
		opt.MaxRetries = -1
		opt.DialTimeout = 10 * time.Second
		opt.ReadTimeout = 30 * time.Second
		opt.WriteTimeout = 30 * time.Second
		opt.OnConnect = func(ctx context.Context, cn *redis.Conn) error {
			clientId, err := cn.ClientID(ctx).Result()
			log.Println("clientId:", clientId)
			return err
		}
	}
	rdb := redis.NewClusterClient(opt)
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return rdb, nil
}

func GetClient() Client {
	if _rclient == nil {
		panic(errors.New("请先调用iotredis.NewClient创建Client"))
	}
	return _rclient
}

