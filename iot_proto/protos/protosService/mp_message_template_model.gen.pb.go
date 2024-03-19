// Code generated by protoc,2022-05-31 10:02:23. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: mp_message_template_model.gen.proto

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
type MpMessageTemplate struct {
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
	// @inject_tag: json:"tplParams" gorm:"default:模板参数;comment:模板参数;size:64;"
	TplParams string `protobuf:"bytes,104,opt,name=tplParams,proto3" json:"tplParams,omitempty"`
	// @inject_tag: json:"pushType" gorm:"default:消息平台[0:APP消息，1：云管平台消息2， 2：开放平台消息];comment:消息平台[0:APP消息，1：云管平台消息2， 2：开放平台消息];size:10;"
	PushType int32 `protobuf:"varint,105,opt,name=pushType,proto3" json:"pushType,omitempty"`
	// @inject_tag: json:"messageType" gorm:"default:消息类型[0：系统提醒，1：设备提醒，....自定义];comment:消息类型[0：系统提醒，1：设备提醒，....自定义];size:10;"
	MessageType int32 `protobuf:"varint,106,opt,name=messageType,proto3" json:"messageType,omitempty"`
	// @inject_tag: json:"agentType" gorm:"default:接收终端类型[0:所有终端， 1：IOS端， 2:android端];comment:接收终端类型[0:所有终端， 1：IOS端， 2:android端];size:10;"
	AgentType int32 `protobuf:"varint,107,opt,name=agentType,proto3" json:"agentType,omitempty"`
	// @inject_tag: json:"lang" gorm:"default:语言编码[zh_CN:简体中文,zh_TW:繁体中文,en_US:英语,es:西班牙语];comment:语言编码[zh_CN:简体中文,zh_TW:繁体中文,en_US:英语,es:西班牙语];size:10;"
	Lang string `protobuf:"bytes,108,opt,name=lang,proto3" json:"lang,omitempty"`
	// @inject_tag: json:"expireHour" gorm:"default:有效时间(单位小时);comment:有效时间(单位小时);size:10;"
	ExpireHour int32 `protobuf:"varint,109,opt,name=expireHour,proto3" json:"expireHour,omitempty"`
	// @inject_tag: json:"createdBy" gorm:"default:创建人;comment:创建人;size:20;"
	CreatedBy int64 `protobuf:"varint,110,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	// @inject_tag: json:"updatedBy" gorm:"default:修改人;comment:修改人;size:20;"
	UpdatedBy int64 `protobuf:"varint,111,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	// @inject_tag: json:"createdAt" gorm:"default:创建时间;comment:创建时间;"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,112,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// @inject_tag: json:"updatedAt" gorm:"default:修改时间;comment:修改时间;"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,113,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// @inject_tag: json:"deletedAt" gorm:"default:删除时间;comment:删除时间;"
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,114,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *MpMessageTemplate) Reset() {
	*x = MpMessageTemplate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mp_message_template_model_gen_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MpMessageTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MpMessageTemplate) ProtoMessage() {}

func (x *MpMessageTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_mp_message_template_model_gen_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MpMessageTemplate.ProtoReflect.Descriptor instead.
func (*MpMessageTemplate) Descriptor() ([]byte, []int) {
	return file_mp_message_template_model_gen_proto_rawDescGZIP(), []int{0}
}

func (x *MpMessageTemplate) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MpMessageTemplate) GetTplCode() string {
	if x != nil {
		return x.TplCode
	}
	return ""
}

func (x *MpMessageTemplate) GetTplName() string {
	if x != nil {
		return x.TplName
	}
	return ""
}

func (x *MpMessageTemplate) GetTplContent() string {
	if x != nil {
		return x.TplContent
	}
	return ""
}

func (x *MpMessageTemplate) GetTplParams() string {
	if x != nil {
		return x.TplParams
	}
	return ""
}

func (x *MpMessageTemplate) GetPushType() int32 {
	if x != nil {
		return x.PushType
	}
	return 0
}

func (x *MpMessageTemplate) GetMessageType() int32 {
	if x != nil {
		return x.MessageType
	}
	return 0
}

func (x *MpMessageTemplate) GetAgentType() int32 {
	if x != nil {
		return x.AgentType
	}
	return 0
}

func (x *MpMessageTemplate) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *MpMessageTemplate) GetExpireHour() int32 {
	if x != nil {
		return x.ExpireHour
	}
	return 0
}

func (x *MpMessageTemplate) GetCreatedBy() int64 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *MpMessageTemplate) GetUpdatedBy() int64 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

func (x *MpMessageTemplate) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *MpMessageTemplate) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *MpMessageTemplate) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type MpMessageTemplateFilter struct {
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
	// @inject_tag: json:"tplParams" gorm:"default:模板参数;comment:模板参数;size:64;"
	TplParams string `protobuf:"bytes,104,opt,name=tplParams,proto3" json:"tplParams,omitempty"`
	// @inject_tag: json:"pushType" gorm:"default:消息平台[0:APP消息，1：云管平台消息2， 2：开放平台消息];comment:消息平台[0:APP消息，1：云管平台消息2， 2：开放平台消息];size:10;"
	PushType int32 `protobuf:"varint,105,opt,name=pushType,proto3" json:"pushType,omitempty"`
	// @inject_tag: json:"messageType" gorm:"default:消息类型[0：系统提醒，1：设备提醒，....自定义];comment:消息类型[0：系统提醒，1：设备提醒，....自定义];size:10;"
	MessageType int32 `protobuf:"varint,106,opt,name=messageType,proto3" json:"messageType,omitempty"`
	// @inject_tag: json:"agentType" gorm:"default:接收终端类型[0:所有终端， 1：IOS端， 2:android端];comment:接收终端类型[0:所有终端， 1：IOS端， 2:android端];size:10;"
	AgentType int32 `protobuf:"varint,107,opt,name=agentType,proto3" json:"agentType,omitempty"`
	// @inject_tag: json:"lang" gorm:"default:语言编码[zh_CN:简体中文,zh_TW:繁体中文,en_US:英语,es:西班牙语];comment:语言编码[zh_CN:简体中文,zh_TW:繁体中文,en_US:英语,es:西班牙语];size:10;"
	Lang string `protobuf:"bytes,108,opt,name=lang,proto3" json:"lang,omitempty"`
	// @inject_tag: json:"expireHour" gorm:"default:有效时间(单位小时);comment:有效时间(单位小时);size:10;"
	ExpireHour int32 `protobuf:"varint,109,opt,name=expireHour,proto3" json:"expireHour,omitempty"`
	// @inject_tag: json:"createdBy" gorm:"default:创建人;comment:创建人;size:20;"
	CreatedBy int64 `protobuf:"varint,110,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	// @inject_tag: json:"updatedBy" gorm:"default:修改人;comment:修改人;size:20;"
	UpdatedBy int64 `protobuf:"varint,111,opt,name=updatedBy,proto3" json:"updatedBy,omitempty"`
	// @inject_tag: json:"createdAt" gorm:"default:创建时间;comment:创建时间;"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,112,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// @inject_tag: json:"updatedAt" gorm:"default:修改时间;comment:修改时间;"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,113,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// @inject_tag: json:"deletedAt" gorm:"default:删除时间;comment:删除时间;"
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,114,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *MpMessageTemplateFilter) Reset() {
	*x = MpMessageTemplateFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mp_message_template_model_gen_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MpMessageTemplateFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MpMessageTemplateFilter) ProtoMessage() {}

func (x *MpMessageTemplateFilter) ProtoReflect() protoreflect.Message {
	mi := &file_mp_message_template_model_gen_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MpMessageTemplateFilter.ProtoReflect.Descriptor instead.
func (*MpMessageTemplateFilter) Descriptor() ([]byte, []int) {
	return file_mp_message_template_model_gen_proto_rawDescGZIP(), []int{1}
}

func (x *MpMessageTemplateFilter) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MpMessageTemplateFilter) GetTplCode() string {
	if x != nil {
		return x.TplCode
	}
	return ""
}

func (x *MpMessageTemplateFilter) GetTplName() string {
	if x != nil {
		return x.TplName
	}
	return ""
}

func (x *MpMessageTemplateFilter) GetTplContent() string {
	if x != nil {
		return x.TplContent
	}
	return ""
}

func (x *MpMessageTemplateFilter) GetTplParams() string {
	if x != nil {
		return x.TplParams
	}
	return ""
}

func (x *MpMessageTemplateFilter) GetPushType() int32 {
	if x != nil {
		return x.PushType
	}
	return 0
}

func (x *MpMessageTemplateFilter) GetMessageType() int32 {
	if x != nil {
		return x.MessageType
	}
	return 0
}

func (x *MpMessageTemplateFilter) GetAgentType() int32 {
	if x != nil {
		return x.AgentType
	}
	return 0
}

func (x *MpMessageTemplateFilter) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *MpMessageTemplateFilter) GetExpireHour() int32 {
	if x != nil {
		return x.ExpireHour
	}
	return 0
}

func (x *MpMessageTemplateFilter) GetCreatedBy() int64 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *MpMessageTemplateFilter) GetUpdatedBy() int64 {
	if x != nil {
		return x.UpdatedBy
	}
	return 0
}

func (x *MpMessageTemplateFilter) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *MpMessageTemplateFilter) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *MpMessageTemplateFilter) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type MpMessageTemplateListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query     *MpMessageTemplate `protobuf:"bytes,101,opt,name=query,proto3" json:"query,omitempty"`
	Page      int64              `protobuf:"varint,102,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int64              `protobuf:"varint,103,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	OrderKey  string             `protobuf:"bytes,104,opt,name=orderKey,proto3" json:"orderKey,omitempty"`
	OrderDesc string             `protobuf:"bytes,105,opt,name=orderDesc,proto3" json:"orderDesc,omitempty"`
	SearchKey string             `protobuf:"bytes,106,opt,name=searchKey,proto3" json:"searchKey,omitempty"`
}

func (x *MpMessageTemplateListRequest) Reset() {
	*x = MpMessageTemplateListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mp_message_template_model_gen_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MpMessageTemplateListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MpMessageTemplateListRequest) ProtoMessage() {}

func (x *MpMessageTemplateListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mp_message_template_model_gen_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MpMessageTemplateListRequest.ProtoReflect.Descriptor instead.
func (*MpMessageTemplateListRequest) Descriptor() ([]byte, []int) {
	return file_mp_message_template_model_gen_proto_rawDescGZIP(), []int{2}
}

func (x *MpMessageTemplateListRequest) GetQuery() *MpMessageTemplate {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *MpMessageTemplateListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *MpMessageTemplateListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *MpMessageTemplateListRequest) GetOrderKey() string {
	if x != nil {
		return x.OrderKey
	}
	return ""
}

func (x *MpMessageTemplateListRequest) GetOrderDesc() string {
	if x != nil {
		return x.OrderDesc
	}
	return ""
}

func (x *MpMessageTemplateListRequest) GetSearchKey() string {
	if x != nil {
		return x.SearchKey
	}
	return ""
}

type MpMessageTemplateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32                `protobuf:"varint,101,opt,name=code,proto3" json:"code,omitempty"`
	Message string               `protobuf:"bytes,102,opt,name=message,proto3" json:"message,omitempty"`
	Total   int64                `protobuf:"varint,103,opt,name=total,proto3" json:"total,omitempty"`
	Data    []*MpMessageTemplate `protobuf:"bytes,104,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *MpMessageTemplateResponse) Reset() {
	*x = MpMessageTemplateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mp_message_template_model_gen_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MpMessageTemplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MpMessageTemplateResponse) ProtoMessage() {}

func (x *MpMessageTemplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mp_message_template_model_gen_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MpMessageTemplateResponse.ProtoReflect.Descriptor instead.
func (*MpMessageTemplateResponse) Descriptor() ([]byte, []int) {
	return file_mp_message_template_model_gen_proto_rawDescGZIP(), []int{3}
}

func (x *MpMessageTemplateResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *MpMessageTemplateResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MpMessageTemplateResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *MpMessageTemplateResponse) GetData() []*MpMessageTemplate {
	if x != nil {
		return x.Data
	}
	return nil
}

type MpMessageTemplateUpdateFieldsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []string           `protobuf:"bytes,101,rep,name=fields,proto3" json:"fields,omitempty"`
	Data   *MpMessageTemplate `protobuf:"bytes,102,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *MpMessageTemplateUpdateFieldsRequest) Reset() {
	*x = MpMessageTemplateUpdateFieldsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mp_message_template_model_gen_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MpMessageTemplateUpdateFieldsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MpMessageTemplateUpdateFieldsRequest) ProtoMessage() {}

func (x *MpMessageTemplateUpdateFieldsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mp_message_template_model_gen_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MpMessageTemplateUpdateFieldsRequest.ProtoReflect.Descriptor instead.
func (*MpMessageTemplateUpdateFieldsRequest) Descriptor() ([]byte, []int) {
	return file_mp_message_template_model_gen_proto_rawDescGZIP(), []int{4}
}

func (x *MpMessageTemplateUpdateFieldsRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *MpMessageTemplateUpdateFieldsRequest) GetData() *MpMessageTemplate {
	if x != nil {
		return x.Data
	}
	return nil
}

// 表模型主键，只能是整数或字符串，若生成错误，请检查数据库表主键是否正确。支持组合主键
type MpMessageTemplatePrimarykey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MpMessageTemplatePrimarykey) Reset() {
	*x = MpMessageTemplatePrimarykey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mp_message_template_model_gen_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MpMessageTemplatePrimarykey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MpMessageTemplatePrimarykey) ProtoMessage() {}

func (x *MpMessageTemplatePrimarykey) ProtoReflect() protoreflect.Message {
	mi := &file_mp_message_template_model_gen_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MpMessageTemplatePrimarykey.ProtoReflect.Descriptor instead.
func (*MpMessageTemplatePrimarykey) Descriptor() ([]byte, []int) {
	return file_mp_message_template_model_gen_proto_rawDescGZIP(), []int{5}
}

func (x *MpMessageTemplatePrimarykey) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type MpMessageTemplateBatchDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []*MpMessageTemplatePrimarykey `protobuf:"bytes,101,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *MpMessageTemplateBatchDeleteRequest) Reset() {
	*x = MpMessageTemplateBatchDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mp_message_template_model_gen_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MpMessageTemplateBatchDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MpMessageTemplateBatchDeleteRequest) ProtoMessage() {}

func (x *MpMessageTemplateBatchDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mp_message_template_model_gen_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MpMessageTemplateBatchDeleteRequest.ProtoReflect.Descriptor instead.
func (*MpMessageTemplateBatchDeleteRequest) Descriptor() ([]byte, []int) {
	return file_mp_message_template_model_gen_proto_rawDescGZIP(), []int{6}
}

func (x *MpMessageTemplateBatchDeleteRequest) GetKeys() []*MpMessageTemplatePrimarykey {
	if x != nil {
		return x.Keys
	}
	return nil
}

var File_mp_message_template_model_gen_proto protoreflect.FileDescriptor

var file_mp_message_template_model_gen_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6d, 0x70, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x67, 0x65, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x8f, 0x04, 0x0a, 0x11, 0x4d, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x70, 0x6c, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x70, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x74, 0x70, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x74, 0x70, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x70, 0x6c,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x67, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74,
	0x70, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x70, 0x6c,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x70,
	0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x75, 0x73, 0x68, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x69, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x75, 0x73, 0x68, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x6b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x6c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x48, 0x6f, 0x75, 0x72, 0x18, 0x6d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x48, 0x6f, 0x75, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x42, 0x79, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x18, 0x6f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x70, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x71, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x72, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x95, 0x04, 0x0a, 0x17, 0x4d, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x74, 0x70, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x74, 0x70, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x70, 0x6c, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x70, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x70, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x67, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x70, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x70, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x68,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x70, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x75, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x18, 0x69, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x75, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x6b, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c,
	0x61, 0x6e, 0x67, 0x18, 0x6c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12,
	0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x48, 0x6f, 0x75, 0x72, 0x18, 0x6d, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x48, 0x6f, 0x75, 0x72, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x6e, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x6f, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x70, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x71, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x38, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x72, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xd8, 0x01, 0x0a, 0x1c, 0x4d, 0x70,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x4d, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x67, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x44, 0x65, 0x73, 0x63, 0x18, 0x69, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x4b, 0x65, 0x79, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x4b, 0x65, 0x79, 0x22, 0x8f, 0x01, 0x0a, 0x19, 0x4d, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x68,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d,
	0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x6e, 0x0a, 0x24, 0x4d, 0x70, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x66,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d,
	0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2d, 0x0a, 0x1b, 0x4d, 0x70, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x6b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5f, 0x0a, 0x23, 0x4d, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x04,
	0x6b, 0x65, 0x79, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79,
	0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_mp_message_template_model_gen_proto_rawDescOnce sync.Once
	file_mp_message_template_model_gen_proto_rawDescData = file_mp_message_template_model_gen_proto_rawDesc
)

func file_mp_message_template_model_gen_proto_rawDescGZIP() []byte {
	file_mp_message_template_model_gen_proto_rawDescOnce.Do(func() {
		file_mp_message_template_model_gen_proto_rawDescData = protoimpl.X.CompressGZIP(file_mp_message_template_model_gen_proto_rawDescData)
	})
	return file_mp_message_template_model_gen_proto_rawDescData
}

var file_mp_message_template_model_gen_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_mp_message_template_model_gen_proto_goTypes = []interface{}{
	(*MpMessageTemplate)(nil),                    // 0: service.MpMessageTemplate
	(*MpMessageTemplateFilter)(nil),              // 1: service.MpMessageTemplateFilter
	(*MpMessageTemplateListRequest)(nil),         // 2: service.MpMessageTemplateListRequest
	(*MpMessageTemplateResponse)(nil),            // 3: service.MpMessageTemplateResponse
	(*MpMessageTemplateUpdateFieldsRequest)(nil), // 4: service.MpMessageTemplateUpdateFieldsRequest
	(*MpMessageTemplatePrimarykey)(nil),          // 5: service.MpMessageTemplatePrimarykey
	(*MpMessageTemplateBatchDeleteRequest)(nil),  // 6: service.MpMessageTemplateBatchDeleteRequest
	(*timestamppb.Timestamp)(nil),                // 7: google.protobuf.Timestamp
}
var file_mp_message_template_model_gen_proto_depIdxs = []int32{
	7,  // 0: service.MpMessageTemplate.createdAt:type_name -> google.protobuf.Timestamp
	7,  // 1: service.MpMessageTemplate.updatedAt:type_name -> google.protobuf.Timestamp
	7,  // 2: service.MpMessageTemplate.deletedAt:type_name -> google.protobuf.Timestamp
	7,  // 3: service.MpMessageTemplateFilter.createdAt:type_name -> google.protobuf.Timestamp
	7,  // 4: service.MpMessageTemplateFilter.updatedAt:type_name -> google.protobuf.Timestamp
	7,  // 5: service.MpMessageTemplateFilter.deletedAt:type_name -> google.protobuf.Timestamp
	0,  // 6: service.MpMessageTemplateListRequest.query:type_name -> service.MpMessageTemplate
	0,  // 7: service.MpMessageTemplateResponse.data:type_name -> service.MpMessageTemplate
	0,  // 8: service.MpMessageTemplateUpdateFieldsRequest.data:type_name -> service.MpMessageTemplate
	5,  // 9: service.MpMessageTemplateBatchDeleteRequest.keys:type_name -> service.MpMessageTemplatePrimarykey
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_mp_message_template_model_gen_proto_init() }
func file_mp_message_template_model_gen_proto_init() {
	if File_mp_message_template_model_gen_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mp_message_template_model_gen_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MpMessageTemplate); i {
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
		file_mp_message_template_model_gen_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MpMessageTemplateFilter); i {
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
		file_mp_message_template_model_gen_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MpMessageTemplateListRequest); i {
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
		file_mp_message_template_model_gen_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MpMessageTemplateResponse); i {
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
		file_mp_message_template_model_gen_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MpMessageTemplateUpdateFieldsRequest); i {
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
		file_mp_message_template_model_gen_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MpMessageTemplatePrimarykey); i {
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
		file_mp_message_template_model_gen_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MpMessageTemplateBatchDeleteRequest); i {
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
			RawDescriptor: file_mp_message_template_model_gen_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mp_message_template_model_gen_proto_goTypes,
		DependencyIndexes: file_mp_message_template_model_gen_proto_depIdxs,
		MessageInfos:      file_mp_message_template_model_gen_proto_msgTypes,
	}.Build()
	File_mp_message_template_model_gen_proto = out.File
	file_mp_message_template_model_gen_proto_rawDesc = nil
	file_mp_message_template_model_gen_proto_goTypes = nil
	file_mp_message_template_model_gen_proto_depIdxs = nil
}
