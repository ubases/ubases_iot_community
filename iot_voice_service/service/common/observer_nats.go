package common

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotnatsjs"
	"cloud_platform/iot_device_service/config"
	"context"
	"errors"
	"sync"
	"time"

	"github.com/nats-io/nats.go/jetstream"

	"github.com/nats-io/nats.go"
)

var _suberonce sync.Once
var _subersingle *NatsSubscriber

func GetNatsSubscriber() *NatsSubscriber {
	_suberonce.Do(func() {
		suber, err := NewBuildSubscriber()
		if err != nil {
			panic(err)
		}
		_subersingle = suber
	})
	return _subersingle
}

type NatsSubscriber struct {
	suber      *iotnatsjs.JsClient
	concurrent int
	ctx        context.Context
	cancel     context.CancelFunc
	publisher  *Publisher
}

func NewBuildSubscriber() (*NatsSubscriber, error) {
	appName := "iot_voice_servicex"
	suber, err := iotnatsjs.NewJsClient(config.Global.Nats.Addrs...)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	err = suber.CreateOrUpdateConsumer(ctx, iotconst.NATS_STREAM_ORIGINAL_REDIS, []string{iotconst.HKEY_ACK_DATA_PUB_PREFIX + ".*"}, appName)
	if err != nil {
		cancel()
		return nil, err
	}
	Concurrent := 4
	return &NatsSubscriber{
		suber,
		Concurrent,
		ctx,
		cancel,
		NewPublisher(),
	}, nil
}

func (bs NatsSubscriber) GetPublisher() *Publisher {
	return bs.publisher
}

func (bs NatsSubscriber) Run() {
	// 从nats消息队列拉取数据，通过rpc接口，发布数据到iot_mqtt_service服务
	for i := 0; i < bs.concurrent; i++ {
		jsctx, err := bs.suber.Consume(bs.MessageHandler, bs.ErrorHandler)
		if err == nil {
			defer jsctx.Stop()
		}
	}
	for {
		select {
		case <-bs.ctx.Done():
			return
		}
	}
}

func (bs NatsSubscriber) Close() {
	bs.cancel()
	bs.suber.Close()
}

func (bs NatsSubscriber) MessageHandler(msg jetstream.Msg) {
	bs.publisher.Publish(EventData{
		Subject: msg.Subject(),
		Data:    string(msg.Data()),
	})
}

func (bs NatsSubscriber) ErrorHandler(consumeCtx jetstream.ConsumeContext, err error) {
	if errors.Is(err, nats.ErrConnectionClosed) {
		iotlogger.LogHelper.Errorf("拉取消息失败,原因:%s", err.Error())
		time.Sleep(3 * time.Second)
	} else if !errors.Is(err, nats.ErrTimeout) {
		iotlogger.LogHelper.Errorf("拉取消息失败,原因:%s", err.Error())
	}
}
