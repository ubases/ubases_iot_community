// Code generated by sgen.exe,2022-07-25 09:29:22. DO NOT EDIT.
// versions: v1.0.0
//  API结构体封装，请自行根据需要删减字段修改

package entitys

import (
	proto "cloud_platform/iot_proto/protos/protosService"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

// 词条列表查询条件
type SysAppEntryListReq struct {
	DirId    string `json:"dirId"`
	Lang     string `json:"lang"`
	Title    string `json:"title"`
	IsNormal int32  `json:"isNormal"`
	IsEnable int32  `json:"isEnable"`
	Page     int64  `json:"pageNum"`
	PageSize int64  `json:"pageSize"`
}

// 词条列表查询条件
type SysAppEntryListRes struct {
	SetingId  string `json:"setingId"`  // 设置id
	DirId     string `json:"dirId"`     // 目录id
	Lang      string `json:"lang"`      // 语种编码
	Title     string `json:"title"`     // 标题
	IsNormal  int32  `json:"isNormal"`  // 是否常见
	IsEnable  int32  `json:"isEnable"`  // 是否启用
	DirName   string `json:"dirName"`   // 目录名称
	Sort      int64  `json:"sort"`      //排序
	UpdatedAt int64  `json:"updatedAt"` // 更新时间
}

type SysAppEntrySetingSaveReq struct {
	IsEnable int64  `json:"isEnable"` // 是否启用
	IsNormal int64  `json:"isNormal"` // 是否设为常见问题
	DirId    string `json:"dirId"`    // 目录id
	SetingId string `json:"setingId"` //设置id
	Sort     int64  `json:"sort"`     // 排序
}

// 词条新增参数
type SysAppEntrySaveReq struct {
	Content  string `json:"content"`
	Lang     string `json:"lang"`
	DirId    string `json:"dirId"` //目录id
	Title    string `json:"title"`
	SetingId string `json:"setingId"` //设置id. 默认值零
}

// 增、删、改及查询返回
type SysAppEntryEntitys struct {
	Id        int64     `json:"id,omitempty"`
	Lang      string    `json:"lang,omitempty"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	SetingId  int64     `json:"setingId,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

// 新增参数非空检查
func (s *SysAppEntryEntitys) AddCheck() error {
	return nil
}

// 修改参数非空检查
func (s *SysAppEntryEntitys) UpdateCheck() error {
	return nil
}

// 查询参数必填检查
func (*SysAppEntryQuery) QueryCheck() error {
	return nil
}

// 查询条件
type SysAppEntryQuery struct {
	Page      uint64             `json:"page,omitempty"`
	Limit     uint64             `json:"limit,omitempty"`
	Sort      string             `json:"sort,omitempty"`
	SortField string             `json:"sortField,omitempty"`
	SearchKey string             `json:"searchKey,omitempty"`
	Query     *SysAppEntryFilter `json:"query,omitempty"`
}

// SysAppEntryFilter，查询条件，字段请根据需要自行增减
type SysAppEntryFilter struct {
	Id        int64     `json:"id,omitempty"`
	Lang      string    `json:"lang,omitempty"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	SetingId  int64     `json:"setingId,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

// 实体转pb对象
func SysAppEntry_e2pb(src *SysAppEntryEntitys) *proto.SysAppEntry {
	if src == nil {
		return nil
	}
	pbObj := proto.SysAppEntry{
		Id:        src.Id,
		Lang:      src.Lang,
		Title:     src.Title,
		Content:   src.Content,
		SetingId:  src.SetingId,
		UpdatedAt: timestamppb.New(src.UpdatedAt),
	}
	return &pbObj
}

// pb对象转实体
func SysAppEntry_pb2e(src *proto.SysAppEntry) *SysAppEntryEntitys {
	if src == nil {
		return nil
	}
	entitysObj := SysAppEntryEntitys{
		Id:        src.Id,
		Lang:      src.Lang,
		Title:     src.Title,
		Content:   src.Content,
		SetingId:  src.SetingId,
		UpdatedAt: src.UpdatedAt.AsTime(),
	}
	return &entitysObj
}
