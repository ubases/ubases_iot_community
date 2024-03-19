// Code generated by protoc,2022-04-18 19:12:08. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: sys_dept_model.gen.proto

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
type SysDept struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"deptId" gorm:"default:部门id;comment:部门id;size:19;"
	DeptId int64 `protobuf:"varint,100,opt,name=deptId,proto3" json:"deptId,omitempty"`
	// @inject_tag: json:"parentId" gorm:"default:父部门id;comment:父部门id;size:19;"
	ParentId int64 `protobuf:"varint,101,opt,name=parentId,proto3" json:"parentId,omitempty"`
	// @inject_tag: json:"ancestors" gorm:"default:祖级列表;comment:祖级列表;size:50;"
	Ancestors string `protobuf:"bytes,102,opt,name=ancestors,proto3" json:"ancestors,omitempty"`
	// @inject_tag: json:"deptName" gorm:"default:部门名称;comment:部门名称;size:30;"
	DeptName string `protobuf:"bytes,103,opt,name=deptName,proto3" json:"deptName,omitempty"`
	// @inject_tag: json:"orderNum" gorm:"default:显示顺序;comment:显示顺序;size:10;"
	OrderNum int32 `protobuf:"varint,104,opt,name=orderNum,proto3" json:"orderNum,omitempty"`
	// @inject_tag: json:"leader" gorm:"default:负责人;comment:负责人;size:20;"
	Leader string `protobuf:"bytes,105,opt,name=leader,proto3" json:"leader,omitempty"`
	// @inject_tag: json:"phone" gorm:"default:联系电话;comment:联系电话;size:11;"
	Phone string `protobuf:"bytes,106,opt,name=phone,proto3" json:"phone,omitempty"`
	// @inject_tag: json:"email" gorm:"default:邮箱;comment:邮箱;size:50;"
	Email string `protobuf:"bytes,107,opt,name=email,proto3" json:"email,omitempty"`
	// @inject_tag: json:"status" gorm:"default:部门状态（0正常 1停用）;comment:部门状态（0正常 1停用）;"
	Status string `protobuf:"bytes,108,opt,name=status,proto3" json:"status,omitempty"`
	// @inject_tag: json:"createdBy" gorm:"default:创建人;comment:创建人;size:20;"
	CreatedBy int64 `protobuf:"varint,109,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	// @inject_tag: json:"updatedBy" gorm:"default:修改人;comment:修改人;size:19;"
	UpdatedBy int64 `protobuf:"varint,110,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	// @inject_tag: json:"createdAt" gorm:"default:创建时间;comment:创建时间;"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,111,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// @inject_tag: json:"updatedAt" gorm:"default:修改时间;comment:修改时间;"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,112,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// @inject_tag: json:"deletedAt" gorm:"default:删除时间;comment:删除时间;"
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,113,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *SysDept) Reset() {
	*x = SysDept{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_dept_model_gen_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysDept) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysDept) ProtoMessage() {}

func (x *SysDept) ProtoReflect() protoreflect.Message {
	mi := &file_sys_dept_model_gen_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysDept.ProtoReflect.Descriptor instead.
func (*SysDept) Descriptor() ([]byte, []int) {
	return file_sys_dept_model_gen_proto_rawDescGZIP(), []int{0}
}

func (x *SysDept) GetDeptId() int64 {
	if x != nil {
		return x.DeptId
	}
	return 0
}

func (x *SysDept) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *SysDept) GetAncestors() string {
	if x != nil {
		return x.Ancestors
	}
	return ""
}

func (x *SysDept) GetDeptName() string {
	if x != nil {
		return x.DeptName
	}
	return ""
}

func (x *SysDept) GetOrderNum() int32 {
	if x != nil {
		return x.OrderNum
	}
	return 0
}

func (x *SysDept) GetLeader() string {
	if x != nil {
		return x.Leader
	}
	return ""
}

func (x *SysDept) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *SysDept) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SysDept) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *SysDept) GetCreatedBy() int64 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *SysDept) GetUpdatedBy() int64 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

func (x *SysDept) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *SysDept) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *SysDept) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type SysDeptFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"deptId" gorm:"default:部门id;comment:部门id;size:19;"
	DeptId int64 `protobuf:"varint,100,opt,name=deptId,proto3" json:"deptId,omitempty"`
	// @inject_tag: json:"parentId" gorm:"default:父部门id;comment:父部门id;size:19;"
	ParentId int64 `protobuf:"varint,101,opt,name=parentId,proto3" json:"parentId,omitempty"`
	// @inject_tag: json:"ancestors" gorm:"default:祖级列表;comment:祖级列表;size:50;"
	Ancestors string `protobuf:"bytes,102,opt,name=ancestors,proto3" json:"ancestors,omitempty"`
	// @inject_tag: json:"deptName" gorm:"default:部门名称;comment:部门名称;size:30;"
	DeptName string `protobuf:"bytes,103,opt,name=deptName,proto3" json:"deptName,omitempty"`
	// @inject_tag: json:"orderNum" gorm:"default:显示顺序;comment:显示顺序;size:10;"
	OrderNum int32 `protobuf:"varint,104,opt,name=orderNum,proto3" json:"orderNum,omitempty"`
	// @inject_tag: json:"leader" gorm:"default:负责人;comment:负责人;size:20;"
	Leader string `protobuf:"bytes,105,opt,name=leader,proto3" json:"leader,omitempty"`
	// @inject_tag: json:"phone" gorm:"default:联系电话;comment:联系电话;size:11;"
	Phone string `protobuf:"bytes,106,opt,name=phone,proto3" json:"phone,omitempty"`
	// @inject_tag: json:"email" gorm:"default:邮箱;comment:邮箱;size:50;"
	Email string `protobuf:"bytes,107,opt,name=email,proto3" json:"email,omitempty"`
	// @inject_tag: json:"status" gorm:"default:部门状态（0正常 1停用）;comment:部门状态（0正常 1停用）;"
	Status string `protobuf:"bytes,108,opt,name=status,proto3" json:"status,omitempty"`
	// @inject_tag: json:"createdBy" gorm:"default:创建人;comment:创建人;size:20;"
	CreatedBy int64 `protobuf:"varint,109,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	// @inject_tag: json:"updatedBy" gorm:"default:修改人;comment:修改人;size:19;"
	UpdatedBy int64 `protobuf:"varint,110,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	// @inject_tag: json:"createdAt" gorm:"default:创建时间;comment:创建时间;"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,111,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// @inject_tag: json:"updatedAt" gorm:"default:修改时间;comment:修改时间;"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,112,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// @inject_tag: json:"deletedAt" gorm:"default:删除时间;comment:删除时间;"
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,113,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *SysDeptFilter) Reset() {
	*x = SysDeptFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_dept_model_gen_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysDeptFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysDeptFilter) ProtoMessage() {}

func (x *SysDeptFilter) ProtoReflect() protoreflect.Message {
	mi := &file_sys_dept_model_gen_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysDeptFilter.ProtoReflect.Descriptor instead.
func (*SysDeptFilter) Descriptor() ([]byte, []int) {
	return file_sys_dept_model_gen_proto_rawDescGZIP(), []int{1}
}

func (x *SysDeptFilter) GetDeptId() int64 {
	if x != nil {
		return x.DeptId
	}
	return 0
}

func (x *SysDeptFilter) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *SysDeptFilter) GetAncestors() string {
	if x != nil {
		return x.Ancestors
	}
	return ""
}

func (x *SysDeptFilter) GetDeptName() string {
	if x != nil {
		return x.DeptName
	}
	return ""
}

func (x *SysDeptFilter) GetOrderNum() int32 {
	if x != nil {
		return x.OrderNum
	}
	return 0
}

func (x *SysDeptFilter) GetLeader() string {
	if x != nil {
		return x.Leader
	}
	return ""
}

func (x *SysDeptFilter) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *SysDeptFilter) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SysDeptFilter) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *SysDeptFilter) GetCreatedBy() int64 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *SysDeptFilter) GetUpdatedBy() int64 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

func (x *SysDeptFilter) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *SysDeptFilter) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *SysDeptFilter) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type SysDeptListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query     *SysDept `protobuf:"bytes,101,opt,name=query,proto3" json:"query,omitempty"`
	Page      int64    `protobuf:"varint,102,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int64    `protobuf:"varint,103,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	OrderKey  string   `protobuf:"bytes,104,opt,name=orderKey,proto3" json:"orderKey,omitempty"`
	OrderDesc string   `protobuf:"bytes,105,opt,name=orderDesc,proto3" json:"orderDesc,omitempty"`
	SearchKey string   `protobuf:"bytes,106,opt,name=searchKey,proto3" json:"searchKey,omitempty"`
	BeginTime string   `protobuf:"bytes,107,opt,name=beginTime,proto3" json:"beginTime,omitempty"`
	EndTime   string   `protobuf:"bytes,108,opt,name=endTime,proto3" json:"endTime,omitempty"`
}

func (x *SysDeptListRequest) Reset() {
	*x = SysDeptListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_dept_model_gen_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysDeptListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysDeptListRequest) ProtoMessage() {}

func (x *SysDeptListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sys_dept_model_gen_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysDeptListRequest.ProtoReflect.Descriptor instead.
func (*SysDeptListRequest) Descriptor() ([]byte, []int) {
	return file_sys_dept_model_gen_proto_rawDescGZIP(), []int{2}
}

func (x *SysDeptListRequest) GetQuery() *SysDept {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *SysDeptListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SysDeptListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SysDeptListRequest) GetOrderKey() string {
	if x != nil {
		return x.OrderKey
	}
	return ""
}

func (x *SysDeptListRequest) GetOrderDesc() string {
	if x != nil {
		return x.OrderDesc
	}
	return ""
}

func (x *SysDeptListRequest) GetSearchKey() string {
	if x != nil {
		return x.SearchKey
	}
	return ""
}

func (x *SysDeptListRequest) GetBeginTime() string {
	if x != nil {
		return x.BeginTime
	}
	return ""
}

func (x *SysDeptListRequest) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

type SysDeptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32      `protobuf:"varint,101,opt,name=code,proto3" json:"code,omitempty"`
	Message string     `protobuf:"bytes,102,opt,name=message,proto3" json:"message,omitempty"`
	Total   int64      `protobuf:"varint,103,opt,name=total,proto3" json:"total,omitempty"`
	Data    []*SysDept `protobuf:"bytes,104,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *SysDeptResponse) Reset() {
	*x = SysDeptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_dept_model_gen_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysDeptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysDeptResponse) ProtoMessage() {}

func (x *SysDeptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sys_dept_model_gen_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysDeptResponse.ProtoReflect.Descriptor instead.
func (*SysDeptResponse) Descriptor() ([]byte, []int) {
	return file_sys_dept_model_gen_proto_rawDescGZIP(), []int{3}
}

func (x *SysDeptResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SysDeptResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SysDeptResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *SysDeptResponse) GetData() []*SysDept {
	if x != nil {
		return x.Data
	}
	return nil
}

type SysDeptUpdateFieldsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []string `protobuf:"bytes,101,rep,name=fields,proto3" json:"fields,omitempty"`
	Data   *SysDept `protobuf:"bytes,102,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SysDeptUpdateFieldsRequest) Reset() {
	*x = SysDeptUpdateFieldsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_dept_model_gen_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysDeptUpdateFieldsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysDeptUpdateFieldsRequest) ProtoMessage() {}

func (x *SysDeptUpdateFieldsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sys_dept_model_gen_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysDeptUpdateFieldsRequest.ProtoReflect.Descriptor instead.
func (*SysDeptUpdateFieldsRequest) Descriptor() ([]byte, []int) {
	return file_sys_dept_model_gen_proto_rawDescGZIP(), []int{4}
}

func (x *SysDeptUpdateFieldsRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *SysDeptUpdateFieldsRequest) GetData() *SysDept {
	if x != nil {
		return x.Data
	}
	return nil
}

// 表模型主键，只能是整数或字符串，若生成错误，请检查数据库表主键是否正确。支持组合主键
type SysDeptPrimarykey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeptId int64 `protobuf:"varint,100,opt,name=deptId,proto3" json:"deptId,omitempty"`
}

func (x *SysDeptPrimarykey) Reset() {
	*x = SysDeptPrimarykey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_dept_model_gen_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysDeptPrimarykey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysDeptPrimarykey) ProtoMessage() {}

func (x *SysDeptPrimarykey) ProtoReflect() protoreflect.Message {
	mi := &file_sys_dept_model_gen_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysDeptPrimarykey.ProtoReflect.Descriptor instead.
func (*SysDeptPrimarykey) Descriptor() ([]byte, []int) {
	return file_sys_dept_model_gen_proto_rawDescGZIP(), []int{5}
}

func (x *SysDeptPrimarykey) GetDeptId() int64 {
	if x != nil {
		return x.DeptId
	}
	return 0
}

type SysDeptBatchDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []*SysDeptPrimarykey `protobuf:"bytes,101,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *SysDeptBatchDeleteRequest) Reset() {
	*x = SysDeptBatchDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_dept_model_gen_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysDeptBatchDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysDeptBatchDeleteRequest) ProtoMessage() {}

func (x *SysDeptBatchDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sys_dept_model_gen_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysDeptBatchDeleteRequest.ProtoReflect.Descriptor instead.
func (*SysDeptBatchDeleteRequest) Descriptor() ([]byte, []int) {
	return file_sys_dept_model_gen_proto_rawDescGZIP(), []int{6}
}

func (x *SysDeptBatchDeleteRequest) GetKeys() []*SysDeptPrimarykey {
	if x != nil {
		return x.Keys
	}
	return nil
}

var File_sys_dept_model_gen_proto protoreflect.FileDescriptor

var file_sys_dept_model_gen_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x79, 0x73, 0x5f, 0x64, 0x65, 0x70, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x03, 0x0a, 0x07, 0x53, 0x79, 0x73, 0x44, 0x65, 0x70, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x70, 0x74, 0x49, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x64, 0x65, 0x70, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x18, 0x65, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x73, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x70, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x67,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x70, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x18, 0x68, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x18, 0x69, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x6a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x6b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x6c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x18, 0x6d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x6f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x70, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x71, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0xdf, 0x03, 0x0a, 0x0d, 0x53, 0x79, 0x73, 0x44, 0x65, 0x70, 0x74, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x70, 0x74, 0x49, 0x64, 0x18, 0x64, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x64, 0x65, 0x70, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x65, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x73, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x6e, 0x63, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x70, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x67, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x70, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x18, 0x68, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x16, 0x0a, 0x06,
	0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x69, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x6a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x6b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x6c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x6d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x42, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x6f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38,
	0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x70, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x71, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0xfc, 0x01, 0x0a, 0x12, 0x53, 0x79, 0x73, 0x44, 0x65, 0x70, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x79, 0x73, 0x44, 0x65, 0x70, 0x74, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x18, 0x68, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a,
	0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x18, 0x69, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x65, 0x67,
	0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x6b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x65,
	0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x6c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x7b, 0x0a, 0x0f, 0x53, 0x79, 0x73, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x65, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x67, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x68, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x53, 0x79, 0x73, 0x44, 0x65, 0x70, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x5a,
	0x0a, 0x1a, 0x53, 0x79, 0x73, 0x44, 0x65, 0x70, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x66, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x79, 0x73,
	0x44, 0x65, 0x70, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2b, 0x0a, 0x11, 0x53, 0x79,
	0x73, 0x44, 0x65, 0x70, 0x74, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x65, 0x70, 0x74, 0x49, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x64, 0x65, 0x70, 0x74, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x19, 0x53, 0x79, 0x73, 0x44, 0x65,
	0x70, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x65, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x79, 0x73,
	0x44, 0x65, 0x70, 0x74, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x52, 0x04,
	0x6b, 0x65, 0x79, 0x73, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sys_dept_model_gen_proto_rawDescOnce sync.Once
	file_sys_dept_model_gen_proto_rawDescData = file_sys_dept_model_gen_proto_rawDesc
)

func file_sys_dept_model_gen_proto_rawDescGZIP() []byte {
	file_sys_dept_model_gen_proto_rawDescOnce.Do(func() {
		file_sys_dept_model_gen_proto_rawDescData = protoimpl.X.CompressGZIP(file_sys_dept_model_gen_proto_rawDescData)
	})
	return file_sys_dept_model_gen_proto_rawDescData
}

var file_sys_dept_model_gen_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_sys_dept_model_gen_proto_goTypes = []interface{}{
	(*SysDept)(nil),                    // 0: service.SysDept
	(*SysDeptFilter)(nil),              // 1: service.SysDeptFilter
	(*SysDeptListRequest)(nil),         // 2: service.SysDeptListRequest
	(*SysDeptResponse)(nil),            // 3: service.SysDeptResponse
	(*SysDeptUpdateFieldsRequest)(nil), // 4: service.SysDeptUpdateFieldsRequest
	(*SysDeptPrimarykey)(nil),          // 5: service.SysDeptPrimarykey
	(*SysDeptBatchDeleteRequest)(nil),  // 6: service.SysDeptBatchDeleteRequest
	(*timestamppb.Timestamp)(nil),      // 7: google.protobuf.Timestamp
}
var file_sys_dept_model_gen_proto_depIdxs = []int32{
	7,  // 0: service.SysDept.createdAt:type_name -> google.protobuf.Timestamp
	7,  // 1: service.SysDept.updatedAt:type_name -> google.protobuf.Timestamp
	7,  // 2: service.SysDept.deletedAt:type_name -> google.protobuf.Timestamp
	7,  // 3: service.SysDeptFilter.createdAt:type_name -> google.protobuf.Timestamp
	7,  // 4: service.SysDeptFilter.updatedAt:type_name -> google.protobuf.Timestamp
	7,  // 5: service.SysDeptFilter.deletedAt:type_name -> google.protobuf.Timestamp
	0,  // 6: service.SysDeptListRequest.query:type_name -> service.SysDept
	0,  // 7: service.SysDeptResponse.data:type_name -> service.SysDept
	0,  // 8: service.SysDeptUpdateFieldsRequest.data:type_name -> service.SysDept
	5,  // 9: service.SysDeptBatchDeleteRequest.keys:type_name -> service.SysDeptPrimarykey
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_sys_dept_model_gen_proto_init() }
func file_sys_dept_model_gen_proto_init() {
	if File_sys_dept_model_gen_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sys_dept_model_gen_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysDept); i {
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
		file_sys_dept_model_gen_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysDeptFilter); i {
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
		file_sys_dept_model_gen_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysDeptListRequest); i {
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
		file_sys_dept_model_gen_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysDeptResponse); i {
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
		file_sys_dept_model_gen_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysDeptUpdateFieldsRequest); i {
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
		file_sys_dept_model_gen_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysDeptPrimarykey); i {
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
		file_sys_dept_model_gen_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysDeptBatchDeleteRequest); i {
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
			RawDescriptor: file_sys_dept_model_gen_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sys_dept_model_gen_proto_goTypes,
		DependencyIndexes: file_sys_dept_model_gen_proto_depIdxs,
		MessageInfos:      file_sys_dept_model_gen_proto_msgTypes,
	}.Build()
	File_sys_dept_model_gen_proto = out.File
	file_sys_dept_model_gen_proto_rawDesc = nil
	file_sys_dept_model_gen_proto_goTypes = nil
	file_sys_dept_model_gen_proto_depIdxs = nil
}
