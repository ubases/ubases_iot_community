package service

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	proto "cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"

	"github.com/golang-jwt/jwt"
)

type OpenAuthSvc struct {
	Ctx context.Context
}

func (s *OpenAuthSvc) MiniProgramLogin(request *proto.MiniProgramLoginRequest) (*proto.CloudLoginResponse, error) {
	return nil, errors.New("not implement")
}

func (s *OpenAuthSvc) PhoneCodeLogin(request *proto.PhoneCodeLoginRequest) (*proto.CloudLoginResponse, error) {
	return nil, errors.New("not implement")
}

func (s *OpenAuthSvc) EmailCodeLogin(request *proto.EmailCodeLoginRequest) (*proto.CloudLoginResponse, error) {
	return nil, errors.New("not implement")
}

func (s *OpenAuthSvc) PasswordLogin(request *proto.PasswordLoginRequest) (*proto.CloudLoginResponse, error) {
	if request.LoginName == "" || request.Password == "" {
		err := errors.New("用户名或密码为空")
		iotlogger.LogHelper.Errorf("PasswordLogin: 用户%s登录失败，原因:%s", request.LoginName, err.Error())
		return nil, err
	}
	ou := OpenUserSvc{s.Ctx}
	req := proto.OpenUserFilter{UserName: request.LoginName}
	resp, err := ou.FindOpenUser(&req)
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
	var userinfo OpenUserInfo
	userinfo.UserID = resp.Id
	userinfo.Nickname = resp.UserNickname
	userinfo.Avatar = resp.Avatar
	userinfo.AccountType = resp.AccountType
	//userinfo.TenantId = resp.Data[0].TenantId
	//TODO 需要返回用户所有的空间
	ouc := OpenUserCompanySvc{s.Ctx}
	respCompanys, _, err := ouc.GetListOpenUserCompany(&proto.OpenUserCompanyListRequest{
		Query: &proto.OpenUserCompany{UserId: userinfo.UserID},
	})
	if err != nil {
		return nil, err
	}
	if len(respCompanys) == 0 {
		return nil, errors.New("当前用户数据异常，请联系平台客服！")
	}
	//获取主账号（每个开放者必定一有个主账号），登录之后默认进入主账号
	for _, comp := range respCompanys {
		if comp.UserId == userinfo.UserID && comp.UserType == iotconst.OPEN_USER_MAIN_ACCOUNT {
			userinfo.TenantId = comp.TenantId
			ocs := OpenCompanySvc{Ctx: s.Ctx}
			if oc, err := ocs.FindOpenCompanyNoCtx(&proto.OpenCompanyFilter{TenantId: userinfo.TenantId}); err == nil && oc != nil {
				userinfo.Company = oc.Name
			}
			break
		}
	}
	//如果没有获取到主账号，返回提示数据异常
	if userinfo.TenantId == "" {
		return nil, errors.New("当前用户数据异常，请联系平台客服！！！")
	}
	ret, err := createOpenToken(&userinfo)
	if err != nil {
		iotlogger.LogHelper.Errorf("创建token失败:%s", err.Error())
		return nil, err
	}
	saveOpenUserOnline(request, ret)
	return ret, nil
}

func (s *OpenAuthSvc) Logout(request *proto.LogoutRequest) (*proto.LogoutResponse, error) {
	return nil, errors.New("not implement")
}

func (s *OpenAuthSvc) VerifyToken(request *proto.VerifyTokenRequest) (*proto.CloudVerifyTokenResponse, error) {
	ret := proto.CloudVerifyTokenResponse{Valid: false}
	openClaims := &OpenClaims{}
	ret.Valid = openClaims.VerifyToken(request.Token)
	if ret.Valid {
		ret.ExpiresAt = openClaims.ExpiresAt
		ret.UserInfo = &proto.CloudUserInfo{
			UserId:   openClaims.UserID,
			NickName: openClaims.Nickname,
			Avatar:   openClaims.Avatar,
			Company:  openClaims.Company,
		}
	}
	return &ret, nil
}

func (s *OpenAuthSvc) RefreshToken(request *proto.RefreshTokenRequest) (*proto.CloudRefreshTokenResponse, error) {
	openClaims := &OpenClaims{}
	bValid := openClaims.VerifyToken(request.GetRefreshToken())
	if !bValid {
		err := errors.New("refresh token 已失效")
		iotlogger.LogHelper.Errorf("RefreshToken: %s", err.Error())
		return &proto.CloudRefreshTokenResponse{Valid: false}, err
	}
	resp, err := createOpenToken(&openClaims.OpenUserInfo)
	if err != nil {
		iotlogger.LogHelper.Errorf("创建token失败:%s", err.Error())
		return nil, err
	}
	return &proto.CloudRefreshTokenResponse{
		Valid: true,
		Data:  resp,
	}, nil
}

func (s *OpenAuthSvc) ChangeTenant(request *proto.RefreshTokenRequest) (*proto.CloudRefreshTokenResponse, error) {
	openClaims := &OpenClaims{}
	bValid := openClaims.VerifyToken(request.GetRefreshToken())
	if !bValid {
		err := errors.New("refresh token 已失效")
		iotlogger.LogHelper.Errorf("RefreshToken: %s", err.Error())
		return &proto.CloudRefreshTokenResponse{Valid: false}, err
	}
	//变更租户ID. 重新生成token 返回
	openClaims.OpenUserInfo.TenantId = request.TenantId
	resp, err := createOpenToken(&openClaims.OpenUserInfo)
	if err != nil {
		iotlogger.LogHelper.Errorf("创建token失败:%s", err.Error())
		return nil, err
	}
	return &proto.CloudRefreshTokenResponse{
		Valid: true,
		Data:  resp,
	}, nil
}

func OpenUserInfo2pb(src *OpenUserInfo) *proto.CloudUserInfo {
	if src == nil {
		return nil
	}
	pbObj := proto.CloudUserInfo{
		UserId:   src.UserID,
		NickName: src.Nickname,
		Avatar:   src.Avatar,
		TenantId: src.TenantId,
		Company:  src.Company,
	}
	return &pbObj
}

func createOpenToken(userInfo *OpenUserInfo) (*proto.CloudLoginResponse, error) {
	openClaims := &OpenClaims{jwt.StandardClaims{}, *userInfo}
	//创建访问token
	tokenStr, err1 := openClaims.GenerateToken()
	if err1 != nil {
		return nil, err1
	}
	expiresAt := openClaims.ExpiresAt
	//创建刷新token
	refreshTokenStr, err2 := openClaims.GenerateRefreshToken()
	if err2 != nil {
		return nil, err2
	}
	ret := proto.CloudLoginResponse{
		Token:        tokenStr,
		RefreshToken: refreshTokenStr,
		UserInfo:     OpenUserInfo2pb(userInfo),
		ExpiresAt:    expiresAt,
	}
	return &ret, nil
}

func saveOpenUserOnline(request *proto.PasswordLoginRequest, resq *proto.CloudLoginResponse) {
	defer iotutil.PanicHandler(resq)
	ouos := OpenUserOnlineSvc{context.Background()}
	_, err := ouos.CreateOpenUserOnline(&proto.OpenUserOnline{
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
