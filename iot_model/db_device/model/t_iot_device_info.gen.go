// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTIotDeviceInfo = "t_iot_device_info"

// TIotDeviceInfo mapped from table <t_iot_device_info>
type TIotDeviceInfo struct {
	Id                int64          `gorm:"column:id;primaryKey" json:"id"`                        // 主键ID
	Did               string         `gorm:"column:did;not null" json:"did"`                        // 设备唯一ID（14位 1~9 A~Z随机）
	ProductId         int64          `gorm:"column:product_id;not null" json:"productId"`           // 产品ID(t_cloud_product.id)
	ProductKey        string         `gorm:"column:product_key" json:"productKey"`                  // 产品key
	OnlineStatus      int32          `gorm:"column:online_status" json:"onlineStatus"`              // 在线状态（0 在线 1 不在线）
	DeviceName        string         `gorm:"column:device_name" json:"deviceName"`                  // 设备名称
	DeviceNature      string         `gorm:"column:device_nature" json:"deviceNature"`              // 设备性质
	Sn                string         `gorm:"column:sn" json:"sn"`                                   // 设备SN
	BatchId           int64          `gorm:"column:batch_id" json:"batchId"`                        // 批次ID(t_cloud_device_batch.id)
	GroupId           int64          `gorm:"column:group_id" json:"groupId"`                        // 设备组ID（t_cloud_device_group.id）
	DeviceModel       string         `gorm:"column:device_model" json:"deviceModel"`                // 设备型号
	UserName          string         `gorm:"column:user_name" json:"userName"`                      // 用户名
	Passward          string         `gorm:"column:passward" json:"passward"`                       // 设备密码
	Salt              string         `gorm:"column:salt" json:"salt"`                               // 盐值
	DeviceSecretHttp  string         `gorm:"column:device_secret_http" json:"deviceSecretHttp"`     // 设备密钥（http）
	DeviceSecretMqtt  string         `gorm:"column:device_secret_mqtt" json:"deviceSecretMqtt"`     // 设备密钥（mqtt）
	IpAddress         string         `gorm:"column:ip_address" json:"ipAddress"`                    // ip地址
	Lat               float64        `gorm:"column:lat" json:"lat"`                                 // 纬度
	Lng               float64        `gorm:"column:lng" json:"lng"`                                 // 经度
	Country           string         `gorm:"column:country" json:"country"`                         // 国家编码
	Province          string         `gorm:"column:province" json:"province"`                       // 省份编码
	City              string         `gorm:"column:city" json:"city"`                               // 城市编码
	District          string         `gorm:"column:district" json:"district"`                       // 地区编码
	LastActivatedTime time.Time      `gorm:"column:last_activated_time" json:"lastActivatedTime"`   // 最后激活时间
	ActivatedTime     time.Time      `gorm:"column:activated_time" json:"activatedTime"`            // 激活时间
	MacAddress        string         `gorm:"column:mac_address" json:"macAddress"`                  // mac地址
	DeviceVersion     string         `gorm:"column:device_version" json:"deviceVersion"`            // 设备版本
	ActiveStatus      string         `gorm:"column:active_status" json:"activeStatus"`              // 激活状态[0:未激活,1:已激活]
	CreatedBy         int64          `gorm:"column:created_by;not null;default:0" json:"createdBy"` // 创建人
	UpdatedBy         int64          `gorm:"column:updated_by;not null;default:0" json:"updatedBy"` // 修改人
	CreatedAt         time.Time      `gorm:"column:created_at" json:"createdAt"`                    // 创建时间
	UpdatedAt         time.Time      `gorm:"column:updated_at" json:"updatedAt"`                    // 修改时间
	DeletedAt         gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`                    // 删除时间
	ActiveUserId      int64          `gorm:"column:active_user_id" json:"activeUserId"`             // 激活用户编号
	ActiveUserName    string         `gorm:"column:active_user_name" json:"activeUserName"`         // 激活用户名称
	TenantId          string         `gorm:"column:tenant_id" json:"tenantId"`                      // 租户编号
	AppKey            string         `gorm:"column:app_key" json:"appKey"`                          // APP Key
	ActiveChannel     int32          `gorm:"column:active_channel" json:"activeChannel"`            // 激活渠道
	ModuleVersion     string         `gorm:"column:module_version" json:"moduleVersion"`            // 模组版本
	Sid               int64          `gorm:"column:sid" json:"sid"`                                 // 绑定服务地址
	UseType           int32          `gorm:"column:use_type;default:0" json:"useType"`              // 使用类型（=1 虚拟测试设备）
}

// TableName TIotDeviceInfo's table name
func (*TIotDeviceInfo) TableName() string {
	return TableNameTIotDeviceInfo
}
