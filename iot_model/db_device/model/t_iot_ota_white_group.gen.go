// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTIotOtaWhiteGroup = "t_iot_ota_white_group"

// TIotOtaWhiteGroup mapped from table <t_iot_ota_white_group>
type TIotOtaWhiteGroup struct {
	Id        int64          `gorm:"column:id;primaryKey" json:"id"`                        // 主键ID
	GroupName string         `gorm:"column:group_name;not null" json:"groupName"`           // 测试白名单组名
	CreatedBy int64          `gorm:"column:created_by;not null;default:0" json:"createdBy"` // 创建人
	UpdatedBy int64          `gorm:"column:updated_by;not null;default:0" json:"updatedBy"` // 修改人
	CreatedAt time.Time      `gorm:"column:created_at" json:"createdAt"`                    // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updatedAt"`                    // 修改时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`                    // 删除时间
}

// TableName TIotOtaWhiteGroup's table name
func (*TIotOtaWhiteGroup) TableName() string {
	return TableNameTIotOtaWhiteGroup
}
