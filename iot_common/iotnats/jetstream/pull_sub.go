package jetstream

import (
	"errors"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

type PullSubIface interface {
	Fetch(count int, timeout time.Duration) ([]*nats.Msg, error)
	Close()
	GetJetStreamContext() nats.JetStreamContext
}

type PullSubscriber struct {
	js  JetStreamIface
	sub *nats.Subscription
}

func (jsub *PullSubscriber) Close() {
	if jsub.js != nil {
		jsub.js.Close()
	}
}

func (jsub *PullSubscriber) GetJetStreamContext() nats.JetStreamContext {
	return jsub.js.GetJetStreamContext()
}

func (jsub *PullSubscriber) Fetch(count int, timeout time.Duration) ([]*nats.Msg, error) {
	msgs, err := jsub.sub.Fetch(count, nats.MaxWait(timeout))
	if err != nil {
		if errors.Is(err, nats.ErrConnectionClosed) {
			time.Sleep(10 * time.Millisecond)
		}
		return nil, err
	}

	for _, m := range msgs {
		_ = m.Ack()
	}
	return msgs, nil
}

func NewPullSubscriber(cfg *Config, subject string) (PullSubIface, error) {
	j, err := NewJetStream(cfg)
	if err != nil {
		return nil, err
	}

	opts := []nats.SubOpt{
		nats.BindStream(cfg.StreamConfig.Name),
		nats.DeliverAll(),
		nats.AckExplicit(),
	}

	sub, err := j.GetJetStreamContext().PullSubscribe(subject, cfg.Name, opts...)
	if err != nil {
		j.Close()
		if err.Error() == nats.ErrPullSubscribeToPushConsumer.Error() {
			err = fmt.Errorf("%w; possibly an existing push-based consumer exists. Try `nats con info %s %s`", err, cfg.StreamConfig.Name, cfg.Name)
		}
		return nil, err
	}

	jsub := &PullSubscriber{
		js:  j,
		sub: sub,
	}
	return jsub, nil
}
