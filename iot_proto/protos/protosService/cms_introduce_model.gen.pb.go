// Code generated by protoc,2022-05-11 22:57:20. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: cms_introduce_model.gen.proto

package protosService

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//来源于数据表
type CmsIntroduce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:;comment:;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"platformType" gorm:"default:平台类型 =1 云管平台 =2 开放平台 =3 APP;comment:平台类型 =1 云管平台 =2 开放平台 =3 APP;"
	PlatformType int32 `protobuf:"varint,101,opt,name=platformType,proto3" json:"platformType,omitempty"`
	// @inject_tag: json:"title" gorm:"default:标题;comment:标题;size:100;"
	Title string `protobuf:"bytes,102,opt,name=title,proto3" json:"title,omitempty"`
	// @inject_tag: json:"code" gorm:"default:编码;comment:编码;size:20;"
	Code string `protobuf:"bytes,103,opt,name=code,proto3" json:"code,omitempty"`
	// @inject_tag: json:"contentMode" gorm:"default:内容模式 =0 内容 =1 url;comment:内容模式 =0 内容 =1 url;"
	ContentMode int32 `protobuf:"varint,104,opt,name=contentMode,proto3" json:"contentMode,omitempty"`
	// @inject_tag: json:"content" gorm:"default:内容;comment:内容;"
	Content string `protobuf:"bytes,105,opt,name=content,proto3" json:"content,omitempty"`
	// @inject_tag: json:"contentUrl" gorm:"default:URL内容;comment:URL内容;size:255;"
	ContentUrl string `protobuf:"bytes,106,opt,name=contentUrl,proto3" json:"contentUrl,omitempty"`
	// @inject_tag: json:"lang" gorm:"default:语言;comment:语言;size:10;"
	Lang string `protobuf:"bytes,107,opt,name=lang,proto3" json:"lang,omitempty"`
	// @inject_tag: json:"status" gorm:"default:状态;comment:状态;"
	Status int32 `protobuf:"varint,108,opt,name=status,proto3" json:"status,omitempty"`
	// @inject_tag: json:"createdBy" gorm:"default:创建人;comment:创建人;size:20;"
	CreatedBy int64 `protobuf:"varint,109,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	// @inject_tag: json:"updatedBy" gorm:"default:修改人;comment:修改人;size:20;"
	UpdatedBy int64 `protobuf:"varint,110,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	// @inject_tag: json:"createdAt" gorm:"default:创建时间;comment:创建时间;"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,111,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// @inject_tag: json:"updatedAt" gorm:"default:修改时间;comment:修改时间;"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,112,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// @inject_tag: json:"deletedAt" gorm:"default:删除时间;comment:删除时间;"
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,113,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *CmsIntroduce) Reset() {
	*x = CmsIntroduce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cms_introduce_model_gen_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmsIntroduce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmsIntroduce) ProtoMessage() {}

func (x *CmsIntroduce) ProtoReflect() protoreflect.Message {
	mi := &file_cms_introduce_model_gen_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmsIntroduce.ProtoReflect.Descriptor instead.
func (*CmsIntroduce) Descriptor() ([]byte, []int) {
	return file_cms_introduce_model_gen_proto_rawDescGZIP(), []int{0}
}

func (x *CmsIntroduce) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CmsIntroduce) GetPlatformType() int32 {
	if x != nil {
		return x.PlatformType
	}
	return 0
}

func (x *CmsIntroduce) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CmsIntroduce) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CmsIntroduce) GetContentMode() int32 {
	if x != nil {
		return x.ContentMode
	}
	return 0
}

func (x *CmsIntroduce) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CmsIntroduce) GetContentUrl() string {
	if x != nil {
		return x.ContentUrl
	}
	return ""
}

func (x *CmsIntroduce) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *CmsIntroduce) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CmsIntroduce) GetCreatedBy() int64 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *CmsIntroduce) GetUpdatedBy() int64 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

func (x *CmsIntroduce) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *CmsIntroduce) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *CmsIntroduce) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type CmsIntroduceFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:;comment:;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"platformType" gorm:"default:平台类型 =1 云管平台 =2 开放平台 =3 APP;comment:平台类型 =1 云管平台 =2 开放平台 =3 APP;"
	PlatformType int32 `protobuf:"varint,101,opt,name=platformType,proto3" json:"platformType,omitempty"`
	// @inject_tag: json:"title" gorm:"default:标题;comment:标题;size:100;"
	Title string `protobuf:"bytes,102,opt,name=title,proto3" json:"title,omitempty"`
	// @inject_tag: json:"code" gorm:"default:编码;comment:编码;size:20;"
	Code string `protobuf:"bytes,103,opt,name=code,proto3" json:"code,omitempty"`
	// @inject_tag: json:"contentMode" gorm:"default:内容模式 =0 内容 =1 url;comment:内容模式 =0 内容 =1 url;"
	ContentMode int32 `protobuf:"varint,104,opt,name=contentMode,proto3" json:"contentMode,omitempty"`
	// @inject_tag: json:"content" gorm:"default:内容;comment:内容;"
	Content string `protobuf:"bytes,105,opt,name=content,proto3" json:"content,omitempty"`
	// @inject_tag: json:"contentUrl" gorm:"default:URL内容;comment:URL内容;size:255;"
	ContentUrl string `protobuf:"bytes,106,opt,name=contentUrl,proto3" json:"contentUrl,omitempty"`
	// @inject_tag: json:"lang" gorm:"default:语言;comment:语言;size:10;"
	Lang string `protobuf:"bytes,107,opt,name=lang,proto3" json:"lang,omitempty"`
	// @inject_tag: json:"status" gorm:"default:状态;comment:状态;"
	Status int32 `protobuf:"varint,108,opt,name=status,proto3" json:"status,omitempty"`
	// @inject_tag: json:"createdBy" gorm:"default:创建人;comment:创建人;size:20;"
	CreatedBy int64 `protobuf:"varint,109,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	// @inject_tag: json:"updatedBy" gorm:"default:修改人;comment:修改人;size:20;"
	UpdatedBy int64 `protobuf:"varint,110,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	// @inject_tag: json:"createdAt" gorm:"default:创建时间;comment:创建时间;"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,111,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// @inject_tag: json:"updatedAt" gorm:"default:修改时间;comment:修改时间;"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,112,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// @inject_tag: json:"deletedAt" gorm:"default:删除时间;comment:删除时间;"
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,113,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *CmsIntroduceFilter) Reset() {
	*x = CmsIntroduceFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cms_introduce_model_gen_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmsIntroduceFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmsIntroduceFilter) ProtoMessage() {}

func (x *CmsIntroduceFilter) ProtoReflect() protoreflect.Message {
	mi := &file_cms_introduce_model_gen_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmsIntroduceFilter.ProtoReflect.Descriptor instead.
func (*CmsIntroduceFilter) Descriptor() ([]byte, []int) {
	return file_cms_introduce_model_gen_proto_rawDescGZIP(), []int{1}
}

func (x *CmsIntroduceFilter) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CmsIntroduceFilter) GetPlatformType() int32 {
	if x != nil {
		return x.PlatformType
	}
	return 0
}

func (x *CmsIntroduceFilter) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CmsIntroduceFilter) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CmsIntroduceFilter) GetContentMode() int32 {
	if x != nil {
		return x.ContentMode
	}
	return 0
}

func (x *CmsIntroduceFilter) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CmsIntroduceFilter) GetContentUrl() string {
	if x != nil {
		return x.ContentUrl
	}
	return ""
}

func (x *CmsIntroduceFilter) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *CmsIntroduceFilter) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CmsIntroduceFilter) GetCreatedBy() int64 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *CmsIntroduceFilter) GetUpdatedBy() int64 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

func (x *CmsIntroduceFilter) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *CmsIntroduceFilter) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *CmsIntroduceFilter) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type CmsIntroduceListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query     *CmsIntroduce `protobuf:"bytes,101,opt,name=query,proto3" json:"query,omitempty"`
	Page      int64         `protobuf:"varint,102,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int64         `protobuf:"varint,103,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	OrderKey  string        `protobuf:"bytes,104,opt,name=orderKey,proto3" json:"orderKey,omitempty"`
	OrderDesc string        `protobuf:"bytes,105,opt,name=orderDesc,proto3" json:"orderDesc,omitempty"`
	SearchKey string        `protobuf:"bytes,106,opt,name=searchKey,proto3" json:"searchKey,omitempty"`
}

func (x *CmsIntroduceListRequest) Reset() {
	*x = CmsIntroduceListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cms_introduce_model_gen_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmsIntroduceListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmsIntroduceListRequest) ProtoMessage() {}

func (x *CmsIntroduceListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cms_introduce_model_gen_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmsIntroduceListRequest.ProtoReflect.Descriptor instead.
func (*CmsIntroduceListRequest) Descriptor() ([]byte, []int) {
	return file_cms_introduce_model_gen_proto_rawDescGZIP(), []int{2}
}

func (x *CmsIntroduceListRequest) GetQuery() *CmsIntroduce {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *CmsIntroduceListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *CmsIntroduceListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *CmsIntroduceListRequest) GetOrderKey() string {
	if x != nil {
		return x.OrderKey
	}
	return ""
}

func (x *CmsIntroduceListRequest) GetOrderDesc() string {
	if x != nil {
		return x.OrderDesc
	}
	return ""
}

func (x *CmsIntroduceListRequest) GetSearchKey() string {
	if x != nil {
		return x.SearchKey
	}
	return ""
}

type CmsIntroduceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32           `protobuf:"varint,101,opt,name=code,proto3" json:"code,omitempty"`
	Message string          `protobuf:"bytes,102,opt,name=message,proto3" json:"message,omitempty"`
	Total   int64           `protobuf:"varint,103,opt,name=total,proto3" json:"total,omitempty"`
	Data    []*CmsIntroduce `protobuf:"bytes,104,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *CmsIntroduceResponse) Reset() {
	*x = CmsIntroduceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cms_introduce_model_gen_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmsIntroduceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmsIntroduceResponse) ProtoMessage() {}

func (x *CmsIntroduceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cms_introduce_model_gen_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmsIntroduceResponse.ProtoReflect.Descriptor instead.
func (*CmsIntroduceResponse) Descriptor() ([]byte, []int) {
	return file_cms_introduce_model_gen_proto_rawDescGZIP(), []int{3}
}

func (x *CmsIntroduceResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CmsIntroduceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CmsIntroduceResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *CmsIntroduceResponse) GetData() []*CmsIntroduce {
	if x != nil {
		return x.Data
	}
	return nil
}

type CmsIntroduceUpdateFieldsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []string      `protobuf:"bytes,101,rep,name=fields,proto3" json:"fields,omitempty"`
	Data   *CmsIntroduce `protobuf:"bytes,102,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CmsIntroduceUpdateFieldsRequest) Reset() {
	*x = CmsIntroduceUpdateFieldsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cms_introduce_model_gen_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmsIntroduceUpdateFieldsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmsIntroduceUpdateFieldsRequest) ProtoMessage() {}

func (x *CmsIntroduceUpdateFieldsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cms_introduce_model_gen_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmsIntroduceUpdateFieldsRequest.ProtoReflect.Descriptor instead.
func (*CmsIntroduceUpdateFieldsRequest) Descriptor() ([]byte, []int) {
	return file_cms_introduce_model_gen_proto_rawDescGZIP(), []int{4}
}

func (x *CmsIntroduceUpdateFieldsRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *CmsIntroduceUpdateFieldsRequest) GetData() *CmsIntroduce {
	if x != nil {
		return x.Data
	}
	return nil
}

// 表模型主键，只能是整数或字符串，若生成错误，请检查数据库表主键是否正确。支持组合主键
type CmsIntroducePrimarykey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CmsIntroducePrimarykey) Reset() {
	*x = CmsIntroducePrimarykey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cms_introduce_model_gen_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmsIntroducePrimarykey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmsIntroducePrimarykey) ProtoMessage() {}

func (x *CmsIntroducePrimarykey) ProtoReflect() protoreflect.Message {
	mi := &file_cms_introduce_model_gen_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmsIntroducePrimarykey.ProtoReflect.Descriptor instead.
func (*CmsIntroducePrimarykey) Descriptor() ([]byte, []int) {
	return file_cms_introduce_model_gen_proto_rawDescGZIP(), []int{5}
}

func (x *CmsIntroducePrimarykey) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CmsIntroduceBatchDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []*CmsIntroducePrimarykey `protobuf:"bytes,101,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *CmsIntroduceBatchDeleteRequest) Reset() {
	*x = CmsIntroduceBatchDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cms_introduce_model_gen_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmsIntroduceBatchDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmsIntroduceBatchDeleteRequest) ProtoMessage() {}

func (x *CmsIntroduceBatchDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cms_introduce_model_gen_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmsIntroduceBatchDeleteRequest.ProtoReflect.Descriptor instead.
func (*CmsIntroduceBatchDeleteRequest) Descriptor() ([]byte, []int) {
	return file_cms_introduce_model_gen_proto_rawDescGZIP(), []int{6}
}

func (x *CmsIntroduceBatchDeleteRequest) GetKeys() []*CmsIntroducePrimarykey {
	if x != nil {
		return x.Keys
	}
	return nil
}

var File_cms_introduce_model_gen_proto protoreflect.FileDescriptor

var file_cms_introduce_model_gen_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x6d, 0x73, 0x5f, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xde, 0x03, 0x0a, 0x0c, 0x43, 0x6d,
	0x73, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x67, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x68, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x69, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x55,
	0x72, 0x6c, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x6b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x6c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x6d, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1c,
	0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x6e, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x38, 0x0a, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x6f, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x70, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x38, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x71, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xe4, 0x03, 0x0a, 0x12, 0x43,
	0x6d, 0x73, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x66,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x68,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x69, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x55, 0x72, 0x6c, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6c,
	0x61, 0x6e, 0x67, 0x18, 0x6b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x6c, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x42, 0x79, 0x18, 0x6d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x6f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x70, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x71, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0xce, 0x01, 0x0a, 0x17, 0x43, 0x6d, 0x73, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6d, 0x73, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x65, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44,
	0x65, 0x73, 0x63, 0x18, 0x69, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x44, 0x65, 0x73, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65,
	0x79, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b,
	0x65, 0x79, 0x22, 0x85, 0x01, 0x0a, 0x14, 0x43, 0x6d, 0x73, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x68, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6d, 0x73, 0x49, 0x6e, 0x74, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x64, 0x0a, 0x1f, 0x43, 0x6d,
	0x73, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x66, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6d,
	0x73, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x28, 0x0a, 0x16, 0x43, 0x6d, 0x73, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65,
	0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x55, 0x0a, 0x1e, 0x43, 0x6d,
	0x73, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x04,
	0x6b, 0x65, 0x79, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6d, 0x73, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x65, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x52, 0x04, 0x6b, 0x65, 0x79,
	0x73, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cms_introduce_model_gen_proto_rawDescOnce sync.Once
	file_cms_introduce_model_gen_proto_rawDescData = file_cms_introduce_model_gen_proto_rawDesc
)

func file_cms_introduce_model_gen_proto_rawDescGZIP() []byte {
	file_cms_introduce_model_gen_proto_rawDescOnce.Do(func() {
		file_cms_introduce_model_gen_proto_rawDescData = protoimpl.X.CompressGZIP(file_cms_introduce_model_gen_proto_rawDescData)
	})
	return file_cms_introduce_model_gen_proto_rawDescData
}

var file_cms_introduce_model_gen_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cms_introduce_model_gen_proto_goTypes = []interface{}{
	(*CmsIntroduce)(nil),                    // 0: service.CmsIntroduce
	(*CmsIntroduceFilter)(nil),              // 1: service.CmsIntroduceFilter
	(*CmsIntroduceListRequest)(nil),         // 2: service.CmsIntroduceListRequest
	(*CmsIntroduceResponse)(nil),            // 3: service.CmsIntroduceResponse
	(*CmsIntroduceUpdateFieldsRequest)(nil), // 4: service.CmsIntroduceUpdateFieldsRequest
	(*CmsIntroducePrimarykey)(nil),          // 5: service.CmsIntroducePrimarykey
	(*CmsIntroduceBatchDeleteRequest)(nil),  // 6: service.CmsIntroduceBatchDeleteRequest
	(*timestamppb.Timestamp)(nil),           // 7: google.protobuf.Timestamp
}
var file_cms_introduce_model_gen_proto_depIdxs = []int32{
	7,  // 0: service.CmsIntroduce.createdAt:type_name -> google.protobuf.Timestamp
	7,  // 1: service.CmsIntroduce.updatedAt:type_name -> google.protobuf.Timestamp
	7,  // 2: service.CmsIntroduce.deletedAt:type_name -> google.protobuf.Timestamp
	7,  // 3: service.CmsIntroduceFilter.createdAt:type_name -> google.protobuf.Timestamp
	7,  // 4: service.CmsIntroduceFilter.updatedAt:type_name -> google.protobuf.Timestamp
	7,  // 5: service.CmsIntroduceFilter.deletedAt:type_name -> google.protobuf.Timestamp
	0,  // 6: service.CmsIntroduceListRequest.query:type_name -> service.CmsIntroduce
	0,  // 7: service.CmsIntroduceResponse.data:type_name -> service.CmsIntroduce
	0,  // 8: service.CmsIntroduceUpdateFieldsRequest.data:type_name -> service.CmsIntroduce
	5,  // 9: service.CmsIntroduceBatchDeleteRequest.keys:type_name -> service.CmsIntroducePrimarykey
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_cms_introduce_model_gen_proto_init() }
func file_cms_introduce_model_gen_proto_init() {
	if File_cms_introduce_model_gen_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cms_introduce_model_gen_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmsIntroduce); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cms_introduce_model_gen_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmsIntroduceFilter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cms_introduce_model_gen_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmsIntroduceListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cms_introduce_model_gen_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmsIntroduceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cms_introduce_model_gen_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmsIntroduceUpdateFieldsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cms_introduce_model_gen_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmsIntroducePrimarykey); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cms_introduce_model_gen_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmsIntroduceBatchDeleteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cms_introduce_model_gen_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cms_introduce_model_gen_proto_goTypes,
		DependencyIndexes: file_cms_introduce_model_gen_proto_depIdxs,
		MessageInfos:      file_cms_introduce_model_gen_proto_msgTypes,
	}.Build()
	File_cms_introduce_model_gen_proto = out.File
	file_cms_introduce_model_gen_proto_rawDesc = nil
	file_cms_introduce_model_gen_proto_goTypes = nil
	file_cms_introduce_model_gen_proto_depIdxs = nil
}
