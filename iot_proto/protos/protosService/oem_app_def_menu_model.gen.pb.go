// Code generated by protoc,2022-05-31 16:29:04. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: oem_app_def_menu_model.gen.proto

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
type OemAppDefMenu struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:;comment:;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"name" gorm:"default:菜单名称;comment:菜单名称;size:32;"
	Name string `protobuf:"bytes,101,opt,name=name,proto3" json:"name,omitempty"`
	// @inject_tag: json:"menuKey" gorm:"default:菜单标识;comment:菜单标识;size:32;"
	MenuKey string `protobuf:"bytes,102,opt,name=menuKey,proto3" json:"menuKey,omitempty"`
	// @inject_tag: json:"position" gorm:"default:菜单位置，左到右，1开始;comment:菜单位置，左到右，1开始;size:10;"
	Position int32 `protobuf:"varint,103,opt,name=position,proto3" json:"position,omitempty"`
	// @inject_tag: json:"defImage" gorm:"default:默认图标;comment:默认图标;size:255;"
	DefImage string `protobuf:"bytes,104,opt,name=defImage,proto3" json:"defImage,omitempty"`
	// @inject_tag: json:"selImage" gorm:"default:选中图标;comment:选中图标;size:255;"
	SelImage string `protobuf:"bytes,105,opt,name=selImage,proto3" json:"selImage,omitempty"`
	// @inject_tag: json:"required" gorm:"default:是否必须 1 必须,  2 非必须;comment:是否必须 1 必须,  2 非必须;"
	Required int32 `protobuf:"varint,106,opt,name=required,proto3" json:"required,omitempty"`
	// @inject_tag: json:"status" gorm:"default:1启用2禁用;comment:1启用2禁用;"
	Status int32 `protobuf:"varint,107,opt,name=status,proto3" json:"status,omitempty"`
	// @inject_tag: json:"createdBy" gorm:"default:;comment:;size:19;"
	CreatedBy int64 `protobuf:"varint,108,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	// @inject_tag: json:"updatedBy" gorm:"default:;comment:;size:19;"
	UpdatedBy int64 `protobuf:"varint,109,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	// @inject_tag: json:"createdAt" gorm:"default:;comment:;"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,110,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// @inject_tag: json:"updatedAt" gorm:"default:;comment:;"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,111,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// @inject_tag: json:"deletedAt" gorm:"default:;comment:;"
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,112,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *OemAppDefMenu) Reset() {
	*x = OemAppDefMenu{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oem_app_def_menu_model_gen_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OemAppDefMenu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OemAppDefMenu) ProtoMessage() {}

func (x *OemAppDefMenu) ProtoReflect() protoreflect.Message {
	mi := &file_oem_app_def_menu_model_gen_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OemAppDefMenu.ProtoReflect.Descriptor instead.
func (*OemAppDefMenu) Descriptor() ([]byte, []int) {
	return file_oem_app_def_menu_model_gen_proto_rawDescGZIP(), []int{0}
}

func (x *OemAppDefMenu) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OemAppDefMenu) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OemAppDefMenu) GetMenuKey() string {
	if x != nil {
		return x.MenuKey
	}
	return ""
}

func (x *OemAppDefMenu) GetPosition() int32 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *OemAppDefMenu) GetDefImage() string {
	if x != nil {
		return x.DefImage
	}
	return ""
}

func (x *OemAppDefMenu) GetSelImage() string {
	if x != nil {
		return x.SelImage
	}
	return ""
}

func (x *OemAppDefMenu) GetRequired() int32 {
	if x != nil {
		return x.Required
	}
	return 0
}

func (x *OemAppDefMenu) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *OemAppDefMenu) GetCreatedBy() int64 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *OemAppDefMenu) GetUpdatedBy() int64 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

func (x *OemAppDefMenu) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *OemAppDefMenu) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *OemAppDefMenu) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type OemAppDefMenuFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:;comment:;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"name" gorm:"default:菜单名称;comment:菜单名称;size:32;"
	Name string `protobuf:"bytes,101,opt,name=name,proto3" json:"name,omitempty"`
	// @inject_tag: json:"menuKey" gorm:"default:菜单标识;comment:菜单标识;size:32;"
	MenuKey string `protobuf:"bytes,102,opt,name=menuKey,proto3" json:"menuKey,omitempty"`
	// @inject_tag: json:"position" gorm:"default:菜单位置，左到右，1开始;comment:菜单位置，左到右，1开始;size:10;"
	Position int32 `protobuf:"varint,103,opt,name=position,proto3" json:"position,omitempty"`
	// @inject_tag: json:"defImage" gorm:"default:默认图标;comment:默认图标;size:255;"
	DefImage string `protobuf:"bytes,104,opt,name=defImage,proto3" json:"defImage,omitempty"`
	// @inject_tag: json:"selImage" gorm:"default:选中图标;comment:选中图标;size:255;"
	SelImage string `protobuf:"bytes,105,opt,name=selImage,proto3" json:"selImage,omitempty"`
	// @inject_tag: json:"required" gorm:"default:是否必须 1 必须,  2 非必须;comment:是否必须 1 必须,  2 非必须;"
	Required int32 `protobuf:"varint,106,opt,name=required,proto3" json:"required,omitempty"`
	// @inject_tag: json:"status" gorm:"default:1启用2禁用;comment:1启用2禁用;"
	Status int32 `protobuf:"varint,107,opt,name=status,proto3" json:"status,omitempty"`
	// @inject_tag: json:"createdBy" gorm:"default:;comment:;size:19;"
	CreatedBy int64 `protobuf:"varint,108,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	// @inject_tag: json:"updatedBy" gorm:"default:;comment:;size:19;"
	UpdatedBy int64 `protobuf:"varint,109,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	// @inject_tag: json:"createdAt" gorm:"default:;comment:;"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,110,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// @inject_tag: json:"updatedAt" gorm:"default:;comment:;"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,111,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// @inject_tag: json:"deletedAt" gorm:"default:;comment:;"
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,112,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *OemAppDefMenuFilter) Reset() {
	*x = OemAppDefMenuFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oem_app_def_menu_model_gen_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OemAppDefMenuFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OemAppDefMenuFilter) ProtoMessage() {}

func (x *OemAppDefMenuFilter) ProtoReflect() protoreflect.Message {
	mi := &file_oem_app_def_menu_model_gen_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OemAppDefMenuFilter.ProtoReflect.Descriptor instead.
func (*OemAppDefMenuFilter) Descriptor() ([]byte, []int) {
	return file_oem_app_def_menu_model_gen_proto_rawDescGZIP(), []int{1}
}

func (x *OemAppDefMenuFilter) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OemAppDefMenuFilter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OemAppDefMenuFilter) GetMenuKey() string {
	if x != nil {
		return x.MenuKey
	}
	return ""
}

func (x *OemAppDefMenuFilter) GetPosition() int32 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *OemAppDefMenuFilter) GetDefImage() string {
	if x != nil {
		return x.DefImage
	}
	return ""
}

func (x *OemAppDefMenuFilter) GetSelImage() string {
	if x != nil {
		return x.SelImage
	}
	return ""
}

func (x *OemAppDefMenuFilter) GetRequired() int32 {
	if x != nil {
		return x.Required
	}
	return 0
}

func (x *OemAppDefMenuFilter) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *OemAppDefMenuFilter) GetCreatedBy() int64 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *OemAppDefMenuFilter) GetUpdatedBy() int64 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

func (x *OemAppDefMenuFilter) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *OemAppDefMenuFilter) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *OemAppDefMenuFilter) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type OemAppDefMenuListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query     *OemAppDefMenu `protobuf:"bytes,101,opt,name=query,proto3" json:"query,omitempty"`
	Page      int64          `protobuf:"varint,102,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int64          `protobuf:"varint,103,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	OrderKey  string         `protobuf:"bytes,104,opt,name=orderKey,proto3" json:"orderKey,omitempty"`
	OrderDesc string         `protobuf:"bytes,105,opt,name=orderDesc,proto3" json:"orderDesc,omitempty"`
	SearchKey string         `protobuf:"bytes,106,opt,name=searchKey,proto3" json:"searchKey,omitempty"`
}

func (x *OemAppDefMenuListRequest) Reset() {
	*x = OemAppDefMenuListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oem_app_def_menu_model_gen_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OemAppDefMenuListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OemAppDefMenuListRequest) ProtoMessage() {}

func (x *OemAppDefMenuListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oem_app_def_menu_model_gen_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OemAppDefMenuListRequest.ProtoReflect.Descriptor instead.
func (*OemAppDefMenuListRequest) Descriptor() ([]byte, []int) {
	return file_oem_app_def_menu_model_gen_proto_rawDescGZIP(), []int{2}
}

func (x *OemAppDefMenuListRequest) GetQuery() *OemAppDefMenu {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *OemAppDefMenuListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *OemAppDefMenuListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *OemAppDefMenuListRequest) GetOrderKey() string {
	if x != nil {
		return x.OrderKey
	}
	return ""
}

func (x *OemAppDefMenuListRequest) GetOrderDesc() string {
	if x != nil {
		return x.OrderDesc
	}
	return ""
}

func (x *OemAppDefMenuListRequest) GetSearchKey() string {
	if x != nil {
		return x.SearchKey
	}
	return ""
}

type OemAppDefMenuResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32            `protobuf:"varint,101,opt,name=code,proto3" json:"code,omitempty"`
	Message string           `protobuf:"bytes,102,opt,name=message,proto3" json:"message,omitempty"`
	Total   int64            `protobuf:"varint,103,opt,name=total,proto3" json:"total,omitempty"`
	Data    []*OemAppDefMenu `protobuf:"bytes,104,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *OemAppDefMenuResponse) Reset() {
	*x = OemAppDefMenuResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oem_app_def_menu_model_gen_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OemAppDefMenuResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OemAppDefMenuResponse) ProtoMessage() {}

func (x *OemAppDefMenuResponse) ProtoReflect() protoreflect.Message {
	mi := &file_oem_app_def_menu_model_gen_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OemAppDefMenuResponse.ProtoReflect.Descriptor instead.
func (*OemAppDefMenuResponse) Descriptor() ([]byte, []int) {
	return file_oem_app_def_menu_model_gen_proto_rawDescGZIP(), []int{3}
}

func (x *OemAppDefMenuResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *OemAppDefMenuResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *OemAppDefMenuResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *OemAppDefMenuResponse) GetData() []*OemAppDefMenu {
	if x != nil {
		return x.Data
	}
	return nil
}

type OemAppDefMenuUpdateFieldsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []string       `protobuf:"bytes,101,rep,name=fields,proto3" json:"fields,omitempty"`
	Data   *OemAppDefMenu `protobuf:"bytes,102,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OemAppDefMenuUpdateFieldsRequest) Reset() {
	*x = OemAppDefMenuUpdateFieldsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oem_app_def_menu_model_gen_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OemAppDefMenuUpdateFieldsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OemAppDefMenuUpdateFieldsRequest) ProtoMessage() {}

func (x *OemAppDefMenuUpdateFieldsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oem_app_def_menu_model_gen_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OemAppDefMenuUpdateFieldsRequest.ProtoReflect.Descriptor instead.
func (*OemAppDefMenuUpdateFieldsRequest) Descriptor() ([]byte, []int) {
	return file_oem_app_def_menu_model_gen_proto_rawDescGZIP(), []int{4}
}

func (x *OemAppDefMenuUpdateFieldsRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *OemAppDefMenuUpdateFieldsRequest) GetData() *OemAppDefMenu {
	if x != nil {
		return x.Data
	}
	return nil
}

// 表模型主键，只能是整数或字符串，若生成错误，请检查数据库表主键是否正确。支持组合主键
type OemAppDefMenuPrimarykey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OemAppDefMenuPrimarykey) Reset() {
	*x = OemAppDefMenuPrimarykey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oem_app_def_menu_model_gen_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OemAppDefMenuPrimarykey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OemAppDefMenuPrimarykey) ProtoMessage() {}

func (x *OemAppDefMenuPrimarykey) ProtoReflect() protoreflect.Message {
	mi := &file_oem_app_def_menu_model_gen_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OemAppDefMenuPrimarykey.ProtoReflect.Descriptor instead.
func (*OemAppDefMenuPrimarykey) Descriptor() ([]byte, []int) {
	return file_oem_app_def_menu_model_gen_proto_rawDescGZIP(), []int{5}
}

func (x *OemAppDefMenuPrimarykey) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type OemAppDefMenuBatchDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []*OemAppDefMenuPrimarykey `protobuf:"bytes,101,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *OemAppDefMenuBatchDeleteRequest) Reset() {
	*x = OemAppDefMenuBatchDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oem_app_def_menu_model_gen_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OemAppDefMenuBatchDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OemAppDefMenuBatchDeleteRequest) ProtoMessage() {}

func (x *OemAppDefMenuBatchDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oem_app_def_menu_model_gen_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OemAppDefMenuBatchDeleteRequest.ProtoReflect.Descriptor instead.
func (*OemAppDefMenuBatchDeleteRequest) Descriptor() ([]byte, []int) {
	return file_oem_app_def_menu_model_gen_proto_rawDescGZIP(), []int{6}
}

func (x *OemAppDefMenuBatchDeleteRequest) GetKeys() []*OemAppDefMenuPrimarykey {
	if x != nil {
		return x.Keys
	}
	return nil
}

var File_oem_app_def_menu_model_gen_proto protoreflect.FileDescriptor

var file_oem_app_def_menu_model_gen_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6f, 0x65, 0x6d, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x6d, 0x65,
	0x6e, 0x75, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x03, 0x0a,
	0x0d, 0x4f, 0x65, 0x6d, 0x41, 0x70, 0x70, 0x44, 0x65, 0x66, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6e, 0x75, 0x4b, 0x65, 0x79, 0x18, 0x66, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x6e, 0x75, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x67, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x66, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x66, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x69, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x6a, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x6b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x18, 0x6c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18,
	0x6d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79,
	0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x6e, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x6f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x70, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xc5,
	0x03, 0x0a, 0x13, 0x4f, 0x65, 0x6d, 0x41, 0x70, 0x70, 0x44, 0x65, 0x66, 0x4d, 0x65, 0x6e, 0x75,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x65,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x6e, 0x75, 0x4b, 0x65, 0x79, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x6e,
	0x75, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x67, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x66, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x68, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x66, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x65, 0x6c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x69, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x65, 0x6c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x6b,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09,
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
	0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x70, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xd0, 0x01, 0x0a, 0x18, 0x4f, 0x65, 0x6d, 0x41, 0x70,
	0x70, 0x44, 0x65, 0x66, 0x4d, 0x65, 0x6e, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x65, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x65, 0x6d,
	0x41, 0x70, 0x70, 0x44, 0x65, 0x66, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x18, 0x68, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a,
	0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x18, 0x69, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x22, 0x87, 0x01, 0x0a, 0x15, 0x4f, 0x65,
	0x6d, 0x41, 0x70, 0x70, 0x44, 0x65, 0x66, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x68, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x4f, 0x65, 0x6d, 0x41, 0x70, 0x70, 0x44, 0x65, 0x66, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x66, 0x0a, 0x20, 0x4f, 0x65, 0x6d, 0x41, 0x70, 0x70, 0x44, 0x65, 0x66,
	0x4d, 0x65, 0x6e, 0x75, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12,
	0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x65, 0x6d, 0x41, 0x70, 0x70, 0x44, 0x65,
	0x66, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x29, 0x0a, 0x17, 0x4f,
	0x65, 0x6d, 0x41, 0x70, 0x70, 0x44, 0x65, 0x66, 0x4d, 0x65, 0x6e, 0x75, 0x50, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x57, 0x0a, 0x1f, 0x4f, 0x65, 0x6d, 0x41, 0x70, 0x70,
	0x44, 0x65, 0x66, 0x4d, 0x65, 0x6e, 0x75, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x04, 0x6b, 0x65, 0x79,
	0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x4f, 0x65, 0x6d, 0x41, 0x70, 0x70, 0x44, 0x65, 0x66, 0x4d, 0x65, 0x6e, 0x75, 0x50,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x42,
	0x11, 0x5a, 0x0f, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_oem_app_def_menu_model_gen_proto_rawDescOnce sync.Once
	file_oem_app_def_menu_model_gen_proto_rawDescData = file_oem_app_def_menu_model_gen_proto_rawDesc
)

func file_oem_app_def_menu_model_gen_proto_rawDescGZIP() []byte {
	file_oem_app_def_menu_model_gen_proto_rawDescOnce.Do(func() {
		file_oem_app_def_menu_model_gen_proto_rawDescData = protoimpl.X.CompressGZIP(file_oem_app_def_menu_model_gen_proto_rawDescData)
	})
	return file_oem_app_def_menu_model_gen_proto_rawDescData
}

var file_oem_app_def_menu_model_gen_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_oem_app_def_menu_model_gen_proto_goTypes = []interface{}{
	(*OemAppDefMenu)(nil),                    // 0: service.OemAppDefMenu
	(*OemAppDefMenuFilter)(nil),              // 1: service.OemAppDefMenuFilter
	(*OemAppDefMenuListRequest)(nil),         // 2: service.OemAppDefMenuListRequest
	(*OemAppDefMenuResponse)(nil),            // 3: service.OemAppDefMenuResponse
	(*OemAppDefMenuUpdateFieldsRequest)(nil), // 4: service.OemAppDefMenuUpdateFieldsRequest
	(*OemAppDefMenuPrimarykey)(nil),          // 5: service.OemAppDefMenuPrimarykey
	(*OemAppDefMenuBatchDeleteRequest)(nil),  // 6: service.OemAppDefMenuBatchDeleteRequest
	(*timestamppb.Timestamp)(nil),            // 7: google.protobuf.Timestamp
}
var file_oem_app_def_menu_model_gen_proto_depIdxs = []int32{
	7,  // 0: service.OemAppDefMenu.createdAt:type_name -> google.protobuf.Timestamp
	7,  // 1: service.OemAppDefMenu.updatedAt:type_name -> google.protobuf.Timestamp
	7,  // 2: service.OemAppDefMenu.deletedAt:type_name -> google.protobuf.Timestamp
	7,  // 3: service.OemAppDefMenuFilter.createdAt:type_name -> google.protobuf.Timestamp
	7,  // 4: service.OemAppDefMenuFilter.updatedAt:type_name -> google.protobuf.Timestamp
	7,  // 5: service.OemAppDefMenuFilter.deletedAt:type_name -> google.protobuf.Timestamp
	0,  // 6: service.OemAppDefMenuListRequest.query:type_name -> service.OemAppDefMenu
	0,  // 7: service.OemAppDefMenuResponse.data:type_name -> service.OemAppDefMenu
	0,  // 8: service.OemAppDefMenuUpdateFieldsRequest.data:type_name -> service.OemAppDefMenu
	5,  // 9: service.OemAppDefMenuBatchDeleteRequest.keys:type_name -> service.OemAppDefMenuPrimarykey
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_oem_app_def_menu_model_gen_proto_init() }
func file_oem_app_def_menu_model_gen_proto_init() {
	if File_oem_app_def_menu_model_gen_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oem_app_def_menu_model_gen_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OemAppDefMenu); i {
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
		file_oem_app_def_menu_model_gen_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OemAppDefMenuFilter); i {
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
		file_oem_app_def_menu_model_gen_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OemAppDefMenuListRequest); i {
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
		file_oem_app_def_menu_model_gen_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OemAppDefMenuResponse); i {
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
		file_oem_app_def_menu_model_gen_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OemAppDefMenuUpdateFieldsRequest); i {
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
		file_oem_app_def_menu_model_gen_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OemAppDefMenuPrimarykey); i {
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
		file_oem_app_def_menu_model_gen_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OemAppDefMenuBatchDeleteRequest); i {
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
			RawDescriptor: file_oem_app_def_menu_model_gen_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_oem_app_def_menu_model_gen_proto_goTypes,
		DependencyIndexes: file_oem_app_def_menu_model_gen_proto_depIdxs,
		MessageInfos:      file_oem_app_def_menu_model_gen_proto_msgTypes,
	}.Build()
	File_oem_app_def_menu_model_gen_proto = out.File
	file_oem_app_def_menu_model_gen_proto_rawDesc = nil
	file_oem_app_def_menu_model_gen_proto_goTypes = nil
	file_oem_app_def_menu_model_gen_proto_depIdxs = nil
}
