package iotnatsjs

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"context"
	"sync"
	"time"
)

type NatsPubData struct {
	Subject string
	Data    string
}

var _puberonce sync.Once
var _pubersingle *JsClientPub

func GetJsClientPub() *JsClientPub {
	_puberonce.Do(func() {
		ctx, cancel := context.WithCancel(context.Background())
		_pubersingle = &JsClientPub{
			queue:  iotutil.NewQueue(10240),
			ctx:    ctx,
			cancel: cancel,
		}
	})
	return _pubersingle
}

type JsClientPub struct {
	JsClient *JsClient
	queue    *iotutil.MlQueue
	ctx      context.Context
	cancel   context.CancelFunc
}

func (jpm *JsClientPub) InitJsClient(addrs []string) error {
	client, err := NewJsClient(addrs...)
	if err != nil {
		return err
	}
	jpm.JsClient = client
	return nil
}

func (jpm *JsClientPub) Run() {
	for {
		select {
		case <-jpm.ctx.Done():
		default:
			jpm.Handler()
			time.Sleep(time.Millisecond)
		}
	}
}

func (jpm *JsClientPub) Handler() {
	var err error
	for {
		m, ok, _ := jpm.queue.Get()
		if !ok {
			break
		}
		data, ok := m.(*NatsPubData)
		if !ok {
			continue
		}
		err = jpm.JsClient.Publish(jpm.ctx, data.Subject, []byte(data.Data))
		if err != nil {
			iotlogger.LogHelper.Errorf("JsClientPub.Handler:Publish error,%s", err.Error())
		}
	}
}

func (jpm *JsClientPub) PushData(data *NatsPubData) {
	jpm.queue.Put(data)
}

func (jpm *JsClientPub) Close() {
	jpm.cancel()
	jpm.Close()
}
