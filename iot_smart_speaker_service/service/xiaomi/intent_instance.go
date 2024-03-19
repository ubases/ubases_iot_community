package xiaomi

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_smart_speaker_service/cached"
	"cloud_platform/iot_smart_speaker_service/entitys"
	"context"
	"fmt"
	"time"
)

const (
	XiaomiDiscoverDevicesResponse = "get-devices"
	XiaomiSubscribeResponse       = "subscribe"
	XiaomiInvokeActionResponse    = "invoke-action"
	XiaomiUnsubscribeResponse     = "unsubscribe"
	XiaomiGetDeviceStatusResponse = "get-device-status"
	XiaomiGetProperties           = "get-properties"
	XiaomiSetProperties           = "set-properties"
)

func RunIntent(res *entitys.XiaomiRequest, data []byte, userId, token string) error {
	err := cached.RedisStore.GetClient().Set(context.Background(), fmt.Sprintf(iotconst.XiaomiVoiceUserTokenKey, userId), token, time.Hour*24*2).Err()
	if err != nil {
		iotlogger.LogHelper.Helper.Error("set voice user token error: ", err)
	}
	switch res.Intent {
	case XiaomiDiscoverDevicesResponse:
		return DiscoverDevices(res, userId, token)
	case XiaomiSubscribeResponse:
		return SubscribeDevices(res, userId, token)
	case XiaomiInvokeActionResponse:
	case XiaomiUnsubscribeResponse:
		return UnSubscribeDevices(res, userId, token)
	case XiaomiGetDeviceStatusResponse:
		return StatusDevices(res, userId, token)
	case XiaomiGetProperties:
		return GetProperties(res, data, userId, token)
	case XiaomiSetProperties:
		return SetProperties(res, data, userId, token)
	}
	return nil
}
