// Code generated by protoc,2022-07-14 15:10:36. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: oem_app_entry_seting_model.proto

package protosService

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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
type OemAppEntrySeting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:设置id;comment:设置id;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"dirId" gorm:"default:目录id;comment:目录id;size:19;"
	DirId int64 `protobuf:"varint,101,opt,name=dirId,proto3" json:"dirId,omitempty"`
	// @inject_tag: json:"sort" gorm:"default:排序;comment:排序;size:10;"
	Sort int32 `protobuf:"varint,102,opt,name=sort,proto3" json:"sort,omitempty"`
	// @inject_tag: json:"isEnable" gorm:"default:是否启用;comment:是否启用;"
	IsEnable int32 `protobuf:"varint,103,opt,name=isEnable,proto3" json:"isEnable,omitempty"`
	// @inject_tag: json:"isNormal" gorm:"default:是否设置为常见;comment:是否设置为常见;"
	IsNormal int32 `protobuf:"varint,104,opt,name=isNormal,proto3" json:"isNormal,omitempty"`
}

func (x *OemAppEntrySeting) Reset() {
	*x = OemAppEntrySeting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oem_app_entry_seting_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OemAppEntrySeting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OemAppEntrySeting) ProtoMessage() {}

func (x *OemAppEntrySeting) ProtoReflect() protoreflect.Message {
	mi := &file_oem_app_entry_seting_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OemAppEntrySeting.ProtoReflect.Descriptor instead.
func (*OemAppEntrySeting) Descriptor() ([]byte, []int) {
	return file_oem_app_entry_seting_model_proto_rawDescGZIP(), []int{0}
}

func (x *OemAppEntrySeting) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OemAppEntrySeting) GetDirId() int64 {
	if x != nil {
		return x.DirId
	}
	return 0
}

func (x *OemAppEntrySeting) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *OemAppEntrySeting) GetIsEnable() int32 {
	if x != nil {
		return x.IsEnable
	}
	return 0
}

func (x *OemAppEntrySeting) GetIsNormal() int32 {
	if x != nil {
		return x.IsNormal
	}
	return 0
}

type OemAppEntrySetingFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:设置id;comment:设置id;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"dirId" gorm:"default:目录id;comment:目录id;size:19;"
	DirId int64 `protobuf:"varint,101,opt,name=dirId,proto3" json:"dirId,omitempty"`
	// @inject_tag: json:"sort" gorm:"default:排序;comment:排序;size:10;"
	Sort int32 `protobuf:"varint,102,opt,name=sort,proto3" json:"sort,omitempty"`
	// @inject_tag: json:"isEnable" gorm:"default:是否启用;comment:是否启用;"
	IsEnable int32 `protobuf:"varint,103,opt,name=isEnable,proto3" json:"isEnable,omitempty"`
	// @inject_tag: json:"isNormal" gorm:"default:是否设置为常见;comment:是否设置为常见;"
	IsNormal int32 `protobuf:"varint,104,opt,name=isNormal,proto3" json:"isNormal,omitempty"`
}

func (x *OemAppEntrySetingFilter) Reset() {
	*x = OemAppEntrySetingFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oem_app_entry_seting_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OemAppEntrySetingFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OemAppEntrySetingFilter) ProtoMessage() {}

func (x *OemAppEntrySetingFilter) ProtoReflect() protoreflect.Message {
	mi := &file_oem_app_entry_seting_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OemAppEntrySetingFilter.ProtoReflect.Descriptor instead.
func (*OemAppEntrySetingFilter) Descriptor() ([]byte, []int) {
	return file_oem_app_entry_seting_model_proto_rawDescGZIP(), []int{1}
}

func (x *OemAppEntrySetingFilter) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OemAppEntrySetingFilter) GetDirId() int64 {
	if x != nil {
		return x.DirId
	}
	return 0
}

func (x *OemAppEntrySetingFilter) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *OemAppEntrySetingFilter) GetIsEnable() int32 {
	if x != nil {
		return x.IsEnable
	}
	return 0
}

func (x *OemAppEntrySetingFilter) GetIsNormal() int32 {
	if x != nil {
		return x.IsNormal
	}
	return 0
}

type OemAppEntrySetingListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query     *OemAppEntrySeting `protobuf:"bytes,101,opt,name=query,proto3" json:"query,omitempty"`
	Page      int64              `protobuf:"varint,102,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int64              `protobuf:"varint,103,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	OrderKey  string             `protobuf:"bytes,104,opt,name=orderKey,proto3" json:"orderKey,omitempty"`
	OrderDesc string             `protobuf:"bytes,105,opt,name=orderDesc,proto3" json:"orderDesc,omitempty"`
	SearchKey string             `protobuf:"bytes,106,opt,name=searchKey,proto3" json:"searchKey,omitempty"`
}

func (x *OemAppEntrySetingListRequest) Reset() {
	*x = OemAppEntrySetingListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oem_app_entry_seting_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OemAppEntrySetingListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OemAppEntrySetingListRequest) ProtoMessage() {}

func (x *OemAppEntrySetingListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oem_app_entry_seting_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OemAppEntrySetingListRequest.ProtoReflect.Descriptor instead.
func (*OemAppEntrySetingListRequest) Descriptor() ([]byte, []int) {
	return file_oem_app_entry_seting_model_proto_rawDescGZIP(), []int{2}
}

func (x *OemAppEntrySetingListRequest) GetQuery() *OemAppEntrySeting {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *OemAppEntrySetingListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *OemAppEntrySetingListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *OemAppEntrySetingListRequest) GetOrderKey() string {
	if x != nil {
		return x.OrderKey
	}
	return ""
}

func (x *OemAppEntrySetingListRequest) GetOrderDesc() string {
	if x != nil {
		return x.OrderDesc
	}
	return ""
}

func (x *OemAppEntrySetingListRequest) GetSearchKey() string {
	if x != nil {
		return x.SearchKey
	}
	return ""
}

type OemAppEntrySetingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32                `protobuf:"varint,101,opt,name=code,proto3" json:"code,omitempty"`
	Message string               `protobuf:"bytes,102,opt,name=message,proto3" json:"message,omitempty"`
	Total   int64                `protobuf:"varint,103,opt,name=total,proto3" json:"total,omitempty"`
	Data    []*OemAppEntrySeting `protobuf:"bytes,104,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *OemAppEntrySetingResponse) Reset() {
	*x = OemAppEntrySetingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oem_app_entry_seting_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OemAppEntrySetingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OemAppEntrySetingResponse) ProtoMessage() {}

func (x *OemAppEntrySetingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_oem_app_entry_seting_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OemAppEntrySetingResponse.ProtoReflect.Descriptor instead.
func (*OemAppEntrySetingResponse) Descriptor() ([]byte, []int) {
	return file_oem_app_entry_seting_model_proto_rawDescGZIP(), []int{3}
}

func (x *OemAppEntrySetingResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *OemAppEntrySetingResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *OemAppEntrySetingResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *OemAppEntrySetingResponse) GetData() []*OemAppEntrySeting {
	if x != nil {
		return x.Data
	}
	return nil
}

type OemAppEntrySetingUpdateFieldsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []string           `protobuf:"bytes,101,rep,name=fields,proto3" json:"fields,omitempty"`
	Data   *OemAppEntrySeting `protobuf:"bytes,102,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OemAppEntrySetingUpdateFieldsRequest) Reset() {
	*x = OemAppEntrySetingUpdateFieldsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oem_app_entry_seting_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OemAppEntrySetingUpdateFieldsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OemAppEntrySetingUpdateFieldsRequest) ProtoMessage() {}

func (x *OemAppEntrySetingUpdateFieldsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oem_app_entry_seting_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OemAppEntrySetingUpdateFieldsRequest.ProtoReflect.Descriptor instead.
func (*OemAppEntrySetingUpdateFieldsRequest) Descriptor() ([]byte, []int) {
	return file_oem_app_entry_seting_model_proto_rawDescGZIP(), []int{4}
}

func (x *OemAppEntrySetingUpdateFieldsRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *OemAppEntrySetingUpdateFieldsRequest) GetData() *OemAppEntrySeting {
	if x != nil {
		return x.Data
	}
	return nil
}

// 表模型主键，只能是整数或字符串，若生成错误，请检查数据库表主键是否正确。支持组合主键
type OemAppEntrySetingPrimarykey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OemAppEntrySetingPrimarykey) Reset() {
	*x = OemAppEntrySetingPrimarykey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oem_app_entry_seting_model_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OemAppEntrySetingPrimarykey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OemAppEntrySetingPrimarykey) ProtoMessage() {}

func (x *OemAppEntrySetingPrimarykey) ProtoReflect() protoreflect.Message {
	mi := &file_oem_app_entry_seting_model_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OemAppEntrySetingPrimarykey.ProtoReflect.Descriptor instead.
func (*OemAppEntrySetingPrimarykey) Descriptor() ([]byte, []int) {
	return file_oem_app_entry_seting_model_proto_rawDescGZIP(), []int{5}
}

func (x *OemAppEntrySetingPrimarykey) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type OemAppEntrySetingBatchDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []*OemAppEntrySetingPrimarykey `protobuf:"bytes,101,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *OemAppEntrySetingBatchDeleteRequest) Reset() {
	*x = OemAppEntrySetingBatchDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oem_app_entry_seting_model_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OemAppEntrySetingBatchDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OemAppEntrySetingBatchDeleteRequest) ProtoMessage() {}

func (x *OemAppEntrySetingBatchDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oem_app_entry_seting_model_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OemAppEntrySetingBatchDeleteRequest.ProtoReflect.Descriptor instead.
func (*OemAppEntrySetingBatchDeleteRequest) Descriptor() ([]byte, []int) {
	return file_oem_app_entry_seting_model_proto_rawDescGZIP(), []int{6}
}

func (x *OemAppEntrySetingBatchDeleteRequest) GetKeys() []*OemAppEntrySetingPrimarykey {
	if x != nil {
		return x.Keys
	}
	return nil
}

type OemAppEntrySetingBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OemAppEntrySetings []*OemAppEntrySeting `protobuf:"bytes,101,rep,name=OemAppEntrySetings,proto3" json:"OemAppEntrySetings,omitempty"`
}

func (x *OemAppEntrySetingBatchRequest) Reset() {
	*x = OemAppEntrySetingBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oem_app_entry_seting_model_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OemAppEntrySetingBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OemAppEntrySetingBatchRequest) ProtoMessage() {}

func (x *OemAppEntrySetingBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oem_app_entry_seting_model_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OemAppEntrySetingBatchRequest.ProtoReflect.Descriptor instead.
func (*OemAppEntrySetingBatchRequest) Descriptor() ([]byte, []int) {
	return file_oem_app_entry_seting_model_proto_rawDescGZIP(), []int{7}
}

func (x *OemAppEntrySetingBatchRequest) GetOemAppEntrySetings() []*OemAppEntrySeting {
	if x != nil {
		return x.OemAppEntrySetings
	}
	return nil
}

var File_oem_app_entry_seting_model_proto protoreflect.FileDescriptor

var file_oem_app_entry_seting_model_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6f, 0x65, 0x6d, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f,
	0x73, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x85, 0x01, 0x0a, 0x11,
	0x4f, 0x65, 0x6d, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x69, 0x72, 0x49, 0x64, 0x18, 0x65, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x64, 0x69, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18,
	0x66, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69,
	0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69,
	0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x4e, 0x6f, 0x72,
	0x6d, 0x61, 0x6c, 0x18, 0x68, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x73, 0x4e, 0x6f, 0x72,
	0x6d, 0x61, 0x6c, 0x22, 0x8b, 0x01, 0x0a, 0x17, 0x4f, 0x65, 0x6d, 0x41, 0x70, 0x70, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x53, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x64, 0x69, 0x72, 0x49, 0x64, 0x18, 0x65, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x64, 0x69, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x66, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x45,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x73, 0x45,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x4e, 0x6f, 0x72, 0x6d, 0x61,
	0x6c, 0x18, 0x68, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x73, 0x4e, 0x6f, 0x72, 0x6d, 0x61,
	0x6c, 0x22, 0xd8, 0x01, 0x0a, 0x1c, 0x4f, 0x65, 0x6d, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x53, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x65, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x65, 0x6d, 0x41,
	0x70, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x66, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79,
	0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79,
	0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x18, 0x69, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x18, 0x6a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x22, 0x8f, 0x01, 0x0a,
	0x19, 0x4f, 0x65, 0x6d, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2e,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x68, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x65, 0x6d, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x53, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x6e,
	0x0a, 0x24, 0x4f, 0x65, 0x6d, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x18, 0x65, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x2e,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x65, 0x6d, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x53, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2d,
	0x0a, 0x1b, 0x4f, 0x65, 0x6d, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5f, 0x0a,
	0x23, 0x4f, 0x65, 0x6d, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x65, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x65, 0x6d,
	0x41, 0x70, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x22, 0x6b,
	0x0a, 0x1d, 0x4f, 0x65, 0x6d, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x4a, 0x0a, 0x12, 0x4f, 0x65, 0x6d, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x65, 0x6d, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x53, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x12, 0x4f, 0x65, 0x6d, 0x41, 0x70, 0x70, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x53, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x11, 0x5a, 0x0f, 0x2e,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_oem_app_entry_seting_model_proto_rawDescOnce sync.Once
	file_oem_app_entry_seting_model_proto_rawDescData = file_oem_app_entry_seting_model_proto_rawDesc
)

func file_oem_app_entry_seting_model_proto_rawDescGZIP() []byte {
	file_oem_app_entry_seting_model_proto_rawDescOnce.Do(func() {
		file_oem_app_entry_seting_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_oem_app_entry_seting_model_proto_rawDescData)
	})
	return file_oem_app_entry_seting_model_proto_rawDescData
}

var file_oem_app_entry_seting_model_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_oem_app_entry_seting_model_proto_goTypes = []interface{}{
	(*OemAppEntrySeting)(nil),                    // 0: service.OemAppEntrySeting
	(*OemAppEntrySetingFilter)(nil),              // 1: service.OemAppEntrySetingFilter
	(*OemAppEntrySetingListRequest)(nil),         // 2: service.OemAppEntrySetingListRequest
	(*OemAppEntrySetingResponse)(nil),            // 3: service.OemAppEntrySetingResponse
	(*OemAppEntrySetingUpdateFieldsRequest)(nil), // 4: service.OemAppEntrySetingUpdateFieldsRequest
	(*OemAppEntrySetingPrimarykey)(nil),          // 5: service.OemAppEntrySetingPrimarykey
	(*OemAppEntrySetingBatchDeleteRequest)(nil),  // 6: service.OemAppEntrySetingBatchDeleteRequest
	(*OemAppEntrySetingBatchRequest)(nil),        // 7: service.OemAppEntrySetingBatchRequest
}
var file_oem_app_entry_seting_model_proto_depIdxs = []int32{
	0, // 0: service.OemAppEntrySetingListRequest.query:type_name -> service.OemAppEntrySeting
	0, // 1: service.OemAppEntrySetingResponse.data:type_name -> service.OemAppEntrySeting
	0, // 2: service.OemAppEntrySetingUpdateFieldsRequest.data:type_name -> service.OemAppEntrySeting
	5, // 3: service.OemAppEntrySetingBatchDeleteRequest.keys:type_name -> service.OemAppEntrySetingPrimarykey
	0, // 4: service.OemAppEntrySetingBatchRequest.OemAppEntrySetings:type_name -> service.OemAppEntrySeting
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_oem_app_entry_seting_model_proto_init() }
func file_oem_app_entry_seting_model_proto_init() {
	if File_oem_app_entry_seting_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oem_app_entry_seting_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OemAppEntrySeting); i {
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
		file_oem_app_entry_seting_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OemAppEntrySetingFilter); i {
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
		file_oem_app_entry_seting_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OemAppEntrySetingListRequest); i {
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
		file_oem_app_entry_seting_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OemAppEntrySetingResponse); i {
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
		file_oem_app_entry_seting_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OemAppEntrySetingUpdateFieldsRequest); i {
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
		file_oem_app_entry_seting_model_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OemAppEntrySetingPrimarykey); i {
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
		file_oem_app_entry_seting_model_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OemAppEntrySetingBatchDeleteRequest); i {
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
		file_oem_app_entry_seting_model_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OemAppEntrySetingBatchRequest); i {
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
			RawDescriptor: file_oem_app_entry_seting_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_oem_app_entry_seting_model_proto_goTypes,
		DependencyIndexes: file_oem_app_entry_seting_model_proto_depIdxs,
		MessageInfos:      file_oem_app_entry_seting_model_proto_msgTypes,
	}.Build()
	File_oem_app_entry_seting_model_proto = out.File
	file_oem_app_entry_seting_model_proto_rawDesc = nil
	file_oem_app_entry_seting_model_proto_goTypes = nil
	file_oem_app_entry_seting_model_proto_depIdxs = nil
}
