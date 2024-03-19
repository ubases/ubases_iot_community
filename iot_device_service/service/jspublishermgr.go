package service

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotnats/jetstream"
	"cloud_platform/iot_common/iotutil"
	"context"
	"log"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
)

type NatsPubData struct {
	Subject string
	Data    string
}

var _puberonce sync.Once
var _pubersingle *JsPublisherMgr

func GetJsPublisherMgr() *JsPublisherMgr {
	_puberonce.Do(func() {
		ctx, cancel := context.WithCancel(context.Background())
		_pubersingle = &JsPublisherMgr{
			queue:  iotutil.NewQueue(10240),
			ctx:    ctx,
			cancel: cancel,
		}
	})
	return _pubersingle
}

type JsPublisherMgr struct {
	jsPublisher *jetstream.JSPublisher
	queue       *iotutil.MlQueue
	ctx         context.Context
	cancel      context.CancelFunc
}

func (jpm *JsPublisherMgr) AddPublisher(stream string, subject string, addrs []string) error {
	puber, err := jetstream.NewJSPublisher("iot_device_service", stream, subject, pubconnerrhandler, addrs...)
	if err != nil {
		return err
	}
	jpm.jsPublisher = puber
	return nil
}

func (jpm *JsPublisherMgr) AddPublisherEx(appName, stream string, subject string, addrs []string) error {
	puber, err := jetstream.NewJSPublisher(appName, stream, subject, pubconnerrhandler, addrs...)
	if err != nil {
		return err
	}
	jpm.jsPublisher = puber
	return nil
}

func (jpm *JsPublisherMgr) Run() {
	for {
		select {
		case <-jpm.ctx.Done():
		default:
			jpm.Handler()
			time.Sleep(10 * time.Microsecond)
		}
	}
}

func (jpm *JsPublisherMgr) Handler() {
	var err error
	for {
		m, ok, _ := jpm.queue.Get()
		if !ok {
			break
		}
		data, ok1 := m.(*NatsPubData)
		if !ok1 {
			continue
		}
		err = jpm.jsPublisher.PublishEx([]byte(data.Data), handler)
		if err != nil {
			iotlogger.LogHelper.Errorf("JsPublisherMgr.Handler:PublishEx error,%s", err.Error())
		}
	}
}

func (jpm *JsPublisherMgr) PushData(data *NatsPubData) {
	jpm.queue.Put(data)
}

func (jpm *JsPublisherMgr) Close() {
	jpm.cancel()
	jpm.jsPublisher.Close()
}

func handler(msg *nats.Msg, err error) {
	if err != nil {
		log.Println(err.Error())
	}
}

func pubconnerrhandler(conn *nats.Conn, err error) {
	if err != nil {
		iotlogger.LogHelper.Errorf("nats连接错误:%s", err.Error())
	}
}
