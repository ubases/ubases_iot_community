// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTOpmCommunityProductLanguage = "t_opm_community_product_language"

// TOpmCommunityProductLanguage mapped from table <t_opm_community_product_language>
type TOpmCommunityProductLanguage struct {
	Id                 int64  `gorm:"column:id;primaryKey" json:"id"`                                 // 主键
	CommunityProductId int64  `gorm:"column:community_product_id;not null" json:"communityProductId"` // 设备产品编号
	Lang               string `gorm:"column:lang;not null" json:"lang"`                               // 语言分类
	Name               string `gorm:"column:name;not null" json:"name"`                               // 产品名称
	Description        string `gorm:"column:description" json:"description"`                          // 产品描述
}

// TableName TOpmCommunityProductLanguage's table name
func (*TOpmCommunityProductLanguage) TableName() string {
	return TableNameTOpmCommunityProductLanguage
}
