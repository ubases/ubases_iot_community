package apis

import (
	"cloud_platform/iot_app_api_service/cached"
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/common/commonGlobal"
	_const "cloud_platform/iot_app_api_service/controls/user/const"
	"cloud_platform/iot_app_api_service/controls/user/entitys"
	"cloud_platform/iot_app_api_service/controls/user/services"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_model/db_app/model"
	"cloud_platform/iot_proto/protos/protosService"
	proto "cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

var Usercontroller UserController

type UserController struct {
} //用户操作控制器

var userServices = services.AppUserService{}

// @Summary 获取用户信息
// @Description 获取用户信息
// @Tags user
// @Accept application/json
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/user/detail [get]
func (UserController) GetUserDetail(c *gin.Context) {
	userid := controls.GetUserId(c)
	if userid == 0 {
		iotgin.ResBusinessP(c, "header.userId not found")
		return
	}
	data, code, err := userServices.SetContext(controls.WithUserContext(c)).GetUser(iotutil.ToString(userid))
	if code != 0 {
		iotgin.ResBusinessP(c, err.Error())
		return
	}
	iotgin.ResSuccess(c, data)
}

// Login 登录
// @Summary user
// @Description 登录
// @Tags user
// @Accept application/json
// @Param data body entitys.LoginInput true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/user/login [post]
func (UserController) Login(c *gin.Context) {
	bm := entitys.LoginInput{}
	if err := c.BindJSON(&bm); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	var (
		account     = bm.Account
		accountType = bm.Type
		password    = bm.Password
		loginType   = bm.LoginType
		appKey      = controls.GetAppKey(c)
		tenantId    = controls.GetTenantId(c)
		regionId    = controls.GetRegionInt(c)
		appPushId   = controls.GetAppPushId(c)
	)
	// 检查登录接口中的参数
	code, msg := userServices.CheckAuthParams(account, password, accountType)
	if code != 0 {
		iotgin.ResBusinessP(c, msg)
		return
	}
	// 登录检验手机和邮箱
	_, code, msg = userServices.AuthCheckPhoneAndEmail(account, bm.AreaPhoneNumber, accountType)
	if code != 0 {
		iotgin.ResBusinessP(c, msg)
		return
	}
	c.Set("Account", bm.Account)
	// 获取登录信息
	ctx := controls.WithUserContext(c)
	data, code, msg := userServices.SetContext(ctx).AppUserLogin(accountType, account, password, appKey, tenantId, regionId, bm.RegisterRegion, loginType, c.ClientIP())
	if code != 0 {
		iotgin.ResBusinessP(c, msg)
		return
	}
	//清理推送别名和注册新对的推送别名
	go clearUserAlias(ctx, data.UserId, appKey, tenantId, appPushId)
	iotgin.ResSuccess(c, data)
}

func clearUserAlias(ctx context.Context, userId, appKey, tenantId, appPushId string) {
	defer iotutil.PanicHandler()
	//检查别名并删除
	_, err := rpc.MessageService.ClearAlias(ctx, &proto.ClearAliasRequest{
		UserId:    userId,
		AppKey:    appKey,
		TenantId:  tenantId,
		AppPushId: appPushId,
	})
	if err != nil {
		iotlogger.LogHelper.Error("清理别名:%s", err.Error())
		return
	}
	iotlogger.LogHelper.Debugf("clearUserAlias end")
}

// Register 注册（邮箱&手机）
// @Summary 注册（邮箱&手机）
// @Description 注册（邮箱&手机）
// @Tags user
// @Accept application/json
// @Param data body entitys.Registerparam true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/user/register [post]
func (UserController) Register(c *gin.Context) {
	bm := entitys.Registerparam{}
	if err := c.BindJSON(&bm); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	phone := bm.Phone
	email := bm.Email
	password := bm.Password
	smsCode := bm.Code
	registerRegion := bm.RegisterRegion
	registerRegionId := controls.GetRegionInt(c)
	appKey := controls.GetAppKey(c)
	tenantId := controls.GetTenantId(c)
	thisContext := controls.WithUserContext(c)
	// 检查注册接口中的参数
	code, msg := userServices.CheckRegisterParams(password, smsCode, registerRegion, appKey, tenantId)
	if code != 0 {
		iotlogger.LogHelper.Errorf(msg)
		iotgin.ResBusinessP(c, msg)
		return
	}
	// 注册检验手机和邮箱
	account, accountParam, code, msg := userServices.SetContext(thisContext).RegisterCheckPhoneAndEmail(phone, email, bm.AreaPhoneNumber)
	if code != 0 {
		iotlogger.LogHelper.Errorf(msg)
		iotgin.ResBusinessP(c, msg)
		return
	}
	registerParams := entitys.UserRegister{
		AccountType:      accountParam,
		Account:          account,
		Password:         password,
		RegisterRegion:   registerRegion,
		RegisterRegionId: registerRegionId,
		Smscode:          smsCode,
		Ip:               c.ClientIP(),
		AppKey:           appKey,
		TenantId:         tenantId,
	}
	// 注册用户
	code, userId, msg := userServices.SetContext(thisContext).Register(registerParams)
	if code != 0 {
		iotlogger.LogHelper.Errorf(msg)
		iotgin.ResBusinessP(c, msg)
		return
	}
	c.Set("Account", account)
	iotgin.ResSuccess(c, userId)
}

// RegisterEx
// @Summary 注册（邮箱&手机，无密码）
// @Description 注册（邮箱&手机，无密码）
// @Tags user
// @Accept application/json
// @Param data body entitys.Registerparam true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/user/registerex [post]
func (UserController) RegisterEx(c *gin.Context) {
	bm := entitys.Registerparam{}
	if err := c.BindJSON(&bm); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	var (
		phone            = bm.Phone
		email            = bm.Email
		password         = bm.Password
		smsCode          = bm.Code
		registerRegion   = bm.RegisterRegion
		registerRegionId = controls.GetRegionInt(c)
		appKey           = controls.GetAppKey(c)
		tenantId         = controls.GetTenantId(c)
		appPushId        = controls.GetAppPushId(c)
		thisContext      = controls.WithUserContext(c)
	)
	// 检查注册接口中的参数
	code, msg := userServices.CheckRegisterParamsEx(password, smsCode, registerRegion, appKey, tenantId)
	if code != 0 {
		iotlogger.LogHelper.Errorf(msg)
		iotgin.ResBusinessP(c, msg)
		return
	}
	// 注册检验手机和邮箱
	account, accountParam, code, msg := userServices.SetContext(thisContext).RegisterCheckPhoneAndEmail(phone, email, bm.AreaPhoneNumber)
	if code != 0 {
		iotlogger.LogHelper.Errorf(msg)
		iotgin.ResBusinessP(c, msg)
		return
	}
	registerParams := entitys.UserRegister{
		AccountType: accountParam,
		Account:     account,
		//Password:         password,
		RegisterRegion:   registerRegion,
		RegisterRegionId: registerRegionId,
		Smscode:          smsCode,
		Ip:               c.ClientIP(),
		AppKey:           appKey,
		TenantId:         tenantId,
	}
	// 注册用户
	data, code, msg := userServices.SetContext(thisContext).RegisterEx(registerParams)
	if code != 0 {
		iotgin.ResBusinessP(c, msg)
		return
	}
	//清理推送别名和注册新对的推送别名
	go clearUserAlias(thisContext, data.UserId, appKey, tenantId, appPushId)
	iotgin.ResSuccess(c, data)
}

// @Summary 发送短信验证码
// @Description 发送短信验证码
// @Tags user
// @Accept application/json
// @Param data body entitys.SendSmsParam true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/user/sendSms [post]
func (UserController) SendSms(c *gin.Context) {
	bm := entitys.SendSmsParam{}
	if err := c.BindJSON(&bm); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	phone := bm.Phone
	if phone == "" {
		phone = bm.Account
	}
	smsType := bm.Type
	areaPhoneNumber := bm.AreaPhoneNumber
	if strings.TrimSpace(phone) == "" {
		iotgin.ResBusinessP(c, "手机号码为空")
		return
	}
	if strings.TrimSpace(areaPhoneNumber) == "" {
		iotgin.ResBusinessP(c, "区域手机号为空")
		return
	}
	if smsType == 0 {
		iotgin.ResBusinessP(c, "验证码类型有误")
		return
	}
	var (
		appKey   = controls.GetAppKey(c)
		lang     = controls.GetLang(c)
		tenantId = controls.GetTenantId(c)
	)
	if iotutil.CheckAllPhone(areaPhoneNumber, phone) == false {
		iotgin.ResBusinessP(c, "手机号码不合法")
		return
	}
	var phoneType int32
	if areaPhoneNumber == "86" {
		phoneType = 1
	} else if areaPhoneNumber == "1" {
		phoneType = 3
		phone = areaPhoneNumber + phone
	} else {
		iotgin.ResBusinessP(c, "手机号码不合法")
		return
	}
	// 发送验证码
	_, code, msg := userServices.SendSms(lang, phone, bm.Phone, phoneType, smsType, tenantId, appKey)
	if code != 0 {
		iotgin.ResBusinessP(c, msg)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// @Summary user
// @Description 发送邮件
// @Tags user
// @Accept application/json
// @Param data body entitys.SendEmailParam true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/user/sendEmail [post]
func (UserController) SendEmail(c *gin.Context) {
	bm := entitys.SendEmailParam{}
	if err := c.BindJSON(&bm); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	email := bm.Email
	emailType := bm.Type
	if strings.TrimSpace(email) == "" {
		iotgin.ResBusinessP(c, "邮箱为空")
		return
	}
	if emailType == 0 {
		iotgin.ResBusinessP(c, "邮件类型为空")
		return
	}
	var (
		lang     = controls.GetLang(c)
		appKey   = controls.GetAppKey(c)
		tenantId = controls.GetTenantId(c)
	)
	// 发送邮件
	_, code, msg := userServices.SendEmail(email, emailType, tenantId, appKey, lang)
	if code != 0 {
		iotgin.ResBusinessP(c, msg)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// @Summary 修改用户信息
// @Description 修改用户信息
// @Tags user
// @Accept application/json
// @Param data body entitys.UpdateUserParam true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/user/updateUser [post]
func (UserController) UpdateUser(c *gin.Context) {
	bm := entitys.UpdateUserParam{}
	if err := c.BindJSON(&bm); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	newpassword := bm.NewPassword
	smscode := bm.Code
	appKey := c.Request.Header.Get("appKey")
	var account string
	if strings.TrimSpace(newpassword) != "" || strings.TrimSpace(smscode) != "" {
		if strings.TrimSpace(newpassword) == "" {
			iotgin.ResBusinessP(c, "密码为空")
			return
		}
		if strings.TrimSpace(smscode) == "" {
			iotgin.ResBusinessP(c, "验证码为空")
			return
		}

		userId := controls.GetUserId(c)
		userInfo, err := userServices.SetContext(controls.WithUserContext(c)).GetUserById(userId)
		if err != nil {
			iotlogger.LogHelper.Errorf("用户不存在, Id: %v", userId)
			iotgin.ResBusinessP(c, "用户不存在")
			return
		}
		account = userInfo.UserName
		resp := iotredis.GetClient().Get(context.Background(), cached.APP+"_"+appKey+"_"+account+"_3")

		if resp.Val() != smscode {
			iotlogger.LogHelper.Errorf("验证码有误")
			iotgin.ResBusinessP(c, "验证码有误")
			return
		}
	}

	userId := controls.GetUserId(c)
	// 修改
	code, msg := userServices.UpdateUser(iotutil.ToString(userId), bm)
	if code != 0 {
		iotgin.ResBusinessP(c, msg)
		return
	}
	iotgin.ResSuccessMsg(c)
	if strings.TrimSpace(newpassword) != "" || strings.TrimSpace(smscode) != "" {
		//删除redis中验证码
		iotredis.GetClient().Del(context.Background(), cached.APP+"_"+appKey+"_"+account+"_3")

		//推送消息
		go services.SendUpdatePasswordMessage(services.SetAppInfo(c), iotutil.ToInt64(userId), account)
		//清除token
		controls.ClearTokenByUserId(iotutil.ToInt64(userId))
	}
	//设置图片状态
	if bm.Photo != "" {
		commonGlobal.SetAttachmentStatus(model.TableNameTUcUser, iotutil.ToString(userId), bm.Photo)
	}
}

// @Summary logout
// @Description user logout
// @Tags APP
// @Accept application/json
// @Param token header string true "token"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/user/logout [post]
func (UserController) Logout(c *gin.Context) {
	err := controls.Logout("用户主动调用 /user/logout 退出", c)
	if err != nil {
		iotlogger.LogHelper.Errorf("Logout.AppPushTokenUserService.Delete err:%s", err.Error())
	}
	iotgin.ResSuccessMsg(c)
	return
}

// @Summary 校验账号是否注册
// @Description 校验账号是否注册
// @Tags user
// @Accept application/json
// @Param data body entitys.ForgetPassword true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/user/checkAccount [post]
func (UserController) ForgetPassword(c *gin.Context) {
	bm := entitys.ForgetPassword{}
	if err := c.BindJSON(&bm); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	appKey := controls.GetAppKey(c)
	tenantId := controls.GetTenantId(c)
	regionServerId, _ := controls.RegionIdToServerId(iotutil.ToString(controls.GetRegionInt(c)))
	c.Set("Account", bm.Account)
	// 忘记密码
	data, code, msg := userServices.SetContext(controls.WithUserContext(c)).ForgetPassword(bm, appKey, tenantId, regionServerId)
	if code != 0 {
		iotgin.ResBusinessP(c, msg)
		return
	}
	iotgin.ResSuccess(c, data)
}

// @Summary 设置用户密码
// @Description 设置用户密码
// @Tags user
// @Accept application/json
// @Param data body entitys.SetPassword true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/user/setPassword [post]
func (UserController) SetUserPassword(c *gin.Context) {
	bm := entitys.SetPassword{}
	if err := c.BindJSON(&bm); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	newPassword := bm.NewPassword
	if strings.TrimSpace(newPassword) == "" {
		iotgin.ResBusinessP(c, "密码为空")
		return
	}
	userId := controls.GetUserId(c)
	// 修改
	code, msg := userServices.UpdateUserPassword(iotutil.ToString(userId), bm)
	if code != 0 {
		iotgin.ResBusinessP(c, msg)
		return
	}
	iotgin.ResSuccessMsg(c)
	//推送消息
	//go services.SendUpdatePasswordMessage(services.SetAppInfo(c), iotutil.ToInt64(userId), account)
	//清除token
	//controls.ClearTokenByUserId(iotutil.ToInt64(userId))
}

// @Summary 校验验证码
// @Description 校验验证码
// @Tags user
// @Accept application/json
// @Param data body entitys.CheckCodeParam true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/user/checkCode [post]
func (UserController) CheckCode(c *gin.Context) {
	bm := entitys.CheckCodeParam{}
	if err := c.BindJSON(&bm); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}

	appKey := c.Request.Header.Get("appKey")
	// 校验验证码
	code, msg := userServices.CheckCode(bm, appKey)
	if code != 0 {
		iotgin.ResBusinessP(c, msg)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// @Summary 校验账号是否注册
// @Description 校验账号是否注册
// @Tags user
// @Accept application/json
// @Param data body entitys.CheckAccount true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /user/checkAccount [post]
func (UserController) CheckAccount(c *gin.Context) {
	bm := entitys.CheckAccount{}
	if err := c.BindJSON(&bm); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	appKey := controls.GetAppKey(c)
	tenantId := controls.GetTenantId(c)
	regionServerId, _ := controls.RegionIdToServerId(iotutil.ToString(controls.GetRegionInt(c)))
	// 校验验证码
	code, msg := userServices.CheckAccount(bm, appKey, tenantId, regionServerId)
	if code != 0 {
		iotgin.ResFailCode(c, msg, code)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// @Router /user/convertMd5 [get]
func (UserController) ConvertToMd5(c *gin.Context) {
	code := c.Query("code")
	c.JSON(200, gin.H{"code": 200, "msg": "", "data": iotutil.EncodeMD5(code)})
}

// @Summary 根据ip获取位置信息
// @Description 根据ip获取位置信息
// @Tags user
// @Accept application/json
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/user/getAddress [get]
//func (UserController) GetAddress(c *gin.Context) {
//	geo, err := iotutil.Geoip(c.ClientIP(), config.Global.IpService.QueryUrl, config.Global.IpService.AppCode) //根据ip获取位置信息
//	if err != nil {
//		logger.Errorf("get address by ip[%s], error:%s", c.ClientIP(), err.Error())
//	}
//	iotgin.ResSuccess(c, geo)
//}

// @Summary user
// @Description 获取用户信息
// @Tags user
// @Accept application/json
// @Param data body entitys.UpdateUserParam true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Failure 400 object iotgin.ResponseModel 失败返回
// @Router /user/detail [get]
func (UserController) Detail(c *gin.Context) {
	bm := entitys.UpdateUserParam{}
	if err := c.BindJSON(&bm); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}

	//if  strings.TrimSpace(smscode) == "" {
	//iotgin.ResBusinessP(c, "验证码为空")
	//return
	//}

	newpassword := bm.NewPassword
	if strings.TrimSpace(newpassword) != "" {
		smscode := bm.Code
		if strings.TrimSpace(smscode) == "" {
			iotgin.ResBusinessP(c, "验证码为空")
			return
		}
		//从redis中校验验证码正确性
		iotgin.ResBusinessP(c, "新密码为空")
		return
	}

	//loginaccount,_:=c.Get("Account")
	//account := loginaccount.(string)

	//rds := util.Redispool.Get()
	//if  rds.Err() != nil {
	//	global.GVA_LOG.Error(fmt.Sprintf("redis连接出错, error:%s",  rds.Err().Error()))
	//	c.JSON(200, sys.ErrorCode.HSet(c, gin.H{"code": 101014, "msg": "验证码异常"}))
	//	return
	//}
	//defer rds.Close()
	//if err:=util.VerificationCodeInputRds(rds, redisKeys.GetUpdatePasswordCacheKey(account), smscode); err !=nil {
	//	global.GVA_LOG.Error(fmt.Sprintf("获取信息失败, params:%s,error:%s", util.GetJsonString(map[string]interface{}{"account": account}), err.Error()))
	//	c.JSON(200, sys.ErrorCode.HSet(c, gin.H{"code": 100001, "msg": "获取信息失败"}))
	//	return
	//}

	userId := controls.GetUserId(c)
	//uidhex := bson.ObjectIdHex(userid.(string))
	// 修改密码
	code, msg := userServices.UpdateUser(iotutil.ToString(userId), bm)
	if code != 0 {
		iotgin.ResBusinessP(c, msg)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// @Summary 查询用户家庭列表
// @Description 查询用户家庭列表
// @Tags user
// @Accept application/json
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/user/homeList [get]
func (UserController) HomeList(c *gin.Context) {
	userId := controls.GetUserId(c)
	// 用户家庭列表
	data, code, msg := userServices.SetContext(controls.WithUserContext(c)).HomeList(iotutil.ToInt64(userId))
	if code != 0 {
		iotgin.ResBusinessP(c, msg)
		return
	}
	iotgin.ResSuccess(c, data)
}

// @Summary 第三方登录
// @Description 第三方登录
// @Tags APP
// @Accept application/json
// @Param data body entitys.ChannelAuth true "请求参数结构体"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/user/channelAuth [post]
func (UserController) ChannelAuth(c *gin.Context) {
	bm := entitys.ChannelAuth{}
	if err := c.BindJSON(&bm); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	var (
		appKey            = controls.GetAppKey(c)
		tenantId          = controls.GetTenantId(c)
		regionId          = controls.GetRegionInt(c)
		appPushId         = controls.GetAppPushId(c)
		thisContext       = controls.WithUserContext(c)
		sysInfo           = controls.GetSystemInfo(c)
		authorizationCode = bm.Code
		channelType       = bm.Type
		channelId         = bm.ChannelId
		nickname          = bm.Nickname
		result            *entitys.LoginUserRes
		resultMsg         string
	)
	if channelType == 0 {
		iotgin.ResBadRequest(c, "type")
		return
	}

	//区域Id转区域服务器Id
	regionServerId, err := controls.RegionIdToServerId(iotutil.ToString(regionId))
	if err != nil {
		iotgin.ResBadRequest(c, "region")
		return
	}

	switch channelType {
	case _const.Wechat:
		data, msgcode, msg := userServices.SetContext(thisContext).WechatLogin(authorizationCode, appKey, tenantId, c.ClientIP(), regionServerId)
		result = data
		resultMsg = msg
		if msgcode != 0 && msgcode != 200 {
			iotgin.ResBusinessP(c, msg)
			return
		}
	case _const.WechatApplet:
		data, msgcode, msg := userServices.SetContext(thisContext).MinProgramLogin(authorizationCode, appKey, tenantId, c.ClientIP(), regionServerId, sysInfo)
		result = data
		resultMsg = msg
		if msgcode != 0 && msgcode != 200 {
			iotgin.ResBusinessP(c, msg)
			return
		}
	case _const.Appleid, _const.Google:
		data, msgcode, msg := userServices.SetContext(thisContext).AppleidLogin(channelId, nickname, c.ClientIP(), appKey, tenantId, channelType, regionServerId)
		result = data
		resultMsg = msg
		if msgcode != 0 && msgcode != 200 {
			iotgin.ResBusinessP(c, msg)
			return
		}
	default:
		iotgin.ResBadRequest(c, "type")
		return
	}
	//清理推送别名和注册新对的推送别名
	go clearUserAlias(thisContext, result.UserId, appKey, tenantId, appPushId)
	iotgin.ResSuccessDataAndMsg(c, result, resultMsg)
}

// MinprogramChannelAuth 微信小程序登录
// @Summary 微信小程序登录
// @Description 微信小程序登录
// @Tags APP
// @Accept application/json
// @Param data body entitys.ChannelAuth true "请求参数结构体"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/wechat/minprogram/login [post]
func (UserController) MinprogramChannelAuth(c *gin.Context) {
	bm := entitys.ChannelAuth{}
	if err := c.BindJSON(&bm); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	var (
		authorizationCode = bm.Code
		appKey            = controls.GetAppKey(c)
		tenantId          = controls.GetTenantId(c)
		regionId          = controls.GetRegionInt(c)
		appPushId         = controls.GetAppPushId(c)
		thisContext       = controls.WithUserContext(c)
		sysInfo           = controls.GetSystemInfo(c)
	)

	//区域Id转区域服务器Id
	regionServerId, err := controls.RegionIdToServerId(iotutil.ToString(regionId))
	if err != nil {
		iotgin.ResBadRequest(c, "region")
		return
	}

	result, msgcode, msg := userServices.SetContext(thisContext).MinProgramLogin(authorizationCode, appKey, tenantId, c.ClientIP(), regionServerId, sysInfo)
	resultMsg := msg
	if msgcode != 0 && msgcode != 200 {
		iotgin.ResBusinessP(c, msg)
		return
	}
	//清理推送别名和注册新对的推送别名
	go clearUserAlias(thisContext, result.UserId, appKey, tenantId, appPushId)
	iotgin.ResSuccessDataAndMsg(c, result, resultMsg)
}

// @Summary 第三方渠道注册
// @Description 第三方渠道注册
// @Tags APP
// @Accept application/json
// @Param data body entitys.ChannelBind true "请求参数结构体"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/user/channelBind [post]
func (UserController) ChannelBind(c *gin.Context) {
	bm := entitys.ChannelBind{}
	if err := c.BindJSON(&bm); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	smsCode := bm.Code
	account := bm.Account
	bindType := bm.BindType
	channelId := bm.ChannelId
	channelType := bm.Type
	code, msg := userServices.CheckChannelBindParams(account, smsCode, bindType, channelType, channelId)
	if code != 0 {
		iotgin.ResBusinessP(c, msg)
		return
	}
	// 登录检验手机和邮箱
	_, code, msg = userServices.AuthCheckPhoneAndEmail(account, bm.AreaPhoneNumber, bindType)
	if code != 0 {
		iotgin.ResBusinessP(c, msg)
		return
	}
	bm.RegisterRegionId = controls.GetRegionInt(c)
	bm.AppKey = controls.GetAppKey(c)
	bm.TenantId = controls.GetTenantId(c)

	data, msgcode, msg := userServices.SetContext(controls.WithUserContext(c)).ChannelBind(c.ClientIP(), bm)
	if msgcode != 0 && msgcode != 200 {
		iotgin.ResBusinessP(c, msg)
		return
	}

	iotgin.ResSuccess(c, data)
}

// @Summary 增加第三方渠道账号绑定
// @Description 增加第三方渠道账号绑定
// @Tags user
// @Accept application/json
// @Param data body entitys.AddChannelBind true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/user/addChannelBind [post]
func (UserController) AddChannelBind(c *gin.Context) {
	bm := entitys.AddChannelBind{}
	if err := c.BindJSON(&bm); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	if bm.Code == "" {
		iotgin.ResBadRequest(c, "code")
		return
	}
	channelType, _ := iotutil.ToInt32ErrNew(bm.Type)
	if channelType == 0 {
		iotgin.ResBadRequest(c, "type")
		return
	}
	userId := controls.GetUserId(c)
	appKey := controls.GetAppKey(c)
	tenantId := controls.GetTenantId(c)
	regionId := controls.GetRegionInt(c)
	nickName := controls.GetNickName(c)
	data, msgcode, msg := userServices.SetContext(controls.WithUserContext(c)).AddChannelBind(bm, iotutil.ToInt64(userId), appKey, tenantId, regionId, nickName)
	if msgcode != 0 && msgcode != 200 {
		iotgin.ResFailCode(c, msg, msgcode)
		return
	}

	iotgin.ResSuccess(c, data)
}

// @Summary 第三方渠道账号解绑
// @Description 第三方渠道账号解绑
// @Tags user
// @Accept application/json
// @Param data body entitys.UnbundlingChannel true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/user/unbindChannel [post]
func (UserController) UnbindChannel(c *gin.Context) {
	bm := entitys.UnbundlingChannel{}
	if err := c.BindJSON(&bm); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	if bm.ChannelId == "" {
		iotgin.ResBadRequest(c, "channelId")
		return
	}
	channelType, _ := iotutil.ToInt32ErrNew(bm.Type)
	if channelType == 0 {
		iotgin.ResBadRequest(c, "type")
		return
	}
	userId := controls.GetUserId(c)
	appKey := controls.GetAppKey(c)
	tenantId := controls.GetTenantId(c)
	regionServerId, _ := controls.RegionIdToServerId(iotutil.ToString(controls.GetRegionInt(c)))
	msgcode, msg := userServices.SetContext(controls.WithUserContext(c)).UnbindChannel(bm, iotutil.ToInt64(userId), appKey, tenantId, regionServerId)
	if msgcode != 0 {
		iotgin.ResBusinessP(c, msg)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// @Summary 登录账号信息绑定
// @Description 登录账号信息绑定
// @Tags user
// @Accept application/json
// @Param data body entitys.AccountBind true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/user/accountBind [post]
func (UserController) AccountBind(c *gin.Context) {
	bm := entitys.AccountBind{}
	if err := c.BindJSON(&bm); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	if strings.TrimSpace(bm.Account) == "" {
		iotgin.ResBusinessP(c, "account")
		return
	}
	if strings.TrimSpace(bm.Code) == "" {
		iotgin.ResBusinessP(c, "code")
		return
	}
	accountType := bm.Type
	if accountType == 0 {
		iotgin.ResBusinessP(c, "type")
		return
	}

	userId := controls.GetUserId(c)
	appKey := controls.GetAppKey(c)
	tenantId := controls.GetTenantId(c)
	regionId := controls.GetRegionInt(c)
	code, msg := userServices.AccountBind(iotutil.ToInt64(userId), bm, appKey, tenantId, regionId)
	if code != 0 {
		iotgin.ResFailCode(c, msg, code)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// @Summary APP注销账号
// @Description APP注销账号
// @Tags user
// @Accept application/json
// @Param data body string true "请求参数结构体 {"account": "", "code":""}"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/user/cancelAccount [post]
func (UserController) CancelAccount(c *gin.Context) {
	bm := entitys.CancelAccount{}
	if err := c.BindJSON(&bm); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	if strings.TrimSpace(bm.Account) == "" || strings.TrimSpace(bm.Code) == "" {
		iotgin.ResErrParams(c)
		return
	}
	userId := controls.GetUserId(c)
	appKey := controls.GetAppKey(c)
	tenantId := controls.GetTenantId(c)
	regionServerId := controls.GetRegionInt(c)
	code, msg := userServices.SetContext(controls.WithUserContext(c)).CancelAccount(true, iotutil.ToInt64(userId), bm, appKey, tenantId, regionServerId, c.ClientIP())
	if code != 0 {
		iotgin.ResFailCode(c, msg, code)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// @Summary 取消第三方验证
// @Description 取消第三方验证
// @Tags user
// @Accept application/json
// @Param data body string true "请求参数结构体 {"account": "", "code":""}"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/user/cancelThirdPartyAccount [post]
func (UserController) CancelThirdPartyAccount(c *gin.Context) {
	bm := entitys.CancelAccount{}
	if err := c.BindJSON(&bm); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	userId := controls.GetUserId(c)
	appKey := controls.GetAppKey(c)
	tenantId := controls.GetTenantId(c)
	regionServerId := controls.GetRegionInt(c)
	code, msg := userServices.SetContext(controls.WithUserContext(c)).CancelAccount(false, iotutil.ToInt64(userId), bm, appKey, tenantId, regionServerId, c.ClientIP())
	if code != 0 {
		iotgin.ResFailCode(c, msg, code)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// RefreshToken 刷新token
// @Summary 刷新token
// @Description
// @Tags 刷新token
// @Accept application/json
// @Param data body entitys.RefreshToken true "请求参数结构体"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/user/refreshToken [post]
func (UserController) RefreshToken(c *gin.Context) {
	var req entitys.RefreshToken
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.RefreshToken == "" {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	appKey := c.Request.Header.Get("appKey")
	tenantId := c.Request.Header.Get("tenantId")
	token, refreshToken, expiresAt, err := userServices.RefreshToken(req, appKey, tenantId)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	var res = map[string]interface{}{
		"token":        token,
		"refreshToken": refreshToken,
		"expiresAt":    expiresAt,
	}
	iotgin.ResSuccess(c, res)
}

// @Summary 获取appId
// @Description 获取appId
// @Tags APP
// @Accept application/json
// @Param data body entitys.GetAppId true "请求参数结构体"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/user/getAppId [post]
func (UserController) GetAppId(c *gin.Context) {
	bm := entitys.GetAppId{}
	if err := c.BindJSON(&bm); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	channelType := bm.Type

	userId := controls.GetUserId(c)
	appKey := c.GetHeader("appKey")

	msgcode, _, result := userServices.SetContext(controls.WithUserContext(c)).GetAppId(channelType, iotutil.ToInt64(userId), appKey)
	if msgcode != 0 {
		iotgin.ResSuccess(c, result)
		//iotgin.ResBusinessP(c, msg)
		return
	}
	iotgin.ResSuccess(c, result)
}

// @Summary 获取语音配置信息
// @Description 获取语音配置信息
// @Tags user
// @Accept application/json
// @Param appKey header string true "APPKEY"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/user/getFunctionConfig [get]
func (UserController) GetFunctionConfigVoice(c *gin.Context) {
	userId := controls.GetUserId(c)
	appKey := c.GetHeader("appKey")

	msgcode, _, result := userServices.SetContext(controls.WithUserContext(c)).GetFunctionConfigVoice(iotutil.ToInt64(userId), appKey)
	if msgcode != 0 {
		iotgin.ResSuccess(c, result)
		//iotgin.ResBusinessP(c, msg)
		return
	}
	iotgin.ResSuccess(c, result)
}

// GetFunctionConfig APP功能配置信息获取
// @Summary APP功能配置信息获取
// @Description 工单信息表获取详情
// @Tags 用户相关接口
// @Param appKey header string true "appKey"
// @Success 0 {object} iotgin.ResponseModel "{"code":0,"msg":"ok","data":{"aboutUs":1,"eula":0,"privacyPolicy":0,"weather":1,"geo":2}}"
// @Router /v1/platform/app/functionConfig [get]
func (UserController) GetFunctionConfig(c *gin.Context) {
	appKey := c.GetHeader("appKey")

	resp, err := userServices.SetContext(controls.WithUserContext(c)).GetFunctionConfig(appKey)
	if err != nil {
		iotgin.ResBadRequest(c, "获取功能配置失败")
		return
	}
	iotgin.ResSuccess(c, resp)
}

// @Summary 切换语言，清理缓存
// @Description
// @Tags APP
// @Accept application/json
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/changeLang [get]
func (UserController) ChangeLang(c *gin.Context) {
	userId := controls.GetUserId(c)
	//用户Id不为空的情况需要清理家庭缓存
	if userId != 0 {
		services.ClearHomeCached(userId, true)
	}
	iotgin.ResSuccessMsg(c)
}

func getAppInfo(ctx context.Context, appKey string) (*protosService.OemApp, error) {
	res, err := rpc.ClientOemAppService.Find(ctx, &protosService.OemAppFilter{
		AppKey: appKey,
	})
	if err != nil && err.Error() != "record not found" {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	if res.Data == nil || len(res.Data) <= 0 {
		return nil, errors.New("参数错误,未找到记录")
	}
	return res.Data[0], nil
}

// @Description APP推送服务注册接口
// @Tags APP
// @Accept application/json
// @Param data body entitys.AppPushRegister true "请求参数结构体"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/push/register [post]
func (UserController) PushRegister(c *gin.Context) {
	bm := entitys.AppPushRegister{}
	if err := c.BindJSON(&bm); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	if bm.AppPushId == "" || bm.AppPushId == "undefined" {
		iotgin.ResBadRequest(c, "appPushId不能为空或者undefined")
		return
	}
	appKey := controls.GetAppKey(c)
	appInfo, err := getAppInfo(c, appKey)
	if err != nil {
		iotgin.ResBadRequest(c, "appKey")
		return
	}
	//APP的包名获取
	appPacketName := ""
	switch bm.Platform {
	case "ios":
		appPacketName = appInfo.IosPkgName
	default:
		appPacketName = appInfo.AndroidPkgName
	}
	userId := controls.GetUserId(c)
	tenantId := controls.GetTenantId(c)
	regionId := controls.GetRegionInt(c)
	lang := controls.GetLang(c)
	//将AppPushRegister保存到iot_message.t_app_push_token
	_, err = rpc.AppPushTokenService.Create(context.Background(), &proto.AppPushToken{
		AppToken:        bm.AppPushToken,
		AppPushId:       bm.AppPushId,
		AppPushPlatform: iotutil.ToString(bm.Platform),
		AppKey:          appKey,
		AppPacketName:   appPacketName,
		UserId:          userId,
		TenantId:        tenantId,
		RegionId:        regionId,
		Lang:            lang,
	})
	if err != nil {
		iotgin.ResBadRequest(c, "注册推送信息失败")
		return
	}
	iotgin.ResSuccessMsg(c)
}

// Subscribe 会员订阅提交
// @Summary 会员订阅提交
// @Description 会员订阅提交
// @Tags user
// @Accept application/json
// @Param data body entitys.UserSubscribeParam true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/user/subscribe [post]
func (UserController) Subscribe(c *gin.Context) {
	bm := entitys.UserSubscribeParam{}
	if err := c.BindJSON(&bm); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	// userId := controls.GetUserId(c)
	// 修改
	//code, msg := userServices.Subscribe(iotutil.ToString(userId), bm)
	//if code != 0 {
	//	iotgin.ResBusinessP(c, msg)
	//	return
	//}
	iotgin.ResSuccessMsg(c)
}
