// Code generated by protoc,2022-09-26 11:41:05. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: uc_appleid_info_model.gen.proto

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
type UcAppleidInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:主键ID;comment:主键ID;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"thirdUserId" gorm:"default:第三方用户ID;comment:第三方用户ID;size:50;"
	ThirdUserId string `protobuf:"bytes,101,opt,name=thirdUserId,proto3" json:"thirdUserId,omitempty"`
	// @inject_tag: json:"nickname" gorm:"default:第三方用户昵称;comment:第三方用户昵称;size:50;"
	Nickname string `protobuf:"bytes,102,opt,name=nickname,proto3" json:"nickname,omitempty"`
	// @inject_tag: json:"createdBy" gorm:"default:创建人;comment:创建人;size:20;"
	CreatedBy int64 `protobuf:"varint,103,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	// @inject_tag: json:"updatedBy" gorm:"default:修改人;comment:修改人;size:20;"
	UpdatedBy int64 `protobuf:"varint,104,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	// @inject_tag: json:"createdAt" gorm:"default:创建时间;comment:创建时间;"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,105,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// @inject_tag: json:"updatedAt" gorm:"default:修改时间;comment:修改时间;"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,106,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// @inject_tag: json:"deletedAt" gorm:"default:删除时间;comment:删除时间;"
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,107,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *UcAppleidInfo) Reset() {
	*x = UcAppleidInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uc_appleid_info_model_gen_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UcAppleidInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UcAppleidInfo) ProtoMessage() {}

func (x *UcAppleidInfo) ProtoReflect() protoreflect.Message {
	mi := &file_uc_appleid_info_model_gen_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UcAppleidInfo.ProtoReflect.Descriptor instead.
func (*UcAppleidInfo) Descriptor() ([]byte, []int) {
	return file_uc_appleid_info_model_gen_proto_rawDescGZIP(), []int{0}
}

func (x *UcAppleidInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UcAppleidInfo) GetThirdUserId() string {
	if x != nil {
		return x.ThirdUserId
	}
	return ""
}

func (x *UcAppleidInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *UcAppleidInfo) GetCreatedBy() int64 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *UcAppleidInfo) GetUpdatedBy() int64 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

func (x *UcAppleidInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UcAppleidInfo) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *UcAppleidInfo) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type UcAppleidInfoFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:主键ID;comment:主键ID;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"thirdUserId" gorm:"default:第三方用户ID;comment:第三方用户ID;size:50;"
	ThirdUserId string `protobuf:"bytes,101,opt,name=thirdUserId,proto3" json:"thirdUserId,omitempty"`
	// @inject_tag: json:"nickname" gorm:"default:第三方用户昵称;comment:第三方用户昵称;size:50;"
	Nickname string `protobuf:"bytes,102,opt,name=nickname,proto3" json:"nickname,omitempty"`
	// @inject_tag: json:"createdBy" gorm:"default:创建人;comment:创建人;size:20;"
	CreatedBy int64 `protobuf:"varint,103,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	// @inject_tag: json:"updatedBy" gorm:"default:修改人;comment:修改人;size:20;"
	UpdatedBy int64 `protobuf:"varint,104,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	// @inject_tag: json:"createdAt" gorm:"default:创建时间;comment:创建时间;"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,105,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// @inject_tag: json:"updatedAt" gorm:"default:修改时间;comment:修改时间;"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,106,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// @inject_tag: json:"deletedAt" gorm:"default:删除时间;comment:删除时间;"
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,107,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *UcAppleidInfoFilter) Reset() {
	*x = UcAppleidInfoFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uc_appleid_info_model_gen_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UcAppleidInfoFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UcAppleidInfoFilter) ProtoMessage() {}

func (x *UcAppleidInfoFilter) ProtoReflect() protoreflect.Message {
	mi := &file_uc_appleid_info_model_gen_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UcAppleidInfoFilter.ProtoReflect.Descriptor instead.
func (*UcAppleidInfoFilter) Descriptor() ([]byte, []int) {
	return file_uc_appleid_info_model_gen_proto_rawDescGZIP(), []int{1}
}

func (x *UcAppleidInfoFilter) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UcAppleidInfoFilter) GetThirdUserId() string {
	if x != nil {
		return x.ThirdUserId
	}
	return ""
}

func (x *UcAppleidInfoFilter) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *UcAppleidInfoFilter) GetCreatedBy() int64 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *UcAppleidInfoFilter) GetUpdatedBy() int64 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

func (x *UcAppleidInfoFilter) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UcAppleidInfoFilter) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *UcAppleidInfoFilter) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type UcAppleidInfoListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query     *UcAppleidInfo `protobuf:"bytes,101,opt,name=query,proto3" json:"query,omitempty"`
	Page      int64          `protobuf:"varint,102,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int64          `protobuf:"varint,103,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	OrderKey  string         `protobuf:"bytes,104,opt,name=orderKey,proto3" json:"orderKey,omitempty"`
	OrderDesc string         `protobuf:"bytes,105,opt,name=orderDesc,proto3" json:"orderDesc,omitempty"`
	SearchKey string         `protobuf:"bytes,106,opt,name=searchKey,proto3" json:"searchKey,omitempty"`
}

func (x *UcAppleidInfoListRequest) Reset() {
	*x = UcAppleidInfoListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uc_appleid_info_model_gen_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UcAppleidInfoListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UcAppleidInfoListRequest) ProtoMessage() {}

func (x *UcAppleidInfoListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_uc_appleid_info_model_gen_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UcAppleidInfoListRequest.ProtoReflect.Descriptor instead.
func (*UcAppleidInfoListRequest) Descriptor() ([]byte, []int) {
	return file_uc_appleid_info_model_gen_proto_rawDescGZIP(), []int{2}
}

func (x *UcAppleidInfoListRequest) GetQuery() *UcAppleidInfo {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *UcAppleidInfoListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *UcAppleidInfoListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *UcAppleidInfoListRequest) GetOrderKey() string {
	if x != nil {
		return x.OrderKey
	}
	return ""
}

func (x *UcAppleidInfoListRequest) GetOrderDesc() string {
	if x != nil {
		return x.OrderDesc
	}
	return ""
}

func (x *UcAppleidInfoListRequest) GetSearchKey() string {
	if x != nil {
		return x.SearchKey
	}
	return ""
}

type UcAppleidInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32            `protobuf:"varint,101,opt,name=code,proto3" json:"code,omitempty"`
	Message string           `protobuf:"bytes,102,opt,name=message,proto3" json:"message,omitempty"`
	Total   int64            `protobuf:"varint,103,opt,name=total,proto3" json:"total,omitempty"`
	Data    []*UcAppleidInfo `protobuf:"bytes,104,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *UcAppleidInfoResponse) Reset() {
	*x = UcAppleidInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uc_appleid_info_model_gen_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UcAppleidInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UcAppleidInfoResponse) ProtoMessage() {}

func (x *UcAppleidInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_uc_appleid_info_model_gen_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UcAppleidInfoResponse.ProtoReflect.Descriptor instead.
func (*UcAppleidInfoResponse) Descriptor() ([]byte, []int) {
	return file_uc_appleid_info_model_gen_proto_rawDescGZIP(), []int{3}
}

func (x *UcAppleidInfoResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *UcAppleidInfoResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UcAppleidInfoResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *UcAppleidInfoResponse) GetData() []*UcAppleidInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type UcAppleidInfoUpdateFieldsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []string       `protobuf:"bytes,101,rep,name=fields,proto3" json:"fields,omitempty"`
	Data   *UcAppleidInfo `protobuf:"bytes,102,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UcAppleidInfoUpdateFieldsRequest) Reset() {
	*x = UcAppleidInfoUpdateFieldsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uc_appleid_info_model_gen_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UcAppleidInfoUpdateFieldsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UcAppleidInfoUpdateFieldsRequest) ProtoMessage() {}

func (x *UcAppleidInfoUpdateFieldsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_uc_appleid_info_model_gen_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UcAppleidInfoUpdateFieldsRequest.ProtoReflect.Descriptor instead.
func (*UcAppleidInfoUpdateFieldsRequest) Descriptor() ([]byte, []int) {
	return file_uc_appleid_info_model_gen_proto_rawDescGZIP(), []int{4}
}

func (x *UcAppleidInfoUpdateFieldsRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *UcAppleidInfoUpdateFieldsRequest) GetData() *UcAppleidInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

// 表模型主键，只能是整数或字符串，若生成错误，请检查数据库表主键是否正确。支持组合主键
type UcAppleidInfoPrimarykey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UcAppleidInfoPrimarykey) Reset() {
	*x = UcAppleidInfoPrimarykey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uc_appleid_info_model_gen_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UcAppleidInfoPrimarykey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UcAppleidInfoPrimarykey) ProtoMessage() {}

func (x *UcAppleidInfoPrimarykey) ProtoReflect() protoreflect.Message {
	mi := &file_uc_appleid_info_model_gen_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UcAppleidInfoPrimarykey.ProtoReflect.Descriptor instead.
func (*UcAppleidInfoPrimarykey) Descriptor() ([]byte, []int) {
	return file_uc_appleid_info_model_gen_proto_rawDescGZIP(), []int{5}
}

func (x *UcAppleidInfoPrimarykey) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UcAppleidInfoBatchDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []*UcAppleidInfoPrimarykey `protobuf:"bytes,101,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *UcAppleidInfoBatchDeleteRequest) Reset() {
	*x = UcAppleidInfoBatchDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uc_appleid_info_model_gen_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UcAppleidInfoBatchDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UcAppleidInfoBatchDeleteRequest) ProtoMessage() {}

func (x *UcAppleidInfoBatchDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_uc_appleid_info_model_gen_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UcAppleidInfoBatchDeleteRequest.ProtoReflect.Descriptor instead.
func (*UcAppleidInfoBatchDeleteRequest) Descriptor() ([]byte, []int) {
	return file_uc_appleid_info_model_gen_proto_rawDescGZIP(), []int{6}
}

func (x *UcAppleidInfoBatchDeleteRequest) GetKeys() []*UcAppleidInfoPrimarykey {
	if x != nil {
		return x.Keys
	}
	return nil
}

var File_uc_appleid_info_model_gen_proto protoreflect.FileDescriptor

var file_uc_appleid_info_model_gen_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x75, 0x63, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x69, 0x64, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc7, 0x02, 0x0a, 0x0d,
	0x55, 0x63, 0x41, 0x70, 0x70, 0x6c, 0x65, 0x69, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x74, 0x68, 0x69, 0x72, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x65, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x74, 0x68, 0x69, 0x72, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x68, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
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
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xcd, 0x02, 0x0a, 0x13, 0x55, 0x63, 0x41, 0x70, 0x70, 0x6c,
	0x65, 0x69, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x74, 0x68, 0x69, 0x72, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x65, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x74, 0x68, 0x69, 0x72, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x68, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
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
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xd0, 0x01, 0x0a, 0x18, 0x55, 0x63, 0x41, 0x70, 0x70, 0x6c,
	0x65, 0x69, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2c, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x65, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x63, 0x41, 0x70,
	0x70, 0x6c, 0x65, 0x69, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x18, 0x68, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x18, 0x69, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x22, 0x87, 0x01, 0x0a, 0x15, 0x55, 0x63, 0x41,
	0x70, 0x70, 0x6c, 0x65, 0x69, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x68,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55,
	0x63, 0x41, 0x70, 0x70, 0x6c, 0x65, 0x69, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x66, 0x0a, 0x20, 0x55, 0x63, 0x41, 0x70, 0x70, 0x6c, 0x65, 0x69, 0x64, 0x49,
	0x6e, 0x66, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x18, 0x65, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x2a,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x63, 0x41, 0x70, 0x70, 0x6c, 0x65, 0x69, 0x64,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x29, 0x0a, 0x17, 0x55, 0x63,
	0x41, 0x70, 0x70, 0x6c, 0x65, 0x69, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x6b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x57, 0x0a, 0x1f, 0x55, 0x63, 0x41, 0x70, 0x70, 0x6c, 0x65,
	0x69, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73,
	0x18, 0x65, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x55, 0x63, 0x41, 0x70, 0x70, 0x6c, 0x65, 0x69, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x42, 0x11,
	0x5a, 0x0f, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_uc_appleid_info_model_gen_proto_rawDescOnce sync.Once
	file_uc_appleid_info_model_gen_proto_rawDescData = file_uc_appleid_info_model_gen_proto_rawDesc
)

func file_uc_appleid_info_model_gen_proto_rawDescGZIP() []byte {
	file_uc_appleid_info_model_gen_proto_rawDescOnce.Do(func() {
		file_uc_appleid_info_model_gen_proto_rawDescData = protoimpl.X.CompressGZIP(file_uc_appleid_info_model_gen_proto_rawDescData)
	})
	return file_uc_appleid_info_model_gen_proto_rawDescData
}

var file_uc_appleid_info_model_gen_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_uc_appleid_info_model_gen_proto_goTypes = []interface{}{
	(*UcAppleidInfo)(nil),                    // 0: service.UcAppleidInfo
	(*UcAppleidInfoFilter)(nil),              // 1: service.UcAppleidInfoFilter
	(*UcAppleidInfoListRequest)(nil),         // 2: service.UcAppleidInfoListRequest
	(*UcAppleidInfoResponse)(nil),            // 3: service.UcAppleidInfoResponse
	(*UcAppleidInfoUpdateFieldsRequest)(nil), // 4: service.UcAppleidInfoUpdateFieldsRequest
	(*UcAppleidInfoPrimarykey)(nil),          // 5: service.UcAppleidInfoPrimarykey
	(*UcAppleidInfoBatchDeleteRequest)(nil),  // 6: service.UcAppleidInfoBatchDeleteRequest
	(*timestamppb.Timestamp)(nil),            // 7: google.protobuf.Timestamp
}
var file_uc_appleid_info_model_gen_proto_depIdxs = []int32{
	7,  // 0: service.UcAppleidInfo.createdAt:type_name -> google.protobuf.Timestamp
	7,  // 1: service.UcAppleidInfo.updatedAt:type_name -> google.protobuf.Timestamp
	7,  // 2: service.UcAppleidInfo.deletedAt:type_name -> google.protobuf.Timestamp
	7,  // 3: service.UcAppleidInfoFilter.createdAt:type_name -> google.protobuf.Timestamp
	7,  // 4: service.UcAppleidInfoFilter.updatedAt:type_name -> google.protobuf.Timestamp
	7,  // 5: service.UcAppleidInfoFilter.deletedAt:type_name -> google.protobuf.Timestamp
	0,  // 6: service.UcAppleidInfoListRequest.query:type_name -> service.UcAppleidInfo
	0,  // 7: service.UcAppleidInfoResponse.data:type_name -> service.UcAppleidInfo
	0,  // 8: service.UcAppleidInfoUpdateFieldsRequest.data:type_name -> service.UcAppleidInfo
	5,  // 9: service.UcAppleidInfoBatchDeleteRequest.keys:type_name -> service.UcAppleidInfoPrimarykey
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_uc_appleid_info_model_gen_proto_init() }
func file_uc_appleid_info_model_gen_proto_init() {
	if File_uc_appleid_info_model_gen_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_uc_appleid_info_model_gen_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UcAppleidInfo); i {
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
		file_uc_appleid_info_model_gen_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UcAppleidInfoFilter); i {
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
		file_uc_appleid_info_model_gen_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UcAppleidInfoListRequest); i {
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
		file_uc_appleid_info_model_gen_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UcAppleidInfoResponse); i {
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
		file_uc_appleid_info_model_gen_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UcAppleidInfoUpdateFieldsRequest); i {
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
		file_uc_appleid_info_model_gen_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UcAppleidInfoPrimarykey); i {
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
		file_uc_appleid_info_model_gen_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UcAppleidInfoBatchDeleteRequest); i {
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
			RawDescriptor: file_uc_appleid_info_model_gen_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_uc_appleid_info_model_gen_proto_goTypes,
		DependencyIndexes: file_uc_appleid_info_model_gen_proto_depIdxs,
		MessageInfos:      file_uc_appleid_info_model_gen_proto_msgTypes,
	}.Build()
	File_uc_appleid_info_model_gen_proto = out.File
	file_uc_appleid_info_model_gen_proto_rawDesc = nil
	file_uc_appleid_info_model_gen_proto_goTypes = nil
	file_uc_appleid_info_model_gen_proto_depIdxs = nil
}
