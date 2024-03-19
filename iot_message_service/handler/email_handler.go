package handler

import (
	"cloud_platform/iot_proto/protos/protosService"
	"context"

	"cloud_platform/iot_message_service/service"
)

type EmailServiceHandler struct{}

func (EmailServiceHandler) SendEmailUserCode(ctx context.Context, request *protosService.SendEmailUserCodeRequest, response *protosService.SendEmailResponse) error {
	s := service.EmailSvc{Ctx: ctx}
	resp, err := s.SendEmailUserCode(request)
	if err == nil {
		*response = *resp
	}
	return nil
}

func (EmailServiceHandler) SendEmailUserLoggedIn(ctx context.Context, request *protosService.SendEmailUserLoggedInRequest, response *protosService.SendEmailResponse) error {
	s := service.EmailSvc{Ctx: ctx}
	resp, err := s.SendEmailUserLoggedIn(request)
	if err == nil {
		*response = *resp
	}
	return nil
}

func (EmailServiceHandler) SendEmailUserRegister(ctx context.Context, request *protosService.SendEmailUserRegisterRequest, response *protosService.SendEmailResponse) error {
	s := service.EmailSvc{Ctx: ctx}
	resp, err := s.SendEmailUserRegister(request)
	if err == nil {
		*response = *resp
	}
	return nil
}
