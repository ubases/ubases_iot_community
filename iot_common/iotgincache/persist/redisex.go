package persist

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"

	"cloud_platform/iot_common/iotredis"
)

//集群版gin缓存

// RedisStoreEx store http response in redis
type RedisStoreEx struct {
	client iotredis.Client
}

// NewBatRedisStore create a redis memory store with redis client
func NewBatRedisStore(cli iotredis.Client) *RedisStoreEx {
	return &RedisStoreEx{
		client: cli,
	}
}

// Set put key value pair to redis, and expire after expireDuration
func (store *RedisStoreEx) Set(key string, value interface{}, expire time.Duration) error {
	payload, err := Serialize(value)
	if err != nil {
		return err
	}

	ctx := context.TODO()
	return store.client.Set(ctx, key, payload, expire).Err()
}

// Delete remove key in redis, do nothing if key doesn't exist
func (store *RedisStoreEx) Delete(key string) error {
	ctx := context.TODO()
	return store.client.Del(ctx, key).Err()
}

// Get get key in redis, if key doesn't exist, return ErrCacheMiss
func (store *RedisStoreEx) Get(key string, value interface{}) error {
	ctx := context.TODO()
	payload, err := store.client.Get(ctx, key).Bytes()

	if errors.Is(err, redis.Nil) {
		return ErrCacheMiss
	}

	if err != nil {
		return err
	}

	return Deserialize(payload, value)
}

// Set put multi hash key value pair to redis
func (store *RedisStoreEx) HMSet(key string, value interface{}) error {
	ctx := context.TODO()
	return store.client.HMSet(ctx, key, value).Err()
}

// Get get key in redis hash, if key doesn't exist, return ""
func (store *RedisStoreEx) HGetCodeMsg(key, field string) string {
	ctx := context.TODO()
	msg, err := store.client.HGet(ctx, key, field).Result()
	if err != nil {
		fmt.Println("get code msg from redis hash error: ", err)
		return err.Error()
	}
	return msg
}

func (store *RedisStoreEx) TxPipeline() redis.Pipeliner {
	return store.client.TxPipeline()
}

func (store *RedisStoreEx) Pipeline() redis.Pipeliner {
	return store.client.Pipeline()
}

func (store *RedisStoreEx) Keys(ctx context.Context, pattern string) []string {
	return store.client.Keys(ctx, pattern).Val()
}

func GetRedisKey(tpl string, params ...interface{}) string {
	return fmt.Sprintf(tpl, params...)
}

func (store *RedisStoreEx) SubscribeNum(ctx context.Context, pattern string) map[string]int64 {
	return store.client.PubSubNumSub(ctx, pattern, pattern).Val()
}

func (store *RedisStoreEx) GetClient() iotredis.Client {
	return store.client
}
