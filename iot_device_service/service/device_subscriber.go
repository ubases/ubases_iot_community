package service

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotnatsjs"
	"cloud_platform/iot_common/iotprotocol"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_device_service/config"
	"context"
	"errors"
	"strings"
	"time"

	"github.com/nats-io/nats.go/jetstream"

	"github.com/nats-io/nats.go"
)

type Subscriber struct {
	Subscribers []*DeviceSubscriber
}

func InitSubscriber() (subOp *Subscriber, err error) {
	subOp = new(Subscriber)
	subOp.SetSubscriber(strings.Join([]string{"iot_device_service", "device", "active"}, "_"),
		iotconst.NATS_STREAM_DEVICE,
		iotconst.NATS_SUBJECT_INFO+".>")
	subOp.SetSubscriber(strings.Join([]string{"iot_device_service", "device", "online"}, "_"),
		iotconst.NATS_STREAM_DEVICE,
		iotconst.NATS_SUBJECT_ONLINE+".>")
	subOp.SetSubscriber(strings.Join([]string{"iot_device_service", "ota", "upgrade"}, "_"),
		iotconst.NATS_STREAM_DEVICE,
		iotconst.NATS_SUBJECT_UPGRADE_REPORT+".>")
	subOp.SetSubscriber(strings.Join([]string{"iot_device_service", "device", "control"}, "_"),
		iotconst.NATS_STREAM_DEVICE,
		iotconst.NATS_SUBJECT_CONTROL+".>")
	subOp.SetSubscriber(strings.Join([]string{"iot_device_service", "control", "ack"}, "_"),
		iotconst.NATS_STREAM_DEVICE,
		iotconst.NATS_SUBJECT_CONTROL_ACK+".>")
	subOp.SetSubscriber(strings.Join([]string{"iot_device_service", "device", "report"}, "_"),
		iotconst.NATS_STREAM_DEVICE,
		iotconst.NATS_SUBJECT_REPORT+".>")
	return
}

func (s *Subscriber) SetSubscriber(appName, stream, subject string) {
	//TODO 临时增加panic
	defer iotutil.PanicHandler()
	sub, err := NewDeviceSubscriber(appName, stream, subject)
	if err != nil {
		iotlogger.LogHelper.Errorf("创建订阅服务错误:", err.Error(), appName, stream, subject)
		return
	}
	s.Subscribers = append(s.Subscribers, sub)
}

func (s *Subscriber) RunSub() {
	for _, subscriberMap := range s.Subscribers {
		go subscriberMap.Run()
	}
}

func (s *Subscriber) CloseSub() {
	for _, subscriberMap := range s.Subscribers {
		go subscriberMap.Close()
	}
}

type DeviceSubscriber struct {
	//suber      *jetstream.JSPullSubscriber
	suber      *iotnatsjs.JsClient
	concurrent int
	ctx        context.Context
	cancel     context.CancelFunc
}

func NewDeviceSubscriber(appname string, stream string, subject string) (*DeviceSubscriber, error) {
	//suber, err := jetstream.NewJSPullSubscriber(appname, stream, subject, connerrhandler, config.Global.Nats.Addrs...)
	//if err != nil {
	//	return nil, err
	//}
	//ctx, cancel := context.WithCancel(context.Background())
	suber, err := iotnatsjs.NewJsClient(config.Global.Nats.Addrs...)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	err = suber.CreateOrUpdateConsumer(ctx, stream, []string{subject}, appname)
	if err != nil {
		cancel()
		return nil, err
	}
	Concurrent := 1
	return &DeviceSubscriber{suber, Concurrent, ctx, cancel}, nil
}

func connerrhandler(conn *nats.Conn, err error) {
	if err != nil {
		iotlogger.LogHelper.Errorf("nats连接错误:%s", err.Error())
	}
}

func (bs DeviceSubscriber) Run() {
	jsctx, err := bs.suber.Consume(MessageHandler, ErrorHandler)
	if err != nil {
		return
	}
	defer jsctx.Stop()
	for {
		select {
		case <-bs.ctx.Done():
			return
		}
	}
	//for {
	//	if bs.ctx.Err() != nil {
	//		break
	//	}
	//	msgList, err := bs.suber.FetchMessageEx(1)
	//	if err != nil {
	//		if errors.Is(err, nats.ErrConnectionClosed) {
	//			iotlogger.LogHelper.Errorf("拉取消息失败,原因:%s", err.Error())
	//			time.Sleep(3 * time.Second)
	//		} else if !errors.Is(err, nats.ErrTimeout) {
	//			iotlogger.LogHelper.Errorf("拉取消息失败,原因:%s", err.Error())
	//		}
	//		continue
	//	}
	//	for _, v := range msgList {
	//		info := iotstruct.MqttToNatsData{}
	//		err = json.Unmarshal(v.Data, &info)
	//		if err != nil {
	//			iotlogger.LogHelper.Errorf("解析激活失败,内容[%s],错误:%s", string(v.Data), err.Error())
	//			continue
	//		}
	//		switch v.Subject {
	//		case iotconst.NATS_SUBJECT_INFO:
	//			//激活逻辑
	//			iotlogger.LogHelper.Info("收到激活信息", string(v.Data))
	//			svc := ActiveDeviceSvc{Data: info}
	//			svc.ActiveDevice()
	//		case iotconst.NATS_SUBJECT_ONLINE:
	//			iotlogger.LogHelper.Info("设备上线消息", string(v.Data))
	//			svc := OnlineDeviceSvc{Data: &info}
	//			svc.OnlineDevice()
	//		case iotconst.NATS_SUBJECT_UPGRADE_REPORT:
	//			iotlogger.LogHelper.Info("设备OTA升级", string(v.Data))
	//			svc := OtaUpgradeDeviceSvc{Data: &info}
	//			svc.OtaUpgradeDevice()
	//		case iotconst.NATS_SUBJECT_CONTROL:
	//			svc := ReportDeviceSvc{Data: &info}
	//			//直接接收control上报
	//			iotlogger.LogHelper.Info("Control下发", string(v.Data))
	//			svc.ReportDevice()
	//		case iotconst.NATS_SUBJECT_REPORT:
	//			iotlogger.LogHelper.Info("控制上报", string(v.Data))
	//			svc := ReportDeviceSvc{Data: &info}
	//			//直接接收control上报
	//			if info.Name == iotprotocol.TP_C_CONTROL {
	//				//记录设备信息
	//				svc.ReportAckDevice()
	//				svc.ReportMsgDevice()
	//			} else {
	//				iotlogger.LogHelper.Info("设备定时上报", string(v.Data))
	//			}
	//			svc.ReportDeviceFault()
	//		}
	//	}
	//}
}

func (bs DeviceSubscriber) Close() {
	bs.cancel()
	bs.suber.Close()
}

func MessageHandler(msg jetstream.Msg) {
	info := iotstruct.MqttToNatsData{}
	err := json.Unmarshal(msg.Data(), &info)
	if err != nil {
		iotlogger.LogHelper.Errorf("解析激活失败,内容[%s],错误:%s", string(msg.Data()), err.Error())
		return
	}
	subject := msg.Subject()
	subprefix := strings.Split(msg.Subject(), ".")
	if len(subprefix) >= 3 {
		subject = subprefix[0] + "." + subprefix[1]
	}
	switch subject {
	case iotconst.NATS_SUBJECT_INFO:
		//激活逻辑
		iotlogger.LogHelper.Info("收到激活信息", string(msg.Data()))
		svc := ActiveDeviceSvc{Data: info}
		svc.ActiveDevice()
	case iotconst.NATS_SUBJECT_ONLINE:
		iotlogger.LogHelper.Info("设备上线消息", string(msg.Data()))
		svc := OnlineDeviceSvc{Data: &info}
		svc.OnlineDevice()
	case iotconst.NATS_SUBJECT_UPGRADE_REPORT:
		iotlogger.LogHelper.Info("设备OTA升级", string(msg.Data()))
		svc := OtaUpgradeDeviceSvc{Data: &info}
		svc.OtaUpgradeDevice()
	case iotconst.NATS_SUBJECT_CONTROL:
		svc := ReportDeviceSvc{Data: &info}
		//直接接收control上报
		iotlogger.LogHelper.Info("Control下发", string(msg.Data()))
		svc.ReportDevice()
	case iotconst.NATS_SUBJECT_REPORT:
		iotlogger.LogHelper.Info("控制上报", string(msg.Data()))
		svc := ReportDeviceSvc{Data: &info}
		//直接接收control上报
		if info.Name == iotprotocol.TP_C_CONTROL {
			//记录设备信息
			svc.ReportAckDevice()
			svc.ReportMsgDevice()
		} else {
			iotlogger.LogHelper.Info("设备定时上报", string(msg.Data()))
		}
		svc.ReportDeviceFault()
	}
}

func ErrorHandler(consumeCtx jetstream.ConsumeContext, err error) {
	if errors.Is(err, nats.ErrConnectionClosed) {
		iotlogger.LogHelper.Errorf("拉取消息失败,原因:%s", err.Error())
		time.Sleep(3 * time.Second)
	} else if !errors.Is(err, nats.ErrTimeout) {
		iotlogger.LogHelper.Errorf("拉取消息失败,原因:%s", err.Error())
	}
}
