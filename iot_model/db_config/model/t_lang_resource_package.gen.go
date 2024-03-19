// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTLangResourcePackage = "t_lang_resource_package"

// TLangResourcePackage mapped from table <t_lang_resource_package>
type TLangResourcePackage struct {
	Id              int64     `gorm:"column:id;primaryKey" json:"id"`                        // 主键ID
	BelongType      int32     `gorm:"column:belong_type" json:"belongType"`                  // 归属应用类型（=1平台 =2 开放平台 =3 APP）
	PackageName     string    `gorm:"column:package_name" json:"packageName"`                // 语言分类
	CreatedBy       int64     `gorm:"column:created_by;not null;default:0" json:"createdBy"` // 创建人
	UpdatedBy       int64     `gorm:"column:updated_by;not null;default:0" json:"updatedBy"` // 修改人
	CreatedAt       time.Time `gorm:"column:created_at" json:"createdAt"`                    // 创建时间
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"updatedAt"`                    // 修改时间
	BelongId        int64     `gorm:"column:belong_id" json:"belongId"`                      // 归属数据编号
	AppTemplateId   int64     `gorm:"column:app_template_id" json:"appTemplateId"`           // APP模板编号
	AppTemplateType int32     `gorm:"column:app_template_type" json:"appTemplateType"`       // APP模板类型
	FileSize        int64     `gorm:"column:file_size" json:"fileSize"`                      // 文件包大小
	FileName        string    `gorm:"column:file_name" json:"fileName"`                      // 文件包名称
}

// TableName TLangResourcePackage's table name
func (*TLangResourcePackage) TableName() string {
	return TableNameTLangResourcePackage
}
