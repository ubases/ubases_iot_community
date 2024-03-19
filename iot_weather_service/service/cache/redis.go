package cache

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/proto"

	"cloud_platform/iot_common/iotredis"
)

type RedisStore struct {
	client iotredis.Client
}

func NewRedisStore(cli iotredis.Client) *RedisStore {
	return &RedisStore{
		client: cli,
	}
}
func (store *RedisStore) Set(key string, m proto.Message) error {
	payload, err := Serialize(m)
	if err != nil {
		return err
	}
	ctx := context.TODO()
	return store.client.Set(ctx, key, payload, 0).Err()
}
func (store *RedisStore) Get(key string, m proto.Message) error {
	ctx := context.TODO()
	payload, err := store.client.Get(ctx, key).Bytes()
	if errors.Is(err, redis.Nil) {
		return ErrCacheMiss
	}
	if err != nil {
		return err
	}
	return Deserialize(payload, m)
}
