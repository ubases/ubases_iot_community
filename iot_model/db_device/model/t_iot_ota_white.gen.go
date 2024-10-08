// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTIotOtaWhite = "t_iot_ota_white"

// TIotOtaWhite mapped from table <t_iot_ota_white>
type TIotOtaWhite struct {
	Id         int64          `gorm:"column:id;primaryKey" json:"id"`                        // 主键ID
	PkgId      int64          `gorm:"column:pkg_id" json:"pkgId"`                            // 包ID(t_cloud_ota_pkg.id)
	ProductId  int64          `gorm:"column:product_id" json:"productId"`                    // 产品ID(t_cloud_product.id)
	GroupId    int64          `gorm:"column:group_id" json:"groupId"`                        // 组ID(t_cloud_ota_white_group.id)
	Did        string         `gorm:"column:did" json:"did"`                                 // 设备ID
	DeviceId   int64          `gorm:"column:device_id" json:"deviceId"`                      // 设备ID(t_iot_device.id)
	BelongType string         `gorm:"column:belong_type" json:"belongType"`                  // 数据归属类型
	BelongId   int64          `gorm:"column:belong_id" json:"belongId"`                      // 数据归属对象ID （开发平台UserId）
	CreatedBy  int64          `gorm:"column:created_by;not null;default:0" json:"createdBy"` // 创建人
	UpdatedBy  int64          `gorm:"column:updated_by;not null;default:0" json:"updatedBy"` // 修改人
	CreatedAt  time.Time      `gorm:"column:created_at" json:"createdAt"`                    // 创建时间
	UpdatedAt  time.Time      `gorm:"column:updated_at" json:"updatedAt"`                    // 修改时间
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`                    // 删除时间
}

// TableName TIotOtaWhite's table name
func (*TIotOtaWhite) TableName() string {
	return TableNameTIotOtaWhite
}
