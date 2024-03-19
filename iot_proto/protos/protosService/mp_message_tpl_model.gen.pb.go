// Code generated by protoc,2022-05-22 07:25:57. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: mp_message_tpl_model.gen.proto

package protosService

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//来源于数据表
type MpMessageTpl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:主键ID;comment:主键ID;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"tplCode" gorm:"default:模板编码;comment:模板编码;size:50;"
	TplCode string `protobuf:"bytes,101,opt,name=tplCode,proto3" json:"tplCode,omitempty"`
	// @inject_tag: json:"tplName" gorm:"default:模板名称;comment:模板名称;size:50;"
	TplName string `protobuf:"bytes,102,opt,name=tplName,proto3" json:"tplName,omitempty"`
	// @inject_tag: json:"tplContent" gorm:"default:模板内容;comment:模板内容;size:255;"
	TplContent string `protobuf:"bytes,103,opt,name=tplContent,proto3" json:"tplContent,omitempty"`
	// @inject_tag: json:"pushType" gorm:"default:发送类型[0:站内且站外，1：站内， 2：站外];comment:发送类型[0:站内且站外，1：站内， 2：站外];size:10;"
	PushType int32 `protobuf:"varint,104,opt,name=pushType,proto3" json:"pushType,omitempty"`
	// @inject_tag: json:"messageType" gorm:"default:消息类型[0：系统提醒，1：设备提醒，....自定义];comment:消息类型[0：系统提醒，1：设备提醒，....自定义];size:10;"
	MessageType int32 `protobuf:"varint,105,opt,name=messageType,proto3" json:"messageType,omitempty"`
	// @inject_tag: json:"agentType" gorm:"default:接收终端类型[0:所有终端， 1：IOS端， 2:android端];comment:接收终端类型[0:所有终端， 1：IOS端， 2:android端];size:10;"
	AgentType int32 `protobuf:"varint,106,opt,name=agentType,proto3" json:"agentType,omitempty"`
	// @inject_tag: json:"lang" gorm:"default:语言编码[zh_CN:简体中文,zh_TW:繁体中文,en_US:英语,es:西班牙语];comment:语言编码[zh_CN:简体中文,zh_TW:繁体中文,en_US:英语,es:西班牙语];size:10;"
	Lang string `protobuf:"bytes,107,opt,name=lang,proto3" json:"lang,omitempty"`
	// @inject_tag: json:"expireHour" gorm:"default:有效时间(单位小时);comment:有效时间(单位小时);size:10;"
	ExpireHour int32 `protobuf:"varint,108,opt,name=expireHour,proto3" json:"expireHour,omitempty"`
	// @inject_tag: json:"createdBy" gorm:"default:创建人;comment:创建人;size:20;"
	CreatedBy int64 `protobuf:"varint,109,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	// @inject_tag: json:"updatedBy" gorm:"default:修改人;comment:修改人;size:20;"
	UpdatedBy int64 `protobuf:"varint,110,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	// @inject_tag: json:"createdAt" gorm:"default:创建时间;comment:创建时间;"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,111,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// @inject_tag: json:"updatedAt" gorm:"default:修改时间;comment:修改时间;"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,112,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// @inject_tag: json:"deletedAt" gorm:"default:删除时间;comment:删除时间;"
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,113,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *MpMessageTpl) Reset() {
	*x = MpMessageTpl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mp_message_tpl_model_gen_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MpMessageTpl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MpMessageTpl) ProtoMessage() {}

func (x *MpMessageTpl) ProtoReflect() protoreflect.Message {
	mi := &file_mp_message_tpl_model_gen_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MpMessageTpl.ProtoReflect.Descriptor instead.
func (*MpMessageTpl) Descriptor() ([]byte, []int) {
	return file_mp_message_tpl_model_gen_proto_rawDescGZIP(), []int{0}
}

func (x *MpMessageTpl) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MpMessageTpl) GetTplCode() string {
	if x != nil {
		return x.TplCode
	}
	return ""
}

func (x *MpMessageTpl) GetTplName() string {
	if x != nil {
		return x.TplName
	}
	return ""
}

func (x *MpMessageTpl) GetTplContent() string {
	if x != nil {
		return x.TplContent
	}
	return ""
}

func (x *MpMessageTpl) GetPushType() int32 {
	if x != nil {
		return x.PushType
	}
	return 0
}

func (x *MpMessageTpl) GetMessageType() int32 {
	if x != nil {
		return x.MessageType
	}
	return 0
}

func (x *MpMessageTpl) GetAgentType() int32 {
	if x != nil {
		return x.AgentType
	}
	return 0
}

func (x *MpMessageTpl) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *MpMessageTpl) GetExpireHour() int32 {
	if x != nil {
		return x.ExpireHour
	}
	return 0
}

func (x *MpMessageTpl) GetCreatedBy() int64 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *MpMessageTpl) GetUpdatedBy() int64 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

func (x *MpMessageTpl) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *MpMessageTpl) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *MpMessageTpl) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type MpMessageTplFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:主键ID;comment:主键ID;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"tplCode" gorm:"default:模板编码;comment:模板编码;size:50;"
	TplCode string `protobuf:"bytes,101,opt,name=tplCode,proto3" json:"tplCode,omitempty"`
	// @inject_tag: json:"tplName" gorm:"default:模板名称;comment:模板名称;size:50;"
	TplName string `protobuf:"bytes,102,opt,name=tplName,proto3" json:"tplName,omitempty"`
	// @inject_tag: json:"tplContent" gorm:"default:模板内容;comment:模板内容;size:255;"
	TplContent string `protobuf:"bytes,103,opt,name=tplContent,proto3" json:"tplContent,omitempty"`
	// @inject_tag: json:"pushType" gorm:"default:发送类型[0:站内且站外，1：站内， 2：站外];comment:发送类型[0:站内且站外，1：站内， 2：站外];size:10;"
	PushType int32 `protobuf:"varint,104,opt,name=pushType,proto3" json:"pushType,omitempty"`
	// @inject_tag: json:"messageType" gorm:"default:消息类型[0：系统提醒，1：设备提醒，....自定义];comment:消息类型[0：系统提醒，1：设备提醒，....自定义];size:10;"
	MessageType int32 `protobuf:"varint,105,opt,name=messageType,proto3" json:"messageType,omitempty"`
	// @inject_tag: json:"agentType" gorm:"default:接收终端类型[0:所有终端， 1：IOS端， 2:android端];comment:接收终端类型[0:所有终端， 1：IOS端， 2:android端];size:10;"
	AgentType int32 `protobuf:"varint,106,opt,name=agentType,proto3" json:"agentType,omitempty"`
	// @inject_tag: json:"lang" gorm:"default:语言编码[zh_CN:简体中文,zh_TW:繁体中文,en_US:英语,es:西班牙语];comment:语言编码[zh_CN:简体中文,zh_TW:繁体中文,en_US:英语,es:西班牙语];size:10;"
	Lang string `protobuf:"bytes,107,opt,name=lang,proto3" json:"lang,omitempty"`
	// @inject_tag: json:"expireHour" gorm:"default:有效时间(单位小时);comment:有效时间(单位小时);size:10;"
	ExpireHour int32 `protobuf:"varint,108,opt,name=expireHour,proto3" json:"expireHour,omitempty"`
	// @inject_tag: json:"createdBy" gorm:"default:创建人;comment:创建人;size:20;"
	CreatedBy int64 `protobuf:"varint,109,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	// @inject_tag: json:"updatedBy" gorm:"default:修改人;comment:修改人;size:20;"
	UpdatedBy int64 `protobuf:"varint,110,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	// @inject_tag: json:"createdAt" gorm:"default:创建时间;comment:创建时间;"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,111,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// @inject_tag: json:"updatedAt" gorm:"default:修改时间;comment:修改时间;"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,112,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// @inject_tag: json:"deletedAt" gorm:"default:删除时间;comment:删除时间;"
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,113,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *MpMessageTplFilter) Reset() {
	*x = MpMessageTplFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mp_message_tpl_model_gen_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MpMessageTplFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MpMessageTplFilter) ProtoMessage() {}

func (x *MpMessageTplFilter) ProtoReflect() protoreflect.Message {
	mi := &file_mp_message_tpl_model_gen_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MpMessageTplFilter.ProtoReflect.Descriptor instead.
func (*MpMessageTplFilter) Descriptor() ([]byte, []int) {
	return file_mp_message_tpl_model_gen_proto_rawDescGZIP(), []int{1}
}

func (x *MpMessageTplFilter) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MpMessageTplFilter) GetTplCode() string {
	if x != nil {
		return x.TplCode
	}
	return ""
}

func (x *MpMessageTplFilter) GetTplName() string {
	if x != nil {
		return x.TplName
	}
	return ""
}

func (x *MpMessageTplFilter) GetTplContent() string {
	if x != nil {
		return x.TplContent
	}
	return ""
}

func (x *MpMessageTplFilter) GetPushType() int32 {
	if x != nil {
		return x.PushType
	}
	return 0
}

func (x *MpMessageTplFilter) GetMessageType() int32 {
	if x != nil {
		return x.MessageType
	}
	return 0
}

func (x *MpMessageTplFilter) GetAgentType() int32 {
	if x != nil {
		return x.AgentType
	}
	return 0
}

func (x *MpMessageTplFilter) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *MpMessageTplFilter) GetExpireHour() int32 {
	if x != nil {
		return x.ExpireHour
	}
	return 0
}

func (x *MpMessageTplFilter) GetCreatedBy() int64 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *MpMessageTplFilter) GetUpdatedBy() int64 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

func (x *MpMessageTplFilter) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *MpMessageTplFilter) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *MpMessageTplFilter) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type MpMessageTplListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query     *MpMessageTpl `protobuf:"bytes,101,opt,name=query,proto3" json:"query,omitempty"`
	Page      int64         `protobuf:"varint,102,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int64         `protobuf:"varint,103,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	OrderKey  string        `protobuf:"bytes,104,opt,name=orderKey,proto3" json:"orderKey,omitempty"`
	OrderDesc string        `protobuf:"bytes,105,opt,name=orderDesc,proto3" json:"orderDesc,omitempty"`
	SearchKey string        `protobuf:"bytes,106,opt,name=searchKey,proto3" json:"searchKey,omitempty"`
}

func (x *MpMessageTplListRequest) Reset() {
	*x = MpMessageTplListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mp_message_tpl_model_gen_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MpMessageTplListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MpMessageTplListRequest) ProtoMessage() {}

func (x *MpMessageTplListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mp_message_tpl_model_gen_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MpMessageTplListRequest.ProtoReflect.Descriptor instead.
func (*MpMessageTplListRequest) Descriptor() ([]byte, []int) {
	return file_mp_message_tpl_model_gen_proto_rawDescGZIP(), []int{2}
}

func (x *MpMessageTplListRequest) GetQuery() *MpMessageTpl {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *MpMessageTplListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *MpMessageTplListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *MpMessageTplListRequest) GetOrderKey() string {
	if x != nil {
		return x.OrderKey
	}
	return ""
}

func (x *MpMessageTplListRequest) GetOrderDesc() string {
	if x != nil {
		return x.OrderDesc
	}
	return ""
}

func (x *MpMessageTplListRequest) GetSearchKey() string {
	if x != nil {
		return x.SearchKey
	}
	return ""
}

type MpMessageTplResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32           `protobuf:"varint,101,opt,name=code,proto3" json:"code,omitempty"`
	Message string          `protobuf:"bytes,102,opt,name=message,proto3" json:"message,omitempty"`
	Total   int64           `protobuf:"varint,103,opt,name=total,proto3" json:"total,omitempty"`
	Data    []*MpMessageTpl `protobuf:"bytes,104,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *MpMessageTplResponse) Reset() {
	*x = MpMessageTplResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mp_message_tpl_model_gen_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MpMessageTplResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MpMessageTplResponse) ProtoMessage() {}

func (x *MpMessageTplResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mp_message_tpl_model_gen_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MpMessageTplResponse.ProtoReflect.Descriptor instead.
func (*MpMessageTplResponse) Descriptor() ([]byte, []int) {
	return file_mp_message_tpl_model_gen_proto_rawDescGZIP(), []int{3}
}

func (x *MpMessageTplResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *MpMessageTplResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MpMessageTplResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *MpMessageTplResponse) GetData() []*MpMessageTpl {
	if x != nil {
		return x.Data
	}
	return nil
}

type MpMessageTplUpdateFieldsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []string      `protobuf:"bytes,101,rep,name=fields,proto3" json:"fields,omitempty"`
	Data   *MpMessageTpl `protobuf:"bytes,102,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *MpMessageTplUpdateFieldsRequest) Reset() {
	*x = MpMessageTplUpdateFieldsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mp_message_tpl_model_gen_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MpMessageTplUpdateFieldsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MpMessageTplUpdateFieldsRequest) ProtoMessage() {}

func (x *MpMessageTplUpdateFieldsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mp_message_tpl_model_gen_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MpMessageTplUpdateFieldsRequest.ProtoReflect.Descriptor instead.
func (*MpMessageTplUpdateFieldsRequest) Descriptor() ([]byte, []int) {
	return file_mp_message_tpl_model_gen_proto_rawDescGZIP(), []int{4}
}

func (x *MpMessageTplUpdateFieldsRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *MpMessageTplUpdateFieldsRequest) GetData() *MpMessageTpl {
	if x != nil {
		return x.Data
	}
	return nil
}

// 表模型主键，只能是整数或字符串，若生成错误，请检查数据库表主键是否正确。支持组合主键
type MpMessageTplPrimarykey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MpMessageTplPrimarykey) Reset() {
	*x = MpMessageTplPrimarykey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mp_message_tpl_model_gen_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MpMessageTplPrimarykey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MpMessageTplPrimarykey) ProtoMessage() {}

func (x *MpMessageTplPrimarykey) ProtoReflect() protoreflect.Message {
	mi := &file_mp_message_tpl_model_gen_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MpMessageTplPrimarykey.ProtoReflect.Descriptor instead.
func (*MpMessageTplPrimarykey) Descriptor() ([]byte, []int) {
	return file_mp_message_tpl_model_gen_proto_rawDescGZIP(), []int{5}
}

func (x *MpMessageTplPrimarykey) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type MpMessageTplBatchDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []*MpMessageTplPrimarykey `protobuf:"bytes,101,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *MpMessageTplBatchDeleteRequest) Reset() {
	*x = MpMessageTplBatchDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mp_message_tpl_model_gen_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MpMessageTplBatchDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MpMessageTplBatchDeleteRequest) ProtoMessage() {}

func (x *MpMessageTplBatchDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mp_message_tpl_model_gen_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MpMessageTplBatchDeleteRequest.ProtoReflect.Descriptor instead.
func (*MpMessageTplBatchDeleteRequest) Descriptor() ([]byte, []int) {
	return file_mp_message_tpl_model_gen_proto_rawDescGZIP(), []int{6}
}

func (x *MpMessageTplBatchDeleteRequest) GetKeys() []*MpMessageTplPrimarykey {
	if x != nil {
		return x.Keys
	}
	return nil
}

var File_mp_message_tpl_model_gen_proto protoreflect.FileDescriptor

var file_mp_message_tpl_model_gen_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6d, 0x70, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x70, 0x6c,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x03, 0x0a, 0x0c, 0x4d,
	0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x70, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x74,
	0x70, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x70,
	0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x70, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x70, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x74, 0x70, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x67, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x70, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x75, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x18, 0x68, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x75, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x69, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c,
	0x61, 0x6e, 0x67, 0x18, 0x6b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12,
	0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x48, 0x6f, 0x75, 0x72, 0x18, 0x6c, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x48, 0x6f, 0x75, 0x72, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x6d, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x6f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x70, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x38, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x71, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xf2, 0x03, 0x0a, 0x12, 0x4d, 0x70,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x70, 0x6c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x74, 0x70, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x74, 0x70, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x70,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x70, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x70, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x67, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x70, 0x6c, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x75, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x68, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x75, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x69, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x6a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x6b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6c, 0x61, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x48, 0x6f,
	0x75, 0x72, 0x18, 0x6c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x48, 0x6f, 0x75, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x18, 0x6d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18,
	0x6e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79,
	0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x6f, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x70, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x71, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xce,
	0x01, 0x0a, 0x17, 0x4d, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x70, 0x6c, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x4d, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x70, 0x6c,
	0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x66, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x4b, 0x65, 0x79, 0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63,
	0x18, 0x69, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x73,
	0x63, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x18, 0x6a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x22,
	0x85, 0x01, 0x0a, 0x14, 0x4d, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x70, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x65, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x29, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x68, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x70,
	0x6c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x64, 0x0a, 0x1f, 0x4d, 0x70, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x54, 0x70, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x70, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x54, 0x70, 0x6c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x28, 0x0a,
	0x16, 0x4d, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x70, 0x6c, 0x50, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x55, 0x0a, 0x1e, 0x4d, 0x70, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x54, 0x70, 0x6c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x04, 0x6b, 0x65, 0x79,
	0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x4d, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x70, 0x6c, 0x50, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x42, 0x11,
	0x5a, 0x0f, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mp_message_tpl_model_gen_proto_rawDescOnce sync.Once
	file_mp_message_tpl_model_gen_proto_rawDescData = file_mp_message_tpl_model_gen_proto_rawDesc
)

func file_mp_message_tpl_model_gen_proto_rawDescGZIP() []byte {
	file_mp_message_tpl_model_gen_proto_rawDescOnce.Do(func() {
		file_mp_message_tpl_model_gen_proto_rawDescData = protoimpl.X.CompressGZIP(file_mp_message_tpl_model_gen_proto_rawDescData)
	})
	return file_mp_message_tpl_model_gen_proto_rawDescData
}

var file_mp_message_tpl_model_gen_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_mp_message_tpl_model_gen_proto_goTypes = []interface{}{
	(*MpMessageTpl)(nil),                    // 0: service.MpMessageTpl
	(*MpMessageTplFilter)(nil),              // 1: service.MpMessageTplFilter
	(*MpMessageTplListRequest)(nil),         // 2: service.MpMessageTplListRequest
	(*MpMessageTplResponse)(nil),            // 3: service.MpMessageTplResponse
	(*MpMessageTplUpdateFieldsRequest)(nil), // 4: service.MpMessageTplUpdateFieldsRequest
	(*MpMessageTplPrimarykey)(nil),          // 5: service.MpMessageTplPrimarykey
	(*MpMessageTplBatchDeleteRequest)(nil),  // 6: service.MpMessageTplBatchDeleteRequest
	(*timestamppb.Timestamp)(nil),           // 7: google.protobuf.Timestamp
}
var file_mp_message_tpl_model_gen_proto_depIdxs = []int32{
	7,  // 0: service.MpMessageTpl.createdAt:type_name -> google.protobuf.Timestamp
	7,  // 1: service.MpMessageTpl.updatedAt:type_name -> google.protobuf.Timestamp
	7,  // 2: service.MpMessageTpl.deletedAt:type_name -> google.protobuf.Timestamp
	7,  // 3: service.MpMessageTplFilter.createdAt:type_name -> google.protobuf.Timestamp
	7,  // 4: service.MpMessageTplFilter.updatedAt:type_name -> google.protobuf.Timestamp
	7,  // 5: service.MpMessageTplFilter.deletedAt:type_name -> google.protobuf.Timestamp
	0,  // 6: service.MpMessageTplListRequest.query:type_name -> service.MpMessageTpl
	0,  // 7: service.MpMessageTplResponse.data:type_name -> service.MpMessageTpl
	0,  // 8: service.MpMessageTplUpdateFieldsRequest.data:type_name -> service.MpMessageTpl
	5,  // 9: service.MpMessageTplBatchDeleteRequest.keys:type_name -> service.MpMessageTplPrimarykey
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_mp_message_tpl_model_gen_proto_init() }
func file_mp_message_tpl_model_gen_proto_init() {
	if File_mp_message_tpl_model_gen_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mp_message_tpl_model_gen_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MpMessageTpl); i {
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
		file_mp_message_tpl_model_gen_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MpMessageTplFilter); i {
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
		file_mp_message_tpl_model_gen_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MpMessageTplListRequest); i {
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
		file_mp_message_tpl_model_gen_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MpMessageTplResponse); i {
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
		file_mp_message_tpl_model_gen_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MpMessageTplUpdateFieldsRequest); i {
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
		file_mp_message_tpl_model_gen_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MpMessageTplPrimarykey); i {
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
		file_mp_message_tpl_model_gen_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MpMessageTplBatchDeleteRequest); i {
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
			RawDescriptor: file_mp_message_tpl_model_gen_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mp_message_tpl_model_gen_proto_goTypes,
		DependencyIndexes: file_mp_message_tpl_model_gen_proto_depIdxs,
		MessageInfos:      file_mp_message_tpl_model_gen_proto_msgTypes,
	}.Build()
	File_mp_message_tpl_model_gen_proto = out.File
	file_mp_message_tpl_model_gen_proto_rawDesc = nil
	file_mp_message_tpl_model_gen_proto_goTypes = nil
	file_mp_message_tpl_model_gen_proto_depIdxs = nil
}
