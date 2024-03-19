// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTOemAppTemplateUi = "t_oem_app_template_ui"

// TOemAppTemplateUi mapped from table <t_oem_app_template_ui>
type TOemAppTemplateUi struct {
	Id            int64     `gorm:"column:id;primaryKey" json:"id"`   // 主键唯一编号
	Type          int32     `gorm:"column:type;not null" json:"type"` // 选择功能类型（字典配置 app_template_ui_type)
	Name          string    `gorm:"column:name;not null" json:"name"` // 中文名称
	NameEn        string    `gorm:"column:name_en" json:"name_en"`    // 英文名称
	Code          string    `gorm:"column:code;not null" json:"code"` // 页面组件id
	Sort          int32     `gorm:"column:sort" json:"sort"`
	PageJson      string    `gorm:"column:page_json" json:"page_json"`
	AppTemplateId int64     `gorm:"column:app_template_id;not null" json:"app_template_id"`
	CreatedBy     int64     `gorm:"column:created_by" json:"created_by"` // 修改人
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"` // 创建时间
	UpdatedBy     int64     `gorm:"column:updated_by" json:"updated_by"` // 修改人
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"` // 修改时间
}

// TableName TOemAppTemplateUi's table name
func (*TOemAppTemplateUi) TableName() string {
	return TableNameTOemAppTemplateUi
}
