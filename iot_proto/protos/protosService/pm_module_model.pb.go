// Code generated by sgen.exe,2022-04-24 10:36:38. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: pm_module_model.proto

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
type PmModule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:唯一主键;comment:唯一主键;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"moduleName" gorm:"default:模组芯片名称;comment:模组芯片名称;size:50;"
	ModuleName string `protobuf:"bytes,101,opt,name=moduleName,proto3" json:"moduleName,omitempty"`
	// @inject_tag: json:"firmwareType" gorm:"default:固件类型;comment:固件类型;"
	FirmwareType string `protobuf:"bytes,102,opt,name=firmwareType,proto3" json:"firmwareType,omitempty"`
	// @inject_tag: json:"firmwareFlag" gorm:"default:固件标识;comment:固件标识;size:50;"
	FirmwareFlag string `protobuf:"bytes,103,opt,name=firmwareFlag,proto3" json:"firmwareFlag,omitempty"`
	// @inject_tag: json:"firmwareId" gorm:"default:固件编号;comment:固件编号;size:19;"
	FirmwareId int64 `protobuf:"varint,104,opt,name=firmwareId,proto3" json:"firmwareId,omitempty"`
	// @inject_tag: json:"status" gorm:"default:状态（=1启用 =2 禁用）;comment:状态（=1启用 =2 禁用）;size:10;"
	Status int32 `protobuf:"varint,105,opt,name=status,proto3" json:"status,omitempty"`
	// @inject_tag: json:"imgUrl" gorm:"default:操作类型;comment:操作类型;size:255;"
	ImgUrl string `protobuf:"bytes,106,opt,name=imgUrl,proto3" json:"imgUrl,omitempty"`
	// @inject_tag: json:"fileUrl" gorm:"default:操作说明;comment:操作说明;size:255;"
	FileUrl string `protobuf:"bytes,107,opt,name=fileUrl,proto3" json:"fileUrl,omitempty"`
	// @inject_tag: json:"remark" gorm:"default:备注;comment:备注;size:255;"
	Remark string `protobuf:"bytes,108,opt,name=remark,proto3" json:"remark,omitempty"`
	// @inject_tag: json:"createdAt" gorm:"default:创建时间;comment:创建时间;"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,109,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// @inject_tag: json:"updatedBy" gorm:"default:修改人;comment:修改人;size:19;"
	UpdatedBy int64 `protobuf:"varint,110,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	// @inject_tag: json:"updatedAt" gorm:"default:修改时间;comment:修改时间;"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,111,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// @inject_tag: json:"deletedAt" gorm:"default:删除标识（0-正常 1-删除）;comment:删除标识（0-正常 1-删除）;"
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,112,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
	// @inject_tag: json:"moduleNameEn" gorm:"default:英文名称;comment:英文名称;size:50;"
	ModuleNameEn string `protobuf:"bytes,113,opt,name=moduleNameEn,proto3" json:"moduleNameEn,omitempty"`
	FirmwareUrl  string `protobuf:"bytes,114,opt,name=firmwareUrl,proto3" json:"firmwareUrl,omitempty"`
	FirmwareName string `protobuf:"bytes,115,opt,name=FirmwareName,proto3" json:"FirmwareName,omitempty"`
	// @inject_tag: json:"fileName" gorm:"default:文档文件名称;comment:文档文件名称;size:255;"
	FileName string `protobuf:"bytes,116,opt,name=fileName,proto3" json:"fileName,omitempty"`
	//默认版本号
	DefaultVersion string `protobuf:"bytes,117,opt,name=defaultVersion,proto3" json:"defaultVersion,omitempty"`
	//授权版本列表
	VersionList []*ModuleAuthFirmwareVersion `protobuf:"bytes,118,rep,name=versionList,proto3" json:"versionList,omitempty"`
	//固件Key
	FirmwareKey string `protobuf:"bytes,119,opt,name=FirmwareKey,proto3" json:"FirmwareKey,omitempty"`
	//固件版本数量
	VersionCount   int32  `protobuf:"varint,120,opt,name=versionCount,proto3" json:"versionCount,omitempty"`
	FirmwareNameEn string `protobuf:"bytes,121,opt,name=FirmwareNameEn,proto3" json:"FirmwareNameEn,omitempty"`
	RelationId     int64  `protobuf:"varint,122,opt,name=relationId,proto3" json:"relationId,omitempty"` //固件与产品关联Id
}

func (x *PmModule) Reset() {
	*x = PmModule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pm_module_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PmModule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PmModule) ProtoMessage() {}

func (x *PmModule) ProtoReflect() protoreflect.Message {
	mi := &file_pm_module_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PmModule.ProtoReflect.Descriptor instead.
func (*PmModule) Descriptor() ([]byte, []int) {
	return file_pm_module_model_proto_rawDescGZIP(), []int{0}
}

func (x *PmModule) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PmModule) GetModuleName() string {
	if x != nil {
		return x.ModuleName
	}
	return ""
}

func (x *PmModule) GetFirmwareType() string {
	if x != nil {
		return x.FirmwareType
	}
	return ""
}

func (x *PmModule) GetFirmwareFlag() string {
	if x != nil {
		return x.FirmwareFlag
	}
	return ""
}

func (x *PmModule) GetFirmwareId() int64 {
	if x != nil {
		return x.FirmwareId
	}
	return 0
}

func (x *PmModule) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *PmModule) GetImgUrl() string {
	if x != nil {
		return x.ImgUrl
	}
	return ""
}

func (x *PmModule) GetFileUrl() string {
	if x != nil {
		return x.FileUrl
	}
	return ""
}

func (x *PmModule) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *PmModule) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *PmModule) GetUpdatedBy() int64 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

func (x *PmModule) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *PmModule) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *PmModule) GetModuleNameEn() string {
	if x != nil {
		return x.ModuleNameEn
	}
	return ""
}

func (x *PmModule) GetFirmwareUrl() string {
	if x != nil {
		return x.FirmwareUrl
	}
	return ""
}

func (x *PmModule) GetFirmwareName() string {
	if x != nil {
		return x.FirmwareName
	}
	return ""
}

func (x *PmModule) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *PmModule) GetDefaultVersion() string {
	if x != nil {
		return x.DefaultVersion
	}
	return ""
}

func (x *PmModule) GetVersionList() []*ModuleAuthFirmwareVersion {
	if x != nil {
		return x.VersionList
	}
	return nil
}

func (x *PmModule) GetFirmwareKey() string {
	if x != nil {
		return x.FirmwareKey
	}
	return ""
}

func (x *PmModule) GetVersionCount() int32 {
	if x != nil {
		return x.VersionCount
	}
	return 0
}

func (x *PmModule) GetFirmwareNameEn() string {
	if x != nil {
		return x.FirmwareNameEn
	}
	return ""
}

func (x *PmModule) GetRelationId() int64 {
	if x != nil {
		return x.RelationId
	}
	return 0
}

//授权固件版本
type ModuleAuthFirmwareVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version   string `protobuf:"bytes,101,opt,name=version,proto3" json:"version,omitempty"`
	VersionId string `protobuf:"bytes,102,opt,name=versionId,proto3" json:"versionId,omitempty"`
}

func (x *ModuleAuthFirmwareVersion) Reset() {
	*x = ModuleAuthFirmwareVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pm_module_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModuleAuthFirmwareVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModuleAuthFirmwareVersion) ProtoMessage() {}

func (x *ModuleAuthFirmwareVersion) ProtoReflect() protoreflect.Message {
	mi := &file_pm_module_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModuleAuthFirmwareVersion.ProtoReflect.Descriptor instead.
func (*ModuleAuthFirmwareVersion) Descriptor() ([]byte, []int) {
	return file_pm_module_model_proto_rawDescGZIP(), []int{1}
}

func (x *ModuleAuthFirmwareVersion) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ModuleAuthFirmwareVersion) GetVersionId() string {
	if x != nil {
		return x.VersionId
	}
	return ""
}

type PmModuleFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:唯一主键;comment:唯一主键;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"moduleName" gorm:"default:模组芯片名称;comment:模组芯片名称;size:50;"
	ModuleName string `protobuf:"bytes,101,opt,name=moduleName,proto3" json:"moduleName,omitempty"`
	// @inject_tag: json:"firmwareType" gorm:"default:固件类型;comment:固件类型;"
	FirmwareType string `protobuf:"bytes,102,opt,name=firmwareType,proto3" json:"firmwareType,omitempty"`
	// @inject_tag: json:"firmwareFlag" gorm:"default:固件标识;comment:固件标识;size:50;"
	FirmwareFlag string `protobuf:"bytes,103,opt,name=firmwareFlag,proto3" json:"firmwareFlag,omitempty"`
	// @inject_tag: json:"firmwareId" gorm:"default:固件编号;comment:固件编号;size:19;"
	FirmwareId int64 `protobuf:"varint,104,opt,name=firmwareId,proto3" json:"firmwareId,omitempty"`
	// @inject_tag: json:"status" gorm:"default:状态（=1启用 =2 禁用）;comment:状态（=1启用 =2 禁用）;size:10;"
	Status int32 `protobuf:"varint,105,opt,name=status,proto3" json:"status,omitempty"`
	// @inject_tag: json:"imgUrl" gorm:"default:操作类型;comment:操作类型;size:255;"
	ImgUrl string `protobuf:"bytes,106,opt,name=imgUrl,proto3" json:"imgUrl,omitempty"`
	// @inject_tag: json:"fileUrl" gorm:"default:操作说明;comment:操作说明;size:255;"
	FileUrl string `protobuf:"bytes,107,opt,name=fileUrl,proto3" json:"fileUrl,omitempty"`
	// @inject_tag: json:"remark" gorm:"default:备注;comment:备注;size:255;"
	Remark string `protobuf:"bytes,108,opt,name=remark,proto3" json:"remark,omitempty"`
	// @inject_tag: json:"createdAt" gorm:"default:创建时间;comment:创建时间;"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,109,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// @inject_tag: json:"updatedBy" gorm:"default:修改人;comment:修改人;size:19;"
	UpdatedBy int64 `protobuf:"varint,110,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	// @inject_tag: json:"updatedAt" gorm:"default:修改时间;comment:修改时间;"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,111,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// @inject_tag: json:"deletedAt" gorm:"default:删除标识（0-正常 1-删除）;comment:删除标识（0-正常 1-删除）;"
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,112,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
	// @inject_tag: json:"nameEn" gorm:"default:英文名称;comment:英文名称;size:50;"
	ModuleNameEn string `protobuf:"bytes,113,opt,name=moduleNameEn,proto3" json:"moduleNameEn,omitempty"`
}

func (x *PmModuleFilter) Reset() {
	*x = PmModuleFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pm_module_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PmModuleFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PmModuleFilter) ProtoMessage() {}

func (x *PmModuleFilter) ProtoReflect() protoreflect.Message {
	mi := &file_pm_module_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PmModuleFilter.ProtoReflect.Descriptor instead.
func (*PmModuleFilter) Descriptor() ([]byte, []int) {
	return file_pm_module_model_proto_rawDescGZIP(), []int{2}
}

func (x *PmModuleFilter) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PmModuleFilter) GetModuleName() string {
	if x != nil {
		return x.ModuleName
	}
	return ""
}

func (x *PmModuleFilter) GetFirmwareType() string {
	if x != nil {
		return x.FirmwareType
	}
	return ""
}

func (x *PmModuleFilter) GetFirmwareFlag() string {
	if x != nil {
		return x.FirmwareFlag
	}
	return ""
}

func (x *PmModuleFilter) GetFirmwareId() int64 {
	if x != nil {
		return x.FirmwareId
	}
	return 0
}

func (x *PmModuleFilter) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *PmModuleFilter) GetImgUrl() string {
	if x != nil {
		return x.ImgUrl
	}
	return ""
}

func (x *PmModuleFilter) GetFileUrl() string {
	if x != nil {
		return x.FileUrl
	}
	return ""
}

func (x *PmModuleFilter) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *PmModuleFilter) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *PmModuleFilter) GetUpdatedBy() int64 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

func (x *PmModuleFilter) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *PmModuleFilter) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *PmModuleFilter) GetModuleNameEn() string {
	if x != nil {
		return x.ModuleNameEn
	}
	return ""
}

type PmModuleListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query     *PmModule `protobuf:"bytes,101,opt,name=query,proto3" json:"query,omitempty"`
	Page      int64     `protobuf:"varint,102,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int64     `protobuf:"varint,103,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	OrderKey  string    `protobuf:"bytes,104,opt,name=orderKey,proto3" json:"orderKey,omitempty"`
	OrderDesc string    `protobuf:"bytes,105,opt,name=orderDesc,proto3" json:"orderDesc,omitempty"`
}

func (x *PmModuleListRequest) Reset() {
	*x = PmModuleListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pm_module_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PmModuleListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PmModuleListRequest) ProtoMessage() {}

func (x *PmModuleListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pm_module_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PmModuleListRequest.ProtoReflect.Descriptor instead.
func (*PmModuleListRequest) Descriptor() ([]byte, []int) {
	return file_pm_module_model_proto_rawDescGZIP(), []int{3}
}

func (x *PmModuleListRequest) GetQuery() *PmModule {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *PmModuleListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PmModuleListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PmModuleListRequest) GetOrderKey() string {
	if x != nil {
		return x.OrderKey
	}
	return ""
}

func (x *PmModuleListRequest) GetOrderDesc() string {
	if x != nil {
		return x.OrderDesc
	}
	return ""
}

type PmModuleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32       `protobuf:"varint,101,opt,name=code,proto3" json:"code,omitempty"`
	Message string      `protobuf:"bytes,102,opt,name=message,proto3" json:"message,omitempty"`
	Total   int64       `protobuf:"varint,103,opt,name=total,proto3" json:"total,omitempty"`
	Data    []*PmModule `protobuf:"bytes,104,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *PmModuleResponse) Reset() {
	*x = PmModuleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pm_module_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PmModuleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PmModuleResponse) ProtoMessage() {}

func (x *PmModuleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pm_module_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PmModuleResponse.ProtoReflect.Descriptor instead.
func (*PmModuleResponse) Descriptor() ([]byte, []int) {
	return file_pm_module_model_proto_rawDescGZIP(), []int{4}
}

func (x *PmModuleResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PmModuleResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PmModuleResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PmModuleResponse) GetData() []*PmModule {
	if x != nil {
		return x.Data
	}
	return nil
}

type PmModuleUpdateFieldsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []string  `protobuf:"bytes,101,rep,name=fields,proto3" json:"fields,omitempty"`
	Data   *PmModule `protobuf:"bytes,102,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PmModuleUpdateFieldsRequest) Reset() {
	*x = PmModuleUpdateFieldsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pm_module_model_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PmModuleUpdateFieldsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PmModuleUpdateFieldsRequest) ProtoMessage() {}

func (x *PmModuleUpdateFieldsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pm_module_model_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PmModuleUpdateFieldsRequest.ProtoReflect.Descriptor instead.
func (*PmModuleUpdateFieldsRequest) Descriptor() ([]byte, []int) {
	return file_pm_module_model_proto_rawDescGZIP(), []int{5}
}

func (x *PmModuleUpdateFieldsRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *PmModuleUpdateFieldsRequest) GetData() *PmModule {
	if x != nil {
		return x.Data
	}
	return nil
}

// 表模型主键，只能是整数或字符串，若生成错误，请检查数据库表主键是否正确。支持组合主键
type PmModulePrimarykey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PmModulePrimarykey) Reset() {
	*x = PmModulePrimarykey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pm_module_model_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PmModulePrimarykey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PmModulePrimarykey) ProtoMessage() {}

func (x *PmModulePrimarykey) ProtoReflect() protoreflect.Message {
	mi := &file_pm_module_model_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PmModulePrimarykey.ProtoReflect.Descriptor instead.
func (*PmModulePrimarykey) Descriptor() ([]byte, []int) {
	return file_pm_module_model_proto_rawDescGZIP(), []int{6}
}

func (x *PmModulePrimarykey) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PmModuleBatchDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []*PmModulePrimarykey `protobuf:"bytes,101,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *PmModuleBatchDeleteRequest) Reset() {
	*x = PmModuleBatchDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pm_module_model_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PmModuleBatchDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PmModuleBatchDeleteRequest) ProtoMessage() {}

func (x *PmModuleBatchDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pm_module_model_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PmModuleBatchDeleteRequest.ProtoReflect.Descriptor instead.
func (*PmModuleBatchDeleteRequest) Descriptor() ([]byte, []int) {
	return file_pm_module_model_proto_rawDescGZIP(), []int{7}
}

func (x *PmModuleBatchDeleteRequest) GetKeys() []*PmModulePrimarykey {
	if x != nil {
		return x.Keys
	}
	return nil
}

var File_pm_module_model_proto protoreflect.FileDescriptor

var file_pm_module_model_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x6d, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd2, 0x06, 0x0a, 0x08, 0x50, 0x6d, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x65, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x66,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x46, 0x6c,
	0x61, 0x67, 0x18, 0x67, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61,
	0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61,
	0x72, 0x65, 0x49, 0x64, 0x18, 0x68, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x69, 0x72, 0x6d,
	0x77, 0x61, 0x72, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x69, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x69, 0x6d, 0x67, 0x55, 0x72, 0x6c, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x69, 0x6d, 0x67, 0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x55, 0x72,
	0x6c, 0x18, 0x6b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x55, 0x72, 0x6c,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x6c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x6d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18,
	0x6e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79,
	0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x6f, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x70, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x45, 0x6e, 0x18, 0x71, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x72, 0x6d,
	0x77, 0x61, 0x72, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x72, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66,
	0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x46, 0x69,
	0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x73, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x46, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x74, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x75, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x0b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x76, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x68, 0x46, 0x69, 0x72, 0x6d,
	0x77, 0x61, 0x72, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x69, 0x72, 0x6d,
	0x77, 0x61, 0x72, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x77, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x46,
	0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x78, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26,
	0x0a, 0x0e, 0x46, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x45, 0x6e,
	0x18, 0x79, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x46, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x18, 0x7a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x53, 0x0a, 0x19, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x41, 0x75, 0x74, 0x68, 0x46, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x65,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xfa, 0x03, 0x0a, 0x0e,
	0x50, 0x6d, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x65, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x66,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x46, 0x6c,
	0x61, 0x67, 0x18, 0x67, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61,
	0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61,
	0x72, 0x65, 0x49, 0x64, 0x18, 0x68, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x69, 0x72, 0x6d,
	0x77, 0x61, 0x72, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x69, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x69, 0x6d, 0x67, 0x55, 0x72, 0x6c, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x69, 0x6d, 0x67, 0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x55, 0x72,
	0x6c, 0x18, 0x6b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x55, 0x72, 0x6c,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x6c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x6d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18,
	0x6e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79,
	0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x6f, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x70, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x45, 0x6e, 0x18, 0x71, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x22, 0xa8, 0x01, 0x0a, 0x13, 0x50, 0x6d, 0x4d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x27, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6d, 0x4d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x4b, 0x65, 0x79, 0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65,
	0x73, 0x63, 0x18, 0x69, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44,
	0x65, 0x73, 0x63, 0x22, 0x7d, 0x0a, 0x10, 0x50, 0x6d, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x65, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x67,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x25, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x68, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x50, 0x6d, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x5c, 0x0a, 0x1b, 0x50, 0x6d, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x50, 0x6d, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x24, 0x0a, 0x12, 0x50, 0x6d, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4d, 0x0a, 0x1a, 0x50, 0x6d, 0x4d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x65, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6d, 0x4d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x52,
	0x04, 0x6b, 0x65, 0x79, 0x73, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pm_module_model_proto_rawDescOnce sync.Once
	file_pm_module_model_proto_rawDescData = file_pm_module_model_proto_rawDesc
)

func file_pm_module_model_proto_rawDescGZIP() []byte {
	file_pm_module_model_proto_rawDescOnce.Do(func() {
		file_pm_module_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_pm_module_model_proto_rawDescData)
	})
	return file_pm_module_model_proto_rawDescData
}

var file_pm_module_model_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pm_module_model_proto_goTypes = []interface{}{
	(*PmModule)(nil),                    // 0: service.PmModule
	(*ModuleAuthFirmwareVersion)(nil),   // 1: service.ModuleAuthFirmwareVersion
	(*PmModuleFilter)(nil),              // 2: service.PmModuleFilter
	(*PmModuleListRequest)(nil),         // 3: service.PmModuleListRequest
	(*PmModuleResponse)(nil),            // 4: service.PmModuleResponse
	(*PmModuleUpdateFieldsRequest)(nil), // 5: service.PmModuleUpdateFieldsRequest
	(*PmModulePrimarykey)(nil),          // 6: service.PmModulePrimarykey
	(*PmModuleBatchDeleteRequest)(nil),  // 7: service.PmModuleBatchDeleteRequest
	(*timestamppb.Timestamp)(nil),       // 8: google.protobuf.Timestamp
}
var file_pm_module_model_proto_depIdxs = []int32{
	8,  // 0: service.PmModule.createdAt:type_name -> google.protobuf.Timestamp
	8,  // 1: service.PmModule.updatedAt:type_name -> google.protobuf.Timestamp
	8,  // 2: service.PmModule.deletedAt:type_name -> google.protobuf.Timestamp
	1,  // 3: service.PmModule.versionList:type_name -> service.ModuleAuthFirmwareVersion
	8,  // 4: service.PmModuleFilter.createdAt:type_name -> google.protobuf.Timestamp
	8,  // 5: service.PmModuleFilter.updatedAt:type_name -> google.protobuf.Timestamp
	8,  // 6: service.PmModuleFilter.deletedAt:type_name -> google.protobuf.Timestamp
	0,  // 7: service.PmModuleListRequest.query:type_name -> service.PmModule
	0,  // 8: service.PmModuleResponse.data:type_name -> service.PmModule
	0,  // 9: service.PmModuleUpdateFieldsRequest.data:type_name -> service.PmModule
	6,  // 10: service.PmModuleBatchDeleteRequest.keys:type_name -> service.PmModulePrimarykey
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_pm_module_model_proto_init() }
func file_pm_module_model_proto_init() {
	if File_pm_module_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pm_module_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PmModule); i {
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
		file_pm_module_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModuleAuthFirmwareVersion); i {
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
		file_pm_module_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PmModuleFilter); i {
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
		file_pm_module_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PmModuleListRequest); i {
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
		file_pm_module_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PmModuleResponse); i {
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
		file_pm_module_model_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PmModuleUpdateFieldsRequest); i {
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
		file_pm_module_model_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PmModulePrimarykey); i {
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
		file_pm_module_model_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PmModuleBatchDeleteRequest); i {
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
			RawDescriptor: file_pm_module_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pm_module_model_proto_goTypes,
		DependencyIndexes: file_pm_module_model_proto_depIdxs,
		MessageInfos:      file_pm_module_model_proto_msgTypes,
	}.Build()
	File_pm_module_model_proto = out.File
	file_pm_module_model_proto_rawDesc = nil
	file_pm_module_model_proto_goTypes = nil
	file_pm_module_model_proto_depIdxs = nil
}
