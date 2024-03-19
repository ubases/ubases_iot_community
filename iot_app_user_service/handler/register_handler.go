package handler

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_proto/protos/protosService"

	"go-micro.dev/v4"
)

// 新增加handler后，请在此注册
func RegisterHandler(s micro.Service) error {
	err := protosService.RegisterAppConfigServiceHandler(s.Server(), new(AppConfigHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterAppConfigServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterAppUpgradeServiceHandler(s.Server(), new(AppUpgradeHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterAppUpgradeServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterUcHomeServiceHandler(s.Server(), new(UcHomeHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterUcHomeServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterUcHomeRoomServiceHandler(s.Server(), new(UcHomeRoomHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterUcHomeRoomServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterUcHomeUserServiceHandler(s.Server(), new(UcHomeUserHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterUcHomeUserServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterUcUserFeedbackServiceHandler(s.Server(), new(UcUserFeedbackHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterUcUserFeedbackServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterUcUserFeedbackResultServiceHandler(s.Server(), new(UcUserFeedbackResultHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterUcUserFeedbackResultServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterUcUserServiceHandler(s.Server(), new(UcUserHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterUcUserServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterUcUserThirdServiceHandler(s.Server(), new(UcUserThirdHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterUcUserThirdServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterUcAppleidInfoServiceHandler(s.Server(), new(UcAppleidInfoHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterUcAppleidInfoServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterUcUserPrizeCollectServiceHandler(s.Server(), new(UcUserPrizeCollectHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterUcUserPrizeCollectServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterAppAuthHandler(s.Server(), new(AppAuthHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterAppAuthHandler:%s", err.Error())
		return err
	}
	return nil
}
