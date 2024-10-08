// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTAppUserActiveDay = "t_app_user_active_day"

// TAppUserActiveDay mapped from table <t_app_user_active_day>
type TAppUserActiveDay struct {
	DataTime  time.Time `gorm:"column:data_time;primaryKey;default:CURRENT_TIMESTAMP" json:"dataTime"` // 月
	TenantId  string    `gorm:"column:tenant_id;primaryKey" json:"tenantId"`                           // 租户ID
	AppKey    string    `gorm:"column:app_key;primaryKey" json:"appKey"`                               // app key
	ActiveSum int64     `gorm:"column:active_sum" json:"activeSum"`                                    // 活跃用户数
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updatedAt"`
}

// TableName TAppUserActiveDay's table name
func (*TAppUserActiveDay) TableName() string {
	return TableNameTAppUserActiveDay
}
