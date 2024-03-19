package iotredis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type Subscriber struct {
	pubsub   *redis.PubSub
	channel  string
	callback func(*redis.Message)
}

func NewSubscriber(channel string, callback func(*redis.Message)) *Subscriber {
	return &Subscriber{
		pubsub:   nil,
		channel:  channel,
		callback: callback,
	}
}

func (o *Subscriber) Subscribe() error {
	o.pubsub = GetClient().Subscribe(context.Background(), o.channel)
	go func() {
		ch := o.pubsub.Channel()
		for msg := range ch {
			o._recover(o.callback, msg)
		}
	}()
	return nil
}

func (o *Subscriber) Close() error {
	if o.pubsub != nil {
		return o.pubsub.Close()
	}
	return nil
}

func (o *Subscriber) _recover(f func(*redis.Message), args *redis.Message) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	f(args)
}

func Publish(channel string, message interface{}) error {
	return GetClient().Publish(context.Background(), channel, message).Err()
}
