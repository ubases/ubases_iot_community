package job

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotprotocol"
	"cloud_platform/iot_device_service/rpc/rpcClient"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
)

func PubControl(productKey, devId string, data map[string]interface{}) (int32, error) {
	var err error
	topic := iotprotocol.GetTopic(iotprotocol.TP_C_CONTROL, productKey, devId)
	packControl := iotprotocol.PackControl{}
	pushData, err := packControl.Encode("", "control", data)
	if err != nil {
		return 0, err
	}
	req := &protosService.PublishMessage{
		TopicFullName:  topic,
		MessageContent: string(pushData),
		Qos:            1,
		Retained:       false,
	}
	_, err = rpcClient.ClientMqttService.Publish(context.Background(), req)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return 0, err
	}
	return 0, nil
}

func RunToMQTT() {
	bs, err := NewBuildSubscriber()
	if err != nil {
		iotlogger.LogHelper.Error("创建设备定时任务执行服务错误:", err.Error())
		return
	}
	defer bs.Close()
	bs.Run()
}
