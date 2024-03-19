package push

import (
	"cloud_platform/iot_message_service/service/push/pushModel"
)

const (
	Gorush = "gorush" //Gorush
	JPush  = "jpush"
	All    = "all"
)

type PushClient interface {
	PushMessage(inputTarget pushModel.MessageTarget, message pushModel.MessageRequestModel) error
	ClearAlias(userId, appKey string) error
}

func NewPushClient(provider string) (PushClient, error) {
	switch provider {
	case Gorush:
		return &JPushClient{}, nil
	case JPush:
		return &GorushClient{}, nil
	case All:
		return &AllClient{}, nil
	}
	return nil, nil
}
