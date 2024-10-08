// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTOpmProductMaterialLanguage = "t_opm_product_material_language"

// TOpmProductMaterialLanguage mapped from table <t_opm_product_material_language>
type TOpmProductMaterialLanguage struct {
	Id             int64  `gorm:"column:id;primaryKey" json:"id"`                        // 主键
	MaterialId     int64  `gorm:"column:material_id;not null" json:"materialId"`         // 耗材编号
	Lang           string `gorm:"column:lang;not null" json:"lang"`                      // 语言分类
	Name           string `gorm:"column:name;not null" json:"name"`                      // 耗材名称
	BrandName      string `gorm:"column:brand_name;not null" json:"brandName"`           // 品牌名称
	FragranceName  string `gorm:"column:fragrance_name;not null" json:"fragranceName"`   // 香型名称
	ProductAddress string `gorm:"column:product_address;not null" json:"productAddress"` // 产地
	Variety        string `gorm:"column:variety;not null" json:"variety"`                // 品种
	Ingredient     string `gorm:"column:ingredient" json:"ingredient"`                   // 成分
	Effect         string `gorm:"column:effect" json:"effect"`                           // 功效
	Description    string `gorm:"column:description" json:"description"`                 // 描述
}

// TableName TOpmProductMaterialLanguage's table name
func (*TOpmProductMaterialLanguage) TableName() string {
	return TableNameTOpmProductMaterialLanguage
}
