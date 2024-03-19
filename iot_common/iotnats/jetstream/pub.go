package jetstream

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/nats-io/nats.go"
)

const (
	PubAckDur = 30 * time.Second
)

type PubIface interface {
	Publish(data []byte, ackif interface{}) error
	PublishSubject(subject string, data []byte, ackif interface{}) error
	Close()
	GetJetStreamContext() nats.JetStreamContext
}

type Publisher struct {
	js JetStreamIface
	publishSubject string
}

func (jp *Publisher) Close() {
	if jp.js != nil {
		jp.js.Close()
	}
}

func (jp *Publisher) GetJetStreamContext() nats.JetStreamContext {
	return jp.js.GetJetStreamContext()
}

func SubjectMatchesStreamConfig(cfg *Config, publishSubject string) bool {
	sok := false
	for _, s := range cfg.StreamConfig.Subjects {
		if sok, _ = regexp.MatchString(s, publishSubject); sok {
			break
		}
	}
	return sok
}

func NewJetStreamPublisher(cfg *Config, publishSubject string) (PubIface, error) {
	if !SubjectMatchesStreamConfig(cfg, publishSubject) {
		return nil,fmt.Errorf("publish subject \"%s\" does not match stream config %v", publishSubject, cfg.StreamConfig.Subjects)
	}
	j, err := NewJetStream(cfg)
	if err != nil {
		return nil, err
	}

	jp := &Publisher{js: j, publishSubject: publishSubject}
	return jp, nil
}

type PubAck struct {
	Msg *nats.Msg
	Err error
}

type PubAckHandler func(msg *nats.Msg, err error)

func (jp *Publisher) publish(subject string, data []byte, ackIf interface{}) error {
	var ackcb PubAckHandler
	var ackchannel chan *PubAck

	if ackIf != nil {
		switch v := ackIf.(type) {
		case func(msg *nats.Msg, err error):
			ackcb = v
		case chan *PubAck:
			ackchannel = v
		default:
			return  fmt.Errorf("ackif unsupported type %T. Expecting func or channel.", v)
		}
	}

	f, err := jp.js.GetJetStreamContext().PublishAsync(subject, data)
	if err != nil {
		return err
	}

	go func() {
		var err error
		ctx, cancel := context.WithTimeout(context.Background(), PubAckDur)
		defer cancel()
		select {
		case err = <-f.Err():
		case <-f.Ok():
		case <-ctx.Done():
			if ctx.Err() != nil {
				err = ctx.Err()
			}
		}
		if ackchannel != nil {
			ackchannel <- &PubAck{
				Msg: f.Msg(),
				Err: err,
			}
		} else if ackcb != nil {
			ackcb(f.Msg(), err)
		}
	}()

	return nil
}

func (jp *Publisher) Publish(data []byte, ackIf interface{}) error {
	return jp.publish(jp.publishSubject, data, ackIf)
}

func (jp *Publisher) PublishSubject(subject string, data []byte, ackIf interface{}) error {
	return jp.publish(subject, data, ackIf)
}
