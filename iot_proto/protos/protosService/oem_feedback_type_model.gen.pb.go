// Code generated by protoc,2022-10-24 08:40:57. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: oem_feedback_type_model.gen.proto

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
type OemFeedbackType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:id;comment:id;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"name" gorm:"default:名称;comment:名称;size:255;"
	Name string `protobuf:"bytes,101,opt,name=name,proto3" json:"name,omitempty"`
	// @inject_tag: json:"appId" gorm:"default:;comment:;size:19;"
	AppId int64 `protobuf:"varint,102,opt,name=appId,proto3" json:"appId,omitempty"`
	// @inject_tag: json:"appKey" gorm:"default:;comment:;size:50;"
	AppKey string `protobuf:"bytes,103,opt,name=appKey,proto3" json:"appKey,omitempty"`
	// @inject_tag: json:"sort" gorm:"default:排序;comment:排序;size:10;"
	Sort int32 `protobuf:"varint,104,opt,name=sort,proto3" json:"sort,omitempty"`
	// @inject_tag: json:"createdAt" gorm:"default:创建时间;comment:创建时间;"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,105,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// @inject_tag: json:"updatedAt" gorm:"default:修改时间;comment:修改时间;"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,106,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// @inject_tag: json:"deletedAt" gorm:"default:删除时间;comment:删除时间;"
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,107,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
	TenantId  string                 `protobuf:"bytes,108,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	AppName   string                 `protobuf:"bytes,109,opt,name=appName,proto3" json:"appName,omitempty"`
}

func (x *OemFeedbackType) Reset() {
	*x = OemFeedbackType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oem_feedback_type_model_gen_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OemFeedbackType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OemFeedbackType) ProtoMessage() {}

func (x *OemFeedbackType) ProtoReflect() protoreflect.Message {
	mi := &file_oem_feedback_type_model_gen_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OemFeedbackType.ProtoReflect.Descriptor instead.
func (*OemFeedbackType) Descriptor() ([]byte, []int) {
	return file_oem_feedback_type_model_gen_proto_rawDescGZIP(), []int{0}
}

func (x *OemFeedbackType) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OemFeedbackType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OemFeedbackType) GetAppId() int64 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *OemFeedbackType) GetAppKey() string {
	if x != nil {
		return x.AppKey
	}
	return ""
}

func (x *OemFeedbackType) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *OemFeedbackType) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *OemFeedbackType) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *OemFeedbackType) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *OemFeedbackType) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *OemFeedbackType) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

type OemFeedbackTypeFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:id;comment:id;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"name" gorm:"default:名称;comment:名称;size:255;"
	Name string `protobuf:"bytes,101,opt,name=name,proto3" json:"name,omitempty"`
	// @inject_tag: json:"appId" gorm:"default:;comment:;size:19;"
	AppId int64 `protobuf:"varint,102,opt,name=appId,proto3" json:"appId,omitempty"`
	// @inject_tag: json:"appKey" gorm:"default:;comment:;size:50;"
	AppKey string `protobuf:"bytes,103,opt,name=appKey,proto3" json:"appKey,omitempty"`
	// @inject_tag: json:"sort" gorm:"default:排序;comment:排序;size:10;"
	Sort int32 `protobuf:"varint,104,opt,name=sort,proto3" json:"sort,omitempty"`
	// @inject_tag: json:"createdAt" gorm:"default:创建时间;comment:创建时间;"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,105,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// @inject_tag: json:"updatedAt" gorm:"default:修改时间;comment:修改时间;"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,106,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// @inject_tag: json:"deletedAt" gorm:"default:删除时间;comment:删除时间;"
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,107,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *OemFeedbackTypeFilter) Reset() {
	*x = OemFeedbackTypeFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oem_feedback_type_model_gen_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OemFeedbackTypeFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OemFeedbackTypeFilter) ProtoMessage() {}

func (x *OemFeedbackTypeFilter) ProtoReflect() protoreflect.Message {
	mi := &file_oem_feedback_type_model_gen_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OemFeedbackTypeFilter.ProtoReflect.Descriptor instead.
func (*OemFeedbackTypeFilter) Descriptor() ([]byte, []int) {
	return file_oem_feedback_type_model_gen_proto_rawDescGZIP(), []int{1}
}

func (x *OemFeedbackTypeFilter) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OemFeedbackTypeFilter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OemFeedbackTypeFilter) GetAppId() int64 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *OemFeedbackTypeFilter) GetAppKey() string {
	if x != nil {
		return x.AppKey
	}
	return ""
}

func (x *OemFeedbackTypeFilter) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *OemFeedbackTypeFilter) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *OemFeedbackTypeFilter) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *OemFeedbackTypeFilter) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type OemFeedbackTypeListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query     *OemFeedbackType `protobuf:"bytes,101,opt,name=query,proto3" json:"query,omitempty"`
	Page      int64            `protobuf:"varint,102,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int64            `protobuf:"varint,103,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	OrderKey  string           `protobuf:"bytes,104,opt,name=orderKey,proto3" json:"orderKey,omitempty"`
	OrderDesc string           `protobuf:"bytes,105,opt,name=orderDesc,proto3" json:"orderDesc,omitempty"`
	SearchKey string           `protobuf:"bytes,106,opt,name=searchKey,proto3" json:"searchKey,omitempty"`
}

func (x *OemFeedbackTypeListRequest) Reset() {
	*x = OemFeedbackTypeListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oem_feedback_type_model_gen_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OemFeedbackTypeListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OemFeedbackTypeListRequest) ProtoMessage() {}

func (x *OemFeedbackTypeListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oem_feedback_type_model_gen_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OemFeedbackTypeListRequest.ProtoReflect.Descriptor instead.
func (*OemFeedbackTypeListRequest) Descriptor() ([]byte, []int) {
	return file_oem_feedback_type_model_gen_proto_rawDescGZIP(), []int{2}
}

func (x *OemFeedbackTypeListRequest) GetQuery() *OemFeedbackType {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *OemFeedbackTypeListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *OemFeedbackTypeListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *OemFeedbackTypeListRequest) GetOrderKey() string {
	if x != nil {
		return x.OrderKey
	}
	return ""
}

func (x *OemFeedbackTypeListRequest) GetOrderDesc() string {
	if x != nil {
		return x.OrderDesc
	}
	return ""
}

func (x *OemFeedbackTypeListRequest) GetSearchKey() string {
	if x != nil {
		return x.SearchKey
	}
	return ""
}

type OemFeedbackTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32              `protobuf:"varint,101,opt,name=code,proto3" json:"code,omitempty"`
	Message string             `protobuf:"bytes,102,opt,name=message,proto3" json:"message,omitempty"`
	Total   int64              `protobuf:"varint,103,opt,name=total,proto3" json:"total,omitempty"`
	Data    []*OemFeedbackType `protobuf:"bytes,104,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *OemFeedbackTypeResponse) Reset() {
	*x = OemFeedbackTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oem_feedback_type_model_gen_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OemFeedbackTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OemFeedbackTypeResponse) ProtoMessage() {}

func (x *OemFeedbackTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_oem_feedback_type_model_gen_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OemFeedbackTypeResponse.ProtoReflect.Descriptor instead.
func (*OemFeedbackTypeResponse) Descriptor() ([]byte, []int) {
	return file_oem_feedback_type_model_gen_proto_rawDescGZIP(), []int{3}
}

func (x *OemFeedbackTypeResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *OemFeedbackTypeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *OemFeedbackTypeResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *OemFeedbackTypeResponse) GetData() []*OemFeedbackType {
	if x != nil {
		return x.Data
	}
	return nil
}

type OemFeedbackTypeUpdateFieldsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []string         `protobuf:"bytes,101,rep,name=fields,proto3" json:"fields,omitempty"`
	Data   *OemFeedbackType `protobuf:"bytes,102,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OemFeedbackTypeUpdateFieldsRequest) Reset() {
	*x = OemFeedbackTypeUpdateFieldsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oem_feedback_type_model_gen_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OemFeedbackTypeUpdateFieldsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OemFeedbackTypeUpdateFieldsRequest) ProtoMessage() {}

func (x *OemFeedbackTypeUpdateFieldsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oem_feedback_type_model_gen_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OemFeedbackTypeUpdateFieldsRequest.ProtoReflect.Descriptor instead.
func (*OemFeedbackTypeUpdateFieldsRequest) Descriptor() ([]byte, []int) {
	return file_oem_feedback_type_model_gen_proto_rawDescGZIP(), []int{4}
}

func (x *OemFeedbackTypeUpdateFieldsRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *OemFeedbackTypeUpdateFieldsRequest) GetData() *OemFeedbackType {
	if x != nil {
		return x.Data
	}
	return nil
}

// 表模型主键，只能是整数或字符串，若生成错误，请检查数据库表主键是否正确。支持组合主键
type OemFeedbackTypePrimarykey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OemFeedbackTypePrimarykey) Reset() {
	*x = OemFeedbackTypePrimarykey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oem_feedback_type_model_gen_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OemFeedbackTypePrimarykey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OemFeedbackTypePrimarykey) ProtoMessage() {}

func (x *OemFeedbackTypePrimarykey) ProtoReflect() protoreflect.Message {
	mi := &file_oem_feedback_type_model_gen_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OemFeedbackTypePrimarykey.ProtoReflect.Descriptor instead.
func (*OemFeedbackTypePrimarykey) Descriptor() ([]byte, []int) {
	return file_oem_feedback_type_model_gen_proto_rawDescGZIP(), []int{5}
}

func (x *OemFeedbackTypePrimarykey) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type OemFeedbackTypeBatchDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []*OemFeedbackTypePrimarykey `protobuf:"bytes,101,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *OemFeedbackTypeBatchDeleteRequest) Reset() {
	*x = OemFeedbackTypeBatchDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oem_feedback_type_model_gen_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OemFeedbackTypeBatchDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OemFeedbackTypeBatchDeleteRequest) ProtoMessage() {}

func (x *OemFeedbackTypeBatchDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oem_feedback_type_model_gen_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OemFeedbackTypeBatchDeleteRequest.ProtoReflect.Descriptor instead.
func (*OemFeedbackTypeBatchDeleteRequest) Descriptor() ([]byte, []int) {
	return file_oem_feedback_type_model_gen_proto_rawDescGZIP(), []int{6}
}

func (x *OemFeedbackTypeBatchDeleteRequest) GetKeys() []*OemFeedbackTypePrimarykey {
	if x != nil {
		return x.Keys
	}
	return nil
}

var File_oem_feedback_type_model_gen_proto protoreflect.FileDescriptor

var file_oem_feedback_type_model_gen_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6f, 0x65, 0x6d, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x02,
	0x0a, 0x0f, 0x4f, 0x65, 0x6d, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x18, 0x66,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x70, 0x70, 0x4b, 0x65, 0x79, 0x18, 0x67, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x70,
	0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x68, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x69, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x6a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x6b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49,
	0x64, 0x18, 0x6c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x6d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xab, 0x02, 0x0a, 0x15,
	0x4f, 0x65, 0x6d, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x65, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70,
	0x49, 0x64, 0x18, 0x66, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x4b, 0x65, 0x79, 0x18, 0x67, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x70, 0x70, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18,
	0x68, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x69, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x38, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x6b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xd4, 0x01, 0x0a, 0x1a, 0x4f, 0x65,
	0x6d, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x4f, 0x65, 0x6d, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x66, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x4b, 0x65, 0x79, 0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x73,
	0x63, 0x18, 0x69, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65,
	0x73, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x18,
	0x6a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79,
	0x22, 0x8b, 0x01, 0x0a, 0x17, 0x4f, 0x65, 0x6d, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x68, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x65, 0x6d, 0x46, 0x65, 0x65, 0x64,
	0x62, 0x61, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x6a,
	0x0a, 0x22, 0x4f, 0x65, 0x6d, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x54, 0x79, 0x70,
	0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x65,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x2c, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x65, 0x6d, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2b, 0x0a, 0x19, 0x4f, 0x65,
	0x6d, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5b, 0x0a, 0x21, 0x4f, 0x65, 0x6d, 0x46, 0x65,
	0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04,
	0x6b, 0x65, 0x79, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x65, 0x6d, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b,
	0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x52, 0x04,
	0x6b, 0x65, 0x79, 0x73, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_oem_feedback_type_model_gen_proto_rawDescOnce sync.Once
	file_oem_feedback_type_model_gen_proto_rawDescData = file_oem_feedback_type_model_gen_proto_rawDesc
)

func file_oem_feedback_type_model_gen_proto_rawDescGZIP() []byte {
	file_oem_feedback_type_model_gen_proto_rawDescOnce.Do(func() {
		file_oem_feedback_type_model_gen_proto_rawDescData = protoimpl.X.CompressGZIP(file_oem_feedback_type_model_gen_proto_rawDescData)
	})
	return file_oem_feedback_type_model_gen_proto_rawDescData
}

var file_oem_feedback_type_model_gen_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_oem_feedback_type_model_gen_proto_goTypes = []interface{}{
	(*OemFeedbackType)(nil),                    // 0: service.OemFeedbackType
	(*OemFeedbackTypeFilter)(nil),              // 1: service.OemFeedbackTypeFilter
	(*OemFeedbackTypeListRequest)(nil),         // 2: service.OemFeedbackTypeListRequest
	(*OemFeedbackTypeResponse)(nil),            // 3: service.OemFeedbackTypeResponse
	(*OemFeedbackTypeUpdateFieldsRequest)(nil), // 4: service.OemFeedbackTypeUpdateFieldsRequest
	(*OemFeedbackTypePrimarykey)(nil),          // 5: service.OemFeedbackTypePrimarykey
	(*OemFeedbackTypeBatchDeleteRequest)(nil),  // 6: service.OemFeedbackTypeBatchDeleteRequest
	(*timestamppb.Timestamp)(nil),              // 7: google.protobuf.Timestamp
}
var file_oem_feedback_type_model_gen_proto_depIdxs = []int32{
	7,  // 0: service.OemFeedbackType.createdAt:type_name -> google.protobuf.Timestamp
	7,  // 1: service.OemFeedbackType.updatedAt:type_name -> google.protobuf.Timestamp
	7,  // 2: service.OemFeedbackType.deletedAt:type_name -> google.protobuf.Timestamp
	7,  // 3: service.OemFeedbackTypeFilter.createdAt:type_name -> google.protobuf.Timestamp
	7,  // 4: service.OemFeedbackTypeFilter.updatedAt:type_name -> google.protobuf.Timestamp
	7,  // 5: service.OemFeedbackTypeFilter.deletedAt:type_name -> google.protobuf.Timestamp
	0,  // 6: service.OemFeedbackTypeListRequest.query:type_name -> service.OemFeedbackType
	0,  // 7: service.OemFeedbackTypeResponse.data:type_name -> service.OemFeedbackType
	0,  // 8: service.OemFeedbackTypeUpdateFieldsRequest.data:type_name -> service.OemFeedbackType
	5,  // 9: service.OemFeedbackTypeBatchDeleteRequest.keys:type_name -> service.OemFeedbackTypePrimarykey
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_oem_feedback_type_model_gen_proto_init() }
func file_oem_feedback_type_model_gen_proto_init() {
	if File_oem_feedback_type_model_gen_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oem_feedback_type_model_gen_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OemFeedbackType); i {
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
		file_oem_feedback_type_model_gen_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OemFeedbackTypeFilter); i {
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
		file_oem_feedback_type_model_gen_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OemFeedbackTypeListRequest); i {
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
		file_oem_feedback_type_model_gen_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OemFeedbackTypeResponse); i {
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
		file_oem_feedback_type_model_gen_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OemFeedbackTypeUpdateFieldsRequest); i {
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
		file_oem_feedback_type_model_gen_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OemFeedbackTypePrimarykey); i {
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
		file_oem_feedback_type_model_gen_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OemFeedbackTypeBatchDeleteRequest); i {
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
			RawDescriptor: file_oem_feedback_type_model_gen_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_oem_feedback_type_model_gen_proto_goTypes,
		DependencyIndexes: file_oem_feedback_type_model_gen_proto_depIdxs,
		MessageInfos:      file_oem_feedback_type_model_gen_proto_msgTypes,
	}.Build()
	File_oem_feedback_type_model_gen_proto = out.File
	file_oem_feedback_type_model_gen_proto_rawDesc = nil
	file_oem_feedback_type_model_gen_proto_goTypes = nil
	file_oem_feedback_type_model_gen_proto_depIdxs = nil
}
