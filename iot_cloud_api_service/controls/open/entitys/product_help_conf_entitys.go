// Code generated by sgen.exe,2022-08-18 20:09:05. DO NOT EDIT.
// versions: v1.0.0
//  API结构体封装，请自行根据需要删减字段修改

package entitys

import (
	"cloud_platform/iot_common/iotutil"
	proto "cloud_platform/iot_proto/protos/protosService"
	"encoding/json"
	"errors"
	"sort"
)

type Lang struct {
	Lang     string `json:"lang"`
	LangName string `json:"langName"`
}

// 增、删、改及查询返回
type ProductHelpConfEntitys struct {
	Id            string `json:"id,omitempty"`
	TenantId      string `json:"tenantId,omitempty"`
	ProductName   string `json:"productName,omitempty"`
	ProductKey    string `json:"productKey,omitempty"`
	ProductTypeId string `json:"productTypeId,omitempty"`
	Langs         []Lang `json:"langs,omitempty"`
	RemainLang    string `json:"remainLang,omitempty"`
	Status        int32  `json:"status,omitempty"`
	CreatedBy     int64  `json:"createdBy,omitempty"`
	CreatedAt     int64  `json:"createdAt,omitempty"`
	UpdatedBy     int64  `json:"updatedBy,omitempty"`
	UpdatedAt     int64  `json:"updatedAt,omitempty"`
}

// 新增参数非空检查
func (s *ProductHelpConfEntitys) AddCheck() error {
	return nil
}

// 修改参数非空检查
func (s *ProductHelpConfEntitys) UpdateCheck() error {
	return nil
}

// 查询参数必填检查
func (*ProductHelpConfQuery) QueryCheck() error {
	return nil
}

// 查询条件
type ProductHelpConfQuery struct {
	Page      uint64                 `json:"page,omitempty"`
	Limit     uint64                 `json:"limit,omitempty"`
	Sort      string                 `json:"sort,omitempty"`
	SortField string                 `json:"sortField,omitempty"`
	SearchKey string                 `json:"searchKey,omitempty"`
	Query     *ProductHelpConfFilter `json:"query,omitempty"`
}

// ProductHelpConfFilter，查询条件，字段请根据需要自行增减
type ProductHelpConfFilter struct {
	Id            string `json:"id,omitempty"`
	TenantId      string `json:"tenantId,omitempty"`
	ProductName   string `json:"productName,omitempty"`
	ProductKey    string `json:"productKey,omitempty"`
	ProductTypeId string `json:"productTypeId,omitempty"`
	Langs         string `json:"langs,omitempty"`
	RemainLang    string `json:"remainLang,omitempty"`
	Status        int32  `json:"status,omitempty"`
	CreatedBy     int64  `json:"createdBy,omitempty"`
	CreatedAt     int64  `json:"createdAt,omitempty"`
	UpdatedBy     int64  `json:"updatedBy,omitempty"`
	UpdatedAt     int64  `json:"updatedAt,omitempty"`
}

// 实体转pb对象
func ProductHelpConf_e2pb(src *ProductHelpConfEntitys) (*proto.ProductHelpConf, error) {
	if src == nil {
		return nil, errors.New("产品配置参数为nil")
	}
	langsBytes, err := json.Marshal(src.Langs)
	if err != nil {
		return nil, err
	}
	pbObj := proto.ProductHelpConf{
		Id:            iotutil.ToInt64(src.Id),
		TenantId:      src.TenantId,
		ProductName:   src.ProductName,
		ProductKey:    src.ProductKey,
		ProductTypeId: iotutil.ToInt64(src.ProductTypeId),
		Langs:         string(langsBytes),
		RemainLang:    src.RemainLang,
		Status:        src.Status,
		CreatedBy:     src.CreatedBy,
		UpdatedBy:     src.UpdatedBy,
	}
	return &pbObj, nil
}

// pb对象转实体
func ProductHelpConf_pb2e(src *proto.ProductHelpConf) (*ProductHelpConfEntitys, error) {
	if src == nil {
		return nil, errors.New("产品配置参数为nil")
	}
	var langs []Lang
	if src.Langs != "" {
		if err := json.Unmarshal([]byte(src.Langs), &langs); err != nil {
			return nil, err
		}
		sort.Slice(langs, func(i, j int) bool {
			return langs[i].Lang > langs[j].Lang
		})
	} else {
		langs = []Lang{}
	}
	entitysObj := ProductHelpConfEntitys{
		Id:            iotutil.ToString(src.Id),
		TenantId:      src.TenantId,
		ProductName:   src.ProductName,
		ProductKey:    src.ProductKey,
		ProductTypeId: iotutil.ToString(src.ProductTypeId),
		Langs:         langs,
		RemainLang:    src.RemainLang,
		Status:        src.Status,
		CreatedBy:     src.CreatedBy,
		CreatedAt:     src.CreatedAt.AsTime().Unix(),
		UpdatedBy:     src.UpdatedBy,
		UpdatedAt:     src.UpdatedAt.AsTime().Unix(),
	}
	return &entitysObj, nil
}
