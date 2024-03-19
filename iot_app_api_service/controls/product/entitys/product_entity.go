package entitys

import (
	"cloud_platform/iot_app_api_service/controls"
	proto "cloud_platform/iot_proto/protos/protosService"
)

// AppProductVo mapped from table <t_opm_product>
type AppProductVo struct {
	Id              int64  `json:"id,omitempty"`                                    // 主键ID
	ProductTypeId   int64  `json:"type"`                                            // 产品品类ID
	ProductTypeName string `gorm:"column:product_type_name" json:"productTypeName"` // 产品分类名称
	Name            string `json:"name"`                                            // 产品名称
	NameEn          string `json:"nameEn"`                                          // 产品名称（英文）
	Model           string `json:"model"`                                           // 产品型号
	ImageUrl        string `json:"listimg"`                                         // 产品图片
	WifiFlag        string `json:"wifiName"`                                        // WIFI标识
	NetworkType     string `json:"networkMode"`                                     // 通信协议（WIFI, BLE, WIFI+BLE）
	IsVirtualTest   int32  `json:"isVirtualTest"`                                   // 是否支持虚拟测试（0：否，1：是）
	Token           string `json:"token"`                                           //token
}

// AppQueryProductForm query OpmProduct  form ;  if some field is required, create binding:"required" to tag by self
type AppQueryProductForm struct {
	ProductTypeId   int64    `json:"productTypeId,string"`                            // 产品品类ID
	BaseProductId   int64    `json:"baseProductId,string"`                            // 产品类型编号
	ProductTypeName string   `gorm:"column:product_type_name" json:"productTypeName"` // 产品分类名称
	WifiFlag        string   `json:"wifiName"`                                        // WIFI标识
	WifiFlags       []string `json:"modelList"`                                       // WIFI标识集合
	Order           int      `json:"order" form:"order"`                              // example: orderMap[column]=desc
	Page            int      `json:"page" form:"page"`                                //
	Limit           int      `json:"limit" form:"limit"`                              //
	AppKey          string   `json:"appKey"`
	TenantId        string   `json:"tenantId"`
}

// 增、删、改及查询返回
type OpmNetworkGuideEntitys struct {
	Id          int64                         `json:"id,string"`
	ProductId   int64                         `json:"productId,string"`
	Type        int32                         `json:"type"`        //配网方式
	NetworkMode int32                         `json:"networkMode"` //通讯协议
	WifiFlag    string                        `json:"wifiFlag"`    //
	ProductKey  string                        `json:"productKey"`
	Steps       []*OpmNetworkGuideStepEntitys `json:"steps"`
}

// 增、删、改及查询返回
type OpmNetworkGuideStepEntitys struct {
	Id             int64  `json:"id,string"`
	NetworkGuideId int64  `json:"networkGuideId,string"`
	Instruction    string `json:"instruction"`
	//InstructionEn  string `json:"instructionEn"`
	ImageUrl string `json:"imageUrl"`
	VideoUrl string `json:"videoUrl"`
	Sort     int32  `json:"sort"`
}

// pb对象转实体
func OpmNetworkGuideStep_pb2e(src *proto.OpmNetworkGuideStep, language string) *OpmNetworkGuideStepEntitys {
	if src == nil {
		return nil
	}
	var instruction string
	if language == "en" {
		instruction = src.InstructionEn
	} else {
		instruction = src.Instruction
	}

	entitysObj := OpmNetworkGuideStepEntitys{
		Id:             src.Id,
		NetworkGuideId: src.NetworkGuideId,
		Instruction:    instruction,
		//InstructionEn:  src.InstructionEn,
		ImageUrl: src.ImageUrl,
		VideoUrl: src.VideoUrl,
		Sort:     src.Sort,
	}
	return &entitysObj
}

type AppProductDto struct {
	Id               string `json:"id,omitempty"` // 主键ID
	ProductTypeId    string `json:"type"`         // 产品品类ID
	Name             string `json:"name"`         // 产品名称
	Model            string `json:"productKey"`   // 产品型号
	ImageUrl         string `json:"imageUrl"`     // 产品图片
	WifiFlag         string `json:"wifiFlag"`     // WIFI标识
	NetworkType      int32  `json:"networkMode"`  // 通信协议（WIFI, BLE, WIFI+BLE）
	Token            string `json:"token"`        //token
	DistributionType int32  `json:"networkType"`  //配网类型 请替换为NetworkType
}

// 增、删、改及查询返回
type OpmProductEntitys struct {
	Id                   int64  `json:"id,string"`
	ProductTypeId        int64  `json:"productTypeId,string"`
	ProductKey           string `json:"productKey"`
	Name                 string `json:"name"`
	NameEn               string `json:"nameEn"`
	Identifier           string `json:"identifier"`
	Model                string `json:"model"`
	ImageUrl             string `json:"imageUrl"`
	WifiFlag             string `json:"wifiFlag"`
	NetworkType          int32  `json:"networkType"`
	NetworkTypeDesc      string `json:"networkTypeDesc"`
	AttributeType        int32  `json:"attributeType"`
	PowerConsumeType     int32  `json:"powerConsumeType"`
	PowerConsumeTypeDesc string `json:"powerConsumeTypeDesc"`
	Status               int32  `json:"status"`
	StatusDesc           string `json:"statusDesc"`
	IsVirtualTest        int32  `json:"isVirtualTest"`
	IsVirtualTestName    string `json:"isVirtualTestName"`
	StatusName           string `json:"statusName"`
	IsScheme             int32  `json:"isScheme"`
	Desc                 string `json:"desc"`
	ProductTypeName      string `json:"productTypeName"`
	TenantId             string `json:"tenantId"`
	SchemeId             int64  `json:"schemeId,string"`
	ControlPanelId       int64  `json:"controlPanelId,string"`
	ModuleId             int64  `json:"moduleId,string"`
	FirmwareId           int64  `json:"firmwareId,string,omitempty"`
	FirmwareVersionId    int64  `json:"firmwareVersionId,string,omitempty"`
	FirmwareVersion      string `json:"firmwareVersion,string,omitempty"`
	UpdatedAt            int64  `json:"updatedAt"`
}

// 查询条件
type OpmProductQuery struct {
	Page       uint64            `json:"page,omitempty"`
	Limit      uint64            `json:"limit,omitempty"`
	Sort       string            `json:"sort,omitempty"`
	SortField  string            `json:"sortField,omitempty"`
	SearchKey  string            `json:"searchKey,omitempty"`
	IsPlatform bool              `json:"isPlatform"`
	Query      *OpmProductFilter `json:"query,omitempty"`
}

// OpmProductFilter，查询条件，字段请根据需要自行增减
type OpmProductFilter struct {
	Id                int64  `json:"id,string,omitempty"`
	ProductTypeId     int64  `json:"productTypeId,string,omitempty"`
	ProductKey        string `json:"productKey,omitempty"`
	Name              string `json:"name,omitempty"`
	NameEn            string `json:"nameEn,omitempty"`
	Identifier        string `json:"identifier,omitempty"`
	Model             string `json:"model,omitempty"`
	ImageUrl          string `json:"imageUrl,omitempty"`
	WifiFlag          string `json:"wifiFlag,omitempty"`
	NetworkType       int32  `json:"networkType,omitempty"`
	AttributeType     int32  `json:"attributeType,omitempty"`
	PowerConsumeType  int32  `json:"powerConsumeType,omitempty"`
	Status            int32  `json:"status,omitempty"`
	IsVirtualTest     int32  `json:"isVirtualTest,omitempty"`
	IsScheme          int32  `json:"isScheme,omitempty"`
	Desc              string `json:"desc,omitempty"`
	ProductTypeName   string `json:"productTypeName,omitempty"`
	TenantId          string `json:"tenantId,omitempty"`
	ProductId         int64  `json:"productId,string,omitempty"`
	ControlPanelId    int64  `json:"controlPanelId,string,omitempty"`
	ModuleId          int64  `json:"moduleId,string,omitempty"`
	FirmwareVersionId int64  `json:"firmwareVersionId,string,omitempty"`
	FirmwareId        int64  `json:"firmwareId,string,omitempty"`
	FirmwareVersion   string `json:"firmwareVersion,omitempty"`
}

// pb对象转实体
func OpmProduct_pb2e(src *proto.OpmProduct) *OpmProductEntitys {
	if src == nil {
		return nil
	}

	entitysObj := OpmProductEntitys{
		Id:               src.Id,
		ProductTypeId:    src.ProductTypeId,
		ProductKey:       src.ProductKey,
		Name:             src.Name,
		Identifier:       src.Identifier,
		Model:            src.Model,
		ImageUrl:         src.ImageUrl,
		WifiFlag:         src.WifiFlag,
		AttributeType:    src.AttributeType,
		PowerConsumeType: src.PowerConsumeType,
		Status:           src.Status,
		IsVirtualTest:    src.IsVirtualTest,
		IsScheme:         src.IsScheme,
		Desc:             src.Desc,
		ProductTypeName:  src.ProductTypeName,
		TenantId:         src.TenantId,
		SchemeId:         src.ProductId,
		ControlPanelId:   src.ControlPanelId,
		ModuleId:         src.ModuleId,
		UpdatedAt:        src.UpdatedAt.AsTime().Unix(),
		NetworkType:      src.NetworkType,
	}
	entitysObj.ImageUrl = controls.ConvertProImg(src.ImageUrl)

	return &entitysObj
}

type ThingModelRuleItem struct {
	DpId       int64       `json:"dpId"`
	Identifier string      `json:"identifier"`
	Operate    int32       `json:"operate"`
	Value      interface{} `json:"value"`
}

type ThingModelRuleItemResponse struct {
	DpId          int64                `json:"dpId"`
	Identifier    string               `json:"identifier"`
	Operate       int32                `json:"operate"`
	Value         interface{}          `json:"value"`
	ConditionType int32                `json:"conditionType"`
	IfSpecs       []ThingModelRuleItem `json:"ifSpecs,omitempty"`
	Specs         []ThingModelRuleItem `json:"specs,omitempty"`
	Sort          int64                `json:"sort"` //排序
}
