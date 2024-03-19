// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTOpenLoginLog = "t_open_login_log"

// TOpenLoginLog mapped from table <t_open_login_log>
type TOpenLoginLog struct {
	InfoId        int64     `gorm:"column:info_id;primaryKey;autoIncrement:true" json:"infoId"` // 访问ID
	LoginName     string    `gorm:"column:login_name;default:''" json:"loginName"`              // 登录账号
	Ipaddr        string    `gorm:"column:ipaddr;default:''" json:"ipaddr"`                     // 登录IP地址
	LoginLocation string    `gorm:"column:login_location;default:''" json:"loginLocation"`      // 登录地点
	Browser       string    `gorm:"column:browser;default:''" json:"browser"`                   // 浏览器类型
	Os            string    `gorm:"column:os;default:''" json:"os"`                             // 操作系统
	Status        int32     `gorm:"column:status;default:0" json:"status"`                      // 登录状态（0成功 1失败）
	Msg           string    `gorm:"column:msg;default:''" json:"msg"`                           // 提示消息
	LoginTime     time.Time `gorm:"column:login_time" json:"loginTime"`                         // 登录时间
	Module        string    `gorm:"column:module;default:''" json:"module"`                     // 登录模块
}

// TableName TOpenLoginLog's table name
func (*TOpenLoginLog) TableName() string {
	return TableNameTOpenLoginLog
}
