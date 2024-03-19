package iotconst

// APP版本
var (
	OS_IOS           int32 = 1 //iOS上架包
	OS_ANDROID_CHINA int32 = 2 //Android国内版
	OS_ANDROID       int32 = 3 //Android海外版
)

// APP所属平台
var (
	APP_GOOGLE int32 = 1 //Google签名
	APP_HAIWEI int32 = 2 //华为签名
)

// APP支持语言
var (
	APP_SUPPORT_LANGUAGE = []string{"zh", "en"}
)

// APP面板类型
var (
	APP_PANEL_PUBLIC         int32 = 1 //公共面板
	APP_PANEL_CUSTOM_STUDIO  int32 = 2 //自定义设计器面板
	APP_PANEL_CUSTOM_DEVELOP int32 = 3 //自定义开发面板
)
