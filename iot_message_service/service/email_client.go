package service

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_message_service/config"
	"cloud_platform/iot_message_service/service/email"
	"context"
	"errors"
	"time"
)

var EmailMgr EmailClientMgr

type EmailClientMgr struct {
	cnf    *email.SMTPConfig
	ctx    context.Context
	cancel context.CancelFunc
	queue  *iotutil.MlQueue
}

func (s *EmailClientMgr) Init() {
	SMTP := config.Global.SMTP
	s.cnf = &email.SMTPConfig{
		Host:           SMTP.Host,
		Port:           SMTP.Port,
		Username:       SMTP.Username,
		Password:       SMTP.Password,
		ConnectTimeout: time.Duration(SMTP.ConnectTimeout) * time.Second,
		SendTimeout:    time.Duration(SMTP.SendTimeout) * time.Second,
		Helo:           SMTP.Helo,
		KeepAlive:      SMTP.KeepAlive,
		Exchange:       SMTP.Exchange,
		AuthType:       SMTP.AuthType,
		Ssl:            SMTP.Ssl,
		From:           SMTP.From,
	}
	ctx, cancel := context.WithCancel(context.Background())
	s.ctx = ctx
	s.cancel = cancel
	s.queue = iotutil.NewQueue(2048)
}

func (e *EmailClientMgr) Send(email *email.SendEmailInput) error {
	ok, quantity := e.queue.Put(email)
	if !ok {
		iotlogger.LogHelper.Errorf("EmailClientMgr.SendEmail 发送队列已满,当前排队数量:%d", quantity)
		return errors.New("发送队列已满,请稍后再试")
	}
	return nil
}

func (e *EmailClientMgr) StartQueueHandle() {
	for i := 0; i < 10; i++ {
		go e.QueueHandle()
	}
}

func (e *EmailClientMgr) QueueHandle() {
	defer func() {
		if err := recover(); err != nil {
			iotlogger.LogHelper.Errorf("EmailClientMgr.QueueHandle发生异常:%v", err)
			//重启携程
			go e.QueueHandle()
		}
	}()
	var err error
	client := e.NewEmailClient()
	for {
		select {
		case <-e.ctx.Done():
			if client != nil {
				client.Close()
			}
			return
		default:
			m, ok, _ := e.queue.Get()
			if !ok {
				//队列没有数据，则休眠
				time.Sleep(10 * time.Millisecond)
				continue
			}
			emailInput, ok1 := m.(*email.SendEmailInput)
			if !ok1 {
				continue
			}
			iotlogger.LogHelper.Infof("EmailClientMgr.QueueHandle: To=%s,Subject=%s", emailInput.To, emailInput.Subject)

			if client != nil {
				if err = client.Noop(); err != nil {
					client.Close()
					client = nil
				}
			}

			if client == nil {
				client = e.NewEmailClient()
				if client == nil {
					continue
				}
			}
			//成功立即返回，失败最多再试1次
			for i := 0; i < 2; i++ {
				if _, err = client.Send(*emailInput); err == nil {
					break
				} else {
					iotlogger.LogHelper.Errorf("EmailClientMgr.QueueHandle,Send error:%s, params: %s", err.Error(), iotutil.ToString(&emailInput))
					client.Close()
					client = e.NewEmailClient()
					if client == nil {
						break
					}
				}
			}
		}
	}
}

func (e *EmailClientMgr) NewEmailClient() *email.Client {
	client, err := email.NewClient(e.cnf)
	if err != nil {
		iotlogger.LogHelper.Errorf("EmailClientMgr.QueueHandle,NewClient error:%s", err.Error())
	}
	return client
}

func (e *EmailClientMgr) Close() {
	e.cancel()
}
