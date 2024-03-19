// Code generated by protoc,2022-05-21 11:08:04. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: sys_area_model.proto

package protosService

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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
type SysArea struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"primary_key;AUTO_INCREMENT;default:主键;comment:主键;size:20;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"pid" gorm:"default:父ID;comment:父ID;size:20;"
	Pid int64 `protobuf:"varint,101,opt,name=pid,proto3" json:"pid,omitempty"`
	// @inject_tag: json:"level" gorm:"default:层级;comment:层级;size:10;"
	Level int32 `protobuf:"varint,102,opt,name=level,proto3" json:"level,omitempty"`
	// @inject_tag: json:"path" gorm:"default:路径;comment:路径;size:16;"
	Path string `protobuf:"bytes,103,opt,name=path,proto3" json:"path,omitempty"`
	// @inject_tag: json:"code" gorm:"default:代码;comment:代码;size:8;"
	Code string `protobuf:"bytes,104,opt,name=code,proto3" json:"code,omitempty"`
	// @inject_tag: json:"areaNumber" gorm:"default:国家区号;comment:国家区号;size:8;"
	AreaNumber string `protobuf:"bytes,105,opt,name=areaNumber,proto3" json:"areaNumber,omitempty"`
	// @inject_tag: json:"areaPhoneNumber" gorm:"default:国家电话区号;comment:国家电话区号;size:8;"
	AreaPhoneNumber string `protobuf:"bytes,106,opt,name=areaPhoneNumber,proto3" json:"areaPhoneNumber,omitempty"`
	// @inject_tag: json:"abbreviation" gorm:"default:国家缩写;comment:国家缩写;size:16;"
	Abbreviation string `protobuf:"bytes,107,opt,name=abbreviation,proto3" json:"abbreviation,omitempty"`
	// @inject_tag: json:"iso" gorm:"default:时区;comment:时区;size:16;"
	Iso string `protobuf:"bytes,108,opt,name=iso,proto3" json:"iso,omitempty"`
	// @inject_tag: json:"chineseName" gorm:"default:中文名称;comment:中文名称;size:64;"
	ChineseName string `protobuf:"bytes,109,opt,name=chineseName,proto3" json:"chineseName,omitempty"`
	// @inject_tag: json:"englishName" gorm:"default:英文名称;comment:英文名称;size:64;"
	EnglishName string `protobuf:"bytes,110,opt,name=englishName,proto3" json:"englishName,omitempty"`
	// @inject_tag: json:"pinyin" gorm:"default:中文拼音;comment:中文拼音;size:64;"
	Pinyin string `protobuf:"bytes,111,opt,name=pinyin,proto3" json:"pinyin,omitempty"`
	// 是否返回子集数据
	ShowChild bool `protobuf:"varint,112,opt,name=showChild,proto3" json:"showChild,omitempty"`
	//启用查询国家、城市和地区名称获取Code
	EnableGetCode bool    `protobuf:"varint,113,opt,name=enableGetCode,proto3" json:"enableGetCode,omitempty"`
	Country       string  `protobuf:"bytes,114,opt,name=country,proto3" json:"country,omitempty"`
	Province      string  `protobuf:"bytes,115,opt,name=province,proto3" json:"province,omitempty"`
	City          string  `protobuf:"bytes,116,opt,name=city,proto3" json:"city,omitempty"`
	District      string  `protobuf:"bytes,117,opt,name=district,proto3" json:"district,omitempty"`
	AreaIds       []int64 `protobuf:"varint,118,rep,packed,name=areaIds,proto3" json:"areaIds,omitempty"`
}

func (x *SysArea) Reset() {
	*x = SysArea{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_area_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysArea) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysArea) ProtoMessage() {}

func (x *SysArea) ProtoReflect() protoreflect.Message {
	mi := &file_sys_area_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysArea.ProtoReflect.Descriptor instead.
func (*SysArea) Descriptor() ([]byte, []int) {
	return file_sys_area_model_proto_rawDescGZIP(), []int{0}
}

func (x *SysArea) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SysArea) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *SysArea) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *SysArea) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *SysArea) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *SysArea) GetAreaNumber() string {
	if x != nil {
		return x.AreaNumber
	}
	return ""
}

func (x *SysArea) GetAreaPhoneNumber() string {
	if x != nil {
		return x.AreaPhoneNumber
	}
	return ""
}

func (x *SysArea) GetAbbreviation() string {
	if x != nil {
		return x.Abbreviation
	}
	return ""
}

func (x *SysArea) GetIso() string {
	if x != nil {
		return x.Iso
	}
	return ""
}

func (x *SysArea) GetChineseName() string {
	if x != nil {
		return x.ChineseName
	}
	return ""
}

func (x *SysArea) GetEnglishName() string {
	if x != nil {
		return x.EnglishName
	}
	return ""
}

func (x *SysArea) GetPinyin() string {
	if x != nil {
		return x.Pinyin
	}
	return ""
}

func (x *SysArea) GetShowChild() bool {
	if x != nil {
		return x.ShowChild
	}
	return false
}

func (x *SysArea) GetEnableGetCode() bool {
	if x != nil {
		return x.EnableGetCode
	}
	return false
}

func (x *SysArea) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *SysArea) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *SysArea) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *SysArea) GetDistrict() string {
	if x != nil {
		return x.District
	}
	return ""
}

func (x *SysArea) GetAreaIds() []int64 {
	if x != nil {
		return x.AreaIds
	}
	return nil
}

type SysAreaFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"primary_key;AUTO_INCREMENT;default:主键;comment:主键;size:20;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"pid" gorm:"default:父ID;comment:父ID;size:20;"
	Pid int64 `protobuf:"varint,101,opt,name=pid,proto3" json:"pid,omitempty"`
	// @inject_tag: json:"level" gorm:"default:层级;comment:层级;size:10;"
	Level int32 `protobuf:"varint,102,opt,name=level,proto3" json:"level,omitempty"`
	// @inject_tag: json:"path" gorm:"default:路径;comment:路径;size:16;"
	Path string `protobuf:"bytes,103,opt,name=path,proto3" json:"path,omitempty"`
	// @inject_tag: json:"code" gorm:"default:代码;comment:代码;size:8;"
	Code string `protobuf:"bytes,104,opt,name=code,proto3" json:"code,omitempty"`
	// @inject_tag: json:"areaNumber" gorm:"default:国家区号;comment:国家区号;size:8;"
	AreaNumber string `protobuf:"bytes,105,opt,name=areaNumber,proto3" json:"areaNumber,omitempty"`
	// @inject_tag: json:"areaPhoneNumber" gorm:"default:国家电话区号;comment:国家电话区号;size:8;"
	AreaPhoneNumber string `protobuf:"bytes,106,opt,name=areaPhoneNumber,proto3" json:"areaPhoneNumber,omitempty"`
	// @inject_tag: json:"abbreviation" gorm:"default:国家缩写;comment:国家缩写;size:16;"
	Abbreviation string `protobuf:"bytes,107,opt,name=abbreviation,proto3" json:"abbreviation,omitempty"`
	// @inject_tag: json:"iso" gorm:"default:时区;comment:时区;size:16;"
	Iso string `protobuf:"bytes,108,opt,name=iso,proto3" json:"iso,omitempty"`
	// @inject_tag: json:"chineseName" gorm:"default:中文名称;comment:中文名称;size:64;"
	ChineseName string `protobuf:"bytes,109,opt,name=chineseName,proto3" json:"chineseName,omitempty"`
	// @inject_tag: json:"englishName" gorm:"default:英文名称;comment:英文名称;size:64;"
	EnglishName string `protobuf:"bytes,110,opt,name=englishName,proto3" json:"englishName,omitempty"`
	// @inject_tag: json:"pinyin" gorm:"default:中文拼音;comment:中文拼音;size:64;"
	Pinyin string `protobuf:"bytes,111,opt,name=pinyin,proto3" json:"pinyin,omitempty"`
}

func (x *SysAreaFilter) Reset() {
	*x = SysAreaFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_area_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysAreaFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysAreaFilter) ProtoMessage() {}

func (x *SysAreaFilter) ProtoReflect() protoreflect.Message {
	mi := &file_sys_area_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysAreaFilter.ProtoReflect.Descriptor instead.
func (*SysAreaFilter) Descriptor() ([]byte, []int) {
	return file_sys_area_model_proto_rawDescGZIP(), []int{1}
}

func (x *SysAreaFilter) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SysAreaFilter) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *SysAreaFilter) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *SysAreaFilter) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *SysAreaFilter) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *SysAreaFilter) GetAreaNumber() string {
	if x != nil {
		return x.AreaNumber
	}
	return ""
}

func (x *SysAreaFilter) GetAreaPhoneNumber() string {
	if x != nil {
		return x.AreaPhoneNumber
	}
	return ""
}

func (x *SysAreaFilter) GetAbbreviation() string {
	if x != nil {
		return x.Abbreviation
	}
	return ""
}

func (x *SysAreaFilter) GetIso() string {
	if x != nil {
		return x.Iso
	}
	return ""
}

func (x *SysAreaFilter) GetChineseName() string {
	if x != nil {
		return x.ChineseName
	}
	return ""
}

func (x *SysAreaFilter) GetEnglishName() string {
	if x != nil {
		return x.EnglishName
	}
	return ""
}

func (x *SysAreaFilter) GetPinyin() string {
	if x != nil {
		return x.Pinyin
	}
	return ""
}

type SysAreaListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query     *SysArea `protobuf:"bytes,101,opt,name=query,proto3" json:"query,omitempty"`
	Page      int64    `protobuf:"varint,102,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int64    `protobuf:"varint,103,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	OrderKey  string   `protobuf:"bytes,104,opt,name=orderKey,proto3" json:"orderKey,omitempty"`
	OrderDesc string   `protobuf:"bytes,105,opt,name=orderDesc,proto3" json:"orderDesc,omitempty"`
	SearchKey string   `protobuf:"bytes,106,opt,name=searchKey,proto3" json:"searchKey,omitempty"`
}

func (x *SysAreaListRequest) Reset() {
	*x = SysAreaListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_area_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysAreaListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysAreaListRequest) ProtoMessage() {}

func (x *SysAreaListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sys_area_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysAreaListRequest.ProtoReflect.Descriptor instead.
func (*SysAreaListRequest) Descriptor() ([]byte, []int) {
	return file_sys_area_model_proto_rawDescGZIP(), []int{2}
}

func (x *SysAreaListRequest) GetQuery() *SysArea {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *SysAreaListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SysAreaListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SysAreaListRequest) GetOrderKey() string {
	if x != nil {
		return x.OrderKey
	}
	return ""
}

func (x *SysAreaListRequest) GetOrderDesc() string {
	if x != nil {
		return x.OrderDesc
	}
	return ""
}

func (x *SysAreaListRequest) GetSearchKey() string {
	if x != nil {
		return x.SearchKey
	}
	return ""
}

type SysAreaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32      `protobuf:"varint,101,opt,name=code,proto3" json:"code,omitempty"`
	Message string     `protobuf:"bytes,102,opt,name=message,proto3" json:"message,omitempty"`
	Total   int64      `protobuf:"varint,103,opt,name=total,proto3" json:"total,omitempty"`
	Data    []*SysArea `protobuf:"bytes,104,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *SysAreaResponse) Reset() {
	*x = SysAreaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_area_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysAreaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysAreaResponse) ProtoMessage() {}

func (x *SysAreaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sys_area_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysAreaResponse.ProtoReflect.Descriptor instead.
func (*SysAreaResponse) Descriptor() ([]byte, []int) {
	return file_sys_area_model_proto_rawDescGZIP(), []int{3}
}

func (x *SysAreaResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SysAreaResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SysAreaResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *SysAreaResponse) GetData() []*SysArea {
	if x != nil {
		return x.Data
	}
	return nil
}

type SysAreaUpdateFieldsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []string `protobuf:"bytes,101,rep,name=fields,proto3" json:"fields,omitempty"`
	Data   *SysArea `protobuf:"bytes,102,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SysAreaUpdateFieldsRequest) Reset() {
	*x = SysAreaUpdateFieldsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_area_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysAreaUpdateFieldsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysAreaUpdateFieldsRequest) ProtoMessage() {}

func (x *SysAreaUpdateFieldsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sys_area_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysAreaUpdateFieldsRequest.ProtoReflect.Descriptor instead.
func (*SysAreaUpdateFieldsRequest) Descriptor() ([]byte, []int) {
	return file_sys_area_model_proto_rawDescGZIP(), []int{4}
}

func (x *SysAreaUpdateFieldsRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *SysAreaUpdateFieldsRequest) GetData() *SysArea {
	if x != nil {
		return x.Data
	}
	return nil
}

// 表模型主键，只能是整数或字符串，若生成错误，请检查数据库表主键是否正确。支持组合主键
type SysAreaPrimarykey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SysAreaPrimarykey) Reset() {
	*x = SysAreaPrimarykey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_area_model_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysAreaPrimarykey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysAreaPrimarykey) ProtoMessage() {}

func (x *SysAreaPrimarykey) ProtoReflect() protoreflect.Message {
	mi := &file_sys_area_model_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysAreaPrimarykey.ProtoReflect.Descriptor instead.
func (*SysAreaPrimarykey) Descriptor() ([]byte, []int) {
	return file_sys_area_model_proto_rawDescGZIP(), []int{5}
}

func (x *SysAreaPrimarykey) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SysAreaBatchDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []*SysAreaPrimarykey `protobuf:"bytes,101,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *SysAreaBatchDeleteRequest) Reset() {
	*x = SysAreaBatchDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_area_model_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysAreaBatchDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysAreaBatchDeleteRequest) ProtoMessage() {}

func (x *SysAreaBatchDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sys_area_model_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysAreaBatchDeleteRequest.ProtoReflect.Descriptor instead.
func (*SysAreaBatchDeleteRequest) Descriptor() ([]byte, []int) {
	return file_sys_area_model_proto_rawDescGZIP(), []int{6}
}

func (x *SysAreaBatchDeleteRequest) GetKeys() []*SysAreaPrimarykey {
	if x != nil {
		return x.Keys
	}
	return nil
}

var File_sys_area_model_proto protoreflect.FileDescriptor

var file_sys_area_model_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x79, 0x73, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22,
	0x89, 0x04, 0x0a, 0x07, 0x53, 0x79, 0x73, 0x41, 0x72, 0x65, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70,
	0x69, 0x64, 0x18, 0x65, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x66, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x67, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61,
	0x72, 0x65, 0x61, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x69, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x61, 0x72, 0x65, 0x61, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x0f, 0x61,
	0x72, 0x65, 0x61, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x6a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x72, 0x65, 0x61, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x62, 0x62, 0x72, 0x65, 0x76, 0x69,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x6b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x62, 0x62,
	0x72, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x73, 0x6f,
	0x18, 0x6c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x73, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x6d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x65, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x6e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x70, 0x69, 0x6e, 0x79, 0x69, 0x6e, 0x18, 0x6f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x70, 0x69, 0x6e, 0x79, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x77, 0x43,
	0x68, 0x69, 0x6c, 0x64, 0x18, 0x70, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x77,
	0x43, 0x68, 0x69, 0x6c, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x71, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x72, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63,
	0x65, 0x18, 0x73, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x74, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x18, 0x75, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x73, 0x18, 0x76, 0x20, 0x03,
	0x28, 0x03, 0x52, 0x07, 0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x73, 0x22, 0xcb, 0x02, 0x0a, 0x0d,
	0x53, 0x79, 0x73, 0x41, 0x72, 0x65, 0x61, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x70, 0x69, 0x64, 0x18, 0x65, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x66, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x67, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x61, 0x72, 0x65, 0x61, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x69, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x61, 0x72, 0x65, 0x61, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x28, 0x0a,
	0x0f, 0x61, 0x72, 0x65, 0x61, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x6a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x72, 0x65, 0x61, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x62, 0x62, 0x72, 0x65,
	0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x6b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61,
	0x62, 0x62, 0x72, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x73, 0x6f, 0x18, 0x6c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x73, 0x6f, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x6d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x6e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x69, 0x6e, 0x79, 0x69, 0x6e, 0x18, 0x6f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x69, 0x6e, 0x79, 0x69, 0x6e, 0x22, 0xc4, 0x01, 0x0a, 0x12, 0x53, 0x79,
	0x73, 0x41, 0x72, 0x65, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x26, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x79, 0x73, 0x41, 0x72, 0x65,
	0x61, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x66, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x4b, 0x65, 0x79, 0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x73,
	0x63, 0x18, 0x69, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65,
	0x73, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x18,
	0x6a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79,
	0x22, 0x7b, 0x0a, 0x0f, 0x53, 0x79, 0x73, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x68, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x53, 0x79, 0x73, 0x41, 0x72, 0x65, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x5a, 0x0a,
	0x1a, 0x53, 0x79, 0x73, 0x41, 0x72, 0x65, 0x61, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x66, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x79, 0x73, 0x41,
	0x72, 0x65, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x23, 0x0a, 0x11, 0x53, 0x79, 0x73,
	0x41, 0x72, 0x65, 0x61, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4b,
	0x0a, 0x19, 0x53, 0x79, 0x73, 0x41, 0x72, 0x65, 0x61, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x6b,
	0x65, 0x79, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x53, 0x79, 0x73, 0x41, 0x72, 0x65, 0x61, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x6b, 0x65, 0x79, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x42, 0x11, 0x5a, 0x0f, 0x2e,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sys_area_model_proto_rawDescOnce sync.Once
	file_sys_area_model_proto_rawDescData = file_sys_area_model_proto_rawDesc
)

func file_sys_area_model_proto_rawDescGZIP() []byte {
	file_sys_area_model_proto_rawDescOnce.Do(func() {
		file_sys_area_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_sys_area_model_proto_rawDescData)
	})
	return file_sys_area_model_proto_rawDescData
}

var file_sys_area_model_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_sys_area_model_proto_goTypes = []interface{}{
	(*SysArea)(nil),                    // 0: service.SysArea
	(*SysAreaFilter)(nil),              // 1: service.SysAreaFilter
	(*SysAreaListRequest)(nil),         // 2: service.SysAreaListRequest
	(*SysAreaResponse)(nil),            // 3: service.SysAreaResponse
	(*SysAreaUpdateFieldsRequest)(nil), // 4: service.SysAreaUpdateFieldsRequest
	(*SysAreaPrimarykey)(nil),          // 5: service.SysAreaPrimarykey
	(*SysAreaBatchDeleteRequest)(nil),  // 6: service.SysAreaBatchDeleteRequest
}
var file_sys_area_model_proto_depIdxs = []int32{
	0, // 0: service.SysAreaListRequest.query:type_name -> service.SysArea
	0, // 1: service.SysAreaResponse.data:type_name -> service.SysArea
	0, // 2: service.SysAreaUpdateFieldsRequest.data:type_name -> service.SysArea
	5, // 3: service.SysAreaBatchDeleteRequest.keys:type_name -> service.SysAreaPrimarykey
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_sys_area_model_proto_init() }
func file_sys_area_model_proto_init() {
	if File_sys_area_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sys_area_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysArea); i {
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
		file_sys_area_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysAreaFilter); i {
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
		file_sys_area_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysAreaListRequest); i {
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
		file_sys_area_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysAreaResponse); i {
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
		file_sys_area_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysAreaUpdateFieldsRequest); i {
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
		file_sys_area_model_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysAreaPrimarykey); i {
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
		file_sys_area_model_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysAreaBatchDeleteRequest); i {
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
			RawDescriptor: file_sys_area_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sys_area_model_proto_goTypes,
		DependencyIndexes: file_sys_area_model_proto_depIdxs,
		MessageInfos:      file_sys_area_model_proto_msgTypes,
	}.Build()
	File_sys_area_model_proto = out.File
	file_sys_area_model_proto_rawDesc = nil
	file_sys_area_model_proto_goTypes = nil
	file_sys_area_model_proto_depIdxs = nil
}
