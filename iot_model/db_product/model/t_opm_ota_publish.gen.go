// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTOpmOtaPublish = "t_opm_ota_publish"

// TOpmOtaPublish mapped from table <t_opm_ota_publish>
type TOpmOtaPublish struct {
	Id                   int64          `gorm:"column:id;primaryKey" json:"id"`                                        // 主键ID
	PkgId                int64          `gorm:"column:pkg_id;not null" json:"pkgId"`                                   // 包id（t_opm_ota_pkg.id）
	PublishAt            time.Time      `gorm:"column:publish_at;not null;default:CURRENT_TIMESTAMP" json:"publishAt"` // 发布时间
	IsGray               int32          `gorm:"column:is_gray;not null" json:"isGray"`                                 // 升级规模 =1  全量升级 =1 灰度升级
	Status               int32          `gorm:"column:status;not null" json:"status"`                                  // 状态[0:待发布,1:已发布,2:已暂停]
	GrayType             int32          `gorm:"column:gray_type" json:"grayType"`                                      // 灰度类型[1:按比例灰度, 2:按数量灰度, 3: 指定设备灰度]
	GrayScale            int32          `gorm:"column:gray_scale" json:"grayScale"`                                    // 灰度比例
	CreatedBy            int64          `gorm:"column:created_by;not null;default:0" json:"createdBy"`                 // 创建人
	UpdatedBy            int64          `gorm:"column:updated_by;not null;default:0" json:"updatedBy"`                 // 修改人
	CreatedAt            time.Time      `gorm:"column:created_at" json:"createdAt"`                                    // 创建时间
	UpdatedAt            time.Time      `gorm:"column:updated_at" json:"updatedAt"`                                    // 修改时间
	DeletedAt            gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`                                    // 删除时间
	TenantId             string         `gorm:"column:tenant_id;not null" json:"tenantId"`                             // 租户id（t_open_company.tenant_id）
	IsAuto               int32          `gorm:"column:is_auto;not null" json:"isAuto"`                                 // 是否自动升级[0:否 1:是]
	UpgradeDesc          string         `gorm:"column:upgrade_desc" json:"upgradeDesc"`                                // 升级文案（中文）
	UpgradeDescEn        string         `gorm:"column:upgrade_desc_en" json:"upgradeDescEn"`                           // 升级文案（英文）
	AutoStartAt          string         `gorm:"column:auto_start_at" json:"autoStartAt"`                               // 自动升级开始时间
	AutoEndAt            string         `gorm:"column:auto_end_at" json:"autoEndAt"`                                   // 自动升级结束时间
	SpecifiedVersionMode int32          `gorm:"column:specified_version_mode" json:"specifiedVersionMode"`             // 指定升级版本（模式选择 1全部版本 2指定版本）
	SpecifiedVersion     string         `gorm:"column:specified_version" json:"specifiedVersion"`                      // 指定升级版本
	SpecifiedAreaMode    int32          `gorm:"column:specified_area_mode" json:"specifiedAreaMode"`                   // 指定升级区域（模式选择 1全部区域 2指定区域）
	SpecifiedArea        string         `gorm:"column:specified_area" json:"specifiedArea"`                            // 指定升级区域
	VersionId            int64          `gorm:"column:version_id;not null" json:"versionId"`                           // 版本号编号
	Version              string         `gorm:"column:version" json:"version"`                                         // 固件版本号
	UpgradeMode          int32          `gorm:"column:upgrade_mode;not null" json:"upgradeMode"`                       // 升级方式[0:静默,1:提醒,2:强制]
	UpgradeTimeMode      int32          `gorm:"column:upgrade_time_mode;default:1" json:"upgradeTimeMode"`             // 升级时间模式 =1 全天 =2 指定时间
	TotalCount           int32          `gorm:"column:total_count" json:"totalCount"`                                  // 总升级条数
	PubResult            string         `gorm:"column:pub_result" json:"pubResult"`                                    // 发布结果
	SuccessCount         int32          `gorm:"column:success_count" json:"successCount"`                              // 成功升级条数
}

// TableName TOpmOtaPublish's table name
func (*TOpmOtaPublish) TableName() string {
	return TableNameTOpmOtaPublish
}
