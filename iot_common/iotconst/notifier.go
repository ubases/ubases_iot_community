package iotconst

//通知类型，对应iot_message表的tpl_type字段,1:验证码
type NotifierBusinesses int

var (
	NB_CODE     NotifierBusinesses = 1 //验证码
	NB_REGISTER NotifierBusinesses = 2 //注册
	NB_LOGGEDIN NotifierBusinesses = 3 //异地登录
)

// 验证码模板类型
var (
	NB_REGISTER_CODE                NotifierBusinesses = 1 //注册验证码
	NB_FORGET_PASSWORD_CODE         NotifierBusinesses = 2 //忘记密码验证码
	NB_CHANGE_PASSWORD_CODE         NotifierBusinesses = 2 //关闭密码验证码
	NB_CANCEL_ACCOUNT_CODE          NotifierBusinesses = 3 //注销账号验证码
	NB_BIND_THIRDPARTY_ACCOUNT_CODE NotifierBusinesses = 4 //绑定第三方账号验证码
	NB_BIND_PHONE_OR_EMAIL_CODE     NotifierBusinesses = 5 //绑定手机或邮箱验证码
	NB_LOGIN_CODE                   NotifierBusinesses = 7 //验证码登录
)
