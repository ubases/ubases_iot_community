// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTLangTranslate = "t_lang_translate"

// TLangTranslate mapped from table <t_lang_translate>
type TLangTranslate struct {
	Id           int64     `gorm:"column:id;primaryKey" json:"id"`                        // 主键ID
	SourceTable  string    `gorm:"column:source_table" json:"sourceTable"`                // 来源的表
	SourceRowId  string    `gorm:"column:source_row_id" json:"sourceRowId"`               // 来源的行id
	Lang         string    `gorm:"column:lang" json:"lang"`                               // 语言分类
	FieldName    string    `gorm:"column:field_name" json:"fieldName"`                    // 字段名
	FieldType    int32     `gorm:"column:field_type" json:"fieldType"`                    // 字段值的类型 string,properties
	FieldValue   string    `gorm:"column:field_value" json:"fieldValue"`                  // 字段值
	CreatedBy    int64     `gorm:"column:created_by;not null;default:0" json:"createdBy"` // 创建人
	UpdatedBy    int64     `gorm:"column:updated_by;not null;default:0" json:"updatedBy"` // 修改人
	CreatedAt    time.Time `gorm:"column:created_at" json:"createdAt"`                    // 创建时间
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updatedAt"`                    // 修改时间
	PlatformType int32     `gorm:"column:platform_type" json:"platformType"`              // 平台类型 =1 云管平台 =2 开放平台
	TenantId     string    `gorm:"column:tenant_id" json:"tenantId"`                      // 租户Id（开放平台数据需要绑定）
}

// TableName TLangTranslate's table name
func (*TLangTranslate) TableName() string {
	return TableNameTLangTranslate
}
