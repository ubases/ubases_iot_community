// Code generated by protoc,2022-10-27 18:44:20. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: opm_voice_product_map_model.gen.proto

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
type OpmVoiceProductMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:语控属性配置id;语控属性配置id;comment:语控属性配置id;语控属性配置id;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"voiceProductId" gorm:"default:产品语控配置id;comment:产品语控配置id;size:19;"
	VoiceProductId int64 `protobuf:"varint,101,opt,name=voiceProductId,proto3" json:"voiceProductId,omitempty"`
	// @inject_tag: json:"voiceNo" gorm:"default:;comment:;size:50;"
	VoiceNo string `protobuf:"bytes,102,opt,name=voiceNo,proto3" json:"voiceNo,omitempty"`
	// @inject_tag: json:"funName" gorm:"default:功能名称;comment:功能名称;size:100;"
	FunName string `protobuf:"bytes,103,opt,name=funName,proto3" json:"funName,omitempty"`
	// @inject_tag: json:"attrCode" gorm:"default:属性标识编码;comment:属性标识编码;size:100;"
	AttrCode string `protobuf:"bytes,104,opt,name=attrCode,proto3" json:"attrCode,omitempty"`
	// @inject_tag: json:"voiceCode" gorm:"default:语控标识编码;comment:语控标识编码;size:100;"
	VoiceCode string `protobuf:"bytes,105,opt,name=voiceCode,proto3" json:"voiceCode,omitempty"`
	// @inject_tag: json:"valueType" gorm:"default:值类型;1 数值范围  2 数值列表 3 字符串类型;comment:值类型;1 数值范围  2 数值列表 3 字符串类型;size:10;"
	ValueType int32 `protobuf:"varint,106,opt,name=valueType,proto3" json:"valueType,omitempty"`
	// @inject_tag: json:"voiceDemo" gorm:"default:语控示例;comment:语控示例;size:2000;"
	VoiceDemo string `protobuf:"bytes,107,opt,name=voiceDemo,proto3" json:"voiceDemo,omitempty"`
	// @inject_tag: json:"attrUnit" gorm:"default:属性单位;comment:属性单位;size:50;"
	AttrUnit string `protobuf:"bytes,108,opt,name=attrUnit,proto3" json:"attrUnit,omitempty"`
	// @inject_tag: json:"voiceAttrUnit" gorm:"default:语控属性单位;comment:语控属性单位;size:50;"
	VoiceAttrUnit string `protobuf:"bytes,109,opt,name=voiceAttrUnit,proto3" json:"voiceAttrUnit,omitempty"`
	// @inject_tag: json:"attrDpid" gorm:"default:功能的dpid;comment:功能的dpid;size:10;"
	AttrDpid int32 `protobuf:"varint,110,opt,name=attrDpid,proto3" json:"attrDpid,omitempty"`
	// @inject_tag: json:"valueMap" gorm:"default:属性映射字符串;comment:属性映射字符串;"
	ValueMap string `protobuf:"bytes,111,opt,name=valueMap,proto3" json:"valueMap,omitempty"`
	// @inject_tag: json:"createdAt" gorm:"default:新增时间;comment:新增时间;"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,112,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// @inject_tag: json:"dataType" gorm:"default:类型有INT、DOUBLE、TEXT、ENUM、BOOL等;comment:类型有INT、DOUBLE、TEXT、ENUM、BOOL等;size:64;"
	DataType string `protobuf:"bytes,113,opt,name=dataType,proto3" json:"dataType,omitempty"`
	// @inject_tag: json:"vDataType" gorm:"default:类型有INT、DOUBLE、TEXT、ENUM、BOOL等;comment:类型有INT、DOUBLE、TEXT、ENUM、BOOL等;size:64;"
	VDataType    string `protobuf:"bytes,114,opt,name=vDataType,proto3" json:"vDataType,omitempty"`
	ProductKey   string `protobuf:"bytes,115,opt,name=productKey,proto3" json:"productKey,omitempty"`
	ProductId    int64  `protobuf:"varint,116,opt,name=productId,proto3" json:"productId,omitempty"`
	VoiceSynonym string `protobuf:"bytes,117,opt,name=voiceSynonym,proto3" json:"voiceSynonym,omitempty"`
	Trait        string `protobuf:"bytes,118,opt,name=trait,proto3" json:"trait,omitempty"`
	Command      string `protobuf:"bytes,119,opt,name=command,proto3" json:"command,omitempty"`
	VoiceOther   string `protobuf:"bytes,120,opt,name=voiceOther,proto3" json:"voiceOther,omitempty"` //其它参数查询
}

func (x *OpmVoiceProductMap) Reset() {
	*x = OpmVoiceProductMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opm_voice_product_map_model_gen_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpmVoiceProductMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpmVoiceProductMap) ProtoMessage() {}

func (x *OpmVoiceProductMap) ProtoReflect() protoreflect.Message {
	mi := &file_opm_voice_product_map_model_gen_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpmVoiceProductMap.ProtoReflect.Descriptor instead.
func (*OpmVoiceProductMap) Descriptor() ([]byte, []int) {
	return file_opm_voice_product_map_model_gen_proto_rawDescGZIP(), []int{0}
}

func (x *OpmVoiceProductMap) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OpmVoiceProductMap) GetVoiceProductId() int64 {
	if x != nil {
		return x.VoiceProductId
	}
	return 0
}

func (x *OpmVoiceProductMap) GetVoiceNo() string {
	if x != nil {
		return x.VoiceNo
	}
	return ""
}

func (x *OpmVoiceProductMap) GetFunName() string {
	if x != nil {
		return x.FunName
	}
	return ""
}

func (x *OpmVoiceProductMap) GetAttrCode() string {
	if x != nil {
		return x.AttrCode
	}
	return ""
}

func (x *OpmVoiceProductMap) GetVoiceCode() string {
	if x != nil {
		return x.VoiceCode
	}
	return ""
}

func (x *OpmVoiceProductMap) GetValueType() int32 {
	if x != nil {
		return x.ValueType
	}
	return 0
}

func (x *OpmVoiceProductMap) GetVoiceDemo() string {
	if x != nil {
		return x.VoiceDemo
	}
	return ""
}

func (x *OpmVoiceProductMap) GetAttrUnit() string {
	if x != nil {
		return x.AttrUnit
	}
	return ""
}

func (x *OpmVoiceProductMap) GetVoiceAttrUnit() string {
	if x != nil {
		return x.VoiceAttrUnit
	}
	return ""
}

func (x *OpmVoiceProductMap) GetAttrDpid() int32 {
	if x != nil {
		return x.AttrDpid
	}
	return 0
}

func (x *OpmVoiceProductMap) GetValueMap() string {
	if x != nil {
		return x.ValueMap
	}
	return ""
}

func (x *OpmVoiceProductMap) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *OpmVoiceProductMap) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *OpmVoiceProductMap) GetVDataType() string {
	if x != nil {
		return x.VDataType
	}
	return ""
}

func (x *OpmVoiceProductMap) GetProductKey() string {
	if x != nil {
		return x.ProductKey
	}
	return ""
}

func (x *OpmVoiceProductMap) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *OpmVoiceProductMap) GetVoiceSynonym() string {
	if x != nil {
		return x.VoiceSynonym
	}
	return ""
}

func (x *OpmVoiceProductMap) GetTrait() string {
	if x != nil {
		return x.Trait
	}
	return ""
}

func (x *OpmVoiceProductMap) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *OpmVoiceProductMap) GetVoiceOther() string {
	if x != nil {
		return x.VoiceOther
	}
	return ""
}

type OpmVoiceProductMapFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:语控属性配置id;语控属性配置id;comment:语控属性配置id;语控属性配置id;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"voiceProductId" gorm:"default:产品语控配置id;comment:产品语控配置id;size:19;"
	VoiceProductId int64 `protobuf:"varint,101,opt,name=voiceProductId,proto3" json:"voiceProductId,omitempty"`
	// @inject_tag: json:"voiceNo" gorm:"default:;comment:;size:50;"
	VoiceNo string `protobuf:"bytes,102,opt,name=voiceNo,proto3" json:"voiceNo,omitempty"`
	// @inject_tag: json:"funName" gorm:"default:功能名称;comment:功能名称;size:100;"
	FunName string `protobuf:"bytes,103,opt,name=funName,proto3" json:"funName,omitempty"`
	// @inject_tag: json:"attrCode" gorm:"default:属性标识编码;comment:属性标识编码;size:100;"
	AttrCode string `protobuf:"bytes,104,opt,name=attrCode,proto3" json:"attrCode,omitempty"`
	// @inject_tag: json:"voiceCode" gorm:"default:语控标识编码;comment:语控标识编码;size:100;"
	VoiceCode string `protobuf:"bytes,105,opt,name=voiceCode,proto3" json:"voiceCode,omitempty"`
	// @inject_tag: json:"valueType" gorm:"default:值类型;1 数值范围  2 数值列表 3 字符串类型;comment:值类型;1 数值范围  2 数值列表 3 字符串类型;size:10;"
	ValueType int32 `protobuf:"varint,106,opt,name=valueType,proto3" json:"valueType,omitempty"`
	// @inject_tag: json:"voiceDemo" gorm:"default:语控示例;comment:语控示例;size:2000;"
	VoiceDemo string `protobuf:"bytes,107,opt,name=voiceDemo,proto3" json:"voiceDemo,omitempty"`
	// @inject_tag: json:"attrUnit" gorm:"default:属性单位;comment:属性单位;size:50;"
	AttrUnit string `protobuf:"bytes,108,opt,name=attrUnit,proto3" json:"attrUnit,omitempty"`
	// @inject_tag: json:"voiceAttrUnit" gorm:"default:语控属性单位;comment:语控属性单位;size:50;"
	VoiceAttrUnit string `protobuf:"bytes,109,opt,name=voiceAttrUnit,proto3" json:"voiceAttrUnit,omitempty"`
	// @inject_tag: json:"attrDpid" gorm:"default:功能的dpid;comment:功能的dpid;size:10;"
	AttrDpid int32 `protobuf:"varint,110,opt,name=attrDpid,proto3" json:"attrDpid,omitempty"`
	// @inject_tag: json:"valueMap" gorm:"default:属性映射字符串;comment:属性映射字符串;"
	ValueMap string `protobuf:"bytes,111,opt,name=valueMap,proto3" json:"valueMap,omitempty"`
	// @inject_tag: json:"createdAt" gorm:"default:新增时间;comment:新增时间;"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,112,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// @inject_tag: json:"dataType" gorm:"default:类型有INT、DOUBLE、TEXT、ENUM、BOOL等;comment:类型有INT、DOUBLE、TEXT、ENUM、BOOL等;size:64;"
	DataType string `protobuf:"bytes,113,opt,name=dataType,proto3" json:"dataType,omitempty"`
	// @inject_tag: json:"vDataType" gorm:"default:类型有INT、DOUBLE、TEXT、ENUM、BOOL等;comment:类型有INT、DOUBLE、TEXT、ENUM、BOOL等;size:64;"
	VDataType string `protobuf:"bytes,114,opt,name=vDataType,proto3" json:"vDataType,omitempty"`
}

func (x *OpmVoiceProductMapFilter) Reset() {
	*x = OpmVoiceProductMapFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opm_voice_product_map_model_gen_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpmVoiceProductMapFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpmVoiceProductMapFilter) ProtoMessage() {}

func (x *OpmVoiceProductMapFilter) ProtoReflect() protoreflect.Message {
	mi := &file_opm_voice_product_map_model_gen_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpmVoiceProductMapFilter.ProtoReflect.Descriptor instead.
func (*OpmVoiceProductMapFilter) Descriptor() ([]byte, []int) {
	return file_opm_voice_product_map_model_gen_proto_rawDescGZIP(), []int{1}
}

func (x *OpmVoiceProductMapFilter) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OpmVoiceProductMapFilter) GetVoiceProductId() int64 {
	if x != nil {
		return x.VoiceProductId
	}
	return 0
}

func (x *OpmVoiceProductMapFilter) GetVoiceNo() string {
	if x != nil {
		return x.VoiceNo
	}
	return ""
}

func (x *OpmVoiceProductMapFilter) GetFunName() string {
	if x != nil {
		return x.FunName
	}
	return ""
}

func (x *OpmVoiceProductMapFilter) GetAttrCode() string {
	if x != nil {
		return x.AttrCode
	}
	return ""
}

func (x *OpmVoiceProductMapFilter) GetVoiceCode() string {
	if x != nil {
		return x.VoiceCode
	}
	return ""
}

func (x *OpmVoiceProductMapFilter) GetValueType() int32 {
	if x != nil {
		return x.ValueType
	}
	return 0
}

func (x *OpmVoiceProductMapFilter) GetVoiceDemo() string {
	if x != nil {
		return x.VoiceDemo
	}
	return ""
}

func (x *OpmVoiceProductMapFilter) GetAttrUnit() string {
	if x != nil {
		return x.AttrUnit
	}
	return ""
}

func (x *OpmVoiceProductMapFilter) GetVoiceAttrUnit() string {
	if x != nil {
		return x.VoiceAttrUnit
	}
	return ""
}

func (x *OpmVoiceProductMapFilter) GetAttrDpid() int32 {
	if x != nil {
		return x.AttrDpid
	}
	return 0
}

func (x *OpmVoiceProductMapFilter) GetValueMap() string {
	if x != nil {
		return x.ValueMap
	}
	return ""
}

func (x *OpmVoiceProductMapFilter) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *OpmVoiceProductMapFilter) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *OpmVoiceProductMapFilter) GetVDataType() string {
	if x != nil {
		return x.VDataType
	}
	return ""
}

type OpmVoiceProductMapListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query     *OpmVoiceProductMap `protobuf:"bytes,101,opt,name=query,proto3" json:"query,omitempty"`
	Page      int64               `protobuf:"varint,102,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int64               `protobuf:"varint,103,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	OrderKey  string              `protobuf:"bytes,104,opt,name=orderKey,proto3" json:"orderKey,omitempty"`
	OrderDesc string              `protobuf:"bytes,105,opt,name=orderDesc,proto3" json:"orderDesc,omitempty"`
	SearchKey string              `protobuf:"bytes,106,opt,name=searchKey,proto3" json:"searchKey,omitempty"`
}

func (x *OpmVoiceProductMapListRequest) Reset() {
	*x = OpmVoiceProductMapListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opm_voice_product_map_model_gen_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpmVoiceProductMapListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpmVoiceProductMapListRequest) ProtoMessage() {}

func (x *OpmVoiceProductMapListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_opm_voice_product_map_model_gen_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpmVoiceProductMapListRequest.ProtoReflect.Descriptor instead.
func (*OpmVoiceProductMapListRequest) Descriptor() ([]byte, []int) {
	return file_opm_voice_product_map_model_gen_proto_rawDescGZIP(), []int{2}
}

func (x *OpmVoiceProductMapListRequest) GetQuery() *OpmVoiceProductMap {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *OpmVoiceProductMapListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *OpmVoiceProductMapListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *OpmVoiceProductMapListRequest) GetOrderKey() string {
	if x != nil {
		return x.OrderKey
	}
	return ""
}

func (x *OpmVoiceProductMapListRequest) GetOrderDesc() string {
	if x != nil {
		return x.OrderDesc
	}
	return ""
}

func (x *OpmVoiceProductMapListRequest) GetSearchKey() string {
	if x != nil {
		return x.SearchKey
	}
	return ""
}

type OpmVoiceProductMapResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32                 `protobuf:"varint,101,opt,name=code,proto3" json:"code,omitempty"`
	Message string                `protobuf:"bytes,102,opt,name=message,proto3" json:"message,omitempty"`
	Total   int64                 `protobuf:"varint,103,opt,name=total,proto3" json:"total,omitempty"`
	Data    []*OpmVoiceProductMap `protobuf:"bytes,104,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *OpmVoiceProductMapResponse) Reset() {
	*x = OpmVoiceProductMapResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opm_voice_product_map_model_gen_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpmVoiceProductMapResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpmVoiceProductMapResponse) ProtoMessage() {}

func (x *OpmVoiceProductMapResponse) ProtoReflect() protoreflect.Message {
	mi := &file_opm_voice_product_map_model_gen_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpmVoiceProductMapResponse.ProtoReflect.Descriptor instead.
func (*OpmVoiceProductMapResponse) Descriptor() ([]byte, []int) {
	return file_opm_voice_product_map_model_gen_proto_rawDescGZIP(), []int{3}
}

func (x *OpmVoiceProductMapResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *OpmVoiceProductMapResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *OpmVoiceProductMapResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *OpmVoiceProductMapResponse) GetData() []*OpmVoiceProductMap {
	if x != nil {
		return x.Data
	}
	return nil
}

type OpmVoiceProductMapUpdateFieldsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []string            `protobuf:"bytes,101,rep,name=fields,proto3" json:"fields,omitempty"`
	Data   *OpmVoiceProductMap `protobuf:"bytes,102,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OpmVoiceProductMapUpdateFieldsRequest) Reset() {
	*x = OpmVoiceProductMapUpdateFieldsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opm_voice_product_map_model_gen_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpmVoiceProductMapUpdateFieldsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpmVoiceProductMapUpdateFieldsRequest) ProtoMessage() {}

func (x *OpmVoiceProductMapUpdateFieldsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_opm_voice_product_map_model_gen_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpmVoiceProductMapUpdateFieldsRequest.ProtoReflect.Descriptor instead.
func (*OpmVoiceProductMapUpdateFieldsRequest) Descriptor() ([]byte, []int) {
	return file_opm_voice_product_map_model_gen_proto_rawDescGZIP(), []int{4}
}

func (x *OpmVoiceProductMapUpdateFieldsRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *OpmVoiceProductMapUpdateFieldsRequest) GetData() *OpmVoiceProductMap {
	if x != nil {
		return x.Data
	}
	return nil
}

// 表模型主键，只能是整数或字符串，若生成错误，请检查数据库表主键是否正确。支持组合主键
type OpmVoiceProductMapPrimarykey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OpmVoiceProductMapPrimarykey) Reset() {
	*x = OpmVoiceProductMapPrimarykey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opm_voice_product_map_model_gen_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpmVoiceProductMapPrimarykey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpmVoiceProductMapPrimarykey) ProtoMessage() {}

func (x *OpmVoiceProductMapPrimarykey) ProtoReflect() protoreflect.Message {
	mi := &file_opm_voice_product_map_model_gen_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpmVoiceProductMapPrimarykey.ProtoReflect.Descriptor instead.
func (*OpmVoiceProductMapPrimarykey) Descriptor() ([]byte, []int) {
	return file_opm_voice_product_map_model_gen_proto_rawDescGZIP(), []int{5}
}

func (x *OpmVoiceProductMapPrimarykey) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type OpmVoiceProductMapBatchDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []*OpmVoiceProductMapPrimarykey `protobuf:"bytes,101,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *OpmVoiceProductMapBatchDeleteRequest) Reset() {
	*x = OpmVoiceProductMapBatchDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opm_voice_product_map_model_gen_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpmVoiceProductMapBatchDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpmVoiceProductMapBatchDeleteRequest) ProtoMessage() {}

func (x *OpmVoiceProductMapBatchDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_opm_voice_product_map_model_gen_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpmVoiceProductMapBatchDeleteRequest.ProtoReflect.Descriptor instead.
func (*OpmVoiceProductMapBatchDeleteRequest) Descriptor() ([]byte, []int) {
	return file_opm_voice_product_map_model_gen_proto_rawDescGZIP(), []int{6}
}

func (x *OpmVoiceProductMapBatchDeleteRequest) GetKeys() []*OpmVoiceProductMapPrimarykey {
	if x != nil {
		return x.Keys
	}
	return nil
}

var File_opm_voice_product_map_model_gen_proto protoreflect.FileDescriptor

var file_opm_voice_product_map_model_gen_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6f, 0x70, 0x6d, 0x5f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x70, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x67, 0x65,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x96, 0x05, 0x0a, 0x12, 0x4f, 0x70, 0x6d, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x61, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x65, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x4e, 0x6f, 0x18, 0x66, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x4e, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x75,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x75, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x74, 0x74, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x74, 0x74, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x69, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x6a, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x18, 0x6b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x74,
	0x74, 0x72, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x6c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x74,
	0x74, 0x72, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x41,
	0x74, 0x74, 0x72, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x6d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x41, 0x74, 0x74, 0x72, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x61, 0x74, 0x74, 0x72, 0x44, 0x70, 0x69, 0x64, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x61, 0x74, 0x74, 0x72, 0x44, 0x70, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x4d, 0x61, 0x70, 0x18, 0x6f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x4d, 0x61, 0x70, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x70, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x71, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x44,
	0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x72, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76,
	0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x73, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x74, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x53,
	0x79, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x18, 0x75, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x53, 0x79, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x72,
	0x61, 0x69, 0x74, 0x18, 0x76, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x72, 0x61, 0x69, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x77, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x18, 0x78, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x22, 0xea, 0x03, 0x0a, 0x18, 0x4f,
	0x70, 0x6d, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x61,
	0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x65, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x4e, 0x6f, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x4e, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x75, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x75, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x74, 0x74, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x74, 0x74, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x69, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x18, 0x6b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x74, 0x74,
	0x72, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x6c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x74, 0x74,
	0x72, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x41, 0x74,
	0x74, 0x72, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x6d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x41, 0x74, 0x74, 0x72, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x61,
	0x74, 0x74, 0x72, 0x44, 0x70, 0x69, 0x64, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61,
	0x74, 0x74, 0x72, 0x44, 0x70, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x4d, 0x61, 0x70, 0x18, 0x6f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x4d, 0x61, 0x70, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x70, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x71, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x44, 0x61,
	0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x72, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x44,
	0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x22, 0xda, 0x01, 0x0a, 0x1d, 0x4f, 0x70, 0x6d, 0x56,
	0x6f, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x61, 0x70, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4d, 0x61, 0x70, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x67, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x44, 0x65, 0x73, 0x63, 0x18, 0x69, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x4b, 0x65, 0x79, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x4b, 0x65, 0x79, 0x22, 0x91, 0x01, 0x0a, 0x1a, 0x4f, 0x70, 0x6d, 0x56, 0x6f, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x68, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x4f, 0x70, 0x6d, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d,
	0x61, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x70, 0x0a, 0x25, 0x4f, 0x70, 0x6d, 0x56,
	0x6f, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x61, 0x70, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4d, 0x61, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2e, 0x0a, 0x1c, 0x4f, 0x70,
	0x6d, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x61, 0x70,
	0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x61, 0x0a, 0x24, 0x4f, 0x70,
	0x6d, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x61, 0x70,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x39, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x56, 0x6f,
	0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x61, 0x70, 0x50, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x42, 0x11, 0x5a,
	0x0f, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_opm_voice_product_map_model_gen_proto_rawDescOnce sync.Once
	file_opm_voice_product_map_model_gen_proto_rawDescData = file_opm_voice_product_map_model_gen_proto_rawDesc
)

func file_opm_voice_product_map_model_gen_proto_rawDescGZIP() []byte {
	file_opm_voice_product_map_model_gen_proto_rawDescOnce.Do(func() {
		file_opm_voice_product_map_model_gen_proto_rawDescData = protoimpl.X.CompressGZIP(file_opm_voice_product_map_model_gen_proto_rawDescData)
	})
	return file_opm_voice_product_map_model_gen_proto_rawDescData
}

var file_opm_voice_product_map_model_gen_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_opm_voice_product_map_model_gen_proto_goTypes = []interface{}{
	(*OpmVoiceProductMap)(nil),                    // 0: service.OpmVoiceProductMap
	(*OpmVoiceProductMapFilter)(nil),              // 1: service.OpmVoiceProductMapFilter
	(*OpmVoiceProductMapListRequest)(nil),         // 2: service.OpmVoiceProductMapListRequest
	(*OpmVoiceProductMapResponse)(nil),            // 3: service.OpmVoiceProductMapResponse
	(*OpmVoiceProductMapUpdateFieldsRequest)(nil), // 4: service.OpmVoiceProductMapUpdateFieldsRequest
	(*OpmVoiceProductMapPrimarykey)(nil),          // 5: service.OpmVoiceProductMapPrimarykey
	(*OpmVoiceProductMapBatchDeleteRequest)(nil),  // 6: service.OpmVoiceProductMapBatchDeleteRequest
	(*timestamppb.Timestamp)(nil),                 // 7: google.protobuf.Timestamp
}
var file_opm_voice_product_map_model_gen_proto_depIdxs = []int32{
	7, // 0: service.OpmVoiceProductMap.createdAt:type_name -> google.protobuf.Timestamp
	7, // 1: service.OpmVoiceProductMapFilter.createdAt:type_name -> google.protobuf.Timestamp
	0, // 2: service.OpmVoiceProductMapListRequest.query:type_name -> service.OpmVoiceProductMap
	0, // 3: service.OpmVoiceProductMapResponse.data:type_name -> service.OpmVoiceProductMap
	0, // 4: service.OpmVoiceProductMapUpdateFieldsRequest.data:type_name -> service.OpmVoiceProductMap
	5, // 5: service.OpmVoiceProductMapBatchDeleteRequest.keys:type_name -> service.OpmVoiceProductMapPrimarykey
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_opm_voice_product_map_model_gen_proto_init() }
func file_opm_voice_product_map_model_gen_proto_init() {
	if File_opm_voice_product_map_model_gen_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_opm_voice_product_map_model_gen_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpmVoiceProductMap); i {
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
		file_opm_voice_product_map_model_gen_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpmVoiceProductMapFilter); i {
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
		file_opm_voice_product_map_model_gen_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpmVoiceProductMapListRequest); i {
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
		file_opm_voice_product_map_model_gen_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpmVoiceProductMapResponse); i {
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
		file_opm_voice_product_map_model_gen_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpmVoiceProductMapUpdateFieldsRequest); i {
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
		file_opm_voice_product_map_model_gen_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpmVoiceProductMapPrimarykey); i {
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
		file_opm_voice_product_map_model_gen_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpmVoiceProductMapBatchDeleteRequest); i {
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
			RawDescriptor: file_opm_voice_product_map_model_gen_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_opm_voice_product_map_model_gen_proto_goTypes,
		DependencyIndexes: file_opm_voice_product_map_model_gen_proto_depIdxs,
		MessageInfos:      file_opm_voice_product_map_model_gen_proto_msgTypes,
	}.Build()
	File_opm_voice_product_map_model_gen_proto = out.File
	file_opm_voice_product_map_model_gen_proto_rawDesc = nil
	file_opm_voice_product_map_model_gen_proto_goTypes = nil
	file_opm_voice_product_map_model_gen_proto_depIdxs = nil
}
