package handler

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_proto/protos/protosService"

	"go-micro.dev/v4"
)

// 新增加handler后，请在此注册
func RegisterHandler(s micro.Service) error {

	err := protosService.RegisterIotDeviceHomeServiceHandler(s.Server(), new(IotDeviceHomeHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterIotDeviceHomeServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterIotDeviceLogServiceHandler(s.Server(), new(IotDeviceLogHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterIotDeviceLogServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterIotDeviceInfoServiceHandler(s.Server(), new(IotDeviceInfoHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterIotDeviceInfoServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterIotDeviceTriadServiceHandler(s.Server(), new(IotDeviceTriadHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterIotDeviceTriadServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterIotOtaPkgServiceHandler(s.Server(), new(IotOtaPkgHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterIotOtaPkgServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterIotOtaVersionServiceHandler(s.Server(), new(IotOtaVersionHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterIotOtaVersionServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterIotOtaVersionPublishServiceHandler(s.Server(), new(IotOtaVersionPublishHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterIotOtaVersionPublishServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterIotOtaVersionVerifyServiceHandler(s.Server(), new(IotOtaVersionVerifyHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterIotOtaVersionVerifyServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterIotOtaWhiteGroupServiceHandler(s.Server(), new(IotOtaWhiteGroupHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterIotOtaWhiteGroupServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterIotOtaWhiteServiceHandler(s.Server(), new(IotOtaWhiteHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterIotOtaWhiteServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterIotDeviceFaultServiceHandler(s.Server(), new(IotDeviceFaultHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterIotDeviceFaultServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterIotDeviceTimerServiceHandler(s.Server(), new(IotDeviceTimerHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterIotDeviceTimerServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterIotDeviceCountdownServiceHandler(s.Server(), new(IotDeviceCountdownHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterIotDeviceCountdownServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterIotDeviceSharedServiceHandler(s.Server(), new(IotDeviceSharedHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterIotDeviceSharedServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterIotDeviceShareReceiveServiceHandler(s.Server(), new(IotDeviceShareReceiveHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterIotDeviceShareReceiveServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterIotDeviceGroupServiceHandler(s.Server(), new(IotDeviceGroupHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterIotDeviceShareReceiveServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterIotDeviceGroupListServiceHandler(s.Server(), new(IotDeviceGroupListHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterIotDeviceShareReceiveServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterIotOtaUpgradeRecordServiceHandler(s.Server(), new(IotOtaUpgradeRecordHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterIotOtaUpgradeRecordServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterIotJobServiceHandler(s.Server(), new(IotJobHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterIotJobServiceHandler 错误:%s", err.Error())
		return err
	}

	return nil
}
