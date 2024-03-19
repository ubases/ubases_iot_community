package service

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	proto "cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"strings"

	"github.com/golang-jwt/jwt"
)

type CloudAuthSvc struct {
	Ctx context.Context
}

func (s *CloudAuthSvc) MiniProgramLogin(request *proto.MiniProgramLoginRequest) (*proto.CloudLoginResponse, error) {
	return nil, errors.New("not implement")
}

func (s *CloudAuthSvc) PhoneCodeLogin(request *proto.PhoneCodeLoginRequest) (*proto.CloudLoginResponse, error) {
	return nil, errors.New("not implement")
}

func (s *CloudAuthSvc) EmailCodeLogin(request *proto.EmailCodeLoginRequest) (*proto.CloudLoginResponse, error) {
	return nil, errors.New("not implement")
}

func (s *CloudAuthSvc) PasswordLogin(request *proto.PasswordLoginRequest) (*proto.CloudLoginResponse, error) {
	if request.LoginName == "" || request.Password == "" {
		err := errors.New("用户名或密码为空")
		iotlogger.LogHelper.Errorf("PasswordLogin: 用户%s登录失败，原因:%s", request.LoginName, err.Error())
		return nil, err
	}
	su := SysUserSvc{s.Ctx}
	req := proto.SysUserFilter{UserName: request.LoginName}
	resp, err := su.FindSysUser(&req)
	if err != nil {
		iotlogger.LogHelper.Errorf("PasswordLogin: 用户%s登录失败，原因:%s", request.LoginName, err.Error())
		return nil, err
	}
	if resp == nil {
		err = errors.New("用户不存在")
		iotlogger.LogHelper.Errorf("PasswordLogin: 用户%s登录失败，原因:%s", request.LoginName, err.Error())
		return nil, err
	}
	//1 正常, 2 停用
	if resp.UserStatus != 1 {
		err = errors.New("用户已经停用")
		iotlogger.LogHelper.Errorf("PasswordLogin: 用户%s登录失败，原因:%s", request.LoginName, err.Error())
		return nil, err
	}

	UserPassword := iotutil.Md5(request.Password + resp.UserSalt)
	if UserPassword != resp.UserPassword {
		err = errors.New("用户名或密码错误")
		iotlogger.LogHelper.Errorf("PasswordLogin: 用户%s登录失败，原因:%s", request.LoginName, err.Error())
		return nil, err
	}
	var userinfo CloudUserInfo
	userinfo.UserID = resp.Id
	userinfo.Nickname = resp.UserNickname
	userinfo.Avatar = resp.Avatar
	//todo 用户角色和用户部门待完成
	ret, err := createCloudToken(&userinfo)
	if err != nil {
		iotlogger.LogHelper.Errorf("创建token失败:%s", err.Error())
		return nil, err
	}
	//将登录状态写入到sys_user_online
	saveUserOnline(request, ret)
	return ret, nil
}

func (s *CloudAuthSvc) Logout(request *proto.LogoutRequest) (*proto.LogoutResponse, error) {
	return nil, errors.New("not implement")
}

func (s *CloudAuthSvc) VerifyToken(request *proto.VerifyTokenRequest) (*proto.CloudVerifyTokenResponse, error) {
	ret := proto.CloudVerifyTokenResponse{Valid: false}
	cloudClaims := &CloudClaims{}
	ret.Valid = cloudClaims.VerifyToken(request.Token)
	if ret.Valid {
		ret.ExpiresAt = cloudClaims.ExpiresAt
		ret.UserInfo = &proto.CloudUserInfo{
			UserId:   cloudClaims.UserID,
			NickName: cloudClaims.Nickname,
			Avatar:   cloudClaims.Avatar,
			DeptId:   cloudClaims.DeptId,
			RoleIds:  strings.Split(cloudClaims.RoleIds, ","),
			PostIds:  strings.Split(cloudClaims.PostIds, ","),
		}
	}
	return &ret, nil
}

func (s *CloudAuthSvc) RefreshToken(request *proto.RefreshTokenRequest) (*proto.CloudRefreshTokenResponse, error) {
	cloudClaims := &CloudClaims{}
	bValid := cloudClaims.VerifyToken(request.GetRefreshToken())
	if !bValid {
		err := errors.New("refresh token 已失效")
		iotlogger.LogHelper.Errorf("RefreshToken: %s", err.Error())
		return &proto.CloudRefreshTokenResponse{Valid: false}, err
	}
	resp, err := createCloudToken(&cloudClaims.CloudUserInfo)
	if err != nil {
		iotlogger.LogHelper.Errorf("创建token失败:%s", err.Error())
		return nil, err
	}
	return &proto.CloudRefreshTokenResponse{
		Valid: true,
		Data:  resp,
	}, nil
}

func CloudUserInfo2pb(src *CloudUserInfo) *proto.CloudUserInfo {
	if src == nil {
		return nil
	}
	pbObj := proto.CloudUserInfo{
		UserId:   src.UserID,
		NickName: src.Nickname,
		Avatar:   src.Avatar,
		DeptId:   src.DeptId,
		RoleIds:  strings.Split(src.RoleIds, ","),
		PostIds:  strings.Split(src.PostIds, ","),
	}
	return &pbObj
}

func createCloudToken(userInfo *CloudUserInfo) (*proto.CloudLoginResponse, error) {
	cloudClaims := &CloudClaims{jwt.StandardClaims{}, *userInfo}
	//创建访问token
	tokenStr, err1 := cloudClaims.GenerateToken()
	if err1 != nil {
		return nil, err1
	}
	expiresAt := cloudClaims.ExpiresAt
	//创建刷新token
	refreshTokenStr, err2 := cloudClaims.GenerateRefreshToken()
	if err2 != nil {
		return nil, err2
	}
	ret := proto.CloudLoginResponse{
		Token:        tokenStr,
		RefreshToken: refreshTokenStr,
		UserInfo:     CloudUserInfo2pb(userInfo),
		ExpiresAt:    expiresAt,
	}
	return &ret, nil
}

func saveUserOnline(request *proto.PasswordLoginRequest, resq *proto.CloudLoginResponse) {
	defer iotutil.PanicHandler(resq)
	suo := SysUserOnlineSvc{context.Background()}
	_, err := suo.CreateSysUserOnline(&proto.SysUserOnline{
		Uuid:     iotutil.ToString(resq.UserInfo.UserId),
		Token:    resq.Token,
		UserName: resq.UserInfo.NickName,
		Ip:       request.ClientIp,
		Explorer: request.Explorer,
		Os:       request.Os,
	})
	if err != nil {
		iotlogger.LogHelper.Errorf("保存用户在线状态失败:%s", err.Error())
		return
	}
}
