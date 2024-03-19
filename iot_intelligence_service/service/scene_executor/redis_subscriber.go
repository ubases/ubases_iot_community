package scene_executor

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotprotocol"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_intelligence_service/cached"
	"strings"

	"context"
	"encoding/json"
)

var WeatherChan chan map[string]interface{} //protosService.WeatherData
var DeviceChan chan iotstruct.DeviceRedisData

func init() {
	WeatherChan = make(chan map[string]interface{})
	DeviceChan = make(chan iotstruct.DeviceRedisData)
}

func InitRedisSub() {
	ctx := context.Background()
	//TestDeviceChan()
	ackCh := strings.Join([]string{iotconst.HKEY_REPORT_DATA_PUB_PREFIX, "*"}, ".")
	ackSub := cached.RedisStore.GetClient().PSubscribe(ctx, ackCh)
	defer ackSub.Close()

	ackChannel := ackSub.Channel()
	for {
		select {
		case msg := <-ackChannel:
			data := iotstruct.DeviceRedisData{}
			if err := json.Unmarshal([]byte(msg.Payload), &data); err != nil {
				iotlogger.LogHelper.Helper.Error("json unmarshal online error: ", err)
				continue
			}
			if data.Name != iotprotocol.REPORT_HEAD_NAME && data.Name != iotprotocol.CONTROL_HEAD_NAME {
				continue
			}
			DeviceChan <- data
		}
	}
}
