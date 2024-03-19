package entitys

import (
	"cloud_platform/iot_proto/protos/protosService"
	"github.com/golang-jwt/jwt"
)

type UpdateUserParam struct {
	NewPassword   string `json:"newPassword"`   //新密码
	Code          string `json:"code"`          //短信验证码
	Nickname      string `json:"nickname"`      //昵称
	DefaultHomeId string `json:"defaultHomeId"` //默认家庭id
	Photo         string `json:"photo"`         //用户头像
	Gender        int32  `json:"gender"`        //性别
	Birthday      string `json:"birthday"`      //生日
	Country       string `json:"country"`       //国家
	Province      string `json:"province"`      //省份
	City          string `json:"city"`          //城市
}

type Registerparam struct {
	RegisterRegion  string `json:"registerRegion"`  //注册地区
	Phone           string `json:"phone"`           //手机号码  手机号码和邮箱必传其一
	Password        string `json:"password"`        //密码
	Code            string `json:"code"`            //短信验证码
	Email           string `json:"email"`           //邮箱  手机号码和邮箱必传其一
	AreaPhoneNumber string `json:"areaPhoneNumber"` //手机区号
	AppPushRegister
}

type LoginInput struct {
	Password        string `json:"password"`                 //密码
	Account         string `json:"account"`                  //用户信息
	Type            int32  `json:"type"`                     //用户类型 1-手机,2-邮箱
	AreaPhoneNumber string `json:"areaPhoneNumber"`          //手机区号
	LoginType       int32  `json:"loginType,omitempty"`      //登录方式 0-密码登录 1-验证码登录
	RegisterRegion  string `json:"registerRegion,omitempty"` //注册地区
	AppPushRegister        //APP推送注册参数
}

// GetStartLoginInput 免登输入参数
type GetStartLoginInput struct {
	VisitId         string `json:"visitId"`                  //密码
	RegisterRegion  string `json:"registerRegion,omitempty"` //注册地区
	AppPushRegister        //APP推送注册参数
}

type SendSmsParam struct {
	AreaPhoneNumber string `json:"areaPhoneNumber"`
	Phone           string `json:"phone"`   //手机号码
	Account         string `json:"account"` //手机号码
	Type            int32  `json:"type"`    //验证码类型 1-注册  2-忘记密码  3-修改密码  4-注销账号  5-绑定第三方账号  6-绑定手机或邮箱 7-验证码登录
}

type SendEmailParam struct {
	Email string `json:"email"` //邮箱
	Type  int32  `json:"type"`  //验证码类型 1-注册  2-忘记密码  3-修改密码  4-注销账号  5-绑定第三方账号  6-绑定手机或邮箱
}

type CheckCodeParam struct {
	Account         string `json:"account"`         //用户信息
	AccountType     int32  `json:"accountType"`     // 1-手机  2-邮箱
	AreaPhoneNumber string `json:"areaPhoneNumber"` //手机区号
	Type            int32  `json:"type"`            //验证码类型 1-注册  2-忘记密码  3-修改密码  4-注销账号  5-绑定第三方账号  6-绑定手机或邮箱 7-验证码登录
	Code            string `json:"code"`            //短信验证码
}

type ForgetPassword struct {
	Account         string `json:"account"`         //登录信息
	NewPassword     string `json:"newPassword"`     //新密码
	Code            string `json:"code"`            //短信验证码
	Type            int32  `json:"type"`            //登录类型 1-手机,2-邮箱
	AreaPhoneNumber string `json:"areaPhoneNumber"` //手机区号
}

type SetPassword struct {
	Account     string `json:"account"`     //登录信息
	NewPassword string `json:"newPassword"` //新密码
}

type QueryUser struct {
	Username string `json:"username" binding:"required" ` //用户名
}

// 用户返回数据结构
type UserInfoDto struct {
	//Id            bson.ObjectId    `json:"_id" bson:"_id"` //用户唯一编号
	//Appleid       AppleIdLoginInfo `json:"appleid"`
	//Wechat        WechatLoginInfo  `json:"wechat"`
	City string `json:"city"`
	//Defaulthomeid bson.ObjectId    `json:"defaulthomeid"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
	//Homelist      []bson.ObjectId  `json:"homelist"`
	//Ip string `json:"ip"`
	Nickname        string  `json:"nickname"`
	Phone           string  `json:"phone"`
	Photo           string  `json:"photo"`
	Username        string  `json:"username"`
	Cancel_time     float64 `json:"cancel_time"`
	Accountstate    string  `json:"accountstate"`
	Openid          string  `json:"openid"`
	Unionid         string  `json:"unionid"`
	Wechatnickname  string  `json:"wechatnickname"`
	Appleiduserid   string  `json:"appleiduserid"`
	Appleidnickname string  `json:"appleidnickname"`
	Registerregion  string  `json:"registerregion"` //注册地区
}

// 用户信息响应数据结构
type LoginUserRes struct {
	//User         AppUserInfo `json:"user"`
	UserId   string `json:"userId,omitempty"`
	NickName string `json:"nickname,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Photo    string `json:"photo,omitempty"`
	//Status   int32  `json:"status"`
	//City          string `json:"city"`
	Gender        int32  `json:"gender"`
	Email         string `json:"email,omitempty"`
	DefaultHomeId string `json:"defaultHomeId,omitempty"`
	//RegisterRegion   string                   `json:"registerRegion"`
	UserName         string                 `json:"userName,omitempty"`
	SubmitCancelTime int64                  `json:"submitCancelTime,omitempty"` // 原注销提交时间，前端提示使用
	AccountCasser    bool                   `json:"accountCasser,omitempty"`
	ThirdPartyLogin  AppUserThirdPartyLogin `json:"thirdPartyLogin,omitempty"` //第三方登录
	Token            string                 `json:"token,omitempty"`
	RefreshToken     string                 `json:"refreshToken,omitempty"`
	ExpiresAt        int64                  `json:"expiresAt,omitempty"`
	LoginKey         string                 `json:"channelId,omitempty"`   //第三方渠道用户id
	ChannelName      string                 `json:"channelName,omitempty"` //第三方渠道昵称
	Password         string                 `json:"password,omitempty"`    //用户密码
	Birthday         string                 `json:"birthday"`              //生日
	Country          string                 `json:"country"`               //国家
	Province         string                 `json:"province"`              //省份
	City             string                 `json:"city"`                  //城市
	CountryName      string                 `json:"countryName"`           //国家名称
	ProvinceName     string                 `json:"provinceName"`          //省份名称
	CityName         string                 `json:"cityName"`              //城市名称
	//ChannelId        string                 `json:"channelId"`             //第三方渠道Id
	PasswordNotSet bool `json:"passwordNotSet"` //密码未设置
	ShowVconsole   bool `json:"showVconsole"`   //显示调试
	IsGuest        bool `json:"isGuest"`        //是否游客
}

func (s *LoginUserRes) SetShowVconsole(showVconsole bool) *LoginUserRes {
	s.ShowVconsole = showVconsole
	return s
}

func (s *LoginUserRes) SetIsGuest(isGuest bool) *LoginUserRes {
	s.IsGuest = isGuest
	return s
}

type AppUserInfo struct {
	UserId   string `json:"userId"`
	NickName string `json:"nickname"`
	Phone    string `json:"phone"`
	Photo    string `json:"photo"`
	//Status   int32  `json:"status"`
	//City          string `json:"city"`
	//Gender        int32  `json:"gender"`
	Email         string `json:"email"`
	DefaultHomeId string `json:"defaultHomeId"`
	//RegisterRegion   string                   `json:"registerRegion"`
	UserName         string                 `json:"userName"`
	SubmitCancelTime int64                  `json:"submitCancelTime"` // 原注销提交时间，前端提示使用
	AccountCasser    bool                   `json:"accountCasser"`
	ThirdPartyLogin  AppUserThirdPartyLogin `json:"thirdPartyLogin"` //第三方登录
}

type AppUserThirdPartyLogin struct {
	AppleId AppUserThirdPartyLoginInfo `json:"appleId,omitempty"`
	Wechat  AppUserThirdPartyLoginInfo `json:"wechat,omitempty"`
	Guest   AppUserThirdPartyLoginInfo `json:"guest,omitempty"`
}

func (s *AppUserThirdPartyLoginInfo) Set(third *protosService.UcUserThird) AppUserThirdPartyLoginInfo {
	s.Mode = third.ThirdType
	s.LoginKey = third.ThirdUserId
	s.Nickname = third.Nickname
	return *s
}

type AppUserThirdPartyLoginInfo struct {
	Mode     int32  `json:"type,omitempty"`        //登录方式 =1 微信 , =5 AppleId
	LoginKey string `json:"channelId,omitempty"`   //第三方渠道用户id
	Nickname string `json:"channelName,omitempty"` //第三方渠道昵称
}

type CheckAccount struct {
	Account         string `json:"account"`         //用户参数
	Type            int32  `json:"type"`            //类型 1-手机,2-邮箱
	AreaPhoneNumber string `json:"areaPhoneNumber"` //手机区号
}

type ChannelAuth struct {
	Code            string `json:"code"`        //授权Code  微信和支付宝渠道,传code
	ChannelId       string `json:"channelId"`   //第三方userid     AppleId渠道，传ChannelId
	Type            int32  `json:"type"`        //第三方类型 1-微信,5-AppleId
	Nickname        string `json:"channelName"` //第三方昵称  AppleId渠道，传Nickname
	AppPushRegister        //APP推送注册参数
}

type AccountBind struct {
	Code            string `json:"code"`            //验证码
	Account         string `json:"account"`         //登录参数
	Type            int32  `json:"type"`            //登录类型 1-手机,2-邮箱
	AreaPhoneNumber string `json:"areaPhoneNumber"` //手机区号
}

type CancelAccount struct {
	Account string `json:"account"` //登录信息
	Code    string `json:"code"`    //验证码
}

// 刷新tokne
type RefreshToken struct {
	RefreshToken string `json:"refreshToken"`
}

type ChannelBind struct {
	Code             string `json:"code"`             //验证码
	Account          string `json:"account"`          //登录参数
	BindType         int32  `json:"bindType"`         //登录类型 1-手机,2-邮箱 3-其它
	ChannelId        string `json:"channelId"`        //第三方平台用户唯一的id
	Type             int32  `json:"type"`             //渠道类型 1-微信,5-AppleId
	Password         string `json:"password"`         //密码
	Nickname         string `json:"channelName"`      //第三方昵称
	RegisterRegion   string `json:"registerRegion"`   //注册地区
	RegisterRegionId int64  `json:"registerRegionId"` //注册地区
	AreaPhoneNumber  string `json:"areaPhoneNumber"`  //手机区号
	AppKey           string `json:"appKey"`
	TenantId         string `json:"tenantId"`
}

type AppUserData struct {
	UserID   int64  `json:"userId"`
	Nickname string `json:"nickName"`
	Avatar   string `json:"avatar"`
	Account  string `json:"account"`
	//HomeIds  string `json:"homeIds"`
}

type AppClaims struct {
	jwt.StandardClaims
	AppUserData
}

type AddChannelBind struct {
	Code      string `json:"code"`        //授权Code  微信和支付宝渠道,传code
	Type      int32  `json:"type"`        //第三方类型 1-微信,5-AppleId
	ChannelId string `json:"channelId"`   //第三方userid    AppleId渠道，传ChannelId
	Nickname  string `json:"channelName"` //第三方昵称    AppleId渠道，传Nickname
}

type UnbundlingChannel struct {
	Type      int32  `json:"type"`      //第三方类型 1-微信,2-支付宝,3-微博,4-Facebook,5-AppleId
	ChannelId string `json:"channelId"` //第三方userid
}

type GetAppId struct {
	Type int32 `json:"type"` //第三方类型 1-微信,5-AppleId
}

type RegisterNewUser struct {
	Phone          string `json:"phone"`          //手机号码
	Password       string `json:"password"`       //密码
	Email          string `json:"email"`          //邮箱
	Ip             string `json:"ip"`             //IP
	AppKey         string `json:"appKey"`         //APPKey
	TenantId       string `json:"tenantId"`       //租户
	ThirdType      int32  `json:"thirdType"`      //第三方类型
	ThirdUserId    string `json:"thirdUserId"`    //第三方用户
	ThirdNickname  string `json:"thirdNickname"`  //第三方用户昵称
	RegionServerId int64  `json:"regionServerId"` //区域服务Id
}

type AppPushRegister struct {
	AppPushId    string      `json:"appPushId"`    //APP推送Id
	AppPushToken string      `json:"appPushToken"` //APP推送Token
	Platform     interface{} `json:"platform"`     //APP推送所属平台
}

// UserSubscribeParam 会员订阅
type UserSubscribeParam struct {
	payMethod int32  `json:"payMethod"` //支付方式
	useModel  int32  `json:"useModel"`  //使用机型 1 国内安卓、2 国外安卓、3 IOS系列
	priceType int32  `json:"priceType"` //价格类型
	priceId   string `json:"priceId"`   //价格Id
	price     int32  `json:"price"`     //价格验证
}
