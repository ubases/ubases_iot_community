package entitys

import "time"

// AppProductTypeVo mapped from table <t_pm_product_type>
type AppProductTypeVo struct {
	Id     string `json:"id,omitempty"`     // 主键（雪花算法19位）
	Name   string `json:"name,omitempty"`   // 分类名称
	NameEn string `json:"nameEn,omitempty"` // 分类名称（英文）
	Sort   int32  `json:"sort"`             // 排序
	Desc   string `json:"desc,omitempty"`   // 描述
}

// QueryProductTypeForm query ProductType  form ;  if some field is required, create binding:"required" to tag by self
type AppQueryProductTypeForm struct {
	Id         int64  `json:"id" form:"id"`                 // cond Id
	Name       string `json:"name" form:"name"`             // cond Name
	NameEn     string `json:"nameEn" form:"nameEn"`         // cond NameEn
	Identifier string `json:"identifier" form:"identifier"` // cond Identifier
	Order      int    `json:"order" form:"order"`           // example: orderMap[column]=desc
	Page       int    `json:"page" form:"page"`             // get all without uploading
	Limit      int    `json:"limit" form:"limit"`           // get all without uploading
	SearchKey  string `json:"searchKey" form:"searchKey"`   // Search key
}

// TPmProductType mapped from table <t_pm_product_type>
type TPmProductTypeVo struct {
	Id          string              `json:"id,omitempty"`         // 主键（雪花算法19位）
	Name        string              `json:"name,omitempty"`       // 分类名称
	NameEn      string              `json:"nameEn,omitempty"`     // 分类名称（英文）
	Identifier  string              `json:"identifier,omitempty"` // 属性的标识符。可包含大小写英文字母、数字、下划线（_），长度不超过50个字符。
	Sort        int32               `json:"sort"`                 // 排序
	ParentId    string              `json:"parentId"`             // 父ID
	Desc        string              `json:"desc,omitempty"`       // 描述
	ImgFullPath string              `json:"pic,omitempty"`        // 图片
	ParentName  string              `json:"parentName,omitempty"` //父级分类名称
	Children    []*TPmProductTypeVo `json:"children"`             //递归引用链
	Products    []*TPmProductVo     `json:"products"`             //递归引用链-产品
}

// TPmProductVo mapped from table <t_pm_product>
type TPmProductVo struct {
	Id               string                   `json:"id"`                               // 主键（雪花算法19位）
	ProductTypeId    string                   `json:"productTypeId,omitempty"`          // 产品类型ID
	ProductKey       string                   `json:"productKey,omitempty"`             // 产品唯一标识
	Name             string                   `json:"name"`                             // 产品名称
	NameEn           string                   `json:"nameEn,omitempty"`                 // 产品名称（英文）
	Identifier       string                   `json:"identifier,omitempty"`             // 属性的标识符。可包含大小写英文字母、数字、下划线（_），长度不超过50个字符。
	Model            string                   `json:"model,omitempty"`                  // 产品型号
	ImageURL         string                   `json:"imageUrl,omitempty"`               // 产品图片
	WifiFlag         string                   `json:"wifiFlag,omitempty"`               // WIFI标识
	NetworkType      string                   `json:"networkType,omitempty"`            // 通信协议（WIFI, BLE, WIFI+BLE）
	AttributeType    int32                    `json:"attributeType"`                    // 设备性质（0:普通设备，1：网关设备）
	Status           int32                    `json:"status"`                           // 状态（0：未发布，1：已发布，2：停用）
	IsVirtualTest    int32                    `json:"isVirtualTest"`                    // 是否支持虚拟测试（0：否，1：是）
	Desc             string                   `json:"desc,omitempty"`                   // 描述
	CreatedBy        int64                    `json:"createdBy,omitempty"`              // 创建人
	CreatedAt        string                   `json:"createdAt,omitempty"`              // 创建时间
	UpdatedBy        int64                    `json:"updatedBy,omitempty"`              // 修改人
	UpdatedAt        string                   `json:"updatedAt,omitempty"`              // 修改时间
	DeletedAt        int32                    `json:"deletedAt,omitempty"`              // 删除的标识 0-正常 1-删除
	ProductTypeName  string                   `json:"productTypeName,omitempty"`        // 产品分类名称
	ModelId          string                   `json:"modelId,omitempty" form:"modelId"` // modelId
	ThingModels      []*TPmThingModelVo       `json:"thingModels,omitempty"`            //物模型列表
	ControlPanelIds  []string                 `json:"controlPanelIds,omitempty"`        //控制面板
	ModuleIds        []string                 `json:"moduleIds,omitempty"`              //模组
	FirmwareIds      []string                 `json:"firmwareIds,omitempty"`            //固件
	PowerConsumeType int32                    `json:"powerConsumeType"`
	NetworkGuides    []*PmNetworkGuideEntitys `json:"networkGuides,omitempty"`
}

// TPmProductVo mapped from table <t_pm_product>
type TPmThingModelVo struct {
	Id          string `json:"id,omitempty"`         // 主键ID
	FuncType    string `json:"funcType,omitempty"`   // 功能类型
	FuncName    string `json:"funcName,omitempty"`   // 功能名称
	Required    int32  `json:"required"`             // 必选
	Identifier  string `json:"identifier,omitempty"` // 标识符
	RwFlag      string `json:"rwFlag,omitempty"`     // 数据传输类型
	DataType    string `json:"dataType,omitempty"`   // 数据类型
	Attribute   string `json:"attribute,omitempty"`  // 功能点属性
	Desc        string `json:"desc,omitempty"`
	TriggerCond bool   `json:"triggerCond"`
	ExecCond    bool   `json:"execCond"`
	Valid       bool   `json:"valid"`
}

type PmNetworkGuideEntitys struct {
	Id            int64                        `json:"id,omitempty"`
	ProductId     int64                        `json:"productId,omitempty"`
	ProductTypeId int64                        `json:"productTypeId,omitempty"`
	Type          int32                        `json:"type"`
	CreatedBy     int64                        `json:"createdBy,omitempty"`
	CreatedAt     time.Time                    `json:"createdAt,omitempty"`
	UpdatedBy     int64                        `json:"updatedBy,omitempty"`
	UpdatedAt     time.Time                    `json:"updatedAt,omitempty"`
	DeletedAt     time.Time                    `json:"deletedAt,omitempty"`
	Steps         []*PmNetworkGuideStepEntitys `json:"steps,omitempty"`
}

type PmNetworkGuideStepEntitys struct {
	Id             int64     `json:"id,omitempty"`
	NetworkGuideId int64     `json:"networkGuideId,omitempty"`
	Instruction    string    `json:"instruction,omitempty"`
	InstructionEn  string    `json:"instructionEn,omitempty"`
	ImageUrl       string    `json:"imageUrl,omitempty"`
	VideoUrl       string    `json:"videoUrl,omitempty"`
	Sort           int32     `json:"sort,omitempty"`
	CreatedBy      int64     `json:"createdBy,omitempty"`
	CreatedAt      time.Time `json:"createdAt,omitempty"`
	UpdatedBy      int64     `json:"updatedBy,omitempty"`
	UpdatedAt      time.Time `json:"updatedAt,omitempty"`
	DeletedAt      time.Time `json:"deletedAt,omitempty"`
}
