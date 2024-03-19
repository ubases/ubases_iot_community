// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTOpenRole = "t_open_role"

// TOpenRole mapped from table <t_open_role>
type TOpenRole struct {
	Id        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	TenantId  string         `gorm:"column:tenant_id" json:"tenantId"`
	Status    int32          `gorm:"column:status;not null;default:0" json:"status"` // 状态;1: 正常; 2:禁用;
	ListOrder int32          `gorm:"column:list_order;default:0" json:"listOrder"`   // 排序
	IsDefault int32          `gorm:"column:is_default" json:"isDefault"`             // 是否默认角色(1 默认, 2 非默认)
	Name      string         `gorm:"column:name;not null;default:''" json:"name"`    // 角色名称
	Remark    string         `gorm:"column:remark;default:''" json:"remark"`         // 备注
	DataScope int32          `gorm:"column:data_scope;default:3" json:"dataScope"`   // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	IsAdmin   int32          `gorm:"column:is_admin" json:"isAdmin"`                 // 是否管理员角色(每一个空间都默认一个管理员角色,并且有所有菜单权限)[1 是, 2 否]
	CreatedAt time.Time      `gorm:"column:created_at" json:"createdAt"`
	CreatedBy int64          `gorm:"column:created_by" json:"createdBy"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	UpdatedBy int64          `gorm:"column:updated_by" json:"updatedBy"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`
}

// TableName TOpenRole's table name
func (*TOpenRole) TableName() string {
	return TableNameTOpenRole
}
