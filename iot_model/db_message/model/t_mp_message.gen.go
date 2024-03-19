// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTMpMessage = "t_mp_message"

// TMpMessage mapped from table <t_mp_message>
type TMpMessage struct {
	Id          int64     `gorm:"column:id;primaryKey" json:"id"`
	PushType    int32     `gorm:"column:push_type;not null" json:"pushType"`       // 推送类型[0:站内且站外，1：站内， 2：站外]
	MessageType int32     `gorm:"column:message_type;not null" json:"messageType"` // 消息类型[0：系统提醒，1：设备提醒，....自定义]
	PushTo      string    `gorm:"column:push_to;not null" json:"pushTo"`           // 接收对象[数据字典配置， 比如全用户：all、user、home、device]
	PushMode    int32     `gorm:"column:push_mode;not null" json:"pushMode"`       // 发送模式[0:实时发送， 1：定时发送，2：轮询发送]
	PushStatus  int32     `gorm:"column:push_status;not null" json:"pushStatus"`   // 发送状态[0:已发送， 1：待发送， 2：发送失败， 3：已删除]
	AgentType   int32     `gorm:"column:agent_type;not null" json:"agentType"`     // 终端类型[0:所有终端， 1：IOS端， 2:android端]
	PushTime    time.Time `gorm:"column:push_time;not null" json:"pushTime"`       // 推送时间
	ExpireHour  int32     `gorm:"column:expire_hour" json:"expireHour"`            // 有效时间(单位小时)
	ActionType  int32     `gorm:"column:action_type;not null" json:"actionType"`   // 行为类型[0:主动发送， 1:被动发送]
	TplCode     string    `gorm:"column:tpl_code" json:"tplCode"`                  // 消息模板编号
	PushParams  string    `gorm:"column:push_params" json:"pushParams"`            // 消息模板的参数
	TargetIds   string    `gorm:"column:target_ids" json:"targetIds"`              // 推送的目标编号集合(push_to = user  为userid集合）,使用逗号分割
	Did         string    `gorm:"column:did" json:"did"`
	ProductKey  string    `gorm:"column:product_key" json:"productKey"`                  // 推送目标产品
	Content     string    `gorm:"column:content" json:"content"`                         // 消息内容
	CreatedBy   int64     `gorm:"column:created_by;not null;default:0" json:"createdBy"` // 创建人
	CreatedAt   time.Time `gorm:"column:created_at" json:"createdAt"`                    // 创建时间
}

// TableName TMpMessage's table name
func (*TMpMessage) TableName() string {
	return TableNameTMpMessage
}
