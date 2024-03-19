// Code generated by protoc,2022-06-17 09:58:13. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: data_overview_hour_model.proto

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
type DataOverviewHour struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"dataTime" gorm:"default:时间;comment:时间;"
	DataTime *timestamppb.Timestamp `protobuf:"bytes,100,opt,name=dataTime,proto3" json:"dataTime,omitempty"`
	// @inject_tag: json:"tenantId" gorm:"default:租户ID，为空表示所有;comment:租户ID，为空表示所有;size:20;"
	TenantId string `protobuf:"bytes,101,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	// @inject_tag: json:"deviceActiveSum" gorm:"default:该月激活设备数;comment:该月激活设备数;size:19;"
	DeviceActiveSum int64 `protobuf:"varint,102,opt,name=deviceActiveSum,proto3" json:"deviceActiveSum,omitempty"`
	// @inject_tag: json:"deviceFaultSum" gorm:"default:该月设备故障数;comment:该月设备故障数;size:19;"
	DeviceFaultSum int64 `protobuf:"varint,103,opt,name=deviceFaultSum,proto3" json:"deviceFaultSum,omitempty"`
	// @inject_tag: json:"developerRegisterSum" gorm:"default:该月开发者注册数;comment:该月开发者注册数;size:19;"
	DeveloperRegisterSum int64 `protobuf:"varint,104,opt,name=developerRegisterSum,proto3" json:"developerRegisterSum,omitempty"`
	// @inject_tag: json:"userRegisterSum" gorm:"default:该月APP用户注册数;comment:该月APP用户注册数;size:19;"
	UserRegisterSum int64 `protobuf:"varint,105,opt,name=userRegisterSum,proto3" json:"userRegisterSum,omitempty"`
	// @inject_tag: json:"updatedAt" gorm:"default:;comment:;"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,106,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *DataOverviewHour) Reset() {
	*x = DataOverviewHour{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_overview_hour_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataOverviewHour) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataOverviewHour) ProtoMessage() {}

func (x *DataOverviewHour) ProtoReflect() protoreflect.Message {
	mi := &file_data_overview_hour_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataOverviewHour.ProtoReflect.Descriptor instead.
func (*DataOverviewHour) Descriptor() ([]byte, []int) {
	return file_data_overview_hour_model_proto_rawDescGZIP(), []int{0}
}

func (x *DataOverviewHour) GetDataTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DataTime
	}
	return nil
}

func (x *DataOverviewHour) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *DataOverviewHour) GetDeviceActiveSum() int64 {
	if x != nil {
		return x.DeviceActiveSum
	}
	return 0
}

func (x *DataOverviewHour) GetDeviceFaultSum() int64 {
	if x != nil {
		return x.DeviceFaultSum
	}
	return 0
}

func (x *DataOverviewHour) GetDeveloperRegisterSum() int64 {
	if x != nil {
		return x.DeveloperRegisterSum
	}
	return 0
}

func (x *DataOverviewHour) GetUserRegisterSum() int64 {
	if x != nil {
		return x.UserRegisterSum
	}
	return 0
}

func (x *DataOverviewHour) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type DataOverviewHourFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"dataTime" gorm:"default:时间;comment:时间;"
	DataTime *timestamppb.Timestamp `protobuf:"bytes,100,opt,name=dataTime,proto3" json:"dataTime,omitempty"`
	// @inject_tag: json:"tenantId" gorm:"default:租户ID，为空表示所有;comment:租户ID，为空表示所有;size:20;"
	TenantId string `protobuf:"bytes,101,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	// @inject_tag: json:"deviceActiveSum" gorm:"default:该月激活设备数;comment:该月激活设备数;size:19;"
	DeviceActiveSum int64 `protobuf:"varint,102,opt,name=deviceActiveSum,proto3" json:"deviceActiveSum,omitempty"`
	// @inject_tag: json:"deviceFaultSum" gorm:"default:该月设备故障数;comment:该月设备故障数;size:19;"
	DeviceFaultSum int64 `protobuf:"varint,103,opt,name=deviceFaultSum,proto3" json:"deviceFaultSum,omitempty"`
	// @inject_tag: json:"developerRegisterSum" gorm:"default:该月开发者注册数;comment:该月开发者注册数;size:19;"
	DeveloperRegisterSum int64 `protobuf:"varint,104,opt,name=developerRegisterSum,proto3" json:"developerRegisterSum,omitempty"`
	// @inject_tag: json:"userRegisterSum" gorm:"default:该月APP用户注册数;comment:该月APP用户注册数;size:19;"
	UserRegisterSum int64 `protobuf:"varint,105,opt,name=userRegisterSum,proto3" json:"userRegisterSum,omitempty"`
	// @inject_tag: json:"updatedAt" gorm:"default:;comment:;"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,106,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *DataOverviewHourFilter) Reset() {
	*x = DataOverviewHourFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_overview_hour_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataOverviewHourFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataOverviewHourFilter) ProtoMessage() {}

func (x *DataOverviewHourFilter) ProtoReflect() protoreflect.Message {
	mi := &file_data_overview_hour_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataOverviewHourFilter.ProtoReflect.Descriptor instead.
func (*DataOverviewHourFilter) Descriptor() ([]byte, []int) {
	return file_data_overview_hour_model_proto_rawDescGZIP(), []int{1}
}

func (x *DataOverviewHourFilter) GetDataTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DataTime
	}
	return nil
}

func (x *DataOverviewHourFilter) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *DataOverviewHourFilter) GetDeviceActiveSum() int64 {
	if x != nil {
		return x.DeviceActiveSum
	}
	return 0
}

func (x *DataOverviewHourFilter) GetDeviceFaultSum() int64 {
	if x != nil {
		return x.DeviceFaultSum
	}
	return 0
}

func (x *DataOverviewHourFilter) GetDeveloperRegisterSum() int64 {
	if x != nil {
		return x.DeveloperRegisterSum
	}
	return 0
}

func (x *DataOverviewHourFilter) GetUserRegisterSum() int64 {
	if x != nil {
		return x.UserRegisterSum
	}
	return 0
}

func (x *DataOverviewHourFilter) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type DataOverviewHourListFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantId  string                 `protobuf:"bytes,100,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	StartTime *timestamppb.Timestamp `protobuf:"bytes,101,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime   *timestamppb.Timestamp `protobuf:"bytes,102,opt,name=endTime,proto3" json:"endTime,omitempty"`
}

func (x *DataOverviewHourListFilter) Reset() {
	*x = DataOverviewHourListFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_overview_hour_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataOverviewHourListFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataOverviewHourListFilter) ProtoMessage() {}

func (x *DataOverviewHourListFilter) ProtoReflect() protoreflect.Message {
	mi := &file_data_overview_hour_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataOverviewHourListFilter.ProtoReflect.Descriptor instead.
func (*DataOverviewHourListFilter) Descriptor() ([]byte, []int) {
	return file_data_overview_hour_model_proto_rawDescGZIP(), []int{2}
}

func (x *DataOverviewHourListFilter) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *DataOverviewHourListFilter) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *DataOverviewHourListFilter) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

type DataOverviewHourListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query     *DataOverviewHourListFilter `protobuf:"bytes,101,opt,name=query,proto3" json:"query,omitempty"`
	Page      int64                       `protobuf:"varint,102,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int64                       `protobuf:"varint,103,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	OrderKey  string                      `protobuf:"bytes,104,opt,name=orderKey,proto3" json:"orderKey,omitempty"`
	OrderDesc string                      `protobuf:"bytes,105,opt,name=orderDesc,proto3" json:"orderDesc,omitempty"`
	SearchKey string                      `protobuf:"bytes,106,opt,name=searchKey,proto3" json:"searchKey,omitempty"`
}

func (x *DataOverviewHourListRequest) Reset() {
	*x = DataOverviewHourListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_overview_hour_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataOverviewHourListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataOverviewHourListRequest) ProtoMessage() {}

func (x *DataOverviewHourListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_data_overview_hour_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataOverviewHourListRequest.ProtoReflect.Descriptor instead.
func (*DataOverviewHourListRequest) Descriptor() ([]byte, []int) {
	return file_data_overview_hour_model_proto_rawDescGZIP(), []int{3}
}

func (x *DataOverviewHourListRequest) GetQuery() *DataOverviewHourListFilter {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *DataOverviewHourListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *DataOverviewHourListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *DataOverviewHourListRequest) GetOrderKey() string {
	if x != nil {
		return x.OrderKey
	}
	return ""
}

func (x *DataOverviewHourListRequest) GetOrderDesc() string {
	if x != nil {
		return x.OrderDesc
	}
	return ""
}

func (x *DataOverviewHourListRequest) GetSearchKey() string {
	if x != nil {
		return x.SearchKey
	}
	return ""
}

type DataOverviewHourResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32               `protobuf:"varint,101,opt,name=code,proto3" json:"code,omitempty"`
	Message string              `protobuf:"bytes,102,opt,name=message,proto3" json:"message,omitempty"`
	Total   int64               `protobuf:"varint,103,opt,name=total,proto3" json:"total,omitempty"`
	Data    []*DataOverviewHour `protobuf:"bytes,104,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *DataOverviewHourResponse) Reset() {
	*x = DataOverviewHourResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_overview_hour_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataOverviewHourResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataOverviewHourResponse) ProtoMessage() {}

func (x *DataOverviewHourResponse) ProtoReflect() protoreflect.Message {
	mi := &file_data_overview_hour_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataOverviewHourResponse.ProtoReflect.Descriptor instead.
func (*DataOverviewHourResponse) Descriptor() ([]byte, []int) {
	return file_data_overview_hour_model_proto_rawDescGZIP(), []int{4}
}

func (x *DataOverviewHourResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DataOverviewHourResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *DataOverviewHourResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *DataOverviewHourResponse) GetData() []*DataOverviewHour {
	if x != nil {
		return x.Data
	}
	return nil
}

type DataOverviewHourUpdateFieldsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []string          `protobuf:"bytes,101,rep,name=fields,proto3" json:"fields,omitempty"`
	Data   *DataOverviewHour `protobuf:"bytes,102,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DataOverviewHourUpdateFieldsRequest) Reset() {
	*x = DataOverviewHourUpdateFieldsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_overview_hour_model_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataOverviewHourUpdateFieldsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataOverviewHourUpdateFieldsRequest) ProtoMessage() {}

func (x *DataOverviewHourUpdateFieldsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_data_overview_hour_model_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataOverviewHourUpdateFieldsRequest.ProtoReflect.Descriptor instead.
func (*DataOverviewHourUpdateFieldsRequest) Descriptor() ([]byte, []int) {
	return file_data_overview_hour_model_proto_rawDescGZIP(), []int{5}
}

func (x *DataOverviewHourUpdateFieldsRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *DataOverviewHourUpdateFieldsRequest) GetData() *DataOverviewHour {
	if x != nil {
		return x.Data
	}
	return nil
}

// 表模型主键，只能是整数或字符串，若生成错误，请检查数据库表主键是否正确。支持组合主键
type DataOverviewHourPrimarykey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataTime *timestamppb.Timestamp `protobuf:"bytes,100,opt,name=dataTime,proto3" json:"dataTime,omitempty"`
	TenantId string                 `protobuf:"bytes,101,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
}

func (x *DataOverviewHourPrimarykey) Reset() {
	*x = DataOverviewHourPrimarykey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_overview_hour_model_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataOverviewHourPrimarykey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataOverviewHourPrimarykey) ProtoMessage() {}

func (x *DataOverviewHourPrimarykey) ProtoReflect() protoreflect.Message {
	mi := &file_data_overview_hour_model_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataOverviewHourPrimarykey.ProtoReflect.Descriptor instead.
func (*DataOverviewHourPrimarykey) Descriptor() ([]byte, []int) {
	return file_data_overview_hour_model_proto_rawDescGZIP(), []int{6}
}

func (x *DataOverviewHourPrimarykey) GetDataTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DataTime
	}
	return nil
}

func (x *DataOverviewHourPrimarykey) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

type DataOverviewHourBatchDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []*DataOverviewHourPrimarykey `protobuf:"bytes,101,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *DataOverviewHourBatchDeleteRequest) Reset() {
	*x = DataOverviewHourBatchDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_overview_hour_model_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataOverviewHourBatchDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataOverviewHourBatchDeleteRequest) ProtoMessage() {}

func (x *DataOverviewHourBatchDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_data_overview_hour_model_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataOverviewHourBatchDeleteRequest.ProtoReflect.Descriptor instead.
func (*DataOverviewHourBatchDeleteRequest) Descriptor() ([]byte, []int) {
	return file_data_overview_hour_model_proto_rawDescGZIP(), []int{7}
}

func (x *DataOverviewHourBatchDeleteRequest) GetKeys() []*DataOverviewHourPrimarykey {
	if x != nil {
		return x.Keys
	}
	return nil
}

var File_data_overview_hour_model_proto protoreflect.FileDescriptor

var file_data_overview_hour_model_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x5f,
	0x68, 0x6f, 0x75, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x02, 0x0a, 0x10, 0x44,
	0x61, 0x74, 0x61, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x48, 0x6f, 0x75, 0x72, 0x12,
	0x36, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x64,
	0x61, 0x74, 0x61, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x49, 0x64, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x53, 0x75, 0x6d, 0x18, 0x66, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53, 0x75, 0x6d, 0x12, 0x26, 0x0a,
	0x0e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x75, 0x6d, 0x18,
	0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x46, 0x61, 0x75,
	0x6c, 0x74, 0x53, 0x75, 0x6d, 0x12, 0x32, 0x0a, 0x14, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x75, 0x6d, 0x18, 0x68, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x14, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x75, 0x6d, 0x12, 0x28, 0x0a, 0x0f, 0x75, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x75, 0x6d, 0x18, 0x69, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x53, 0x75, 0x6d, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x6a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xd6, 0x02,
	0x0a, 0x16, 0x44, 0x61, 0x74, 0x61, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x48, 0x6f,
	0x75, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x65, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53, 0x75, 0x6d, 0x18,
	0x66, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x53, 0x75, 0x6d, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x46, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x75, 0x6d, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x75, 0x6d, 0x12, 0x32,
	0x0a, 0x14, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x53, 0x75, 0x6d, 0x18, 0x68, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x64, 0x65,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53,
	0x75, 0x6d, 0x12, 0x28, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x53, 0x75, 0x6d, 0x18, 0x69, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x75, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x75, 0x6d, 0x12, 0x38, 0x0a, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xa8, 0x01, 0x0a, 0x1a, 0x44, 0x61, 0x74, 0x61, 0x4f,
	0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x48, 0x6f, 0x75, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49,
	0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x38, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x65,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x65,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0xe0, 0x01, 0x0a, 0x1b, 0x44, 0x61, 0x74, 0x61, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69,
	0x65, 0x77, 0x48, 0x6f, 0x75, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x39, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4f,
	0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x48, 0x6f, 0x75, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x67, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x44, 0x65, 0x73, 0x63, 0x18, 0x69, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x4b, 0x65, 0x79, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x4b, 0x65, 0x79, 0x22, 0x8d, 0x01, 0x0a, 0x18, 0x44, 0x61, 0x74, 0x61, 0x4f, 0x76, 0x65,
	0x72, 0x76, 0x69, 0x65, 0x77, 0x48, 0x6f, 0x75, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x68, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x48, 0x6f, 0x75, 0x72, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x6c, 0x0a, 0x23, 0x44, 0x61, 0x74, 0x61, 0x4f, 0x76, 0x65, 0x72,
	0x76, 0x69, 0x65, 0x77, 0x48, 0x6f, 0x75, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x66, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x48, 0x6f, 0x75, 0x72, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x70, 0x0a, 0x1a, 0x44, 0x61, 0x74, 0x61, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69,
	0x65, 0x77, 0x48, 0x6f, 0x75, 0x72, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79,
	0x12, 0x36, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x64, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08,
	0x64, 0x61, 0x74, 0x61, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x49, 0x64, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x49, 0x64, 0x22, 0x5d, 0x0a, 0x22, 0x44, 0x61, 0x74, 0x61, 0x4f, 0x76, 0x65, 0x72,
	0x76, 0x69, 0x65, 0x77, 0x48, 0x6f, 0x75, 0x72, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x04, 0x6b, 0x65,
	0x79, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x48,
	0x6f, 0x75, 0x72, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x52, 0x04, 0x6b,
	0x65, 0x79, 0x73, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_overview_hour_model_proto_rawDescOnce sync.Once
	file_data_overview_hour_model_proto_rawDescData = file_data_overview_hour_model_proto_rawDesc
)

func file_data_overview_hour_model_proto_rawDescGZIP() []byte {
	file_data_overview_hour_model_proto_rawDescOnce.Do(func() {
		file_data_overview_hour_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_overview_hour_model_proto_rawDescData)
	})
	return file_data_overview_hour_model_proto_rawDescData
}

var file_data_overview_hour_model_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_data_overview_hour_model_proto_goTypes = []interface{}{
	(*DataOverviewHour)(nil),                    // 0: service.DataOverviewHour
	(*DataOverviewHourFilter)(nil),              // 1: service.DataOverviewHourFilter
	(*DataOverviewHourListFilter)(nil),          // 2: service.DataOverviewHourListFilter
	(*DataOverviewHourListRequest)(nil),         // 3: service.DataOverviewHourListRequest
	(*DataOverviewHourResponse)(nil),            // 4: service.DataOverviewHourResponse
	(*DataOverviewHourUpdateFieldsRequest)(nil), // 5: service.DataOverviewHourUpdateFieldsRequest
	(*DataOverviewHourPrimarykey)(nil),          // 6: service.DataOverviewHourPrimarykey
	(*DataOverviewHourBatchDeleteRequest)(nil),  // 7: service.DataOverviewHourBatchDeleteRequest
	(*timestamppb.Timestamp)(nil),               // 8: google.protobuf.Timestamp
}
var file_data_overview_hour_model_proto_depIdxs = []int32{
	8,  // 0: service.DataOverviewHour.dataTime:type_name -> google.protobuf.Timestamp
	8,  // 1: service.DataOverviewHour.updatedAt:type_name -> google.protobuf.Timestamp
	8,  // 2: service.DataOverviewHourFilter.dataTime:type_name -> google.protobuf.Timestamp
	8,  // 3: service.DataOverviewHourFilter.updatedAt:type_name -> google.protobuf.Timestamp
	8,  // 4: service.DataOverviewHourListFilter.startTime:type_name -> google.protobuf.Timestamp
	8,  // 5: service.DataOverviewHourListFilter.endTime:type_name -> google.protobuf.Timestamp
	2,  // 6: service.DataOverviewHourListRequest.query:type_name -> service.DataOverviewHourListFilter
	0,  // 7: service.DataOverviewHourResponse.data:type_name -> service.DataOverviewHour
	0,  // 8: service.DataOverviewHourUpdateFieldsRequest.data:type_name -> service.DataOverviewHour
	8,  // 9: service.DataOverviewHourPrimarykey.dataTime:type_name -> google.protobuf.Timestamp
	6,  // 10: service.DataOverviewHourBatchDeleteRequest.keys:type_name -> service.DataOverviewHourPrimarykey
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_data_overview_hour_model_proto_init() }
func file_data_overview_hour_model_proto_init() {
	if File_data_overview_hour_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_data_overview_hour_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataOverviewHour); i {
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
		file_data_overview_hour_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataOverviewHourFilter); i {
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
		file_data_overview_hour_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataOverviewHourListFilter); i {
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
		file_data_overview_hour_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataOverviewHourListRequest); i {
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
		file_data_overview_hour_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataOverviewHourResponse); i {
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
		file_data_overview_hour_model_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataOverviewHourUpdateFieldsRequest); i {
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
		file_data_overview_hour_model_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataOverviewHourPrimarykey); i {
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
		file_data_overview_hour_model_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataOverviewHourBatchDeleteRequest); i {
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
			RawDescriptor: file_data_overview_hour_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_data_overview_hour_model_proto_goTypes,
		DependencyIndexes: file_data_overview_hour_model_proto_depIdxs,
		MessageInfos:      file_data_overview_hour_model_proto_msgTypes,
	}.Build()
	File_data_overview_hour_model_proto = out.File
	file_data_overview_hour_model_proto_rawDesc = nil
	file_data_overview_hour_model_proto_goTypes = nil
	file_data_overview_hour_model_proto_depIdxs = nil
}
