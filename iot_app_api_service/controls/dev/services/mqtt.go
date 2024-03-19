package services

import (
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotprotocol"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
)

func PubControl(productKey, devId string, data map[string]interface{}) (string, int32, error) {
	var err error
	topic := iotprotocol.GetTopic(iotprotocol.TP_C_CONTROL, productKey, devId)
	packControl := iotprotocol.PackControl{}
	pushData, err := packControl.Encode("", iotprotocol.CONTROL_HEAD_NAME, data)
	if err != nil {
		return "", 0, err
	}
	req := &protosService.PublishMessage{
		TopicFullName:  topic,
		MessageContent: string(pushData),
		Qos:            0,
		Retained:       false,
	}
	_, err = rpc.ClientMqttService.Publish(context.Background(), req)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return "", 0, err
	}
	return packControl.Header.Mid, 0, nil
}
