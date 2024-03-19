// Code generated by sgen.exe,2022-10-24 09:40:34. DO NOT EDIT.
// versions: v1.0.0
//  API结构体封装，请自行根据需要删减字段修改

package entitys

import (
	proto "cloud_platform/iot_proto/protos/protosService"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// 增、删、改及查询返回
type PublicAppVersionEntitys struct {
	Id            int64     `json:"id,string,omitempty"`
	AppKey        string    `json:"appKey,omitempty"`
	ReleaseTime   time.Time `json:"releaseTime,omitempty"`
	ReleaseMarket int32     `json:"releaseMarket,omitempty"`
	Version       string    `json:"version,omitempty"`
	AppType       int32     `json:"appType,omitempty"`
	Status        int32     `json:"status,omitempty"`
	Remark        string    `json:"remark,omitempty"`
	CreatedBy     int64     `json:"createdBy,string,omitempty"`
	CreatedAt     time.Time `json:"createdAt,omitempty"`
	UpdatedBy     int64     `json:"updatedBy,string,omitempty"`
	UpdatedAt     time.Time `json:"updatedAt,omitempty"`
	DeletedAt     time.Time `json:"deletedAt,omitempty"`
}

// 新增参数非空检查
func (s *PublicAppVersionEntitys) AddCheck() error {
	return nil
}

// 修改参数非空检查
func (s *PublicAppVersionEntitys) UpdateCheck() error {
	return nil
}

// 查询参数必填检查
func (*PublicAppVersionQuery) QueryCheck() error {
	return nil
}

// 查询条件
type PublicAppVersionQuery struct {
	Page      uint64                  `json:"page,omitempty"`
	Limit     uint64                  `json:"limit,omitempty"`
	Sort      string                  `json:"sort,omitempty"`
	SortField string                  `json:"sortField,omitempty"`
	SearchKey string                  `json:"searchKey,omitempty"`
	Query     *PublicAppVersionFilter `json:"query,omitempty"`
}

// PublicAppVersionFilter，查询条件，字段请根据需要自行增减
type PublicAppVersionFilter struct {
	Id            int64     `json:"id,string,omitempty"`
	AppKey        string    `json:"appKey,omitempty"`
	ReleaseTime   time.Time `json:"releaseTime,omitempty"`
	ReleaseMarket int32     `json:"releaseMarket,omitempty"`
	Version       string    `json:"version,omitempty"`
	AppType       int32     `json:"appType,omitempty"`
	Status        int32     `json:"status,omitempty"`
	Remark        string    `json:"remark,omitempty"`
	CreatedBy     int64     `json:"createdBy,string,omitempty"`
	CreatedAt     time.Time `json:"createdAt,omitempty"`
	UpdatedBy     int64     `json:"updatedBy,string,omitempty"`
	UpdatedAt     time.Time `json:"updatedAt,omitempty"`
	DeletedAt     time.Time `json:"deletedAt,omitempty"`
}

// 实体转pb对象
func PublicAppVersion_e2pb(src *PublicAppVersionEntitys) *proto.PublicAppVersion {
	if src == nil {
		return nil
	}
	pbObj := proto.PublicAppVersion{
		Id:            src.Id,
		AppKey:        src.AppKey,
		ReleaseTime:   timestamppb.New(src.ReleaseTime),
		ReleaseMarket: src.ReleaseMarket,
		Version:       src.Version,
		AppType:       src.AppType,
		Status:        src.Status,
		Remark:        src.Remark,
		CreatedBy:     src.CreatedBy,
		CreatedAt:     timestamppb.New(src.CreatedAt),
		UpdatedBy:     src.UpdatedBy,
		UpdatedAt:     timestamppb.New(src.UpdatedAt),
	}
	return &pbObj
}

// pb对象转实体
func PublicAppVersion_pb2e(src *proto.PublicAppVersion) *PublicAppVersionEntitys {
	if src == nil {
		return nil
	}
	entitysObj := PublicAppVersionEntitys{
		Id:            src.Id,
		AppKey:        src.AppKey,
		ReleaseTime:   src.ReleaseTime.AsTime(),
		ReleaseMarket: src.ReleaseMarket,
		Version:       src.Version,
		AppType:       src.AppType,
		Status:        src.Status,
		Remark:        src.Remark,
		CreatedBy:     src.CreatedBy,
		CreatedAt:     src.CreatedAt.AsTime(),
		UpdatedBy:     src.UpdatedBy,
		UpdatedAt:     src.UpdatedAt.AsTime(),
	}
	return &entitysObj
}
