package sync_update

import (
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_voice_service/service/google"
	"context"
)

type UpdateDeviceSvc struct {
	Data  *iotstruct.DeviceRedisUpdate
	DevId string
}

// UpdateDevice 更新设备列表
func (s *UpdateDeviceSvc) UpdateDevice() error {
	////同步在线状态
	//go google.GetSmartHome().RequestSync(context.Background(), s.Data.UserId)

	//同步设备列表（删除设备、切换家庭）
	go RunDeviceListSync(*s)

	//同步在线状态
	go google.GetSmartHome().RequestSync(context.Background(), s.Data.UserId)
	return nil
}
