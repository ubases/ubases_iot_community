// Code generated by sgen,2023-09-26 13:54:18. DO NOT EDIT.
// versions: v1.0.0
//  API结构体封装，请自行根据需要删减字段修改

package entitys

import (
	proto "cloud_platform/iot_proto/protos/protosService"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

//增、删、改及查询返回
type OpmPanelImageAssetEntitys struct {
	Id          int64     `json:"id,string,omitempty"`
	Resolution  string    `json:"resolution,omitempty"`
	AssetFormat string    `json:"assetFormat,omitempty"`
	Size        int64     `json:"size,string,omitempty"`
	AssetType   int32     `json:"assetType,omitempty"`
	Builtin     int32     `json:"builtin,omitempty"`
	IconType    int32     `json:"iconType,omitempty"`
	IconSubType int32     `json:"iconSubType,omitempty"`
	AssetName   string    `json:"assetName,omitempty"`
	Thumbnail   string    `json:"thumbnail,omitempty"`
	AssetUrl    string    `json:"assetUrl,omitempty"`
	TenantId    string    `json:"tenantId,omitempty"`
	CreatedBy   int64     `json:"createdBy,string,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	DeletedAt   time.Time `json:"deletedAt,omitempty"`
}

//新增参数非空检查
func (s *OpmPanelImageAssetEntitys) AddCheck() error {
	return nil
}

//修改参数非空检查
func (s *OpmPanelImageAssetEntitys) UpdateCheck() error {
	return nil
}

//查询参数必填检查
func (*OpmPanelImageAssetQuery) QueryCheck() error {
	return nil
}

//查询条件
type OpmPanelImageAssetQuery struct {
	Page      uint64                    `json:"page,omitempty"`
	Limit     uint64                    `json:"limit,omitempty"`
	Sort      string                    `json:"sort,omitempty"`
	SortField string                    `json:"sortField,omitempty"`
	SearchKey string                    `json:"searchKey,omitempty"`
	Query     *OpmPanelImageAssetFilter `json:"query,omitempty"`
}

//OpmPanelImageAssetFilter，查询条件，字段请根据需要自行增减
type OpmPanelImageAssetFilter struct {
	//Id int64  `json:"id,string,omitempty"`
	//Resolution string `json:"resolution,omitempty"`
	//AssetFormat string `json:"assetFormat,omitempty"`
	//Size int64  `json:"size,string,omitempty"`
	AssetType int32 `json:"assetType,omitempty"`
	//Builtin int32  `json:"builtin,omitempty"`
	IconType    int32 `json:"iconType,omitempty"`
	IconSubType int32 `json:"iconSubType,omitempty"`
	//AssetName string `json:"assetName,omitempty"`
	//Thumbnail string `json:"thumbnail,omitempty"`
	//AssetUrl string `json:"assetUrl,omitempty"`
	TenantId string `json:"tenantId,omitempty"`
	//CreatedBy int64  `json:"createdBy,string,omitempty"`
	//CreatedAt time.Time `json:"createdAt,omitempty"`
	//DeletedAt time.Time `json:"deletedAt,omitempty"`
}

//实体转pb对象
func OpmPanelImageAsset_e2pb(src *OpmPanelImageAssetEntitys) *proto.OpmPanelImageAsset {
	if src == nil {
		return nil
	}
	pbObj := proto.OpmPanelImageAsset{
		Id:          src.Id,
		Resolution:  src.Resolution,
		AssetFormat: src.AssetFormat,
		Size:        src.Size,
		AssetType:   src.AssetType,
		Builtin:     src.Builtin,
		IconType:    src.IconType,
		IconSubType: src.IconSubType,
		AssetName:   src.AssetName,
		Thumbnail:   src.Thumbnail,
		AssetUrl:    src.AssetUrl,
		TenantId:    src.TenantId,
		CreatedBy:   src.CreatedBy,
		CreatedAt:   timestamppb.New(src.CreatedAt),
	}
	return &pbObj
}

//pb对象转实体
func OpmPanelImageAsset_pb2e(src *proto.OpmPanelImageAsset) *OpmPanelImageAssetEntitys {
	if src == nil {
		return nil
	}
	entitysObj := OpmPanelImageAssetEntitys{
		Id:          src.Id,
		Resolution:  src.Resolution,
		AssetFormat: src.AssetFormat,
		Size:        src.Size,
		AssetType:   src.AssetType,
		Builtin:     src.Builtin,
		IconType:    src.IconType,
		IconSubType: src.IconSubType,
		AssetName:   src.AssetName,
		Thumbnail:   src.Thumbnail,
		AssetUrl:    src.AssetUrl,
		TenantId:    src.TenantId,
		CreatedBy:   src.CreatedBy,
		CreatedAt:   src.CreatedAt.AsTime(),
	}
	return &entitysObj
}
