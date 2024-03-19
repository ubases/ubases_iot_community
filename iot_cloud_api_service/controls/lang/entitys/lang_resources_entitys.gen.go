// Code generated by sgen.exe,2022-05-31 13:46:36. DO NOT EDIT.
// versions: v1.0.0
//  API结构体封装，请自行根据需要删减字段修改

package entitys

import (
	proto "cloud_platform/iot_proto/protos/protosService"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// 增、删、改及查询返回
type LangResourcesEntitys struct {
	Id         int64     `json:"id,omitempty"`
	BelongType int32     `json:"belongType,omitempty"`
	Lang       string    `json:"lang,omitempty"`
	Code       string    `json:"code,omitempty"`
	Value      string    `json:"value,omitempty"`
	CreatedBy  int64     `json:"createdBy,omitempty"`
	UpdatedBy  int64     `json:"updatedBy,omitempty"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty"`
}

// 新增参数非空检查
func (s *LangResourcesEntitys) AddCheck() error {
	return nil
}

// 修改参数非空检查
func (s *LangResourcesEntitys) UpdateCheck() error {
	return nil
}

// 查询参数必填检查
func (*LangResourcesQuery) QueryCheck() error {
	return nil
}

// 查询条件
type LangResourcesQuery struct {
	Page      uint64               `json:"page,omitempty"`
	Limit     uint64               `json:"limit,omitempty"`
	Sort      string               `json:"sort,omitempty"`
	SortField string               `json:"sortField,omitempty"`
	SearchKey string               `json:"searchKey,omitempty"`
	Query     *LangResourcesFilter `json:"query,omitempty"`
}

// LangResourcesFilter，查询条件，字段请根据需要自行增减
type LangResourcesFilter struct {
	Id         int64     `json:"id,omitempty"`
	BelongType int32     `json:"belongType,omitempty"`
	Lang       string    `json:"lang,omitempty"`
	Code       string    `json:"code,omitempty"`
	Value      string    `json:"value,omitempty"`
	CreatedBy  int64     `json:"createdBy,omitempty"`
	UpdatedBy  int64     `json:"updatedBy,omitempty"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty"`
}

// 实体转pb对象
func LangResources_e2pb(src *LangResourcesEntitys) *proto.LangResources {
	if src == nil {
		return nil
	}
	pbObj := proto.LangResources{
		Id:         src.Id,
		BelongType: src.BelongType,
		Lang:       src.Lang,
		Code:       src.Code,
		Value:      src.Value,
		CreatedBy:  src.CreatedBy,
		UpdatedBy:  src.UpdatedBy,
		CreatedAt:  timestamppb.New(src.CreatedAt),
		UpdatedAt:  timestamppb.New(src.UpdatedAt),
	}
	return &pbObj
}

// pb对象转实体
func LangResources_pb2e(src *proto.LangResources) *LangResourcesEntitys {
	if src == nil {
		return nil
	}
	entitysObj := LangResourcesEntitys{
		Id:         src.Id,
		BelongType: src.BelongType,
		Lang:       src.Lang,
		Code:       src.Code,
		Value:      src.Value,
		CreatedBy:  src.CreatedBy,
		UpdatedBy:  src.UpdatedBy,
		CreatedAt:  src.CreatedAt.AsTime(),
		UpdatedAt:  src.UpdatedAt.AsTime(),
	}
	return &entitysObj
}
