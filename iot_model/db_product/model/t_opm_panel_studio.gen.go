// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTOpmPanelStudio = "t_opm_panel_studio"

// TOpmPanelStudio mapped from table <t_opm_panel_studio>
type TOpmPanelStudio struct {
	Id           int64     `gorm:"column:id;primaryKey" json:"id"`                        // 主键ID
	PanelId      int64     `gorm:"column:panel_id;not null" json:"panelId"`               // 面板Id
	PageName     string    `gorm:"column:page_name" json:"pageName"`                      // 页面名称
	PageIdentify string    `gorm:"column:page_identify" json:"pageIdentify"`              // 页面英文名称
	JsonContent  string    `gorm:"column:json_content;not null" json:"jsonContent"`       // JSON内容数据
	PopupContent string    `gorm:"column:popup_content" json:"popupContent"`              // 弹框内容部分
	VueContent   string    `gorm:"column:vue_content" json:"vueContent"`                  // VUE内容数据
	StyleContent string    `gorm:"column:style_content" json:"styleContent"`              // Style内容数据
	IsHome       int32     `gorm:"column:is_home" json:"isHome"`                          // 是否主页
	Sort         int32     `gorm:"column:sort" json:"sort"`                               // 排序
	CreatedAt    time.Time `gorm:"column:created_at" json:"createdAt"`                    // 创建时间
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updatedAt"`                    // 修改时间
	UpdatedBy    int64     `gorm:"column:updated_by;not null;default:0" json:"updatedBy"` // 修改人
}

// TableName TOpmPanelStudio's table name
func (*TOpmPanelStudio) TableName() string {
	return TableNameTOpmPanelStudio
}
