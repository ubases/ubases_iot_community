// Code generated by protoc,2022-05-31 10:08:59. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: config_platform_model.gen.proto

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
type ConfigPlatform struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:主键ID;comment:主键ID;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"code" gorm:"default:平台配置编码码;comment:平台配置编码码;size:50;"
	Code string `protobuf:"bytes,101,opt,name=code,proto3" json:"code,omitempty"`
	// @inject_tag: json:"value" gorm:"default:平台配置值;comment:平台配置值;size:100;"
	Value string `protobuf:"bytes,102,opt,name=value,proto3" json:"value,omitempty"`
	// @inject_tag: json:"remark" gorm:"default:备注;comment:备注;size:255;"
	Remark string `protobuf:"bytes,103,opt,name=remark,proto3" json:"remark,omitempty"`
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

func (x *ConfigPlatform) Reset() {
	*x = ConfigPlatform{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_platform_model_gen_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigPlatform) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigPlatform) ProtoMessage() {}

func (x *ConfigPlatform) ProtoReflect() protoreflect.Message {
	mi := &file_config_platform_model_gen_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigPlatform.ProtoReflect.Descriptor instead.
func (*ConfigPlatform) Descriptor() ([]byte, []int) {
	return file_config_platform_model_gen_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigPlatform) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ConfigPlatform) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ConfigPlatform) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *ConfigPlatform) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *ConfigPlatform) GetCreatedBy() int64 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *ConfigPlatform) GetUpdatedBy() int64 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

func (x *ConfigPlatform) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ConfigPlatform) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *ConfigPlatform) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type ConfigPlatformFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:主键ID;comment:主键ID;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"code" gorm:"default:平台配置编码码;comment:平台配置编码码;size:50;"
	Code string `protobuf:"bytes,101,opt,name=code,proto3" json:"code,omitempty"`
	// @inject_tag: json:"value" gorm:"default:平台配置值;comment:平台配置值;size:100;"
	Value string `protobuf:"bytes,102,opt,name=value,proto3" json:"value,omitempty"`
	// @inject_tag: json:"remark" gorm:"default:备注;comment:备注;size:255;"
	Remark string `protobuf:"bytes,103,opt,name=remark,proto3" json:"remark,omitempty"`
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

func (x *ConfigPlatformFilter) Reset() {
	*x = ConfigPlatformFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_platform_model_gen_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigPlatformFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigPlatformFilter) ProtoMessage() {}

func (x *ConfigPlatformFilter) ProtoReflect() protoreflect.Message {
	mi := &file_config_platform_model_gen_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigPlatformFilter.ProtoReflect.Descriptor instead.
func (*ConfigPlatformFilter) Descriptor() ([]byte, []int) {
	return file_config_platform_model_gen_proto_rawDescGZIP(), []int{1}
}

func (x *ConfigPlatformFilter) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ConfigPlatformFilter) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ConfigPlatformFilter) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *ConfigPlatformFilter) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *ConfigPlatformFilter) GetCreatedBy() int64 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *ConfigPlatformFilter) GetUpdatedBy() int64 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

func (x *ConfigPlatformFilter) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ConfigPlatformFilter) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *ConfigPlatformFilter) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type ConfigPlatformListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query     *ConfigPlatform `protobuf:"bytes,101,opt,name=query,proto3" json:"query,omitempty"`
	Page      int64           `protobuf:"varint,102,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int64           `protobuf:"varint,103,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	OrderKey  string          `protobuf:"bytes,104,opt,name=orderKey,proto3" json:"orderKey,omitempty"`
	OrderDesc string          `protobuf:"bytes,105,opt,name=orderDesc,proto3" json:"orderDesc,omitempty"`
	SearchKey string          `protobuf:"bytes,106,opt,name=searchKey,proto3" json:"searchKey,omitempty"`
}

func (x *ConfigPlatformListRequest) Reset() {
	*x = ConfigPlatformListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_platform_model_gen_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigPlatformListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigPlatformListRequest) ProtoMessage() {}

func (x *ConfigPlatformListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_platform_model_gen_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigPlatformListRequest.ProtoReflect.Descriptor instead.
func (*ConfigPlatformListRequest) Descriptor() ([]byte, []int) {
	return file_config_platform_model_gen_proto_rawDescGZIP(), []int{2}
}

func (x *ConfigPlatformListRequest) GetQuery() *ConfigPlatform {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *ConfigPlatformListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ConfigPlatformListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ConfigPlatformListRequest) GetOrderKey() string {
	if x != nil {
		return x.OrderKey
	}
	return ""
}

func (x *ConfigPlatformListRequest) GetOrderDesc() string {
	if x != nil {
		return x.OrderDesc
	}
	return ""
}

func (x *ConfigPlatformListRequest) GetSearchKey() string {
	if x != nil {
		return x.SearchKey
	}
	return ""
}

type ConfigPlatformResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32             `protobuf:"varint,101,opt,name=code,proto3" json:"code,omitempty"`
	Message string            `protobuf:"bytes,102,opt,name=message,proto3" json:"message,omitempty"`
	Total   int64             `protobuf:"varint,103,opt,name=total,proto3" json:"total,omitempty"`
	Data    []*ConfigPlatform `protobuf:"bytes,104,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ConfigPlatformResponse) Reset() {
	*x = ConfigPlatformResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_platform_model_gen_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigPlatformResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigPlatformResponse) ProtoMessage() {}

func (x *ConfigPlatformResponse) ProtoReflect() protoreflect.Message {
	mi := &file_config_platform_model_gen_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigPlatformResponse.ProtoReflect.Descriptor instead.
func (*ConfigPlatformResponse) Descriptor() ([]byte, []int) {
	return file_config_platform_model_gen_proto_rawDescGZIP(), []int{3}
}

func (x *ConfigPlatformResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ConfigPlatformResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ConfigPlatformResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ConfigPlatformResponse) GetData() []*ConfigPlatform {
	if x != nil {
		return x.Data
	}
	return nil
}

type ConfigPlatformUpdateFieldsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []string        `protobuf:"bytes,101,rep,name=fields,proto3" json:"fields,omitempty"`
	Data   *ConfigPlatform `protobuf:"bytes,102,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ConfigPlatformUpdateFieldsRequest) Reset() {
	*x = ConfigPlatformUpdateFieldsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_platform_model_gen_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigPlatformUpdateFieldsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigPlatformUpdateFieldsRequest) ProtoMessage() {}

func (x *ConfigPlatformUpdateFieldsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_platform_model_gen_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigPlatformUpdateFieldsRequest.ProtoReflect.Descriptor instead.
func (*ConfigPlatformUpdateFieldsRequest) Descriptor() ([]byte, []int) {
	return file_config_platform_model_gen_proto_rawDescGZIP(), []int{4}
}

func (x *ConfigPlatformUpdateFieldsRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *ConfigPlatformUpdateFieldsRequest) GetData() *ConfigPlatform {
	if x != nil {
		return x.Data
	}
	return nil
}

// 表模型主键，只能是整数或字符串，若生成错误，请检查数据库表主键是否正确。支持组合主键
type ConfigPlatformPrimarykey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ConfigPlatformPrimarykey) Reset() {
	*x = ConfigPlatformPrimarykey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_platform_model_gen_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigPlatformPrimarykey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigPlatformPrimarykey) ProtoMessage() {}

func (x *ConfigPlatformPrimarykey) ProtoReflect() protoreflect.Message {
	mi := &file_config_platform_model_gen_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigPlatformPrimarykey.ProtoReflect.Descriptor instead.
func (*ConfigPlatformPrimarykey) Descriptor() ([]byte, []int) {
	return file_config_platform_model_gen_proto_rawDescGZIP(), []int{5}
}

func (x *ConfigPlatformPrimarykey) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ConfigPlatformBatchDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []*ConfigPlatformPrimarykey `protobuf:"bytes,101,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *ConfigPlatformBatchDeleteRequest) Reset() {
	*x = ConfigPlatformBatchDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_platform_model_gen_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigPlatformBatchDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigPlatformBatchDeleteRequest) ProtoMessage() {}

func (x *ConfigPlatformBatchDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_platform_model_gen_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigPlatformBatchDeleteRequest.ProtoReflect.Descriptor instead.
func (*ConfigPlatformBatchDeleteRequest) Descriptor() ([]byte, []int) {
	return file_config_platform_model_gen_proto_rawDescGZIP(), []int{6}
}

func (x *ConfigPlatformBatchDeleteRequest) GetKeys() []*ConfigPlatformPrimarykey {
	if x != nil {
		return x.Keys
	}
	return nil
}

var File_config_platform_model_gen_proto protoreflect.FileDescriptor

var file_config_platform_model_gen_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x02, 0x0a, 0x0e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61,
	0x72, 0x6b, 0x18, 0x67, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x68, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1c,
	0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x69, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x38, 0x0a, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x6b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x38, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x6c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xd2, 0x02, 0x0a, 0x14, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x67, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
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
	0xd2, 0x01, 0x0a, 0x19, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x67, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x44, 0x65, 0x73, 0x63, 0x18, 0x69, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x4b, 0x65, 0x79, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x4b, 0x65, 0x79, 0x22, 0x89, 0x01, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x66,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x68, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x68, 0x0a, 0x21, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18,
	0x65, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x2b, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2a, 0x0a, 0x18, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x50, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x59, 0x0a, 0x20, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x04, 0x6b, 0x65,
	0x79, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x52, 0x04, 0x6b, 0x65, 0x79,
	0x73, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_platform_model_gen_proto_rawDescOnce sync.Once
	file_config_platform_model_gen_proto_rawDescData = file_config_platform_model_gen_proto_rawDesc
)

func file_config_platform_model_gen_proto_rawDescGZIP() []byte {
	file_config_platform_model_gen_proto_rawDescOnce.Do(func() {
		file_config_platform_model_gen_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_platform_model_gen_proto_rawDescData)
	})
	return file_config_platform_model_gen_proto_rawDescData
}

var file_config_platform_model_gen_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_config_platform_model_gen_proto_goTypes = []interface{}{
	(*ConfigPlatform)(nil),                    // 0: service.ConfigPlatform
	(*ConfigPlatformFilter)(nil),              // 1: service.ConfigPlatformFilter
	(*ConfigPlatformListRequest)(nil),         // 2: service.ConfigPlatformListRequest
	(*ConfigPlatformResponse)(nil),            // 3: service.ConfigPlatformResponse
	(*ConfigPlatformUpdateFieldsRequest)(nil), // 4: service.ConfigPlatformUpdateFieldsRequest
	(*ConfigPlatformPrimarykey)(nil),          // 5: service.ConfigPlatformPrimarykey
	(*ConfigPlatformBatchDeleteRequest)(nil),  // 6: service.ConfigPlatformBatchDeleteRequest
	(*timestamppb.Timestamp)(nil),             // 7: google.protobuf.Timestamp
}
var file_config_platform_model_gen_proto_depIdxs = []int32{
	7,  // 0: service.ConfigPlatform.createdAt:type_name -> google.protobuf.Timestamp
	7,  // 1: service.ConfigPlatform.updatedAt:type_name -> google.protobuf.Timestamp
	7,  // 2: service.ConfigPlatform.deletedAt:type_name -> google.protobuf.Timestamp
	7,  // 3: service.ConfigPlatformFilter.createdAt:type_name -> google.protobuf.Timestamp
	7,  // 4: service.ConfigPlatformFilter.updatedAt:type_name -> google.protobuf.Timestamp
	7,  // 5: service.ConfigPlatformFilter.deletedAt:type_name -> google.protobuf.Timestamp
	0,  // 6: service.ConfigPlatformListRequest.query:type_name -> service.ConfigPlatform
	0,  // 7: service.ConfigPlatformResponse.data:type_name -> service.ConfigPlatform
	0,  // 8: service.ConfigPlatformUpdateFieldsRequest.data:type_name -> service.ConfigPlatform
	5,  // 9: service.ConfigPlatformBatchDeleteRequest.keys:type_name -> service.ConfigPlatformPrimarykey
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_config_platform_model_gen_proto_init() }
func file_config_platform_model_gen_proto_init() {
	if File_config_platform_model_gen_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_platform_model_gen_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigPlatform); i {
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
		file_config_platform_model_gen_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigPlatformFilter); i {
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
		file_config_platform_model_gen_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigPlatformListRequest); i {
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
		file_config_platform_model_gen_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigPlatformResponse); i {
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
		file_config_platform_model_gen_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigPlatformUpdateFieldsRequest); i {
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
		file_config_platform_model_gen_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigPlatformPrimarykey); i {
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
		file_config_platform_model_gen_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigPlatformBatchDeleteRequest); i {
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
			RawDescriptor: file_config_platform_model_gen_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_platform_model_gen_proto_goTypes,
		DependencyIndexes: file_config_platform_model_gen_proto_depIdxs,
		MessageInfos:      file_config_platform_model_gen_proto_msgTypes,
	}.Build()
	File_config_platform_model_gen_proto = out.File
	file_config_platform_model_gen_proto_rawDesc = nil
	file_config_platform_model_gen_proto_goTypes = nil
	file_config_platform_model_gen_proto_depIdxs = nil
}
