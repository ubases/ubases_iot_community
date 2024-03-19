// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTSysArea = "t_sys_area"

// TSysArea mapped from table <t_sys_area>
type TSysArea struct {
	Id              int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`          // 主键
	Pid             int64  `gorm:"column:pid;default:0" json:"pid"`                            // 父ID
	Level           int32  `gorm:"column:level;default:0" json:"level"`                        // 层级
	Path            string `gorm:"column:path;default:''" json:"path"`                         // 路径
	Code            string `gorm:"column:code;default:''" json:"code"`                         // 代码
	AreaNumber      string `gorm:"column:area_number;default:''" json:"areaNumber"`            // 国家区号
	AreaPhoneNumber string `gorm:"column:area_phone_number;default:''" json:"areaPhoneNumber"` // 国家电话区号
	Abbreviation    string `gorm:"column:abbreviation;default:''" json:"abbreviation"`         // 国家缩写
	Iso             string `gorm:"column:iso;default:''" json:"iso"`                           // 时区
	ChineseName     string `gorm:"column:chinese_name;default:''" json:"chineseName"`          // 中文名称
	EnglishName     string `gorm:"column:english_name;default:''" json:"englishName"`          // 英文名称
	Pinyin          string `gorm:"column:pinyin;default:''" json:"pinyin"`                     // 中文拼音
}

// TableName TSysArea's table name
func (*TSysArea) TableName() string {
	return TableNameTSysArea
}
