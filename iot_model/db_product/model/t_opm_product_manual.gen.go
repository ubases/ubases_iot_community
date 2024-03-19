// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTOpmProductManual = "t_opm_product_manual"

// TOpmProductManual mapped from table <t_opm_product_manual>
type TOpmProductManual struct {
	Id         int64     `gorm:"column:id;primaryKey" json:"id"`                        // 主键
	TenantId   string    `gorm:"column:tenant_id;not null" json:"tenantId"`             // 租户id
	ProductKey string    `gorm:"column:product_key;not null" json:"productKey"`         // 产品唯一标识
	FileName   string    `gorm:"column:file_name;not null" json:"fileName"`             // 文件名称
	FileUrl    string    `gorm:"column:file_url;not null" json:"fileUrl"`               // 文件链接
	FileMd5    string    `gorm:"column:file_md5;not null" json:"fileMd5"`               // 文件md5
	CreatedBy  int64     `gorm:"column:created_by;not null;default:0" json:"createdBy"` // 创建人
	CreatedAt  time.Time `gorm:"column:created_at" json:"createdAt"`                    // 创建时间
	UpdatedBy  int64     `gorm:"column:updated_by;not null;default:0" json:"updatedBy"` // 修改人
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updatedAt"`                    // 修改时间
}

// TableName TOpmProductManual's table name
func (*TOpmProductManual) TableName() string {
	return TableNameTOpmProductManual
}
