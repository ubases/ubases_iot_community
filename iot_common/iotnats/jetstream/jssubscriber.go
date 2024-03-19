package jetstream

import (
	"context"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
)

type JSPullSubscriber struct {
	PullSubIface
}

//appname 用于nats管理，建议填写应用的名称
func NewJSPullSubscriber(appname string, stream string, subject string, errHandler nats.ConnErrHandler, addrs ...string) (*JSPullSubscriber, error) {
	cfg := Config{
		Ctx:         context.Background(),
		Name:        appname,
		ClusterURLs: strings.Join(addrs, ","), //addrs包含token
		StreamConfig: &nats.StreamConfig{
			Name:      stream,
			Retention: nats.WorkQueuePolicy,
			Subjects:  []string{stream + ".*"},
			MaxAge:    24 * 7 * time.Hour,
		},
		ConnErrHandler: errHandler,
	}
	s, err := NewPullSubscriber(&cfg, subject)
	if err != nil {
		return nil, err
	}
	return &JSPullSubscriber{s}, nil
}

//appname 用于nats管理，建议填写应用的名称
func NewJSPullArrSubscriber(appname string, streamName string, stream []string, subject string, errHandler nats.ConnErrHandler, addrs ...string) (*JSPullSubscriber, error) {
	cfg := Config{
		Ctx:         context.Background(),
		Name:        appname,
		ClusterURLs: strings.Join(addrs, ","), //addrs包含token
		StreamConfig: &nats.StreamConfig{
			Name:      streamName,
			Retention: nats.WorkQueuePolicy,
			Subjects:  stream,
			MaxAge:    24 * 7 * time.Hour,
		},
		ConnErrHandler: errHandler,
	}
	s, err := NewPullSubscriber(&cfg, subject)
	if err != nil {
		return nil, err
	}
	return &JSPullSubscriber{s}, nil
}

func (jspsub *JSPullSubscriber) FetchMessage() ([]*nats.Msg, error) {
	return jspsub.Fetch(5, 5*time.Second)
}

func (jspsub *JSPullSubscriber) FetchMessageEx(count int) ([]*nats.Msg, error) {
	return jspsub.Fetch(count, 5*time.Second)
}
