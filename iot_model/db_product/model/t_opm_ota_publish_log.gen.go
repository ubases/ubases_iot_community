// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTOpmOtaPublishLog = "t_opm_ota_publish_log"

// TOpmOtaPublishLog mapped from table <t_opm_ota_publish_log>
type TOpmOtaPublishLog struct {
	Id         int64          `gorm:"column:id;primaryKey" json:"id"`                        // 主键ID
	ProductId  int64          `gorm:"column:product_id" json:"productId"`                    // 产品编号
	FirmwareId int64          `gorm:"column:firmware_id" json:"firmwareId"`                  // 固件编号
	PkgId      int64          `gorm:"column:pkg_id;not null" json:"pkgId"`                   // 包id（t_opm_ota_pkg.id）
	IsGray     int32          `gorm:"column:is_gray;not null" json:"isGray"`                 // 是否灰度发布 0：否，1：是
	GrayType   int32          `gorm:"column:gray_type" json:"grayType"`                      // 灰度类型[0:按比例灰度, 1:按数量灰度]
	GrayScale  int32          `gorm:"column:gray_scale" json:"grayScale"`                    // 灰度比例
	Type       int32          `gorm:"column:type;not null" json:"type"`                      // OTA类型[0:固件]
	Version    string         `gorm:"column:version;not null" json:"version"`                // OTA版本号
	Did        string         `gorm:"column:did;not null" json:"did"`                        // 设备唯一ID（14位 1~9 A~Z随机）
	Status     int32          `gorm:"column:status;not null" json:"status"`                  // 状态[0:升级成功, 1:升级失败]
	DeviceLog  string         `gorm:"column:device_log" json:"deviceLog"`                    // 设备日志
	CreatedBy  int64          `gorm:"column:created_by;not null;default:0" json:"createdBy"` // 创建人
	UpdatedBy  int64          `gorm:"column:updated_by;not null;default:0" json:"updatedBy"` // 修改人
	CreatedAt  time.Time      `gorm:"column:created_at" json:"createdAt"`                    // 创建时间
	UpdatedAt  time.Time      `gorm:"column:updated_at" json:"updatedAt"`                    // 修改时间
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`                    // 删除时间
	TenantId   string         `gorm:"column:tenant_id;not null" json:"tenantId"`             // 租户id（t_open_company.tenant_id）
}

// TableName TOpmOtaPublishLog's table name
func (*TOpmOtaPublishLog) TableName() string {
	return TableNameTOpmOtaPublishLog
}
