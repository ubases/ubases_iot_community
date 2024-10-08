// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTCloudAppBuildAuth = "t_cloud_app_build_auth"

// TCloudAppBuildAuth mapped from table <t_cloud_app_build_auth>
type TCloudAppBuildAuth struct {
	Id        int64     `gorm:"column:id;primaryKey" json:"id"`            // 唯一Id
	Name      string    `gorm:"column:name;not null" json:"name"`          // 申请名称
	UserId    int64     `gorm:"column:user_id;not null" json:"userId"`     // 用户Id
	AuthKey   string    `gorm:"column:auth_key;not null" json:"authKey"`   // 授权Key
	TenantId  string    `gorm:"column:tenant_id;not null" json:"tenantId"` // 租户ID
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
	UpdateAt  time.Time `gorm:"column:update_at" json:"updateAt"`
}

// TableName TCloudAppBuildAuth's table name
func (*TCloudAppBuildAuth) TableName() string {
	return TableNameTCloudAppBuildAuth
}
