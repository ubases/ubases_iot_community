package service

import (
	"cloud_platform/iot_basic_service/config"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotnatsjs"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_common/iotutil"
	proto "cloud_platform/iot_proto/protos/protosService"
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/nats-io/nats.go/jetstream"
	"github.com/nats-io/nats.go"
)

type BuildSubscriber struct {
	//suber      *jetstream.JSPullSubscriber
	suber      *iotnatsjs.JsClient
	concurrent int
	ctx        context.Context
	cancel     context.CancelFunc
}

func NewBuildSubscriber() (*BuildSubscriber, error) {
	appName := "iot_basic_service_sub"
	//suber, err := jetstream.NewJSPullSubscriber(appName, iotconst.NATS_LANGUAGE, iotconst.NATS_SUBJECT_LANGUAGE_UPDATE, connerrhandler, config.Global.Nats.Addrs...)
	//if err != nil {
	//	return nil, err
	//}
	//ctx, cancel := context.WithCancel(context.Background())
	suber, err := iotnatsjs.NewJsClient(config.Global.Nats.Addrs...)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	err = suber.CreateOrUpdateConsumer(ctx, iotconst.NATS_LANGUAGE, []string{iotconst.NATS_SUBJECT_LANGUAGE_UPDATE}, appName)
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
	//		info := iotstruct.TranslatePush{}
	//		err = json.Unmarshal(v.Data, &info)
	//		if err != nil {
	//			iotlogger.LogHelper.Errorf("翻译新增失败,内容[%s],错误:%s", string(v.Data), err.Error())
	//			continue
	//		}
	//		//保存翻译
	//		translateList := []*proto.BatchSaveTranslateItem{}
	//		for _, item := range info.TranslateList {
	//			translateItem := proto.BatchSaveTranslateItem{
	//				Lang:        item.Lang,
	//				FieldName:   item.FieldName,
	//				FieldType:   item.FieldType,
	//				FieldValue:  item.FieldValue,
	//				SourceRowId: item.SourceRowId,
	//			}
	//			if item.Id != "" {
	//				translateItem.Id = iotutil.ToInt64(item.Id)
	//			}
	//			translateList = append(translateList, &translateItem)
	//		}
	//
	//		svc := LangTranslateSvc{Ctx: context.Background()}
	//		err := svc.BatchCreateLangTranslate(&proto.BatchSaveTranslate{
	//			SourceTable:  info.SourceTable,
	//			SourceRowId:  info.SourceRowId,
	//			List:         translateList,
	//			PlatformType: info.PlatformType,
	//		})
	//		if err != nil {
	//			iotlogger.LogHelper.Error(err)
	//		} else {
	//			iotlogger.LogHelper.Infof("接收载荷信息:%s", string(v.Data))
	//		}
	//	}
	//}
			}

func (bs BuildSubscriber) Close() {
	bs.cancel()
	bs.suber.Close()
		}

func MessageHandler(msg jetstream.Msg) {
			info := iotstruct.TranslatePush{}
	err := json.Unmarshal(msg.Data(), &info)
			if err != nil {
		iotlogger.LogHelper.Errorf("翻译新增失败,内容[%s],错误:%s", string(msg.Data()), err.Error())
		return
			}
			//保存翻译
			translateList := []*proto.BatchSaveTranslateItem{}
			for _, item := range info.TranslateList {
				translateItem := proto.BatchSaveTranslateItem{
					Lang:        item.Lang,
					FieldName:   item.FieldName,
					FieldType:   item.FieldType,
					FieldValue:  item.FieldValue,
					SourceRowId: item.SourceRowId,
				}
				if item.Id != "" {
					translateItem.Id = iotutil.ToInt64(item.Id)
				}
				translateList = append(translateList, &translateItem)
			}

			svc := LangTranslateSvc{Ctx: context.Background()}
	err = svc.BatchCreateLangTranslate(&proto.BatchSaveTranslate{
				SourceTable:  info.SourceTable,
				SourceRowId:  info.SourceRowId,
				List:         translateList,
				PlatformType: info.PlatformType,
			})
			if err != nil {
				iotlogger.LogHelper.Error(err)
			} else {
		iotlogger.LogHelper.Infof("接收载荷信息:%s", string(msg.Data()))
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
