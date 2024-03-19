package handler

import (
	"cloud_platform/iot_message_service/service"
	proto "cloud_platform/iot_proto/protos/protosService"
	"context"
)

type SmsServiceHandler struct{}

func (SmsServiceHandler) SendCode(ctx context.Context, request *proto.SendSMSCodeRequest, response *proto.SendSMSResponse) error {
	s := service.SmsSvc{Ctx: ctx}
	resp, err := s.SendCode(request)
	if err == nil {
		*response = *resp
	}
	return nil
}

func (SmsServiceHandler) SendSMSVerifyCode(ctx context.Context, request *proto.SendSMSVerifyCodeRequest, response *proto.SendSMSResponse) error {
	s := service.SmsSvc{Ctx: ctx}
	resp, err := s.SendSMSVerifyCode(request)
	if err == nil {
		*response = *resp
	}
	return nil
}

func (SmsServiceHandler) SendLoggedIn(ctx context.Context, request *proto.SendSMSLoggedInRequest, response *proto.SendSMSResponse) error {
	s := service.SmsSvc{Ctx: ctx}
	resp, err := s.SendLoggedIn(request)
	if err == nil {
		*response = *resp
	}
	return nil
}

func (SmsServiceHandler) SendRegister(ctx context.Context, request *proto.SendSMSRegisterRequest, response *proto.SendSMSResponse) error {
	s := service.SmsSvc{Ctx: ctx}
	resp, err := s.SendRegister(request)
	if err == nil {
		*response = *resp
	}
	return nil
}
