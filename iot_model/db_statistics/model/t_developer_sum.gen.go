// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTDeveloperSum = "t_developer_sum"

// TDeveloperSum mapped from table <t_developer_sum>
type TDeveloperSum struct {
	DeveloperSum int64     `gorm:"column:developer_sum" json:"developerSum"` // 开发者累计
	UpdatedAt    time.Time `gorm:"column:updated_at;not null" json:"updatedAt"`
}

// TableName TDeveloperSum's table name
func (*TDeveloperSum) TableName() string {
	return TableNameTDeveloperSum
}
