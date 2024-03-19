package jetstream

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"context"
	"log"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
)

type NatsPubData struct {
	Subject string
	Data    string //iotstruct.MqttToNatsData
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
	mapJsPublisher sync.Map //subject --> JsPublisher
	queue          *iotutil.MlQueue
	ctx            context.Context
	cancel         context.CancelFunc
}

func (jpm *JsPublisherMgr) AddPublisher(appName, natsStream, subject string, addrs []string) error {
	puber, err := NewJSPublisher(appName, natsStream, subject, connerrhandler, addrs...)
	if err != nil {
		return err
	}
	jpm.mapJsPublisher.Store(subject, puber)
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
		p, ok2 := jpm.mapJsPublisher.Load(data.Subject)
		if !ok2 {
			continue
		}
		puber := p.(*JSPublisher)
		if puber == nil {
			continue
		}
		err = puber.PublishEx([]byte(data.Data), handler)
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
	jpm.mapJsPublisher.Range(func(key, value interface{}) bool {
		jsp := value.(*JSPublisher)
		if jsp != nil {
			jsp.Close()
		}
		jpm.mapJsPublisher.Delete(key)
		return true
	})
}

func handler(msg *nats.Msg, err error) {
	if err != nil {
		log.Println(err.Error())
	}
}

func connerrhandler(conn *nats.Conn, err error) {
	if err != nil {
		iotlogger.LogHelper.Errorf("connerrhandler:nats连接错误,addr:%s,error:%s", conn.ConnectedAddr(), err.Error())
	}
}

func reconnectHandler(conn *nats.Conn) {
	iotlogger.LogHelper.Warnf("reconnectHandler:addr:%s", conn.ConnectedAddr())
}
