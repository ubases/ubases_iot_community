package iotnatsjs

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type JsClient struct {
	Nc       *nats.Conn
	Js       jetstream.JetStream
	Consumer jetstream.Consumer
}

func NewJsClient(addrs ...string) (*JsClient, error) {
	if len(addrs) == 0 {
		return nil, fmt.Errorf("addrs is empty")
	}
	nc, err := nats.Connect(addrs[0],
		nats.RetryOnFailedConnect(true),
		nats.ReconnectWait(5*time.Second),
	)
	if err != nil {
		return nil, err
	}
	js, err := jetstream.New(nc)
	if err != nil {
		return nil, err
	}
	return &JsClient{Nc: nc, Js: js}, nil
}

// NOTICE 部署专用,微服务禁止调用;否则未消费的消息会丢失!!!
func (o *JsClient) CreateOrUpdateStream(stream string, subject []string) error {
	err := o.Js.DeleteStream(context.Background(), stream)
	//if err != nil {
	//	return err
	//}
	s, err := o.Js.CreateOrUpdateStream(
		context.Background(),
		jetstream.StreamConfig{
			Name:     stream,
			Subjects: subject,
			MaxAge:   7 * 24 * time.Hour,
		},
	)
	fmt.Println(s)
	if err != nil {
		return err
	}
	return nil
}

func (o *JsClient) CreateOrUpdateConsumer(ctx context.Context, stream string, subject []string, Durable string) error {
	s, err := o.Js.Stream(ctx, stream)
	if err != nil {
		log.Fatal(err)
	}
	cons, err := s.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Durable:        Durable,
		AckPolicy:      jetstream.AckExplicitPolicy,
		FilterSubjects: subject,
	})
	if err != nil {
		return err
	}
	o.Consumer = cons
	return nil
}

func (o *JsClient) Consume(
	handler func(msg jetstream.Msg),
	errHandler func(consumeCtx jetstream.ConsumeContext, err error),
) (jetstream.ConsumeContext, error) {
	cc, err := o.Consumer.Consume(func(msg jetstream.Msg) {
		msg.Ack()
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()
		handler(msg)
	}, jetstream.ConsumeErrHandler(errHandler))
	if err != nil {
		return nil, err
	}
	return cc, nil
}

func (o *JsClient) ConsumeStop(ctx jetstream.ConsumeContext) {
	if ctx != nil {
		ctx.Stop()
	}
}

func (o *JsClient) Fetch(batch int) (jetstream.MessageBatch, error) {
	return o.Consumer.Fetch(batch)
}

func (o *JsClient) Publish(ctx context.Context, subject string, payload []byte) error {
	_, err := o.Js.Publish(ctx, subject, payload)
	return err
}

func (o *JsClient) Close() {
	if o.Nc != nil {
		o.Nc.Close()
	}
}
