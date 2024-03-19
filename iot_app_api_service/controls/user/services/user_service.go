package services

import (
	"cloud_platform/iot_app_api_service/cached"
	"cloud_platform/iot_app_api_service/controls"
	_const "cloud_platform/iot_app_api_service/controls/user/const"
	"cloud_platform/iot_app_api_service/controls/user/entitys"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_common/iotgincache/persist"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/go-pay/gopay/wechat"
	"go-micro.dev/v4/metadata"
)

//var (
//	tokenUserKey = "%s_tokens"
//)

type AppUserService struct {
	Ctx context.Context
}

func (s AppUserService) SetContext(ctx context.Context) AppUserService {
	s.Ctx = ctx
	return s
}

// GetUserById 通过用户Id获取用户信息
func (s AppUserService) GetUserById(userid int64) (*protosService.UcUser, error) {
	res, err := rpc.TUcUserService.FindById(s.Ctx, &protosService.UcUserFilter{
		Id: userid,
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	if len(res.Data) == 0 {
		return nil, errors.New("未找到用户信息")
	}
	return res.Data[0], nil
}

// GetUser 获取用户信息
func (s AppUserService) GetUser(userid string) (*entitys.LoginUserRes, int, error) {
	result := entitys.LoginUserRes{}
	res, err := rpc.TUcUserService.FindById(s.Ctx, &protosService.UcUserFilter{
		Id: iotutil.ToInt64(userid),
	})
	if err != nil {
		return &result, -1, err
	}
	if res.Code != 200 {
		return &result, -1, errors.New(res.Message)
	}
	if len(res.Data) == 0 {
		return &result, -1, errors.New("未找到用户信息")
	}

	resp, err := rpc.ClientUcUserThirdService.Lists(s.Ctx, &protosService.UcUserThirdListRequest{Query: &protosService.UcUserThird{UserId: iotutil.ToInt64(userid)}})
	if err != nil {
		iotlogger.LogHelper.Errorf("读取用户所有第三方登录方式失败，原因:%s", err.Error())
		return &result, -1, err
	}

	user := res.Data[0]
	isGuest := false
	var thirdPartyLogin = entitys.AppUserThirdPartyLogin{}
	for _, thirdPartys := range resp.Data {
		switch thirdPartys.ThirdType {
		case iotconst.Wechat: //微信登录
			thirdPartyLogin.Wechat.Set(thirdPartys)
		case iotconst.Appleid: //苹果Id登录
			thirdPartyLogin.Wechat.Set(thirdPartys)
		case iotconst.Guest: //游客（不会同时存在）
			thirdPartyLogin.Guest.Set(thirdPartys)
			isGuest = true
		}
	}

	lang, _ := metadata.Get(s.Ctx, "lang")
	// 查询国家名称，省份名称，城市名称
	var countryName, provinceName, cityName string
	if len(user.Country) != 0 {
		id, err := iotutil.ToInt32Err(user.Country)
		if err != nil {
			return &result, -1, err
		}
		if id == 0 {
			countryName = user.Country
		} else {
			areaData, err := rpc.ClientAreaService.FindById(s.Ctx, &protosService.SysAreaFilter{
				Id: int64(id),
			})
			if err != nil {
				return &result, -1, err
			}
			if res.Code != 200 {
				return &result, -1, errors.New(areaData.Message)
			}
			countryName = getNameByLang(lang, areaData.Data[0].ChineseName, areaData.Data[0].EnglishName)
		}
	}
	if len(user.Province) != 0 {
		id, err := iotutil.ToInt32Err(user.Province)
		if err != nil {
			return &result, -1, err
		}
		if id == 0 {
			provinceName = user.Province
		} else {
			areaData, err := rpc.ClientAreaService.FindById(s.Ctx, &protosService.SysAreaFilter{
				Id: int64(id),
			})
			if err != nil {
				return &result, -1, err
			}
			if res.Code != 200 {
				return &result, -1, errors.New(areaData.Message)
			}
			provinceName = getNameByLang(lang, areaData.Data[0].ChineseName, areaData.Data[0].EnglishName)
		}
	}
	if len(user.City) != 0 {
		id, err := iotutil.ToInt32Err(user.City)
		if err != nil {
			return &result, -1, err
		}
		if id == 0 {
			cityName = user.City
		} else {
			areaData, err := rpc.ClientAreaService.FindById(s.Ctx, &protosService.SysAreaFilter{
				Id: int64(id),
			})
			if err != nil {
				return &result, -1, err
			}
			if res.Code != 200 {
				return &result, -1, errors.New(areaData.Message)
			}
			cityName = getNameByLang(lang, areaData.Data[0].ChineseName, areaData.Data[0].EnglishName)
		}
	}

	pbObj := &entitys.LoginUserRes{
		UserId:           iotutil.ToString(user.Id),
		NickName:         user.NickName,
		Phone:            user.Phone,
		Photo:            user.Photo,
		Email:            user.Email,
		DefaultHomeId:    user.DefaultHomeId,
		UserName:         user.UserName,
		SubmitCancelTime: user.CancelTime,
		AccountCasser:    false,
		ThirdPartyLogin:  thirdPartyLogin,
		Token:            "",
		RefreshToken:     "",
		ExpiresAt:        0,
		Password:         user.Password,
		Country:          user.Country,
		Province:         user.Province,
		City:             user.City,
		Gender:           user.Gender,
		Birthday:         iotutil.DateFormat(user.Birthday.AsTime().Local()),
		CountryName:      countryName,
		ProvinceName:     provinceName,
		CityName:         cityName,
		PasswordNotSet:   len(user.Password) == 0,
		IsGuest:          isGuest,
	}
	pbObj = pbObj.SetShowVconsole(getShowVconsole(user.Id))
	if user.Birthday.AsTime().IsZero() {
		pbObj.Birthday = ""
	}
	return pbObj, 0, nil
}

// 检查登录接口中的参数
func (s AppUserService) CheckAuthParams(account string, password string, accountType int32) (code int, msg string) {
	if strings.TrimSpace(account) == "" {
		code = -1
		msg = "账号为空"
		return
	}
	if strings.TrimSpace(password) == "" {
		code = -1
		msg = "密码为空"
		return
	}
	if accountType == 0 {
		code = -1
		msg = "类型为空"
		return
	}
	return
}

// 登录检验手机和邮箱
func (s AppUserService) AuthCheckPhoneAndEmail(account, areaPhoneNumber string, accountType int32) (accountParam string, code int, msg string) {
	if accountType == 1 {
		if iotutil.CheckAllPhone(areaPhoneNumber, account) == false {
			code = -1
			msg = "手机号码不合法"
			return
		}
		accountParam = "phone"
	} else if accountType == 2 {
		if iotutil.VerifyEmailFormat(account) == false {
			code = -1
			msg = "邮箱不合法"
			return
		}
		accountParam = "email"
	} else if accountType == 3 {
		if iotutil.VerifyAccount(account) == false {
			code = -1
			msg = "账号不合法"
			return
		}

	} else {
		code = -1
		msg = "用户类型有误"
		return
	}
	return
}

// 检查注册接口中的参数
func (s AppUserService) CheckRegisterParams(password string, smsCode string, registerRegion, appKey, tenantId string) (code int, msg string) {
	if strings.TrimSpace(password) == "" {
		code = 100009
		msg = "密码不能为空"
		return
	}
	if strings.TrimSpace(smsCode) == "" {
		code = 100010
		msg = "验证码不能为空"
		return
	}
	if strings.TrimSpace(registerRegion) == "" {
		code = 100010
		msg = "注册地区不能为空"
		return
	}
	if strings.TrimSpace(appKey) == "" {
		code = 100010
		msg = "appKey不能为空"
		return
	}
	if strings.TrimSpace(tenantId) == "" {
		code = 100010
		msg = "tenantId不能为空"
		return
	}
	return
}

// 检查注册接口中的参数
func (s AppUserService) CheckRegisterParamsEx(password string, smsCode string, registerRegion, appKey, tenantId string) (code int, msg string) {
	if strings.TrimSpace(smsCode) == "" {
		code = 100010
		msg = "验证码不能为空"
		return
	}
	if strings.TrimSpace(registerRegion) == "" {
		code = 100010
		msg = "注册地区不能为空"
		return
	}
	if strings.TrimSpace(appKey) == "" {
		code = 100010
		msg = "appKey不能为空"
		return
	}
	if strings.TrimSpace(tenantId) == "" {
		code = 100010
		msg = "tenantId不能为空"
		return
	}
	return
}

// 注册检验手机和邮箱
func (s AppUserService) RegisterCheckPhoneAndEmail(phone, email, areaPhoneNumber string) (param string, accountParam string, code int, msg string) {
	if strings.TrimSpace(phone) == "" && strings.TrimSpace(email) == "" {
		code = 101001
		msg = "手机或邮箱不能为空"
		return
	} else if strings.TrimSpace(phone) != "" {
		if iotutil.CheckAllPhone(areaPhoneNumber, phone) == false {
			code = 100006
			msg = "手机格式不合法"
			return
		}
		param = phone
		accountParam = "phone"
	} else if strings.TrimSpace(email) != "" {
		if iotutil.VerifyEmailFormat(email) == false {
			code = 100007
			msg = "邮箱格式不合法"
			return
		}
		param = email
		accountParam = "email"
	} else {
		code = 100018
		msg = "传参有误"
		return
	}
	return
}

func AppUserInfo_pb2eModel(account string, src *protosService.AppUser) *controls.UserInfo {
	if src == nil {
		return nil
	}
	uiObj := controls.UserInfo{
		UserID:         src.Id,
		Nickname:       src.NickName,
		Avatar:         "",
		Account:        src.Account,
		RegionServerId: src.RegionServerId,
		TenantId:       src.TenantId,
		AppKey:         src.AppKey,
	}
	return &uiObj
}

// AppUserLogin App用户登录获取登录信息
func (s AppUserService) AppUserLogin(accountType int32, account string, password string, appKey string, tenantId string, regionServerId int64, registerRegion string, loginType int32, ip string) (*entitys.LoginUserRes, int, string) {
	if accountType == 0 {
		return nil, -1, "账号类型不能为空"
	}
	if account == "" {
		return nil, -1, "账号不能为空"
	}
	if password == "" {
		return nil, -1, "密码不能为空"
	}
	var AppLoginResponse *protosService.AppLoginResponse
	var err error
	if loginType == 0 {
		AppLoginResponse, err = rpc.AppAuthService.PasswordLogin(s.Ctx, &protosService.PasswordLoginRequest{
			LoginName:      account,
			Password:       password,
			Channel:        iotutil.ToString(accountType),
			AppKey:         appKey,
			TenantId:       tenantId,
			RegionServerId: regionServerId,
		})
	} else {
		//1、验证码登录，先校验验证码,此时password为验证码
		verificationCode, _ := iotredis.GetClient().Get(s.Ctx, cached.APP+"_"+appKey+"_"+account+"_7").Result()
		if verificationCode != password {
			return nil, -1, "验证码不对"
		}

		//2、调用用户服务接口，验证用户是否存在，不存在则创建用户
		cond := protosService.UcUserFilter{
			AppKey:         appKey,
			TenantId:       tenantId,
			RegionServerId: regionServerId,
		}
		if accountType == 1 {
			cond.Phone = account
		} else if accountType == 2 {
			cond.Email = account
		}

		resp, err := rpc.TUcUserService.Find(s.Ctx, &cond)
		if err == nil && resp.Total > 0 && len(resp.Data) > 0 {
			//用户存在，啥也不做
		} else {
			//用户不存在，自动注册
			lang, _ := metadata.Get(s.Ctx, "lang")
			langMap, err := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_DFAULT_DATA).Result()
			if err != nil {
				langMap = make(map[string]string)
			}
			defaultHomeName := iotutil.MapGetStringVal(langMap[fmt.Sprintf("%s_default_home_name", lang)], _const.DefaultHomeName)

			//区域服务器Id
			ret, err := rpc.TUcUserService.Register(s.Ctx, &protosService.UcUserRegisterRequest{
				Phone:           cond.Phone,
				Password:        iotutil.EncodeMD5(iotutil.EncodeMD5(iotutil.GetSecret(6))),
				Email:           cond.Email,
				Code:            "",
				RegisterRegion:  registerRegion,
				Ip:              ip,
				AppKey:          appKey,
				TenantId:        tenantId,
				DefaultHomeName: defaultHomeName,
				Lang:            lang,
				RegionServerId:  regionServerId,
			})
			if err != nil {
				iotlogger.LogHelper.Error("AppUserLogin.用户%s注册失败，原因:%s", account, err.Error())
				return nil, -1, fmt.Sprintf("%s", err.Error())
			}
			if ret.Code != 200 {
				iotlogger.LogHelper.Error("AppUserLogin.用户%s注册失败，原因:%s", account, ret.Message)
				return nil, -1, ret.Message
			}
		}
		AppLoginResponse, err = rpc.AppAuthService.PasswordLogin(s.Ctx, &protosService.PasswordLoginRequest{
			LoginName: account,
			//Password:       password,
			Channel:        iotutil.ToString(accountType),
			AppKey:         appKey,
			TenantId:       tenantId,
			RegionServerId: regionServerId,
			VerifyCode:     verificationCode, //传验证码，用于auth服务忽略密码验证
		})
	}

	if err != nil {
		iotlogger.LogHelper.Error("用户%s登录失败，原因:%s", account, err.Error())
		return nil, -1, iotutil.ResErrToString(err)
	}
	userinfo := AppUserInfo_pb2eModel(account, AppLoginResponse.GetUserInfo())
	userinfo.TenantId = tenantId
	userinfo.AppKey = appKey
	iotlogger.LogHelper.Infof("AppUserLogin,userinfo:%s--%s", userinfo.AppKey, userinfo.TenantId)
	expires := time.Unix(AppLoginResponse.ExpiresAt, 0).Sub(time.Now())
	_, err = iotredis.GetClient().Set(s.Ctx, AppLoginResponse.Token, iotutil.ToString(userinfo), expires).Result()
	if err != nil {
		iotlogger.LogHelper.Errorf("AppUserLogin,缓存token失败:%s", err.Error())
		return nil, -1, iotutil.ResErrToString(err)
	}
	controls.CacheTokenByUserId(userinfo.UserID, AppLoginResponse.Token, AppLoginResponse.ExpiresAt)
	return AppLoginInfo_pb2e(AppLoginResponse).SetShowVconsole(getShowVconsole(AppLoginResponse.UserInfo.Id)), 0, ""
}

func getShowVconsole(userId int64) bool {
	//获取用户调试数据
	var showVconsole bool
	debuggers, _ := rpc.ClientOemAppDebuggerService.Find(context.Background(), &protosService.OemAppDebuggerFilter{UserId: userId})
	if debuggers != nil && debuggers.Code == 200 && debuggers.Data != nil && len(debuggers.Data) > 0 {
		showVconsole = true
	}
	return showVconsole
}

func (s AppUserService) Register(params entitys.UserRegister) (code int, userId int64, msg string) {
	lang, _ := metadata.Get(s.Ctx, "lang")
	smsCodeValue := iotredis.GetClient().Get(s.Ctx, cached.APP+"_"+params.AppKey+"_"+params.Account+"_1")

	if smsCodeValue.Val() != params.Smscode {
		iotlogger.LogHelper.Error("验证码有误")
		return -1, 0, "验证码有误"
	}

	var phone, email string
	if params.AccountType == "phone" {
		phone = params.Account
	} else if params.AccountType == "email" {
		email = params.Account
	}

	//读取翻译内容
	langMap, err := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_DFAULT_DATA).Result()
	if err != nil {
		langMap = make(map[string]string)
	}
	defaultHomeName := iotutil.MapGetStringVal(langMap[fmt.Sprintf("%s_default_home_name", lang)], _const.DefaultHomeName)

	//区域服务器Id
	ret, err := rpc.TUcUserService.Register(s.Ctx, &protosService.UcUserRegisterRequest{
		Phone:           phone,
		Password:        iotutil.EncodeMD5(params.Password),
		Email:           email,
		Code:            "",
		RegisterRegion:  params.RegisterRegion,
		Ip:              params.Ip,
		AppKey:          params.AppKey,
		TenantId:        params.TenantId,
		DefaultHomeName: defaultHomeName,
		Lang:            lang,
		RegionServerId:  params.RegisterRegionId,
	})
	if err != nil {
		iotlogger.LogHelper.Error("用户%s注册失败，原因:%s", params.Account, err.Error())
		return -1, 0, err.Error()
	}
	if ret.Code != 200 {
		iotlogger.LogHelper.Error("用户%s注册失败，原因:%s", params.Account, ret.Message)
		return -1, 0, ret.Message
	}
	var newUserId int64
	if len(ret.Data) > 0 {
		newUserId = ret.Data[0].Id
	}
	//删除redis中验证码
	iotredis.GetClient().Del(context.Background(), cached.APP+"_"+params.AppKey+"_"+params.Account+"_1")
	return 0, newUserId, ""
}

// 不设密码注册，并返回token
func (s AppUserService) RegisterEx(params entitys.UserRegister) (*entitys.LoginUserRes, int, string) {
	lang, _ := metadata.Get(s.Ctx, "lang")
	smsCodeValue := iotredis.GetClient().Get(s.Ctx, cached.APP+"_"+params.AppKey+"_"+params.Account+"_1")

	if smsCodeValue.Val() != params.Smscode {
		iotlogger.LogHelper.Error("验证码有误")
		return nil, -1, "验证码有误"
	}

	var phone, email string
	if params.AccountType == "phone" {
		phone = params.Account
	} else if params.AccountType == "email" {
		email = params.Account
	}

	//读取翻译内容
	langMap, err := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_DFAULT_DATA).Result()
	if err != nil {
		langMap = make(map[string]string)
	}
	defaultHomeName := iotutil.MapGetStringVal(langMap[fmt.Sprintf("%s_default_home_name", lang)], _const.DefaultHomeName)
	var password string
	if params.Password == "" {
		password = iotutil.EncodeMD5(iotutil.EncodeMD5(iotutil.GetSecret(6)))
	} else {
		password = iotutil.EncodeMD5(params.Password)
	}

	//区域服务器Id
	ret, err := rpc.TUcUserService.Register(s.Ctx, &protosService.UcUserRegisterRequest{
		Phone:           phone,
		Password:        password,
		Email:           email,
		Code:            "",
		RegisterRegion:  params.RegisterRegion,
		Ip:              params.Ip,
		AppKey:          params.AppKey,
		TenantId:        params.TenantId,
		DefaultHomeName: defaultHomeName,
		Lang:            lang,
		RegionServerId:  params.RegisterRegionId,
	})
	if err != nil {
		iotlogger.LogHelper.Error("用户%s注册失败，原因:%s", params.Account, err.Error())
		return nil, -1, err.Error()
	}
	if ret.Code != 200 {
		iotlogger.LogHelper.Error("用户%s注册失败，原因:%s", params.Account, ret.Message)
		return nil, -1, ret.Message
	}
	//var newUserId int64
	//if len(ret.Data) > 0 {
	//	newUserId = ret.Data[0].Id
	//}
	//删除redis中验证码
	iotredis.GetClient().Del(context.Background(), cached.APP+"_"+params.AppKey+"_"+params.Account+"_1")

	//注册成功后，走自动登录逻辑
	var accountType int
	if params.AccountType == "phone" {
		accountType = 1
	} else if params.AccountType == "email" {
		accountType = 2
	}
	AppLoginResponse, err := rpc.AppAuthService.PasswordLogin(s.Ctx, &protosService.PasswordLoginRequest{
		LoginName: params.Account,
		//Password:       password,
		Channel:        iotutil.ToString(accountType),
		AppKey:         params.AppKey,
		TenantId:       params.TenantId,
		RegionServerId: params.RegisterRegionId,
		VerifyCode:     params.Smscode,
	})
	if err != nil {
		iotlogger.LogHelper.Error("用户%s登录失败，原因:%s", params.Account, err.Error())
		return nil, -1, iotutil.ResErrToString(err)
	}
	userinfo := AppUserInfo_pb2eModel(params.Account, AppLoginResponse.GetUserInfo())
	userinfo.TenantId = params.TenantId
	userinfo.AppKey = params.AppKey
	iotlogger.LogHelper.Infof("RegisterEx,userinfo:%s--%s", userinfo.AppKey, userinfo.TenantId)
	expires := time.Unix(AppLoginResponse.ExpiresAt, 0).Sub(time.Now())
	_, err = iotredis.GetClient().Set(s.Ctx, AppLoginResponse.Token, iotutil.ToString(userinfo), expires).Result()
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterEx,缓存token失败:%s", err.Error())
		return nil, -1, iotutil.ResErrToString(err)
	}
	controls.CacheTokenByUserId(userinfo.UserID, AppLoginResponse.Token, AppLoginResponse.ExpiresAt)
	return AppLoginInfo_pb2e(AppLoginResponse).SetShowVconsole(getShowVconsole(AppLoginResponse.UserInfo.Id)), 0, ""
}

func (s AppUserService) UpdateUser(userId string, updateUserParam entitys.UpdateUserParam) (code int, msg string) {
	req := &protosService.UcUser{
		Id:            iotutil.ToInt64(userId),
		NickName:      updateUserParam.Nickname,
		Photo:         updateUserParam.Photo,
		Gender:        updateUserParam.Gender,
		Password:      updateUserParam.NewPassword,
		DefaultHomeId: updateUserParam.DefaultHomeId,
		Country:       updateUserParam.Country,
		Province:      updateUserParam.Province,
		City:          updateUserParam.City,
	}
	if updateUserParam.Birthday != "" {
		birthday, err := iotutil.GetDateByStr(updateUserParam.Birthday)
		if err != nil {
			return -1, err.Error()
		}
		req.Birthday = timestamppb.New(birthday)
	}
	ret, err := rpc.TUcUserService.Update(context.Background(), req)
	if err != nil {
		iotlogger.LogHelper.Error("修改用户信息失败，原因:%s", err.Error())
		return -1, err.Error()
	}
	if ret.Code != 200 {
		iotlogger.LogHelper.Error("修改用户信息失败，原因:%s", ret.Message)
		return -1, ret.Message
	}
	du := iotstruct.DeviceRedisUpdate{
		UserId: userId,
		HomeId: updateUserParam.DefaultHomeId,
	}
	duBytes, err := json.Marshal(du)
	if err != nil {
		return -1, err.Error()
	}
	if len(updateUserParam.DefaultHomeId) != 0 {
		err = iotredis.GetClient().Publish(context.Background(), strings.Join([]string{iotconst.HKEY_UPDATE_DATA_PUB_PREFIX, userId}, "."), string(duBytes)).Err()
		if err != nil {
			return -1, err.Error()
		}
	}
	return 0, ""
}

func (s AppUserService) UpdateUserPassword(userId string, updateUserParam entitys.SetPassword) (code int, msg string) {
	req := &protosService.UcUser{
		Id:       iotutil.ToInt64(userId),
		Password: updateUserParam.NewPassword,
	}
	ret, err := rpc.TUcUserService.Update(context.Background(), req)
	if err != nil {
		iotlogger.LogHelper.Error("修改用户信息失败，原因:%s", err.Error())
		return -1, err.Error()
	}
	if ret.Code != 200 {
		iotlogger.LogHelper.Error("修改用户信息失败，原因:%s", ret.Message)
		return -1, ret.Message
	}
	return 0, ""
}

func (s AppUserService) SendEmail(email string, emailType int32, appKey, lang string) (emailCode string, code int, msg string) {
	emailCode = iotutil.Getcode()
	iotlogger.LogHelper.Info(emailCode)
	res := iotredis.GetClient().Set(context.Background(), cached.APP+"_"+appKey+"_"+email+"_"+iotutil.ToString(emailType), emailCode, 600*time.Second) //有效期10分钟
	if res.Err() != nil {
		iotlogger.LogHelper.Errorf("SendEmail,缓存emailCode失败:%s", res.Err().Error())
		return "", -1, res.Err().Error()
	}
	_, err := rpc.EmailService.SendEmailUserCode(context.Background(), &protosService.SendEmailUserCodeRequest{
		Email:   email,
		Code:    emailCode,
		Lang:    lang,
		TplType: emailType,
	})
	if err != nil {
		iotlogger.LogHelper.Error("发送邮件失败，原因:%s", err.Error())
		return "", -1, err.Error()
	}
	return emailCode, 0, ""
}

// 发送短信验证码
func (s AppUserService) SendSms(lang, areaPhone, phone string, phoneType, smsType int32, appKey string) (smsCode string, code int, msg string) {
	smsCode = iotutil.Getcode()
	res := iotredis.GetClient().Set(context.Background(), cached.APP+"_"+appKey+"_"+phone+"_"+iotutil.ToString(smsType), smsCode, 600*time.Second) //有效期10分钟
	if res.Err() != nil {
		iotlogger.LogHelper.Errorf("SendSms,缓存smsCodeInt失败:%s", res.Err().Error())
		return "", -1, res.Err().Error()
	}
	_, err := rpc.SmsService.SendSMSVerifyCode(context.Background(), &protosService.SendSMSVerifyCodeRequest{
		PhoneNumber: areaPhone,
		UserName:    areaPhone,
		Code:        smsCode,
		Lang:        lang,
		TplType:     smsType,
		PhoneType:   phoneType,
	})
	if err != nil {
		iotlogger.LogHelper.Error("发送短信验证码失败，原因:%s", err.Error())
		return "", -1, err.Error()
	}
	return smsCode, 0, ""
}

func (s AppUserService) CheckCode(checkCodeParam entitys.CheckCodeParam, appKey string) (code int, msg string) {
	codeType := checkCodeParam.Type
	smsCode := checkCodeParam.Code
	account := checkCodeParam.Account
	accountType := checkCodeParam.AccountType

	if strings.TrimSpace(account) == "" {
		iotlogger.LogHelper.Errorf("用户信息为空")
		return -1, "用户信息为空"
	}
	if accountType == 0 {
		iotlogger.LogHelper.Errorf("用户类型为空")
		return -1, "用户类型为空"
	}
	if codeType == 0 {
		iotlogger.LogHelper.Errorf("类型为空")
		return -1, "类型为空"
	}
	if strings.TrimSpace(smsCode) == "" {
		iotlogger.LogHelper.Errorf("验证码为空")
		return -1, "验证码为空"
	}

	if accountType == 1 {
		if iotutil.CheckAllPhone(checkCodeParam.AreaPhoneNumber, account) == false {
			iotlogger.LogHelper.Errorf("手机号码不合法")
			return -1, "手机号码不合法"
		}
	} else if accountType == 2 {
		if iotutil.VerifyEmailFormat(account) == false {
			iotlogger.LogHelper.Errorf("电子邮箱不合法")
			return -1, "电子邮箱不合法"
		}
	} else {
		iotlogger.LogHelper.Errorf("用户类型有误")
		return -1, "用户类型有误"
	}

	resp := iotredis.GetClient().Get(context.Background(), cached.APP+"_"+appKey+"_"+account+"_"+iotutil.ToString(codeType))
	if resp.Val() != smsCode {
		iotlogger.LogHelper.Error("验证码有误")
		return -1, "验证码有误"
	}
	return 0, ""
}

func (s AppUserService) ForgetPassword(updateUserParam entitys.ForgetPassword, appKey, tenantId string, regionServerId int64) (response *entitys.LoginUserRes, code int, msg string) {
	smsCode := updateUserParam.Code
	if strings.TrimSpace(smsCode) == "" {
		return nil, -1, "验证码为空"
	}

	newPassword := updateUserParam.NewPassword
	if strings.TrimSpace(newPassword) == "" {
		return nil, -1, "密码为空"
	}

	account := updateUserParam.Account
	if strings.TrimSpace(account) == "" {
		return nil, -1, "用户信息为空"
	}

	accountType := updateUserParam.Type
	if accountType == 0 {
		return nil, -1, "用户类型为空"
	}

	if accountType == 1 {
		if iotutil.CheckAllPhone(updateUserParam.AreaPhoneNumber, account) == false {
			return nil, -1, "手机号码不合法"
		}
	} else if accountType == 2 {
		if iotutil.VerifyEmailFormat(account) == false {
			return nil, -1, "电子邮箱不合法"
		}
	} else {
		return nil, -1, "用户类型有误"
	}

	resp := iotredis.GetClient().Get(context.Background(), cached.APP+"_"+appKey+"_"+account+"_2")

	if resp.Val() != smsCode {
		iotlogger.LogHelper.Error("验证码有误")
		return nil, -1, "验证码有误"
	}

	ret, err := rpc.TUcUserService.ForgetPassword(context.Background(), &protosService.UcForgetPasswordReq{
		Account:        account,
		AccountType:    iotutil.ToString(accountType),
		NewPassword:    newPassword,
		AppKey:         appKey,
		TenantId:       tenantId,
		RegionServerId: regionServerId,
	})
	if err != nil {
		iotlogger.LogHelper.Error("修改密码失败，原因:%s", err.Error())
		return nil, -1, err.Error()
	}
	//删除redis中验证码
	iotredis.GetClient().Del(context.Background(), cached.APP+"_"+appKey+"_"+account+"_2")
	if len(ret.Data) > 0 {
		go SendUpdatePasswordMessage(SetAppInfoByContext(s.Ctx), ret.Data[0].Id, account)
	}

	//清除token
	controls.ClearTokenByUserId(ret.Data[0].Id)

	//重置密码后返回token
	AppLoginResponse, err := rpc.AppAuthService.GetTokenByAccount(s.Ctx, &protosService.GetTokenByAccountRequest{
		LoginName:      account,
		Password:       newPassword,
		TenantId:       tenantId,
		AppKey:         appKey,
		RegionServerId: regionServerId,
	})
	if err != nil {
		iotlogger.LogHelper.Error("用户%s登录失败，原因:%s", account, err.Error())
		return nil, -1, iotutil.ResErrToString(err)
	}
	userinfo := AppUserInfo_pb2eModel(account, AppLoginResponse.GetUserInfo())
	userinfo.TenantId = tenantId
	userinfo.AppKey = appKey
	expires := time.Unix(AppLoginResponse.ExpiresAt, 0).Sub(time.Now())
	cacheResp := iotredis.GetClient().Set(s.Ctx, AppLoginResponse.Token, iotutil.ToString(userinfo), expires)
	if cacheResp.Err() != nil {
		iotlogger.LogHelper.Errorf("AppUserLogin,缓存token失败:%s", cacheResp.Err().Error())
		return nil, -1, iotutil.ResErrToString(cacheResp.Err())
	}
	return AppLoginInfo_pb2e(AppLoginResponse).SetShowVconsole(getShowVconsole(AppLoginResponse.UserInfo.Id)), 0, ""
}

func (s AppUserService) HomeList(userId int64) (data []map[string]interface{}, code int, msg string) {
	var (
		lang, _ = metadata.Get(s.Ctx, "lang")
	)
	respData := make([]map[string]interface{}, 0)
	err := cached.RedisStore.Get(persist.GetRedisKey(iotconst.APP_HOME_LIST_DATA, iotutil.ToString(userId)), &respData)
	if err == nil {
		for i, d := range respData {
			respData[i]["name"] = HomeLanguage(lang, d["name"])
		}
		return respData, 0, ""
	}
	homeListInfo, err := rpc.TUcUserService.HomeList(context.Background(), &protosService.UcUser{
		Id: userId,
	})
	if err != nil {
		iotlogger.LogHelper.Error("获取用户家庭列表失败，原因:%s", err.Error())
		return nil, -1, err.Error()
	}
	homeList := homeListInfo.GetHomeUsers()

	for _, v := range homeList {
		homeInfo := make(map[string]interface{})
		homeInfo["type"] = v.Type
		homeInfo["id"] = iotutil.ToString(v.Id)
		homeInfo["name"] = HomeLanguage(lang, v.Name) // v.Name
		homeInfo["userNum"] = v.UserList
		homeInfo["roomNum"] = v.RoomList
		homeInfo["deviceNum"] = v.DeviceList
		respData = append(respData, homeInfo)
	}

	err = cached.RedisStore.Set(persist.GetRedisKey(iotconst.APP_HOME_LIST_DATA, iotutil.ToString(userId)), &respData, 600*time.Second)
	if err != nil {
		return respData, 0, err.Error()
	}
	return respData, 0, ""
}

func (s AppUserService) CheckAccount(checkAccountBm entitys.CheckAccount, appKey, tenantId string, regionServerId int64) (code int, msg string) {
	account := checkAccountBm.Account
	accountType := checkAccountBm.Type

	var phone, email string
	switch accountType {
	case 1:
		if iotutil.CheckAllPhone(checkAccountBm.AreaPhoneNumber, account) == false {
			return -1, "手机号码不合法"
		}
		phone = account
	case 2:
		if iotutil.VerifyEmailFormat(account) == false {
			return -1, "电子邮箱不合法"
		}
		email = account
	default: //default case
		return -1, "用户类型有误"
	}

	data, err := rpc.TUcUserService.Find(context.Background(), &protosService.UcUserFilter{
		Phone:          phone,
		Email:          email,
		AppKey:         appKey,
		TenantId:       tenantId,
		RegionServerId: regionServerId,
	})
	if err == nil && data.Total != 0 {
		iotlogger.LogHelper.Error("当前用户%s已注册过", account)
		return 2, ioterrs.ERROR_EXIST_ACOUNT.Msg
	}
	return 0, ""
}

// AppLoginInfo_pb2e 登录数据实体转换
func AppLoginInfo_pb2e(src *protosService.AppLoginResponse) *entitys.LoginUserRes {
	pbObj := &entitys.LoginUserRes{
		UserId:           iotutil.ToString(src.UserInfo.Id),
		NickName:         src.UserInfo.NickName,
		Phone:            src.UserInfo.Phone,
		Photo:            src.UserInfo.Photo,
		Email:            src.UserInfo.Email,
		DefaultHomeId:    src.UserInfo.DefaultHomeId,
		UserName:         src.UserInfo.UserName,
		SubmitCancelTime: src.UserInfo.SubmitCancelTime,
		AccountCasser:    src.UserInfo.AccountCasser,
		Token:            src.Token,
		RefreshToken:     src.RefreshToken,
		ExpiresAt:        src.ExpiresAt,
		Password:         src.UserInfo.Password,
	}
	thirdPartyLogin := entitys.AppUserThirdPartyLogin{}
	for _, partyLogin := range src.UserInfo.ThirdPartyLogin {
		if partyLogin.Mode == 1 {
			thirdPartyLogin.Wechat.Mode = partyLogin.Mode
			thirdPartyLogin.Wechat.LoginKey = partyLogin.LoginKey
			thirdPartyLogin.Wechat.Nickname = partyLogin.Nickname
		} else {
			thirdPartyLogin.AppleId.Mode = partyLogin.Mode
			thirdPartyLogin.AppleId.LoginKey = partyLogin.LoginKey
			thirdPartyLogin.AppleId.Nickname = partyLogin.Nickname
		}
	}
	pbObj.ThirdPartyLogin = thirdPartyLogin
	return pbObj
}

// AppUserInfo_pb2e 用户数据实体转换
func AppUserInfo_pb2e(src *protosService.UcUser) *entitys.AppUserInfo {
	thirdPartyLogin := entitys.AppUserThirdPartyLogin{}
	pbObj := entitys.AppUserInfo{
		UserId:          iotutil.ToString(src.Id),
		NickName:        src.NickName,
		Phone:           src.Phone,
		Photo:           src.Photo,
		Email:           src.Email,
		DefaultHomeId:   src.DefaultHomeId,
		UserName:        src.Phone,
		ThirdPartyLogin: thirdPartyLogin,
	}
	return &pbObj
}

func (s AppUserService) ChannelBind(ip string, req entitys.ChannelBind) (*entitys.LoginUserRes, int, string) {
	//lang, _ := metadata.Get(s.Ctx, "lang")
	appKey := req.AppKey
	tenantId := req.TenantId
	regionServerId := req.RegisterRegionId
	response := iotredis.GetClient().Get(context.Background(), cached.APP+"_"+appKey+"_"+req.Account+"_5")

	if response.Val() != req.Code {
		iotlogger.LogHelper.Errorf("验证码有误")
		return nil, -1, "验证码有误"
	}

	phone, email, err := s.checkUserName(req.Account)
	if err != nil {
		return nil, -1, err.Error()
	}

	reqParam := protosService.UcUserByLoginRequest{
		Phone:          phone,
		Email:          email,
		AppKey:         appKey,
		TenantId:       tenantId,
		RegionServerId: regionServerId,
		Status:         1, //1:正常，2:待注销，3:已注销
	}
	//通过手机号码/邮箱查询用户信息，因为存在注销账号所以有可能会返回多条
	resp, err := s.queryUserByLogin(req.Account, reqParam)
	if err != nil {
		return nil, -1, err.Error()
	}

	if len(resp) > 0 { //已经注册的用户信息的直接登录
		userId := resp[0].Id
		//todo 检查当前第三方userid是否被其他用户账号绑定
		ucUserThirdResp, err := rpc.ClientUcUserThirdService.Find(context.Background(), &protosService.UcUserThirdFilter{
			UserId:         resp[0].Id,
			ThirdType:      req.Type,
			AppKey:         appKey,
			TenantId:       tenantId,
			RegionServerId: regionServerId,
		})
		if err == nil && ucUserThirdResp.Data != nil && len(ucUserThirdResp.Data) > 0 {
			return nil, -1, "当前账号已被第三方用户账号绑定"
		}

		if len(req.Nickname) == 0 {
			AppleidResponse, err := rpc.UcAppleidInfoService.Find(s.Ctx, &protosService.UcAppleidInfoFilter{
				ThirdUserId: req.ChannelId,
			})
			if err != nil {
				return nil, -1, err.Error()
			}
			if AppleidResponse.Data != nil && len(AppleidResponse.Data) > 0 {
				req.Nickname = AppleidResponse.Data[0].Nickname
			}
		}

		_, err = rpc.ClientUcUserThirdService.Create(context.Background(), &protosService.UcUserThird{
			Id:             iotutil.GetNextSeqInt64(),
			UserId:         userId,
			ThirdType:      req.Type,
			ThirdUserId:    req.ChannelId,
			Nickname:       req.Nickname,
			CreatedBy:      userId,
			UpdatedBy:      userId,
			AppKey:         appKey,
			TenantId:       tenantId,
			RegionServerId: regionServerId,
		})
		if err != nil {
			return nil, -1, err.Error()
		}

		userInfo := resp[0]
		//获取token
		AppLoginResponse, err := rpc.AppAuthService.GetToken(context.Background(), &protosService.GetTokenRequest{
			UserId:         userInfo.Id,
			Nickname:       userInfo.NickName,
			Avatar:         "",
			Account:        userInfo.UserName,
			RegionServerId: userInfo.RegionServerId,
			AppKey:         userInfo.AppKey,
			TenantId:       userInfo.TenantId,
		})
		if err != nil {
			iotlogger.LogHelper.Error("获取token失败，原因:%s", req.Account, err.Error())
			return nil, -1, iotutil.ResErrToString(err)
		}
		AppLoginResponse.UserInfo = &protosService.AppUser{
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
			Account:          req.Account,
			Token:            AppLoginResponse.Token,
			AccountCasser:    false,
			SubmitCancelTime: userInfo.CancelTime,
			ThirdPartyLogin:  AppLoginResponse.UserInfo.ThirdPartyLogin,
			UserName:         userInfo.UserName,
			Password:         userInfo.Password,
			RegionServerId:   regionServerId,
		}

		//if err != nil {
		//	iotlogger.LogHelper.Error("用户%s登录失败，原因:%s", "", err.Error())
		//	return &entitys.LoginUserRes{}, -1, iotutil.ResErrToString(err)
		//}
		userinfo := AppUserInfo_pb2eModel("", AppLoginResponse.GetUserInfo())
		userinfo.TenantId = tenantId
		userinfo.AppKey = appKey
		expires := time.Unix(AppLoginResponse.ExpiresAt, 0).Sub(time.Now())
		cacheResp := iotredis.GetClient().Set(context.Background(), AppLoginResponse.Token, iotutil.ToString(userinfo), expires)
		if cacheResp.Err() != nil {
			iotlogger.LogHelper.Errorf("AppUserLogin,缓存token失败:%s", cacheResp.Err().Error())
			return &entitys.LoginUserRes{}, -1, iotutil.ResErrToString(cacheResp.Err())
		}
		//删除redis中验证码
		iotredis.GetClient().Del(context.Background(), cached.APP+"_"+appKey+"_"+req.Account+"_5")
		return AppLoginInfo_pb2e(AppLoginResponse).SetShowVconsole(getShowVconsole(AppLoginResponse.UserInfo.Id)), 0, ""
	} else { //未注册用户信息的先进行注册
		lang, _ := metadata.Get(s.Ctx, "lang")
		//读取翻译内容
		langMap, err := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_DFAULT_DATA).Result()
		if err != nil {
			langMap = make(map[string]string)
		}
		defaultHomeName := iotutil.MapGetStringVal(langMap[fmt.Sprintf("%s_default_home_name", lang)], _const.DefaultHomeName)
		ret, err := rpc.TUcUserService.Register(s.Ctx, &protosService.UcUserRegisterRequest{
			Phone:           phone,
			Password:        iotutil.EncodeMD5(req.Password),
			Email:           email,
			Code:            "",
			RegisterRegion:  req.RegisterRegion,
			Ip:              ip,
			AppKey:          appKey,
			TenantId:        tenantId,
			DefaultHomeName: defaultHomeName,
			Lang:            lang,
			ThirdType:       iotutil.ToString(req.Type),
			ThirdUserId:     req.ChannelId,
			ThirdNickname:   req.Nickname,
			RegionServerId:  req.RegisterRegionId,
		})
		if err != nil {
			iotlogger.LogHelper.Error("用户%s注册失败，原因:%s", req.Account, err.Error())
			return &entitys.LoginUserRes{}, -1, err.Error()
		}
		if ret.Code != 200 {
			iotlogger.LogHelper.Error("用户%s注册失败，原因:%s", req.Account, ret.Message)
			return &entitys.LoginUserRes{}, -1, ret.Message
		}
	}

	responseInfo, err := s.queryUserByLogin(req.Account, reqParam)
	if err != nil {
		return nil, -1, err.Error()
	}

	userData := responseInfo[0]
	//获取登录的token
	AppLoginResponse, err := rpc.AppAuthService.GetToken(context.Background(), &protosService.GetTokenRequest{
		UserId:         userData.Id,
		Nickname:       userData.NickName,
		Avatar:         "",
		Account:        userData.UserName,
		TenantId:       tenantId,
		AppKey:         appKey,
		RegionServerId: regionServerId,
	})
	if err != nil {
		iotlogger.LogHelper.Error("获取token失败，原因:%s", req.Account, err.Error())
		return nil, -1, iotutil.ResErrToString(err)
	}
	//AppUser
	AppLoginResponse.UserInfo = &protosService.AppUser{
		Id:               userData.Id,
		NickName:         userData.NickName,
		Phone:            userData.Phone,
		Photo:            userData.Photo,
		Status:           userData.Status,
		City:             userData.City,
		Gender:           userData.Gender,
		Email:            userData.Email,
		DefaultHomeId:    userData.DefaultHomeId,
		RegisterRegion:   userData.RegisterRegion,
		Account:          req.Account,
		Token:            AppLoginResponse.Token,
		AccountCasser:    false,
		SubmitCancelTime: userData.CancelTime,
		ThirdPartyLogin:  AppLoginResponse.UserInfo.ThirdPartyLogin,
		UserName:         userData.UserName,
		Password:         userData.Password,
	}

	userinfo := AppUserInfo_pb2eModel("", AppLoginResponse.GetUserInfo())
	userinfo.TenantId = tenantId
	userinfo.AppKey = appKey
	expires := time.Unix(AppLoginResponse.ExpiresAt, 0).Sub(time.Now())
	cacheResp := iotredis.GetClient().Set(context.Background(), AppLoginResponse.Token, iotutil.ToString(userinfo), expires)
	if cacheResp.Err() != nil {
		iotlogger.LogHelper.Errorf("AppUserLogin,缓存token失败:%s", cacheResp.Err().Error())
		return &entitys.LoginUserRes{}, -1, iotutil.ResErrToString(cacheResp.Err())
	}
	//删除redis中验证码
	iotredis.GetClient().Del(context.Background(), cached.APP+"_"+appKey+"_"+req.Account+"_5")
	return AppLoginInfo_pb2e(AppLoginResponse).SetShowVconsole(getShowVconsole(AppLoginResponse.UserInfo.Id)), 0, ""
}

func (s *AppUserService) RegisterAndLogin(req entitys.RegisterNewUser) (*entitys.LoginUserRes, int, string) {
	lang, _ := metadata.Get(s.Ctx, "lang")
	//读取翻译内容
	langMap, err := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_DFAULT_DATA).Result()
	if err != nil {
		langMap = make(map[string]string)
	}
	defaultHomeName := iotutil.MapGetStringVal(langMap[fmt.Sprintf("%s_default_home_name", lang)], _const.DefaultHomeName)
	ret, err := rpc.TUcUserService.Register(s.Ctx, &protosService.UcUserRegisterRequest{
		Phone:           req.Phone,
		Password:        iotutil.EncodeMD5(req.Password),
		Email:           req.Email,
		Code:            "",
		RegionServerId:  req.RegionServerId,
		Ip:              req.Ip,
		AppKey:          req.AppKey,
		TenantId:        req.TenantId,
		DefaultHomeName: defaultHomeName,
		Lang:            lang,
		ThirdType:       iotutil.ToString(req.ThirdType),
		ThirdUserId:     req.ThirdUserId,
		ThirdNickname:   req.ThirdNickname,
	})
	if err != nil {
		iotlogger.LogHelper.Error("第三方登录类型%v，用户%s注册失败，原因:%s", req.ThirdType, req.ThirdUserId, err.Error())
		return &entitys.LoginUserRes{}, -1, err.Error()
	}
	if ret.Code != 200 {
		iotlogger.LogHelper.Error("第三方登录类型%v，用户%s注册失败，原因:%s", req.ThirdType, req.ThirdUserId, ret.Message)
		return &entitys.LoginUserRes{}, -1, ret.Message
	}

	userData := ret.Data[0]
	//获取登录的token
	AppLoginResponse, err := rpc.AppAuthService.GetToken(context.Background(), &protosService.GetTokenRequest{
		UserId:         ret.Data[0].Id,
		Nickname:       req.ThirdNickname,
		Avatar:         "",
		Account:        ret.Data[0].UserName,
		TenantId:       req.TenantId,
		AppKey:         req.AppKey,
		RegionServerId: req.RegionServerId,
	})
	if err != nil {
		iotlogger.LogHelper.Error("获取token失败%v，用户%s注册失败，原因:%s", req.ThirdType, req.ThirdUserId, err.Error())
		return nil, -1, iotutil.ResErrToString(err)
	}
	//AppUser
	AppLoginResponse.UserInfo = &protosService.AppUser{
		Id:               userData.Id,
		NickName:         userData.NickName,
		Phone:            userData.Phone,
		Photo:            userData.Photo,
		Status:           userData.Status,
		City:             userData.City,
		Gender:           userData.Gender,
		Email:            userData.Email,
		DefaultHomeId:    userData.DefaultHomeId,
		RegisterRegion:   userData.RegisterRegion,
		Account:          ret.Data[0].UserName,
		Token:            AppLoginResponse.Token,
		AccountCasser:    false,
		SubmitCancelTime: userData.CancelTime,
		ThirdPartyLogin:  AppLoginResponse.UserInfo.ThirdPartyLogin,
		UserName:         userData.UserName,
		Password:         userData.Password,
	}

	userinfo := AppUserInfo_pb2eModel("", AppLoginResponse.GetUserInfo())
	userinfo.TenantId = req.TenantId
	userinfo.AppKey = req.AppKey
	expires := time.Unix(AppLoginResponse.ExpiresAt, 0).Sub(time.Now())
	cacheResp := iotredis.GetClient().Set(context.Background(), AppLoginResponse.Token, iotutil.ToString(userinfo), expires)
	if cacheResp.Err() != nil {
		iotlogger.LogHelper.Errorf("AppUserLogin,缓存token失败:%s", cacheResp.Err().Error())
		return &entitys.LoginUserRes{}, -1, iotutil.ResErrToString(cacheResp.Err())
	}
	//删除redis中验证码
	iotredis.GetClient().Del(context.Background(), cached.APP+"_"+req.AppKey+"_"+ret.Data[0].UserName+"_5")
	return AppLoginInfo_pb2e(AppLoginResponse).SetShowVconsole(getShowVconsole(AppLoginResponse.UserInfo.Id)), 0, ""
}

// checkUserName 检查用户名称，判断为手机、邮箱
func (s *AppUserService) checkUserName(userName string) (phone, email string, err error) {
	if iotutil.CheckAllPhone("", userName) {
		phone = userName
	} else if iotutil.IsEmail(userName) {
		email = userName
	} else {
		err = errors.New("无法识别账号格式")
	}
	return
}

// queryUserByLogin 通过登录信息查询用户信息
func (s *AppUserService) queryUserByLogin(loginName string, req protosService.UcUserByLoginRequest) (res []*protosService.UcUser, err error) {
	//通过手机号码/邮箱查询用户信息，因为存在注销账号所以有可能会返回多条
	resp, err := rpc.TUcUserService.GetUserByLogin(s.Ctx, &req)
	if err != nil {
		iotlogger.LogHelper.Errorf("PasswordLogin: 用户%s登录失败，原因:%s", loginName, err.Error())
		return nil, err
	}
	if resp.Code != 200 {
		//return nil, errors.New(resp.Message)
		return []*protosService.UcUser{}, nil
	}
	if len(resp.Data) == 0 {
		//err = errors.New("用户不存在")
		//iotlogger.LogHelper.Errorf("PasswordLogin: 用户%s登录失败，原因:%s", loginName, err.Error())
		//return nil, err
		return []*protosService.UcUser{}, nil
	}
	return resp.Data, nil
}

// 检查账号
func (s *AppUserService) verityAccount(userlist []*protosService.UcUser) (userinfo *protosService.UcUser, err error) {
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

// TODO 迁移到auth微服务中
// 微信登录
func (s AppUserService) WechatLogin(authorizationCode, appKey, tenantId, ip string, regionServerId int64) (*entitys.LoginUserRes, int, string) {
	////根据渠道类型、code查出渠道信息(第三方userid和昵称)
	//channelUserId, channelIdName, channelNickname, channelName, code, msg := GetWechatInfo(authorizationCode)
	//if code != 0 {
	//	return
	//}
	//obj, code, msg = GetChannelAuthData(channelUserId, channelIdName, channelNickname, channelName, _const.Wechat)

	oemAppResult, err := rpc.ClientOemAppService.Find(context.Background(), &protosService.OemAppFilter{
		AppKey: appKey,
	})
	if err != nil {
		iotlogger.LogHelper.Errorf("GetFunctionConfig error")
		return &entitys.LoginUserRes{}, -1, iotutil.ResErrToString(err)
	}
	if oemAppResult.Code != 200 {
		iotlogger.LogHelper.Errorf(oemAppResult.Message)
		return &entitys.LoginUserRes{}, -1, oemAppResult.Message
	}
	oemAppInfo := oemAppResult.Data[0]
	oemAppFunctionConfig, err := rpc.ClientOemAppFunctionConfigService.Find(context.Background(), &protosService.OemAppFunctionConfigFilter{
		AppId:   oemAppInfo.Id,
		Version: oemAppInfo.Version,
	})
	if err != nil {
		iotlogger.LogHelper.Errorf("GetFunctionConfig error")
		return &entitys.LoginUserRes{}, -1, iotutil.ResErrToString(err)
	}
	if oemAppFunctionConfig.Code != 200 {
		iotlogger.LogHelper.Errorf(oemAppFunctionConfig.Message)
		return &entitys.LoginUserRes{}, -1, oemAppFunctionConfig.Message
	}

	if oemAppFunctionConfig.Data[0].Thirds == "" {
		return &entitys.LoginUserRes{}, -1, ""
	}
	thirds := iotutil.JsonToMapArray(oemAppFunctionConfig.Data[0].Thirds)

	var appId, appSecret string
	for _, third := range thirds {
		if third["code"] == "wechat" {
			appId = iotutil.ToString(third["appId"])
			appSecret = iotutil.ToString(third["appKey"])
			break
		}
	}

	AppLoginResponse, err := rpc.AppAuthService.WechatLogin(context.Background(), &protosService.AppThirdRequest{
		Code:           authorizationCode,
		ChannelId:      "",
		Mode:           "1",
		Nickname:       "",
		AppId:          appId,
		AppSecret:      appSecret,
		AppKey:         appKey,
		TenantId:       tenantId,
		RegionServerId: regionServerId,
		Ip:             ip,
	})

	if AppLoginResponse == nil && err != nil {
		iotlogger.LogHelper.Error("用户%s登录失败，原因:%s", "", err.Error())
		return &entitys.LoginUserRes{}, -1, iotutil.ResErrToString(err)
	}
	//if AppLoginResponse == nil {
	//	iotlogger.LogHelper.Error("用户%s登录失败，原因:%s", "", err.Error())
	//	return nil, -1, iotutil.ResErrToString(err)
	//}
	if AppLoginResponse.Token == "" {
		loginResponse := AppLoginInfo_pb2e(AppLoginResponse).SetShowVconsole(getShowVconsole(AppLoginResponse.UserInfo.Id))
		loginResponse.LoginKey = loginResponse.ThirdPartyLogin.Wechat.LoginKey
		loginResponse.ChannelName = loginResponse.ThirdPartyLogin.Wechat.Nickname
		loginResponse.UserId = ""
		return loginResponse, 0, "没有绑定用户请前往绑定"
	}

	userinfo := AppUserInfo_pb2eModel("", AppLoginResponse.GetUserInfo())
	userinfo.TenantId = tenantId
	userinfo.AppKey = appKey
	expires := time.Unix(AppLoginResponse.ExpiresAt, 0).Sub(time.Now())
	cacheResp := iotredis.GetClient().Set(context.Background(), AppLoginResponse.Token, iotutil.ToString(userinfo), expires)
	if cacheResp.Err() != nil {
		iotlogger.LogHelper.Errorf("AppUserLogin,缓存token失败:%s", cacheResp.Err().Error())
		return &entitys.LoginUserRes{}, -1, iotutil.ResErrToString(cacheResp.Err())
	}
	return AppLoginInfo_pb2e(AppLoginResponse).SetShowVconsole(getShowVconsole(AppLoginResponse.UserInfo.Id)), 0, "ok"
}

// Appleid登录 （GOOGlE、APPLEID 通用接口）
func (s AppUserService) AppleidLogin(channelId string, nickname, ip, appKey, tenantId string, channelType int32, regionServerId int64) (*entitys.LoginUserRes, int, string) {
	lang, _ := metadata.Get(s.Ctx, "lang")
	//读取翻译内容
	langMap, err := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_DFAULT_DATA).Result()
	if err != nil {
		langMap = make(map[string]string)
	}
	defaultHomeName := iotutil.MapGetStringVal(langMap[fmt.Sprintf("%s_default_home_name", lang)], _const.DefaultHomeName)

	AppLoginResponse, err := rpc.AppAuthService.AppleidLogin(s.Ctx, &protosService.AppThirdRequest{
		Code:           "",
		ChannelId:      channelId,
		Mode:           iotutil.ToString(channelType),
		Nickname:       nickname,
		Ip:             ip,
		HomeName:       defaultHomeName,
		AppKey:         appKey,
		TenantId:       tenantId,
		RegionServerId: regionServerId,
	})
	if err != nil {
		iotlogger.LogHelper.Error("用户%s登录失败，原因:%s", "", err.Error())
		return &entitys.LoginUserRes{}, -1, iotutil.ResErrToString(err)
	}
	//if AppLoginResponse.UserInfo == nil {
	//	iotlogger.LogHelper.Error("账号或密码错误")
	//	return &entitys.LoginUserRes{}, -1, "账号或密码错误"
	//}
	if AppLoginResponse.Token == "" {
		loginResponse := AppLoginInfo_pb2e(AppLoginResponse).SetShowVconsole(getShowVconsole(AppLoginResponse.UserInfo.Id))
		loginResponse.LoginKey = loginResponse.ThirdPartyLogin.Wechat.LoginKey
		loginResponse.ChannelName = loginResponse.ThirdPartyLogin.Wechat.Nickname
		loginResponse.UserId = ""
		return loginResponse, 0, "没有绑定用户请前往绑定"
		//if config.Global.Service.ThirdPartyRequirePwd {
		//	loginResponse := AppLoginInfo_pb2e(AppLoginResponse)
		//	loginResponse.LoginKey = loginResponse.ThirdPartyLogin.Wechat.LoginKey
		//	loginResponse.ChannelName = loginResponse.ThirdPartyLogin.Wechat.Nickname
		//	loginResponse.UserId = ""
		//	return loginResponse, 0, "没有绑定用户请前往绑定"
		//} else {
		//	//此处可以进行第三方登录，自动注册功能
		//	loginResponse, code, msg := s.RegisterAndLogin(entitys.RegisterNewUser{
		//		Ip:             ip,
		//		AppKey:         appKey,
		//		TenantId:       tenantId,
		//		ThirdType:      _const.Appleid,
		//		ThirdUserId:    channelId,
		//		ThirdNickname:  nickname,
		//		RegionServerId: regionServerId,
		//	})
		//	return loginResponse, code, msg
		//}
	}

	userinfo := AppUserInfo_pb2eModel("", AppLoginResponse.GetUserInfo())
	userinfo.TenantId = tenantId
	userinfo.AppKey = appKey
	expires := time.Unix(AppLoginResponse.ExpiresAt, 0).Sub(time.Now())
	cacheResp := iotredis.GetClient().Set(context.Background(), AppLoginResponse.Token, iotutil.ToString(userinfo), expires)
	if cacheResp.Err() != nil {
		iotlogger.LogHelper.Errorf("AppUserLogin,缓存token失败:%s", cacheResp.Err().Error())
		return &entitys.LoginUserRes{}, -1, iotutil.ResErrToString(cacheResp.Err())
	}
	res := AppLoginInfo_pb2e(AppLoginResponse).SetShowVconsole(getShowVconsole(AppLoginResponse.UserInfo.Id))
	res.LoginKey = res.ThirdPartyLogin.AppleId.LoginKey
	//第一次nickname才有值，如果没有值需要从历史记录中获取
	res.ChannelName = nickname
	if res.ChannelName == "" {
		res.ChannelName = res.ThirdPartyLogin.AppleId.Nickname
	}
	return res, 0, ""
}

// 账号绑定
func (s AppUserService) AccountBind(userId int64, accountBindParam entitys.AccountBind, appKey, tenantId string, regionId int64) (code int, msg string) {
	account := accountBindParam.Account
	accountType := accountBindParam.Type
	verificationCode := accountBindParam.Code

	resp := iotredis.GetClient().Get(context.Background(), cached.APP+"_"+appKey+"_"+account+"_6")

	if resp.Val() != verificationCode {
		iotlogger.LogHelper.Error("验证码有误")
		return -1, "验证码有误"
	}

	var phone, email string
	switch accountType {
	case 1:
		if iotutil.CheckAllPhone(accountBindParam.AreaPhoneNumber, account) == false {
			return -1, "手机号码不合法"
		}
		phone = account
	case 2:
		if iotutil.VerifyEmailFormat(account) == false {
			return -1, "电子邮箱不合法"
		}
		email = account
	default: //default case
		return -1, "用户类型有误"
	}

	data, err := rpc.TUcUserService.Find(context.Background(), &protosService.UcUserFilter{
		Phone:          phone,
		Email:          email,
		AppKey:         appKey,
		TenantId:       tenantId,
		RegionServerId: regionId,
	})
	if err == nil && data.Total != 0 {
		iotlogger.LogHelper.Error("该账号被注册了,无法进行绑定")
		return ioterrs.ERROR_EXIST_ACOUNT_NOT_BIND.Code, ioterrs.ERROR_EXIST_ACOUNT_NOT_BIND.Msg
	}
	_, err = rpc.TUcUserService.Update(context.Background(), &protosService.UcUser{
		Id:       userId,
		Phone:    phone,
		Email:    email,
		UserName: account,
	})
	if err != nil {
		iotlogger.LogHelper.Error("AccountBind error")
		return -1, "AccountBind error"
	}
	//删除redis中验证码
	iotredis.GetClient().Del(context.Background(), cached.APP+"_"+appKey+"_"+account+"_6")
	return 0, ""
}

// TODO 此接口调用rpc接口次数太多，需要进行优化
func (s AppUserService) CancelAccount(isVerifyCode bool, userId int64, cancelAccountParam entitys.CancelAccount, appKey, tenantId string, regionServerId int64, ip string) (code int, msg string) {
	account := cancelAccountParam.Account
	var err error
	if isVerifyCode {
		verificationCode := cancelAccountParam.Code

		_, _, err := iotutil.CheckUserName(account)
		if err != nil {
			return -1, err.Error()
		}

		resp := iotredis.GetClient().Get(context.Background(), cached.APP+"_"+appKey+"_"+account+"_4")
		if resp.Val() != verificationCode {
			iotlogger.LogHelper.Errorf("验证码有误[cached:%s-req:%s]", resp.Val(), verificationCode)
			return -1, "验证码有误"
		}
	}

	userInfo, err := rpc.TUcUserService.Lists(s.Ctx, &protosService.UcUserListRequest{
		Query: &protosService.UcUser{
			Id:             userId,
			Status:         1,
			AppKey:         appKey,
			TenantId:       tenantId,
			RegionServerId: regionServerId,
		},
	})
	if err != nil {
		iotlogger.LogHelper.Errorf("QueryQuestionTypeList error,%s", err.Error())
		return -1, err.Error()
	}
	if userInfo.Code != 200 || userInfo.Data == nil {
		return ioterrs.ERROR_NOT_BELONG_TO_USER.Code, ioterrs.ERROR_NOT_BELONG_TO_USER.Msg
	}

	_, err = rpc.TUcUserService.Update(context.Background(), &protosService.UcUser{
		Id:         userId,
		Status:     iotutil.ToInt32(_const.AccountAlreadyCancel),
		CancelTime: time.Now().Unix(),
	})
	if err != nil {
		return -1, "CancelAccount error"
	}

	//清理用户资产
	go clearUserAssets(userId, ip)

	//清除token
	controls.ClearTokenByUserId(userId)

	//删除redis中验证码
	iotredis.GetClient().Del(context.Background(), cached.APP+"_"+appKey+"_"+account+"_4")
	return 0, ""
}

func clearUserAssets(userId int64, ip string) {
	defer iotutil.PanicHandler("CancelAccount.clearUserAssets", userId)
	ucHomeUsersResponse, err := rpc.UcHomeUserService.Lists(context.Background(), &protosService.UcHomeUserListRequest{
		Query: &protosService.UcHomeUser{UserId: userId},
	})
	if err != nil {
		iotlogger.LogHelper.Errorf("注销账号获取家庭用户给列表失败, err: %v", err.Error())
		return
	}

	geo, err := controls.Geoip(ip)
	if err != nil {
		iotlogger.LogHelper.Errorf("注销账号，get address by ip[%s], error:%s", ip, err.Error())
	}

	//删除家庭用户信息
	_, err = rpc.UcHomeUserService.Delete(context.Background(), &protosService.UcHomeUser{
		UserId: userId,
	})
	if err != nil {
		iotlogger.LogHelper.Errorf("注销账号删除家庭用户失败, err: %v", err.Error())
	}
	//用户所在的家庭，并且当前用户为家庭所有者，则需要进行家庭数据清理
	homeIds := make([]int64, 0)
	for _, v := range ucHomeUsersResponse.Data {
		//owner-所有者, admin-管理者, member-成员
		//家庭所有者才进行数据清理
		if v.RoleType == iotconst.HOME_USER_ROLE_1 {
			homeIds = append(homeIds, v.HomeId)
		}
	}
	devList, err := rpc.IotDeviceHomeService.HomeDevList(context.Background(), &protosService.IotDeviceHomeHomeId{
		HomeIds: homeIds,
	})
	if len(devList.DevList) > 0 {
		homeDevIds := make(map[string][]string, 0)
		for _, d := range devList.DevList {
			if _, ok := homeDevIds[d.HomeId]; !ok {
				homeDevIds[d.HomeId] = make([]string, 0)
			}
			homeDevIds[d.HomeId] = append(homeDevIds[d.HomeId], d.Did)
		}
		//按照家庭分组删除设备
		for homeId, devs := range homeDevIds {
			_, err = rpc.UcHomeService.ChangeAllUserDefaultHomeId(context.Background(), &protosService.UcHome{
				Id:       iotutil.ToInt64(homeId),
				Lat:      geo.Lat,
				Lng:      geo.Lng,
				Country:  geo.Country,
				Province: geo.Province,
				City:     geo.City,
				District: geo.District,
			})
			if err != nil {
				iotlogger.LogHelper.Error("CancelAccount error")
				return
			}
			_, err = rpc.IotDeviceHomeService.RemoveDev(context.Background(), &protosService.RemoveDevRequest{
				HomeId:    homeId,
				UserId:    iotutil.ToString(userId),
				DevIdList: devs,
			})
		}
	}
	//清理用户的分享记录
	_, err = rpc.IotDeviceSharedService.Delete(context.Background(), &protosService.IotDeviceShared{
		BelongUserId: userId,
	})
	//清理用户的第三发登录信息
	_, err = rpc.ClientUcUserThirdService.Delete(context.Background(), &protosService.UcUserThird{
		UserId: userId,
	})
	if err != nil {
		iotlogger.LogHelper.Error("CancelAccount error")
	}
}

func AppRefreshToken_pb2e(src *protosService.AppUser) *controls.UserInfo {
	if src == nil {
		return nil
	}
	uiObj := controls.UserInfo{
		UserID:   src.Id,
		Nickname: src.NickName,
		Avatar:   src.Photo,
		Account:  src.Account,
	}
	return &uiObj
}

// RefreshToken 用户登录
func (s AppUserService) RefreshToken(req entitys.RefreshToken, appKey, tenantId string) (string, string, int64, error) {
	resp, err := rpc.AppAuthService.RefreshToken(context.Background(), &protosService.RefreshTokenRequest{
		RefreshToken: req.RefreshToken,
	})
	if err != nil {
		return "", "", 0, err
	}
	if resp == nil || (resp.GetValid() && resp.GetData() == nil) {
		return "", "", 0, errors.New("RefreshToken:内部服务错误")
	}
	if !resp.GetValid() {
		return "", "", 0, errors.New("RefreshToken:refresh token 已失效")
	}
	userinfo := AppRefreshToken_pb2e(resp.GetData().GetUserInfo())
	userinfo.TenantId = tenantId
	userinfo.AppKey = appKey
	expires := time.Unix(resp.GetData().GetExpiresAt(), 0).Sub(time.Now())
	cacheResp := iotredis.GetClient().Set(context.Background(), resp.GetData().GetToken(), iotutil.ToString(userinfo), expires)
	if cacheResp.Err() != nil {
		iotlogger.LogHelper.Errorf("RefreshToken:缓存token失败:%s", cacheResp.Err().Error())
	}
	controls.CacheTokenByUserId(userinfo.UserID, resp.GetData().GetToken(), resp.GetData().GetExpiresAt())
	return resp.GetData().GetToken(), resp.GetData().GetRefreshToken(), resp.GetData().GetExpiresAt(), nil
}

// CheckChannelBindParams 检查第三方渠道绑定接口中的参数
func (s AppUserService) CheckChannelBindParams(account string, smsCode string, channelType int32, bindType int32, channelId string) (code int, msg string) {
	if strings.TrimSpace(account) == "" {
		code = -1
		msg = "用户信息为空"
		return
	}
	if strings.TrimSpace(smsCode) == "" {
		code = -1
		msg = "验证码为空"
		return
	}
	if channelType == 0 {
		code = -1
		msg = "第三方渠道类型为空"
		return
	}
	if bindType == 0 {
		code = -1
		msg = "用户类型为空"
		return
	}
	if strings.TrimSpace(channelId) == "" {
		code = -1
		msg = "渠道userid为空"
		return
	}
	return
}

func (s AppUserService) AddChannelBind(req entitys.AddChannelBind, userId int64, appKey, tenantId string, regionServerId int64, nickName string) (*entitys.LoginUserRes, int, string) {
	//根据渠道类型、code查出渠道信息(第三方userid和昵称)
	channelUserId, channelNickname, code, msg := s.GetChannelInfo(req.Type, req.Code, req.ChannelId, req.Nickname, appKey)
	if code != 0 {
		//c.JSON(200,  sys.ErrorCode.HSet(c, gin.H{"code": code,  "msg": msg}))
		return nil, 2, msg
	}
	//todo 需判断当前第三方账号是否被其他用户账号绑定
	ucUserThirdResp, err := rpc.ClientUcUserThirdService.Find(context.Background(), &protosService.UcUserThirdFilter{
		ThirdType:      req.Type,
		ThirdUserId:    channelUserId,
		AppKey:         appKey,
		RegionServerId: regionServerId,
	})
	if err == nil && ucUserThirdResp.Data != nil && len(ucUserThirdResp.Data) > 0 {
		return nil, ioterrs.ERROR_BIND_BY_OTHER_USER.Code, ioterrs.ERROR_BIND_BY_OTHER_USER.Msg
	}
	//如果授权过了就获取不到昵称了
	if channelNickname == "" {
		ucUserThirdResp, err = rpc.ClientUcUserThirdService.Find(context.Background(), &protosService.UcUserThirdFilter{
			ThirdType:      req.Type,
			ThirdUserId:    channelUserId,
			AppKey:         appKey,
			RegionServerId: regionServerId,
			QueryDelete:    true,
		})
		if err == nil && len(ucUserThirdResp.Data) > 0 {
			channelNickname = ucUserThirdResp.Data[0].Nickname
		}
	}
	channelNickname = iotutil.IfStringEmpty(channelNickname, nickName)

	_, err = rpc.ClientUcUserThirdService.Create(context.Background(), &protosService.UcUserThird{
		Id:             iotutil.GetNextSeqInt64(),
		UserId:         userId,
		ThirdType:      req.Type,
		ThirdUserId:    channelUserId,
		Nickname:       channelNickname,
		CreatedBy:      userId,
		UpdatedBy:      userId,
		RegionServerId: regionServerId,
		AppKey:         appKey,
		TenantId:       tenantId,
		CreatedAt:      timestamppb.Now(),
		UpdatedAt:      timestamppb.Now(),
	})
	if err != nil {
		iotlogger.LogHelper.Errorf("addChannelBind error")
		return nil, 2, "addChannelBind error"
	}

	userInfo, err := rpc.TUcUserService.FindById(context.Background(), &protosService.UcUserFilter{
		Id: userId,
	})
	if err != nil {
		iotlogger.LogHelper.Errorf("getUserInfo error")
		return nil, 2, "getUserInfo error"
	}

	resp, err := rpc.ClientUcUserThirdService.Lists(context.Background(), &protosService.UcUserThirdListRequest{Query: &protosService.UcUserThird{UserId: userId}})
	if err != nil {
		iotlogger.LogHelper.Errorf("读取用户所有第三方登录方式失败，原因:%s", err.Error())
		return nil, 2, err.Error()
	}

	var thirdPartyLogin = entitys.AppUserThirdPartyLogin{}
	var channelId string
	for _, thirdPartys := range resp.Data {
		if thirdPartys.ThirdType == 1 {
			thirdPartyLogin.Wechat.Mode = thirdPartys.ThirdType
			thirdPartyLogin.Wechat.LoginKey = thirdPartys.ThirdUserId
			thirdPartyLogin.Wechat.Nickname = thirdPartys.Nickname
			channelId = thirdPartys.ThirdUserId
		} else {
			thirdPartyLogin.AppleId.Mode = thirdPartys.ThirdType
			thirdPartyLogin.AppleId.LoginKey = thirdPartys.ThirdUserId
			thirdPartyLogin.AppleId.Nickname = thirdPartys.Nickname
			channelId = thirdPartys.ThirdUserId
		}
	}
	user := userInfo.Data[0]

	pbObj := entitys.LoginUserRes{
		UserId:           iotutil.ToString(user.Id),
		NickName:         user.NickName,
		Phone:            user.Phone,
		Photo:            user.Photo,
		Email:            user.Email,
		DefaultHomeId:    user.DefaultHomeId,
		UserName:         user.UserName,
		SubmitCancelTime: user.CancelTime,
		AccountCasser:    false,
		ThirdPartyLogin:  thirdPartyLogin,
		Token:            "",
		RefreshToken:     "",
		ExpiresAt:        0,
		LoginKey:         channelId,
	}
	return &pbObj, 0, ""
}

// GetChannelInfo 获取第三方渠道信息
func (s AppUserService) GetChannelInfo(types int32, authorizationCode string, channelId string, nickname, appKey string) (channelUserId string, channelNickname string, msgCode int, msg string) {
	if types == 1 { //微信登录
		channelUserId, channelNickname, msgCode, msg = s.GetWechatInfo(authorizationCode, appKey)
	} else if types == 5 { //appleid登录
		if channelId == "" {
			msgCode = 100024
			msg = "第三方userid为空"
			return
		}
		channelUserId = channelId
		if nickname == "" {
			AppleidInfo, _ := rpc.UcAppleidInfoService.Find(context.Background(), &protosService.UcAppleidInfoFilter{
				ThirdUserId: channelUserId,
			})
			if AppleidInfo != nil && AppleidInfo.Data != nil {
				channelNickname = AppleidInfo.Data[0].Nickname
			}
		} else {
			channelNickname = nickname
		}

	}
	return
}

// 获取微信渠道信息
func (s AppUserService) GetWechatInfo(authorizationCode, appKey string) (channelUserId string, channelNickname string, msgCode int, msg string) {
	if authorizationCode == "" {
		msgCode = 100024
		msg = "授权Code为空"
		return
	}
	oemAppResult, err := rpc.ClientOemAppService.Find(context.Background(), &protosService.OemAppFilter{
		AppKey: appKey,
	})
	if err != nil {
		iotlogger.LogHelper.Errorf("GetFunctionConfig error")
		msgCode = 100024
		msg = err.Error()
		return
	}
	if oemAppResult.Code != 200 {
		iotlogger.LogHelper.Errorf(oemAppResult.Message)
		msgCode = 100024
		msg = oemAppResult.Message
		return
	}
	oemAppInfo := oemAppResult.Data[0]
	oemAppFunctionConfig, err := rpc.ClientOemAppFunctionConfigService.Find(context.Background(), &protosService.OemAppFunctionConfigFilter{
		AppId:   oemAppInfo.Id,
		Version: oemAppInfo.Version,
	})
	if err != nil {
		iotlogger.LogHelper.Errorf("GetFunctionConfig error")
		msgCode = 100024
		msg = err.Error()
		return
	}
	if oemAppFunctionConfig.Code != 200 {
		iotlogger.LogHelper.Errorf(oemAppFunctionConfig.Message)
		msgCode = 100024
		msg = oemAppFunctionConfig.Message
		return
	}

	if oemAppFunctionConfig.Data[0].Thirds == "" {
		msgCode = 100024
		msg = "appId、appSecret为空"
		return
	}
	thirds := iotutil.JsonToMapArray(oemAppFunctionConfig.Data[0].Thirds)

	var appId, appSecret string
	for _, third := range thirds {
		if third["code"] == "wechat" {
			appId = iotutil.ToString(third["appId"])
			appSecret = iotutil.ToString(third["appKey"])
			break
		}
	}

	//appId, appSecret, authorizationCode查询accessToken
	accessToken, err := wechat.GetOauth2AccessToken(context.Background(), appId, appSecret, authorizationCode)
	if err != nil {
		msgCode = 100024
		msg = err.Error()
		return
	}
	if accessToken.Openid == "" {
		msgCode = 100025
		msg = accessToken.Errmsg
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

func (s AppUserService) UnbindChannel(req entitys.UnbundlingChannel, userId int64, appKey, tenantId string, regionServerId int64) (int, string) {
	_, err := rpc.ClientUcUserThirdService.Delete(context.Background(), &protosService.UcUserThird{
		UserId:         userId,
		ThirdUserId:    req.ChannelId,
		AppKey:         appKey,
		RegionServerId: regionServerId,
		TenantId:       tenantId,
	})
	if err != nil {
		iotlogger.LogHelper.Errorf("UnbindChannel error")
		return -1, "UnbindChannel error"
	}

	return 0, ""
}

func (s AppUserService) GetAppId(channelType int32, userId int64, appKey string) (int, string, interface{}) {
	resultList := []map[string]interface{}{}
	oemAppResult, err := rpc.ClientOemAppService.Find(context.Background(), &protosService.OemAppFilter{
		AppKey: appKey,
	})
	if err != nil {
		iotlogger.LogHelper.Errorf("GetAppId error")
		return -1, err.Error(), resultList
	}
	if oemAppResult.Code != 200 {
		iotlogger.LogHelper.Errorf(oemAppResult.Message)
		return -1, oemAppResult.Message, resultList
	}
	oemAppInfo := oemAppResult.Data[0]
	oemAppFunctionConfig, err := rpc.ClientOemAppFunctionConfigService.Find(context.Background(), &protosService.OemAppFunctionConfigFilter{
		AppId:   oemAppInfo.Id,
		Version: oemAppInfo.Version,
	})
	if err != nil {
		iotlogger.LogHelper.Errorf("GetAppId error")
		return -1, err.Error(), resultList
	}
	if oemAppFunctionConfig.Code != 200 {
		iotlogger.LogHelper.Errorf(oemAppFunctionConfig.Message)
		return -1, oemAppFunctionConfig.Message, resultList
	}

	if oemAppFunctionConfig.Data[0].Thirds == "" {
		return -1, "thirds is empty", resultList
	}
	thirds := iotutil.JsonToMapArray(oemAppFunctionConfig.Data[0].Thirds)

	for _, third := range thirds {
		//此写法为了兼容未配置isCheck之前的数据
		if v, ok := third["isCheck"]; ok {
			isCheck := iotutil.ToString(v)
			if isCheck != "1" {
				continue
			}
		}

		var code, appId string
		switch third["code"].(string) {
		case "wechat":
			code = "wechatAppID"
			appId = iotutil.ToString(third["appId"])
		case "apple", "appleid": //兼容处理
			code = "appleAppID"
			appId = "-" //解决IOSJSON处理空值的Key丢失问题
		case "google":
			code = "googleClientID"
			appId = iotutil.ToString(third["appId"])
		default:
			code = iotutil.ToString(third["code"])
			appId = iotutil.ToString(third["appId"])
		}
		result := map[string]interface{}{}
		result["type"] = channelType
		result["code"] = code
		result["appId"] = appId
		resultList = append(resultList, result)
	}
	return 0, "", resultList
}

func (s AppUserService) GetFunctionConfigVoice(userId int64, appKey string) (int, string, []map[string]interface{}) {
	lang, _ := metadata.Get(s.Ctx, "lang")
	resultList := []map[string]interface{}{}
	oemAppResult, err := rpc.ClientOemAppService.Find(context.Background(), &protosService.OemAppFilter{
		AppKey: appKey,
	})
	if err != nil {
		iotlogger.LogHelper.Errorf("GetFunctionConfig error")
		return -1, err.Error(), resultList
	}
	if oemAppResult.Code != 200 {
		iotlogger.LogHelper.Errorf(oemAppResult.Message)
		return -1, oemAppResult.Message, resultList
	}
	oemAppInfo := oemAppResult.Data[0]
	oemAppFunctionConfig, err := rpc.ClientOemAppFunctionConfigService.Find(context.Background(), &protosService.OemAppFunctionConfigFilter{
		AppId:   oemAppInfo.Id,
		Version: oemAppInfo.Version,
	})
	if err != nil {
		iotlogger.LogHelper.Errorf("GetFunctionConfig error")
		return -1, err.Error(), resultList
	}
	if oemAppFunctionConfig.Code != 200 {
		iotlogger.LogHelper.Errorf(oemAppFunctionConfig.Message)
		return -1, oemAppFunctionConfig.Message, resultList
	}

	if oemAppFunctionConfig.Data[0].Voices == "" {
		return -1, "voices is empty", resultList
	}
	voices := iotutil.JsonToMapArray(oemAppFunctionConfig.Data[0].Voices)

	oemAppIntroduce, err := rpc.ClientOemAppIntroduceService.Lists(context.Background(), &protosService.OemAppIntroduceListRequest{
		Query: &protosService.OemAppIntroduce{
			AppId: oemAppInfo.Id,
			// Version:     oemAppInfo.Version,
			Lang: lang,
			// Status:      2,
			ContentType: 4,
		},
	})
	if err != nil {
		iotlogger.LogHelper.Errorf("GetOemAppIntroduce error")
		//return -1, err.Error(), resultList
	}

	appIntroduceKeyValue := map[string]string{}
	if oemAppIntroduce.Code == 200 && oemAppIntroduce.Data != nil {
		for _, appIntroduce := range oemAppIntroduce.Data {
			appIntroduceKeyValue[appIntroduce.VoiceCode] = appIntroduce.Abstract
		}
	}

	opmVoice, err := rpc.ClientOpmVoiceService.Lists(context.Background(), nil)
	if err != nil {
		iotlogger.LogHelper.Errorf("GetOemAppIntroduce error")
		//return -1, err.Error(), resultList
	}

	opmVoiceKeyValue := map[string]string{}
	opmVoiceName := map[string]string{}
	if opmVoice.Code == 200 && opmVoice.Data != nil {
		for _, opmVoiceInfo := range opmVoice.Data {
			opmVoiceKeyValue[opmVoiceInfo.VoiceNo] = opmVoiceInfo.VoiceLogo
			if lang == "zh" {
				opmVoiceName[opmVoiceInfo.VoiceNo] = opmVoiceInfo.VoiceName
			} else {
				opmVoiceName[opmVoiceInfo.VoiceNo] = opmVoiceInfo.VoiceEnName
			}
		}
	}

	for _, voice := range voices {
		var title string
		if voice["code"] == "tianmao" {
			title = "连接天猫精灵"
		} else if voice["code"] == "xiaodu" {
			title = "连接小度音响"
		} else if voice["code"] == "alexa" {
			title = "连接Alexa"
		} else if voice["code"] == "google" {
			title = "连接google"
		} else if voice["code"] == "xiaoai" {
			title = "连接小爱音响"
		} else {
			title = "未知"
		}
		result := map[string]interface{}{}
		result["title"] = title
		code := iotutil.ToString(voice["code"])
		result["name"] = opmVoiceName[code]
		result["url"] = opmVoiceKeyValue[code]
		result["code"] = code
		result["abstract"] = appIntroduceKeyValue[code]
		resultList = append(resultList, result)
	}
	return 0, "", resultList
}

// 获取功能配置
func (s AppUserService) GetFunctionConfig(appKey string) (*entitys.OemAppFunctionConfig, error) {
	oemAppResult, err := rpc.ClientOemAppService.Find(context.Background(), &protosService.OemAppFilter{
		AppKey: appKey,
	})
	if err != nil {
		iotlogger.LogHelper.Errorf("GetFunctionConfig error")
		return nil, err
	}
	if oemAppResult.Code != 200 {
		iotlogger.LogHelper.Errorf(oemAppResult.Message)
		return nil, err
	}
	res, err := rpc.ClientOemAppFunctionConfigService.Find(s.Ctx, &protosService.OemAppFunctionConfigFilter{
		AppId:   oemAppResult.Data[0].Id,
		Version: oemAppResult.Data[0].Version,
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return nil, errors.New(res.Message)
	}
	if len(res.Data) == 0 || res.Data == nil {
		var nodata = entitys.OemAppFunctionConfig{}
		//nodata.AppId = req.AppId
		return &nodata, nil
	}
	var data = entitys.OemAppFunctionConfig{}

	//data.Eula = res.Data[0].Eula
	data.Aboutus = res.Data[0].AboutUs
	data.Weather = res.Data[0].Weather
	data.Geo = res.Data[0].Geo
	//data.Privacypolicy = res.Data[0].PrivacyPolicy
	data.Id = iotutil.ToString(res.Data[0].Id)
	//data.AppId = req.AppId

	return &data, err
}

func getNameByLang(lang, cnName, enName string) string {
	switch lang {
	case "zh":
		return cnName
	case "en":
		return enName
	}
	return ""
}
