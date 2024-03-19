package service

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_message_service/config"
	"cloud_platform/iot_message_service/service/sms"
	"context"
	"errors"
	"time"
)

var SmdMgr SMSClientMgr

type SmsInfo struct {
	Template          string
	Param             map[string]string
	TargetPhoneNumber []string
}

type SMSClientMgr struct {
	provider  string
	smsClient sms.SmsClient
	ctx       context.Context
	cancel    context.CancelFunc
	queue     *iotutil.MlQueue
}

func (s *SMSClientMgr) Init() error {
	var err error
	cnf := config.Global.SMS
	s.smsClient, err = sms.NewSmsClient(cnf.Provider, cnf.AccessId, cnf.AccessKey, cnf.Sign, cnf.Other...)
	s.provider = cnf.Provider
	ctx, cancel := context.WithCancel(context.Background())
	s.ctx = ctx
	s.cancel = cancel
	s.queue = iotutil.NewQueue(2048)
	return err
}

func (s *SMSClientMgr) SendSMS(data interface{}, template string, targetPhoneNumber ...string) error {
	var param map[string]string
	param = VariablesToMap(data)
	smsinfo := SmsInfo{
		Template:          template,
		Param:             param,
		TargetPhoneNumber: targetPhoneNumber,
	}
	ok, quantity := s.queue.Put(smsinfo)
	if !ok {
		iotlogger.LogHelper.Errorf("SMSClientMgr.SendSMS 发送队列已满,当前排队数量:%d", quantity)
		return errors.New("发送队列已满,请稍后再试")
	}
	return nil
}

func (s *SMSClientMgr) QueueHandle() {
	defer func() {
		if err := recover(); err != nil {
			iotlogger.LogHelper.Errorf("SMSClientMgr.QueueHandle发生异常:%v", err)
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
				if smsinfo, ok1 := m.(SmsInfo); ok1 {
					iotlogger.LogHelper.Infof("SMSClientMgr.SendMessage: template=%s,param=%v,phone=%v",
						smsinfo.Template, smsinfo.Param, smsinfo.TargetPhoneNumber)
					err = s.smsClient.SendMessage(smsinfo.Template, smsinfo.Param, smsinfo.TargetPhoneNumber...)
					if err != nil {
						iotlogger.LogHelper.Errorf("SMSClientMgr.SendMessage error:%s", err.Error())
					}
				}
			}
			time.Sleep(10 * time.Millisecond)
		}
	}
}
func (s *SMSClientMgr) Close() {
	s.cancel()
}
