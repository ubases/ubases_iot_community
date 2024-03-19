package push

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_message_service/service/push/pushModel"
	"context"
	"errors"
	"time"
)

var PushMgr PushClientMgr

type PushInfo struct {
	InputTarget pushModel.MessageTarget
	Message     pushModel.MessageRequestModel
}

type PushClientMgr struct {
	provider   string
	pushClient PushClient
	ctx        context.Context
	cancel     context.CancelFunc
	queue      *iotutil.MlQueue
}

func (s *PushClientMgr) Init() error {
	var err error
	s.pushClient, err = NewPushClient(All)
	ctx, cancel := context.WithCancel(context.Background())
	s.ctx = ctx
	s.cancel = cancel
	s.queue = iotutil.NewQueue(2048)
	return err
}

func (s *PushClientMgr) PushPush(data interface{}) error {
	ok, quantity := s.queue.Put(data)
	if !ok {
		iotlogger.LogHelper.Errorf("PushClientMgr.SendSMS 发送队列已满,当前排队数量:%d", quantity)
		return errors.New("发送队列已满,请稍后再试")
	}
	return nil
}

func (s *PushClientMgr) QueueHandle() {
	defer func() {
		if err := recover(); err != nil {
			iotlogger.LogHelper.Errorf("PushClientMgr.QueueHandle发生异常:%v", err)
			time.Sleep(2 * time.Second)
			//重启携程
			go s.QueueHandle()
		}
	}()
	var err error
	for {
		select {
		case <-s.ctx.Done():
			return
		default:
			m, ok, _ := s.queue.Get()
			if ok {
				if pushInfo, ok1 := m.(PushInfo); ok1 {
					iotlogger.LogHelper.Infof("PushClientMgr.data: template=%v", iotutil.ToString(pushInfo))
					err = s.pushClient.PushMessage(pushInfo.InputTarget, pushInfo.Message)
					if err != nil {
						iotlogger.LogHelper.Errorf("PushClientMgr.SendPush error:%s", err.Error())
					}
				}
			}
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func (s *PushClientMgr) Close() {
	s.cancel()
}
