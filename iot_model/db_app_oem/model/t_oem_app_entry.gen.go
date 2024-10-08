// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTOemAppEntry = "t_oem_app_entry"

// TOemAppEntry mapped from table <t_oem_app_entry>
type TOemAppEntry struct {
	Id        int64     `gorm:"column:id;primaryKey" json:"id"`   // 词条id
	Lang      string    `gorm:"column:lang" json:"lang"`          // 语种编码
	Title     string    `gorm:"column:title" json:"title"`        // 标题
	Content   string    `gorm:"column:content" json:"content"`    // 内容
	SetingId  int64     `gorm:"column:seting_id" json:"setingId"` // 设置id
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName TOemAppEntry's table name
func (*TOemAppEntry) TableName() string {
	return TableNameTOemAppEntry
}
