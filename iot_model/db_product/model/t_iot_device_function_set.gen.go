// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTIotDeviceFunctionSet = "t_iot_device_function_set"

// TIotDeviceFunctionSet mapped from table <t_iot_device_function_set>
type TIotDeviceFunctionSet struct {
	Id             int64     `gorm:"column:id;primaryKey" json:"id"`                        // 主键ID
	DeviceId       string    `gorm:"column:device_id;not null" json:"deviceId"`             // 设备Id
	ProductKey     string    `gorm:"column:product_key;not null" json:"productKey"`         // 产品Key
	FuncKey        string    `gorm:"column:func_key;not null" json:"funcKey"`               // 功能Dpid
	FuncIdentifier string    `gorm:"column:func_identifier" json:"funcIdentifier"`          // 功能标识符
	FuncValue      string    `gorm:"column:func_value" json:"funcValue"`                    // 功能值
	CustomType     int32     `gorm:"column:custom_type" json:"customType"`                  // 自定义的类型（1=物模型的属性  2=物模型的值）
	CustomDesc     string    `gorm:"column:custom_desc" json:"customDesc"`                  // 功能值描述
	CreatedBy      int64     `gorm:"column:created_by;not null;default:0" json:"createdBy"` // 创建人
	UpdatedBy      int64     `gorm:"column:updated_by;not null;default:0" json:"updatedBy"` // 修改人
	CreatedAt      time.Time `gorm:"column:created_at" json:"createdAt"`                    // 创建时间
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updatedAt"`                    // 修改时间
}

// TableName TIotDeviceFunctionSet's table name
func (*TIotDeviceFunctionSet) TableName() string {
	return TableNameTIotDeviceFunctionSet
}
