package cache

import (
	"errors"
	"github.com/golang/protobuf/proto"
)

var ErrCacheMiss = errors.New("persist cache miss error")

type CacheStore interface {
	Get(key string, m proto.Message) error
	Set(key string, m proto.Message) error
}

func Serialize(m proto.Message) ([]byte, error) {
	buf, err := proto.Marshal(m)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func Deserialize(byt []byte, m proto.Message) (err error) {
	return proto.Unmarshal(byt, m)
}
