package services

import (
	"cloud_platform/iot_cloud_api_service/cached"
	"cloud_platform/iot_cloud_api_service/controls"
	services "cloud_platform/iot_cloud_api_service/controls/global"
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"net"
	"strconv"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type OpenUserService struct {
	Ctx context.Context
}

func (s OpenUserService) SetContext(ctx context.Context) OpenUserService {
	s.Ctx = ctx
	return s
}

func (s OpenUserService) OpenUserError(res *protosService.Response, err error) error {
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	return nil
}

// 注册用户
func (s OpenUserService) RegisterUser(req entitys.OpenUserRegisterReq) (string, error) {
	//1.验证码检查
	isVer, msgVer := s.VerificationCode(req.UserName, req.VerifyCode)
	if !isVer {
		return "", errors.New(msgVer)
	}

	if err := req.CheckRegister(); err != nil {
		return "", err
	}
	res, err := rpc.ClientOpenUserService.Register(s.Ctx, &protosService.OpenUserRegisterRequest{
		Account:     req.UserName,
		UserType:    req.AccountType,
		CompanyName: req.CompanyName,
		Password:    req.UserPassword,
		VerifyCode:  req.VerifyCode,
		IP:          req.IP,
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	//刷新缓存
	services.RefreshDevelopCache()
	return iotutil.ToString(res.Data), nil
}

// UserLogin 用户登录
func (s OpenUserService) UserLogin(req entitys.UserLoginReq, ip string) (string, string, int64, error) {
	resp, err := rpc.ClientOpenAuthService.PasswordLogin(s.Ctx, &protosService.PasswordLoginRequest{
		Channel:      "PC",
		LoginName:    req.Username,
		Password:     req.Password,
		VerifyCode:   req.Verifycode,
		PlatformType: string(iotconst.PLATFORMTYPE_OPEN),
		ClientIp:     req.ClientIp,
		Explorer:     req.Explorer,
		Os:           req.Os,
	})
	if err != nil {
		return "", "", 0, err
	}
	if resp == nil || resp.UserInfo == nil {
		return "", "", 0, errors.New("内部服务错误")
	}
	//return "", "", 0, err
	userinfo := OpenUserInfo_pb2e(resp.GetUserInfo())
	expires := time.Unix(resp.ExpiresAt, 0).Sub(time.Now())
	err = cached.RedisStore.Set(resp.Token, *userinfo, expires)
	if err != nil {
		iotlogger.LogHelper.Errorf("UserLogin,缓存token失败:%s", err.Error())
	}
	controls.CacheTokenByUserId(userinfo.UserID, resp.Token, resp.GetExpiresAt())
	//for debug
	//var userInfo controls.OpenUserInfo
	//err = cached.RedisStore.Get(resp.Token, &userInfo)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for debug
	loginInfo := iotstruct.OpenUserLogin{UserId: userinfo.UserID, TenantId: userinfo.TenantId, LoginTime: time.Now().Unix(), ExpiresAt: resp.GetExpiresAt(), Ip: ip, Addr: "", Token: resp.Token}
	s.CacheLoginInfo(&loginInfo)

	return resp.GetToken(), resp.GetRefreshToken(), resp.GetExpiresAt(), nil
}
func (s OpenUserService) UserLogout(token string) error {
	if token == "" {
		return nil
	}
	//删除redis即可
	tokens := strings.Split(token, " ")
	if len(tokens) == 2 {
		err := cached.RedisStore.Delete(tokens[1])
		if err != nil {
			return err
		}
		s.ClearLoginInfo(tokens[1])
	}
	return nil
}

func OpenUserInfo_pb2e(src *protosService.CloudUserInfo) *controls.UserInfo {
	if src == nil {
		return nil
	}
	uiObj := controls.UserInfo{
		UserID:   src.UserId,
		Nickname: src.NickName,
		Avatar:   src.Avatar,
		TenantId: src.TenantId,
	}
	return &uiObj
}

// 修改用户密码
func (s OpenUserService) UpdatePassword(userId string, req entitys.OpenUserUpdatePasswordReq) error {
	//todo 验证用户密码是否正确
	userid := iotutil.ToInt64(userId)
	//获取原用户信息
	user, errUser := rpc.ClientOpenUserService.FindById(context.Background(), &protosService.OpenUserFilter{
		Id: userid,
	})
	if errUser != nil {
		return errUser
	}
	if user.Code != 200 {
		return errors.New(user.Message)
	}
	if len(user.Data) <= 0 {
		return errors.New("userid not exists")
	}
	if user.Data[0].UserName != req.UserName {
		return errors.New("参数错误, 只能修改自己的密码.")
	}
	isVer, msgVer := s.VerificationCode(req.UserName, req.VerifyCode)
	if !isVer {
		return errors.New(msgVer)
	}
	salt := user.Data[0].UserSalt
	password := iotutil.Md5(req.NewPassword + salt)
	res, err := rpc.ClientOpenUserService.UpdateFields(context.Background(), &protosService.OpenUserUpdateFieldsRequest{
		Fields: []string{"user_password", "updated_at"},
		Data:   &protosService.OpenUser{Id: userid, UserPassword: password, UpdatedAt: timestamppb.Now()},
	})

	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	return err
}

func (s OpenUserService) ForgetPassword(req entitys.OpenUserUpdatePasswordReq) error {
	//获取原用户信息
	user, errUser := rpc.ClientOpenUserService.Find(context.Background(), &protosService.OpenUserFilter{
		UserName:   req.UserName,
		UserStatus: 1,
	})
	if errUser != nil {
		return errUser
	}
	if user.Code != 200 {
		return errors.New(user.Message)
	}
	if len(user.Data) <= 0 {
		return errors.New("用户名不存在")
	}
	isVer, msgVer := s.VerificationCode(req.UserName, req.VerifyCode)
	if !isVer {
		return errors.New(msgVer)
	}
	salt := user.Data[0].UserSalt
	password := iotutil.Md5(req.NewPassword + salt)
	res, err := rpc.ClientOpenUserService.UpdateFields(context.Background(), &protosService.OpenUserUpdateFieldsRequest{
		Fields: []string{"user_password", "updated_at"},
		Data:   &protosService.OpenUser{Id: user.Data[0].Id, UserPassword: password, UpdatedAt: timestamppb.Now()},
	})

	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	//清除token
	controls.ClearTokenByUserId(user.Data[0].Id)
	return err
}

// QueryUserList 查询用户列表
func (s OpenUserService) GetUserProfile(userId string, tenantId string) (*entitys.OpenUserProfileRes, error) {
	var resUser entitys.OpenUserProfileRes
	rep, err := rpc.ClientOpenUserService.FindById(s.Ctx, &protosService.OpenUserFilter{
		Id: iotutil.ToInt64(userId),
	})
	if err != nil {
		return nil, err
	}
	if rep.Code != 200 {
		return nil, errors.New(rep.Message)
	}
	if rep.Data == nil || len(rep.Data) == 0 {
		return nil, errors.New("record not found")
	}

	resCompany, errCompany := rpc.ClientOpenUserCompanyService.Lists(s.Ctx, &protosService.OpenUserCompanyListRequest{
		Page:     1,
		PageSize: 1000000,
		Query: &protosService.OpenUserCompany{
			UserId: iotutil.ToInt64(userId),
		},
	})
	if errCompany != nil {
		return nil, err
	}
	if resCompany.Code != 200 {
		return nil, errors.New(resCompany.Message)
	}
	if len(resCompany.Data) == 0 {
		return nil, errors.New("未查到任何关联公司")
	}

	var list []*entitys.OpenProfileUserCompany
	currentCompanyName := ""
	currentUserType := 1

	for _, v := range resCompany.Data {

		//TODO 后续优化考虑缓存.或是数据冗余
		company, _ := rpc.ClientOpenCompanyService.FindById(s.Ctx, &protosService.OpenCompanyFilter{Id: v.CompanyId})

		companyName := company.Data[0].Name
		//如果是个人账号. 空间名称则显示对方的账号名.
		if company.Data[0].AccountType == 2 {
			//要用companyId  查询对应的用户名 TODO  有时间考虑存入字段
			companyName = company.Data[0].UserName //company.Data[0].UserName
		}
		tmp := entitys.OpenProfileUserCompany{
			Id:          iotutil.ToString(v.Id),
			CompanyName: companyName, //TODO  需要改为公司名
			TenantId:    v.TenantId,
			UserType:    v.UserType,
		}
		list = append(list, &tmp)
		if v.TenantId == tenantId {
			//如果当前是主账号.companyName设置空. 这样前端则显示我的空间.
			if v.UserType == 1 {
				currentCompanyName = ""
			} else if v.UserType == 2 {
				//如果是子账号. 企业账号显示公司名, 个人账号显示用户名
				if company.Data[0].AccountType == 1 {
					currentCompanyName = company.Data[0].Name
				} else {
					currentCompanyName = company.Data[0].UserName
				}
			}
			currentUserType = int(v.UserType)
		}
	}

	resUser = entitys.OpenUserProfileRes{
		Id:          iotutil.ToString(rep.Data[0].Id),
		UserName:    rep.Data[0].UserName,
		UserStatus:  rep.Data[0].UserStatus,
		AccountType: rep.Data[0].AccountType,
		Avatar:      rep.Data[0].Avatar,
		TenantId:    tenantId,
		CompanyName: currentCompanyName,
		UserType:    int32(currentUserType),
		TenantList:  list,
	}
	return &resUser, nil
}

func (s OpenUserService) GetUserCompanyTenantIds(userName string) ([]string, error) {
	resCompany, errCompany := rpc.ClientOpenUserCompanyService.Lists(s.Ctx, &protosService.OpenUserCompanyListRequest{
		Query: &protosService.OpenUserCompany{
			UserName: userName,
			UserType: 1,
		},
	})
	if errCompany != nil {
		return nil, errCompany
	}
	if resCompany.Code != 200 {
		return nil, errors.New(resCompany.Message)
	}
	var tenantIds []string
	for _, item := range resCompany.Data {
		tenantIds = append(tenantIds, item.TenantId)
	}
	return tenantIds, nil
}

func (s OpenUserService) RefreshToken(req entitys.RefreshToken, ip string) (string, string, int64, error) {
	resp, err := rpc.ClientOpenAuthService.RefreshToken(context.Background(), &protosService.RefreshTokenRequest{
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
	userinfo := OpenUserInfo_pb2e(resp.GetData().GetUserInfo())
	expires := time.Unix(resp.GetData().GetExpiresAt(), 0).Sub(time.Now())
	err = cached.RedisStore.Set(resp.GetData().GetToken(), *userinfo, expires)
	if err != nil {
		iotlogger.LogHelper.Errorf("RefreshToken:缓存token失败:%s", err.Error())
	}
	controls.CacheTokenByUserId(userinfo.UserID, resp.GetData().GetToken(), resp.GetData().GetExpiresAt())

	loginInfo := iotstruct.OpenUserLogin{UserId: userinfo.UserID, TenantId: userinfo.TenantId, LoginTime: time.Now().Unix(), ExpiresAt: resp.GetData().GetExpiresAt(), Ip: ip, Addr: "", Token: resp.GetData().GetToken()}
	s.CacheLoginInfo(&loginInfo)

	return resp.GetData().GetToken(), resp.GetData().GetRefreshToken(), resp.GetData().GetExpiresAt(), nil
}

func (s OpenUserService) ChangeTenant(req entitys.RefreshToken) (string, string, int64, error) {
	resp, err := rpc.ClientOpenAuthService.ChangeTenant(context.Background(), &protosService.RefreshTokenRequest{
		RefreshToken: req.RefreshToken,
		TenantId:     req.TenantId,
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
	userinfo := OpenUserInfo_pb2e(resp.GetData().GetUserInfo())
	expires := time.Unix(resp.GetData().GetExpiresAt(), 0).Sub(time.Now())
	err = cached.RedisStore.Set(resp.GetData().GetToken(), *userinfo, expires)
	if err != nil {
		iotlogger.LogHelper.Errorf("RefreshToken:缓存token失败:%s", err.Error())
	}
	return resp.GetData().GetToken(), resp.GetData().GetRefreshToken(), resp.GetData().GetExpiresAt(), nil
}

func (s OpenUserService) GetRouters() (*[]*entitys.Open2AuthRuleEntitysTree, error) {
	res, err := rpc.ClientOpenUserService.GetRouters(s.Ctx, &protosService.OpenUserPrimarykey{})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	if len(res.MenuTree) == 0 {
		var nodata *[]*entitys.Open2AuthRuleEntitysTree = &[]*entitys.Open2AuthRuleEntitysTree{}
		return nodata, nil
	}
	var resMenuList *[]*entitys.Open2AuthRuleEntitysTree = &[]*entitys.Open2AuthRuleEntitysTree{}
	//  for _, v := range res.MenuTree {
	//  	menu := entitys.Open2AuthRule_pb2e(v)
	//  	*resMenuList = append(*resMenuList, menu)
	// }

	s.GetMenuTreeList(res.MenuTree, resMenuList)

	result := s.SetUserMenuTree("0", *resMenuList)

	return result, nil
}

func (s OpenUserService) GetMenuTreeList(list []*protosService.OpenMenuTree, res *[]*entitys.Open2AuthRuleEntitysTree) {
	for _, v := range list {
		menu := entitys.Open2AuthRule_pb2e(v)
		*(res) = append(*(res), menu)
		if v.Children != nil && len(v.Children) > 0 {
			s.GetMenuTreeList(v.Children, res)
		}
	}
}

// 递归形成菜单树
func (s *OpenUserService) SetUserMenuTree(pid string, menuList []*entitys.Open2AuthRuleEntitysTree) *[]*entitys.Open2AuthRuleEntitysTree {
	//tree := make(*[]*entitys.Open2AuthRuleEntitysTree, 0, len(menuList))
	var tree *[]*entitys.Open2AuthRuleEntitysTree = &[]*entitys.Open2AuthRuleEntitysTree{}
	for _, v := range menuList {
		if v.Pid == pid {
			child := s.SetUserMenuTree(v.Id, menuList)
			if child != nil {
				v.Children = child
			}
			*(tree) = append((*tree), v)
		}
	}
	return tree
}

// 发送验证码(会验证用户名是否存在) 忘记密码和登录的时候使用.
func (s OpenUserService) SendVerificationCodeForExists(userName, lang string, codeType int32) (string, error) {
	res, err := rpc.ClientOpenUserService.Find(context.Background(), &protosService.OpenUserFilter{
		UserName:   userName,
		UserStatus: 1,
	})
	if err != nil {
		return "", err
	}

	if res.Code != 200 {
		if res.Message == "record not found" {
			return "", errors.New("用户名不存在")
		}
		return "", errors.New(res.Message)
	}
	if len(res.Data) == 0 {
		return "", errors.New("用户名不存在")
	}
	return s.SendVerificationCode(userName, lang, codeType)
}

// 发送验证码
func (s OpenUserService) SendVerificationCode(userName, lang string, codeType int32) (string, error) {
	if lang == "" {
		lang = "zh"
	}
	index := strings.Index(userName, "@")
	code := iotutil.GetRandomNumber(4) //租户Id
	if index > 0 {
		res, err := rpc.ClientEmailService.SendEmailUserCode(context.Background(), &protosService.SendEmailUserCodeRequest{
			Email:    userName,
			UserName: userName,
			Code:     code,
			Lang:     lang,
			TplType:  codeType,
		})
		if err != nil {
			return "", err
		}
		if !res.Status {
			return "error", errors.New("验证码获取错误")
		}
	} else {
		var phoneType int32
		if iotutil.CheckPhone(userName) == true {
			phoneType = 1
		} else {
			return "", errors.New("手机号不合法")
		}
		res, err := rpc.ClientSmsService.SendSMSVerifyCode(context.Background(), &protosService.SendSMSVerifyCodeRequest{
			PhoneNumber: userName,
			UserName:    userName,
			Code:        code,
			Lang:        lang,
			TplType:     codeType,
			PhoneType:   phoneType,
		})
		if err != nil {
			return "", err
		}
		if !res.Status {
			return "error", errors.New("验证码获取错误")
		}

	}
	//验证码10分钟失效
	expires := time.Minute * 2 //time.Unix(600, 0).Sub(time.Now())
	cached.RedisStore.Set(userName, code, expires)
	return "ok", nil
}

// 验证验证码
func (s OpenUserService) VerificationCode(userName string, code string) (bool, string) {
	var codeRedis string
	cached.RedisStore.Get(userName, &codeRedis)
	if codeRedis == "" {
		return false, "验证码已失效"
	}

	if codeRedis != code {
		return false, "验证码错误"
	}
	return true, "ok"

}

// //递归进行转换
// func (s OpenUserService) ConvertTree(menuList []*protosService.OpenMenuTree) []*entitys.OpenAuthRuleEntitysTree {
// 	var resMenuList  []*entitys.OpenAuthRuleEntitysTree
// 	for _,v := range menuList{
// 		menu := entitys.OpenAuthRule_pb2e(v)
// 		resMenuList = append(resMenuList, menu)
// 		if v.Children != nil || len(v.Children) > 0{
// 			return s.ConvertTree(v.Children)
// 		}
// 	}
// 	return resMenuList
// }

// QueryUserList 查询用户列表
func (s OpenUserService) QueryAppUserList(pageNum, pageSize int64, userMobile, userAccount, userEmail, appType string) ([]entitys.QueryUserListRsp, int64, error) {
	userListRsp := []entitys.QueryUserListRsp{}
	appKeyByOemId := ""
	if appType != "" {
		appTypeInt64, _ := iotutil.ToInt64AndErr(appType)
		oemAppInfo, err := rpc.ClientOemAppService.FindById(s.Ctx, &protosService.OemAppFilter{
			Id: appTypeInt64,
		})
		if err != nil {
			return userListRsp, 0, nil
		}
		if oemAppInfo.Code != 200 {
			return userListRsp, 0, nil
		}
		appKeyByOemId = oemAppInfo.Data[0].AppKey
	}
	reponse, err := rpc.UcUserService.Lists(s.Ctx, &protosService.UcUserListRequest{
		Page:      pageNum,
		PageSize:  pageSize,
		OrderKey:  "created_at",
		OrderDesc: "desc",
		Query: &protosService.UcUser{
			Phone:    userMobile,
			UserName: userAccount,
			Email:    userEmail,
			Status:   1, //主要有效的账号才显示再列表中
			//AppOrigin: appType,
			AppKey: appKeyByOemId,
			//Status:    1,
		},
	})
	if err != nil {
		return userListRsp, 0, nil
	}
	if reponse.Code != 200 {
		return userListRsp, 0, nil
	}
	if reponse.Data == nil || len(reponse.Data) == 0 {
		return userListRsp, 0, nil
	}

	// 创建一个临时map用来存储数组元素,去重复数据
	temp := make(map[string]bool)
	appKeys := []string{}
	for _, v := range reponse.Data {
		if v.AppKey == "" {
			continue
		}
		_, ok := temp[v.AppKey]
		if ok {
			continue
		} else {
			temp[v.AppKey] = true
		}
		appKeys = append(appKeys, v.AppKey)
	}

	oemAppData, err := rpc.ClientOemAppService.ListsByAppKeys(s.Ctx, &protosService.ListsByAppKeysRequest{
		AppKeys: appKeys,
	})
	if err != nil {
		return userListRsp, 0, nil
	}
	keyValue := make(map[string]string)
	if oemAppData != nil {
		for _, oemApp := range oemAppData.Data {
			keyValue[oemApp.AppKey] = oemApp.Name
		}
	}

	for _, v := range reponse.Data {
		var appName string
		val, ok := keyValue[v.AppKey]
		if !ok {
			continue
		}
		appName = val
		resUser := entitys.QueryUserListRsp{
			UserId:       iotutil.ToString(v.Id),
			UserMobile:   v.Phone,
			UserAccount:  v.UserName,
			UserEmail:    v.Email,
			Nickname:     v.NickName,
			AppName:      appName,
			RegisterTime: v.CreatedAt.AsTime().Unix(),
		}
		userListRsp = append(userListRsp, resUser)
	}
	return userListRsp, reponse.Total, nil
}

// QueryUserDeviceList app用户绑定的设备
func (s OpenUserService) QueryUserDeviceList(pageNum, pageSize int64, customerUserId string) ([]entitys.QueryUserDeviceList, int64, error) {
	queryUserDeviceList := []entitys.QueryUserDeviceList{}
	rep, err := rpc.UcHomeUserService.Lists(s.Ctx, &protosService.UcHomeUserListRequest{
		Query: &protosService.UcHomeUser{
			UserId: iotutil.ToInt64(customerUserId),
		},
	})
	if err != nil {
		return queryUserDeviceList, 0, nil
	}
	if rep.Code != 200 {
		return queryUserDeviceList, 0, nil
	}
	if rep.Data == nil || len(rep.Data) == 0 {
		return queryUserDeviceList, 0, nil
	}

	var homeIdList []int64
	roleTypeMap := map[int64]int32{}
	for _, v := range rep.Data {
		homeIdList = append(homeIdList, v.HomeId)
		roleTypeMap[v.HomeId] = v.RoleType
	}

	userHomeRep, err := rpc.IotDeviceHomeService.UserDev(s.Ctx, &protosService.IotUserHomeDev{ //获取用户所有家庭中所有设备id
		Page:     pageNum,
		PageSize: pageSize,
		HomeIds:  homeIdList,
	})
	if err != nil || userHomeRep.Total == 0 {
		return queryUserDeviceList, 0, nil
	}
	for _, devObj := range userHomeRep.Data { //循环设备id
		userDevice := entitys.QueryUserDeviceList{}
		userDevice.AddMethod = 1 //用户自己添加
		if roleTypeMap[devObj.HomeId] == 3 {
			userDevice.AddMethod = 2 //他人共享
		}

		IotDeviceTriadRep, err := rpc.ClientIotDeviceServer.Find(s.Ctx, &protosService.IotDeviceTriadFilter{ //
			Did: devObj.DeviceId,
			//UseType: 1, //虚拟设备
		})
		if err == nil && IotDeviceTriadRep.Total > 0 {
			if IotDeviceTriadRep.Data[0].UseType == 1 {
				userDevice.AddMethod = 3 //虚拟设备
			}
		}

		IotDeviceInfoRep, err := rpc.ClientIotDeviceInfoServer.Find(s.Ctx, &protosService.IotDeviceInfoFilter{ //获取设备名称
			Did: devObj.DeviceId,
		})
		if err == nil && IotDeviceInfoRep.Total > 0 {
			userDevice.DeviceName = IotDeviceInfoRep.Data[0].DeviceName
			userDevice.DevicePid = iotutil.ToString(IotDeviceInfoRep.Data[0].Did)

			productRes, err := rpc.ClientOpmProductService.FindById(s.Ctx, &protosService.OpmProductFilter{ //获取产品名称
				Id: IotDeviceInfoRep.Data[0].ProductId,
			})
			if err == nil && productRes.Data != nil {
				productInfo := productRes.Data[0]
				userDevice.ProductName = productInfo.Name                                                             //产品类型
				productTypeRes, err := rpc.ClientProductService.GetTPmProduct(s.Ctx, &protosService.TPmProductFilter{ //所属产品
					Id: productInfo.ProductId,
				})
				if err == nil && productTypeRes.Data != nil {
					userDevice.ProductTypeName = productTypeRes.Data.Name
				}
			}
		}
		if userDevice.DeviceName == "" {
			continue
		}
		queryUserDeviceList = append(queryUserDeviceList, userDevice)
	}

	return queryUserDeviceList, userHomeRep.Total, nil
}

func (s OpenUserService) QueryLangBaseDataList(filter entitys.LangBaseDataQuery) ([]*protosService.ConfigDictData, error) {
	QueryObj := protosService.ConfigDictData{
		DictLabel: filter.DictLabel,
		DictType:  filter.DictType,
	}
	ret, err := rpc.TConfigDictDataServerService.Lists(context.Background(), &protosService.ConfigDictDataListRequest{
		Query: &QueryObj,
	})
	if err != nil {
		return nil, err
	}
	if ret.Code != 200 {
		return nil, errors.New(ret.Message)
	}
	return ret.Data, err
}

// 缓存登录信息
func (s OpenUserService) CacheLoginInfo(login *iotstruct.OpenUserLogin) {
	login.Addr = GetAddrByIp(login.Ip)
	_, err := iotredis.GetClient().HSet(context.Background(), iotconst.USERLASTLOGIN, strconv.Itoa(int(login.UserId)), login).Result()
	if err != nil {
		return
	}
	var ids []string = []string{"22", "5083013370315964416", "33333"}
	GetDeveloperLoginInfo(ids)
}

// 缓存登录信息
func (s OpenUserService) ClearLoginInfo(token string) {
	resp, err := rpc.ClientOpenAuthService.RefreshToken(context.Background(), &protosService.RefreshTokenRequest{
		RefreshToken: token,
	})
	if err != nil || resp == nil || resp.GetData() == nil || resp.GetData().GetUserInfo() == nil {
		return
	}
	userId := resp.GetData().GetUserInfo().UserId
	_, err = iotredis.GetClient().HDel(context.Background(), iotconst.USERLASTLOGIN, strconv.Itoa(int(userId))).Result()
	if err != nil {
		return
	}
	//
	rpc.ClientOpenUserOnlineService.Delete(context.Background(), &protosService.OpenUserOnline{Token: token})
}

func GetDeveloperLoginInfo(ids []string) map[string]iotstruct.OpenUserLogin {
	//填充在线状态或登录地区
	mapData := make(map[string]iotstruct.OpenUserLogin)
	logins, err := iotredis.GetClient().HMGet(context.Background(), iotconst.USERLASTLOGIN, ids...).Result()
	if err != nil {
		return mapData
	}
	for _, v := range logins {
		if v == nil {
			continue
		}
		if str, ok := v.(string); ok && str != "" {
			var info iotstruct.OpenUserLogin
			if err := info.UnmarshalBinary([]byte(str)); err == nil && info.UserId > 0 {
				mapData[strconv.Itoa(int(info.UserId))] = info
			}
		}
	}
	return mapData
}

func ConcatStringSkipEmpty(a, b string) string {
	if b == "" {
		return a
	}
	if a != "" {
		return a + "," + b
	} else {
		return b
	}
}

func GetAddrByIp(ipStr string) string {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return ipStr
	}
	if ip.IsLoopback() { //过滤掉本地地址
		return "localhost"
	}
	if ip.IsPrivate() { //过滤掉本网段局域网地址
		return ipStr
	}
	//geo, err := iotutil.Geoip(ipStr, config.Global.IpService.QueryUrl, config.Global.IpService.AppCode)
	geo, err := controls.Geoip(ipStr)
	if err != nil {
		return ipStr
	}
	var addr string
	addr = ConcatStringSkipEmpty(addr, geo.Country)
	addr = ConcatStringSkipEmpty(addr, geo.Province)
	addr = ConcatStringSkipEmpty(addr, geo.City)
	addr = ConcatStringSkipEmpty(addr, geo.District)
	return addr
}
