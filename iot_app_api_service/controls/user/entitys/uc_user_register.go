package entitys

type UserRegister struct {
	AccountType      string `json:"accountType"`      //账号类型 =phone， =email
	Account          string `json:"account"`          //账号
	Password         string `json:"password"`         //密码
	RegisterRegion   string `json:"registerRegion"`   //注册区域名称
	RegisterRegionId int64  `json:"registerRegionId"` //注册区域Id
	Smscode          string `json:"smsCode"`          //短信验证码
	Ip               string `json:"ip"`               //Ip
	AppKey           string `json:"appKey"`           //APPKey
	TenantId         string `json:"tenantId"`         //租户Id
}
