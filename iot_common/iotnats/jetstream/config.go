package jetstream

import (
	"context"

	"github.com/nats-io/nats.go"
)

type Config struct {
	Ctx            context.Context
	Name           string
	ClusterURLs    string
	StreamConfig   *nats.StreamConfig
	ConnErrHandler nats.ConnErrHandler //连接错误回调，用于排查问题
}
