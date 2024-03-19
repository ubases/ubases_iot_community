package google

import "cloud_platform/iot_smart_speaker_service/service/google/proto"

type State struct {
	Name  string
	Value interface{}
	Error proto.ErrorCode
}
