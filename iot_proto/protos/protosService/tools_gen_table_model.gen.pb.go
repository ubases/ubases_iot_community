// Code generated by protoc,2022-04-18 19:12:11. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: tools_gen_table_model.gen.proto

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
type ToolsGenTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"tableId" gorm:"primary_key;AUTO_INCREMENT;default:编号;comment:编号;size:19;"
	TableId int64 `protobuf:"varint,100,opt,name=tableId,proto3" json:"tableId,omitempty"`
	// @inject_tag: json:"tableNameE" gorm:"default:表名称;comment:表名称;size:200;"
	TableNameE string `protobuf:"bytes,101,opt,name=tableNameE,proto3" json:"tableNameE,omitempty"`
	// @inject_tag: json:"tableComment" gorm:"default:表描述;comment:表描述;size:500;"
	TableComment string `protobuf:"bytes,102,opt,name=tableComment,proto3" json:"tableComment,omitempty"`
	// @inject_tag: json:"className" gorm:"default:实体类名称;comment:实体类名称;size:100;"
	ClassName string `protobuf:"bytes,103,opt,name=className,proto3" json:"className,omitempty"`
	// @inject_tag: json:"tplCategory" gorm:"default:使用的模板（crud单表操作 tree树表操作）;comment:使用的模板（crud单表操作 tree树表操作）;size:200;"
	TplCategory string `protobuf:"bytes,104,opt,name=tplCategory,proto3" json:"tplCategory,omitempty"`
	// @inject_tag: json:"packageName" gorm:"default:生成包路径;comment:生成包路径;size:100;"
	PackageName string `protobuf:"bytes,105,opt,name=packageName,proto3" json:"packageName,omitempty"`
	// @inject_tag: json:"moduleName" gorm:"default:生成模块名;comment:生成模块名;size:30;"
	ModuleName string `protobuf:"bytes,106,opt,name=moduleName,proto3" json:"moduleName,omitempty"`
	// @inject_tag: json:"businessName" gorm:"default:生成业务名;comment:生成业务名;size:30;"
	BusinessName string `protobuf:"bytes,107,opt,name=businessName,proto3" json:"businessName,omitempty"`
	// @inject_tag: json:"functionName" gorm:"default:生成功能名;comment:生成功能名;size:50;"
	FunctionName string `protobuf:"bytes,108,opt,name=functionName,proto3" json:"functionName,omitempty"`
	// @inject_tag: json:"functionAuthor" gorm:"default:生成功能作者;comment:生成功能作者;size:50;"
	FunctionAuthor string `protobuf:"bytes,109,opt,name=functionAuthor,proto3" json:"functionAuthor,omitempty"`
	// @inject_tag: json:"options" gorm:"default:其它生成选项;comment:其它生成选项;size:1000;"
	Options string `protobuf:"bytes,110,opt,name=options,proto3" json:"options,omitempty"`
	// @inject_tag: json:"createTime" gorm:"default:创建时间;comment:创建时间;"
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,111,opt,name=createTime,proto3" json:"createTime,omitempty"`
	// @inject_tag: json:"updateTime" gorm:"default:更新时间;comment:更新时间;"
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,112,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	// @inject_tag: json:"remark" gorm:"default:备注;comment:备注;size:500;"
	Remark string `protobuf:"bytes,113,opt,name=remark,proto3" json:"remark,omitempty"`
}

func (x *ToolsGenTable) Reset() {
	*x = ToolsGenTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tools_gen_table_model_gen_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToolsGenTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToolsGenTable) ProtoMessage() {}

func (x *ToolsGenTable) ProtoReflect() protoreflect.Message {
	mi := &file_tools_gen_table_model_gen_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToolsGenTable.ProtoReflect.Descriptor instead.
func (*ToolsGenTable) Descriptor() ([]byte, []int) {
	return file_tools_gen_table_model_gen_proto_rawDescGZIP(), []int{0}
}

func (x *ToolsGenTable) GetTableId() int64 {
	if x != nil {
		return x.TableId
	}
	return 0
}

func (x *ToolsGenTable) GetTableNameE() string {
	if x != nil {
		return x.TableNameE
	}
	return ""
}

func (x *ToolsGenTable) GetTableComment() string {
	if x != nil {
		return x.TableComment
	}
	return ""
}

func (x *ToolsGenTable) GetClassName() string {
	if x != nil {
		return x.ClassName
	}
	return ""
}

func (x *ToolsGenTable) GetTplCategory() string {
	if x != nil {
		return x.TplCategory
	}
	return ""
}

func (x *ToolsGenTable) GetPackageName() string {
	if x != nil {
		return x.PackageName
	}
	return ""
}

func (x *ToolsGenTable) GetModuleName() string {
	if x != nil {
		return x.ModuleName
	}
	return ""
}

func (x *ToolsGenTable) GetBusinessName() string {
	if x != nil {
		return x.BusinessName
	}
	return ""
}

func (x *ToolsGenTable) GetFunctionName() string {
	if x != nil {
		return x.FunctionName
	}
	return ""
}

func (x *ToolsGenTable) GetFunctionAuthor() string {
	if x != nil {
		return x.FunctionAuthor
	}
	return ""
}

func (x *ToolsGenTable) GetOptions() string {
	if x != nil {
		return x.Options
	}
	return ""
}

func (x *ToolsGenTable) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *ToolsGenTable) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *ToolsGenTable) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type ToolsGenTableFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"tableId" gorm:"primary_key;AUTO_INCREMENT;default:编号;comment:编号;size:19;"
	TableId int64 `protobuf:"varint,100,opt,name=tableId,proto3" json:"tableId,omitempty"`
	// @inject_tag: json:"tableNameE" gorm:"default:表名称;comment:表名称;size:200;"
	TableNameE string `protobuf:"bytes,101,opt,name=tableNameE,proto3" json:"tableNameE,omitempty"`
	// @inject_tag: json:"tableComment" gorm:"default:表描述;comment:表描述;size:500;"
	TableComment string `protobuf:"bytes,102,opt,name=tableComment,proto3" json:"tableComment,omitempty"`
	// @inject_tag: json:"className" gorm:"default:实体类名称;comment:实体类名称;size:100;"
	ClassName string `protobuf:"bytes,103,opt,name=className,proto3" json:"className,omitempty"`
	// @inject_tag: json:"tplCategory" gorm:"default:使用的模板（crud单表操作 tree树表操作）;comment:使用的模板（crud单表操作 tree树表操作）;size:200;"
	TplCategory string `protobuf:"bytes,104,opt,name=tplCategory,proto3" json:"tplCategory,omitempty"`
	// @inject_tag: json:"packageName" gorm:"default:生成包路径;comment:生成包路径;size:100;"
	PackageName string `protobuf:"bytes,105,opt,name=packageName,proto3" json:"packageName,omitempty"`
	// @inject_tag: json:"moduleName" gorm:"default:生成模块名;comment:生成模块名;size:30;"
	ModuleName string `protobuf:"bytes,106,opt,name=moduleName,proto3" json:"moduleName,omitempty"`
	// @inject_tag: json:"businessName" gorm:"default:生成业务名;comment:生成业务名;size:30;"
	BusinessName string `protobuf:"bytes,107,opt,name=businessName,proto3" json:"businessName,omitempty"`
	// @inject_tag: json:"functionName" gorm:"default:生成功能名;comment:生成功能名;size:50;"
	FunctionName string `protobuf:"bytes,108,opt,name=functionName,proto3" json:"functionName,omitempty"`
	// @inject_tag: json:"functionAuthor" gorm:"default:生成功能作者;comment:生成功能作者;size:50;"
	FunctionAuthor string `protobuf:"bytes,109,opt,name=functionAuthor,proto3" json:"functionAuthor,omitempty"`
	// @inject_tag: json:"options" gorm:"default:其它生成选项;comment:其它生成选项;size:1000;"
	Options string `protobuf:"bytes,110,opt,name=options,proto3" json:"options,omitempty"`
	// @inject_tag: json:"createTime" gorm:"default:创建时间;comment:创建时间;"
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,111,opt,name=createTime,proto3" json:"createTime,omitempty"`
	// @inject_tag: json:"updateTime" gorm:"default:更新时间;comment:更新时间;"
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,112,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	// @inject_tag: json:"remark" gorm:"default:备注;comment:备注;size:500;"
	Remark string `protobuf:"bytes,113,opt,name=remark,proto3" json:"remark,omitempty"`
}

func (x *ToolsGenTableFilter) Reset() {
	*x = ToolsGenTableFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tools_gen_table_model_gen_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToolsGenTableFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToolsGenTableFilter) ProtoMessage() {}

func (x *ToolsGenTableFilter) ProtoReflect() protoreflect.Message {
	mi := &file_tools_gen_table_model_gen_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToolsGenTableFilter.ProtoReflect.Descriptor instead.
func (*ToolsGenTableFilter) Descriptor() ([]byte, []int) {
	return file_tools_gen_table_model_gen_proto_rawDescGZIP(), []int{1}
}

func (x *ToolsGenTableFilter) GetTableId() int64 {
	if x != nil {
		return x.TableId
	}
	return 0
}

func (x *ToolsGenTableFilter) GetTableNameE() string {
	if x != nil {
		return x.TableNameE
	}
	return ""
}

func (x *ToolsGenTableFilter) GetTableComment() string {
	if x != nil {
		return x.TableComment
	}
	return ""
}

func (x *ToolsGenTableFilter) GetClassName() string {
	if x != nil {
		return x.ClassName
	}
	return ""
}

func (x *ToolsGenTableFilter) GetTplCategory() string {
	if x != nil {
		return x.TplCategory
	}
	return ""
}

func (x *ToolsGenTableFilter) GetPackageName() string {
	if x != nil {
		return x.PackageName
	}
	return ""
}

func (x *ToolsGenTableFilter) GetModuleName() string {
	if x != nil {
		return x.ModuleName
	}
	return ""
}

func (x *ToolsGenTableFilter) GetBusinessName() string {
	if x != nil {
		return x.BusinessName
	}
	return ""
}

func (x *ToolsGenTableFilter) GetFunctionName() string {
	if x != nil {
		return x.FunctionName
	}
	return ""
}

func (x *ToolsGenTableFilter) GetFunctionAuthor() string {
	if x != nil {
		return x.FunctionAuthor
	}
	return ""
}

func (x *ToolsGenTableFilter) GetOptions() string {
	if x != nil {
		return x.Options
	}
	return ""
}

func (x *ToolsGenTableFilter) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *ToolsGenTableFilter) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *ToolsGenTableFilter) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type ToolsGenTableListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query     *ToolsGenTable `protobuf:"bytes,101,opt,name=query,proto3" json:"query,omitempty"`
	Page      int64          `protobuf:"varint,102,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int64          `protobuf:"varint,103,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	OrderKey  string         `protobuf:"bytes,104,opt,name=orderKey,proto3" json:"orderKey,omitempty"`
	OrderDesc string         `protobuf:"bytes,105,opt,name=orderDesc,proto3" json:"orderDesc,omitempty"`
}

func (x *ToolsGenTableListRequest) Reset() {
	*x = ToolsGenTableListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tools_gen_table_model_gen_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToolsGenTableListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToolsGenTableListRequest) ProtoMessage() {}

func (x *ToolsGenTableListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tools_gen_table_model_gen_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToolsGenTableListRequest.ProtoReflect.Descriptor instead.
func (*ToolsGenTableListRequest) Descriptor() ([]byte, []int) {
	return file_tools_gen_table_model_gen_proto_rawDescGZIP(), []int{2}
}

func (x *ToolsGenTableListRequest) GetQuery() *ToolsGenTable {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *ToolsGenTableListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ToolsGenTableListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ToolsGenTableListRequest) GetOrderKey() string {
	if x != nil {
		return x.OrderKey
	}
	return ""
}

func (x *ToolsGenTableListRequest) GetOrderDesc() string {
	if x != nil {
		return x.OrderDesc
	}
	return ""
}

type ToolsGenTableResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32            `protobuf:"varint,101,opt,name=code,proto3" json:"code,omitempty"`
	Message string           `protobuf:"bytes,102,opt,name=message,proto3" json:"message,omitempty"`
	Total   int64            `protobuf:"varint,103,opt,name=total,proto3" json:"total,omitempty"`
	Data    []*ToolsGenTable `protobuf:"bytes,104,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ToolsGenTableResponse) Reset() {
	*x = ToolsGenTableResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tools_gen_table_model_gen_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToolsGenTableResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToolsGenTableResponse) ProtoMessage() {}

func (x *ToolsGenTableResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tools_gen_table_model_gen_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToolsGenTableResponse.ProtoReflect.Descriptor instead.
func (*ToolsGenTableResponse) Descriptor() ([]byte, []int) {
	return file_tools_gen_table_model_gen_proto_rawDescGZIP(), []int{3}
}

func (x *ToolsGenTableResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ToolsGenTableResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ToolsGenTableResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ToolsGenTableResponse) GetData() []*ToolsGenTable {
	if x != nil {
		return x.Data
	}
	return nil
}

type ToolsGenTableUpdateFieldsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []string       `protobuf:"bytes,101,rep,name=fields,proto3" json:"fields,omitempty"`
	Data   *ToolsGenTable `protobuf:"bytes,102,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ToolsGenTableUpdateFieldsRequest) Reset() {
	*x = ToolsGenTableUpdateFieldsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tools_gen_table_model_gen_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToolsGenTableUpdateFieldsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToolsGenTableUpdateFieldsRequest) ProtoMessage() {}

func (x *ToolsGenTableUpdateFieldsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tools_gen_table_model_gen_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToolsGenTableUpdateFieldsRequest.ProtoReflect.Descriptor instead.
func (*ToolsGenTableUpdateFieldsRequest) Descriptor() ([]byte, []int) {
	return file_tools_gen_table_model_gen_proto_rawDescGZIP(), []int{4}
}

func (x *ToolsGenTableUpdateFieldsRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *ToolsGenTableUpdateFieldsRequest) GetData() *ToolsGenTable {
	if x != nil {
		return x.Data
	}
	return nil
}

// 表模型主键，只能是整数或字符串，若生成错误，请检查数据库表主键是否正确。支持组合主键
type ToolsGenTablePrimarykey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableId int64 `protobuf:"varint,100,opt,name=tableId,proto3" json:"tableId,omitempty"`
}

func (x *ToolsGenTablePrimarykey) Reset() {
	*x = ToolsGenTablePrimarykey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tools_gen_table_model_gen_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToolsGenTablePrimarykey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToolsGenTablePrimarykey) ProtoMessage() {}

func (x *ToolsGenTablePrimarykey) ProtoReflect() protoreflect.Message {
	mi := &file_tools_gen_table_model_gen_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToolsGenTablePrimarykey.ProtoReflect.Descriptor instead.
func (*ToolsGenTablePrimarykey) Descriptor() ([]byte, []int) {
	return file_tools_gen_table_model_gen_proto_rawDescGZIP(), []int{5}
}

func (x *ToolsGenTablePrimarykey) GetTableId() int64 {
	if x != nil {
		return x.TableId
	}
	return 0
}

type ToolsGenTableBatchDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []*ToolsGenTablePrimarykey `protobuf:"bytes,101,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *ToolsGenTableBatchDeleteRequest) Reset() {
	*x = ToolsGenTableBatchDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tools_gen_table_model_gen_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToolsGenTableBatchDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToolsGenTableBatchDeleteRequest) ProtoMessage() {}

func (x *ToolsGenTableBatchDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tools_gen_table_model_gen_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToolsGenTableBatchDeleteRequest.ProtoReflect.Descriptor instead.
func (*ToolsGenTableBatchDeleteRequest) Descriptor() ([]byte, []int) {
	return file_tools_gen_table_model_gen_proto_rawDescGZIP(), []int{6}
}

func (x *ToolsGenTableBatchDeleteRequest) GetKeys() []*ToolsGenTablePrimarykey {
	if x != nil {
		return x.Keys
	}
	return nil
}

var File_tools_gen_table_model_gen_proto protoreflect.FileDescriptor

var file_tools_gen_table_model_gen_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x5f, 0x67, 0x65, 0x6e, 0x5f, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x04, 0x0a, 0x0d,
	0x54, 0x6f, 0x6f, 0x6c, 0x73, 0x47, 0x65, 0x6e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x45, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x45, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x70, 0x6c,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x74, 0x70, 0x6c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x69, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x6a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x6b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x6c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x6d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x66,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x6f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x70, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x71, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x8f, 0x04, 0x0a, 0x13, 0x54, 0x6f, 0x6f, 0x6c,
	0x73, 0x47, 0x65, 0x6e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x45, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x45, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74,
	0x70, 0x6c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x68, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x74, 0x70, 0x6c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x20, 0x0a,
	0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x69, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x6a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x6b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x6c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x6d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x6f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x70, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x71, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0xb2, 0x01, 0x0a, 0x18, 0x54, 0x6f,
	0x6f, 0x6c, 0x73, 0x47, 0x65, 0x6e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x54, 0x6f, 0x6f, 0x6c, 0x73, 0x47, 0x65, 0x6e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x66, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79,
	0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79,
	0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x18, 0x69, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x22, 0x87,
	0x01, 0x0a, 0x15, 0x54, 0x6f, 0x6f, 0x6c, 0x73, 0x47, 0x65, 0x6e, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x65, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2a, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x68, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x6f, 0x6f, 0x6c, 0x73, 0x47, 0x65, 0x6e, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x66, 0x0a, 0x20, 0x54, 0x6f, 0x6f, 0x6c,
	0x73, 0x47, 0x65, 0x6e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x66, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x6f, 0x6f,
	0x6c, 0x73, 0x47, 0x65, 0x6e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x33, 0x0a, 0x17, 0x54, 0x6f, 0x6f, 0x6c, 0x73, 0x47, 0x65, 0x6e, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x57, 0x0a, 0x1f, 0x54, 0x6f, 0x6f, 0x6c, 0x73, 0x47, 0x65,
	0x6e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73,
	0x18, 0x65, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x54, 0x6f, 0x6f, 0x6c, 0x73, 0x47, 0x65, 0x6e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x42, 0x11,
	0x5a, 0x0f, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tools_gen_table_model_gen_proto_rawDescOnce sync.Once
	file_tools_gen_table_model_gen_proto_rawDescData = file_tools_gen_table_model_gen_proto_rawDesc
)

func file_tools_gen_table_model_gen_proto_rawDescGZIP() []byte {
	file_tools_gen_table_model_gen_proto_rawDescOnce.Do(func() {
		file_tools_gen_table_model_gen_proto_rawDescData = protoimpl.X.CompressGZIP(file_tools_gen_table_model_gen_proto_rawDescData)
	})
	return file_tools_gen_table_model_gen_proto_rawDescData
}

var file_tools_gen_table_model_gen_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_tools_gen_table_model_gen_proto_goTypes = []interface{}{
	(*ToolsGenTable)(nil),                    // 0: service.ToolsGenTable
	(*ToolsGenTableFilter)(nil),              // 1: service.ToolsGenTableFilter
	(*ToolsGenTableListRequest)(nil),         // 2: service.ToolsGenTableListRequest
	(*ToolsGenTableResponse)(nil),            // 3: service.ToolsGenTableResponse
	(*ToolsGenTableUpdateFieldsRequest)(nil), // 4: service.ToolsGenTableUpdateFieldsRequest
	(*ToolsGenTablePrimarykey)(nil),          // 5: service.ToolsGenTablePrimarykey
	(*ToolsGenTableBatchDeleteRequest)(nil),  // 6: service.ToolsGenTableBatchDeleteRequest
	(*timestamppb.Timestamp)(nil),            // 7: google.protobuf.Timestamp
}
var file_tools_gen_table_model_gen_proto_depIdxs = []int32{
	7, // 0: service.ToolsGenTable.createTime:type_name -> google.protobuf.Timestamp
	7, // 1: service.ToolsGenTable.updateTime:type_name -> google.protobuf.Timestamp
	7, // 2: service.ToolsGenTableFilter.createTime:type_name -> google.protobuf.Timestamp
	7, // 3: service.ToolsGenTableFilter.updateTime:type_name -> google.protobuf.Timestamp
	0, // 4: service.ToolsGenTableListRequest.query:type_name -> service.ToolsGenTable
	0, // 5: service.ToolsGenTableResponse.data:type_name -> service.ToolsGenTable
	0, // 6: service.ToolsGenTableUpdateFieldsRequest.data:type_name -> service.ToolsGenTable
	5, // 7: service.ToolsGenTableBatchDeleteRequest.keys:type_name -> service.ToolsGenTablePrimarykey
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_tools_gen_table_model_gen_proto_init() }
func file_tools_gen_table_model_gen_proto_init() {
	if File_tools_gen_table_model_gen_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tools_gen_table_model_gen_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToolsGenTable); i {
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
		file_tools_gen_table_model_gen_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToolsGenTableFilter); i {
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
		file_tools_gen_table_model_gen_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToolsGenTableListRequest); i {
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
		file_tools_gen_table_model_gen_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToolsGenTableResponse); i {
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
		file_tools_gen_table_model_gen_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToolsGenTableUpdateFieldsRequest); i {
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
		file_tools_gen_table_model_gen_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToolsGenTablePrimarykey); i {
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
		file_tools_gen_table_model_gen_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToolsGenTableBatchDeleteRequest); i {
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
			RawDescriptor: file_tools_gen_table_model_gen_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tools_gen_table_model_gen_proto_goTypes,
		DependencyIndexes: file_tools_gen_table_model_gen_proto_depIdxs,
		MessageInfos:      file_tools_gen_table_model_gen_proto_msgTypes,
	}.Build()
	File_tools_gen_table_model_gen_proto = out.File
	file_tools_gen_table_model_gen_proto_rawDesc = nil
	file_tools_gen_table_model_gen_proto_goTypes = nil
	file_tools_gen_table_model_gen_proto_depIdxs = nil
}
