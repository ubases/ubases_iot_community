package service

import (
	"cloud_platform/iot_app_user_service/config"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	proto "cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/go-pay/gopay/wechat"
	"github.com/golang-jwt/jwt"
)

type AppAuthSvc struct {
	Ctx context.Context
}

func (s AppAuthSvc) SetContext(ctx context.Context) AppAuthSvc {
	s.Ctx = ctx
	return s
}

func (s *AppAuthSvc) MiniProgramLogin(request *proto.MiniProgramLoginRequest) (*proto.AppLoginResponse, error) {
	//根据渠道类型、code查出渠道信息(第三方userid和昵称)
	channelUserId, channelNickname, code, msg := GetMiniProgram(request)
	if code != 0 {
		return nil, errors.New(msg)
	}
	var appId, appSecret string
	appId = config.Global.ThirdPartyLogin.MiniProgram.AppId
	appSecret = config.Global.ThirdPartyLogin.MiniProgram.AppSecret
	obj, code, msg := s.GetChannelAuthData(&proto.AppThirdRequest{
		Code:           request.Code,
		ChannelId:      channelUserId,
		Mode:           iotutil.ToString(iotconst.WechatMiniProgram),
		Nickname:       channelNickname,
		Ip:             request.ClientIp,
		AppId:          appId,
		AppSecret:      appSecret,
		TenantId:       request.TenantId,
		AppKey:         request.AppKey,
		RegionServerId: request.RegionServerId,
	}, channelUserId, channelNickname, iotconst.WechatMiniProgram)
	if code != 200 {
		return nil, errors.New(msg)
	}
	return &obj, nil
}

func (s *AppAuthSvc) PhoneCodeLogin(request *proto.PhoneCodeLoginRequest) (*proto.AppLoginResponse, error) {
	return nil, errors.New("not implement")
}

func (s *AppAuthSvc) EmailCodeLogin(request *proto.EmailCodeLoginRequest) (*proto.AppLoginResponse, error) {
	return nil, errors.New("not implement")
}

// checkUserName 检查用户名称，判断为手机、邮箱
func (s *AppAuthSvc) checkUserName(userName string) (phone, email string, err error) {
	if iotutil.CheckAllPhone("", userName) {
		phone = userName
	} else if iotutil.IsEmail(userName) {
		email = userName
	} else {
		err = errors.New("无法识别账号格式")
	}
	return
}

// queryUserByLogin  通过登录信息查询用户信息
func (s *AppAuthSvc) queryUserByLogin(loginName string, req proto.UcUserByLoginRequest) (res []*proto.UcUser, err error) {
	//通过手机号码/邮箱查询用户信息，因为存在注销账号所以有可能会返回多条
	ucs := UcUserSvc{s.Ctx}
	resp, err := ucs.GetUserByLogin(&req)
	if err != nil {
		iotlogger.LogHelper.Errorf("PasswordLogin: 用户%s登录失败，原因:%s", loginName, err.Error())
		return nil, err
	}
	if len(resp) == 0 {
		err = errors.New("用户不存在")
		iotlogger.LogHelper.Errorf("PasswordLogin: 用户%s登录失败，原因:%s", loginName, err.Error())
		return nil, err
	}
	return resp, nil
}

// queryThirdPartyLogin 通过登录信息查询用户第三方登录信息
func (s *AppAuthSvc) queryThirdPartyLogin(userId int64) (res []*proto.UcUserThird, err error) {
	//resp, err := rpcclient.ClientUcUserThirdService.Lists(context.Background(), &proto.UcUserThirdListRequest{Query: &proto.UcUserThird{UserId: userId}})
	ucts := UcUserThirdSvc{Ctx: s.Ctx}
	data, _, err := ucts.GetListUcUserThird(&proto.UcUserThirdListRequest{Query: &proto.UcUserThird{UserId: userId}})
	if err != nil {
		iotlogger.LogHelper.Errorf("读取用户所有第三方登录方式失败，原因:%s", err.Error())
		return nil, err
	}
	return data, nil
}

// pushLoginSuccess 推送登录成功
func (s *AppAuthSvc) pushLoginSuccess(user *proto.UcUserLoginSuccessRequest) (err error) {
	ucs := UcUserSvc{s.Ctx}
	err = ucs.LoginSuccess(user)
	if err != nil {
		return err
	}

	return nil
}

// 检查账号
func (s *AppAuthSvc) verityAccount(userlist []*proto.UcUser) (userinfo *proto.UcUser, err error) {
	hasCancel := false
	for _, user := range userlist {
		if user.Status == iotconst.ACCOUNT_NORMAL || user.Status == iotconst.ACCOUNT_CANCELING {
			userinfo = user
			break
		} else if user.Status == iotconst.ACCOUNT_CANCELED {
			hasCancel = true
		}
	}
	if userinfo == nil {
		if hasCancel {
			err = errors.New("账号已注销")
		} else {
			err = errors.New("账号或者密码错误")
		}
	}
	return
}

func (s *AppAuthSvc) PasswordLogin(request *proto.PasswordLoginRequest) (*proto.AppLoginResponse, error) {
	var (
		ip string = ""
	)
	phone, email, err := s.checkUserName(request.LoginName)
	if err != nil {
		return nil, err
	}

	req := proto.UcUserByLoginRequest{
		Phone:          phone,
		Email:          email,
		AppKey:         request.AppKey,
		TenantId:       request.TenantId,
		RegionServerId: request.RegionServerId,
		//Password: request.Password,
		//Status:   1, //1:正常，2:待注销，3:已注销
	}
	//通过手机号码/邮箱查询用户信息，因为存在注销账号所以有可能会返回多条
	resp, err := s.queryUserByLogin(request.LoginName, req)
	iotlogger.LogHelper.Helper.Debugf("resp: %v, err: %v", resp, err)
	if err != nil {
		return nil, err
	}
	//账号检查
	dbUser, err := s.verityAccount(resp)
	if err != nil {
		return nil, err
	}
	var accountCasser = false                // 账号待注销状态，如果为true则App将弹出注销恢复提醒
	var submitCancelTime = dbUser.CancelTime // 原注销提交时间，前端提示使用
	if dbUser.Status == iotconst.ACCOUNT_CANCELING {
		accountCasser = true
	}

	//非验证码登录要验证密码
	if len(request.VerifyCode) == 0 {
		loginPwd := iotutil.Md5(request.Password + dbUser.UserSalt) // + dbUser.Uid)
		iotlogger.LogHelper.Helper.Debugf("loginPwd: %v, Password: %v", loginPwd, dbUser.Password)
		if loginPwd != dbUser.Password {
			return nil, errors.New("账号或密码错误")
		}
	}

	//读取用户所有第三方登录方式
	dbThirdPartyInfo, err := s.queryThirdPartyLogin(dbUser.Id)
	if err != nil {
		//iotlogger.LogHelper.Errorf("创建token失败:%s", err.Error())
		return nil, err
	}

	var thirdPartys = make([]*proto.AppUserThirdPartyLogin, 0)
	for _, thirdPartyLogin := range dbThirdPartyInfo {
		thirdPartys = append(thirdPartys, &proto.AppUserThirdPartyLogin{
			Mode: thirdPartyLogin.ThirdType, LoginKey: thirdPartyLogin.ThirdUserId, Nickname: thirdPartyLogin.Nickname})
	}

	//if dbUser.OpenId != "" {
	//	thirdPartys = append(thirdPartys, &proto.AppUserThirdPartyLogin{
	//		Mode: iotconst.THIRD_PARTY_WECHAT, LoginKey: dbUser.OpenId, NickName: dbUser.WechatNickName})
	//}
	//if dbUser.AppleIdUserId != "" {
	//	thirdPartys = append(thirdPartys, &proto.AppUserThirdPartyLogin{
	//		Mode: iotconst.THIRD_PARTY_APPLEID, LoginKey: dbUser.AppleIdUserId, NickName: dbUser.AppleIdNickName})
	//}
	//登录用户信息
	var user = AppUserInfo{
		UserID:         dbUser.Id,
		Nickname:       dbUser.NickName,
		Avatar:         dbUser.Photo,
		Account:        dbUser.UserName, // request.LoginName,
		AppKey:         dbUser.AppKey,
		TenantId:       dbUser.TenantId,
		RegionServerId: dbUser.RegionServerId,
	}
	ret, err := createAppToken(&user)
	if err != nil {
		iotlogger.LogHelper.Errorf("创建token失败:%s", err.Error())
		return nil, err
	}

	//登录成功，异步调用登录成功
	s.pushLoginSuccess(&proto.UcUserLoginSuccessRequest{
		User:          dbUser,
		Ip:            ip,
		AccountCasser: accountCasser,
		Token:         ret.Token,
	})

	//返回登录用户信息
	ret.UserInfo = &proto.AppUser{
		Id:               dbUser.Id,
		NickName:         dbUser.NickName,
		Phone:            dbUser.Phone,
		Photo:            dbUser.Photo,
		Status:           dbUser.Status,
		City:             dbUser.City,
		Gender:           dbUser.Gender,
		Email:            dbUser.Email,
		UserName:         dbUser.UserName,
		Account:          dbUser.UserName,
		DefaultHomeId:    dbUser.DefaultHomeId,
		RegisterRegion:   dbUser.RegisterRegion,
		AccountCasser:    accountCasser,
		SubmitCancelTime: submitCancelTime,
		ThirdPartyLogin:  thirdPartys,
		Password:         dbUser.Password,
		RegionServerId:   dbUser.RegionServerId,
	}

	return ret, nil
}

func AppUserInfo2pb(src *AppUserInfo) *proto.AppUser {
	if src == nil {
		return nil
	}
	pbObj := proto.AppUser{
		Id:             src.UserID,
		NickName:       src.Nickname,
		Photo:          src.Avatar,
		Account:        src.Account,
		RegionServerId: src.RegionServerId,
		AppKey:         src.AppKey,
		TenantId:       src.TenantId,
		//HomeIds:  strings.Split(src.HomeIds, ","),
	}
	return &pbObj
}

func createAppToken(userInfo *AppUserInfo) (*proto.AppLoginResponse, error) {
	appClaims := &AppClaims{jwt.StandardClaims{}, *userInfo}
	//创建访问token
	tokenStr, err1 := appClaims.GenerateToken()
	if err1 != nil {
		return nil, err1
	}
	expiresAt := appClaims.ExpiresAt
	//创建刷新token
	refreshTokenStr, err2 := appClaims.GenerateRefreshToken()
	if err2 != nil {
		return nil, err2
	}
	ret := proto.AppLoginResponse{
		Token:        tokenStr,
		RefreshToken: refreshTokenStr,
		UserInfo:     AppUserInfo2pb(userInfo),
		ExpiresAt:    expiresAt,
	}
	return &ret, nil
}

// UserList中获取有效状态的用户信息
func ValidUserByList(userlist []*proto.UcUser) (userinfo proto.UcUser, accountstate int32) {
	for _, user := range userlist {
		accountstate = user.Status
		if accountstate == 1 || accountstate == 2 {
			userinfo = *user
			break
		}
	}
	return
}

func (s *AppAuthSvc) Logout(request *proto.LogoutRequest) (*proto.LogoutResponse, error) {
	return nil, errors.New("not implement")
}

func (s *AppAuthSvc) VerifyToken(request *proto.VerifyTokenRequest) (*proto.APPVerifyTokenResponse, error) {
	bValid := (&AppClaims{}).VerifyToken(request.Token)
	return &proto.APPVerifyTokenResponse{Valid: bValid}, nil
}

func (s *AppAuthSvc) RefreshToken(request *proto.RefreshTokenRequest) (*proto.AppRefreshTokenResponse, error) {
	appClaims := &AppClaims{}
	bValid := appClaims.VerifyToken(request.GetRefreshToken())
	if !bValid {
		err := errors.New("refresh token 已失效")
		iotlogger.LogHelper.Errorf("RefreshToken: %s", err.Error())
		return &proto.AppRefreshTokenResponse{Valid: false}, err
	}
	resp, err := createAppToken(&appClaims.AppUserInfo)
	if err != nil {
		iotlogger.LogHelper.Errorf("创建token失败:%s", err.Error())
		return nil, err
	}
	return &proto.AppRefreshTokenResponse{
		Valid: true,
		Data:  resp,
	}, nil
}

func (s *AppAuthSvc) WechatLogin(request *proto.AppThirdRequest) (*proto.AppLoginResponse, error) {
	//根据渠道类型、code查出渠道信息(第三方userid和昵称)
	channelUserId, channelNickname, code, msg := GetWechatInfo(request)
	if code != 0 {
		return nil, errors.New(msg)
	}
	obj, code, msg := s.GetChannelAuthData(request, channelUserId, channelNickname, iotconst.Wechat)
	if code != 200 {
		return nil, errors.New(msg)
	}
	return &obj, nil
}

// 获取微信渠道信息
func GetWechatInfo(request *proto.AppThirdRequest) (channelUserId string, channelNickname string, msgCode int, msg string) {
	if request.Code == "" {
		msgCode = 100024
		msg = "授权Code为空"
		return
	}

	var appId, appSecret string
	appId = request.AppId
	appSecret = request.AppSecret
	iotlogger.LogHelper.Info("appId---" + appId)
	iotlogger.LogHelper.Info("appSecret---" + appSecret)

	//appId, appSecret, authorizationCode查询accessToken
	accessToken, err := wechat.GetOauth2AccessToken(context.Background(), appId, appSecret, request.Code)
	if err != nil || accessToken.Openid == "" {
		msgCode = 100024
		msg = "获取openId失败"
		return
	}
	//accessToken查询userInfo
	userInfo, _ := wechat.GetOauth2UserInfo(context.Background(), accessToken.AccessToken, accessToken.Openid, "")
	if accessToken.Openid != "" && userInfo != nil {
		channelUserId = accessToken.Openid
		channelNickname = userInfo.Nickname
	}
	return
}

func (s *AppAuthSvc) AppleidLogin(request *proto.AppThirdRequest) (*proto.AppLoginResponse, error) {
	//根据渠道类型、code查出渠道信息(第三方userid和昵称)
	channelUserId, channelNickname, code, _ := s.GetAppleidInfo(request.ChannelId, request.Nickname, request.Ip, request.HomeName, request.RegionServerId)
	if code != 0 {
		return nil, nil
	}
	obj, code, _ := s.GetChannelAuthData(request, channelUserId, channelNickname, iotconst.Appleid)
	return &obj, nil
}

// 获取Appleid渠道信息
func (s *AppAuthSvc) GetAppleidInfo(channelId string, nickname, ip, homeName string, regionServerId int64) (channelUserId string, channelNickname string, msgCode int, msg string) {
	//appKey, err := CheckAppKey(s.Ctx)
	//if err != nil {
	//	msgCode = 100024
	//	msg = "appKey为空"
	//	return
	//}
	//tenantId, err := CheckTenantId(s.Ctx)
	//if err != nil {
	//	msgCode = 100024
	//	msg = "tenantId为空"
	//	return
	//}
	//appleid登录
	if channelId == "" {
		msgCode = 100024
		msg = "第三方userid为空"
		return
	}
	channelUserId = channelId
	ucais := UcAppleidInfoSvc{s.Ctx}
	if nickname != "" {
		channelNickname = nickname
		AppleidInfo, _ := ucais.FindUcAppleidInfo(&proto.UcAppleidInfoFilter{
			ThirdUserId: channelUserId,
		})
		if AppleidInfo != nil {
			_, err := ucais.UpdateUcAppleidInfo(&proto.UcAppleidInfo{
				ThirdUserId: channelUserId,
				Nickname:    channelNickname,
			})
			if err != nil {
				msgCode = 100024
				msg = err.Error()
				return
			}
		} else {
			_, err := ucais.CreateUcAppleidInfo(&proto.UcAppleidInfo{
				Id:          iotutil.GetNextSeqInt64(),
				ThirdUserId: channelUserId,
				Nickname:    channelNickname,
				CreatedAt:   timestamppb.Now(),
				UpdatedAt:   timestamppb.Now(),
			})
			if err != nil {
				msgCode = 100024
				msg = err.Error()
				return
			}
		}

	} else {
		AppleidInfo, _ := ucais.FindUcAppleidInfo(&proto.UcAppleidInfoFilter{
			ThirdUserId: channelUserId,
		})
		if AppleidInfo != nil {
			channelNickname = AppleidInfo.Nickname
		}
	}

	return
}

// 获取ChannelAuth数据  *proto.AppThirdResponse
func (s *AppAuthSvc) GetChannelAuthData(request *proto.AppThirdRequest, channelUserId string, channelNickname string, channelType int32) (ret proto.AppLoginResponse, msgCode int, msg string) {
	ucts := UcUserThirdSvc{Ctx: s.Ctx}
	ucs := UcUserSvc{s.Ctx}
	msgCode, msg = 0, ""
	userThirdResp, _ := ucts.FindUcUserThird(&proto.UcUserThirdFilter{
		ThirdType:      channelType,
		ThirdUserId:    channelUserId,
		AppKey:         request.AppKey,
		TenantId:       request.TenantId,
		RegionServerId: request.RegionServerId,
	})

	result := &proto.AppLoginResponse{}
	//增加开关，默认为不需要默认密码
	//if !config.Global.Service.ThirdPartyRequirePwd {
	//默认创建用户
	if userThirdResp == nil {
		iotlogger.LogHelper.Info("第三方登录自动注册账号， channelUserId：%s", channelUserId)
		resR, err := ucs.Register(&proto.UcUserRegisterRequest{
			Phone:           "",
			Password:        iotutil.EncodeMD5(iotutil.EncodeMD5(iotutil.GetSecret(6))),
			Email:           "",
			Code:            "",
			RegisterRegion:  "",
			Ip:              request.Ip,
			ThirdType:       iotutil.ToString(channelType),
			ThirdUserId:     channelUserId,
			ThirdNickname:   channelNickname,
			AppKey:          request.AppKey,
			TenantId:        request.TenantId,
			RegionServerId:  request.RegionServerId,
			DefaultHomeName: request.HomeName,
		})
		if err != nil {
			return *result, -1, err.Error()
		}
		if resR == nil {
			return *result, -1, "failed"
		}
		userThirdResp, _ = ucts.FindUcUserThird(&proto.UcUserThirdFilter{
			ThirdType:      channelType,
			ThirdUserId:    channelUserId,
			AppKey:         request.AppKey,
			TenantId:       request.TenantId,
			RegionServerId: request.RegionServerId,
		})
	}
	//}

	if userThirdResp == nil {
		//没有绑定用户请前往绑定
		var thirdPartys = make([]*proto.AppUserThirdPartyLogin, 0)
		thirdPartys = append(thirdPartys, &proto.AppUserThirdPartyLogin{Mode: iotutil.ToInt32(channelType), LoginKey: channelUserId, Nickname: channelNickname})
		ret = proto.AppLoginResponse{
			Token:        "",
			RefreshToken: "",
			UserInfo: &proto.AppUser{
				ThirdPartyLogin: thirdPartys,
			},
			ExpiresAt: 0,
		}
		return ret, 200, ""
	} else { //已经绑定用户信息的直接登录
		userId := userThirdResp.UserId

		req := proto.UcUserFilter{
			Id:     userId,
			Status: 1,
		}
		//通过手机号码/邮箱查询用户信息，因为存在注销账号所以有可能会返回多条
		resp, err := ucs.FindByIdUcUser(&req)
		if err != nil {
			//iotlogger.LogHelper.Errorf("PasswordLogin: 用户%s登录失败，原因:%s", loginName, err.Error())
			return *result, -1, err.Error()
		}
		if resp == nil {
			//恢复数据作用，正常是不会过来的，历史数据出现账号注销之后第三方信息未注销的问题；
			//rpcclient.ClientUcUserThirdService.Delete(context.Background(), &proto.UcUserThird{Id: userThirdResp.Data[0].Id})
			err = errors.New("用户不存在")
			return *result, -1, err.Error()
		}

		//读取用户所有第三方登录方式
		dbThirdPartyInfo, err := s.queryThirdPartyLogin(userId)
		if err != nil {
			//iotlogger.LogHelper.Errorf("创建token失败:%s", err.Error())
			return *result, -1, err.Error()
		}

		var thirdPartys = make([]*proto.AppUserThirdPartyLogin, 0)
		for _, thirdPartyLogin := range dbThirdPartyInfo {
			thirdPartys = append(thirdPartys, &proto.AppUserThirdPartyLogin{
				Mode: thirdPartyLogin.ThirdType, LoginKey: thirdPartyLogin.ThirdUserId, Nickname: thirdPartyLogin.Nickname})
		}

		userInfo := resp

		//登录用户信息
		var user = AppUserInfo{
			UserID:         userInfo.Id,
			Nickname:       userInfo.NickName,
			Avatar:         userInfo.Photo,
			Account:        userInfo.UserName,
			RegionServerId: userInfo.RegionServerId,
			AppKey:         userInfo.AppKey,
			TenantId:       userInfo.TenantId,
		}

		result, err = createAppToken(&user)
		if err != nil {
			iotlogger.LogHelper.Errorf("创建token失败:%s", err.Error())
			return *result, -1, err.Error()
		}

		//登录成功，异步调用登录成功
		//s.pushLoginSuccess(&proto.UcUserLoginSuccessRequest{
		//	User:          dbUser,
		//	Ip:            ip,
		//	AccountCasser: accountCasser,
		//	Token:         ret.Token,
		//})

		var accountCasser = false                  // 账号待注销状态，如果为true则App将弹出注销恢复提醒
		var submitCancelTime = userInfo.CancelTime // 原注销提交时间，前端提示使用
		if userInfo.Status == iotconst.ACCOUNT_CANCELING {
			accountCasser = true
		}

		//返回登录用户信息
		result.UserInfo = &proto.AppUser{
			Id:               userInfo.Id,
			NickName:         userInfo.NickName,
			Phone:            userInfo.Phone,
			Photo:            userInfo.Photo,
			Status:           userInfo.Status,
			City:             userInfo.City,
			Gender:           userInfo.Gender,
			Email:            userInfo.Email,
			DefaultHomeId:    userInfo.DefaultHomeId,
			RegisterRegion:   userInfo.RegisterRegion,
			AccountCasser:    accountCasser,
			SubmitCancelTime: submitCancelTime,
			ThirdPartyLogin:  thirdPartys,
			Password:         userInfo.Password,
			UserName:         userInfo.UserName,
			Account:          userInfo.UserName,
			AppKey:           userInfo.AppKey,
			TenantId:         userInfo.TenantId,
			RegionServerId:   userInfo.RegionServerId,
		}

		//todo 待处理
		//if  channelNickname!= "" { //appleid第一次登录须缓存用户昵称
		//	err := table.Update(bson.M{"_id": user.Id}, bson.M{"$set": bson.M{channelName: channelNickname}})
		//	if err!=nil {
		//		global.GVA_LOG.Error("绑定渠道信息失败")
		//		msgCode = 100002
		//		msg = "绑定渠道信息失败"
		//		return
		//	}
		//}

		//已经绑定用户信息直接登录
		//obj = map[string]interface{}{"openid": channelUserId,"token": ret.Token,"user": ret.UserInfo,"type":channelType}
		msgCode = 200
		msg = "ok"
	}
	return *result, msgCode, msg
}

func (s *AppAuthSvc) GetUserAndToken(userId int64, nickname, avatar, account string) (*proto.AppLoginResponse, error) {
	//读取用户所有第三方登录方式
	dbThirdPartyInfo, err := s.queryThirdPartyLogin(userId)
	if err != nil {
		//iotlogger.LogHelper.Errorf("创建token失败:%s", err.Error())
		return nil, err
	}

	var thirdPartys = make([]*proto.AppUserThirdPartyLogin, 0)
	for _, thirdPartyLogin := range dbThirdPartyInfo {
		thirdPartys = append(thirdPartys, &proto.AppUserThirdPartyLogin{
			Mode: thirdPartyLogin.ThirdType, LoginKey: thirdPartyLogin.ThirdUserId, Nickname: thirdPartyLogin.Nickname})
	}

	//登录用户信息
	var user = AppUserInfo{
		UserID:   userId,
		Nickname: nickname,
		Avatar:   avatar,
		Account:  account,
	}
	ret, err := createAppToken(&user)
	if err != nil {
		iotlogger.LogHelper.Errorf("创建token失败:%s", err.Error())
		return nil, err
	}

	//登录成功，异步调用登录成功
	s.pushLoginSuccess(&proto.UcUserLoginSuccessRequest{
		User: &proto.UcUser{
			Id: userId,
		},
		Ip:            "",
		AccountCasser: false,
		Token:         ret.Token,
	})

	//返回登录用户信息
	ret.UserInfo = &proto.AppUser{
		Id:               0,
		NickName:         "",
		Phone:            "",
		Photo:            "",
		Status:           0,
		City:             "",
		Gender:           0,
		Email:            "",
		UserName:         "",
		Account:          "",
		DefaultHomeId:    "",
		RegisterRegion:   "",
		AccountCasser:    false,
		SubmitCancelTime: 0,
		ThirdPartyLogin:  thirdPartys,
	}

	return ret, nil
}

func (s *AppAuthSvc) GetTokenByAccount(request *proto.GetTokenByAccountRequest) (*proto.AppLoginResponse, error) {
	var (
		ip string = ""
	)
	phone, email, err := s.checkUserName(request.LoginName)
	if err != nil {
		return nil, err
	}

	req := proto.UcUserByLoginRequest{
		Phone:          phone,
		Email:          email,
		AppKey:         request.AppKey,
		TenantId:       request.TenantId,
		RegionServerId: request.RegionServerId,
		//Password:       iotutil.Md5(request.Password),
		Status: 1, //1:正常，2:待注销，3:已注销
	}
	//通过手机号码/邮箱查询用户信息，因为存在注销账号所以有可能会返回多条
	resp, err := s.queryUserByLogin(request.LoginName, req)
	if err != nil {
		return nil, err
	}

	theUser := resp[0]
	if theUser.Password != iotutil.Md5(request.Password+theUser.UserSalt) {
		return nil, errors.New("用户名密码错误")
	}

	//账号检查
	dbUser, err := s.verityAccount(resp)
	if err != nil {
		return nil, err
	}
	var accountCasser = false                // 账号待注销状态，如果为true则App将弹出注销恢复提醒
	var submitCancelTime = dbUser.CancelTime // 原注销提交时间，前端提示使用
	if dbUser.Status == iotconst.ACCOUNT_CANCELING {
		accountCasser = true
	}

	//密码加密
	//loginPwd := iotutil.Md5(request.Password) // + dbUser.Uid)
	//if loginPwd != dbUser.Password {
	//	return nil, errors.New("账号或密码错误")
	//}

	//读取用户所有第三方登录方式
	dbThirdPartyInfo, err := s.queryThirdPartyLogin(dbUser.Id)
	if err != nil {
		//iotlogger.LogHelper.Errorf("创建token失败:%s", err.Error())
		return nil, err
	}

	var thirdPartys = make([]*proto.AppUserThirdPartyLogin, 0)
	for _, thirdPartyLogin := range dbThirdPartyInfo {
		thirdPartys = append(thirdPartys, &proto.AppUserThirdPartyLogin{
			Mode: thirdPartyLogin.ThirdType, LoginKey: thirdPartyLogin.ThirdUserId, Nickname: thirdPartyLogin.Nickname})
	}

	//if dbUser.OpenId != "" {
	//	thirdPartys = append(thirdPartys, &proto.AppUserThirdPartyLogin{
	//		Mode: iotconst.THIRD_PARTY_WECHAT, LoginKey: dbUser.OpenId, NickName: dbUser.WechatNickName})
	//}
	//if dbUser.AppleIdUserId != "" {
	//	thirdPartys = append(thirdPartys, &proto.AppUserThirdPartyLogin{
	//		Mode: iotconst.THIRD_PARTY_APPLEID, LoginKey: dbUser.AppleIdUserId, NickName: dbUser.AppleIdNickName})
	//}
	//登录用户信息
	var user = AppUserInfo{
		UserID:   dbUser.Id,
		Nickname: dbUser.NickName,
		Avatar:   dbUser.Photo,
		//Account:  request.LoginName,
		Account:  dbUser.UserName,
		AppKey:   dbUser.AppKey,
		TenantId: dbUser.TenantId,
	}
	ret, err := createAppToken(&user)
	if err != nil {
		iotlogger.LogHelper.Errorf("创建token失败:%s", err.Error())
		return nil, err
	}

	//登录成功，异步调用登录成功
	s.pushLoginSuccess(&proto.UcUserLoginSuccessRequest{
		User:          dbUser,
		Ip:            ip,
		AccountCasser: accountCasser,
		Token:         ret.Token,
	})

	//返回登录用户信息
	ret.UserInfo = &proto.AppUser{
		Id:               dbUser.Id,
		NickName:         dbUser.NickName,
		Phone:            dbUser.Phone,
		Photo:            dbUser.Photo,
		Status:           dbUser.Status,
		City:             dbUser.City,
		Gender:           dbUser.Gender,
		Email:            dbUser.Email,
		UserName:         dbUser.UserName,
		Account:          dbUser.UserName,
		DefaultHomeId:    dbUser.DefaultHomeId,
		RegisterRegion:   dbUser.RegisterRegion,
		AccountCasser:    accountCasser,
		SubmitCancelTime: submitCancelTime,
		ThirdPartyLogin:  thirdPartys,
	}

	return ret, nil
}
