// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTDemoDataAuth = "t_demo_data_auth"

// TDemoDataAuth mapped from table <t_demo_data_auth>
type TDemoDataAuth struct {
	Id        int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Title     string         `gorm:"column:title;default:''" json:"title"`         // 标题
	CreatedBy int32          `gorm:"column:created_by;default:0" json:"createdBy"` // 创建人
	UpdatedBy int32          `gorm:"column:updated_by;default:0" json:"updatedBy"` // 修改人
	CreatedAt time.Time      `gorm:"column:created_at" json:"createdAt"`           // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updatedAt"`           // 修改时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`           // 删除时间
}

// TableName TDemoDataAuth's table name
func (*TDemoDataAuth) TableName() string {
	return TableNameTDemoDataAuth
}
