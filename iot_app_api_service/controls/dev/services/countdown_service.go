package services

import (
	"cloud_platform/iot_app_api_service/controls/dev/entitys"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"strconv"
	"time"

	goerrors "go-micro.dev/v4/errors"
)

type AppCountdownService struct {
}

// AddCountdown 添加倒计时，每个设备一个倒计时
func (s AppCountdownService) AddCountdown(req entitys.IotDeviceCountdownEntitys) error {
	_hour := iotutil.ToInt(req.Hour)
	_minute := iotutil.ToInt(req.Minute)
	totalSeconds := (_hour*60 + _minute) * 60
	totalMinutes := _hour*60 + _minute
	strMinutes := strconv.Itoa(totalMinutes)
	addMinutes, err := time.ParseDuration(strMinutes + "m")
	if err != nil {
		return goerrors.New("", err.Error(), ioterrs.ErrAppCountDownAddParam)
	}
	//服务器本地时间增加倒计时分钟数，获取执行时间；
	runTime := time.Now().Add(addMinutes)
	req.Cron = ConvertSpecOnlyOne(runTime)
	req.TotalSecond = int32(totalSeconds)
	req.ExecutionTime = runTime
	req.Enabled = 1

	saveObj := entitys.IotDeviceCountdown_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()
	_, err = rpc.IotDeviceCountdownService.Create(context.Background(), saveObj)
	if err != nil {
		return err
	}
	return err
}

// DeleteCountdown 删除倒计时
func (s AppCountdownService) DeleteCountdown(id string) error {
	_, err := rpc.IotDeviceCountdownService.Delete(context.Background(), &protosService.IotDeviceCountdown{
		DeviceId: id,
	})
	if err != nil {
		return err
	}
	return nil
}

// DisabledCountdown 停用倒计时
func (s AppCountdownService) DisabledCountdown(id int64) error {
	_, err := rpc.IotDeviceCountdownService.StopIotDeviceCountdownJob(context.Background(), &protosService.IotDeviceCountdownJobReq{
		Id: id,
	})
	if err != nil {
		return err
	}
	return nil
}

// EnabledCountdown 启用倒计时
func (s AppCountdownService) EnabledCountdown(id int64) error {
	_, err := rpc.IotDeviceCountdownService.StartIotDeviceCountdownJob(context.Background(), &protosService.IotDeviceCountdownJobReq{
		Id: id,
	})
	if err != nil {
		return err
	}
	return nil
}

// CountdownInfo 获取倒计时信息
func (s AppCountdownService) CountdownInfo(id string) (*entitys.IotDeviceCountdownVo, error) {
	res, err := rpc.IotDeviceCountdownService.Find(context.Background(), &protosService.IotDeviceCountdownFilter{
		DeviceId: id,
	})
	if err != nil {
		return nil, err
	}
	if len(res.Data) == 0 {
		return &entitys.IotDeviceCountdownVo{DeviceId: id, Enabled: 2}, nil
	}
	return entitys.IotDeviceCountdown_vo(res.Data[0]), nil
}
