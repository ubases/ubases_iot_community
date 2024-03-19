package handler

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_proto/protos/protosService"

	"go-micro.dev/v4"
)

// 新增加handler后，请在此注册
func RegisterHandler(s micro.Service) error {
	err := protosService.RegisterDataOverviewMonthServiceHandler(s.Server(), new(DataOverviewMonthHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterDataOverviewMonthServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterDataOverviewHourServiceHandler(s.Server(), new(DataOverviewHourHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterDataOverviewMonthServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterPmAppDataServiceHandler(s.Server(), new(PmAppDataHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterPmAppDataServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterPmDevelopDataServiceHandler(s.Server(), new(PmDevelopDataHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterPmDevelopDataServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterProductFaultMonthServiceHandler(s.Server(), new(ProductFaultMonthHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterProductFaultMonthServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterProductFaultTypeServiceHandler(s.Server(), new(ProductFaultTypeHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterProductFaultTypeServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterDeviceActiveDayServiceHandler(s.Server(), new(DeviceActiveDayHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterDeviceActiveDayServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterDeviceActiveHourServiceHandler(s.Server(), new(DeviceActiveHourHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterDeviceActiveHourServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterDeviceActiveMonthServiceHandler(s.Server(), new(DeviceActiveMonthHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterDeviceActiveMonthServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterDeviceActiveServiceHandler(s.Server(), new(DeviceActiveHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterDeviceActiveServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterStatisticsServiceHandler(s.Server(), new(StatisticsHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterStatisticsServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterAppUserActiveDayServiceHandler(s.Server(), new(AppUserActiveDayHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterAppUserActiveDayServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterAppUserMonthServiceHandler(s.Server(), new(AppUserMonthHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterAppUserMonthServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterAppUserStatisticsServiceHandler(s.Server(), new(AppUserStatisticsHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterAppUserStatisticsServiceHandler 错误:%s", err.Error())
		return err
	}
	return nil
}
