package handler

import (
	"cloud_platform/iot_open_system_service/service"
	proto "cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

var ErrUserPassword = errors.New("账号或密码错误")

type OpenAuthHandler struct{}

func (h *OpenAuthHandler) MiniProgramLogin(ctx context.Context, request *proto.MiniProgramLoginRequest, response *proto.CloudLoginResponse) error {
	s := service.OpenAuthSvc{Ctx: ctx}
	resp, err := s.MiniProgramLogin(request)
	if err == nil {
		*response = *resp
	}
	return nil
}

func (h *OpenAuthHandler) PhoneCodeLogin(ctx context.Context, request *proto.PhoneCodeLoginRequest, response *proto.CloudLoginResponse) error {
	s := service.OpenAuthSvc{Ctx: ctx}
	resp, err := s.PhoneCodeLogin(request)
	if err == nil {
		*response = *resp
	}
	return err
}

func (h *OpenAuthHandler) EmailCodeLogin(ctx context.Context, request *proto.EmailCodeLoginRequest, response *proto.CloudLoginResponse) error {
	s := service.OpenAuthSvc{Ctx: ctx}
	resp, err := s.EmailCodeLogin(request)
	if err == nil {
		*response = *resp
	}
	return err
}

func (h *OpenAuthHandler) PasswordLogin(ctx context.Context, request *proto.PasswordLoginRequest, response *proto.CloudLoginResponse) error {
	s := service.OpenAuthSvc{Ctx: ctx}
	resp, err := s.PasswordLogin(request)
	if err == nil {
		*response = *resp
	} else {
		err = ErrUserPassword
	}
	return err
}

func (h *OpenAuthHandler) Logout(ctx context.Context, request *proto.LogoutRequest, response *proto.LogoutResponse) error {
	s := service.OpenAuthSvc{Ctx: ctx}
	resp, err := s.Logout(request)
	if err == nil {
		response.Result = resp.Result
	}
	return err
}

func (h *OpenAuthHandler) VerifyToken(ctx context.Context, request *proto.VerifyTokenRequest, response *proto.CloudVerifyTokenResponse) error {
	s := service.OpenAuthSvc{Ctx: ctx}
	resp, err := s.VerifyToken(request)
	if err == nil {
		*response = *resp
	}
	return err
}

func (h *OpenAuthHandler) RefreshToken(ctx context.Context, request *proto.RefreshTokenRequest, response *proto.CloudRefreshTokenResponse) error {
	s := service.OpenAuthSvc{Ctx: ctx}
	resp, err := s.RefreshToken(request)
	if err == nil {
		*response = *resp
	}
	return err
}

func (h *OpenAuthHandler) ChangeTenant(ctx context.Context, request *proto.RefreshTokenRequest, response *proto.CloudRefreshTokenResponse) error {
	s := service.OpenAuthSvc{Ctx: ctx}
	resp, err := s.ChangeTenant(request)
	if err == nil {
		*response = *resp
	}
	return err
}
