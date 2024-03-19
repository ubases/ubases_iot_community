// Code generated by protoc,2022-04-19 09:58:34. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: config_oss_model.gen.proto

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
type ConfigOss struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:主键ID;comment:主键ID;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"regionId" gorm:"default:区域ID（t_config_area.id);comment:区域ID（t_config_area.id);size:19;"
	RegionId int64 `protobuf:"varint,101,opt,name=regionId,proto3" json:"regionId,omitempty"`
	// @inject_tag: json:"accessKeyId" gorm:"default:oss配置access_key_id;comment:oss配置access_key_id;size:50;"
	AccessKeyId string `protobuf:"bytes,102,opt,name=accessKeyId,proto3" json:"accessKeyId,omitempty"`
	// @inject_tag: json:"accessKeySecret" gorm:"default:oss配置AccessKeySecret;comment:oss配置AccessKeySecret;size:100;"
	AccessKeySecret string `protobuf:"bytes,103,opt,name=accessKeySecret,proto3" json:"accessKeySecret,omitempty"`
	// @inject_tag: json:"roleArn" gorm:"default:角色Arn;comment:角色Arn;size:50;"
	RoleArn string `protobuf:"bytes,104,opt,name=roleArn,proto3" json:"roleArn,omitempty"`
	// @inject_tag: json:"roleSessionName" gorm:"default:角色名称;comment:角色名称;size:50;"
	RoleSessionName string `protobuf:"bytes,105,opt,name=roleSessionName,proto3" json:"roleSessionName,omitempty"`
	// @inject_tag: json:"bucket" gorm:"default:Bucket名称;comment:Bucket名称;size:20;"
	Bucket string `protobuf:"bytes,106,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// @inject_tag: json:"endpoint" gorm:"default:地域节点;comment:地域节点;size:20;"
	Endpoint string `protobuf:"bytes,107,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// @inject_tag: json:"ossName" gorm:"default:OSS名称;comment:OSS名称;size:20;"
	OssName string `protobuf:"bytes,108,opt,name=ossName,proto3" json:"ossName,omitempty"`
	// @inject_tag: json:"remark" gorm:"default:OSS描述;comment:OSS描述;size:255;"
	Remark string `protobuf:"bytes,109,opt,name=remark,proto3" json:"remark,omitempty"`
	// @inject_tag: json:"status" gorm:"default:启用状态 =1启用=0禁用;comment:启用状态 =1启用=0禁用;size:10;"
	Status int32 `protobuf:"varint,110,opt,name=status,proto3" json:"status,omitempty"`
	// @inject_tag: json:"createdBy" gorm:"default:创建人;comment:创建人;size:20;"
	CreatedBy int64 `protobuf:"varint,111,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	// @inject_tag: json:"updatedBy" gorm:"default:修改人;comment:修改人;size:20;"
	UpdatedBy int64 `protobuf:"varint,112,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	// @inject_tag: json:"createdAt" gorm:"default:创建时间;comment:创建时间;"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,113,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// @inject_tag: json:"updatedAt" gorm:"default:修改时间;comment:修改时间;"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,114,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// @inject_tag: json:"deletedAt" gorm:"default:删除时间;comment:删除时间;"
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,115,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *ConfigOss) Reset() {
	*x = ConfigOss{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_oss_model_gen_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigOss) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigOss) ProtoMessage() {}

func (x *ConfigOss) ProtoReflect() protoreflect.Message {
	mi := &file_config_oss_model_gen_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigOss.ProtoReflect.Descriptor instead.
func (*ConfigOss) Descriptor() ([]byte, []int) {
	return file_config_oss_model_gen_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigOss) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ConfigOss) GetRegionId() int64 {
	if x != nil {
		return x.RegionId
	}
	return 0
}

func (x *ConfigOss) GetAccessKeyId() string {
	if x != nil {
		return x.AccessKeyId
	}
	return ""
}

func (x *ConfigOss) GetAccessKeySecret() string {
	if x != nil {
		return x.AccessKeySecret
	}
	return ""
}

func (x *ConfigOss) GetRoleArn() string {
	if x != nil {
		return x.RoleArn
	}
	return ""
}

func (x *ConfigOss) GetRoleSessionName() string {
	if x != nil {
		return x.RoleSessionName
	}
	return ""
}

func (x *ConfigOss) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *ConfigOss) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *ConfigOss) GetOssName() string {
	if x != nil {
		return x.OssName
	}
	return ""
}

func (x *ConfigOss) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *ConfigOss) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ConfigOss) GetCreatedBy() int64 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *ConfigOss) GetUpdatedBy() int64 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

func (x *ConfigOss) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ConfigOss) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *ConfigOss) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type ConfigOssFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:主键ID;comment:主键ID;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"regionId" gorm:"default:区域ID（t_config_area.id);comment:区域ID（t_config_area.id);size:19;"
	RegionId int64 `protobuf:"varint,101,opt,name=regionId,proto3" json:"regionId,omitempty"`
	// @inject_tag: json:"accessKeyId" gorm:"default:oss配置access_key_id;comment:oss配置access_key_id;size:50;"
	AccessKeyId string `protobuf:"bytes,102,opt,name=accessKeyId,proto3" json:"accessKeyId,omitempty"`
	// @inject_tag: json:"accessKeySecret" gorm:"default:oss配置AccessKeySecret;comment:oss配置AccessKeySecret;size:100;"
	AccessKeySecret string `protobuf:"bytes,103,opt,name=accessKeySecret,proto3" json:"accessKeySecret,omitempty"`
	// @inject_tag: json:"roleArn" gorm:"default:角色Arn;comment:角色Arn;size:50;"
	RoleArn string `protobuf:"bytes,104,opt,name=roleArn,proto3" json:"roleArn,omitempty"`
	// @inject_tag: json:"roleSessionName" gorm:"default:角色名称;comment:角色名称;size:50;"
	RoleSessionName string `protobuf:"bytes,105,opt,name=roleSessionName,proto3" json:"roleSessionName,omitempty"`
	// @inject_tag: json:"bucket" gorm:"default:Bucket名称;comment:Bucket名称;size:20;"
	Bucket string `protobuf:"bytes,106,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// @inject_tag: json:"endpoint" gorm:"default:地域节点;comment:地域节点;size:20;"
	Endpoint string `protobuf:"bytes,107,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// @inject_tag: json:"ossName" gorm:"default:OSS名称;comment:OSS名称;size:20;"
	OssName string `protobuf:"bytes,108,opt,name=ossName,proto3" json:"ossName,omitempty"`
	// @inject_tag: json:"remark" gorm:"default:OSS描述;comment:OSS描述;size:255;"
	Remark string `protobuf:"bytes,109,opt,name=remark,proto3" json:"remark,omitempty"`
	// @inject_tag: json:"status" gorm:"default:启用状态 =1启用=0禁用;comment:启用状态 =1启用=0禁用;size:10;"
	Status int32 `protobuf:"varint,110,opt,name=status,proto3" json:"status,omitempty"`
	// @inject_tag: json:"createdBy" gorm:"default:创建人;comment:创建人;size:20;"
	CreatedBy int64 `protobuf:"varint,111,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	// @inject_tag: json:"updatedBy" gorm:"default:修改人;comment:修改人;size:20;"
	UpdatedBy int64 `protobuf:"varint,112,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	// @inject_tag: json:"createdAt" gorm:"default:创建时间;comment:创建时间;"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,113,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// @inject_tag: json:"updatedAt" gorm:"default:修改时间;comment:修改时间;"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,114,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// @inject_tag: json:"deletedAt" gorm:"default:删除时间;comment:删除时间;"
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,115,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *ConfigOssFilter) Reset() {
	*x = ConfigOssFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_oss_model_gen_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigOssFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigOssFilter) ProtoMessage() {}

func (x *ConfigOssFilter) ProtoReflect() protoreflect.Message {
	mi := &file_config_oss_model_gen_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigOssFilter.ProtoReflect.Descriptor instead.
func (*ConfigOssFilter) Descriptor() ([]byte, []int) {
	return file_config_oss_model_gen_proto_rawDescGZIP(), []int{1}
}

func (x *ConfigOssFilter) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ConfigOssFilter) GetRegionId() int64 {
	if x != nil {
		return x.RegionId
	}
	return 0
}

func (x *ConfigOssFilter) GetAccessKeyId() string {
	if x != nil {
		return x.AccessKeyId
	}
	return ""
}

func (x *ConfigOssFilter) GetAccessKeySecret() string {
	if x != nil {
		return x.AccessKeySecret
	}
	return ""
}

func (x *ConfigOssFilter) GetRoleArn() string {
	if x != nil {
		return x.RoleArn
	}
	return ""
}

func (x *ConfigOssFilter) GetRoleSessionName() string {
	if x != nil {
		return x.RoleSessionName
	}
	return ""
}

func (x *ConfigOssFilter) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *ConfigOssFilter) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *ConfigOssFilter) GetOssName() string {
	if x != nil {
		return x.OssName
	}
	return ""
}

func (x *ConfigOssFilter) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *ConfigOssFilter) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ConfigOssFilter) GetCreatedBy() int64 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *ConfigOssFilter) GetUpdatedBy() int64 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

func (x *ConfigOssFilter) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ConfigOssFilter) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *ConfigOssFilter) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type ConfigOssListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query     *ConfigOss `protobuf:"bytes,101,opt,name=query,proto3" json:"query,omitempty"`
	Page      int64      `protobuf:"varint,102,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int64      `protobuf:"varint,103,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	OrderKey  string     `protobuf:"bytes,104,opt,name=orderKey,proto3" json:"orderKey,omitempty"`
	OrderDesc string     `protobuf:"bytes,105,opt,name=orderDesc,proto3" json:"orderDesc,omitempty"`
}

func (x *ConfigOssListRequest) Reset() {
	*x = ConfigOssListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_oss_model_gen_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigOssListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigOssListRequest) ProtoMessage() {}

func (x *ConfigOssListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_oss_model_gen_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigOssListRequest.ProtoReflect.Descriptor instead.
func (*ConfigOssListRequest) Descriptor() ([]byte, []int) {
	return file_config_oss_model_gen_proto_rawDescGZIP(), []int{2}
}

func (x *ConfigOssListRequest) GetQuery() *ConfigOss {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *ConfigOssListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ConfigOssListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ConfigOssListRequest) GetOrderKey() string {
	if x != nil {
		return x.OrderKey
	}
	return ""
}

func (x *ConfigOssListRequest) GetOrderDesc() string {
	if x != nil {
		return x.OrderDesc
	}
	return ""
}

type ConfigOssResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32        `protobuf:"varint,101,opt,name=code,proto3" json:"code,omitempty"`
	Message string       `protobuf:"bytes,102,opt,name=message,proto3" json:"message,omitempty"`
	Total   int64        `protobuf:"varint,103,opt,name=total,proto3" json:"total,omitempty"`
	Data    []*ConfigOss `protobuf:"bytes,104,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ConfigOssResponse) Reset() {
	*x = ConfigOssResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_oss_model_gen_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigOssResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigOssResponse) ProtoMessage() {}

func (x *ConfigOssResponse) ProtoReflect() protoreflect.Message {
	mi := &file_config_oss_model_gen_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigOssResponse.ProtoReflect.Descriptor instead.
func (*ConfigOssResponse) Descriptor() ([]byte, []int) {
	return file_config_oss_model_gen_proto_rawDescGZIP(), []int{3}
}

func (x *ConfigOssResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ConfigOssResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ConfigOssResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ConfigOssResponse) GetData() []*ConfigOss {
	if x != nil {
		return x.Data
	}
	return nil
}

type ConfigOssUpdateFieldsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []string   `protobuf:"bytes,101,rep,name=fields,proto3" json:"fields,omitempty"`
	Data   *ConfigOss `protobuf:"bytes,102,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ConfigOssUpdateFieldsRequest) Reset() {
	*x = ConfigOssUpdateFieldsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_oss_model_gen_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigOssUpdateFieldsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigOssUpdateFieldsRequest) ProtoMessage() {}

func (x *ConfigOssUpdateFieldsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_oss_model_gen_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigOssUpdateFieldsRequest.ProtoReflect.Descriptor instead.
func (*ConfigOssUpdateFieldsRequest) Descriptor() ([]byte, []int) {
	return file_config_oss_model_gen_proto_rawDescGZIP(), []int{4}
}

func (x *ConfigOssUpdateFieldsRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *ConfigOssUpdateFieldsRequest) GetData() *ConfigOss {
	if x != nil {
		return x.Data
	}
	return nil
}

// 表模型主键，只能是整数或字符串，若生成错误，请检查数据库表主键是否正确。支持组合主键
type ConfigOssPrimarykey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ConfigOssPrimarykey) Reset() {
	*x = ConfigOssPrimarykey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_oss_model_gen_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigOssPrimarykey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigOssPrimarykey) ProtoMessage() {}

func (x *ConfigOssPrimarykey) ProtoReflect() protoreflect.Message {
	mi := &file_config_oss_model_gen_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigOssPrimarykey.ProtoReflect.Descriptor instead.
func (*ConfigOssPrimarykey) Descriptor() ([]byte, []int) {
	return file_config_oss_model_gen_proto_rawDescGZIP(), []int{5}
}

func (x *ConfigOssPrimarykey) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ConfigOssBatchDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []*ConfigOssPrimarykey `protobuf:"bytes,101,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *ConfigOssBatchDeleteRequest) Reset() {
	*x = ConfigOssBatchDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_oss_model_gen_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigOssBatchDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigOssBatchDeleteRequest) ProtoMessage() {}

func (x *ConfigOssBatchDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_oss_model_gen_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigOssBatchDeleteRequest.ProtoReflect.Descriptor instead.
func (*ConfigOssBatchDeleteRequest) Descriptor() ([]byte, []int) {
	return file_config_oss_model_gen_proto_rawDescGZIP(), []int{6}
}

func (x *ConfigOssBatchDeleteRequest) GetKeys() []*ConfigOssPrimarykey {
	if x != nil {
		return x.Keys
	}
	return nil
}

var File_config_oss_model_gen_proto protoreflect.FileDescriptor

var file_config_oss_model_gen_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x6f, 0x73, 0x73, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x04, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x4f, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x18, 0x65, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x18,
	0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79,
	0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x67, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x6f, 0x6c, 0x65, 0x41, 0x72, 0x6e, 0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72,
	0x6f, 0x6c, 0x65, 0x41, 0x72, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x6f, 0x6c, 0x65, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x69, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x72, 0x6f, 0x6c, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x18, 0x6b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x6c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x6d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x6e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x6f, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x70, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x71, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x72, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38,
	0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x73, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xb5, 0x04, 0x0a, 0x0f, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x4f, 0x73, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x65, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x67, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x41, 0x72, 0x6e, 0x18,
	0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x41, 0x72, 0x6e, 0x12, 0x28,
	0x0a, 0x0f, 0x72, 0x6f, 0x6c, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x69, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x6f, 0x6c, 0x65, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x6b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6f, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x6c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b,
	0x18, 0x6d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x18, 0x6f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x18, 0x70, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x71, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x72, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x73, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0xaa, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f, 0x73, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f, 0x73, 0x73, 0x52, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x18,
	0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12,
	0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x18, 0x69, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x22, 0x7f, 0x0a,
	0x11, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x68,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f, 0x73, 0x73, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x5e,
	0x0a, 0x1c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f, 0x73, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x66,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f, 0x73, 0x73, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x25,
	0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f, 0x73, 0x73, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x6b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4f, 0x0a, 0x1b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f,
	0x73, 0x73, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x65, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x4f, 0x73, 0x73, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79,
	0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_config_oss_model_gen_proto_rawDescOnce sync.Once
	file_config_oss_model_gen_proto_rawDescData = file_config_oss_model_gen_proto_rawDesc
)

func file_config_oss_model_gen_proto_rawDescGZIP() []byte {
	file_config_oss_model_gen_proto_rawDescOnce.Do(func() {
		file_config_oss_model_gen_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_oss_model_gen_proto_rawDescData)
	})
	return file_config_oss_model_gen_proto_rawDescData
}

var file_config_oss_model_gen_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_config_oss_model_gen_proto_goTypes = []interface{}{
	(*ConfigOss)(nil),                    // 0: service.ConfigOss
	(*ConfigOssFilter)(nil),              // 1: service.ConfigOssFilter
	(*ConfigOssListRequest)(nil),         // 2: service.ConfigOssListRequest
	(*ConfigOssResponse)(nil),            // 3: service.ConfigOssResponse
	(*ConfigOssUpdateFieldsRequest)(nil), // 4: service.ConfigOssUpdateFieldsRequest
	(*ConfigOssPrimarykey)(nil),          // 5: service.ConfigOssPrimarykey
	(*ConfigOssBatchDeleteRequest)(nil),  // 6: service.ConfigOssBatchDeleteRequest
	(*timestamppb.Timestamp)(nil),        // 7: google.protobuf.Timestamp
}
var file_config_oss_model_gen_proto_depIdxs = []int32{
	7,  // 0: service.ConfigOss.createdAt:type_name -> google.protobuf.Timestamp
	7,  // 1: service.ConfigOss.updatedAt:type_name -> google.protobuf.Timestamp
	7,  // 2: service.ConfigOss.deletedAt:type_name -> google.protobuf.Timestamp
	7,  // 3: service.ConfigOssFilter.createdAt:type_name -> google.protobuf.Timestamp
	7,  // 4: service.ConfigOssFilter.updatedAt:type_name -> google.protobuf.Timestamp
	7,  // 5: service.ConfigOssFilter.deletedAt:type_name -> google.protobuf.Timestamp
	0,  // 6: service.ConfigOssListRequest.query:type_name -> service.ConfigOss
	0,  // 7: service.ConfigOssResponse.data:type_name -> service.ConfigOss
	0,  // 8: service.ConfigOssUpdateFieldsRequest.data:type_name -> service.ConfigOss
	5,  // 9: service.ConfigOssBatchDeleteRequest.keys:type_name -> service.ConfigOssPrimarykey
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_config_oss_model_gen_proto_init() }
func file_config_oss_model_gen_proto_init() {
	if File_config_oss_model_gen_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_oss_model_gen_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigOss); i {
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
		file_config_oss_model_gen_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigOssFilter); i {
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
		file_config_oss_model_gen_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigOssListRequest); i {
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
		file_config_oss_model_gen_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigOssResponse); i {
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
		file_config_oss_model_gen_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigOssUpdateFieldsRequest); i {
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
		file_config_oss_model_gen_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigOssPrimarykey); i {
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
		file_config_oss_model_gen_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigOssBatchDeleteRequest); i {
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
			RawDescriptor: file_config_oss_model_gen_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_oss_model_gen_proto_goTypes,
		DependencyIndexes: file_config_oss_model_gen_proto_depIdxs,
		MessageInfos:      file_config_oss_model_gen_proto_msgTypes,
	}.Build()
	File_config_oss_model_gen_proto = out.File
	file_config_oss_model_gen_proto_rawDesc = nil
	file_config_oss_model_gen_proto_goTypes = nil
	file_config_oss_model_gen_proto_depIdxs = nil
}
