package entitys

import (
	"cloud_platform/iot_common/iotutil"
	"errors"
	"fmt"
	"time"
)

// Product table
type Product struct {
	Id            int64     `gorm:"PRIMARY_KEY"`
	ProductTypeId string    `gorm:"Column:product_type_id"`
	ProductKey    string    `gorm:"Column:product_key"`
	Name          string    `gorm:"Column:name"`
	NameEn        string    `gorm:"Column:name_en"`
	Identifier    string    `gorm:"Column:identifier"`
	Model         string    `gorm:"Column:model"`
	ImageUrl      string    `gorm:"Column:image_url"`
	WifiFlag      int32     `gorm:"Column:wifi_flag"`
	NetworkType   int32     `gorm:"Column:network_type"`
	AttributeType int32     `gorm:"Column:attribute_type"`
	Status        int32     `gorm:"Column:status"`
	IsVirtualTest int32     `gorm:"Column:is_virtual_test"`
	Desc          string    `gorm:"Column:desc"`
	CreatedBy     string    `gorm:"Column:created_by"`
	CreatedAt     time.Time `gorm:"Column:created_at"`
	UpdatedBy     string    `gorm:"Column:updated_by"`
	UpdatedAt     time.Time `gorm:"Column:updated_at"`
	DeletedAt     int32     `gorm:"Column:deleted_at"`
}

// QueryProductForm query Product  form ;  if some field is required, create binding:"required" to tag by self
type QueryProductForm struct {
	Id            string `json:"id" form:"id"`                       // cond Id
	ProductTypeId string `json:"productTypeId" form:"productTypeId"` // cond ProductTypeId
	ProductKey    string `json:"productKey" form:"productKey"`       // cond ProductKey
	Name          string `json:"name" form:"name"`                   // cond Name
	NameEn        string `json:"nameEn" form:"nameEn"`               // cond NameEn
	Identifier    string `json:"identifier" form:"identifier"`       // cond Identifier
	Model         string `json:"model" form:"model"`                 // cond Model
	ImageUrl      string `json:"imageUrl" form:"imageUrl"`           // cond ImageUrl
	WifiFlag      string `json:"wifiFlag" form:"wifiFlag"`           // cond WifiFlag
	NetworkType   int32  `json:"networkType" form:"networkType"`     // cond NetworkType
	AttributeType *int32 `json:"attributeType" form:"attributeType"` // example: AttributeType[>]=some value&AttributeType[<]=some value; key must be ">,>=,<,<=,="
	Status        *int32 `json:"status" form:"status"`               // example: Status[>]=some value&Status[<]=some value; key must be ">,>=,<,<=,="
	IsVirtualTest int32  `json:"isVirtualTest" form:"isVirtualTest"` // example: IsVirtualTest[>]=some value&IsVirtualTest[<]=some value; key must be ">,>=,<,<=,="
	Desc          string `json:"desc" form:"desc"`                   // cond Desc
	CreatedBy     string `json:"createdBy" form:"createdBy"`         // cond CreatedBy
	CreatedAt     string `json:"createdAt" form:"createdAt"`         // example: CreatedAt[>]=some value&CreatedAt[<]=some value; key must be ">,>=,<,<=,="
	UpdatedBy     string `json:"updatedBy" form:"updatedBy"`         // cond UpdatedBy
	UpdatedAt     string `json:"updatedAt" form:"updatedAt"`         // example: UpdatedAt[>]=some value&UpdatedAt[<]=some value; key must be ">,>=,<,<=,="
	DeletedAt     int32  `json:"deletedAt" form:"deletedAt"`         // example: DeletedAt[>]=some value&DeletedAt[<]=some value; key must be ">,>=,<,<=,="
	Order         int    `json:"order" form:"order"`                 // example: order[column]=desc
	Page          int    `json:"page" form:"page"`                   // get all without uploading
	Limit         int    `json:"limit" form:"limit"`                 // get all without uploading
	SearchKey     string `json:"searchKey" form:"searchKey"`         // Search key
}

// CreateProductForm create Product form
type CreateProductForm struct {
	Id                string                   `json:"id" form:"id"`                               // id
	ProductTypeId     string                   `json:"productTypeId" form:"productTypeId"`         // productTypeId
	ProductTypeIdPath []string                 `json:"productTypeIdPath" form:"productTypeIdPath"` // productTypeIdPath
	ProductKey        string                   `json:"productKey" form:"productKey"`               // productKey
	Name              string                   `json:"name" form:"name"`                           // name
	NameEn            string                   `json:"nameEn" form:"nameEn"`                       // nameEn
	Identifier        string                   `json:"identifier" form:"identifier"`               // identifier
	Model             string                   `json:"model" form:"model"`                         // model
	ImageUrl          string                   `json:"imageUrl" form:"imageUrl"`                   // imageUrl
	WifiFlag          string                   `json:"wifiFlag" form:"wifiFlag"`                   // wifiFlag
	NetworkType       int32                    `json:"networkType" form:"networkType"`             // networkType
	AttributeType     int32                    `json:"attributeType" form:"attributeType"`         // attributeType
	Status            int32                    `json:"status" form:"status"`                       // status
	IsVirtualTest     int32                    `json:"isVirtualTest" form:"isVirtualTest"`         // isVirtualTest
	Remark            string                   `json:"remark" form:"remark"`                       // remark
	CreatedBy         int64                    `json:"createdBy" form:"createdBy"`                 // createdBy
	UpdatedBy         int64                    `json:"updatedBy" form:"updatedBy"`                 // updatedBy
	DeletedAt         int32                    `json:"deletedAt" form:"deletedAt"`                 // deletedAt
	IsPublish         int32                    `json:"isPublish" form:"isPublish"`                 // isPublish
	ThingModels       []*TPmThingModelVo       `json:"thingModels,omitempty"`                      //物模型列表
	ControlPanelIds   []string                 `json:"controlPanelIds,omitempty"`                  //控制面板
	ModuleIds         []string                 `json:"moduleIds,omitempty"`                        //模组
	FirmwareIds       []string                 `json:"firmwareIds,omitempty"`                      //固件
	PowerConsumeType  int32                    `json:"powerConsumeType"`
	NetworkGuides     []*PmNetworkGuideEntitys `json:"networkGuides,omitempty"`
}

// Valid create Product  form verify
func (a *CreateProductForm) Valid() error {
	_, err := iotutil.ToInt64AndErr(a.ProductTypeId)
	if err != nil {
		return errors.New("productTypeId异常")
	}
	if a.ProductTypeIdPath == nil || len(a.ProductTypeIdPath) == 0 {
		return errors.New("productTypeId路径不能为空")
	}
	return nil
}

type CreateProductBatchForm []*CreateProductForm

// UpProductForm  edit Product form
type UpProductForm struct {
	Id                string                   `json:"id" form:"id"`                               // id
	ProductTypeId     string                   `json:"productTypeId" form:"productTypeId"`         // productTypeId
	ProductTypeIdPath []string                 `json:"productTypeIdPath" form:"productTypeIdPath"` // productTypeIdPath
	ProductKey        string                   `json:"productKey" form:"productKey"`               // productKey
	Name              string                   `json:"name" form:"name"`                           // name
	NameEn            string                   `json:"nameEn" form:"nameEn"`                       // nameEn
	Identifier        string                   `json:"identifier" form:"identifier"`               // identifier
	Model             string                   `json:"model" form:"model"`                         // model
	ImageUrl          string                   `json:"imageUrl" form:"imageUrl"`                   // imageUrl
	WifiFlag          string                   `json:"wifiFlag" form:"wifiFlag"`                   // wifiFlag
	NetworkType       int32                    `json:"networkType" form:"networkType"`             // networkType
	AttributeType     int32                    `json:"attributeType" form:"attributeType"`         // attributeType
	Status            int32                    `json:"status" form:"status"`                       // status
	IsVirtualTest     int32                    `json:"isVirtualTest" form:"isVirtualTest"`         // isVirtualTest
	Remark            string                   `json:"remark" form:"remark"`                       // remark
	CreatedBy         int64                    `json:"createdBy" form:"createdBy"`                 // createdBy
	UpdatedBy         int64                    `json:"updatedBy" form:"updatedBy"`                 // updatedBy
	DeletedAt         int32                    `json:"deletedAt" form:"deletedAt"`                 // deletedAt
	IsPublish         int32                    `json:"isPublish" form:"isPublish"`                 // isPublish
	ModelId           string                   `json:"modelId,omitempty" form:"modelId"`           // modelId
	ThingModels       []*TPmThingModelVo       `json:"thingModels,omitempty"`                      //物模型列表
	ControlPanelIds   []string                 `json:"controlPanelIds,omitempty"`                  //控制面板
	ModuleIds         []string                 `json:"moduleIds,omitempty"`                        //模组
	FirmwareIds       []string                 `json:"firmwareIds,omitempty"`                      //固件
	PowerConsumeType  int32                    `json:"powerConsumeType"`
	NetworkGuides     []*PmNetworkGuideEntitys `json:"networkGuides,omitempty"`
	Step              int32                    `json:"step"` //保存步骤
}

// Valid  edit Product form verify
func (a *UpProductForm) Valid() (err error) {
	if iotutil.ToInt64(a.Id) <= 0 {
		err = fmt.Errorf("%s 不能为空", "id")
		return
	}
	return
}

// TPmProductVo mapped from table <t_pm_product>
type TPmProductVo struct {
	Id                string                   `json:"id,omitempty"`                               // 主键（雪花算法19位）
	ProductTypeId     string                   `json:"productTypeId,omitempty"`                    // 产品类型ID
	ProductTypeIdPath []string                 `json:"productTypeIdPath" form:"productTypeIdPath"` // productTypeIdPath
	ProductKey        string                   `json:"productKey,omitempty"`                       // 产品唯一标识
	Name              string                   `json:"name,omitempty"`                             // 产品名称
	NameEn            string                   `json:"nameEn,omitempty"`                           // 产品名称（英文）
	Identifier        string                   `json:"identifier,omitempty"`                       // 属性的标识符。可包含大小写英文字母、数字、下划线（_），长度不超过50个字符。
	Model             string                   `json:"model,omitempty"`                            // 产品型号
	ImageURL          string                   `json:"imageUrl,omitempty"`                         // 产品图片
	WifiFlag          string                   `json:"wifiFlag,omitempty"`                         // WIFI标识
	NetworkType       int32                    `json:"networkType,omitempty"`                      // 通信协议（WIFI, BLE, WIFI+BLE）
	NetworkTypeName   string                   `json:"networkTypeName,omitempty"`                  // 通信协议（WIFI, BLE, WIFI+BLE）
	AttributeType     int32                    `json:"attributeType"`                              // 设备性质（0:普通设备，1：网关设备）
	Status            int32                    `json:"status"`                                     // 状态（0：未发布，1：已发布，2：停用）
	IsVirtualTest     int32                    `json:"isVirtualTest"`                              // 是否支持虚拟测试（0：否，1：是）
	Remark            string                   `json:"remark" form:"remark"`                       // remark
	CreatedBy         int64                    `json:"createdBy,omitempty"`                        // 创建人
	CreatedAt         string                   `json:"createdAt,omitempty"`                        // 创建时间
	UpdatedBy         int64                    `json:"updatedBy,omitempty"`                        // 修改人
	UpdatedAt         string                   `json:"updatedAt,omitempty"`                        // 修改时间
	DeletedAt         int32                    `json:"deletedAt,omitempty"`                        // 删除的标识 0-正常 1-删除
	ProductTypeName   string                   `json:"productTypeName,omitempty"`                  // 产品分类名称
	ModelId           string                   `json:"modelId,omitempty" form:"modelId"`           // modelId
	ThingModels       []*TPmThingModelVo       `json:"thingModels,omitempty"`                      //物模型列表
	ControlPanelIds   []string                 `json:"controlPanelIds,omitempty"`                  //控制面板
	ModuleIds         []string                 `json:"moduleIds,omitempty"`                        //模组
	FirmwareIds       []string                 `json:"firmwareIds,omitempty"`                      //固件
	PowerConsumeType  int32                    `json:"powerConsumeType"`
	NetworkGuides     []*PmNetworkGuideEntitys `json:"networkGuides,omitempty"`
}

// QueryThingModelForm
type QueryThingModelForm struct {
	ProductKey string `gorm:"column:product_key;not null" json:"productKey"` // 产品唯一标识
	Version    string `gorm:"column:version;not null" json:"version"`        // 物模型版本号
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
	Dpid        int32  `json:"dpid"`
}

// 数值型DataSpecs（INT & FLOAT & DOUBLE）
// 示例：{"dataType":"INT","max":"10000000","min":"0","step":"1","unit":"mg"}
type ValueDataSpecs struct {
	DataType     string      `json:"dataType,omitempty"`     // 取值为INT、FLOAT或DOUBLE。
	Max          interface{} `json:"max,omitempty"`          // 最大值。取值为INT、FLOAT或DOUBLE，必须与dataType设置一致
	Min          interface{} `json:"min,omitempty"`          // 最小值。取值为INT、FLOAT或DOUBLE，必须与dataType设置一致
	Step         interface{} `json:"step,omitempty"`         // 步长，数据每次变化的增量。取值为INT、FLOAT或DOUBLE，必须与dataType设置一致
	Precise      string      `json:"precise,omitempty"`      // 精度。当dataType取值为FLOAT或DOUBLE时，可传入的参数
	DefaultValue string      `json:"defaultValue,omitempty"` // 传入此参数，可存入一个默认值
	Unit         interface{} `json:"unit,omitempty"`         // 单位的符号
	UnitName     string      `json:"unitName,omitempty"`     // 单位的名称
	Custom       int32       `json:"custom,omitempty"`       // 是否是自定义功能。1：是 0：否
}

// 字符型DataSpecs（TEXT & DATE & JSON）
type StringDataSpecs struct {
	DataType     string      `json:"dataType,omitempty"`     // 取值为DATE或TEXT。
	Length       interface{} `json:"length,omitempty"`       // 数据长度，取值不能超过2048，单位：字节。dataType取值为TEXT时，需传入该参数。
	DefaultValue interface{} `json:"defaultValue,omitempty"` // 传入此参数，可存入一个默认值。
	Custom       int32       `json:"custom,omitempty"`       // 是否是自定义功能。1：是 0：否
}

// 枚举型DataSpaceList（ENUM）
type EnumDataSpaces struct {
	DataType string `json:"dataType,omitempty"` // 数据类型
	Name     string `json:"name,omitempty"`     // 名称
	Value    int64  `json:"value,omitempty"`    // 数值
	Custom   int32  `json:"custom,omitempty"`   // 是否是自定义功能。1：是 0：否
	Desc     string `json:"desc,omitempty"`     // 描述
}

// 布尔型BoolDataSpaces（BOOL）
type BoolDataSpaces struct {
	DataType string `json:"dataType,omitempty"` // 数据类型
	Name     string `json:"name,omitempty"`     // 名称
	Value    int64  `json:"value,omitempty"`    // 数值
	Custom   int32  `json:"custom,omitempty"`   // 是否是自定义功能。1：是 0：否
}

// TPmProductVo mapped from table <t_pm_product>
type TOpmThingModel struct {
	Id           string `json:"id,omitempty"`           // 主键ID
	FuncType     string `json:"funcType,omitempty"`     // 功能类型
	FuncName     string `json:"funcName,omitempty"`     // 功能名称
	Required     int32  `json:"required"`               // 必选
	Identifier   string `json:"identifier,omitempty"`   // 标识符
	RwFlag       string `json:"rwFlag,omitempty"`       // 数据传输类型
	DataType     string `json:"dataType,omitempty"`     // 数据类型
	Attribute    string `json:"attribute,omitempty"`    // 功能点属性
	Space        string `json:"space,omitempty"`        // 属性数据
	InputParams  string `json:"inputParams,omitempty"`  // 服务输入参数
	OutputParams string `json:"outputParams,omitempty"` // 服务输出参数
	Outputdata   string `json:"outputdata,omitempty"`   // 事件输出数据
	ModelId      int64  `json:"modelId,omitempty"`      // 物模型ID
	ProductKey   string `json:"productKey,omitempty"`   // 产品唯一标识
	Specs        string `json:"specs,omitempty"`        //非列表型数据
	SpecsList    string `json:"specsList,omitempty"`    //列表型数据
	Custom       int32  `json:"custom,omitempty"`       //是否是自定义功能（0:否, 1:是）
	CallType     int32  `json:"callType,omitempty"`     //服务的调用方式。1：异步调用, 0：同步调用
	EventType    string `json:"eventType,omitempty"`    // 事件类型。INFO_EVENT_TYPE：信息。ALERT_EVENT_TYPE：告警。ERROR_EVENT_TYPE：故障。
}

// TPmProductVo mapped from table <t_pm_product>
type TOpmThingModelVo struct {
	Id           string `json:"id,omitempty"`           // 主键ID
	FuncType     string `json:"funcType,omitempty"`     // 功能类型
	FuncName     string `json:"funcName,omitempty"`     // 功能名称
	Required     int32  `json:"required"`               // 必选
	Identifier   string `json:"identifier,omitempty"`   // 标识符
	RwFlag       string `json:"rwFlag,omitempty"`       // 数据传输类型
	DataType     string `json:"dataType,omitempty"`     // 数据类型
	Attribute    string `json:"attribute,omitempty"`    // 功能点属性
	Space        string `json:"space,omitempty"`        // 属性数据
	InputParams  string `json:"inputParams,omitempty"`  // 服务输入参数
	OutputParams string `json:"outputParams,omitempty"` // 服务输出参数
	Outputdata   string `json:"outputdata,omitempty"`   // 事件输出数据
	EventType    string `json:"eventType,omitempty"`    // 事件类型。INFO_EVENT_TYPE：信息。ALERT_EVENT_TYPE：告警。ERROR_EVENT_TYPE：故障。
}
