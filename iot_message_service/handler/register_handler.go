package handler

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_proto/protos/protosService"

	"go-micro.dev/v4"
)

// 新增加handler后，请在此注册
func RegisterHandler(s micro.Service) error {
	err := protosService.RegisterMpMessageUserInServiceHandler(s.Server(), new(MpMessageUserInHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterCmsDocumentFilesServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterMpMessageUserOutServiceHandler(s.Server(), new(MpMessageUserOutHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterMpMessageUserOutServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterMpMessageRedDotServiceHandler(s.Server(), new(MpMessageRedDotHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterCmsDocumentFilesServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterMpMessageServiceHandler(s.Server(), new(MpMessageHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterMpMessageServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterMpMessageTemplateServiceHandler(s.Server(), new(MpMessageTemplateHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterMpMessageTemplateServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterAppPushTokenServiceHandler(s.Server(), new(AppPushTokenHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterAppPushTokenServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterAppPushTokenUserServiceHandler(s.Server(), new(AppPushTokenUserHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterAppPushTokenUserServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterEmailServiceHandler(s.Server(), new(EmailServiceHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterEmailServiceHandler 错误:%s", err.Error())
		return err
	}
	err = protosService.RegisterMsNoticeTemplateServiceHandler(s.Server(), new(MsNoticeTemplateHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterMsNoticeTemplateServiceHandler 错误:%s", err.Error())
		return err
	}

	err = protosService.RegisterSmsServiceHandler(s.Server(), new(SmsServiceHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterSmsServiceHandler 错误:%s", err.Error())
		return err
	}

	return nil
}
