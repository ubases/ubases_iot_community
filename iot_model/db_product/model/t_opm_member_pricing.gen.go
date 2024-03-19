// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTOpmMemberPricing = "t_opm_member_pricing"

// TOpmMemberPricing mapped from table <t_opm_member_pricing>
type TOpmMemberPricing struct {
	Id        int64     `gorm:"column:id;primaryKey" json:"id"`                               // 主键ID
	AppKey    string    `gorm:"column:app_key;not null" json:"appKey"`                        // APP Key
	PayMethod int32     `gorm:"column:pay_method;not null" json:"payMethod"`                  // 支付方式
	UseModel  string    `gorm:"column:use_model;not null" json:"useModel"`                    // 授权使用机型（1 国内安卓、2 国外安卓、3 IOS系列）可设置多个，使用逗号分隔
	PriceType int32     `gorm:"column:price_type;not null" json:"priceType"`                  // 价格类型
	PriceId   string    `gorm:"column:price_id;not null" json:"priceId"`                      // 价格Id，对应支付平台配置的价格id
	Price     float64   `gorm:"column:price;not null" json:"price"`                           // 价格
	Status    int32     `gorm:"column:status;not null" json:"status"`                         // 生效状态- 考虑不需要改状态（1 生效中 2 待生效 3 失效）
	StartTime time.Time `gorm:"column:start_time;default:CURRENT_TIMESTAMP" json:"startTime"` // 生效开始时间（生效模式=1 则为当前时间）
	EndTime   time.Time `gorm:"column:end_time;default:CURRENT_TIMESTAMP" json:"endTime"`     // 生效结束时间
	CreatedBy int64     `gorm:"column:created_by;not null;default:0" json:"createdBy"`        // 创建人
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`                           // 创建时间
	UpdatedBy int64     `gorm:"column:updated_by;not null;default:0" json:"updatedBy"`        // 修改人
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`                           // 修改时间
}

// TableName TOpmMemberPricing's table name
func (*TOpmMemberPricing) TableName() string {
	return TableNameTOpmMemberPricing
}
