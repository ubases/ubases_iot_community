// Code generated by protoc,2022-04-18 19:12:08. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: sys_dict_type_model.gen.proto

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
type SysDictType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"dictId" gorm:"primary_key;AUTO_INCREMENT;default:字典主键;comment:字典主键;size:20;"
	DictId int64 `protobuf:"varint,100,opt,name=dictId,proto3" json:"dictId,omitempty"`
	// @inject_tag: json:"dictName" gorm:"default:字典名称;comment:字典名称;size:100;"
	DictName string `protobuf:"bytes,101,opt,name=dictName,proto3" json:"dictName,omitempty"`
	// @inject_tag: json:"dictType" gorm:"default:字典类型;comment:字典类型;size:100;"
	DictType string `protobuf:"bytes,102,opt,name=dictType,proto3" json:"dictType,omitempty"`
	// @inject_tag: json:"status" gorm:"default:状态（0正常 1停用）;comment:状态（0正常 1停用）;"
	Status int32 `protobuf:"varint,103,opt,name=status,proto3" json:"status,omitempty"`
	// @inject_tag: json:"createBy" gorm:"default:创建者;comment:创建者;size:10;"
	CreateBy int32 `protobuf:"varint,104,opt,name=createBy,proto3" json:"createBy,omitempty"`
	// @inject_tag: json:"updateBy" gorm:"default:更新者;comment:更新者;size:10;"
	UpdateBy int32 `protobuf:"varint,105,opt,name=updateBy,proto3" json:"updateBy,omitempty"`
	// @inject_tag: json:"remark" gorm:"default:备注;comment:备注;size:500;"
	Remark string `protobuf:"bytes,106,opt,name=remark,proto3" json:"remark,omitempty"`
	// @inject_tag: json:"createdAt" gorm:"default:创建日期;comment:创建日期;"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,107,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// @inject_tag: json:"updatedAt" gorm:"default:修改日期;comment:修改日期;"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,108,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// @inject_tag: json:"deletedAt" gorm:"default:删除日期;comment:删除日期;"
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,109,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *SysDictType) Reset() {
	*x = SysDictType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_dict_type_model_gen_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysDictType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysDictType) ProtoMessage() {}

func (x *SysDictType) ProtoReflect() protoreflect.Message {
	mi := &file_sys_dict_type_model_gen_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysDictType.ProtoReflect.Descriptor instead.
func (*SysDictType) Descriptor() ([]byte, []int) {
	return file_sys_dict_type_model_gen_proto_rawDescGZIP(), []int{0}
}

func (x *SysDictType) GetDictId() int64 {
	if x != nil {
		return x.DictId
	}
	return 0
}

func (x *SysDictType) GetDictName() string {
	if x != nil {
		return x.DictName
	}
	return ""
}

func (x *SysDictType) GetDictType() string {
	if x != nil {
		return x.DictType
	}
	return ""
}

func (x *SysDictType) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SysDictType) GetCreateBy() int32 {
	if x != nil {
		return x.CreateBy
	}
	return 0
}

func (x *SysDictType) GetUpdateBy() int32 {
	if x != nil {
		return x.UpdateBy
	}
	return 0
}

func (x *SysDictType) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *SysDictType) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *SysDictType) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *SysDictType) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type SysDictTypeFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"dictId" gorm:"primary_key;AUTO_INCREMENT;default:字典主键;comment:字典主键;size:20;"
	DictId int64 `protobuf:"varint,100,opt,name=dictId,proto3" json:"dictId,omitempty"`
	// @inject_tag: json:"dictName" gorm:"default:字典名称;comment:字典名称;size:100;"
	DictName string `protobuf:"bytes,101,opt,name=dictName,proto3" json:"dictName,omitempty"`
	// @inject_tag: json:"dictType" gorm:"default:字典类型;comment:字典类型;size:100;"
	DictType string `protobuf:"bytes,102,opt,name=dictType,proto3" json:"dictType,omitempty"`
	// @inject_tag: json:"status" gorm:"default:状态（0正常 1停用）;comment:状态（0正常 1停用）;"
	Status int32 `protobuf:"varint,103,opt,name=status,proto3" json:"status,omitempty"`
	// @inject_tag: json:"createBy" gorm:"default:创建者;comment:创建者;size:10;"
	CreateBy int32 `protobuf:"varint,104,opt,name=createBy,proto3" json:"createBy,omitempty"`
	// @inject_tag: json:"updateBy" gorm:"default:更新者;comment:更新者;size:10;"
	UpdateBy int32 `protobuf:"varint,105,opt,name=updateBy,proto3" json:"updateBy,omitempty"`
	// @inject_tag: json:"remark" gorm:"default:备注;comment:备注;size:500;"
	Remark string `protobuf:"bytes,106,opt,name=remark,proto3" json:"remark,omitempty"`
	// @inject_tag: json:"createdAt" gorm:"default:创建日期;comment:创建日期;"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,107,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// @inject_tag: json:"updatedAt" gorm:"default:修改日期;comment:修改日期;"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,108,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// @inject_tag: json:"deletedAt" gorm:"default:删除日期;comment:删除日期;"
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,109,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *SysDictTypeFilter) Reset() {
	*x = SysDictTypeFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_dict_type_model_gen_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysDictTypeFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysDictTypeFilter) ProtoMessage() {}

func (x *SysDictTypeFilter) ProtoReflect() protoreflect.Message {
	mi := &file_sys_dict_type_model_gen_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysDictTypeFilter.ProtoReflect.Descriptor instead.
func (*SysDictTypeFilter) Descriptor() ([]byte, []int) {
	return file_sys_dict_type_model_gen_proto_rawDescGZIP(), []int{1}
}

func (x *SysDictTypeFilter) GetDictId() int64 {
	if x != nil {
		return x.DictId
	}
	return 0
}

func (x *SysDictTypeFilter) GetDictName() string {
	if x != nil {
		return x.DictName
	}
	return ""
}

func (x *SysDictTypeFilter) GetDictType() string {
	if x != nil {
		return x.DictType
	}
	return ""
}

func (x *SysDictTypeFilter) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SysDictTypeFilter) GetCreateBy() int32 {
	if x != nil {
		return x.CreateBy
	}
	return 0
}

func (x *SysDictTypeFilter) GetUpdateBy() int32 {
	if x != nil {
		return x.UpdateBy
	}
	return 0
}

func (x *SysDictTypeFilter) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *SysDictTypeFilter) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *SysDictTypeFilter) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *SysDictTypeFilter) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type SysDictTypeListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query     *SysDictType `protobuf:"bytes,101,opt,name=query,proto3" json:"query,omitempty"`
	Page      int64        `protobuf:"varint,102,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int64        `protobuf:"varint,103,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	OrderKey  string       `protobuf:"bytes,104,opt,name=orderKey,proto3" json:"orderKey,omitempty"`
	OrderDesc string       `protobuf:"bytes,105,opt,name=orderDesc,proto3" json:"orderDesc,omitempty"`
}

func (x *SysDictTypeListRequest) Reset() {
	*x = SysDictTypeListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_dict_type_model_gen_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysDictTypeListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysDictTypeListRequest) ProtoMessage() {}

func (x *SysDictTypeListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sys_dict_type_model_gen_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysDictTypeListRequest.ProtoReflect.Descriptor instead.
func (*SysDictTypeListRequest) Descriptor() ([]byte, []int) {
	return file_sys_dict_type_model_gen_proto_rawDescGZIP(), []int{2}
}

func (x *SysDictTypeListRequest) GetQuery() *SysDictType {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *SysDictTypeListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SysDictTypeListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SysDictTypeListRequest) GetOrderKey() string {
	if x != nil {
		return x.OrderKey
	}
	return ""
}

func (x *SysDictTypeListRequest) GetOrderDesc() string {
	if x != nil {
		return x.OrderDesc
	}
	return ""
}

type SysDictTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32          `protobuf:"varint,101,opt,name=code,proto3" json:"code,omitempty"`
	Message string         `protobuf:"bytes,102,opt,name=message,proto3" json:"message,omitempty"`
	Total   int64          `protobuf:"varint,103,opt,name=total,proto3" json:"total,omitempty"`
	Data    []*SysDictType `protobuf:"bytes,104,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *SysDictTypeResponse) Reset() {
	*x = SysDictTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_dict_type_model_gen_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysDictTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysDictTypeResponse) ProtoMessage() {}

func (x *SysDictTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sys_dict_type_model_gen_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysDictTypeResponse.ProtoReflect.Descriptor instead.
func (*SysDictTypeResponse) Descriptor() ([]byte, []int) {
	return file_sys_dict_type_model_gen_proto_rawDescGZIP(), []int{3}
}

func (x *SysDictTypeResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SysDictTypeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SysDictTypeResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *SysDictTypeResponse) GetData() []*SysDictType {
	if x != nil {
		return x.Data
	}
	return nil
}

type SysDictTypeUpdateFieldsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []string     `protobuf:"bytes,101,rep,name=fields,proto3" json:"fields,omitempty"`
	Data   *SysDictType `protobuf:"bytes,102,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SysDictTypeUpdateFieldsRequest) Reset() {
	*x = SysDictTypeUpdateFieldsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_dict_type_model_gen_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysDictTypeUpdateFieldsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysDictTypeUpdateFieldsRequest) ProtoMessage() {}

func (x *SysDictTypeUpdateFieldsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sys_dict_type_model_gen_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysDictTypeUpdateFieldsRequest.ProtoReflect.Descriptor instead.
func (*SysDictTypeUpdateFieldsRequest) Descriptor() ([]byte, []int) {
	return file_sys_dict_type_model_gen_proto_rawDescGZIP(), []int{4}
}

func (x *SysDictTypeUpdateFieldsRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *SysDictTypeUpdateFieldsRequest) GetData() *SysDictType {
	if x != nil {
		return x.Data
	}
	return nil
}

// 表模型主键，只能是整数或字符串，若生成错误，请检查数据库表主键是否正确。支持组合主键
type SysDictTypePrimarykey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DictId int64 `protobuf:"varint,100,opt,name=dictId,proto3" json:"dictId,omitempty"`
}

func (x *SysDictTypePrimarykey) Reset() {
	*x = SysDictTypePrimarykey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_dict_type_model_gen_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysDictTypePrimarykey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysDictTypePrimarykey) ProtoMessage() {}

func (x *SysDictTypePrimarykey) ProtoReflect() protoreflect.Message {
	mi := &file_sys_dict_type_model_gen_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysDictTypePrimarykey.ProtoReflect.Descriptor instead.
func (*SysDictTypePrimarykey) Descriptor() ([]byte, []int) {
	return file_sys_dict_type_model_gen_proto_rawDescGZIP(), []int{5}
}

func (x *SysDictTypePrimarykey) GetDictId() int64 {
	if x != nil {
		return x.DictId
	}
	return 0
}

type SysDictTypeBatchDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []*SysDictTypePrimarykey `protobuf:"bytes,101,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *SysDictTypeBatchDeleteRequest) Reset() {
	*x = SysDictTypeBatchDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_dict_type_model_gen_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysDictTypeBatchDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysDictTypeBatchDeleteRequest) ProtoMessage() {}

func (x *SysDictTypeBatchDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sys_dict_type_model_gen_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysDictTypeBatchDeleteRequest.ProtoReflect.Descriptor instead.
func (*SysDictTypeBatchDeleteRequest) Descriptor() ([]byte, []int) {
	return file_sys_dict_type_model_gen_proto_rawDescGZIP(), []int{6}
}

func (x *SysDictTypeBatchDeleteRequest) GetKeys() []*SysDictTypePrimarykey {
	if x != nil {
		return x.Keys
	}
	return nil
}

var File_sys_dict_type_model_gen_proto protoreflect.FileDescriptor

var file_sys_dict_type_model_gen_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x79, 0x73, 0x5f, 0x64, 0x69, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x02, 0x0a, 0x0b, 0x53, 0x79,
	0x73, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x63,
	0x74, 0x49, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x64, 0x69, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x65, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x67, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x18, 0x68, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x18, 0x69, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d,
	0x61, 0x72, 0x6b, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72,
	0x6b, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x6b,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x6c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x6d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0xf9, 0x02, 0x0a, 0x11, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x63, 0x74, 0x49, 0x64, 0x18,
	0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x64, 0x69, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x69, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x69, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x63,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x63,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x67, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x18, 0x68, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x79, 0x18, 0x69, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18,
	0x6a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x38, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x6b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x6c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x38, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x6d,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xae, 0x01, 0x0a, 0x16,
	0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x18, 0x68,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x1c,
	0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x18, 0x69, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x22, 0x83, 0x01, 0x0a,
	0x13, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x65, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x67, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x68, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x62, 0x0a, 0x1e, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x65,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x28, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2f, 0x0a, 0x15, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x69, 0x63, 0x74, 0x49, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x64, 0x69, 0x63, 0x74, 0x49, 0x64, 0x22, 0x53, 0x0a, 0x1d, 0x53, 0x79, 0x73, 0x44, 0x69,
	0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73,
	0x18, 0x65, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x42, 0x11, 0x5a, 0x0f,
	0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sys_dict_type_model_gen_proto_rawDescOnce sync.Once
	file_sys_dict_type_model_gen_proto_rawDescData = file_sys_dict_type_model_gen_proto_rawDesc
)

func file_sys_dict_type_model_gen_proto_rawDescGZIP() []byte {
	file_sys_dict_type_model_gen_proto_rawDescOnce.Do(func() {
		file_sys_dict_type_model_gen_proto_rawDescData = protoimpl.X.CompressGZIP(file_sys_dict_type_model_gen_proto_rawDescData)
	})
	return file_sys_dict_type_model_gen_proto_rawDescData
}

var file_sys_dict_type_model_gen_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_sys_dict_type_model_gen_proto_goTypes = []interface{}{
	(*SysDictType)(nil),                    // 0: service.SysDictType
	(*SysDictTypeFilter)(nil),              // 1: service.SysDictTypeFilter
	(*SysDictTypeListRequest)(nil),         // 2: service.SysDictTypeListRequest
	(*SysDictTypeResponse)(nil),            // 3: service.SysDictTypeResponse
	(*SysDictTypeUpdateFieldsRequest)(nil), // 4: service.SysDictTypeUpdateFieldsRequest
	(*SysDictTypePrimarykey)(nil),          // 5: service.SysDictTypePrimarykey
	(*SysDictTypeBatchDeleteRequest)(nil),  // 6: service.SysDictTypeBatchDeleteRequest
	(*timestamppb.Timestamp)(nil),          // 7: google.protobuf.Timestamp
}
var file_sys_dict_type_model_gen_proto_depIdxs = []int32{
	7,  // 0: service.SysDictType.createdAt:type_name -> google.protobuf.Timestamp
	7,  // 1: service.SysDictType.updatedAt:type_name -> google.protobuf.Timestamp
	7,  // 2: service.SysDictType.deletedAt:type_name -> google.protobuf.Timestamp
	7,  // 3: service.SysDictTypeFilter.createdAt:type_name -> google.protobuf.Timestamp
	7,  // 4: service.SysDictTypeFilter.updatedAt:type_name -> google.protobuf.Timestamp
	7,  // 5: service.SysDictTypeFilter.deletedAt:type_name -> google.protobuf.Timestamp
	0,  // 6: service.SysDictTypeListRequest.query:type_name -> service.SysDictType
	0,  // 7: service.SysDictTypeResponse.data:type_name -> service.SysDictType
	0,  // 8: service.SysDictTypeUpdateFieldsRequest.data:type_name -> service.SysDictType
	5,  // 9: service.SysDictTypeBatchDeleteRequest.keys:type_name -> service.SysDictTypePrimarykey
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_sys_dict_type_model_gen_proto_init() }
func file_sys_dict_type_model_gen_proto_init() {
	if File_sys_dict_type_model_gen_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sys_dict_type_model_gen_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysDictType); i {
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
		file_sys_dict_type_model_gen_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysDictTypeFilter); i {
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
		file_sys_dict_type_model_gen_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysDictTypeListRequest); i {
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
		file_sys_dict_type_model_gen_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysDictTypeResponse); i {
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
		file_sys_dict_type_model_gen_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysDictTypeUpdateFieldsRequest); i {
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
		file_sys_dict_type_model_gen_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysDictTypePrimarykey); i {
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
		file_sys_dict_type_model_gen_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysDictTypeBatchDeleteRequest); i {
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
			RawDescriptor: file_sys_dict_type_model_gen_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sys_dict_type_model_gen_proto_goTypes,
		DependencyIndexes: file_sys_dict_type_model_gen_proto_depIdxs,
		MessageInfos:      file_sys_dict_type_model_gen_proto_msgTypes,
	}.Build()
	File_sys_dict_type_model_gen_proto = out.File
	file_sys_dict_type_model_gen_proto_rawDesc = nil
	file_sys_dict_type_model_gen_proto_goTypes = nil
	file_sys_dict_type_model_gen_proto_depIdxs = nil
}
