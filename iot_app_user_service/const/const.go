package _const

var (
	NoticeMessageType = "notice" //通知消息类型
	HomeMessageType   = "home"   //家庭消息类型
	AlarmMessageType  = "alarm"  //告警消息类型

	TypeIsPhone = 1 //手机
	TypeIsEmail = 2 //邮箱

	AccountNormal        = 1 //正常
	AccountWaitCancel    = 2 //待注销
	AccountAlreadyCancel = 3 //已注销

	SevenDaysSeconds = 7 * 24 * 60 * 60 //7天时间的秒数

	GenderMan   = 1 // 男
	GenderWoman = 2 // 女
	GenderOther = 3 // 其它

	DefaultHomeName = "我的家庭"

	DefaultPhotoUrl = "http://qiniuoss.aithinker.com/avatar/c80351d7-8e9b-4759-a350-958eff48dd49.jpg?e=5255781625&token=_SRlsiDrTatwIIKLM84nINyCg0T25sA99B8GfTRF:QRqgG9dJeveBQYuGeP2UxhhowIU="

	Wechat  = 1 //微信
	Appleid = 5 //appleid

	RoleSuperAdministrator = 1 //超级管理员
	RoleAdministrator      = 2 //管理员
	RoleMember             = 3 //成员

	NormalHome = 1 //正常家庭
	ShareHome  = 2 //分享家庭
)
