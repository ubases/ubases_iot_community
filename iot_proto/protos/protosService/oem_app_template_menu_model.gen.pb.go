// Code generated by protoc,2022-10-24 08:40:56. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: oem_app_template_menu_model.gen.proto

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
type OemAppTemplateMenu struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:主键唯一编号;comment:主键唯一编号;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"name" gorm:"default:菜单名称（中文）;comment:菜单名称（中文）;size:50;"
	Name string `protobuf:"bytes,101,opt,name=name,proto3" json:"name,omitempty"`
	// @inject_tag: json:"nameEn" gorm:"default:菜单名称（英文）;comment:菜单名称（英文）;size:50;"
	NameEn string `protobuf:"bytes,102,opt,name=nameEn,proto3" json:"nameEn,omitempty"`
	// @inject_tag: json:"code" gorm:"default:功能Key;comment:功能Key;size:20;"
	Code string `protobuf:"bytes,103,opt,name=code,proto3" json:"code,omitempty"`
	// @inject_tag: json:"sort" gorm:"default:展示序号;comment:展示序号;size:10;"
	Sort int32 `protobuf:"varint,104,opt,name=sort,proto3" json:"sort,omitempty"`
	// @inject_tag: json:"appTemplateId" gorm:"default:;comment:;size:19;"
	AppTemplateId int64 `protobuf:"varint,105,opt,name=appTemplateId,proto3" json:"appTemplateId,omitempty"`
	// @inject_tag: json:"pagePath" gorm:"default:访问地址;comment:访问地址;size:255;"
	PagePath string `protobuf:"bytes,106,opt,name=pagePath,proto3" json:"pagePath,omitempty"`
	// @inject_tag: json:"required" gorm:"default:;comment:;size:255;"
	Required int32 `protobuf:"varint,107,opt,name=required,proto3" json:"required,omitempty"`
	// @inject_tag: json:"defaultIcon" gorm:"default:默认图标;comment:默认图标;size:255;"
	DefaultIcon string `protobuf:"bytes,108,opt,name=defaultIcon,proto3" json:"defaultIcon,omitempty"`
	// @inject_tag: json:"selectedIcon" gorm:"default:选择图标;comment:选择图标;size:255;"
	SelectedIcon string `protobuf:"bytes,109,opt,name=selectedIcon,proto3" json:"selectedIcon,omitempty"`
	// @inject_tag: json:"createdBy" gorm:"default:修改人;comment:修改人;size:19;"
	CreatedBy int64 `protobuf:"varint,110,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	// @inject_tag: json:"createdAt" gorm:"default:创建时间;comment:创建时间;"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,111,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// @inject_tag: json:"updatedBy" gorm:"default:修改人;comment:修改人;size:19;"
	UpdatedBy int64 `protobuf:"varint,112,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	// @inject_tag: json:"updatedAt" gorm:"default:修改时间;comment:修改时间;"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,113,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *OemAppTemplateMenu) Reset() {
	*x = OemAppTemplateMenu{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oem_app_template_menu_model_gen_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OemAppTemplateMenu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OemAppTemplateMenu) ProtoMessage() {}

func (x *OemAppTemplateMenu) ProtoReflect() protoreflect.Message {
	mi := &file_oem_app_template_menu_model_gen_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OemAppTemplateMenu.ProtoReflect.Descriptor instead.
func (*OemAppTemplateMenu) Descriptor() ([]byte, []int) {
	return file_oem_app_template_menu_model_gen_proto_rawDescGZIP(), []int{0}
}

func (x *OemAppTemplateMenu) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OemAppTemplateMenu) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OemAppTemplateMenu) GetNameEn() string {
	if x != nil {
		return x.NameEn
	}
	return ""
}

func (x *OemAppTemplateMenu) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *OemAppTemplateMenu) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *OemAppTemplateMenu) GetAppTemplateId() int64 {
	if x != nil {
		return x.AppTemplateId
	}
	return 0
}

func (x *OemAppTemplateMenu) GetPagePath() string {
	if x != nil {
		return x.PagePath
	}
	return ""
}

func (x *OemAppTemplateMenu) GetRequired() int32 {
	if x != nil {
		return x.Required
	}
	return 0
}

func (x *OemAppTemplateMenu) GetDefaultIcon() string {
	if x != nil {
		return x.DefaultIcon
	}
	return ""
}

func (x *OemAppTemplateMenu) GetSelectedIcon() string {
	if x != nil {
		return x.SelectedIcon
	}
	return ""
}

func (x *OemAppTemplateMenu) GetCreatedBy() int64 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *OemAppTemplateMenu) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *OemAppTemplateMenu) GetUpdatedBy() int64 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

func (x *OemAppTemplateMenu) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type OemAppTemplateMenuFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:主键唯一编号;comment:主键唯一编号;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"name" gorm:"default:菜单名称（中文）;comment:菜单名称（中文）;size:50;"
	Name string `protobuf:"bytes,101,opt,name=name,proto3" json:"name,omitempty"`
	// @inject_tag: json:"nameEn" gorm:"default:菜单名称（英文）;comment:菜单名称（英文）;size:50;"
	NameEn string `protobuf:"bytes,102,opt,name=nameEn,proto3" json:"nameEn,omitempty"`
	// @inject_tag: json:"code" gorm:"default:功能Key;comment:功能Key;size:20;"
	Code string `protobuf:"bytes,103,opt,name=code,proto3" json:"code,omitempty"`
	// @inject_tag: json:"sort" gorm:"default:展示序号;comment:展示序号;size:10;"
	Sort int32 `protobuf:"varint,104,opt,name=sort,proto3" json:"sort,omitempty"`
	// @inject_tag: json:"appTemplateId" gorm:"default:;comment:;size:19;"
	AppTemplateId int64 `protobuf:"varint,105,opt,name=appTemplateId,proto3" json:"appTemplateId,omitempty"`
	// @inject_tag: json:"pagePath" gorm:"default:访问地址;comment:访问地址;size:255;"
	PagePath string `protobuf:"bytes,106,opt,name=pagePath,proto3" json:"pagePath,omitempty"`
	// @inject_tag: json:"required" gorm:"default:;comment:;size:255;"
	Required int32 `protobuf:"varint,107,opt,name=required,proto3" json:"required,omitempty"`
	// @inject_tag: json:"defaultIcon" gorm:"default:默认图标;comment:默认图标;size:255;"
	DefaultIcon string `protobuf:"bytes,108,opt,name=defaultIcon,proto3" json:"defaultIcon,omitempty"`
	// @inject_tag: json:"selectedIcon" gorm:"default:选择图标;comment:选择图标;size:255;"
	SelectedIcon string `protobuf:"bytes,109,opt,name=selectedIcon,proto3" json:"selectedIcon,omitempty"`
	// @inject_tag: json:"createdBy" gorm:"default:修改人;comment:修改人;size:19;"
	CreatedBy int64 `protobuf:"varint,110,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	// @inject_tag: json:"createdAt" gorm:"default:创建时间;comment:创建时间;"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,111,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// @inject_tag: json:"updatedBy" gorm:"default:修改人;comment:修改人;size:19;"
	UpdatedBy int64 `protobuf:"varint,112,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	// @inject_tag: json:"updatedAt" gorm:"default:修改时间;comment:修改时间;"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,113,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *OemAppTemplateMenuFilter) Reset() {
	*x = OemAppTemplateMenuFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oem_app_template_menu_model_gen_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OemAppTemplateMenuFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OemAppTemplateMenuFilter) ProtoMessage() {}

func (x *OemAppTemplateMenuFilter) ProtoReflect() protoreflect.Message {
	mi := &file_oem_app_template_menu_model_gen_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OemAppTemplateMenuFilter.ProtoReflect.Descriptor instead.
func (*OemAppTemplateMenuFilter) Descriptor() ([]byte, []int) {
	return file_oem_app_template_menu_model_gen_proto_rawDescGZIP(), []int{1}
}

func (x *OemAppTemplateMenuFilter) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OemAppTemplateMenuFilter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OemAppTemplateMenuFilter) GetNameEn() string {
	if x != nil {
		return x.NameEn
	}
	return ""
}

func (x *OemAppTemplateMenuFilter) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *OemAppTemplateMenuFilter) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *OemAppTemplateMenuFilter) GetAppTemplateId() int64 {
	if x != nil {
		return x.AppTemplateId
	}
	return 0
}

func (x *OemAppTemplateMenuFilter) GetPagePath() string {
	if x != nil {
		return x.PagePath
	}
	return ""
}

func (x *OemAppTemplateMenuFilter) GetRequired() int32 {
	if x != nil {
		return x.Required
	}
	return 0
}

func (x *OemAppTemplateMenuFilter) GetDefaultIcon() string {
	if x != nil {
		return x.DefaultIcon
	}
	return ""
}

func (x *OemAppTemplateMenuFilter) GetSelectedIcon() string {
	if x != nil {
		return x.SelectedIcon
	}
	return ""
}

func (x *OemAppTemplateMenuFilter) GetCreatedBy() int64 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *OemAppTemplateMenuFilter) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *OemAppTemplateMenuFilter) GetUpdatedBy() int64 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

func (x *OemAppTemplateMenuFilter) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type OemAppTemplateMenuListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query     *OemAppTemplateMenu `protobuf:"bytes,101,opt,name=query,proto3" json:"query,omitempty"`
	Page      int64               `protobuf:"varint,102,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int64               `protobuf:"varint,103,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	OrderKey  string              `protobuf:"bytes,104,opt,name=orderKey,proto3" json:"orderKey,omitempty"`
	OrderDesc string              `protobuf:"bytes,105,opt,name=orderDesc,proto3" json:"orderDesc,omitempty"`
	SearchKey string              `protobuf:"bytes,106,opt,name=searchKey,proto3" json:"searchKey,omitempty"`
}

func (x *OemAppTemplateMenuListRequest) Reset() {
	*x = OemAppTemplateMenuListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oem_app_template_menu_model_gen_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OemAppTemplateMenuListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OemAppTemplateMenuListRequest) ProtoMessage() {}

func (x *OemAppTemplateMenuListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oem_app_template_menu_model_gen_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OemAppTemplateMenuListRequest.ProtoReflect.Descriptor instead.
func (*OemAppTemplateMenuListRequest) Descriptor() ([]byte, []int) {
	return file_oem_app_template_menu_model_gen_proto_rawDescGZIP(), []int{2}
}

func (x *OemAppTemplateMenuListRequest) GetQuery() *OemAppTemplateMenu {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *OemAppTemplateMenuListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *OemAppTemplateMenuListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *OemAppTemplateMenuListRequest) GetOrderKey() string {
	if x != nil {
		return x.OrderKey
	}
	return ""
}

func (x *OemAppTemplateMenuListRequest) GetOrderDesc() string {
	if x != nil {
		return x.OrderDesc
	}
	return ""
}

func (x *OemAppTemplateMenuListRequest) GetSearchKey() string {
	if x != nil {
		return x.SearchKey
	}
	return ""
}

type OemAppTemplateMenuResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32                 `protobuf:"varint,101,opt,name=code,proto3" json:"code,omitempty"`
	Message string                `protobuf:"bytes,102,opt,name=message,proto3" json:"message,omitempty"`
	Total   int64                 `protobuf:"varint,103,opt,name=total,proto3" json:"total,omitempty"`
	Data    []*OemAppTemplateMenu `protobuf:"bytes,104,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *OemAppTemplateMenuResponse) Reset() {
	*x = OemAppTemplateMenuResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oem_app_template_menu_model_gen_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OemAppTemplateMenuResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OemAppTemplateMenuResponse) ProtoMessage() {}

func (x *OemAppTemplateMenuResponse) ProtoReflect() protoreflect.Message {
	mi := &file_oem_app_template_menu_model_gen_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OemAppTemplateMenuResponse.ProtoReflect.Descriptor instead.
func (*OemAppTemplateMenuResponse) Descriptor() ([]byte, []int) {
	return file_oem_app_template_menu_model_gen_proto_rawDescGZIP(), []int{3}
}

func (x *OemAppTemplateMenuResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *OemAppTemplateMenuResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *OemAppTemplateMenuResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *OemAppTemplateMenuResponse) GetData() []*OemAppTemplateMenu {
	if x != nil {
		return x.Data
	}
	return nil
}

type OemAppTemplateMenuUpdateFieldsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []string            `protobuf:"bytes,101,rep,name=fields,proto3" json:"fields,omitempty"`
	Data   *OemAppTemplateMenu `protobuf:"bytes,102,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OemAppTemplateMenuUpdateFieldsRequest) Reset() {
	*x = OemAppTemplateMenuUpdateFieldsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oem_app_template_menu_model_gen_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OemAppTemplateMenuUpdateFieldsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OemAppTemplateMenuUpdateFieldsRequest) ProtoMessage() {}

func (x *OemAppTemplateMenuUpdateFieldsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oem_app_template_menu_model_gen_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OemAppTemplateMenuUpdateFieldsRequest.ProtoReflect.Descriptor instead.
func (*OemAppTemplateMenuUpdateFieldsRequest) Descriptor() ([]byte, []int) {
	return file_oem_app_template_menu_model_gen_proto_rawDescGZIP(), []int{4}
}

func (x *OemAppTemplateMenuUpdateFieldsRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *OemAppTemplateMenuUpdateFieldsRequest) GetData() *OemAppTemplateMenu {
	if x != nil {
		return x.Data
	}
	return nil
}

// 表模型主键，只能是整数或字符串，若生成错误，请检查数据库表主键是否正确。支持组合主键
type OemAppTemplateMenuPrimarykey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OemAppTemplateMenuPrimarykey) Reset() {
	*x = OemAppTemplateMenuPrimarykey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oem_app_template_menu_model_gen_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OemAppTemplateMenuPrimarykey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OemAppTemplateMenuPrimarykey) ProtoMessage() {}

func (x *OemAppTemplateMenuPrimarykey) ProtoReflect() protoreflect.Message {
	mi := &file_oem_app_template_menu_model_gen_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OemAppTemplateMenuPrimarykey.ProtoReflect.Descriptor instead.
func (*OemAppTemplateMenuPrimarykey) Descriptor() ([]byte, []int) {
	return file_oem_app_template_menu_model_gen_proto_rawDescGZIP(), []int{5}
}

func (x *OemAppTemplateMenuPrimarykey) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type OemAppTemplateMenuBatchDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []*OemAppTemplateMenuPrimarykey `protobuf:"bytes,101,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *OemAppTemplateMenuBatchDeleteRequest) Reset() {
	*x = OemAppTemplateMenuBatchDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oem_app_template_menu_model_gen_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OemAppTemplateMenuBatchDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OemAppTemplateMenuBatchDeleteRequest) ProtoMessage() {}

func (x *OemAppTemplateMenuBatchDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oem_app_template_menu_model_gen_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OemAppTemplateMenuBatchDeleteRequest.ProtoReflect.Descriptor instead.
func (*OemAppTemplateMenuBatchDeleteRequest) Descriptor() ([]byte, []int) {
	return file_oem_app_template_menu_model_gen_proto_rawDescGZIP(), []int{6}
}

func (x *OemAppTemplateMenuBatchDeleteRequest) GetKeys() []*OemAppTemplateMenuPrimarykey {
	if x != nil {
		return x.Keys
	}
	return nil
}

var File_oem_app_template_menu_model_gen_proto protoreflect.FileDescriptor

var file_oem_app_template_menu_model_gen_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6f, 0x65, 0x6d, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x5f, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x67, 0x65,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xcc, 0x03, 0x0a, 0x12, 0x4f, 0x65, 0x6d, 0x41, 0x70, 0x70, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x61,
	0x6d, 0x65, 0x45, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x67, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74,
	0x18, 0x68, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x24, 0x0a, 0x0d,
	0x61, 0x70, 0x70, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18, 0x69, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x6a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x6b, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x49, 0x63, 0x6f, 0x6e, 0x18, 0x6c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x49, 0x63, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c,
	0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x49, 0x63, 0x6f, 0x6e, 0x18, 0x6d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x49, 0x63, 0x6f, 0x6e,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x6e, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x38,
	0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x6f, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x70, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x71, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0xd2, 0x03, 0x0a, 0x18, 0x4f, 0x65, 0x6d, 0x41, 0x70, 0x70, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x18, 0x66, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6e, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x68, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72,
	0x74, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x70, 0x70, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x49, 0x64, 0x18, 0x69, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x50,
	0x61, 0x74, 0x68, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18,
	0x6b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x49, 0x63, 0x6f, 0x6e, 0x18, 0x6c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x49, 0x63, 0x6f,
	0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x49, 0x63, 0x6f,
	0x6e, 0x18, 0x6d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x49, 0x63, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x6f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x70, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x71, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xda, 0x01, 0x0a, 0x1d, 0x4f, 0x65, 0x6d, 0x41, 0x70, 0x70,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4f, 0x65, 0x6d, 0x41, 0x70, 0x70, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x6e, 0x75, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44,
	0x65, 0x73, 0x63, 0x18, 0x69, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x44, 0x65, 0x73, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65,
	0x79, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b,
	0x65, 0x79, 0x22, 0x91, 0x01, 0x0a, 0x1a, 0x4f, 0x65, 0x6d, 0x41, 0x70, 0x70, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x68, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x65,
	0x6d, 0x41, 0x70, 0x70, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x70, 0x0a, 0x25, 0x4f, 0x65, 0x6d, 0x41, 0x70, 0x70,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x4f, 0x65, 0x6d, 0x41, 0x70, 0x70, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x65,
	0x6e, 0x75, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2e, 0x0a, 0x1c, 0x4f, 0x65, 0x6d, 0x41,
	0x70, 0x70, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x50, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x61, 0x0a, 0x24, 0x4f, 0x65, 0x6d, 0x41,
	0x70, 0x70, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x39, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x65, 0x6d, 0x41, 0x70, 0x70, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x6b, 0x65, 0x79, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x42, 0x11, 0x5a, 0x0f, 0x2e,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_oem_app_template_menu_model_gen_proto_rawDescOnce sync.Once
	file_oem_app_template_menu_model_gen_proto_rawDescData = file_oem_app_template_menu_model_gen_proto_rawDesc
)

func file_oem_app_template_menu_model_gen_proto_rawDescGZIP() []byte {
	file_oem_app_template_menu_model_gen_proto_rawDescOnce.Do(func() {
		file_oem_app_template_menu_model_gen_proto_rawDescData = protoimpl.X.CompressGZIP(file_oem_app_template_menu_model_gen_proto_rawDescData)
	})
	return file_oem_app_template_menu_model_gen_proto_rawDescData
}

var file_oem_app_template_menu_model_gen_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_oem_app_template_menu_model_gen_proto_goTypes = []interface{}{
	(*OemAppTemplateMenu)(nil),                    // 0: service.OemAppTemplateMenu
	(*OemAppTemplateMenuFilter)(nil),              // 1: service.OemAppTemplateMenuFilter
	(*OemAppTemplateMenuListRequest)(nil),         // 2: service.OemAppTemplateMenuListRequest
	(*OemAppTemplateMenuResponse)(nil),            // 3: service.OemAppTemplateMenuResponse
	(*OemAppTemplateMenuUpdateFieldsRequest)(nil), // 4: service.OemAppTemplateMenuUpdateFieldsRequest
	(*OemAppTemplateMenuPrimarykey)(nil),          // 5: service.OemAppTemplateMenuPrimarykey
	(*OemAppTemplateMenuBatchDeleteRequest)(nil),  // 6: service.OemAppTemplateMenuBatchDeleteRequest
	(*timestamppb.Timestamp)(nil),                 // 7: google.protobuf.Timestamp
}
var file_oem_app_template_menu_model_gen_proto_depIdxs = []int32{
	7, // 0: service.OemAppTemplateMenu.createdAt:type_name -> google.protobuf.Timestamp
	7, // 1: service.OemAppTemplateMenu.updatedAt:type_name -> google.protobuf.Timestamp
	7, // 2: service.OemAppTemplateMenuFilter.createdAt:type_name -> google.protobuf.Timestamp
	7, // 3: service.OemAppTemplateMenuFilter.updatedAt:type_name -> google.protobuf.Timestamp
	0, // 4: service.OemAppTemplateMenuListRequest.query:type_name -> service.OemAppTemplateMenu
	0, // 5: service.OemAppTemplateMenuResponse.data:type_name -> service.OemAppTemplateMenu
	0, // 6: service.OemAppTemplateMenuUpdateFieldsRequest.data:type_name -> service.OemAppTemplateMenu
	5, // 7: service.OemAppTemplateMenuBatchDeleteRequest.keys:type_name -> service.OemAppTemplateMenuPrimarykey
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_oem_app_template_menu_model_gen_proto_init() }
func file_oem_app_template_menu_model_gen_proto_init() {
	if File_oem_app_template_menu_model_gen_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oem_app_template_menu_model_gen_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OemAppTemplateMenu); i {
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
		file_oem_app_template_menu_model_gen_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OemAppTemplateMenuFilter); i {
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
		file_oem_app_template_menu_model_gen_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OemAppTemplateMenuListRequest); i {
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
		file_oem_app_template_menu_model_gen_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OemAppTemplateMenuResponse); i {
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
		file_oem_app_template_menu_model_gen_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OemAppTemplateMenuUpdateFieldsRequest); i {
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
		file_oem_app_template_menu_model_gen_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OemAppTemplateMenuPrimarykey); i {
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
		file_oem_app_template_menu_model_gen_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OemAppTemplateMenuBatchDeleteRequest); i {
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
			RawDescriptor: file_oem_app_template_menu_model_gen_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_oem_app_template_menu_model_gen_proto_goTypes,
		DependencyIndexes: file_oem_app_template_menu_model_gen_proto_depIdxs,
		MessageInfos:      file_oem_app_template_menu_model_gen_proto_msgTypes,
	}.Build()
	File_oem_app_template_menu_model_gen_proto = out.File
	file_oem_app_template_menu_model_gen_proto_rawDesc = nil
	file_oem_app_template_menu_model_gen_proto_goTypes = nil
	file_oem_app_template_menu_model_gen_proto_depIdxs = nil
}
