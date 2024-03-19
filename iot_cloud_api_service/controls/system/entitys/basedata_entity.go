package entitys

import (
	"time"
)

type BaseDataQuery struct {
	Page      int64  `json:"page,omitempty"`
	Limit     int64  `json:"limit,omitempty"`
	Sort      string `json:"sort,omitempty"`
	SortField string `json:"sortField,omitempty"`
	SearchKey string `json:"searchKey,omitempty"`
	DictLabel string `json:"dict_label,omitempty" form:"dictLabel"`
	DictType  string `json:"dict_type,omitempty" form:"dictType"`
	DictValue string `json:"dictValue"`
}

type BaseDataTypeQuery struct {
	Page      int64  `json:"page,omitempty" form:"pageNum"`
	Limit     int64  `json:"limit,omitempty" form:"pageSize"`
	Sort      string `json:"sort,omitempty"`
	SortField string `json:"sortField,omitempty"`
	SearchKey string `json:"searchKey,omitempty"`
	DictName  string `json:"dict_name,omitempty" form:"dictName"`
	DictType  string `json:"dict_type,omitempty" form:"dictType"`
}

type BaseData struct {
	DictId      string      `gorm:"column:dictId;primaryKey;autoIncrement:true" json:"dictId" form:"dictId"` // 字典编码
	DictSort    int32       `gorm:"column:dict_sort;default:0" json:"dictSort"`                              // 字典排序
	DictLabel   string      `gorm:"column:dict_label" json:"dictLabel"`                                      // 字典标签
	DictValue   interface{} `gorm:"column:dict_value" json:"dictValue"`                                      // 字典键值
	DictType    string      `gorm:"column:dict_type" json:"dictType"`                                        // 字典类型
	ValueType   int32       `json:"valueType,omitempty"`                                                     // 1-整形，2-浮点,  3-字符串
	CSSClass    string      `gorm:"column:css_class" json:"cssClass"`                                        // 样式属性（其他样式扩展）
	ListClass   string      `gorm:"column:list_class" json:"listClass"`                                      // 表格回显样式
	IsDefault   int32       `gorm:"column:is_default;default:0" json:"isDefault"`                            // 是否默认（1是 0否）
	Status      int32       `gorm:"column:status;default:0" json:"status"`                                   // 状态（0正常 1停用）
	CreateBy    string      `gorm:"column:create_by;default:0" json:"createBy"`                              // 创建者
	UpdateBy    string      `gorm:"column:update_by;default:0" json:"updateBy"`                              // 更新者
	Remark      string      `gorm:"column:remark" json:"remark"`                                             // 备注
	Pinyin      string      `gorm:"column:pinyin" json:"pinyin"`                                             // 拼音
	Firstletter string      `gorm:"column:firstletter" json:"firstletter"`                                   // 首字母
	Listimg     string      `gorm:"column:listimg" json:"listimg"`                                           // 图片
	CreatedAt   time.Time   `gorm:"column:created_at" json:"createdAt"`                                      // 创建时间
	UpdatedAt   time.Time   `gorm:"column:updated_at" json:"updatedAt"`                                      // 修改时间
}

type BaseDataType struct {
	DictID    string    `gorm:"column:dict_id;primaryKey;autoIncrement:true" json:"dictId"` // 字典主键
	DictName  string    `gorm:"column:dict_name" json:"dictName"`                           // 字典名称
	DictType  string    `gorm:"column:dict_type" json:"dictType"`                           // 字典类型
	ValueType int32     `json:"valueType,omitempty"`                                        // 1-整形，2-浮点,  3-字符串
	IsSystem  int32     `json:"isSystem,omitempty"`                                         // 是否为系统参数 =1 是 =2 否
	Status    int32     `gorm:"column:status;default:0" json:"status"`                      // 状态（0正常 1停用）
	CreateBy  string    `gorm:"column:create_by;default:0" json:"createBy"`                 // 创建者
	UpdateBy  string    `gorm:"column:update_by;default:0" json:"updateBy"`                 // 更新者
	Remark    string    `gorm:"column:remark" json:"remark"`                                // 备注
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`                         // 创建日期
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`                         // 修改日期
}

type TranslateParam struct {
	ID   string `gorm:"column:id;primaryKey" json:"id"` // 主键
	Code string `gorm:"column:code" json:"code"`        // 名称
	En   string `gorm:"column:en" json:"en"`            // 英文
	Zh   string `gorm:"column:zh" json:"zh"`            // 中文
	Jp   string `gorm:"column:jp" json:"jp"`            // 日文
}

type DictKeyVal struct {
	Name string `json:"name"`
	Code string `json:"code"`
}
