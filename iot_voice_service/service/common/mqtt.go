package common

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotprotocol"
	"cloud_platform/iot_proto/protos/protosService"
	"cloud_platform/iot_voice_service/rpc/rpcclient"
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
	_, err = rpcclient.ClientMqttMessage.Publish(context.Background(), req)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return 0, err
	}
	return 0, nil
}

func PubQuery(productKey, devId string, data []string) (int32, error) {
	var err error
	topic := iotprotocol.GetTopic(iotprotocol.TP_C_QUERY, productKey, devId)
	packControl := iotprotocol.PackQuery{}
	pushData, err := packControl.Encode("", "query", data)
	if err != nil {
		return 0, err
	}
	req := &protosService.PublishMessage{
		TopicFullName:  topic,
		MessageContent: string(pushData),
		Qos:            1,
		Retained:       false,
	}
	_, err = rpcclient.ClientMqttMessage.Publish(context.Background(), req)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return 0, err
	}
	return 0, nil
}
