package iotstruct

type BuildInfo struct {
	ID        string `json:"id"`        //构建编号
	OS        string `json:"os"`        //操作系统
	ZipUrl    string `json:"zipUrl"`    //zip下载地址
	ZipMd5    string `json:"zipMd5"`    //zip MD5校验码
	NotifUrl  string `json:"notifyUrl"` //打包回调接口地址
	Time      int64  `json:"time"`      //时间戳
	BuildType int32  `json:"buildType"` //编译类型 = 2 APP面板编译
	Type      int32  `json:"type"`      //方法类型 = 1 构建  =2 取消构建
}
