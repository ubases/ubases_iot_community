// Code generated by protoc,2022-04-18 19:12:09. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: sys_oper_log_model.gen.proto

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
type SysOperLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"operId" gorm:"primary_key;AUTO_INCREMENT;default:日志主键;comment:日志主键;size:20;"
	OperId int64 `protobuf:"varint,100,opt,name=operId,proto3" json:"operId,omitempty"`
	// @inject_tag: json:"title" gorm:"default:模块标题;comment:模块标题;size:50;"
	Title string `protobuf:"bytes,101,opt,name=title,proto3" json:"title,omitempty"`
	// @inject_tag: json:"businessType" gorm:"default:业务类型（0其它 1新增 2修改 3删除）;comment:业务类型（0其它 1新增 2修改 3删除）;size:10;"
	BusinessType int32 `protobuf:"varint,102,opt,name=businessType,proto3" json:"businessType,omitempty"`
	// @inject_tag: json:"method" gorm:"default:方法名称;comment:方法名称;size:100;"
	Method string `protobuf:"bytes,103,opt,name=method,proto3" json:"method,omitempty"`
	// @inject_tag: json:"requestMethod" gorm:"default:请求方式;comment:请求方式;size:10;"
	RequestMethod string `protobuf:"bytes,104,opt,name=requestMethod,proto3" json:"requestMethod,omitempty"`
	// @inject_tag: json:"operatorType" gorm:"default:操作类别（0其它 1后台用户 2手机端用户）;comment:操作类别（0其它 1后台用户 2手机端用户）;size:10;"
	OperatorType int32 `protobuf:"varint,105,opt,name=operatorType,proto3" json:"operatorType,omitempty"`
	// @inject_tag: json:"operName" gorm:"default:操作人员;comment:操作人员;size:50;"
	OperName string `protobuf:"bytes,106,opt,name=operName,proto3" json:"operName,omitempty"`
	// @inject_tag: json:"deptName" gorm:"default:部门名称;comment:部门名称;size:50;"
	DeptName string `protobuf:"bytes,107,opt,name=deptName,proto3" json:"deptName,omitempty"`
	// @inject_tag: json:"operUrl" gorm:"default:请求URL;comment:请求URL;size:500;"
	OperUrl string `protobuf:"bytes,108,opt,name=operUrl,proto3" json:"operUrl,omitempty"`
	// @inject_tag: json:"operIp" gorm:"default:主机地址;comment:主机地址;size:50;"
	OperIp string `protobuf:"bytes,109,opt,name=operIp,proto3" json:"operIp,omitempty"`
	// @inject_tag: json:"operLocation" gorm:"default:操作地点;comment:操作地点;size:255;"
	OperLocation string `protobuf:"bytes,110,opt,name=operLocation,proto3" json:"operLocation,omitempty"`
	// @inject_tag: json:"operParam" gorm:"default:请求参数;comment:请求参数;"
	OperParam string `protobuf:"bytes,111,opt,name=operParam,proto3" json:"operParam,omitempty"`
	// @inject_tag: json:"jsonResult" gorm:"default:返回参数;comment:返回参数;"
	JsonResult string `protobuf:"bytes,112,opt,name=jsonResult,proto3" json:"jsonResult,omitempty"`
	// @inject_tag: json:"status" gorm:"default:操作状态（0正常 1异常）;comment:操作状态（0正常 1异常）;size:10;"
	Status int32 `protobuf:"varint,113,opt,name=status,proto3" json:"status,omitempty"`
	// @inject_tag: json:"errorMsg" gorm:"default:错误消息;comment:错误消息;size:2000;"
	ErrorMsg string `protobuf:"bytes,114,opt,name=errorMsg,proto3" json:"errorMsg,omitempty"`
	// @inject_tag: json:"operTime" gorm:"default:操作时间;comment:操作时间;"
	OperTime *timestamppb.Timestamp `protobuf:"bytes,115,opt,name=operTime,proto3" json:"operTime,omitempty"`
}

func (x *SysOperLog) Reset() {
	*x = SysOperLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_oper_log_model_gen_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysOperLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysOperLog) ProtoMessage() {}

func (x *SysOperLog) ProtoReflect() protoreflect.Message {
	mi := &file_sys_oper_log_model_gen_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysOperLog.ProtoReflect.Descriptor instead.
func (*SysOperLog) Descriptor() ([]byte, []int) {
	return file_sys_oper_log_model_gen_proto_rawDescGZIP(), []int{0}
}

func (x *SysOperLog) GetOperId() int64 {
	if x != nil {
		return x.OperId
	}
	return 0
}

func (x *SysOperLog) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SysOperLog) GetBusinessType() int32 {
	if x != nil {
		return x.BusinessType
	}
	return 0
}

func (x *SysOperLog) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *SysOperLog) GetRequestMethod() string {
	if x != nil {
		return x.RequestMethod
	}
	return ""
}

func (x *SysOperLog) GetOperatorType() int32 {
	if x != nil {
		return x.OperatorType
	}
	return 0
}

func (x *SysOperLog) GetOperName() string {
	if x != nil {
		return x.OperName
	}
	return ""
}

func (x *SysOperLog) GetDeptName() string {
	if x != nil {
		return x.DeptName
	}
	return ""
}

func (x *SysOperLog) GetOperUrl() string {
	if x != nil {
		return x.OperUrl
	}
	return ""
}

func (x *SysOperLog) GetOperIp() string {
	if x != nil {
		return x.OperIp
	}
	return ""
}

func (x *SysOperLog) GetOperLocation() string {
	if x != nil {
		return x.OperLocation
	}
	return ""
}

func (x *SysOperLog) GetOperParam() string {
	if x != nil {
		return x.OperParam
	}
	return ""
}

func (x *SysOperLog) GetJsonResult() string {
	if x != nil {
		return x.JsonResult
	}
	return ""
}

func (x *SysOperLog) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SysOperLog) GetErrorMsg() string {
	if x != nil {
		return x.ErrorMsg
	}
	return ""
}

func (x *SysOperLog) GetOperTime() *timestamppb.Timestamp {
	if x != nil {
		return x.OperTime
	}
	return nil
}

type SysOperLogFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"operId" gorm:"primary_key;AUTO_INCREMENT;default:日志主键;comment:日志主键;size:20;"
	OperId int64 `protobuf:"varint,100,opt,name=operId,proto3" json:"operId,omitempty"`
	// @inject_tag: json:"title" gorm:"default:模块标题;comment:模块标题;size:50;"
	Title string `protobuf:"bytes,101,opt,name=title,proto3" json:"title,omitempty"`
	// @inject_tag: json:"businessType" gorm:"default:业务类型（0其它 1新增 2修改 3删除）;comment:业务类型（0其它 1新增 2修改 3删除）;size:10;"
	BusinessType int32 `protobuf:"varint,102,opt,name=businessType,proto3" json:"businessType,omitempty"`
	// @inject_tag: json:"method" gorm:"default:方法名称;comment:方法名称;size:100;"
	Method string `protobuf:"bytes,103,opt,name=method,proto3" json:"method,omitempty"`
	// @inject_tag: json:"requestMethod" gorm:"default:请求方式;comment:请求方式;size:10;"
	RequestMethod string `protobuf:"bytes,104,opt,name=requestMethod,proto3" json:"requestMethod,omitempty"`
	// @inject_tag: json:"operatorType" gorm:"default:操作类别（0其它 1后台用户 2手机端用户）;comment:操作类别（0其它 1后台用户 2手机端用户）;size:10;"
	OperatorType int32 `protobuf:"varint,105,opt,name=operatorType,proto3" json:"operatorType,omitempty"`
	// @inject_tag: json:"operName" gorm:"default:操作人员;comment:操作人员;size:50;"
	OperName string `protobuf:"bytes,106,opt,name=operName,proto3" json:"operName,omitempty"`
	// @inject_tag: json:"deptName" gorm:"default:部门名称;comment:部门名称;size:50;"
	DeptName string `protobuf:"bytes,107,opt,name=deptName,proto3" json:"deptName,omitempty"`
	// @inject_tag: json:"operUrl" gorm:"default:请求URL;comment:请求URL;size:500;"
	OperUrl string `protobuf:"bytes,108,opt,name=operUrl,proto3" json:"operUrl,omitempty"`
	// @inject_tag: json:"operIp" gorm:"default:主机地址;comment:主机地址;size:50;"
	OperIp string `protobuf:"bytes,109,opt,name=operIp,proto3" json:"operIp,omitempty"`
	// @inject_tag: json:"operLocation" gorm:"default:操作地点;comment:操作地点;size:255;"
	OperLocation string `protobuf:"bytes,110,opt,name=operLocation,proto3" json:"operLocation,omitempty"`
	// @inject_tag: json:"operParam" gorm:"default:请求参数;comment:请求参数;"
	OperParam string `protobuf:"bytes,111,opt,name=operParam,proto3" json:"operParam,omitempty"`
	// @inject_tag: json:"jsonResult" gorm:"default:返回参数;comment:返回参数;"
	JsonResult string `protobuf:"bytes,112,opt,name=jsonResult,proto3" json:"jsonResult,omitempty"`
	// @inject_tag: json:"status" gorm:"default:操作状态（0正常 1异常）;comment:操作状态（0正常 1异常）;size:10;"
	Status int32 `protobuf:"varint,113,opt,name=status,proto3" json:"status,omitempty"`
	// @inject_tag: json:"errorMsg" gorm:"default:错误消息;comment:错误消息;size:2000;"
	ErrorMsg string `protobuf:"bytes,114,opt,name=errorMsg,proto3" json:"errorMsg,omitempty"`
	// @inject_tag: json:"operTime" gorm:"default:操作时间;comment:操作时间;"
	OperTime *timestamppb.Timestamp `protobuf:"bytes,115,opt,name=operTime,proto3" json:"operTime,omitempty"`
}

func (x *SysOperLogFilter) Reset() {
	*x = SysOperLogFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_oper_log_model_gen_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysOperLogFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysOperLogFilter) ProtoMessage() {}

func (x *SysOperLogFilter) ProtoReflect() protoreflect.Message {
	mi := &file_sys_oper_log_model_gen_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysOperLogFilter.ProtoReflect.Descriptor instead.
func (*SysOperLogFilter) Descriptor() ([]byte, []int) {
	return file_sys_oper_log_model_gen_proto_rawDescGZIP(), []int{1}
}

func (x *SysOperLogFilter) GetOperId() int64 {
	if x != nil {
		return x.OperId
	}
	return 0
}

func (x *SysOperLogFilter) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SysOperLogFilter) GetBusinessType() int32 {
	if x != nil {
		return x.BusinessType
	}
	return 0
}

func (x *SysOperLogFilter) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *SysOperLogFilter) GetRequestMethod() string {
	if x != nil {
		return x.RequestMethod
	}
	return ""
}

func (x *SysOperLogFilter) GetOperatorType() int32 {
	if x != nil {
		return x.OperatorType
	}
	return 0
}

func (x *SysOperLogFilter) GetOperName() string {
	if x != nil {
		return x.OperName
	}
	return ""
}

func (x *SysOperLogFilter) GetDeptName() string {
	if x != nil {
		return x.DeptName
	}
	return ""
}

func (x *SysOperLogFilter) GetOperUrl() string {
	if x != nil {
		return x.OperUrl
	}
	return ""
}

func (x *SysOperLogFilter) GetOperIp() string {
	if x != nil {
		return x.OperIp
	}
	return ""
}

func (x *SysOperLogFilter) GetOperLocation() string {
	if x != nil {
		return x.OperLocation
	}
	return ""
}

func (x *SysOperLogFilter) GetOperParam() string {
	if x != nil {
		return x.OperParam
	}
	return ""
}

func (x *SysOperLogFilter) GetJsonResult() string {
	if x != nil {
		return x.JsonResult
	}
	return ""
}

func (x *SysOperLogFilter) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SysOperLogFilter) GetErrorMsg() string {
	if x != nil {
		return x.ErrorMsg
	}
	return ""
}

func (x *SysOperLogFilter) GetOperTime() *timestamppb.Timestamp {
	if x != nil {
		return x.OperTime
	}
	return nil
}

type SysOperLogListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query     *SysOperLog `protobuf:"bytes,101,opt,name=query,proto3" json:"query,omitempty"`
	Page      int64       `protobuf:"varint,102,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int64       `protobuf:"varint,103,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	OrderKey  string      `protobuf:"bytes,104,opt,name=orderKey,proto3" json:"orderKey,omitempty"`
	OrderDesc string      `protobuf:"bytes,105,opt,name=orderDesc,proto3" json:"orderDesc,omitempty"`
}

func (x *SysOperLogListRequest) Reset() {
	*x = SysOperLogListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_oper_log_model_gen_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysOperLogListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysOperLogListRequest) ProtoMessage() {}

func (x *SysOperLogListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sys_oper_log_model_gen_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysOperLogListRequest.ProtoReflect.Descriptor instead.
func (*SysOperLogListRequest) Descriptor() ([]byte, []int) {
	return file_sys_oper_log_model_gen_proto_rawDescGZIP(), []int{2}
}

func (x *SysOperLogListRequest) GetQuery() *SysOperLog {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *SysOperLogListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SysOperLogListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SysOperLogListRequest) GetOrderKey() string {
	if x != nil {
		return x.OrderKey
	}
	return ""
}

func (x *SysOperLogListRequest) GetOrderDesc() string {
	if x != nil {
		return x.OrderDesc
	}
	return ""
}

type SysOperLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32         `protobuf:"varint,101,opt,name=code,proto3" json:"code,omitempty"`
	Message string        `protobuf:"bytes,102,opt,name=message,proto3" json:"message,omitempty"`
	Total   int64         `protobuf:"varint,103,opt,name=total,proto3" json:"total,omitempty"`
	Data    []*SysOperLog `protobuf:"bytes,104,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *SysOperLogResponse) Reset() {
	*x = SysOperLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_oper_log_model_gen_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysOperLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysOperLogResponse) ProtoMessage() {}

func (x *SysOperLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sys_oper_log_model_gen_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysOperLogResponse.ProtoReflect.Descriptor instead.
func (*SysOperLogResponse) Descriptor() ([]byte, []int) {
	return file_sys_oper_log_model_gen_proto_rawDescGZIP(), []int{3}
}

func (x *SysOperLogResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SysOperLogResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SysOperLogResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *SysOperLogResponse) GetData() []*SysOperLog {
	if x != nil {
		return x.Data
	}
	return nil
}

type SysOperLogUpdateFieldsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []string    `protobuf:"bytes,101,rep,name=fields,proto3" json:"fields,omitempty"`
	Data   *SysOperLog `protobuf:"bytes,102,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SysOperLogUpdateFieldsRequest) Reset() {
	*x = SysOperLogUpdateFieldsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_oper_log_model_gen_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysOperLogUpdateFieldsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysOperLogUpdateFieldsRequest) ProtoMessage() {}

func (x *SysOperLogUpdateFieldsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sys_oper_log_model_gen_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysOperLogUpdateFieldsRequest.ProtoReflect.Descriptor instead.
func (*SysOperLogUpdateFieldsRequest) Descriptor() ([]byte, []int) {
	return file_sys_oper_log_model_gen_proto_rawDescGZIP(), []int{4}
}

func (x *SysOperLogUpdateFieldsRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *SysOperLogUpdateFieldsRequest) GetData() *SysOperLog {
	if x != nil {
		return x.Data
	}
	return nil
}

// 表模型主键，只能是整数或字符串，若生成错误，请检查数据库表主键是否正确。支持组合主键
type SysOperLogPrimarykey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperId int64 `protobuf:"varint,100,opt,name=operId,proto3" json:"operId,omitempty"`
}

func (x *SysOperLogPrimarykey) Reset() {
	*x = SysOperLogPrimarykey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_oper_log_model_gen_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysOperLogPrimarykey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysOperLogPrimarykey) ProtoMessage() {}

func (x *SysOperLogPrimarykey) ProtoReflect() protoreflect.Message {
	mi := &file_sys_oper_log_model_gen_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysOperLogPrimarykey.ProtoReflect.Descriptor instead.
func (*SysOperLogPrimarykey) Descriptor() ([]byte, []int) {
	return file_sys_oper_log_model_gen_proto_rawDescGZIP(), []int{5}
}

func (x *SysOperLogPrimarykey) GetOperId() int64 {
	if x != nil {
		return x.OperId
	}
	return 0
}

type SysOperLogBatchDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []*SysOperLogPrimarykey `protobuf:"bytes,101,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *SysOperLogBatchDeleteRequest) Reset() {
	*x = SysOperLogBatchDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_oper_log_model_gen_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysOperLogBatchDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysOperLogBatchDeleteRequest) ProtoMessage() {}

func (x *SysOperLogBatchDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sys_oper_log_model_gen_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysOperLogBatchDeleteRequest.ProtoReflect.Descriptor instead.
func (*SysOperLogBatchDeleteRequest) Descriptor() ([]byte, []int) {
	return file_sys_oper_log_model_gen_proto_rawDescGZIP(), []int{6}
}

func (x *SysOperLogBatchDeleteRequest) GetKeys() []*SysOperLogPrimarykey {
	if x != nil {
		return x.Keys
	}
	return nil
}

var File_sys_oper_log_model_gen_proto protoreflect.FileDescriptor

var file_sys_oper_log_model_gen_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x79, 0x73, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x03, 0x0a, 0x0a, 0x53, 0x79, 0x73,
	0x4f, 0x70, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x70, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x54, 0x79, 0x70, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x67, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x69, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f,
	0x70, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f,
	0x70, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x70, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x6b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x70, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x18, 0x6c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x70, 0x65, 0x72, 0x49, 0x70, 0x18, 0x6d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f,
	0x70, 0x65, 0x72, 0x49, 0x70, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x70, 0x65,
	0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65,
	0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x6f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x70,
	0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x6a, 0x73, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x70, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6a, 0x73, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x71, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x72, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x36, 0x0a, 0x08, 0x6f,
	0x70, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x73, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x54,
	0x69, 0x6d, 0x65, 0x22, 0xfe, 0x03, 0x0a, 0x10, 0x53, 0x79, 0x73, 0x4f, 0x70, 0x65, 0x72, 0x4c,
	0x6f, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x70, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x62, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x67, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x69, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6f, 0x70, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6f, 0x70, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x70, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x6b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x70, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x18,
	0x6c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x70, 0x65, 0x72, 0x49, 0x70, 0x18, 0x6d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6f, 0x70, 0x65, 0x72, 0x49, 0x70, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x70,
	0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70,
	0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x6f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f,
	0x70, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x6a, 0x73, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x70, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6a, 0x73,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x71, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x72, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x36, 0x0a, 0x08,
	0x6f, 0x70, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x73, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0xac, 0x01, 0x0a, 0x15, 0x53, 0x79, 0x73, 0x4f, 0x70, 0x65, 0x72,
	0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29,
	0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x79, 0x73, 0x4f, 0x70, 0x65, 0x72, 0x4c,
	0x6f, 0x67, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x4b, 0x65, 0x79, 0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65,
	0x73, 0x63, 0x18, 0x69, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44,
	0x65, 0x73, 0x63, 0x22, 0x81, 0x01, 0x0a, 0x12, 0x53, 0x79, 0x73, 0x4f, 0x70, 0x65, 0x72, 0x4c,
	0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x27,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x68, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x79, 0x73, 0x4f, 0x70, 0x65, 0x72, 0x4c, 0x6f,
	0x67, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x60, 0x0a, 0x1d, 0x53, 0x79, 0x73, 0x4f, 0x70,
	0x65, 0x72, 0x4c, 0x6f, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x12, 0x27, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x79, 0x73, 0x4f, 0x70, 0x65, 0x72,
	0x4c, 0x6f, 0x67, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2e, 0x0a, 0x14, 0x53, 0x79, 0x73,
	0x4f, 0x70, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x65, 0x72, 0x49, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x6f, 0x70, 0x65, 0x72, 0x49, 0x64, 0x22, 0x51, 0x0a, 0x1c, 0x53, 0x79, 0x73,
	0x4f, 0x70, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x04, 0x6b, 0x65, 0x79,
	0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x53, 0x79, 0x73, 0x4f, 0x70, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x50, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x42, 0x11, 0x5a, 0x0f,
	0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sys_oper_log_model_gen_proto_rawDescOnce sync.Once
	file_sys_oper_log_model_gen_proto_rawDescData = file_sys_oper_log_model_gen_proto_rawDesc
)

func file_sys_oper_log_model_gen_proto_rawDescGZIP() []byte {
	file_sys_oper_log_model_gen_proto_rawDescOnce.Do(func() {
		file_sys_oper_log_model_gen_proto_rawDescData = protoimpl.X.CompressGZIP(file_sys_oper_log_model_gen_proto_rawDescData)
	})
	return file_sys_oper_log_model_gen_proto_rawDescData
}

var file_sys_oper_log_model_gen_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_sys_oper_log_model_gen_proto_goTypes = []interface{}{
	(*SysOperLog)(nil),                    // 0: service.SysOperLog
	(*SysOperLogFilter)(nil),              // 1: service.SysOperLogFilter
	(*SysOperLogListRequest)(nil),         // 2: service.SysOperLogListRequest
	(*SysOperLogResponse)(nil),            // 3: service.SysOperLogResponse
	(*SysOperLogUpdateFieldsRequest)(nil), // 4: service.SysOperLogUpdateFieldsRequest
	(*SysOperLogPrimarykey)(nil),          // 5: service.SysOperLogPrimarykey
	(*SysOperLogBatchDeleteRequest)(nil),  // 6: service.SysOperLogBatchDeleteRequest
	(*timestamppb.Timestamp)(nil),         // 7: google.protobuf.Timestamp
}
var file_sys_oper_log_model_gen_proto_depIdxs = []int32{
	7, // 0: service.SysOperLog.operTime:type_name -> google.protobuf.Timestamp
	7, // 1: service.SysOperLogFilter.operTime:type_name -> google.protobuf.Timestamp
	0, // 2: service.SysOperLogListRequest.query:type_name -> service.SysOperLog
	0, // 3: service.SysOperLogResponse.data:type_name -> service.SysOperLog
	0, // 4: service.SysOperLogUpdateFieldsRequest.data:type_name -> service.SysOperLog
	5, // 5: service.SysOperLogBatchDeleteRequest.keys:type_name -> service.SysOperLogPrimarykey
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_sys_oper_log_model_gen_proto_init() }
func file_sys_oper_log_model_gen_proto_init() {
	if File_sys_oper_log_model_gen_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sys_oper_log_model_gen_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysOperLog); i {
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
		file_sys_oper_log_model_gen_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysOperLogFilter); i {
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
		file_sys_oper_log_model_gen_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysOperLogListRequest); i {
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
		file_sys_oper_log_model_gen_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysOperLogResponse); i {
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
		file_sys_oper_log_model_gen_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysOperLogUpdateFieldsRequest); i {
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
		file_sys_oper_log_model_gen_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysOperLogPrimarykey); i {
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
		file_sys_oper_log_model_gen_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysOperLogBatchDeleteRequest); i {
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
			RawDescriptor: file_sys_oper_log_model_gen_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sys_oper_log_model_gen_proto_goTypes,
		DependencyIndexes: file_sys_oper_log_model_gen_proto_depIdxs,
		MessageInfos:      file_sys_oper_log_model_gen_proto_msgTypes,
	}.Build()
	File_sys_oper_log_model_gen_proto = out.File
	file_sys_oper_log_model_gen_proto_rawDesc = nil
	file_sys_oper_log_model_gen_proto_goTypes = nil
	file_sys_oper_log_model_gen_proto_depIdxs = nil
}
