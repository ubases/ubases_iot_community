// Code generated by protoc,2022-05-13 13:53:35. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: opm_ota_publish_log_model.gen.proto

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
type OpmOtaPublishLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:主键ID;comment:主键ID;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"productId" gorm:"default:产品编号;comment:产品编号;size:19;"
	ProductId int64 `protobuf:"varint,101,opt,name=productId,proto3" json:"productId,omitempty"`
	// @inject_tag: json:"firmwareId" gorm:"default:固件编号;comment:固件编号;size:19;"
	FirmwareId int64 `protobuf:"varint,102,opt,name=firmwareId,proto3" json:"firmwareId,omitempty"`
	// @inject_tag: json:"pkgId" gorm:"default:包id（t_opm_ota_pkg.id）;comment:包id（t_opm_ota_pkg.id）;size:19;"
	PkgId int64 `protobuf:"varint,103,opt,name=pkgId,proto3" json:"pkgId,omitempty"`
	// @inject_tag: json:"isGray" gorm:"default:是否灰度发布 0：否，1：是;comment:是否灰度发布 0：否，1：是;"
	IsGray int32 `protobuf:"varint,104,opt,name=isGray,proto3" json:"isGray,omitempty"`
	// @inject_tag: json:"grayType" gorm:"default:灰度类型[0:按比例灰度, 1:按数量灰度];comment:灰度类型[0:按比例灰度, 1:按数量灰度];"
	GrayType int32 `protobuf:"varint,105,opt,name=grayType,proto3" json:"grayType,omitempty"`
	// @inject_tag: json:"grayScale" gorm:"default:灰度比例;comment:灰度比例;size:10;"
	GrayScale int32 `protobuf:"varint,106,opt,name=grayScale,proto3" json:"grayScale,omitempty"`
	// @inject_tag: json:"type" gorm:"default:OTA类型[0:固件];comment:OTA类型[0:固件];size:10;"
	Type int32 `protobuf:"varint,107,opt,name=type,proto3" json:"type,omitempty"`
	// @inject_tag: json:"version" gorm:"default:OTA版本号;comment:OTA版本号;size:20;"
	Version string `protobuf:"bytes,108,opt,name=version,proto3" json:"version,omitempty"`
	// @inject_tag: json:"did" gorm:"default:设备唯一ID（14位 1~9 A~Z随机）;comment:设备唯一ID（14位 1~9 A~Z随机）;size:36;"
	Did string `protobuf:"bytes,109,opt,name=did,proto3" json:"did,omitempty"`
	// @inject_tag: json:"status" gorm:"default:状态[0:升级成功, 1:升级失败];comment:状态[0:升级成功, 1:升级失败];"
	Status int32 `protobuf:"varint,110,opt,name=status,proto3" json:"status,omitempty"`
	// @inject_tag: json:"deviceLog" gorm:"default:设备日志;comment:设备日志;"
	DeviceLog string `protobuf:"bytes,111,opt,name=deviceLog,proto3" json:"deviceLog,omitempty"`
	// @inject_tag: json:"createdBy" gorm:"default:创建人;comment:创建人;size:20;"
	CreatedBy int64 `protobuf:"varint,112,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	// @inject_tag: json:"updatedBy" gorm:"default:修改人;comment:修改人;size:20;"
	UpdatedBy int64 `protobuf:"varint,113,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	// @inject_tag: json:"createdAt" gorm:"default:创建时间;comment:创建时间;"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,114,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// @inject_tag: json:"updatedAt" gorm:"default:修改时间;comment:修改时间;"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,115,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// @inject_tag: json:"deletedAt" gorm:"default:删除时间;comment:删除时间;"
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,116,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
	// @inject_tag: json:"tenantId" gorm:"default:租户id（t_open_company.tenant_id）;comment:租户id（t_open_company.tenant_id）;size:6;"
	TenantId string `protobuf:"bytes,117,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
}

func (x *OpmOtaPublishLog) Reset() {
	*x = OpmOtaPublishLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opm_ota_publish_log_model_gen_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpmOtaPublishLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpmOtaPublishLog) ProtoMessage() {}

func (x *OpmOtaPublishLog) ProtoReflect() protoreflect.Message {
	mi := &file_opm_ota_publish_log_model_gen_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpmOtaPublishLog.ProtoReflect.Descriptor instead.
func (*OpmOtaPublishLog) Descriptor() ([]byte, []int) {
	return file_opm_ota_publish_log_model_gen_proto_rawDescGZIP(), []int{0}
}

func (x *OpmOtaPublishLog) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OpmOtaPublishLog) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *OpmOtaPublishLog) GetFirmwareId() int64 {
	if x != nil {
		return x.FirmwareId
	}
	return 0
}

func (x *OpmOtaPublishLog) GetPkgId() int64 {
	if x != nil {
		return x.PkgId
	}
	return 0
}

func (x *OpmOtaPublishLog) GetIsGray() int32 {
	if x != nil {
		return x.IsGray
	}
	return 0
}

func (x *OpmOtaPublishLog) GetGrayType() int32 {
	if x != nil {
		return x.GrayType
	}
	return 0
}

func (x *OpmOtaPublishLog) GetGrayScale() int32 {
	if x != nil {
		return x.GrayScale
	}
	return 0
}

func (x *OpmOtaPublishLog) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *OpmOtaPublishLog) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *OpmOtaPublishLog) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *OpmOtaPublishLog) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *OpmOtaPublishLog) GetDeviceLog() string {
	if x != nil {
		return x.DeviceLog
	}
	return ""
}

func (x *OpmOtaPublishLog) GetCreatedBy() int64 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *OpmOtaPublishLog) GetUpdatedBy() int64 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

func (x *OpmOtaPublishLog) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *OpmOtaPublishLog) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *OpmOtaPublishLog) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *OpmOtaPublishLog) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

type OpmOtaPublishLogFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:主键ID;comment:主键ID;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"productId" gorm:"default:产品编号;comment:产品编号;size:19;"
	ProductId int64 `protobuf:"varint,101,opt,name=productId,proto3" json:"productId,omitempty"`
	// @inject_tag: json:"firmwareId" gorm:"default:固件编号;comment:固件编号;size:19;"
	FirmwareId int64 `protobuf:"varint,102,opt,name=firmwareId,proto3" json:"firmwareId,omitempty"`
	// @inject_tag: json:"pkgId" gorm:"default:包id（t_opm_ota_pkg.id）;comment:包id（t_opm_ota_pkg.id）;size:19;"
	PkgId int64 `protobuf:"varint,103,opt,name=pkgId,proto3" json:"pkgId,omitempty"`
	// @inject_tag: json:"isGray" gorm:"default:是否灰度发布 0：否，1：是;comment:是否灰度发布 0：否，1：是;"
	IsGray int32 `protobuf:"varint,104,opt,name=isGray,proto3" json:"isGray,omitempty"`
	// @inject_tag: json:"grayType" gorm:"default:灰度类型[0:按比例灰度, 1:按数量灰度];comment:灰度类型[0:按比例灰度, 1:按数量灰度];"
	GrayType int32 `protobuf:"varint,105,opt,name=grayType,proto3" json:"grayType,omitempty"`
	// @inject_tag: json:"grayScale" gorm:"default:灰度比例;comment:灰度比例;size:10;"
	GrayScale int32 `protobuf:"varint,106,opt,name=grayScale,proto3" json:"grayScale,omitempty"`
	// @inject_tag: json:"type" gorm:"default:OTA类型[0:固件];comment:OTA类型[0:固件];size:10;"
	Type int32 `protobuf:"varint,107,opt,name=type,proto3" json:"type,omitempty"`
	// @inject_tag: json:"version" gorm:"default:OTA版本号;comment:OTA版本号;size:20;"
	Version string `protobuf:"bytes,108,opt,name=version,proto3" json:"version,omitempty"`
	// @inject_tag: json:"did" gorm:"default:设备唯一ID（14位 1~9 A~Z随机）;comment:设备唯一ID（14位 1~9 A~Z随机）;size:36;"
	Did string `protobuf:"bytes,109,opt,name=did,proto3" json:"did,omitempty"`
	// @inject_tag: json:"status" gorm:"default:状态[0:升级成功, 1:升级失败];comment:状态[0:升级成功, 1:升级失败];"
	Status int32 `protobuf:"varint,110,opt,name=status,proto3" json:"status,omitempty"`
	// @inject_tag: json:"deviceLog" gorm:"default:设备日志;comment:设备日志;"
	DeviceLog string `protobuf:"bytes,111,opt,name=deviceLog,proto3" json:"deviceLog,omitempty"`
	// @inject_tag: json:"createdBy" gorm:"default:创建人;comment:创建人;size:20;"
	CreatedBy int64 `protobuf:"varint,112,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	// @inject_tag: json:"updatedBy" gorm:"default:修改人;comment:修改人;size:20;"
	UpdatedBy int64 `protobuf:"varint,113,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	// @inject_tag: json:"createdAt" gorm:"default:创建时间;comment:创建时间;"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,114,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// @inject_tag: json:"updatedAt" gorm:"default:修改时间;comment:修改时间;"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,115,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// @inject_tag: json:"deletedAt" gorm:"default:删除时间;comment:删除时间;"
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,116,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
	// @inject_tag: json:"tenantId" gorm:"default:租户id（t_open_company.tenant_id）;comment:租户id（t_open_company.tenant_id）;size:6;"
	TenantId string `protobuf:"bytes,117,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
}

func (x *OpmOtaPublishLogFilter) Reset() {
	*x = OpmOtaPublishLogFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opm_ota_publish_log_model_gen_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpmOtaPublishLogFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpmOtaPublishLogFilter) ProtoMessage() {}

func (x *OpmOtaPublishLogFilter) ProtoReflect() protoreflect.Message {
	mi := &file_opm_ota_publish_log_model_gen_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpmOtaPublishLogFilter.ProtoReflect.Descriptor instead.
func (*OpmOtaPublishLogFilter) Descriptor() ([]byte, []int) {
	return file_opm_ota_publish_log_model_gen_proto_rawDescGZIP(), []int{1}
}

func (x *OpmOtaPublishLogFilter) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OpmOtaPublishLogFilter) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *OpmOtaPublishLogFilter) GetFirmwareId() int64 {
	if x != nil {
		return x.FirmwareId
	}
	return 0
}

func (x *OpmOtaPublishLogFilter) GetPkgId() int64 {
	if x != nil {
		return x.PkgId
	}
	return 0
}

func (x *OpmOtaPublishLogFilter) GetIsGray() int32 {
	if x != nil {
		return x.IsGray
	}
	return 0
}

func (x *OpmOtaPublishLogFilter) GetGrayType() int32 {
	if x != nil {
		return x.GrayType
	}
	return 0
}

func (x *OpmOtaPublishLogFilter) GetGrayScale() int32 {
	if x != nil {
		return x.GrayScale
	}
	return 0
}

func (x *OpmOtaPublishLogFilter) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *OpmOtaPublishLogFilter) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *OpmOtaPublishLogFilter) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *OpmOtaPublishLogFilter) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *OpmOtaPublishLogFilter) GetDeviceLog() string {
	if x != nil {
		return x.DeviceLog
	}
	return ""
}

func (x *OpmOtaPublishLogFilter) GetCreatedBy() int64 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *OpmOtaPublishLogFilter) GetUpdatedBy() int64 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

func (x *OpmOtaPublishLogFilter) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *OpmOtaPublishLogFilter) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *OpmOtaPublishLogFilter) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *OpmOtaPublishLogFilter) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

type OpmOtaPublishLogListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query     *OpmOtaPublishLog `protobuf:"bytes,101,opt,name=query,proto3" json:"query,omitempty"`
	Page      int64             `protobuf:"varint,102,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int64             `protobuf:"varint,103,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	OrderKey  string            `protobuf:"bytes,104,opt,name=orderKey,proto3" json:"orderKey,omitempty"`
	OrderDesc string            `protobuf:"bytes,105,opt,name=orderDesc,proto3" json:"orderDesc,omitempty"`
	SearchKey string            `protobuf:"bytes,106,opt,name=searchKey,proto3" json:"searchKey,omitempty"`
}

func (x *OpmOtaPublishLogListRequest) Reset() {
	*x = OpmOtaPublishLogListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opm_ota_publish_log_model_gen_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpmOtaPublishLogListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpmOtaPublishLogListRequest) ProtoMessage() {}

func (x *OpmOtaPublishLogListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_opm_ota_publish_log_model_gen_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpmOtaPublishLogListRequest.ProtoReflect.Descriptor instead.
func (*OpmOtaPublishLogListRequest) Descriptor() ([]byte, []int) {
	return file_opm_ota_publish_log_model_gen_proto_rawDescGZIP(), []int{2}
}

func (x *OpmOtaPublishLogListRequest) GetQuery() *OpmOtaPublishLog {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *OpmOtaPublishLogListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *OpmOtaPublishLogListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *OpmOtaPublishLogListRequest) GetOrderKey() string {
	if x != nil {
		return x.OrderKey
	}
	return ""
}

func (x *OpmOtaPublishLogListRequest) GetOrderDesc() string {
	if x != nil {
		return x.OrderDesc
	}
	return ""
}

func (x *OpmOtaPublishLogListRequest) GetSearchKey() string {
	if x != nil {
		return x.SearchKey
	}
	return ""
}

type OpmOtaPublishLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32               `protobuf:"varint,101,opt,name=code,proto3" json:"code,omitempty"`
	Message string              `protobuf:"bytes,102,opt,name=message,proto3" json:"message,omitempty"`
	Total   int64               `protobuf:"varint,103,opt,name=total,proto3" json:"total,omitempty"`
	Data    []*OpmOtaPublishLog `protobuf:"bytes,104,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *OpmOtaPublishLogResponse) Reset() {
	*x = OpmOtaPublishLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opm_ota_publish_log_model_gen_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpmOtaPublishLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpmOtaPublishLogResponse) ProtoMessage() {}

func (x *OpmOtaPublishLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_opm_ota_publish_log_model_gen_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpmOtaPublishLogResponse.ProtoReflect.Descriptor instead.
func (*OpmOtaPublishLogResponse) Descriptor() ([]byte, []int) {
	return file_opm_ota_publish_log_model_gen_proto_rawDescGZIP(), []int{3}
}

func (x *OpmOtaPublishLogResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *OpmOtaPublishLogResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *OpmOtaPublishLogResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *OpmOtaPublishLogResponse) GetData() []*OpmOtaPublishLog {
	if x != nil {
		return x.Data
	}
	return nil
}

type OpmOtaPublishLogUpdateFieldsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []string          `protobuf:"bytes,101,rep,name=fields,proto3" json:"fields,omitempty"`
	Data   *OpmOtaPublishLog `protobuf:"bytes,102,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OpmOtaPublishLogUpdateFieldsRequest) Reset() {
	*x = OpmOtaPublishLogUpdateFieldsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opm_ota_publish_log_model_gen_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpmOtaPublishLogUpdateFieldsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpmOtaPublishLogUpdateFieldsRequest) ProtoMessage() {}

func (x *OpmOtaPublishLogUpdateFieldsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_opm_ota_publish_log_model_gen_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpmOtaPublishLogUpdateFieldsRequest.ProtoReflect.Descriptor instead.
func (*OpmOtaPublishLogUpdateFieldsRequest) Descriptor() ([]byte, []int) {
	return file_opm_ota_publish_log_model_gen_proto_rawDescGZIP(), []int{4}
}

func (x *OpmOtaPublishLogUpdateFieldsRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *OpmOtaPublishLogUpdateFieldsRequest) GetData() *OpmOtaPublishLog {
	if x != nil {
		return x.Data
	}
	return nil
}

// 表模型主键，只能是整数或字符串，若生成错误，请检查数据库表主键是否正确。支持组合主键
type OpmOtaPublishLogPrimarykey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OpmOtaPublishLogPrimarykey) Reset() {
	*x = OpmOtaPublishLogPrimarykey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opm_ota_publish_log_model_gen_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpmOtaPublishLogPrimarykey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpmOtaPublishLogPrimarykey) ProtoMessage() {}

func (x *OpmOtaPublishLogPrimarykey) ProtoReflect() protoreflect.Message {
	mi := &file_opm_ota_publish_log_model_gen_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpmOtaPublishLogPrimarykey.ProtoReflect.Descriptor instead.
func (*OpmOtaPublishLogPrimarykey) Descriptor() ([]byte, []int) {
	return file_opm_ota_publish_log_model_gen_proto_rawDescGZIP(), []int{5}
}

func (x *OpmOtaPublishLogPrimarykey) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type OpmOtaPublishLogBatchDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []*OpmOtaPublishLogPrimarykey `protobuf:"bytes,101,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *OpmOtaPublishLogBatchDeleteRequest) Reset() {
	*x = OpmOtaPublishLogBatchDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opm_ota_publish_log_model_gen_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpmOtaPublishLogBatchDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpmOtaPublishLogBatchDeleteRequest) ProtoMessage() {}

func (x *OpmOtaPublishLogBatchDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_opm_ota_publish_log_model_gen_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpmOtaPublishLogBatchDeleteRequest.ProtoReflect.Descriptor instead.
func (*OpmOtaPublishLogBatchDeleteRequest) Descriptor() ([]byte, []int) {
	return file_opm_ota_publish_log_model_gen_proto_rawDescGZIP(), []int{6}
}

func (x *OpmOtaPublishLogBatchDeleteRequest) GetKeys() []*OpmOtaPublishLogPrimarykey {
	if x != nil {
		return x.Keys
	}
	return nil
}

var File_opm_ota_publish_log_model_gen_proto protoreflect.FileDescriptor

var file_opm_ota_publish_log_model_gen_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6f, 0x70, 0x6d, 0x5f, 0x6f, 0x74, 0x61, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x67, 0x65, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc4, 0x04, 0x0a, 0x10, 0x4f, 0x70, 0x6d, 0x4f, 0x74, 0x61, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x4c, 0x6f, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x64, 0x18, 0x65, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x49, 0x64,
	0x18, 0x66, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6b, 0x67, 0x49, 0x64, 0x18, 0x67, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x70, 0x6b, 0x67, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x47, 0x72,
	0x61, 0x79, 0x18, 0x68, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x73, 0x47, 0x72, 0x61, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x67, 0x72, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x69, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x67, 0x72, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x67, 0x72, 0x61, 0x79, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x67, 0x72, 0x61, 0x79, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x6b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x6c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x64, 0x18,
	0x6d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x18,
	0x6f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x70, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1c,
	0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x71, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x38, 0x0a, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x72, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x73, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x38, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x74, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x75, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x22, 0xca, 0x04, 0x0a, 0x16, 0x4f, 0x70, 0x6d, 0x4f, 0x74,
	0x61, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x65,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x49, 0x64, 0x18, 0x66, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x6b, 0x67, 0x49, 0x64, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x70, 0x6b, 0x67, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x47, 0x72, 0x61, 0x79, 0x18,
	0x68, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x73, 0x47, 0x72, 0x61, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x67, 0x72, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x69, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x67, 0x72, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x72, 0x61,
	0x79, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x67, 0x72,
	0x61, 0x79, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x6b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x6c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x64, 0x18, 0x6d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x18, 0x6f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x70, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x71, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x72, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x73, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a,
	0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x74, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x49, 0x64, 0x18, 0x75, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x49, 0x64, 0x22, 0xd6, 0x01, 0x0a, 0x1b, 0x4f, 0x70, 0x6d, 0x4f, 0x74, 0x61, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x65, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d,
	0x4f, 0x74, 0x61, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x6f, 0x67, 0x52, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x66, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79,
	0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79,
	0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x18, 0x69, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x18, 0x6a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x22, 0x8d, 0x01, 0x0a,
	0x18, 0x4f, 0x70, 0x6d, 0x4f, 0x74, 0x61, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x6f,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2d, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x68, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x4f, 0x74, 0x61, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x4c, 0x6f, 0x67, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x6c, 0x0a, 0x23,
	0x4f, 0x70, 0x6d, 0x4f, 0x74, 0x61, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x6f, 0x67,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x65, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x2d, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x4f, 0x74, 0x61, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x4c, 0x6f, 0x67, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2c, 0x0a, 0x1a, 0x4f, 0x70,
	0x6d, 0x4f, 0x74, 0x61, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x6f, 0x67, 0x50, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5d, 0x0a, 0x22, 0x4f, 0x70, 0x6d, 0x4f,
	0x74, 0x61, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x6f, 0x67, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37,
	0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x4f, 0x74, 0x61, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x4c, 0x6f, 0x67, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65,
	0x79, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_opm_ota_publish_log_model_gen_proto_rawDescOnce sync.Once
	file_opm_ota_publish_log_model_gen_proto_rawDescData = file_opm_ota_publish_log_model_gen_proto_rawDesc
)

func file_opm_ota_publish_log_model_gen_proto_rawDescGZIP() []byte {
	file_opm_ota_publish_log_model_gen_proto_rawDescOnce.Do(func() {
		file_opm_ota_publish_log_model_gen_proto_rawDescData = protoimpl.X.CompressGZIP(file_opm_ota_publish_log_model_gen_proto_rawDescData)
	})
	return file_opm_ota_publish_log_model_gen_proto_rawDescData
}

var file_opm_ota_publish_log_model_gen_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_opm_ota_publish_log_model_gen_proto_goTypes = []interface{}{
	(*OpmOtaPublishLog)(nil),                    // 0: service.OpmOtaPublishLog
	(*OpmOtaPublishLogFilter)(nil),              // 1: service.OpmOtaPublishLogFilter
	(*OpmOtaPublishLogListRequest)(nil),         // 2: service.OpmOtaPublishLogListRequest
	(*OpmOtaPublishLogResponse)(nil),            // 3: service.OpmOtaPublishLogResponse
	(*OpmOtaPublishLogUpdateFieldsRequest)(nil), // 4: service.OpmOtaPublishLogUpdateFieldsRequest
	(*OpmOtaPublishLogPrimarykey)(nil),          // 5: service.OpmOtaPublishLogPrimarykey
	(*OpmOtaPublishLogBatchDeleteRequest)(nil),  // 6: service.OpmOtaPublishLogBatchDeleteRequest
	(*timestamppb.Timestamp)(nil),               // 7: google.protobuf.Timestamp
}
var file_opm_ota_publish_log_model_gen_proto_depIdxs = []int32{
	7,  // 0: service.OpmOtaPublishLog.createdAt:type_name -> google.protobuf.Timestamp
	7,  // 1: service.OpmOtaPublishLog.updatedAt:type_name -> google.protobuf.Timestamp
	7,  // 2: service.OpmOtaPublishLog.deletedAt:type_name -> google.protobuf.Timestamp
	7,  // 3: service.OpmOtaPublishLogFilter.createdAt:type_name -> google.protobuf.Timestamp
	7,  // 4: service.OpmOtaPublishLogFilter.updatedAt:type_name -> google.protobuf.Timestamp
	7,  // 5: service.OpmOtaPublishLogFilter.deletedAt:type_name -> google.protobuf.Timestamp
	0,  // 6: service.OpmOtaPublishLogListRequest.query:type_name -> service.OpmOtaPublishLog
	0,  // 7: service.OpmOtaPublishLogResponse.data:type_name -> service.OpmOtaPublishLog
	0,  // 8: service.OpmOtaPublishLogUpdateFieldsRequest.data:type_name -> service.OpmOtaPublishLog
	5,  // 9: service.OpmOtaPublishLogBatchDeleteRequest.keys:type_name -> service.OpmOtaPublishLogPrimarykey
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_opm_ota_publish_log_model_gen_proto_init() }
func file_opm_ota_publish_log_model_gen_proto_init() {
	if File_opm_ota_publish_log_model_gen_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_opm_ota_publish_log_model_gen_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpmOtaPublishLog); i {
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
		file_opm_ota_publish_log_model_gen_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpmOtaPublishLogFilter); i {
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
		file_opm_ota_publish_log_model_gen_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpmOtaPublishLogListRequest); i {
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
		file_opm_ota_publish_log_model_gen_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpmOtaPublishLogResponse); i {
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
		file_opm_ota_publish_log_model_gen_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpmOtaPublishLogUpdateFieldsRequest); i {
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
		file_opm_ota_publish_log_model_gen_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpmOtaPublishLogPrimarykey); i {
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
		file_opm_ota_publish_log_model_gen_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpmOtaPublishLogBatchDeleteRequest); i {
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
			RawDescriptor: file_opm_ota_publish_log_model_gen_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_opm_ota_publish_log_model_gen_proto_goTypes,
		DependencyIndexes: file_opm_ota_publish_log_model_gen_proto_depIdxs,
		MessageInfos:      file_opm_ota_publish_log_model_gen_proto_msgTypes,
	}.Build()
	File_opm_ota_publish_log_model_gen_proto = out.File
	file_opm_ota_publish_log_model_gen_proto_rawDesc = nil
	file_opm_ota_publish_log_model_gen_proto_goTypes = nil
	file_opm_ota_publish_log_model_gen_proto_depIdxs = nil
}
