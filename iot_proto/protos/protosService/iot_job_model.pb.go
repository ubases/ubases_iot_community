// Code generated by protoc,2022-12-23 15:26:00. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: iot_job_model.proto

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
type IotJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:主键ID;comment:主键ID;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"productKey" gorm:"default:产品key;comment:产品key;size:14;"
	ProductKey string `protobuf:"bytes,101,opt,name=productKey,proto3" json:"productKey,omitempty"`
	// @inject_tag: json:"deviceId" gorm:"default:设备唯一ID（14位 1~9 A~Z随机）;comment:设备唯一ID（14位 1~9 A~Z随机）;size:36;"
	DeviceId string `protobuf:"bytes,102,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	// @inject_tag: json:"taskId" gorm:"default:任务编号;comment:任务编号;size:19;"
	TaskId int64 `protobuf:"varint,103,opt,name=taskId,proto3" json:"taskId,omitempty"`
	// @inject_tag: json:"taskType" gorm:"default:任务类型（1 倒计时任务 2 定时任务）;comment:任务类型（1 倒计时任务 2 定时任务）;"
	TaskType int32 `protobuf:"varint,104,opt,name=taskType,proto3" json:"taskType,omitempty"`
	// @inject_tag: json:"enabled" gorm:"default:定时器状态（=1启动 =2 禁用）;comment:定时器状态（=1启动 =2 禁用）;"
	Enabled int32 `protobuf:"varint,105,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// @inject_tag: json:"cron" gorm:"default:cron表达式，自动生成;comment:cron表达式，自动生成;size:255;"
	Cron string `protobuf:"bytes,106,opt,name=cron,proto3" json:"cron,omitempty"`
	// @inject_tag: json:"data" gorm:"default:任务数据;comment:任务数据;"
	Data string `protobuf:"bytes,107,opt,name=data,proto3" json:"data,omitempty"`
	// @inject_tag: json:"createdBy" gorm:"default:创建人;comment:创建人;size:20;"
	CreatedBy int64 `protobuf:"varint,108,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	// @inject_tag: json:"updatedBy" gorm:"default:修改人;comment:修改人;size:20;"
	UpdatedBy int64 `protobuf:"varint,109,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	// @inject_tag: json:"createdAt" gorm:"default:创建时间;comment:创建时间;"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,110,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// @inject_tag: json:"updatedAt" gorm:"default:修改时间;comment:修改时间;"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,111,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// @inject_tag: json:"endData" gorm:"default:结束时间-任务数据;comment:结束时间-任务数据;"
	EndData string `protobuf:"bytes,112,opt,name=endData,proto3" json:"endData,omitempty"`
	// @inject_tag: json:"endCron" gorm:"default:结束时间-cron表达式;comment:结束时间-cron表达式;size:255;"
	EndCron string `protobuf:"bytes,113,opt,name=endCron,proto3" json:"endCron,omitempty"`
	//定时器时区
	Timezone string `protobuf:"bytes,114,opt,name=timezone,proto3" json:"timezone,omitempty"`
	//任务所属区域服务器Id
	RegionServerId int64 `protobuf:"varint,115,opt,name=regionServerId,proto3" json:"regionServerId,omitempty"`
}

func (x *IotJob) Reset() {
	*x = IotJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iot_job_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IotJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IotJob) ProtoMessage() {}

func (x *IotJob) ProtoReflect() protoreflect.Message {
	mi := &file_iot_job_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IotJob.ProtoReflect.Descriptor instead.
func (*IotJob) Descriptor() ([]byte, []int) {
	return file_iot_job_model_proto_rawDescGZIP(), []int{0}
}

func (x *IotJob) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *IotJob) GetProductKey() string {
	if x != nil {
		return x.ProductKey
	}
	return ""
}

func (x *IotJob) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *IotJob) GetTaskId() int64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *IotJob) GetTaskType() int32 {
	if x != nil {
		return x.TaskType
	}
	return 0
}

func (x *IotJob) GetEnabled() int32 {
	if x != nil {
		return x.Enabled
	}
	return 0
}

func (x *IotJob) GetCron() string {
	if x != nil {
		return x.Cron
	}
	return ""
}

func (x *IotJob) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *IotJob) GetCreatedBy() int64 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *IotJob) GetUpdatedBy() int64 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

func (x *IotJob) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *IotJob) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *IotJob) GetEndData() string {
	if x != nil {
		return x.EndData
	}
	return ""
}

func (x *IotJob) GetEndCron() string {
	if x != nil {
		return x.EndCron
	}
	return ""
}

func (x *IotJob) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *IotJob) GetRegionServerId() int64 {
	if x != nil {
		return x.RegionServerId
	}
	return 0
}

type IotJobFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:主键ID;comment:主键ID;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"productKey" gorm:"default:产品key;comment:产品key;size:14;"
	ProductKey string `protobuf:"bytes,101,opt,name=productKey,proto3" json:"productKey,omitempty"`
	// @inject_tag: json:"deviceId" gorm:"default:设备唯一ID（14位 1~9 A~Z随机）;comment:设备唯一ID（14位 1~9 A~Z随机）;size:36;"
	DeviceId string `protobuf:"bytes,102,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	// @inject_tag: json:"taskId" gorm:"default:任务编号;comment:任务编号;size:19;"
	TaskId int64 `protobuf:"varint,103,opt,name=taskId,proto3" json:"taskId,omitempty"`
	// @inject_tag: json:"taskType" gorm:"default:任务类型（1 倒计时任务 2 定时任务）;comment:任务类型（1 倒计时任务 2 定时任务）;"
	TaskType int32 `protobuf:"varint,104,opt,name=taskType,proto3" json:"taskType,omitempty"`
	// @inject_tag: json:"enabled" gorm:"default:定时器状态（=1启动 =2 禁用）;comment:定时器状态（=1启动 =2 禁用）;"
	Enabled int32 `protobuf:"varint,105,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// @inject_tag: json:"cron" gorm:"default:cron表达式，自动生成;comment:cron表达式，自动生成;size:255;"
	Cron string `protobuf:"bytes,106,opt,name=cron,proto3" json:"cron,omitempty"`
	// @inject_tag: json:"data" gorm:"default:任务数据;comment:任务数据;"
	Data string `protobuf:"bytes,107,opt,name=data,proto3" json:"data,omitempty"`
	// @inject_tag: json:"createdBy" gorm:"default:创建人;comment:创建人;size:20;"
	CreatedBy int64 `protobuf:"varint,108,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	// @inject_tag: json:"updatedBy" gorm:"default:修改人;comment:修改人;size:20;"
	UpdatedBy int64 `protobuf:"varint,109,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	// @inject_tag: json:"createdAt" gorm:"default:创建时间;comment:创建时间;"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,110,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// @inject_tag: json:"updatedAt" gorm:"default:修改时间;comment:修改时间;"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,111,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// @inject_tag: json:"endData" gorm:"default:结束时间-任务数据;comment:结束时间-任务数据;"
	EndData string `protobuf:"bytes,112,opt,name=endData,proto3" json:"endData,omitempty"`
	// @inject_tag: json:"endCron" gorm:"default:结束时间-cron表达式;comment:结束时间-cron表达式;size:255;"
	EndCron string `protobuf:"bytes,113,opt,name=endCron,proto3" json:"endCron,omitempty"`
}

func (x *IotJobFilter) Reset() {
	*x = IotJobFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iot_job_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IotJobFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IotJobFilter) ProtoMessage() {}

func (x *IotJobFilter) ProtoReflect() protoreflect.Message {
	mi := &file_iot_job_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IotJobFilter.ProtoReflect.Descriptor instead.
func (*IotJobFilter) Descriptor() ([]byte, []int) {
	return file_iot_job_model_proto_rawDescGZIP(), []int{1}
}

func (x *IotJobFilter) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *IotJobFilter) GetProductKey() string {
	if x != nil {
		return x.ProductKey
	}
	return ""
}

func (x *IotJobFilter) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *IotJobFilter) GetTaskId() int64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *IotJobFilter) GetTaskType() int32 {
	if x != nil {
		return x.TaskType
	}
	return 0
}

func (x *IotJobFilter) GetEnabled() int32 {
	if x != nil {
		return x.Enabled
	}
	return 0
}

func (x *IotJobFilter) GetCron() string {
	if x != nil {
		return x.Cron
	}
	return ""
}

func (x *IotJobFilter) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *IotJobFilter) GetCreatedBy() int64 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *IotJobFilter) GetUpdatedBy() int64 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

func (x *IotJobFilter) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *IotJobFilter) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *IotJobFilter) GetEndData() string {
	if x != nil {
		return x.EndData
	}
	return ""
}

func (x *IotJobFilter) GetEndCron() string {
	if x != nil {
		return x.EndCron
	}
	return ""
}

type IotJobListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query     *IotJob `protobuf:"bytes,101,opt,name=query,proto3" json:"query,omitempty"`
	Page      int64   `protobuf:"varint,102,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int64   `protobuf:"varint,103,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	OrderKey  string  `protobuf:"bytes,104,opt,name=orderKey,proto3" json:"orderKey,omitempty"`
	OrderDesc string  `protobuf:"bytes,105,opt,name=orderDesc,proto3" json:"orderDesc,omitempty"`
	SearchKey string  `protobuf:"bytes,106,opt,name=searchKey,proto3" json:"searchKey,omitempty"`
}

func (x *IotJobListRequest) Reset() {
	*x = IotJobListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iot_job_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IotJobListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IotJobListRequest) ProtoMessage() {}

func (x *IotJobListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iot_job_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IotJobListRequest.ProtoReflect.Descriptor instead.
func (*IotJobListRequest) Descriptor() ([]byte, []int) {
	return file_iot_job_model_proto_rawDescGZIP(), []int{2}
}

func (x *IotJobListRequest) GetQuery() *IotJob {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *IotJobListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *IotJobListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *IotJobListRequest) GetOrderKey() string {
	if x != nil {
		return x.OrderKey
	}
	return ""
}

func (x *IotJobListRequest) GetOrderDesc() string {
	if x != nil {
		return x.OrderDesc
	}
	return ""
}

func (x *IotJobListRequest) GetSearchKey() string {
	if x != nil {
		return x.SearchKey
	}
	return ""
}

type IotJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32     `protobuf:"varint,101,opt,name=code,proto3" json:"code,omitempty"`
	Message string    `protobuf:"bytes,102,opt,name=message,proto3" json:"message,omitempty"`
	Total   int64     `protobuf:"varint,103,opt,name=total,proto3" json:"total,omitempty"`
	Data    []*IotJob `protobuf:"bytes,104,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *IotJobResponse) Reset() {
	*x = IotJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iot_job_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IotJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IotJobResponse) ProtoMessage() {}

func (x *IotJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iot_job_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IotJobResponse.ProtoReflect.Descriptor instead.
func (*IotJobResponse) Descriptor() ([]byte, []int) {
	return file_iot_job_model_proto_rawDescGZIP(), []int{3}
}

func (x *IotJobResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *IotJobResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *IotJobResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *IotJobResponse) GetData() []*IotJob {
	if x != nil {
		return x.Data
	}
	return nil
}

type IotJobUpdateFieldsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []string `protobuf:"bytes,101,rep,name=fields,proto3" json:"fields,omitempty"`
	Data   *IotJob  `protobuf:"bytes,102,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *IotJobUpdateFieldsRequest) Reset() {
	*x = IotJobUpdateFieldsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iot_job_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IotJobUpdateFieldsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IotJobUpdateFieldsRequest) ProtoMessage() {}

func (x *IotJobUpdateFieldsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iot_job_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IotJobUpdateFieldsRequest.ProtoReflect.Descriptor instead.
func (*IotJobUpdateFieldsRequest) Descriptor() ([]byte, []int) {
	return file_iot_job_model_proto_rawDescGZIP(), []int{4}
}

func (x *IotJobUpdateFieldsRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *IotJobUpdateFieldsRequest) GetData() *IotJob {
	if x != nil {
		return x.Data
	}
	return nil
}

// 表模型主键，只能是整数或字符串，若生成错误，请检查数据库表主键是否正确。支持组合主键
type IotJobPrimarykey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IotJobPrimarykey) Reset() {
	*x = IotJobPrimarykey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iot_job_model_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IotJobPrimarykey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IotJobPrimarykey) ProtoMessage() {}

func (x *IotJobPrimarykey) ProtoReflect() protoreflect.Message {
	mi := &file_iot_job_model_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IotJobPrimarykey.ProtoReflect.Descriptor instead.
func (*IotJobPrimarykey) Descriptor() ([]byte, []int) {
	return file_iot_job_model_proto_rawDescGZIP(), []int{5}
}

func (x *IotJobPrimarykey) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type IotJobBatchDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []*IotJobPrimarykey `protobuf:"bytes,101,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *IotJobBatchDeleteRequest) Reset() {
	*x = IotJobBatchDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iot_job_model_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IotJobBatchDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IotJobBatchDeleteRequest) ProtoMessage() {}

func (x *IotJobBatchDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iot_job_model_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IotJobBatchDeleteRequest.ProtoReflect.Descriptor instead.
func (*IotJobBatchDeleteRequest) Descriptor() ([]byte, []int) {
	return file_iot_job_model_proto_rawDescGZIP(), []int{6}
}

func (x *IotJobBatchDeleteRequest) GetKeys() []*IotJobPrimarykey {
	if x != nil {
		return x.Keys
	}
	return nil
}

type JobReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:编号;comment:编号;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *JobReq) Reset() {
	*x = JobReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iot_job_model_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobReq) ProtoMessage() {}

func (x *JobReq) ProtoReflect() protoreflect.Message {
	mi := &file_iot_job_model_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobReq.ProtoReflect.Descriptor instead.
func (*JobReq) Descriptor() ([]byte, []int) {
	return file_iot_job_model_proto_rawDescGZIP(), []int{7}
}

func (x *JobReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_iot_job_model_proto protoreflect.FileDescriptor

var file_iot_job_model_proto_rawDesc = []byte{
	0x0a, 0x13, 0x69, 0x6f, 0x74, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf2, 0x03, 0x0a, 0x06, 0x49, 0x6f, 0x74, 0x4a, 0x6f, 0x62, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64,
	0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x18, 0x68, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x69, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x72, 0x6f, 0x6e, 0x18, 0x6a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x72, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x6b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x6c, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x6d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x6f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x18, 0x70, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65,
	0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x43, 0x72, 0x6f,
	0x6e, 0x18, 0x71, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x43, 0x72, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x72, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x26, 0x0a, 0x0e,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x73,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x49, 0x64, 0x22, 0xb4, 0x03, 0x0a, 0x0c, 0x49, 0x6f, 0x74, 0x4a, 0x6f, 0x62, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4b, 0x65, 0x79, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x64, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x18, 0x67, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x73,
	0x6b, 0x54, 0x79, 0x70, 0x65, 0x18, 0x68, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74, 0x61, 0x73,
	0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x69, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x72, 0x6f, 0x6e, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x72, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x6b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x42, 0x79, 0x18, 0x6c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x18, 0x6d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x6e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x6f, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x70, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x43, 0x72, 0x6f, 0x6e, 0x18, 0x71, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x43, 0x72, 0x6f, 0x6e, 0x22, 0xc2, 0x01, 0x0a, 0x11,
	0x49, 0x6f, 0x74, 0x4a, 0x6f, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x25, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6f, 0x74, 0x4a, 0x6f,
	0x62, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x66, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x4b, 0x65, 0x79, 0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x73,
	0x63, 0x18, 0x69, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65,
	0x73, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x18,
	0x6a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79,
	0x22, 0x79, 0x0a, 0x0e, 0x49, 0x6f, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x68,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49,
	0x6f, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x58, 0x0a, 0x19, 0x49,
	0x6f, 0x74, 0x4a, 0x6f, 0x62, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6f, 0x74, 0x4a, 0x6f, 0x62, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x22, 0x0a, 0x10, 0x49, 0x6f, 0x74, 0x4a, 0x6f, 0x62, 0x50,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x49, 0x0a, 0x18, 0x49, 0x6f, 0x74,
	0x4a, 0x6f, 0x62, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x65, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6f,
	0x74, 0x4a, 0x6f, 0x62, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x52, 0x04,
	0x6b, 0x65, 0x79, 0x73, 0x22, 0x18, 0x0a, 0x06, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x42, 0x11,
	0x5a, 0x0f, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_iot_job_model_proto_rawDescOnce sync.Once
	file_iot_job_model_proto_rawDescData = file_iot_job_model_proto_rawDesc
)

func file_iot_job_model_proto_rawDescGZIP() []byte {
	file_iot_job_model_proto_rawDescOnce.Do(func() {
		file_iot_job_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_iot_job_model_proto_rawDescData)
	})
	return file_iot_job_model_proto_rawDescData
}

var file_iot_job_model_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_iot_job_model_proto_goTypes = []interface{}{
	(*IotJob)(nil),                    // 0: service.IotJob
	(*IotJobFilter)(nil),              // 1: service.IotJobFilter
	(*IotJobListRequest)(nil),         // 2: service.IotJobListRequest
	(*IotJobResponse)(nil),            // 3: service.IotJobResponse
	(*IotJobUpdateFieldsRequest)(nil), // 4: service.IotJobUpdateFieldsRequest
	(*IotJobPrimarykey)(nil),          // 5: service.IotJobPrimarykey
	(*IotJobBatchDeleteRequest)(nil),  // 6: service.IotJobBatchDeleteRequest
	(*JobReq)(nil),                    // 7: service.JobReq
	(*timestamppb.Timestamp)(nil),     // 8: google.protobuf.Timestamp
}
var file_iot_job_model_proto_depIdxs = []int32{
	8, // 0: service.IotJob.createdAt:type_name -> google.protobuf.Timestamp
	8, // 1: service.IotJob.updatedAt:type_name -> google.protobuf.Timestamp
	8, // 2: service.IotJobFilter.createdAt:type_name -> google.protobuf.Timestamp
	8, // 3: service.IotJobFilter.updatedAt:type_name -> google.protobuf.Timestamp
	0, // 4: service.IotJobListRequest.query:type_name -> service.IotJob
	0, // 5: service.IotJobResponse.data:type_name -> service.IotJob
	0, // 6: service.IotJobUpdateFieldsRequest.data:type_name -> service.IotJob
	5, // 7: service.IotJobBatchDeleteRequest.keys:type_name -> service.IotJobPrimarykey
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_iot_job_model_proto_init() }
func file_iot_job_model_proto_init() {
	if File_iot_job_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_iot_job_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IotJob); i {
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
		file_iot_job_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IotJobFilter); i {
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
		file_iot_job_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IotJobListRequest); i {
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
		file_iot_job_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IotJobResponse); i {
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
		file_iot_job_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IotJobUpdateFieldsRequest); i {
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
		file_iot_job_model_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IotJobPrimarykey); i {
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
		file_iot_job_model_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IotJobBatchDeleteRequest); i {
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
		file_iot_job_model_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobReq); i {
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
			RawDescriptor: file_iot_job_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_iot_job_model_proto_goTypes,
		DependencyIndexes: file_iot_job_model_proto_depIdxs,
		MessageInfos:      file_iot_job_model_proto_msgTypes,
	}.Build()
	File_iot_job_model_proto = out.File
	file_iot_job_model_proto_rawDesc = nil
	file_iot_job_model_proto_goTypes = nil
	file_iot_job_model_proto_depIdxs = nil
}
