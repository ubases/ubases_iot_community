// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTDataOverviewHour = "t_data_overview_hour"

// TDataOverviewHour mapped from table <t_data_overview_hour>
type TDataOverviewHour struct {
	DataTime             time.Time `gorm:"column:data_time;primaryKey;default:CURRENT_TIMESTAMP" json:"dataTime"` // 时间
	TenantId             string    `gorm:"column:tenant_id;primaryKey" json:"tenantId"`                           // 租户ID，为空表示所有
	DeviceActiveSum      int64     `gorm:"column:device_active_sum" json:"deviceActiveSum"`                       // 该月激活设备数
	DeviceFaultSum       int64     `gorm:"column:device_fault_sum" json:"deviceFaultSum"`                         // 该月设备故障数
	DeveloperRegisterSum int64     `gorm:"column:developer_register_sum" json:"developerRegisterSum"`             // 该月开发者注册数
	UserRegisterSum      int64     `gorm:"column:user_register_sum" json:"userRegisterSum"`                       // 该月APP用户注册数
	UpdatedAt            time.Time `gorm:"column:updated_at;not null" json:"updatedAt"`
}

// TableName TDataOverviewHour's table name
func (*TDataOverviewHour) TableName() string {
	return TableNameTDataOverviewHour
}
