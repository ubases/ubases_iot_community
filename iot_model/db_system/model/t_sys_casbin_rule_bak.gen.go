// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTSysCasbinRuleBak = "t_sys_casbin_rule_bak"

// TSysCasbinRuleBak mapped from table <t_sys_casbin_rule_bak>
type TSysCasbinRuleBak struct {
	Id    int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Ptype string `gorm:"column:ptype" json:"ptype"`
	V0    string `gorm:"column:v0" json:"v0"`
	V1    string `gorm:"column:v1" json:"v1"`
	V2    string `gorm:"column:v2" json:"v2"`
	V3    string `gorm:"column:v3" json:"v3"`
	V4    string `gorm:"column:v4" json:"v4"`
	V5    string `gorm:"column:v5" json:"v5"`
	V6    string `gorm:"column:v6" json:"v6"`
	V7    string `gorm:"column:v7" json:"v7"`
}

// TableName TSysCasbinRuleBak's table name
func (*TSysCasbinRuleBak) TableName() string {
	return TableNameTSysCasbinRuleBak
}