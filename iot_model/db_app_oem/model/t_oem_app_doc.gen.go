// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTOemAppDoc = "t_oem_app_doc"

// TOemAppDoc mapped from table <t_oem_app_doc>
type TOemAppDoc struct {
	Id              int64          `gorm:"column:id;primaryKey" json:"id"`
	Name            string         `gorm:"column:name" json:"name"`                             // 文档名称
	Apps            string         `gorm:"column:apps" json:"apps"`                             // 文档关联的app,  json字符串字段不做查询, 用于查看. 查询还是使用关联表查询
	Langs           string         `gorm:"column:langs" json:"langs"`                           // 文档对应的语种(json格式字符串)
	RemainLang      string         `gorm:"column:remain_lang" json:"remain_lang"`               // 兜底语种编码
	IsSucceedPubDoc int32          `gorm:"column:is_succeed_pub_doc" json:"is_succeed_pub_doc"` // 是否继承公版文档
	SucceedPubDoc   string         `gorm:"column:succeed_pub_doc" json:"succeed_pub_doc"`       // 继承公版的语种(json格式字符串)
	TenantId        string         `gorm:"column:tenant_id" json:"tenant_id"`                   // 租户id
	CreatedBy       int64          `gorm:"column:created_by" json:"created_by"`                 // 创建人
	CreatedAt       time.Time      `gorm:"column:created_at" json:"created_at"`                 // 创建时间
	UpdatedBy       int64          `gorm:"column:updated_by" json:"updated_by"`                 // 更新人
	UpdatedAt       time.Time      `gorm:"column:updated_at" json:"updated_at"`                 // 更新时间
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`                 // 删除时间
	HelpCenterName  string         `gorm:"column:help_center_name" json:"help_center_name"`     // 帮助中心名称
	Version         string         `gorm:"column:version" json:"version"`                       // 模板版本
}

// TableName TOemAppDoc's table name
func (*TOemAppDoc) TableName() string {
	return TableNameTOemAppDoc
}
