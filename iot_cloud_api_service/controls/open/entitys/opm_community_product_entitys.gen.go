// Code generated by sgen,2023-07-10 11:45:46. DO NOT EDIT.
// versions: v1.0.0
//  API结构体封装，请自行根据需要删减字段修改

package entitys

import (
	proto "cloud_platform/iot_proto/protos/protosService"
)

// 增、删、改及查询返回
type OpmCommunityProductEntitys struct {
	Id          int64                                 `json:"id,string,omitempty"`
	TenantId    string                                `json:"tenantId,omitempty"`
	ImageUrl    string                                `json:"imageUrl,omitempty"`
	ProductPage string                                `json:"productPage,omitempty"`
	ProductName string                                `json:"productName,omitempty"`
	ProductDesc string                                `json:"productDesc,omitempty"`
	Sort        int32                                 `json:"sort,omitempty"`
	Status      int32                                 `json:"status,omitempty"`
	CreatedBy   int64                                 `json:"createdBy,string,omitempty"`
	CreatedAt   int64                                 `json:"createdAt,omitempty"`
	UpdatedBy   int64                                 `json:"updatedBy,string,omitempty"`
	UpdatedAt   int64                                 `json:"updatedAt,omitempty"`
	Langs       []*OpmCommunityProductLanguageEntitys `json:"langs"`
}

// 新增参数非空检查
func (s *OpmCommunityProductEntitys) AddCheck() error {
	return nil
}

// 修改参数非空检查
func (s *OpmCommunityProductEntitys) UpdateCheck() error {
	return nil
}

// 查询参数必填检查
func (*OpmCommunityProductQuery) QueryCheck() error {
	return nil
}

// 查询条件
type OpmCommunityProductQuery struct {
	Page      uint64                     `json:"page,omitempty"`
	Limit     uint64                     `json:"limit,omitempty"`
	Sort      string                     `json:"sort,omitempty"`
	SortField string                     `json:"sortField,omitempty"`
	SearchKey string                     `json:"searchKey,omitempty"`
	Query     *OpmCommunityProductFilter `json:"query,omitempty"`
}

// OpmCommunityProductFilter，查询条件，字段请根据需要自行增减
type OpmCommunityProductFilter struct {
	Id          int64  `json:"id,string,omitempty"`
	TenantId    string `json:"tenantId,omitempty"`
	ImageUrl    string `json:"imageUrl,omitempty"`
	ProductPage string `json:"productPage,omitempty"`
	ProductName string `json:"productName,omitempty"`
	ProductDesc string `json:"productDesc,omitempty"`
	Sort        int32  `json:"sort,omitempty"`
	Status      int32  `json:"status,omitempty"`
}

// 实体转pb对象
func OpmCommunityProduct_e2pb(src *OpmCommunityProductEntitys) *proto.OpmCommunityProduct {
	if src == nil {
		return nil
	}
	pbObj := proto.OpmCommunityProduct{
		Id:          src.Id,
		TenantId:    src.TenantId,
		ImageUrl:    src.ImageUrl,
		ProductPage: src.ProductPage,
		ProductName: src.ProductName,
		ProductDesc: src.ProductDesc,
		Sort:        src.Sort,
		Status:      src.Status,
		CreatedBy:   src.CreatedBy,
		UpdatedBy:   src.UpdatedBy,
	}
	if src.Langs != nil {
		pbObj.Langs = make([]*proto.OpmCommunityProductLanguage, 0)
		for _, lang := range src.Langs {
			pbObj.Langs = append(pbObj.Langs, OpmCommunityProductLanguage_e2pb(lang))
		}
	}
	return &pbObj
}

// pb对象转实体
func OpmCommunityProduct_pb2e(src *proto.OpmCommunityProduct) *OpmCommunityProductEntitys {
	if src == nil {
		return nil
	}
	entitysObj := OpmCommunityProductEntitys{
		Id:          src.Id,
		TenantId:    src.TenantId,
		ImageUrl:    src.ImageUrl,
		ProductPage: src.ProductPage,
		ProductName: src.ProductName,
		ProductDesc: src.ProductDesc,
		Sort:        src.Sort,
		Status:      src.Status,
		CreatedBy:   src.CreatedBy,
		UpdatedBy:   src.UpdatedBy,
		CreatedAt:   src.CreatedAt.AsTime().Unix(),
		UpdatedAt:   src.UpdatedAt.AsTime().Unix(),
	}
	if src.Langs != nil {
		entitysObj.Langs = make([]*OpmCommunityProductLanguageEntitys, 0)
		for _, lang := range src.Langs {
			entitysObj.Langs = append(entitysObj.Langs, OpmCommunityProductLanguage_pb2e(lang))
		}
	}
	return &entitysObj
}

// 查询条件转换proto
func (s *OpmCommunityProductQuery) OpmCommunityProductQuery_e2pb() *proto.OpmCommunityProduct {
	queryObj := &proto.OpmCommunityProduct{}
	if s.Query != nil {
		queryObj.Status = s.Query.Status
		queryObj.ProductName = s.Query.ProductName
	}
	return queryObj
}
