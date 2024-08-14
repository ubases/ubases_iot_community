package google

import "cloud_platform/iot_voice_service/service/google/proto"

type State struct {
	Name  string
	Value interface{}
	Error proto.ErrorCode
}
