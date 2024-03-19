// Code generated by protoc,2022-05-17 13:13:13. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: lang_translate_type_model.gen.proto

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
type LangTranslateType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:唯一编号;comment:唯一编号;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"name" gorm:"default:中文描述;comment:中文描述;size:50;"
	Name string `protobuf:"bytes,101,opt,name=name,proto3" json:"name,omitempty"`
	// @inject_tag: json:"code" gorm:"default:英文编码（作为编码使用）;comment:英文编码（作为编码使用）;size:20;"
	Code string `protobuf:"bytes,102,opt,name=code,proto3" json:"code,omitempty"`
	// @inject_tag: json:"status" gorm:"default:1启用/0禁用;comment:1启用/0禁用;size:10;"
	Status int32 `protobuf:"varint,103,opt,name=status,proto3" json:"status,omitempty"`
	// @inject_tag: json:"createdBy" gorm:"default:创建人;comment:创建人;size:20;"
	CreatedBy int64 `protobuf:"varint,104,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	// @inject_tag: json:"updatedBy" gorm:"default:修改人;comment:修改人;size:20;"
	UpdatedBy int64 `protobuf:"varint,105,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	// @inject_tag: json:"createdAt" gorm:"default:创建时间;comment:创建时间;"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,106,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// @inject_tag: json:"updatedAt" gorm:"default:修改时间;comment:修改时间;"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,107,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// @inject_tag: json:"deletedAt" gorm:"default:删除时间;comment:删除时间;"
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,108,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *LangTranslateType) Reset() {
	*x = LangTranslateType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lang_translate_type_model_gen_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LangTranslateType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LangTranslateType) ProtoMessage() {}

func (x *LangTranslateType) ProtoReflect() protoreflect.Message {
	mi := &file_lang_translate_type_model_gen_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LangTranslateType.ProtoReflect.Descriptor instead.
func (*LangTranslateType) Descriptor() ([]byte, []int) {
	return file_lang_translate_type_model_gen_proto_rawDescGZIP(), []int{0}
}

func (x *LangTranslateType) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LangTranslateType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LangTranslateType) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *LangTranslateType) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *LangTranslateType) GetCreatedBy() int64 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *LangTranslateType) GetUpdatedBy() int64 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

func (x *LangTranslateType) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *LangTranslateType) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *LangTranslateType) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type LangTranslateTypeFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:唯一编号;comment:唯一编号;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"name" gorm:"default:中文描述;comment:中文描述;size:50;"
	Name string `protobuf:"bytes,101,opt,name=name,proto3" json:"name,omitempty"`
	// @inject_tag: json:"code" gorm:"default:英文编码（作为编码使用）;comment:英文编码（作为编码使用）;size:20;"
	Code string `protobuf:"bytes,102,opt,name=code,proto3" json:"code,omitempty"`
	// @inject_tag: json:"status" gorm:"default:1启用/0禁用;comment:1启用/0禁用;size:10;"
	Status int32 `protobuf:"varint,103,opt,name=status,proto3" json:"status,omitempty"`
	// @inject_tag: json:"createdBy" gorm:"default:创建人;comment:创建人;size:20;"
	CreatedBy int64 `protobuf:"varint,104,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	// @inject_tag: json:"updatedBy" gorm:"default:修改人;comment:修改人;size:20;"
	UpdatedBy int64 `protobuf:"varint,105,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	// @inject_tag: json:"createdAt" gorm:"default:创建时间;comment:创建时间;"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,106,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// @inject_tag: json:"updatedAt" gorm:"default:修改时间;comment:修改时间;"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,107,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// @inject_tag: json:"deletedAt" gorm:"default:删除时间;comment:删除时间;"
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,108,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *LangTranslateTypeFilter) Reset() {
	*x = LangTranslateTypeFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lang_translate_type_model_gen_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LangTranslateTypeFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LangTranslateTypeFilter) ProtoMessage() {}

func (x *LangTranslateTypeFilter) ProtoReflect() protoreflect.Message {
	mi := &file_lang_translate_type_model_gen_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LangTranslateTypeFilter.ProtoReflect.Descriptor instead.
func (*LangTranslateTypeFilter) Descriptor() ([]byte, []int) {
	return file_lang_translate_type_model_gen_proto_rawDescGZIP(), []int{1}
}

func (x *LangTranslateTypeFilter) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LangTranslateTypeFilter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LangTranslateTypeFilter) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *LangTranslateTypeFilter) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *LangTranslateTypeFilter) GetCreatedBy() int64 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *LangTranslateTypeFilter) GetUpdatedBy() int64 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

func (x *LangTranslateTypeFilter) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *LangTranslateTypeFilter) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *LangTranslateTypeFilter) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type LangTranslateTypeListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query     *LangTranslateType `protobuf:"bytes,101,opt,name=query,proto3" json:"query,omitempty"`
	Page      int64              `protobuf:"varint,102,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int64              `protobuf:"varint,103,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	OrderKey  string             `protobuf:"bytes,104,opt,name=orderKey,proto3" json:"orderKey,omitempty"`
	OrderDesc string             `protobuf:"bytes,105,opt,name=orderDesc,proto3" json:"orderDesc,omitempty"`
	SearchKey string             `protobuf:"bytes,106,opt,name=searchKey,proto3" json:"searchKey,omitempty"`
}

func (x *LangTranslateTypeListRequest) Reset() {
	*x = LangTranslateTypeListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lang_translate_type_model_gen_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LangTranslateTypeListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LangTranslateTypeListRequest) ProtoMessage() {}

func (x *LangTranslateTypeListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lang_translate_type_model_gen_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LangTranslateTypeListRequest.ProtoReflect.Descriptor instead.
func (*LangTranslateTypeListRequest) Descriptor() ([]byte, []int) {
	return file_lang_translate_type_model_gen_proto_rawDescGZIP(), []int{2}
}

func (x *LangTranslateTypeListRequest) GetQuery() *LangTranslateType {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *LangTranslateTypeListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *LangTranslateTypeListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *LangTranslateTypeListRequest) GetOrderKey() string {
	if x != nil {
		return x.OrderKey
	}
	return ""
}

func (x *LangTranslateTypeListRequest) GetOrderDesc() string {
	if x != nil {
		return x.OrderDesc
	}
	return ""
}

func (x *LangTranslateTypeListRequest) GetSearchKey() string {
	if x != nil {
		return x.SearchKey
	}
	return ""
}

type LangTranslateTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32                `protobuf:"varint,101,opt,name=code,proto3" json:"code,omitempty"`
	Message string               `protobuf:"bytes,102,opt,name=message,proto3" json:"message,omitempty"`
	Total   int64                `protobuf:"varint,103,opt,name=total,proto3" json:"total,omitempty"`
	Data    []*LangTranslateType `protobuf:"bytes,104,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *LangTranslateTypeResponse) Reset() {
	*x = LangTranslateTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lang_translate_type_model_gen_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LangTranslateTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LangTranslateTypeResponse) ProtoMessage() {}

func (x *LangTranslateTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lang_translate_type_model_gen_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LangTranslateTypeResponse.ProtoReflect.Descriptor instead.
func (*LangTranslateTypeResponse) Descriptor() ([]byte, []int) {
	return file_lang_translate_type_model_gen_proto_rawDescGZIP(), []int{3}
}

func (x *LangTranslateTypeResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *LangTranslateTypeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LangTranslateTypeResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *LangTranslateTypeResponse) GetData() []*LangTranslateType {
	if x != nil {
		return x.Data
	}
	return nil
}

type LangTranslateTypeUpdateFieldsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []string           `protobuf:"bytes,101,rep,name=fields,proto3" json:"fields,omitempty"`
	Data   *LangTranslateType `protobuf:"bytes,102,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *LangTranslateTypeUpdateFieldsRequest) Reset() {
	*x = LangTranslateTypeUpdateFieldsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lang_translate_type_model_gen_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LangTranslateTypeUpdateFieldsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LangTranslateTypeUpdateFieldsRequest) ProtoMessage() {}

func (x *LangTranslateTypeUpdateFieldsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lang_translate_type_model_gen_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LangTranslateTypeUpdateFieldsRequest.ProtoReflect.Descriptor instead.
func (*LangTranslateTypeUpdateFieldsRequest) Descriptor() ([]byte, []int) {
	return file_lang_translate_type_model_gen_proto_rawDescGZIP(), []int{4}
}

func (x *LangTranslateTypeUpdateFieldsRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *LangTranslateTypeUpdateFieldsRequest) GetData() *LangTranslateType {
	if x != nil {
		return x.Data
	}
	return nil
}

// 表模型主键，只能是整数或字符串，若生成错误，请检查数据库表主键是否正确。支持组合主键
type LangTranslateTypePrimarykey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *LangTranslateTypePrimarykey) Reset() {
	*x = LangTranslateTypePrimarykey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lang_translate_type_model_gen_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LangTranslateTypePrimarykey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LangTranslateTypePrimarykey) ProtoMessage() {}

func (x *LangTranslateTypePrimarykey) ProtoReflect() protoreflect.Message {
	mi := &file_lang_translate_type_model_gen_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LangTranslateTypePrimarykey.ProtoReflect.Descriptor instead.
func (*LangTranslateTypePrimarykey) Descriptor() ([]byte, []int) {
	return file_lang_translate_type_model_gen_proto_rawDescGZIP(), []int{5}
}

func (x *LangTranslateTypePrimarykey) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type LangTranslateTypeBatchDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []*LangTranslateTypePrimarykey `protobuf:"bytes,101,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *LangTranslateTypeBatchDeleteRequest) Reset() {
	*x = LangTranslateTypeBatchDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lang_translate_type_model_gen_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LangTranslateTypeBatchDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LangTranslateTypeBatchDeleteRequest) ProtoMessage() {}

func (x *LangTranslateTypeBatchDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lang_translate_type_model_gen_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LangTranslateTypeBatchDeleteRequest.ProtoReflect.Descriptor instead.
func (*LangTranslateTypeBatchDeleteRequest) Descriptor() ([]byte, []int) {
	return file_lang_translate_type_model_gen_proto_rawDescGZIP(), []int{6}
}

func (x *LangTranslateTypeBatchDeleteRequest) GetKeys() []*LangTranslateTypePrimarykey {
	if x != nil {
		return x.Keys
	}
	return nil
}

var File_lang_translate_type_model_gen_proto protoreflect.FileDescriptor

var file_lang_translate_type_model_gen_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x67, 0x65, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xcd, 0x02, 0x0a, 0x11, 0x4c, 0x61, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x65, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x67, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x18, 0x68, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79,
	0x18, 0x69, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x6a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x6b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x6c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0xd3, 0x02, 0x0a, 0x17, 0x4c, 0x61, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x67, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x68, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x69, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x6b,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x6c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xd8, 0x01, 0x0a, 0x1c, 0x4c, 0x61, 0x6e, 0x67, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x4c, 0x61, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x66, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x4b, 0x65, 0x79, 0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x73,
	0x63, 0x18, 0x69, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65,
	0x73, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x18,
	0x6a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79,
	0x22, 0x8f, 0x01, 0x0a, 0x19, 0x4c, 0x61, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61,
	0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x66, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x68, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x6e, 0x0a, 0x24, 0x4c, 0x61, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c,
	0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x2d, 0x0a, 0x1b, 0x4c, 0x61, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c,
	0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x5f, 0x0a, 0x23, 0x4c, 0x61, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61,
	0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73,
	0x18, 0x65, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x52, 0x04, 0x6b, 0x65,
	0x79, 0x73, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lang_translate_type_model_gen_proto_rawDescOnce sync.Once
	file_lang_translate_type_model_gen_proto_rawDescData = file_lang_translate_type_model_gen_proto_rawDesc
)

func file_lang_translate_type_model_gen_proto_rawDescGZIP() []byte {
	file_lang_translate_type_model_gen_proto_rawDescOnce.Do(func() {
		file_lang_translate_type_model_gen_proto_rawDescData = protoimpl.X.CompressGZIP(file_lang_translate_type_model_gen_proto_rawDescData)
	})
	return file_lang_translate_type_model_gen_proto_rawDescData
}

var file_lang_translate_type_model_gen_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_lang_translate_type_model_gen_proto_goTypes = []interface{}{
	(*LangTranslateType)(nil),                    // 0: service.LangTranslateType
	(*LangTranslateTypeFilter)(nil),              // 1: service.LangTranslateTypeFilter
	(*LangTranslateTypeListRequest)(nil),         // 2: service.LangTranslateTypeListRequest
	(*LangTranslateTypeResponse)(nil),            // 3: service.LangTranslateTypeResponse
	(*LangTranslateTypeUpdateFieldsRequest)(nil), // 4: service.LangTranslateTypeUpdateFieldsRequest
	(*LangTranslateTypePrimarykey)(nil),          // 5: service.LangTranslateTypePrimarykey
	(*LangTranslateTypeBatchDeleteRequest)(nil),  // 6: service.LangTranslateTypeBatchDeleteRequest
	(*timestamppb.Timestamp)(nil),                // 7: google.protobuf.Timestamp
}
var file_lang_translate_type_model_gen_proto_depIdxs = []int32{
	7,  // 0: service.LangTranslateType.createdAt:type_name -> google.protobuf.Timestamp
	7,  // 1: service.LangTranslateType.updatedAt:type_name -> google.protobuf.Timestamp
	7,  // 2: service.LangTranslateType.deletedAt:type_name -> google.protobuf.Timestamp
	7,  // 3: service.LangTranslateTypeFilter.createdAt:type_name -> google.protobuf.Timestamp
	7,  // 4: service.LangTranslateTypeFilter.updatedAt:type_name -> google.protobuf.Timestamp
	7,  // 5: service.LangTranslateTypeFilter.deletedAt:type_name -> google.protobuf.Timestamp
	0,  // 6: service.LangTranslateTypeListRequest.query:type_name -> service.LangTranslateType
	0,  // 7: service.LangTranslateTypeResponse.data:type_name -> service.LangTranslateType
	0,  // 8: service.LangTranslateTypeUpdateFieldsRequest.data:type_name -> service.LangTranslateType
	5,  // 9: service.LangTranslateTypeBatchDeleteRequest.keys:type_name -> service.LangTranslateTypePrimarykey
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_lang_translate_type_model_gen_proto_init() }
func file_lang_translate_type_model_gen_proto_init() {
	if File_lang_translate_type_model_gen_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lang_translate_type_model_gen_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LangTranslateType); i {
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
		file_lang_translate_type_model_gen_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LangTranslateTypeFilter); i {
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
		file_lang_translate_type_model_gen_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LangTranslateTypeListRequest); i {
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
		file_lang_translate_type_model_gen_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LangTranslateTypeResponse); i {
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
		file_lang_translate_type_model_gen_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LangTranslateTypeUpdateFieldsRequest); i {
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
		file_lang_translate_type_model_gen_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LangTranslateTypePrimarykey); i {
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
		file_lang_translate_type_model_gen_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LangTranslateTypeBatchDeleteRequest); i {
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
			RawDescriptor: file_lang_translate_type_model_gen_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lang_translate_type_model_gen_proto_goTypes,
		DependencyIndexes: file_lang_translate_type_model_gen_proto_depIdxs,
		MessageInfos:      file_lang_translate_type_model_gen_proto_msgTypes,
	}.Build()
	File_lang_translate_type_model_gen_proto = out.File
	file_lang_translate_type_model_gen_proto_rawDesc = nil
	file_lang_translate_type_model_gen_proto_goTypes = nil
	file_lang_translate_type_model_gen_proto_depIdxs = nil
}
