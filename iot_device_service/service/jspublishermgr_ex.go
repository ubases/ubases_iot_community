package service

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotnats/jetstream"
	"cloud_platform/iot_common/iotutil"
	"context"
	"sync"
	"time"
)

type NatsPubDataJob struct {
	Subject string
	Data    string //iotstruct.MqttToNatsData
}

var _puberoncejob sync.Once
var _pubersinglejob *JsPublisherMgrJob

func GetJsPublisherMgrJob() *JsPublisherMgrJob {
	_puberoncejob.Do(func() {
		ctx, cancel := context.WithCancel(context.Background())
		_pubersinglejob = &JsPublisherMgrJob{
			queue:  iotutil.NewQueue(10240),
			ctx:    ctx,
			cancel: cancel,
		}
	})
	return _pubersinglejob
}

type JsPublisherMgrJob struct {
	mapJsPublisher sync.Map //subject --> JsPublisher
	queue          *iotutil.MlQueue
	ctx            context.Context
	cancel         context.CancelFunc
}

func (jpm *JsPublisherMgrJob) AddPublisher(subject string, addrs []string) error {
	appName := "iot_job_service"
	puber, err := jetstream.NewJSPublisher(appName, iotconst.NATS_STREAM_DEVICE, subject, connerrhandler, addrs...)
	if err != nil {
		return err
	}
	jpm.mapJsPublisher.Store(subject, puber)
	return nil
}

func (jpm *JsPublisherMgrJob) Run() {
	for {
		select {
		case <-jpm.ctx.Done():
		default:
			jpm.Handler()
			time.Sleep(10 * time.Microsecond)
		}
	}
}

func (jpm *JsPublisherMgrJob) Handler() {
	var err error
	for {
		m, ok, _ := jpm.queue.Get()
		if !ok {
			break
		}
		data, ok1 := m.(*NatsPubDataJob)
		if !ok1 {
			continue
		}
		p, ok2 := jpm.mapJsPublisher.Load(data.Subject)
		if !ok2 {
			continue
		}
		puber := p.(*jetstream.JSPublisher)
		if puber == nil {
			continue
		}
		err = puber.PublishEx([]byte(data.Data), handler)
		if err != nil {
			iotlogger.LogHelper.Errorf("JsPublisherMgrJob.Handler:PublishEx error,%s", err.Error())
		}
	}
}

func (jpm *JsPublisherMgrJob) PushData(data *NatsPubDataJob) {
	jpm.queue.Put(data)
}

func (jpm *JsPublisherMgrJob) Close() {
	jpm.cancel()
	jpm.mapJsPublisher.Range(func(key, value interface{}) bool {
		jsp := value.(*jetstream.JSPublisher)
		if jsp != nil {
			jsp.Close()
		}
		jpm.mapJsPublisher.Delete(key)
		return true
	})
}

//func handler(msg *nats.Msg, err error) {
//	if err != nil {
//		log.Println(err.Error())
//	}
//}
//
//func connerrhandler(conn *nats.Conn, err error) {
//	if err != nil {
//		iotlogger.LogHelper.Errorf("nats连接错误:%s", err.Error())
//	}
//}
