package iotconst

type AccountType int

var (
	PLATFROM_USER AccountType = 1 // 平台用户
	OPEN_USER     AccountType = 2 // 开发平台用户
)

type PlatformType string

var (
	PLATFORMTYPE_CLOUD PlatformType = "cloud" //云管平台
	PLATFORMTYPE_OPEN  PlatformType = "open"  //开放平台

	//开放平台账号
	OPEN_USER_MAIN_ACCOUNT int32 = 1 // 开发平台主账号
	OPEN_USER_SUB_ACCOUNT  int32 = 2 // 开放平台子账号

	///开放平台账号类型（=1 企业 =2 个人）
	OPEN_USER_ENTERPRISE_ACCOUNT int32 = 1
	OPEN_USER_PERSONAL_ACCOUNT   int32 = 2

	//性别
	SEX_UNKNOWN int32 = 0 // 未知
	SEX_MAN     int32 = 2 // 男
	SEX_WOMAN   int32 = 3 // 女

	//账号的状态 1:正常，2:待注销，3:已注销
	ACCOUNT_NORMAL    int32 = 1 // 正常
	ACCOUNT_CANCELING int32 = 2 // 待注销
	ACCOUNT_CANCELED  int32 = 3 // 已注销

	//第三方登录方式
	THIRD_PARTY_WECHAT  int32 = 1 // 微信登录
	THIRD_PARTY_APPLEID int32 = 2 // AppleId登录

	//第三方登录方式
	START_STATE   int32 = 1 // 启动状态
	UNSTART_STATE int32 = 2 // 未启动状态

	//家庭用户角色类型 1-家庭所有者，2-管理员，3-成员
	HOME_USER_ROLE_1 int32 = 1 //家庭所有者
	HOME_USER_ROLE_2 int32 = 2 //管理员
	HOME_USER_ROLE_3 int32 = 3 //成员

	//开放平台账号
	OPEN_USER_ACCOUNT_TYPE_REGISTER int32 = 1 // 个人注册
	OPEN_USER_ACCOUNT_TYPE_ADD      int32 = 2 // 平台管理员添加

)
