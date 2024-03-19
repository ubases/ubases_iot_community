// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTProductFaultType = "t_product_fault_type"

// TProductFaultType mapped from table <t_product_fault_type>
type TProductFaultType struct {
	Id         int64     `gorm:"column:id;primaryKey" json:"id"`                               // 主键ID
	ProductKey string    `gorm:"column:product_key;not null" json:"productKey"`                // 产品key
	Month      time.Time `gorm:"column:month;not null;default:CURRENT_TIMESTAMP" json:"month"` // 月份
	FaultType  string    `gorm:"column:fault_type;not null" json:"faultType"`                  // 故障类型
	Total      int64     `gorm:"column:total;default:0" json:"total"`                          // 月份故障类型总数
	UpdatedAt  time.Time `gorm:"column:updated_at;not null" json:"updatedAt"`                  // 修改时间
}

// TableName TProductFaultType's table name
func (*TProductFaultType) TableName() string {
	return TableNameTProductFaultType
}
