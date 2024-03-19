package jetstream

import (
	"context"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
)

type JSPublisher struct {
	PubIface
}

// appname 用于nats管理，建议填写应用的名称
func NewJSPublisher(appname, stream string, subject string, errHandler nats.ConnErrHandler, addrs ...string) (*JSPublisher, error) {
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
	p, err := NewJetStreamPublisher(&cfg, subject)
	if err != nil {
		return nil, err
	}

	return &JSPublisher{p}, nil
}

func (jspub *JSPublisher) PublishEx(buf []byte, f func(msg *nats.Msg, err error)) error {
	return jspub.Publish(buf, f)
}

func (jspub *JSPublisher) PublishSubjectEx(subject string, buf []byte, f func(msg *nats.Msg, err error)) error {
	return jspub.PublishSubject(subject, buf, f)
}
