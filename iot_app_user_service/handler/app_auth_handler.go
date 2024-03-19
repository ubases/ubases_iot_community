package handler

import (
	"cloud_platform/iot_app_user_service/service"
	proto "cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

var ErrUserPassword = errors.New("账号或密码错误")

type AppAuthHandler struct {
}

func (AppAuthHandler) MiniProgramLogin(ctx context.Context, request *proto.MiniProgramLoginRequest, response *proto.AppLoginResponse) error {
	s := service.AppAuthSvc{Ctx: ctx}
	resp, err := s.MiniProgramLogin(request)
	if err == nil {
		*response = *resp
	}
	return err
}

func (AppAuthHandler) PhoneCodeLogin(ctx context.Context, request *proto.PhoneCodeLoginRequest, response *proto.AppLoginResponse) error {
	s := service.AppAuthSvc{Ctx: ctx}
	resp, err := s.PhoneCodeLogin(request)
	if err == nil {
		*response = *resp
	}
	return err
}

func (AppAuthHandler) EmailCodeLogin(ctx context.Context, request *proto.EmailCodeLoginRequest, response *proto.AppLoginResponse) error {
	s := service.AppAuthSvc{Ctx: ctx}
	resp, err := s.EmailCodeLogin(request)
	if err == nil {
		*response = *resp
	}
	return err
}

func (AppAuthHandler) PasswordLogin(ctx context.Context, request *proto.PasswordLoginRequest, response *proto.AppLoginResponse) error {
	s := service.AppAuthSvc{Ctx: ctx}
	resp, err := s.PasswordLogin(request)
	if err == nil {
		*response = *resp
	} else {
		err = ErrUserPassword
	}
	return err
}

func (AppAuthHandler) Logout(ctx context.Context, request *proto.LogoutRequest, response *proto.LogoutResponse) error {
	s := service.AppAuthSvc{Ctx: ctx}
	resp, err := s.Logout(request)
	if err == nil {
		response.Result = resp.Result
	}
	return err
}

func (AppAuthHandler) VerifyToken(ctx context.Context, request *proto.VerifyTokenRequest, response *proto.APPVerifyTokenResponse) error {
	s := service.AppAuthSvc{Ctx: ctx}
	resp, err := s.VerifyToken(request)
	if err == nil {
		response.Valid = resp.Valid
	}
	return err
}

func (AppAuthHandler) RefreshToken(ctx context.Context, request *proto.RefreshTokenRequest, response *proto.AppRefreshTokenResponse) error {
	s := service.AppAuthSvc{Ctx: ctx}
	resp, err := s.RefreshToken(request)
	if err == nil {
		*response = *resp
	}
	return err
}

func (AppAuthHandler) WechatLogin(ctx context.Context, request *proto.AppThirdRequest, response *proto.AppLoginResponse) error {
	s := service.AppAuthSvc{Ctx: ctx}
	resp, err := s.WechatLogin(request)
	if err == nil {
		*response = *resp
	}
	return err
}

func (AppAuthHandler) AppleidLogin(ctx context.Context, request *proto.AppThirdRequest, response *proto.AppLoginResponse) error {
	s := service.AppAuthSvc{Ctx: ctx}
	resp, err := s.AppleidLogin(request)
	if err == nil {
		*response = *resp
	}
	return err
}

// userId int64,nickname,avatar,account string
func (AppAuthHandler) GetToken(ctx context.Context, request *proto.GetTokenRequest, response *proto.AppLoginResponse) error {
	s := service.AppAuthSvc{Ctx: ctx}
	resp, err := s.GetUserAndToken(request.UserId, request.Nickname, request.Avatar, request.Account)
	if err == nil {
		*response = *resp
	}
	return err
}

func (AppAuthHandler) GetTokenByAccount(ctx context.Context, request *proto.GetTokenByAccountRequest, response *proto.AppLoginResponse) error {
	s := service.AppAuthSvc{Ctx: ctx}
	resp, err := s.GetTokenByAccount(request)
	if err == nil {
		*response = *resp
	}
	return err
}
