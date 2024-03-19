// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTIotOtaVersionPublish = "t_iot_ota_version_publish"

// TIotOtaVersionPublish mapped from table <t_iot_ota_version_publish>
type TIotOtaVersionPublish struct {
	Id           int64          `gorm:"column:id;primaryKey" json:"id"`                        // 主键ID
	VersionId    int64          `gorm:"column:version_id;not null" json:"versionId"`           // 大版本ID(t_cloud_ota_version.id)
	PublishMode  int32          `gorm:"column:publish_mode;not null" json:"publishMode"`       // 发布模式[0:立即发布,1:定时发布]
	ScheduleTime time.Time      `gorm:"column:schedule_time;not null" json:"scheduleTime"`     // 定时时间
	PublishTime  time.Time      `gorm:"column:publish_time;not null" json:"publishTime"`       // 发布时间
	Status       int32          `gorm:"column:status;not null" json:"status"`                  // 状态[0:已撤消,1:已发布,2:待发布]
	IsGray       int32          `gorm:"column:is_gray;not null" json:"isGray"`                 // 是否灰度发布 0：否，1：是
	UpdateDesc   string         `gorm:"column:update_desc" json:"updateDesc"`                  // 更新描述
	CreatedBy    int64          `gorm:"column:created_by;not null;default:0" json:"createdBy"` // 创建人
	UpdatedBy    int64          `gorm:"column:updated_by;not null;default:0" json:"updatedBy"` // 修改人
	CreatedAt    time.Time      `gorm:"column:created_at" json:"createdAt"`                    // 创建时间
	UpdatedAt    time.Time      `gorm:"column:updated_at" json:"updatedAt"`                    // 修改时间
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`                    // 删除时间
}

// TableName TIotOtaVersionPublish's table name
func (*TIotOtaVersionPublish) TableName() string {
	return TableNameTIotOtaVersionPublish
}
