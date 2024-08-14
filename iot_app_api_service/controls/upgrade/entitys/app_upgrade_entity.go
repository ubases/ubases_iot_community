package entitys

// AppUpgradeVo mapped from table <t_app_upgrade>
type AppUpgradeVo struct {
	Id                  int64  `json:"id,omitempty"`        // 主键ID
	Name                string `json:"name"`                // APP名称
	Type                int64  `json:"type"`                // APP类型[1：安卓，2：iOS]
	Version             string `json:"version"`             // APP版本号
	UpgradeMode         int    `json:"upgradeMode"`         // 升级方式[1：弹框  2： 不弹框提醒]
	UpgradePrompt       string `json:"upgradePrompt"`       // 升级提示语
	PublishType         int64  `json:"publishType"`         // 发布类型[0：appstore]
	Url                 string `json:"url"`                 // APP路径
	AppstoreUrl         string `json:"appstoreUrl"`         // appstore地址
	MinimumRequired     string `json:"minimumRequired"`     // 最低兼容版本
	ForceUpgradeVersion string `json:"forceUpgradeVersion"` // 强制升级版本
	FileMd5             string `json:"fileMd5"`             // 文件MD5值
	AgreementRemind     int32  `json:"agreementRemind"`     //协议提醒标记 =1 弹框提醒
}

// AppQueryAppUpgradeForm query appUpgrade  form ;  if some field is required, create binding:"required" to tag by self
type AppQueryAppUpgradeForm struct {
	Type int32 `json:"type"` // APP类型
}
