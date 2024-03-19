// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTOpmDocumentEvaluate = "t_opm_document_evaluate"

// TOpmDocumentEvaluate mapped from table <t_opm_document_evaluate>
type TOpmDocumentEvaluate struct {
	Id         int64     `gorm:"column:id;primaryKey" json:"id"`                // 主键
	ProductKey string    `gorm:"column:product_key;not null" json:"productKey"` // 产品Key
	DocId      int64     `gorm:"column:doc_id;not null" json:"docId"`           // 文档Id
	Useful     int32     `gorm:"column:useful;not null" json:"useful"`          // 有用
	Useless    int32     `gorm:"column:useless;not null" json:"useless"`        // 无用
	CreatedAt  time.Time `gorm:"column:created_at" json:"createdAt"`            // 创建时间
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updatedAt"`            // 修改时间
}

// TableName TOpmDocumentEvaluate's table name
func (*TOpmDocumentEvaluate) TableName() string {
	return TableNameTOpmDocumentEvaluate
}
