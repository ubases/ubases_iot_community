// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTOemAppAssistRelease = "t_oem_app_assist_release"

// TOemAppAssistRelease mapped from table <t_oem_app_assist_release>
type TOemAppAssistRelease struct {
	Id                 int64          `gorm:"column:id;primaryKey" json:"id"`                          // 唯一主键
	DevelopId          int64          `gorm:"column:develop_id;not null" json:"develop_id"`            // 开发者编号
	TenantId           string         `gorm:"column:tenant_id" json:"tenant_id"`                       // 租户编号
	DevelopPhone       string         `gorm:"column:develop_phone" json:"develop_phone"`               // 开发者手机号码
	AppKey             string         `gorm:"column:app_key" json:"app_key"`                           // APP Key
	AppVersion         string         `gorm:"column:app_version" json:"app_version"`                   // APP的版本号
	AppTemplateId      int64          `gorm:"column:app_template_id" json:"app_template_id"`           // APP模板Id
	AppTemplateVersion string         `gorm:"column:app_template_version" json:"app_template_version"` // APP模板版本号
	SkinId             int64          `gorm:"column:skin_id" json:"skin_id"`                           // 皮肤编号
	StartTime          time.Time      `gorm:"column:start_time" json:"start_time"`                     // 有效开始时间
	EndTime            time.Time      `gorm:"column:end_time" json:"end_time"`                         // 有效结束时间
	Status             int32          `gorm:"column:status" json:"status"`                             // 状态（1=启用 2=禁用）
	Remark             string         `gorm:"column:remark" json:"remark"`                             // 备注
	CreatedBy          int64          `gorm:"column:created_by" json:"created_by"`                     // 修改人
	CreatedAt          time.Time      `gorm:"column:created_at" json:"created_at"`                     // 创建时间
	UpdatedBy          int64          `gorm:"column:updated_by" json:"updated_by"`                     // 修改人
	UpdatedAt          time.Time      `gorm:"column:updated_at" json:"updated_at"`                     // 修改时间
	DeletedAt          gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`                     // 删除标识（0-正常 1-删除）
}

// TableName TOemAppAssistRelease's table name
func (*TOemAppAssistRelease) TableName() string {
	return TableNameTOemAppAssistRelease
}
