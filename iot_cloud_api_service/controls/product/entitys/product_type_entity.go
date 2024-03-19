package entitys

import (
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"strconv"
)

// ProductType table
type ProductType struct {
	Id         string `gorm:"PRIMARY_KEY"`
	Name       string `gorm:"Column:name"`
	NameEn     string `gorm:"Column:name_en"`
	Identifier string `gorm:"Column:Identifier"`
	Sort       int16  `gorm:"Column:sort"`
	ParentId   string `gorm:"Column:parent_id"`
	Desc       string `gorm:"Column:desc"`
	CreatedBy  string `gorm:"Column:created_by"`
	CreatedAt  string `gorm:"Column:created_at"`
	UpdatedBy  string `gorm:"Column:updated_by"`
	UpdatedAt  string `gorm:"Column:updated_at"`
	DeletedAt  string `gorm:"Column:deleted_at"`
}

// QueryProductTypeForm query ProductType  form ;  if some field is required, create binding:"required" to tag by self
type QueryProductTypeForm struct {
	Id          int64        `json:"id" form:"id"`                 // cond Id
	Name        string       `json:"name" form:"name"`             // cond Name
	NameEn      string       `json:"nameEn" form:"nameEn"`         // cond NameEn
	Identifier  string       `json:"identifier" form:"identifier"` // cond Identifier
	Sort        int          `json:"sort" form:"sort"`             // example: SortMap[>]=some value&SortMap[<]=some value; key must be ">,>=,<,<=,="
	ParentId    int64        `json:"parentId" form:"parentId"`     // cond ParentId
	Desc        string       `json:"desc" form:"desc"`             // cond Desc
	ImgSize     int          `json:"imgSize"`                      //added in 1.0.3
	ImgPath     string       `json:"imgPath"`                      //added in 1.0.3
	ImgFullPath string       `json:"imgFullPath"`                  //added in 1.0.3
	ImgName     string       `json:"imgName"`                      //added in 1.0.3
	ImgType     string       `json:"imgType"`                      //added in 1.0.3
	Imgkey      string       `json:"Imgkey"`                       //added in 1.0.3
	ModelItems  []ModelsItem `json:"modelItems"`                   //added in 1.0.3
	CreatedBy   string       `json:"createdBy" form:"createdBy"`   // cond CreatedBy
	UpdatedBy   string       `json:"updatedBy" form:"updatedBy"`   // cond UpdatedBy
	DeletedAt   string       `json:"deletedAt" form:"deletedAt"`   // example: DeletedAtMap[>]=some value&DeletedAtMap[<]=some value; key must be ">,>=,<,<=,="
	Order       int          `json:"order" form:"order"`           // example: orderMap[column]=desc
	Page        int          `json:"page" form:"page"`             // get all without uploading
	Limit       int          `json:"limit" form:"page"`            // get all without uploading
	SearchKey   string       `json:"searchKey" form:"searchKey"`   // Search key
}

// CreateProductTypeForm create ProductType form
type CreateProductTypeForm struct {
	Id          string       `json:"id" form:"id"`                 // id
	Name        string       `json:"name" form:"name"`             // name
	NameEn      string       `json:"nameEn" form:"nameEn"`         // nameEn
	Identifier  string       `json:"identifier" form:"identifier"` // identifier
	Sort        int          `json:"sort" form:"sort"`             // sort
	ParentId    string       `json:"parentId" form:"parentId"`     // parentId
	Desc        string       `json:"desc" form:"desc"`             // desc
	ImgSize     int          `json:"imgSize"`                      //added in 1.0.3
	ImgPath     string       `json:"imgPath"`                      //added in 1.0.3
	ImgFullPath string       `json:"imgFullPath"`                  //added in 1.0.3
	ImgName     string       `json:"imgName"`                      //added in 1.0.3
	ImgType     string       `json:"imgType"`                      //added in 1.0.3
	Imgkey      string       `json:"Imgkey"`                       //added in 1.0.3
	ModelItems  []ModelsItem `json:"modelItems"`                   //added in 1.0.3
	CreatedBy   string       `json:"createdBy" form:"createdBy"`   // createdBy
	UpdatedBy   string       `json:"updatedBy" form:"updatedBy"`   // updatedBy
	DeletedAt   string       `json:"deletedAt" form:"deletedAt"`   // deletedAt
}

// Valid create ProductType  form verify
func (a *CreateProductTypeForm) Valid() (err error) {
	return
}

func (a *CreateProductTypeForm) ToPB() (*protosService.TPmProductTypeRequest, error) {
	var tid int64
	var ParentId int64
	tid = iotutil.GetNextSeqInt64()
	if len(a.ParentId) > 0 {
		n, err := strconv.ParseInt(a.ParentId, 10, 64)
		if err != nil {
			return nil, err
		}
		ParentId = n
	}
	CreatedBy, _ := strconv.ParseInt(a.CreatedBy, 10, 64)
	UpdatedBy, _ := strconv.ParseInt(a.UpdatedBy, 10, 64)
	pbObj := protosService.TPmProductTypeRequest{
		Id:          tid,
		Name:        a.Name,
		NameEn:      a.NameEn,
		Identifier:  a.Identifier,
		Sort:        int32(a.Sort),
		ParentId:    ParentId,
		Desc:        a.Desc,
		CreatedBy:   CreatedBy,
		UpdatedBy:   UpdatedBy,
		ImgSize:     int32(a.ImgSize),
		ImgPath:     a.ImgPath,
		ImgFullPath: a.ImgFullPath,
		ImgName:     a.ImgName,
		ImgType:     a.ImgType,
		Imgkey:      a.Imgkey,
	}
	for _, v := range a.ModelItems {
		pbObj.ModelsItems = append(pbObj.ModelsItems, &protosService.ModelsItem{
			Id:            iotutil.GetNextSeqInt64(),
			ProductTypeId: tid,
			Dpid:          int32(v.Dpid),
			Identifier:    v.Identifier,
			Name:          v.Name,
			RwFlag:        v.RwFlag,
			DataType:      v.DataType,
			Properties:    v.Properties,
			Mark:          v.Mark,
			Required:      int32(v.Required),
		})
	}
	return &pbObj, nil
}

// UpProductTypeForm  edit ProductType form
type UpProductTypeForm struct {
	Id          string       `json:"id" form:"id" binding:"required"` // id
	Name        string       `json:"name" form:"name"`                // name
	NameEn      string       `json:"nameEn" form:"nameEn"`            // nameEn
	Identifier  string       `json:"identifier" form:"identifier"`    // identifier
	Sort        int          `json:"sort" form:"sort"`                // sort
	ParentId    string       `json:"parentId" form:"parentId"`        // parentId
	Desc        string       `json:"desc" form:"desc"`                // desc
	ImgSize     int          `json:"imgSize"`                         //added in 1.0.3
	ImgPath     string       `json:"imgPath"`                         //added in 1.0.3
	ImgFullPath string       `json:"imgFullPath"`                     //added in 1.0.3
	ImgName     string       `json:"imgName"`                         //added in 1.0.3
	ImgType     string       `json:"imgType"`                         //added in 1.0.3
	Imgkey      string       `json:"Imgkey"`                          //added in 1.0.3
	ModelItems  []ModelsItem `json:"modelItems"`                      //added in 1.0.3
	CreatedBy   string       `json:"createdBy" form:"createdBy"`      // createdBy
	UpdatedBy   string       `json:"updatedBy" form:"updatedBy"`      // updatedBy
	DeletedAt   string       `json:"deletedAt" form:"deletedAt"`      // deletedAt
}

// Valid  edit ProductType form verify
func (a *UpProductTypeForm) Valid() (err error) {
	return
}

func (a *UpProductTypeForm) ToPB() (*protosService.TPmProductTypeRequest, error) {
	var tid int64
	var ParentId int64
	n, err := strconv.ParseInt(a.Id, 10, 64)
	if err != nil {
		return nil, err
	}
	tid = n
	if len(a.ParentId) > 0 {
		n, err = strconv.ParseInt(a.ParentId, 10, 64)
		if err != nil {
			return nil, err
		}
		ParentId = n
	}
	UpdatedBy, _ := strconv.ParseInt(a.UpdatedBy, 10, 64)
	pbObj := protosService.TPmProductTypeRequest{
		Id:          tid,
		Name:        a.Name,
		NameEn:      a.NameEn,
		Identifier:  a.Identifier,
		Sort:        int32(a.Sort),
		ParentId:    ParentId,
		Desc:        a.Desc,
		UpdatedBy:   UpdatedBy,
		ImgSize:     int32(a.ImgSize),
		ImgPath:     a.ImgPath,
		ImgFullPath: a.ImgFullPath,
		ImgName:     a.ImgName,
		ImgType:     a.ImgType,
		Imgkey:      a.Imgkey,
	}
	var id int64
	for _, v := range a.ModelItems {
		if len(v.Id) > 0 {
			id, err = strconv.ParseInt(v.Id, 10, 64)
			if err != nil {
				return nil, err
			}
		} else {
			id = iotutil.GetNextSeqInt64()
		}
		pbObj.ModelsItems = append(pbObj.ModelsItems, &protosService.ModelsItem{
			Id:            id,
			ProductTypeId: tid,
			Dpid:          int32(v.Dpid),
			Identifier:    v.Identifier,
			Name:          v.Name,
			RwFlag:        v.RwFlag,
			DataType:      v.DataType,
			Properties:    v.Properties,
			Mark:          v.Mark,
			Required:      int32(v.Required),
		})
	}
	return &pbObj, nil
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
	CreatedBy   int64               `json:"createdBy,omitempty"`  // 创建人
	CreatedAt   string              `json:"createdAt,omitempty"`  // 创建时间
	ImgSize     int                 `json:"imgSize,omitempty"`    //图片大小，added in 1.0.3
	ImgPath     string              `json:"imgPath,omitempty"`    //图片路径，added in 1.0.3
	ImgFullPath string              `json:"imgFullPath"`          //图片完整路径，added in 1.0.3
	ImgName     string              `json:"imgName,omitempty"`    //图片名称，added in 1.0.3
	ImgType     string              `json:"imgType,omitempty"`    //图片类型，added in 1.0.3
	Imgkey      string              `json:"Imgkey,omitempty"`     //图片MD5，added in 1.0.3
	ModelItems  []*ModelsItem       `json:"modelItems,omitempty"` //分类物模型，added in 1.0.3
	UpdatedBy   int64               `json:"updatedBy,omitempty"`  // 修改人
	UpdatedAt   string              `json:"updatedAt,omitempty"`  // 修改时间
	DeletedAt   string              `json:"deletedAt,omitempty"`  // 删除的标识 0-正常 1-删除
	ParentName  string              `json:"parentName,omitempty"` //父级分类名称
	Children    []*TPmProductTypeVo `json:"children"`             //递归引用链
	Products    []*TPmProductVo     `json:"products,omitempty"`   //递归引用链-产品
	Count       int64               `json:"count"`                //关联产品类型数量
}

type ModelsItem struct {
	Id         string `json:"id"`
	Dpid       int    `json:"dpid"`
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	RwFlag     string `json:"rwFlag"`
	DataType   string `json:"dataType"`
	Properties string `json:"properties"`
	Mark       string `json:"mark"`
	Required   int    `json:"required"`
}
