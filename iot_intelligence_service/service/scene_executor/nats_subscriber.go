package scene_executor

import (
	"cloud_platform/iot_intelligence_service/config"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotnats/jetstream"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
)

// var WeatherChan chan protosService.WeatherData
var WeatherChan2 chan protosService.WeatherData
var DeviceChan2 chan iotstruct.DeviceRedisData

//var DeviceChan chan iotstruct.MqttToNatsData
//var DeviceChan chan protosService.DeviceData

type Subscriber struct {
	Subscribers []*DeviceSubscriber
}

func InitSubscriber() (subOp *Subscriber, err error) {
	subOp = new(Subscriber)
	subOp.SetSubscriber(strings.Join([]string{"iot_intelligence_service", "data", "weather"}, "_"),
		iotconst.NATS_WEATHER,
		iotconst.NATS_SUBJECT_WEATHER_DATA)
	subOp.SetSubscriber(strings.Join([]string{"iot_device_service", "device", "report"}, "_"),
		iotconst.NATS_STREAM_DEVICE,
		iotconst.NATS_SUBJECT_REPORT)
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
	suber      *jetstream.JSPullSubscriber
	concurrent int
	ctx        context.Context
	cancel     context.CancelFunc
}

func NewDeviceSubscriber(appname string, stream string, subject string) (*DeviceSubscriber, error) {
	suber, err := jetstream.NewJSPullSubscriber(appname, stream, subject, connerrhandler, config.Global.Nats.Addrs...)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	Concurrent := 1
	return &DeviceSubscriber{suber, Concurrent, ctx, cancel}, nil
}

func connerrhandler(conn *nats.Conn, err error) {
	if err != nil {
		iotlogger.LogHelper.Errorf("nats连接错误:%s", err.Error())
	}
}

func (bs DeviceSubscriber) Run() {
	for {
		if bs.ctx.Err() != nil {
			break
		}
		msgList, err := bs.suber.FetchMessageEx(1)
		if err != nil {
			if errors.Is(err, nats.ErrConnectionClosed) {
				iotlogger.LogHelper.Errorf("拉取消息失败,原因:%s", err.Error())
				time.Sleep(3 * time.Second)
			} else if !errors.Is(err, nats.ErrTimeout) {
				iotlogger.LogHelper.Errorf("拉取消息失败,原因:%s", err.Error())
			}
			continue
		}
		for _, v := range msgList {
			switch v.Subject {
			case iotconst.NATS_SUBJECT_WEATHER_DATA:
				info := protosService.WeatherData{}
				err = json.Unmarshal(v.Data, &info)
				if err != nil {
					iotlogger.LogHelper.Errorf("解析天气失败,内容[%s],错误:%s", string(v.Data), err.Error())
					continue
				}
				//天气变化
				//WeatherChan <- info
				iotlogger.LogHelper.Info("收到天气消息", string(v.Data))
			case iotconst.NATS_SUBJECT_REPORT:
				info := iotstruct.MqttToNatsData{}
				err = json.Unmarshal(v.Data, &info)
				if err != nil {
					iotlogger.LogHelper.Errorf("解析设备上报消息失败,内容[%s],错误:%s", string(v.Data), err.Error())
					continue
				}
				//设备状态变化
				//DeviceChan <- info
				iotlogger.LogHelper.Info("收到设备状态变化消息", string(v.Data))
			}
		}
	}
}

func (bs DeviceSubscriber) Close() {
	bs.cancel()
	bs.suber.Close()
}
