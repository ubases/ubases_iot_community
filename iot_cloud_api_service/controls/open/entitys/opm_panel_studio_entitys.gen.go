// Code generated by sgen,2023-06-02 13:48:11. DO NOT EDIT.
// versions: v1.0.0
//  API结构体封装，请自行根据需要删减字段修改

package entitys

import (
	proto "cloud_platform/iot_proto/protos/protosService"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// 增、删、改及查询返回
type OpmPanelStudioEntitys struct {
	Id           int64     `json:"id,string,omitempty"`
	PanelId      int64     `json:"panelId,string,omitempty"`
	PageName     string    `json:"pageName"`
	PageIdentify string    `json:"pageIdentify"`
	JsonContent  string    `json:"jsonContent"`
	PopupContent string    `json:"popupContent"`
	VueContent   string    `json:"vueContent,omitempty"`
	StyleContent string    `json:"styleContent,omitempty"`
	IsHome       int32     `json:"isHome,omitempty"`
	Sort         int32     `json:"sort,omitempty"`
	CreatedAt    time.Time `json:"createdAt,omitempty"`
	UpdatedAt    time.Time `json:"updatedAt,omitempty"`
	UpdatedBy    int64     `json:"updatedBy,string,omitempty"`
	ProductId    int64     `json:"productId,omitempty"`
}

// 新增参数非空检查
func (s *OpmPanelStudioEntitys) AddCheck() error {
	return nil
}

// 修改参数非空检查
func (s *OpmPanelStudioEntitys) UpdateCheck() error {
	return nil
}

// 查询参数必填检查
func (*OpmPanelStudioQuery) QueryCheck() error {
	return nil
}

// 查询条件
type OpmPanelStudioQuery struct {
	Page      uint64                `json:"page,omitempty"`
	Limit     uint64                `json:"limit,omitempty"`
	Sort      string                `json:"sort,omitempty"`
	SortField string                `json:"sortField,omitempty"`
	SearchKey string                `json:"searchKey,omitempty"`
	Query     *OpmPanelStudioFilter `json:"query,omitempty"`
}

// OpmPanelStudioFilter，查询条件，字段请根据需要自行增减
type OpmPanelStudioFilter struct {
	Id           int64     `json:"id,string,omitempty"`
	PanelId      int64     `json:"panelId,string,omitempty"`
	PageName     string    `json:"pageName,omitempty"`
	PageIdentify string    `json:"pageIdentify,omitempty"`
	JsonContent  string    `json:"jsonContent,omitempty"`
	VueContent   string    `json:"vueContent,omitempty"`
	StyleContent string    `json:"styleContent,omitempty"`
	IsHome       int32     `json:"isHome,omitempty"`
	Sort         int32     `json:"sort,omitempty"`
	CreatedAt    time.Time `json:"createdAt,omitempty"`
	UpdatedAt    time.Time `json:"updatedAt,omitempty"`
	UpdatedBy    int64     `json:"updatedBy,string,omitempty"`
}

// 实体转pb对象
func OpmPanelStudio_e2pb(src *OpmPanelStudioEntitys) *proto.OpmPanelStudio {
	if src == nil {
		return nil
	}
	pbObj := proto.OpmPanelStudio{
		Id:           src.Id,
		PanelId:      src.PanelId,
		PageName:     src.PageName,
		PageIdentify: src.PageIdentify,
		JsonContent:  src.JsonContent,
		VueContent:   src.VueContent,
		StyleContent: src.StyleContent,
		IsHome:       src.IsHome,
		Sort:         src.Sort,
		CreatedAt:    timestamppb.New(src.CreatedAt),
		UpdatedAt:    timestamppb.New(src.UpdatedAt),
		UpdatedBy:    src.UpdatedBy,
		PopupContent: src.PopupContent,
	}
	return &pbObj
}

// pb对象转实体
func OpmPanelStudio_pb2e(src *proto.OpmPanelStudio) *OpmPanelStudioEntitys {
	if src == nil {
		return nil
	}
	entitysObj := OpmPanelStudioEntitys{
		Id:           src.Id,
		PanelId:      src.PanelId,
		PageName:     src.PageName,
		PageIdentify: src.PageIdentify,
		JsonContent:  src.JsonContent,
		VueContent:   src.VueContent,
		StyleContent: src.StyleContent,
		IsHome:       src.IsHome,
		Sort:         src.Sort,
		CreatedAt:    src.CreatedAt.AsTime(),
		UpdatedAt:    src.UpdatedAt.AsTime(),
		UpdatedBy:    src.UpdatedBy,
		PopupContent: src.PopupContent,
	}
	return &entitysObj
}
