package service

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotnatsjs"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_message_service/config"
	proto "cloud_platform/iot_proto/protos/protosService"
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type BuildSubscriber struct {
	//suber      *jetstream.JSPullSubscriber
	suber      *iotnatsjs.JsClient
	concurrent int
	ctx        context.Context
	cancel     context.CancelFunc
}

func NewBuildSubscriber() (*BuildSubscriber, error) {
	appName := "iot_message_service_sub"
	//suber, err := jetstream.NewJSPullSubscriber(appName, iotconst.NATS_MESSAGE, iotconst.NATS_SUBJECT_MESSAGE_UPDATE, connerrhandler, config.Global.Nats.Addrs...)
	//if err != nil {
	//	return nil, err
	//}
	//ctx, cancel := context.WithCancel(context.Background())
	suber, err := iotnatsjs.NewJsClient(config.Global.Nats.Addrs...)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	err = suber.CreateOrUpdateConsumer(ctx, iotconst.NATS_MESSAGE, []string{iotconst.NATS_SUBJECT_MESSAGE_UPDATE}, appName)
	if err != nil {
		cancel()
		return nil, err
	}
	Concurrent := 1
	return &BuildSubscriber{suber, Concurrent, ctx, cancel}, nil
}

//func connerrhandler(conn *nats.Conn, err error) {
//	if err != nil {
//		iotlogger.LogHelper.Errorf("nats连接错误:%s", err.Error())
//	}
//}

func (bs BuildSubscriber) Run() {
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
	//			iotlogger.LogHelper.Errorf("翻译新增失败,内容[%s],错误:%s", string(v.Data), err.Error())
	//			continue
	//		}
	//		data := &proto.SendMessageRequest{}
	//		payload, ok := info.Payload.(string)
	//		if !ok {
	//			iotlogger.LogHelper.Errorf("载荷信息断言失败,内容[%v],错误:%s", info.Payload, err.Error())
	//			continue
	//		}
	//		if err := json.Unmarshal([]byte(payload), &data); err != nil {
	//			iotlogger.LogHelper.Errorf("解析载荷信息失败,内容[%s],错误:%s", payload, err.Error())
	//			continue
	//		}
	//		//TODO 消息增加类型，推送消息
	//		svc := MpMessageSvc{Ctx: context.Background()}
	//		res, err := svc.PushMessage(data)
	//		if err != nil {
	//			return
	//		}
	//		smSvc := SendMessageSvc{}
	//		smSvc.SendMessage(res.Id, data)
	//
	//		iotlogger.LogHelper.Infof("接收载荷信息:%s", info.Payload)
	//	}
	//}
}

func (bs BuildSubscriber) Close() {
	bs.cancel()
	bs.suber.Close()
}

func MessageHandler(msg jetstream.Msg) {
	info := iotstruct.MqttToNatsData{}
	err := json.Unmarshal(msg.Data(), &info)
	if err != nil {
		iotlogger.LogHelper.Errorf("翻译新增失败,内容[%s],错误:%s", string(msg.Data()), err.Error())
		return
	}
	data := &proto.SendMessageRequest{}
	payload, ok := info.Payload.(string)
	if !ok {
		iotlogger.LogHelper.Errorf("载荷信息断言失败,内容[%v],错误:%s", info.Payload, err.Error())
		return
	}
	if err := json.Unmarshal([]byte(payload), &data); err != nil {
		iotlogger.LogHelper.Errorf("解析载荷信息失败,内容[%s],错误:%s", payload, err.Error())
		return
	}
	//TODO 消息增加类型，推送消息
	svc := MpMessageSvc{Ctx: context.Background()}
	res, err := svc.PushMessage(data)
	if err != nil {
		return
	}
	smSvc := SendMessageSvc{}
	smSvc.SendMessage(res.Id, data)

	iotlogger.LogHelper.Infof("接收载荷信息:%s", info.Payload)
}

func ErrorHandler(consumeCtx jetstream.ConsumeContext, err error) {
	if errors.Is(err, nats.ErrConnectionClosed) {
		iotlogger.LogHelper.Errorf("拉取消息失败,原因:%s", err.Error())
		time.Sleep(3 * time.Second)
	} else if !errors.Is(err, nats.ErrTimeout) {
		iotlogger.LogHelper.Errorf("拉取消息失败,原因:%s", err.Error())
	}
}
