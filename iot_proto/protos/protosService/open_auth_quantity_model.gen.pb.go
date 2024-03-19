// Code generated by protoc,2022-05-12 15:44:43. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: open_auth_quantity_model.gen.proto

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
type OpenAuthQuantity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"primary_key;AUTO_INCREMENT;default:主键;comment:主键;size:20;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"companyId" gorm:"default:公司ID;comment:公司ID;size:19;"
	CompanyId int64 `protobuf:"varint,101,opt,name=companyId,proto3" json:"companyId,omitempty"`
	// @inject_tag: json:"tenantId" gorm:"default:租户ID;comment:租户ID;size:6;"
	TenantId string `protobuf:"bytes,102,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	// @inject_tag: json:"authCode" gorm:"default:授权码,自动生成;comment:授权码,自动生成;size:8;"
	AuthCode string `protobuf:"bytes,103,opt,name=authCode,proto3" json:"authCode,omitempty"`
	// @inject_tag: json:"authQuantity" gorm:"default:授权数量;comment:授权数量;size:10;"
	AuthQuantity int32 `protobuf:"varint,104,opt,name=authQuantity,proto3" json:"authQuantity,omitempty"`
	// @inject_tag: json:"authDate" gorm:"default:授权日期;comment:授权日期;"
	AuthDate *timestamppb.Timestamp `protobuf:"bytes,105,opt,name=authDate,proto3" json:"authDate,omitempty"`
	// @inject_tag: json:"status" gorm:"default:状态;1: 正常; 2:禁用;;comment:状态;1: 正常; 2:禁用;;"
	Status int32 `protobuf:"varint,106,opt,name=status,proto3" json:"status,omitempty"`
	// @inject_tag: json:"createdBy" gorm:"default:创建者;comment:创建者;size:19;"
	CreatedBy int64 `protobuf:"varint,107,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	// @inject_tag: json:"createdAt" gorm:"default:创建时间;comment:创建时间;"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,108,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// @inject_tag: json:"updatedBy" gorm:"default:更新者;comment:更新者;size:19;"
	UpdatedBy int64 `protobuf:"varint,109,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	// @inject_tag: json:"updatedAt" gorm:"default:更新时间;comment:更新时间;"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,110,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// @inject_tag: json:"deletedAt" gorm:"default:删除时间,有值表示已删除;comment:删除时间,有值表示已删除;"
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,111,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
	// @inject_tag: json:"createdBy" gorm:"default:创建者;comment:创建者;size:19;"
	UserId int64 `protobuf:"varint,112,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *OpenAuthQuantity) Reset() {
	*x = OpenAuthQuantity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_open_auth_quantity_model_gen_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenAuthQuantity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenAuthQuantity) ProtoMessage() {}

func (x *OpenAuthQuantity) ProtoReflect() protoreflect.Message {
	mi := &file_open_auth_quantity_model_gen_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenAuthQuantity.ProtoReflect.Descriptor instead.
func (*OpenAuthQuantity) Descriptor() ([]byte, []int) {
	return file_open_auth_quantity_model_gen_proto_rawDescGZIP(), []int{0}
}

func (x *OpenAuthQuantity) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OpenAuthQuantity) GetCompanyId() int64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *OpenAuthQuantity) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *OpenAuthQuantity) GetAuthCode() string {
	if x != nil {
		return x.AuthCode
	}
	return ""
}

func (x *OpenAuthQuantity) GetAuthQuantity() int32 {
	if x != nil {
		return x.AuthQuantity
	}
	return 0
}

func (x *OpenAuthQuantity) GetAuthDate() *timestamppb.Timestamp {
	if x != nil {
		return x.AuthDate
	}
	return nil
}

func (x *OpenAuthQuantity) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *OpenAuthQuantity) GetCreatedBy() int64 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *OpenAuthQuantity) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *OpenAuthQuantity) GetUpdatedBy() int64 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

func (x *OpenAuthQuantity) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *OpenAuthQuantity) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *OpenAuthQuantity) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type OpenAuthQuantityFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"primary_key;AUTO_INCREMENT;default:主键;comment:主键;size:20;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"companyId" gorm:"default:公司ID;comment:公司ID;size:19;"
	CompanyId int64 `protobuf:"varint,101,opt,name=companyId,proto3" json:"companyId,omitempty"`
	// @inject_tag: json:"tenantId" gorm:"default:租户ID;comment:租户ID;size:6;"
	TenantId string `protobuf:"bytes,102,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	// @inject_tag: json:"authCode" gorm:"default:授权码,自动生成;comment:授权码,自动生成;size:8;"
	AuthCode string `protobuf:"bytes,103,opt,name=authCode,proto3" json:"authCode,omitempty"`
	// @inject_tag: json:"authQuantity" gorm:"default:授权数量;comment:授权数量;size:10;"
	AuthQuantity int32 `protobuf:"varint,104,opt,name=authQuantity,proto3" json:"authQuantity,omitempty"`
	// @inject_tag: json:"authDate" gorm:"default:授权日期;comment:授权日期;"
	AuthDate *timestamppb.Timestamp `protobuf:"bytes,105,opt,name=authDate,proto3" json:"authDate,omitempty"`
	// @inject_tag: json:"status" gorm:"default:状态;1: 正常; 2:禁用;;comment:状态;1: 正常; 2:禁用;;"
	Status int32 `protobuf:"varint,106,opt,name=status,proto3" json:"status,omitempty"`
	// @inject_tag: json:"createdBy" gorm:"default:创建者;comment:创建者;size:19;"
	CreatedBy int64 `protobuf:"varint,107,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	// @inject_tag: json:"createdAt" gorm:"default:创建时间;comment:创建时间;"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,108,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// @inject_tag: json:"updatedBy" gorm:"default:更新者;comment:更新者;size:19;"
	UpdatedBy int64 `protobuf:"varint,109,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	// @inject_tag: json:"updatedAt" gorm:"default:更新时间;comment:更新时间;"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,110,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// @inject_tag: json:"deletedAt" gorm:"default:删除时间,有值表示已删除;comment:删除时间,有值表示已删除;"
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,111,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
	// @inject_tag: json:"createdBy" gorm:"default:创建者;comment:创建者;size:19;"
	UserId int64 `protobuf:"varint,112,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *OpenAuthQuantityFilter) Reset() {
	*x = OpenAuthQuantityFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_open_auth_quantity_model_gen_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenAuthQuantityFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenAuthQuantityFilter) ProtoMessage() {}

func (x *OpenAuthQuantityFilter) ProtoReflect() protoreflect.Message {
	mi := &file_open_auth_quantity_model_gen_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenAuthQuantityFilter.ProtoReflect.Descriptor instead.
func (*OpenAuthQuantityFilter) Descriptor() ([]byte, []int) {
	return file_open_auth_quantity_model_gen_proto_rawDescGZIP(), []int{1}
}

func (x *OpenAuthQuantityFilter) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OpenAuthQuantityFilter) GetCompanyId() int64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *OpenAuthQuantityFilter) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *OpenAuthQuantityFilter) GetAuthCode() string {
	if x != nil {
		return x.AuthCode
	}
	return ""
}

func (x *OpenAuthQuantityFilter) GetAuthQuantity() int32 {
	if x != nil {
		return x.AuthQuantity
	}
	return 0
}

func (x *OpenAuthQuantityFilter) GetAuthDate() *timestamppb.Timestamp {
	if x != nil {
		return x.AuthDate
	}
	return nil
}

func (x *OpenAuthQuantityFilter) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *OpenAuthQuantityFilter) GetCreatedBy() int64 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *OpenAuthQuantityFilter) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *OpenAuthQuantityFilter) GetUpdatedBy() int64 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

func (x *OpenAuthQuantityFilter) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *OpenAuthQuantityFilter) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *OpenAuthQuantityFilter) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type OpenAuthQuantityListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query     *OpenAuthQuantity `protobuf:"bytes,101,opt,name=query,proto3" json:"query,omitempty"`
	Page      int64             `protobuf:"varint,102,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int64             `protobuf:"varint,103,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	OrderKey  string            `protobuf:"bytes,104,opt,name=orderKey,proto3" json:"orderKey,omitempty"`
	OrderDesc string            `protobuf:"bytes,105,opt,name=orderDesc,proto3" json:"orderDesc,omitempty"`
	SearchKey string            `protobuf:"bytes,106,opt,name=searchKey,proto3" json:"searchKey,omitempty"`
}

func (x *OpenAuthQuantityListRequest) Reset() {
	*x = OpenAuthQuantityListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_open_auth_quantity_model_gen_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenAuthQuantityListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenAuthQuantityListRequest) ProtoMessage() {}

func (x *OpenAuthQuantityListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_open_auth_quantity_model_gen_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenAuthQuantityListRequest.ProtoReflect.Descriptor instead.
func (*OpenAuthQuantityListRequest) Descriptor() ([]byte, []int) {
	return file_open_auth_quantity_model_gen_proto_rawDescGZIP(), []int{2}
}

func (x *OpenAuthQuantityListRequest) GetQuery() *OpenAuthQuantity {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *OpenAuthQuantityListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *OpenAuthQuantityListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *OpenAuthQuantityListRequest) GetOrderKey() string {
	if x != nil {
		return x.OrderKey
	}
	return ""
}

func (x *OpenAuthQuantityListRequest) GetOrderDesc() string {
	if x != nil {
		return x.OrderDesc
	}
	return ""
}

func (x *OpenAuthQuantityListRequest) GetSearchKey() string {
	if x != nil {
		return x.SearchKey
	}
	return ""
}

type OpenAuthQuantityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32               `protobuf:"varint,101,opt,name=code,proto3" json:"code,omitempty"`
	Message string              `protobuf:"bytes,102,opt,name=message,proto3" json:"message,omitempty"`
	Total   int64               `protobuf:"varint,103,opt,name=total,proto3" json:"total,omitempty"`
	Data    []*OpenAuthQuantity `protobuf:"bytes,104,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *OpenAuthQuantityResponse) Reset() {
	*x = OpenAuthQuantityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_open_auth_quantity_model_gen_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenAuthQuantityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenAuthQuantityResponse) ProtoMessage() {}

func (x *OpenAuthQuantityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_open_auth_quantity_model_gen_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenAuthQuantityResponse.ProtoReflect.Descriptor instead.
func (*OpenAuthQuantityResponse) Descriptor() ([]byte, []int) {
	return file_open_auth_quantity_model_gen_proto_rawDescGZIP(), []int{3}
}

func (x *OpenAuthQuantityResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *OpenAuthQuantityResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *OpenAuthQuantityResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *OpenAuthQuantityResponse) GetData() []*OpenAuthQuantity {
	if x != nil {
		return x.Data
	}
	return nil
}

type OpenAuthQuantityUpdateFieldsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []string          `protobuf:"bytes,101,rep,name=fields,proto3" json:"fields,omitempty"`
	Data   *OpenAuthQuantity `protobuf:"bytes,102,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OpenAuthQuantityUpdateFieldsRequest) Reset() {
	*x = OpenAuthQuantityUpdateFieldsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_open_auth_quantity_model_gen_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenAuthQuantityUpdateFieldsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenAuthQuantityUpdateFieldsRequest) ProtoMessage() {}

func (x *OpenAuthQuantityUpdateFieldsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_open_auth_quantity_model_gen_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenAuthQuantityUpdateFieldsRequest.ProtoReflect.Descriptor instead.
func (*OpenAuthQuantityUpdateFieldsRequest) Descriptor() ([]byte, []int) {
	return file_open_auth_quantity_model_gen_proto_rawDescGZIP(), []int{4}
}

func (x *OpenAuthQuantityUpdateFieldsRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *OpenAuthQuantityUpdateFieldsRequest) GetData() *OpenAuthQuantity {
	if x != nil {
		return x.Data
	}
	return nil
}

// 表模型主键，只能是整数或字符串，若生成错误，请检查数据库表主键是否正确。支持组合主键
type OpenAuthQuantityPrimarykey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OpenAuthQuantityPrimarykey) Reset() {
	*x = OpenAuthQuantityPrimarykey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_open_auth_quantity_model_gen_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenAuthQuantityPrimarykey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenAuthQuantityPrimarykey) ProtoMessage() {}

func (x *OpenAuthQuantityPrimarykey) ProtoReflect() protoreflect.Message {
	mi := &file_open_auth_quantity_model_gen_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenAuthQuantityPrimarykey.ProtoReflect.Descriptor instead.
func (*OpenAuthQuantityPrimarykey) Descriptor() ([]byte, []int) {
	return file_open_auth_quantity_model_gen_proto_rawDescGZIP(), []int{5}
}

func (x *OpenAuthQuantityPrimarykey) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type OpenAuthQuantityBatchDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []*OpenAuthQuantityPrimarykey `protobuf:"bytes,101,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *OpenAuthQuantityBatchDeleteRequest) Reset() {
	*x = OpenAuthQuantityBatchDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_open_auth_quantity_model_gen_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenAuthQuantityBatchDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenAuthQuantityBatchDeleteRequest) ProtoMessage() {}

func (x *OpenAuthQuantityBatchDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_open_auth_quantity_model_gen_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenAuthQuantityBatchDeleteRequest.ProtoReflect.Descriptor instead.
func (*OpenAuthQuantityBatchDeleteRequest) Descriptor() ([]byte, []int) {
	return file_open_auth_quantity_model_gen_proto_rawDescGZIP(), []int{6}
}

func (x *OpenAuthQuantityBatchDeleteRequest) GetKeys() []*OpenAuthQuantityPrimarykey {
	if x != nil {
		return x.Keys
	}
	return nil
}

var File_open_auth_quantity_model_gen_proto protoreflect.FileDescriptor

var file_open_auth_quantity_model_gen_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee,
	0x03, 0x0a, 0x10, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x51, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64,
	0x18, 0x65, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x66, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x61, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x61, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x75, 0x74,
	0x68, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x68, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x61, 0x75, 0x74, 0x68, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x36, 0x0a,
	0x08, 0x61, 0x75, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x18, 0x69, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x61, 0x75, 0x74,
	0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x6a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x6b, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x6c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x18, 0x6d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x6e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a,
	0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x6f, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x70, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0xf4, 0x03, 0x0a, 0x16, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x51, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x18, 0x65, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x49, 0x64, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x67, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x68, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x51, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x36, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65,
	0x18, 0x69, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x18, 0x6b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x6c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x6d, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x6f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x70, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xd6, 0x01, 0x0a, 0x1b, 0x4f, 0x70, 0x65, 0x6e, 0x41,
	0x75, 0x74, 0x68, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x4f, 0x70, 0x65, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x66, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x4b, 0x65, 0x79, 0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63,
	0x18, 0x69, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x73,
	0x63, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x18, 0x6a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x22,
	0x8d, 0x01, 0x0a, 0x18, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x51, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x68, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x75, 0x74,
	0x68, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x6c, 0x0a, 0x23, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x51, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x18, 0x65, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x2d,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x51,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2c, 0x0a,
	0x1a, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5d, 0x0a, 0x22, 0x4f,
	0x70, 0x65, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x37, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x75,
	0x74, 0x68, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x6b, 0x65, 0x79, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_open_auth_quantity_model_gen_proto_rawDescOnce sync.Once
	file_open_auth_quantity_model_gen_proto_rawDescData = file_open_auth_quantity_model_gen_proto_rawDesc
)

func file_open_auth_quantity_model_gen_proto_rawDescGZIP() []byte {
	file_open_auth_quantity_model_gen_proto_rawDescOnce.Do(func() {
		file_open_auth_quantity_model_gen_proto_rawDescData = protoimpl.X.CompressGZIP(file_open_auth_quantity_model_gen_proto_rawDescData)
	})
	return file_open_auth_quantity_model_gen_proto_rawDescData
}

var file_open_auth_quantity_model_gen_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_open_auth_quantity_model_gen_proto_goTypes = []interface{}{
	(*OpenAuthQuantity)(nil),                    // 0: service.OpenAuthQuantity
	(*OpenAuthQuantityFilter)(nil),              // 1: service.OpenAuthQuantityFilter
	(*OpenAuthQuantityListRequest)(nil),         // 2: service.OpenAuthQuantityListRequest
	(*OpenAuthQuantityResponse)(nil),            // 3: service.OpenAuthQuantityResponse
	(*OpenAuthQuantityUpdateFieldsRequest)(nil), // 4: service.OpenAuthQuantityUpdateFieldsRequest
	(*OpenAuthQuantityPrimarykey)(nil),          // 5: service.OpenAuthQuantityPrimarykey
	(*OpenAuthQuantityBatchDeleteRequest)(nil),  // 6: service.OpenAuthQuantityBatchDeleteRequest
	(*timestamppb.Timestamp)(nil),               // 7: google.protobuf.Timestamp
}
var file_open_auth_quantity_model_gen_proto_depIdxs = []int32{
	7,  // 0: service.OpenAuthQuantity.authDate:type_name -> google.protobuf.Timestamp
	7,  // 1: service.OpenAuthQuantity.createdAt:type_name -> google.protobuf.Timestamp
	7,  // 2: service.OpenAuthQuantity.updatedAt:type_name -> google.protobuf.Timestamp
	7,  // 3: service.OpenAuthQuantity.deletedAt:type_name -> google.protobuf.Timestamp
	7,  // 4: service.OpenAuthQuantityFilter.authDate:type_name -> google.protobuf.Timestamp
	7,  // 5: service.OpenAuthQuantityFilter.createdAt:type_name -> google.protobuf.Timestamp
	7,  // 6: service.OpenAuthQuantityFilter.updatedAt:type_name -> google.protobuf.Timestamp
	7,  // 7: service.OpenAuthQuantityFilter.deletedAt:type_name -> google.protobuf.Timestamp
	0,  // 8: service.OpenAuthQuantityListRequest.query:type_name -> service.OpenAuthQuantity
	0,  // 9: service.OpenAuthQuantityResponse.data:type_name -> service.OpenAuthQuantity
	0,  // 10: service.OpenAuthQuantityUpdateFieldsRequest.data:type_name -> service.OpenAuthQuantity
	5,  // 11: service.OpenAuthQuantityBatchDeleteRequest.keys:type_name -> service.OpenAuthQuantityPrimarykey
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_open_auth_quantity_model_gen_proto_init() }
func file_open_auth_quantity_model_gen_proto_init() {
	if File_open_auth_quantity_model_gen_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_open_auth_quantity_model_gen_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenAuthQuantity); i {
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
		file_open_auth_quantity_model_gen_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenAuthQuantityFilter); i {
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
		file_open_auth_quantity_model_gen_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenAuthQuantityListRequest); i {
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
		file_open_auth_quantity_model_gen_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenAuthQuantityResponse); i {
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
		file_open_auth_quantity_model_gen_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenAuthQuantityUpdateFieldsRequest); i {
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
		file_open_auth_quantity_model_gen_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenAuthQuantityPrimarykey); i {
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
		file_open_auth_quantity_model_gen_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenAuthQuantityBatchDeleteRequest); i {
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
			RawDescriptor: file_open_auth_quantity_model_gen_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_open_auth_quantity_model_gen_proto_goTypes,
		DependencyIndexes: file_open_auth_quantity_model_gen_proto_depIdxs,
		MessageInfos:      file_open_auth_quantity_model_gen_proto_msgTypes,
	}.Build()
	File_open_auth_quantity_model_gen_proto = out.File
	file_open_auth_quantity_model_gen_proto_rawDesc = nil
	file_open_auth_quantity_model_gen_proto_goTypes = nil
	file_open_auth_quantity_model_gen_proto_depIdxs = nil
}
