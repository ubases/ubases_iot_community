package model

import "time"

type AppLogUser struct {
	Id             int64     `gorm:"column:id;type:Int64;NOT NULL;COMMENT:'编号'" json:"id"`
	Account        string    `gorm:"column:account;type:String;NOT NULL;COMMENT:'用户账号,手机号或者邮箱'" json:"account"`
	AppKey         string    `gorm:"column:app_key;type:String;NOT NULL;COMMENT:'appKey'" json:"appKey"`
	TenantId       string    `gorm:"column:tenant_id;type:String;NOT NULL;COMMENT:'租户id'" json:"tenantId"`
	AppName        string    `gorm:"column:app_name;type:String;NOT NULL;COMMENT:'APP名称'" json:"appName"`
	Region         string    `gorm:"column:region;type:String;NOT NULL;COMMENT:'服务区'" json:"region"`
	RegionServerId int64     `gorm:"column:region_server_id;type:Int64;NOT NULL;COMMENT:'区域服务编号'" json:"regionServerId"`
	LoginTime      time.Time `gorm:"column:login_time;type:Datetime;NOT NULL;COMMENT:'最后登录时间'" json:"loginTime"`
	CreatedAt      time.Time `gorm:"column:created_at;type:Datetime;NOT NULL;COMMENT:'创建时间'" json:"createdAt"`
}

func (al *AppLogUser) TableName() string {
	return "t_iot_log_app_user"
}
