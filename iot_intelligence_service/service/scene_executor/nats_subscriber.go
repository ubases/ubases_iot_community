package scene_executor

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotnatsjs"
	"cloud_platform/iot_common/iotprotocol"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_intelligence_service/config"
	"context"

	"github.com/nats-io/nats.go/jetstream"

	"encoding/json"
)

var WeatherChan chan map[string]interface{} //protosService.WeatherData
var DeviceChan chan iotstruct.DeviceRedisData

func init() {
	WeatherChan = make(chan map[string]interface{}, 100)
	DeviceChan = make(chan iotstruct.DeviceRedisData, 100000)
}

func InitReportSub() {
	ackCh := iotconst.HKEY_REPORT_DATA_PUB_PREFIX + ".>"
	client, err := iotnatsjs.NewJsClient(config.Global.Nats.Addrs...)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer client.Close()
	err = client.CreateOrUpdateConsumer(ctx, iotconst.NATS_STREAM_ORIGINAL_REDIS, []string{ackCh}, "iot_intelligence_service_1")
	if err != nil {
		panic(err)
	}
	//client.PSubscribe(handlerRedisMessage, ackCh)
	jsctx, err := client.Consume(func(msg jetstream.Msg) {
		data := iotstruct.DeviceRedisData{}
		if err := json.Unmarshal([]byte(msg.Data()), &data); err != nil {
			iotlogger.LogHelper.Helper.Error("json unmarshal online error: ", err)
			return
		}
		//非控制类回传通知消息，不处理 update by hogan 20240125
		if data.Name != iotprotocol.REPORT_HEAD_NAME && data.Name != iotprotocol.CONTROL_HEAD_NAME {
			return
		}
		DeviceChan <- data
	}, func(consumeCtx jetstream.ConsumeContext, err error) {

	})
	if err != nil {
		return
	}

	defer jsctx.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}
