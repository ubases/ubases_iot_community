// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTSysDictData = "t_sys_dict_data"

// TSysDictData mapped from table <t_sys_dict_data>
type TSysDictData struct {
	DictCode  int64          `gorm:"column:dict_code;primaryKey;autoIncrement:true" json:"dictCode"` // 字典编码
	DictSort  int32          `gorm:"column:dict_sort;default:0" json:"dictSort"`                     // 字典排序
	DictLabel string         `gorm:"column:dict_label;default:''" json:"dictLabel"`                  // 字典标签
	DictValue string         `gorm:"column:dict_value;default:''" json:"dictValue"`                  // 字典键值
	DictType  string         `gorm:"column:dict_type;default:''" json:"dictType"`                    // 字典类型
	CssClass  string         `gorm:"column:css_class" json:"cssClass"`                               // 样式属性（其他样式扩展）
	ListClass string         `gorm:"column:list_class" json:"listClass"`                             // 表格回显样式
	IsDefault int32          `gorm:"column:is_default;default:0" json:"isDefault"`                   // 是否默认（1是 0否）
	Status    int32          `gorm:"column:status;default:0" json:"status"`                          // 状态（0正常 1停用）
	CreateBy  int64          `gorm:"column:create_by;default:0" json:"createBy"`                     // 创建者
	UpdateBy  int64          `gorm:"column:update_by;default:0" json:"updateBy"`                     // 更新者
	Remark    string         `gorm:"column:remark" json:"remark"`                                    // 备注
	CreatedAt time.Time      `gorm:"column:created_at" json:"createdAt"`                             // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updatedAt"`                             // 修改时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`                             // 删除时间
}

// TableName TSysDictData's table name
func (*TSysDictData) TableName() string {
	return TableNameTSysDictData
}
