// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTAppUserActive30day = "t_app_user_active_30day"

// TAppUserActive30day mapped from table <t_app_user_active_30day>
type TAppUserActive30day struct {
	DataTime  time.Time `gorm:"column:data_time;primaryKey;default:CURRENT_TIMESTAMP" json:"dataTime"` // 日期
	TenantId  string    `gorm:"column:tenant_id;primaryKey" json:"tenantId"`                           // 租户ID
	AppKey    string    `gorm:"column:app_key;primaryKey" json:"appKey"`                               // app key
	ActiveSum int64     `gorm:"column:active_sum" json:"activeSum"`                                    // 活跃用户数
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updatedAt"`
}

// TableName TAppUserActive30day's table name
func (*TAppUserActive30day) TableName() string {
	return TableNameTAppUserActive30day
}
