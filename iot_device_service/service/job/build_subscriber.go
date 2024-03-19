package job

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotnats/jetstream"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_device_service/config"
	"context"
	"errors"
	json "github.com/json-iterator/go"
	"time"

	"github.com/nats-io/nats.go"
)

type BuildSubscriber struct {
	suber      *jetstream.JSPullSubscriber
	concurrent int
	ctx        context.Context
	cancel     context.CancelFunc
}

func NewBuildSubscriber() (*BuildSubscriber, error) {
	appName := "iot_device_job_service"
	suber, err := jetstream.NewJSPullSubscriber(appName, iotconst.NATS_STREAM_DEVICE, iotconst.NATS_SUBJECT_DEVICE_JOB, connerrhandler, config.Global.Nats.Addrs...)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	Concurrent := 1
	return &BuildSubscriber{suber, Concurrent, ctx, cancel}, nil
}

func connerrhandler(conn *nats.Conn, err error) {
	if err != nil {
		iotlogger.LogHelper.Errorf("nats连接错误:%s", err.Error())
	}
}

func (bs BuildSubscriber) Run() {
	// 从nats消息队列拉取数据，通过rpc接口，发布数据到iot_mqtt_service服务
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
			info := iotstruct.MqttToNatsData{}
			err = json.Unmarshal(v.Data, &info)
			if err != nil {
				iotlogger.LogHelper.Errorf("解析构建信息失败,内容[%s],错误:%s", string(v.Data), err.Error())
				continue
			}
			data := map[string]interface{}{}
			payload, ok := info.Payload.(string)
			if !ok {
				iotlogger.LogHelper.Errorf("载荷信息断言失败,内容[%v],错误:%s", info.Payload, err.Error())
				continue
			}
			if err := json.Unmarshal([]byte(payload), &data); err != nil {
				iotlogger.LogHelper.Errorf("解析载荷信息失败,内容[%s],错误:%s", payload, err.Error())
				continue
			}
			//convData := iotutil.MapStringToInterface(data)
			if _, err := PubControl(info.ProductKey, info.DeviceId, data); err != nil {
				iotlogger.LogHelper.Errorf("发布控制指令失败,ProductKey[%s],DeviceId[%s],数据[%s],错误:%s", info.ProductKey, info.DeviceId, payload, err.Error())
				continue
			}
			iotlogger.LogHelper.Infof("接收载荷信息:%s", info.Payload)
		}
	}
}

func (bs BuildSubscriber) Close() {
	bs.cancel()
	bs.suber.Close()
}
