package entitys

type AppCancelAccount struct {
	Account  string `json:"account"` //登录信息
	Code     string `json:"code"`    //验证码
	TenantId string `json:"tenantId"`
	RegionId int64  `json:"regionId"`
	AppKey   string `json:"appKey"`
}

type SendCodeParam struct {
	AreaPhoneNumber string `json:"areaPhoneNumber"`
	Account         string `json:"account"` //手机号码
	//Type            int32  `json:"type"`    //验证码类型 1-注册  2-忘记密码  3-修改密码  4-注销账号  5-绑定第三方账号  6-绑定手机或邮箱 7-验证码登录
	AppKey   string `json:"appKey"` //App Key
	TenantId string `json:"tenantId"`
}
