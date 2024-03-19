package model

import "time"

type AppLogRecords struct {
	Id             int64             `gorm:"column:id;type:Int64;NOT NULL;COMMENT:'编号'" json:"id"`
	Account        string            `gorm:"column:account;type:String;NOT NULL;COMMENT:'用户账号,手机号或者邮箱'" json:"account"`
	AppKey         string            `gorm:"column:app_key;type:String;NOT NULL;COMMENT:'appKey'" json:"appKey"`
	TenantId       string            `gorm:"column:tenant_id;type:String;NOT NULL;COMMENT:'租户id'" json:"tenantId"`
	RegionServerId int64             `gorm:"column:region_server_id;type:Int64;NOT NULL;COMMENT:'区域服务编号'" json:"regionServerId"`
	LogType        string            `gorm:"column:log_type;type:String;NOT NULL;COMMENT:'日志类型'" json:"logType"`
	EventName      string            `gorm:"column:event_name;type:String;NOT NULL;COMMENT:'事件名'" json:"eventName"`
	Details        map[string]string `gorm:"column:details;type:Map;NOT NULL;COMMENT:'详情'" json:"details"`
	CreatedAt      time.Time         `gorm:"column:created_at;type:Datetime;NOT NULL;COMMENT:'创建时间'" json:"createdAt"`
}

func (al *AppLogRecords) TableName() string {
	return "t_iot_log_app_records"
}
