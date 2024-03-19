package cache

import (
	"github.com/golang/protobuf/proto"
)

type MemStore struct {
	repo CacheStore
}

func NewMemStore(repo CacheStore) (*MemStore, error) {
	return &MemStore{repo: repo /*, cache: cache*/}, nil
}

func (mem *MemStore) Set(key string, m proto.Message) error {
	err := mem.repo.Set(key, m)
	if err != nil {
		return err
	}
	return nil
}

func (mem *MemStore) Get(key string, m proto.Message) error {
	err := mem.repo.Get(key, m)
	if err != nil {
		return err
	}
	return nil
}
