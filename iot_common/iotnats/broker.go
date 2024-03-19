package iotnats

import (
	micro_nats "github.com/asim/go-micro/plugins/broker/nats/v4"
	"github.com/nats-io/nats.go"
	"go-micro.dev/v4/broker"
)

type BatBorker struct {
	broker.Broker
	subs map[string]broker.Subscriber
}

func NewBatBorker(token string, addrs ...string) (*BatBorker, error) {
	nopts := nats.GetDefaultOptions()
	nopts.Servers = addrs
	nopts.Token = token
	natsBroker := micro_nats.NewBroker(micro_nats.Options(nopts))
	if err := natsBroker.Init(); err != nil {
		return nil, err
	}
	if err := natsBroker.Connect(); err != nil {
		return nil, err
	}
	return &BatBorker{natsBroker, make(map[string]broker.Subscriber)}, nil
}

func (bb *BatBorker) SubscribeEx(topic string, h broker.Handler, opts ...broker.SubscribeOption) error {
	sub, err := bb.Subscribe(topic, h, opts...)
	if err != nil {
		return err
	}
	bb.subs[topic] = sub
	return nil
}

func (bb *BatBorker) SharedSubscribe(topic, queue string, h broker.Handler) error {
	return bb.SubscribeEx(topic, h, broker.Queue(queue))
}

func (bb *BatBorker) UnSubscribe(topic string) error {
	if sub, ok := bb.subs[topic]; ok {
		err := sub.Unsubscribe()
		if err == nil {
			delete(bb.subs, topic)
		}
	}
	return nil
}

func (bb *BatBorker) Close(unSub bool) error {
	if unSub {
		for _, v := range bb.subs {
			_ = v.Unsubscribe()
		}
	}
	bb.subs = nil
	return bb.Disconnect()
}
