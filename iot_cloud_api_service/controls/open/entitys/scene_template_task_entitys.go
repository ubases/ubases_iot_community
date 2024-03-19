// Code generated by sgen.exe,2022-11-09 16:50:48. DO NOT EDIT.
// versions: v1.0.0
//  API结构体封装，请自行根据需要删减字段修改

package entitys

import (
	proto "cloud_platform/iot_proto/protos/protosService"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// 增、删、改及查询返回
type SceneTemplateTaskEntitys struct {
	Id              int64     `json:"id,string,omitempty"`
	SceneTemplateId int64     `json:"sceneTemplateId,string,omitempty"`
	ProductId       int64     `json:"productId,string,omitempty"`
	ProductKey      string    `json:"productKey,omitempty"`
	Functions       string    `json:"functions,omitempty"`
	FuncKey         string    `json:"funcKey,omitempty"`
	FuncValue       string    `json:"funcValue,omitempty"`
	Sort            int32     `json:"sort,omitempty"`
	CreatedBy       int64     `json:"createdBy,string,omitempty"`
	CreatedAt       time.Time `json:"createdAt,omitempty"`
}

// 新增参数非空检查
func (s *SceneTemplateTaskEntitys) AddCheck() error {
	return nil
}

// 修改参数非空检查
func (s *SceneTemplateTaskEntitys) UpdateCheck() error {
	return nil
}

// 查询参数必填检查
func (*SceneTemplateTaskQuery) QueryCheck() error {
	return nil
}

// 查询条件
type SceneTemplateTaskQuery struct {
	Page      uint64                   `json:"page,omitempty"`
	Limit     uint64                   `json:"limit,omitempty"`
	Sort      string                   `json:"sort,omitempty"`
	SortField string                   `json:"sortField,omitempty"`
	SearchKey string                   `json:"searchKey,omitempty"`
	Query     *SceneTemplateTaskFilter `json:"query,omitempty"`
}

// SceneTemplateTaskFilter，查询条件，字段请根据需要自行增减
type SceneTemplateTaskFilter struct {
	Id              int64     `json:"id,string,omitempty"`
	SceneTemplateId int64     `json:"sceneTemplateId,string,omitempty"`
	ProductId       int64     `json:"productId,string,omitempty"`
	ProductKey      string    `json:"productKey,omitempty"`
	Functions       string    `json:"functions,omitempty"`
	FuncKey         string    `json:"funcKey,omitempty"`
	FuncValue       string    `json:"funcValue,omitempty"`
	Sort            int32     `json:"sort,omitempty"`
	CreatedBy       int64     `json:"createdBy,string,omitempty"`
	CreatedAt       time.Time `json:"createdAt,omitempty"`
}

// 实体转pb对象
func SceneTemplateTask_e2pb(src *SceneTemplateTaskEntitys) *proto.SceneTemplateTask {
	if src == nil {
		return nil
	}
	pbObj := proto.SceneTemplateTask{
		Id:              src.Id,
		SceneTemplateId: src.SceneTemplateId,
		ProductId:       src.ProductId,
		ProductKey:      src.ProductKey,
		Functions:       src.Functions,
		Sort:            src.Sort,
		CreatedBy:       src.CreatedBy,
		CreatedAt:       timestamppb.New(src.CreatedAt),
	}
	return &pbObj
}

// pb对象转实体
func SceneTemplateTask_pb2e(src *proto.SceneTemplateTask) *SceneTemplateTaskEntitys {
	if src == nil {
		return nil
	}
	entitysObj := SceneTemplateTaskEntitys{
		Id:              src.Id,
		SceneTemplateId: src.SceneTemplateId,
		ProductId:       src.ProductId,
		ProductKey:      src.ProductKey,
		Functions:       src.Functions,
		Sort:            src.Sort,
		CreatedBy:       src.CreatedBy,
		CreatedAt:       src.CreatedAt.AsTime(),
	}
	return &entitysObj
}
