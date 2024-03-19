// Code generated by sgen.exe,2022-06-02 11:15:11. DO NOT EDIT.
// versions: v1.0.0
//  API结构体封装，请自行根据需要删减字段修改

package entitys

import (
	proto "cloud_platform/iot_proto/protos/protosService"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// 新增的版本检查的接口请求参数
type OemAppIntroduceVersionReq struct {
	AppId       string `json:"appId"`
	ContentType int32  `json:"contentType"` //1 用户协议,2隐私政策,3关于我们
	Version     string `json:"version"`
}

type OemAppIntroduceSaveReq struct {
	AppId       string `json:"appId"`
	Version     string `json:"version"`
	ContentType int32  `json:"contentType"` //1 用户协议,2隐私政策,3关于我们
	Lang        string `json:"lang"`
	Content     string `json:"content"`
	VioceCode   string `json:"vioceCode"`
	Abstract    string `json:"abstract"`
}

// 详情请求
type OemAppIntroduceDetailReq struct {
	AppId       string `json:"appId" form:"appId"`
	Version     string `json:"version" form:"version"`
	ContentType int32  `json:"contentType" form:"contentType"` //1 用户协议,2隐私政策,3关于我们
	Lang        string `json:"lang" form:"lang"`
	VioceCode   string `json:"vioceCode" form:"vioceCode"`
}

// 详情响应
type OemAppIntroduceDetailRes struct {
	AppId       string `json:"appId"`
	Version     string `json:"version"`
	ContentType int32  `json:"contentType"` //1 用户协议,2隐私政策,3关于我们
	Lang        string `json:"lang"`
	Content     string `json:"content"`
	VioceCode   string `json:"vioceCode"`
	Abstract    string `json:"abstract"`
}

type OemAppIntroduceStatusReq struct {
	AppId       string `json:"appId"`
	Version     string `json:"version"`
	ContentType int32  `json:"contentType"` //1 用户协议,2隐私政策,3关于我们
	//Status int32 `json:"status"`
}

type OemAppIntroduceListReq struct {
	AppId       string `json:"appId" form:"appId"`
	ContentType int32  `json:"contentType" form:"contentType"` //1 用户协议,2隐私政策,3关于我们,4语音文档
}

type OemAppIntroduceListRes struct {
	AppId       string `json:"appId"`
	Version     string `json:"version"`
	ContentType int32  `json:"contentType"` //1 用户协议,2隐私政策,3关于我们
	LangCount   int32  `json:"langCount"`
	Status      int32  `json:"status"`
	CreatedAt   int32  `json:"createdAt"`
	UpdatedAt   int32  `json:"updatedAt"`
}

// 协议模板链接响应参数
type OemAppIntroduceLinkRes struct {
	Url  string `json:"url"`
	Lang string `json:"lang"`
}

// 增、删、改及查询返回
type OemAppIntroduceEntitys struct {
	Id          int64     `json:"id,omitempty"`
	Content     string    `json:"content,omitempty"`
	ContentUrl  string    `json:"contentUrl,omitempty"`
	Lang        string    `json:"lang,omitempty"`
	Status      int32     `json:"status,omitempty"`
	ContentType int32     `json:"contentType,omitempty"`
	AppId       int64     `json:"appId,omitempty"`
	Version     string    `json:"version,omitempty"`
	CreatedBy   int64     `json:"createdBy,omitempty"`
	UpdatedBy   int64     `json:"updatedBy,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
	DeletedAt   time.Time `json:"deletedAt,omitempty"`
}

// 新增参数非空检查
func (s *OemAppIntroduceEntitys) AddCheck() error {
	return nil
}

// 修改参数非空检查
func (s *OemAppIntroduceEntitys) UpdateCheck() error {
	return nil
}

// 查询参数必填检查
func (*OemAppIntroduceQuery) QueryCheck() error {
	return nil
}

// 查询条件
type OemAppIntroduceQuery struct {
	Page      uint64                 `json:"page,omitempty"`
	Limit     uint64                 `json:"limit,omitempty"`
	Sort      string                 `json:"sort,omitempty"`
	SortField string                 `json:"sortField,omitempty"`
	SearchKey string                 `json:"searchKey,omitempty"`
	Query     *OemAppIntroduceFilter `json:"query,omitempty"`
}

// OemAppIntroduceFilter，查询条件，字段请根据需要自行增减
type OemAppIntroduceFilter struct {
	Id          int64     `json:"id,omitempty"`
	Content     string    `json:"content,omitempty"`
	ContentUrl  string    `json:"contentUrl,omitempty"`
	Lang        string    `json:"lang,omitempty"`
	Status      int32     `json:"status,omitempty"`
	ContentType int32     `json:"contentType,omitempty"`
	AppId       int64     `json:"appId,omitempty"`
	Version     string    `json:"version,omitempty"`
	CreatedBy   int64     `json:"createdBy,omitempty"`
	UpdatedBy   int64     `json:"updatedBy,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
	DeletedAt   time.Time `json:"deletedAt,omitempty"`
}

// 实体转pb对象
func OemAppIntroduce_e2pb(src *OemAppIntroduceEntitys) *proto.OemAppIntroduce {
	if src == nil {
		return nil
	}
	pbObj := proto.OemAppIntroduce{
		Id:          src.Id,
		Content:     src.Content,
		ContentUrl:  src.ContentUrl,
		Lang:        src.Lang,
		Status:      src.Status,
		ContentType: src.ContentType,
		AppId:       src.AppId,
		Version:     src.Version,
		CreatedBy:   src.CreatedBy,
		UpdatedBy:   src.UpdatedBy,
		CreatedAt:   timestamppb.New(src.CreatedAt),
		UpdatedAt:   timestamppb.New(src.UpdatedAt),
	}
	return &pbObj
}

// pb对象转实体
func OemAppIntroduce_pb2e(src *proto.OemAppIntroduce) *OemAppIntroduceEntitys {
	if src == nil {
		return nil
	}
	entitysObj := OemAppIntroduceEntitys{
		Id:          src.Id,
		Content:     src.Content,
		ContentUrl:  src.ContentUrl,
		Lang:        src.Lang,
		Status:      src.Status,
		ContentType: src.ContentType,
		AppId:       src.AppId,
		Version:     src.Version,
		CreatedBy:   src.CreatedBy,
		UpdatedBy:   src.UpdatedBy,
		CreatedAt:   src.CreatedAt.AsTime(),
		UpdatedAt:   src.UpdatedAt.AsTime(),
	}
	return &entitysObj
}
