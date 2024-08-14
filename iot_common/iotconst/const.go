package iotconst

var (
	Empty = ""
	NoticeMessageType = "notice" //通知消息类型
	HomeMessageType   = "home"   //家庭消息类型
	AlarmMessageType  = "alarm"  //告警消息类型

	TypeIsPhone = 1 //手机
	TypeIsEmail = 2 //邮箱

	//AccountNormal        = 1 //正常
	//AccountWaitCancel    = 2 //待注销
	//AccountAlreadyCancel = 3 //已注销

	SevenDaysSeconds = 7 * 24 * 60 * 60 //7天时间的秒数

	GenderMan   = 1 // 男
	GenderWoman = 2 // 女
	GenderOther = 3 // 其它

	DefaultHomeName = "我的家庭"

	DefaultPhoto = "default_avatar.jpg"

	Wechat       int32 = 1 //微信
	Alipay       int32 = 2 //支付宝
	Weibo        int32 = 3 //微博
	Facebook     int32 = 4 //facebook
	Appleid      int32 = 5 //appleid
	WechatMiniProgram int32 = 6 //微信小程序
	Google       int32 = 7 //谷歌
	Guest        int32 = 8 //游客

	RoleSuperAdministrator = 1 //超级管理员
	RoleAdministrator      = 2 //管理员
	RoleMember             = 3 //成员

	NormalHome = 1 //正常家庭
	ShareHome  = 2 //分享家庭

	TokenUserKey = "%s_tokens"
)
