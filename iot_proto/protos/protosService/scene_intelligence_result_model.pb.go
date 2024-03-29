// Code generated by protoc,2022-05-20 13:36:03. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: scene_intelligence_result_model.proto

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
type SceneIntelligenceResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:唯一主键;comment:唯一主键;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"runTime" gorm:"default:运行时间;comment:运行时间;"
	RunTime *timestamppb.Timestamp `protobuf:"bytes,101,opt,name=runTime,proto3" json:"runTime,omitempty"`
	// @inject_tag: json:"intelligenceId" gorm:"default:智能场景编号;comment:智能场景编号;size:19;"
	IntelligenceId int64 `protobuf:"varint,102,opt,name=intelligenceId,proto3" json:"intelligenceId,omitempty"`
	// @inject_tag: json:"runStatus" gorm:"default:运行状态 =1 执行中 = 2 执行成功 =3 执行失败;comment:运行状态 =1 执行中 = 2 执行成功 =3 执行失败;size:10;"
	RunStatus int32 `protobuf:"varint,103,opt,name=runStatus,proto3" json:"runStatus,omitempty"`
}

func (x *SceneIntelligenceResult) Reset() {
	*x = SceneIntelligenceResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scene_intelligence_result_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneIntelligenceResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneIntelligenceResult) ProtoMessage() {}

func (x *SceneIntelligenceResult) ProtoReflect() protoreflect.Message {
	mi := &file_scene_intelligence_result_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneIntelligenceResult.ProtoReflect.Descriptor instead.
func (*SceneIntelligenceResult) Descriptor() ([]byte, []int) {
	return file_scene_intelligence_result_model_proto_rawDescGZIP(), []int{0}
}

func (x *SceneIntelligenceResult) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SceneIntelligenceResult) GetRunTime() *timestamppb.Timestamp {
	if x != nil {
		return x.RunTime
	}
	return nil
}

func (x *SceneIntelligenceResult) GetIntelligenceId() int64 {
	if x != nil {
		return x.IntelligenceId
	}
	return 0
}

func (x *SceneIntelligenceResult) GetRunStatus() int32 {
	if x != nil {
		return x.RunStatus
	}
	return 0
}

type SceneIntelligenceResultFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:唯一主键;comment:唯一主键;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"runTime" gorm:"default:运行时间;comment:运行时间;"
	RunTime *timestamppb.Timestamp `protobuf:"bytes,101,opt,name=runTime,proto3" json:"runTime,omitempty"`
	// @inject_tag: json:"intelligenceId" gorm:"default:智能场景编号;comment:智能场景编号;size:19;"
	IntelligenceId int64 `protobuf:"varint,102,opt,name=intelligenceId,proto3" json:"intelligenceId,omitempty"`
	// @inject_tag: json:"runStatus" gorm:"default:运行状态 =1 执行中 = 2 执行成功 =3 执行失败;comment:运行状态 =1 执行中 = 2 执行成功 =3 执行失败;size:10;"
	RunStatus int32 `protobuf:"varint,103,opt,name=runStatus,proto3" json:"runStatus,omitempty"`
}

func (x *SceneIntelligenceResultFilter) Reset() {
	*x = SceneIntelligenceResultFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scene_intelligence_result_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneIntelligenceResultFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneIntelligenceResultFilter) ProtoMessage() {}

func (x *SceneIntelligenceResultFilter) ProtoReflect() protoreflect.Message {
	mi := &file_scene_intelligence_result_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneIntelligenceResultFilter.ProtoReflect.Descriptor instead.
func (*SceneIntelligenceResultFilter) Descriptor() ([]byte, []int) {
	return file_scene_intelligence_result_model_proto_rawDescGZIP(), []int{1}
}

func (x *SceneIntelligenceResultFilter) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SceneIntelligenceResultFilter) GetRunTime() *timestamppb.Timestamp {
	if x != nil {
		return x.RunTime
	}
	return nil
}

func (x *SceneIntelligenceResultFilter) GetIntelligenceId() int64 {
	if x != nil {
		return x.IntelligenceId
	}
	return 0
}

func (x *SceneIntelligenceResultFilter) GetRunStatus() int32 {
	if x != nil {
		return x.RunStatus
	}
	return 0
}

type SceneIntelligenceResultListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query     *SceneIntelligenceResult `protobuf:"bytes,101,opt,name=query,proto3" json:"query,omitempty"`
	Page      int64                    `protobuf:"varint,102,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int64                    `protobuf:"varint,103,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	OrderKey  string                   `protobuf:"bytes,104,opt,name=orderKey,proto3" json:"orderKey,omitempty"`
	OrderDesc string                   `protobuf:"bytes,105,opt,name=orderDesc,proto3" json:"orderDesc,omitempty"`
	SearchKey string                   `protobuf:"bytes,106,opt,name=searchKey,proto3" json:"searchKey,omitempty"`
}

func (x *SceneIntelligenceResultListRequest) Reset() {
	*x = SceneIntelligenceResultListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scene_intelligence_result_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneIntelligenceResultListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneIntelligenceResultListRequest) ProtoMessage() {}

func (x *SceneIntelligenceResultListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scene_intelligence_result_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneIntelligenceResultListRequest.ProtoReflect.Descriptor instead.
func (*SceneIntelligenceResultListRequest) Descriptor() ([]byte, []int) {
	return file_scene_intelligence_result_model_proto_rawDescGZIP(), []int{2}
}

func (x *SceneIntelligenceResultListRequest) GetQuery() *SceneIntelligenceResult {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *SceneIntelligenceResultListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SceneIntelligenceResultListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SceneIntelligenceResultListRequest) GetOrderKey() string {
	if x != nil {
		return x.OrderKey
	}
	return ""
}

func (x *SceneIntelligenceResultListRequest) GetOrderDesc() string {
	if x != nil {
		return x.OrderDesc
	}
	return ""
}

func (x *SceneIntelligenceResultListRequest) GetSearchKey() string {
	if x != nil {
		return x.SearchKey
	}
	return ""
}

type SceneIntelligenceResultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32                      `protobuf:"varint,101,opt,name=code,proto3" json:"code,omitempty"`
	Message string                     `protobuf:"bytes,102,opt,name=message,proto3" json:"message,omitempty"`
	Total   int64                      `protobuf:"varint,103,opt,name=total,proto3" json:"total,omitempty"`
	Data    []*SceneIntelligenceResult `protobuf:"bytes,104,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *SceneIntelligenceResultResponse) Reset() {
	*x = SceneIntelligenceResultResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scene_intelligence_result_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneIntelligenceResultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneIntelligenceResultResponse) ProtoMessage() {}

func (x *SceneIntelligenceResultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scene_intelligence_result_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneIntelligenceResultResponse.ProtoReflect.Descriptor instead.
func (*SceneIntelligenceResultResponse) Descriptor() ([]byte, []int) {
	return file_scene_intelligence_result_model_proto_rawDescGZIP(), []int{3}
}

func (x *SceneIntelligenceResultResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SceneIntelligenceResultResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SceneIntelligenceResultResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *SceneIntelligenceResultResponse) GetData() []*SceneIntelligenceResult {
	if x != nil {
		return x.Data
	}
	return nil
}

type SceneIntelligenceResultUpdateFieldsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []string                 `protobuf:"bytes,101,rep,name=fields,proto3" json:"fields,omitempty"`
	Data   *SceneIntelligenceResult `protobuf:"bytes,102,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SceneIntelligenceResultUpdateFieldsRequest) Reset() {
	*x = SceneIntelligenceResultUpdateFieldsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scene_intelligence_result_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneIntelligenceResultUpdateFieldsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneIntelligenceResultUpdateFieldsRequest) ProtoMessage() {}

func (x *SceneIntelligenceResultUpdateFieldsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scene_intelligence_result_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneIntelligenceResultUpdateFieldsRequest.ProtoReflect.Descriptor instead.
func (*SceneIntelligenceResultUpdateFieldsRequest) Descriptor() ([]byte, []int) {
	return file_scene_intelligence_result_model_proto_rawDescGZIP(), []int{4}
}

func (x *SceneIntelligenceResultUpdateFieldsRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *SceneIntelligenceResultUpdateFieldsRequest) GetData() *SceneIntelligenceResult {
	if x != nil {
		return x.Data
	}
	return nil
}

// 表模型主键，只能是整数或字符串，若生成错误，请检查数据库表主键是否正确。支持组合主键
type SceneIntelligenceResultPrimarykey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SceneIntelligenceResultPrimarykey) Reset() {
	*x = SceneIntelligenceResultPrimarykey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scene_intelligence_result_model_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneIntelligenceResultPrimarykey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneIntelligenceResultPrimarykey) ProtoMessage() {}

func (x *SceneIntelligenceResultPrimarykey) ProtoReflect() protoreflect.Message {
	mi := &file_scene_intelligence_result_model_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneIntelligenceResultPrimarykey.ProtoReflect.Descriptor instead.
func (*SceneIntelligenceResultPrimarykey) Descriptor() ([]byte, []int) {
	return file_scene_intelligence_result_model_proto_rawDescGZIP(), []int{5}
}

func (x *SceneIntelligenceResultPrimarykey) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SceneIntelligenceResultBatchDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []*SceneIntelligenceResultPrimarykey `protobuf:"bytes,101,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *SceneIntelligenceResultBatchDeleteRequest) Reset() {
	*x = SceneIntelligenceResultBatchDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scene_intelligence_result_model_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneIntelligenceResultBatchDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneIntelligenceResultBatchDeleteRequest) ProtoMessage() {}

func (x *SceneIntelligenceResultBatchDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scene_intelligence_result_model_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneIntelligenceResultBatchDeleteRequest.ProtoReflect.Descriptor instead.
func (*SceneIntelligenceResultBatchDeleteRequest) Descriptor() ([]byte, []int) {
	return file_scene_intelligence_result_model_proto_rawDescGZIP(), []int{6}
}

func (x *SceneIntelligenceResultBatchDeleteRequest) GetKeys() []*SceneIntelligenceResultPrimarykey {
	if x != nil {
		return x.Keys
	}
	return nil
}

var File_scene_intelligence_result_model_proto protoreflect.FileDescriptor

var file_scene_intelligence_result_model_proto_rawDesc = []byte{
	0x0a, 0x25, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x67,
	0x65, 0x6e, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x17, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x6c,
	0x6c, 0x69, 0x67, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x34, 0x0a,
	0x07, 0x72, 0x75, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x67, 0x65,
	0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x66, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x69, 0x6e, 0x74,
	0x65, 0x6c, 0x6c, 0x69, 0x67, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72,
	0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x67, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x72, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xab, 0x01, 0x0a, 0x1d, 0x53, 0x63,
	0x65, 0x6e, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x67, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x72,
	0x75, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x26, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x67, 0x65, 0x6e, 0x63,
	0x65, 0x49, 0x64, 0x18, 0x66, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x6c,
	0x6c, 0x69, 0x67, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x75, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x67, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x75,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xe4, 0x01, 0x0a, 0x22, 0x53, 0x63, 0x65, 0x6e,
	0x65, 0x49, 0x6e, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x67, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36,
	0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x74,
	0x65, 0x6c, 0x6c, 0x69, 0x67, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x66,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4b,
	0x65, 0x79, 0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4b,
	0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63, 0x18,
	0x69, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x18, 0x6a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x22, 0x9b,
	0x01, 0x0a, 0x1f, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x67,
	0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x34, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x68,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53,
	0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x67, 0x65, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x7a, 0x0a, 0x2a,
	0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x67, 0x65, 0x6e, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x12, 0x34, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65,
	0x49, 0x6e, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x67, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x33, 0x0a, 0x21, 0x53, 0x63, 0x65, 0x6e,
	0x65, 0x49, 0x6e, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x67, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6b, 0x0a,
	0x29, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x67, 0x65, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x04, 0x6b, 0x65,
	0x79, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x67,
	0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x6b, 0x65, 0x79, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scene_intelligence_result_model_proto_rawDescOnce sync.Once
	file_scene_intelligence_result_model_proto_rawDescData = file_scene_intelligence_result_model_proto_rawDesc
)

func file_scene_intelligence_result_model_proto_rawDescGZIP() []byte {
	file_scene_intelligence_result_model_proto_rawDescOnce.Do(func() {
		file_scene_intelligence_result_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_scene_intelligence_result_model_proto_rawDescData)
	})
	return file_scene_intelligence_result_model_proto_rawDescData
}

var file_scene_intelligence_result_model_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_scene_intelligence_result_model_proto_goTypes = []interface{}{
	(*SceneIntelligenceResult)(nil),                    // 0: service.SceneIntelligenceResult
	(*SceneIntelligenceResultFilter)(nil),              // 1: service.SceneIntelligenceResultFilter
	(*SceneIntelligenceResultListRequest)(nil),         // 2: service.SceneIntelligenceResultListRequest
	(*SceneIntelligenceResultResponse)(nil),            // 3: service.SceneIntelligenceResultResponse
	(*SceneIntelligenceResultUpdateFieldsRequest)(nil), // 4: service.SceneIntelligenceResultUpdateFieldsRequest
	(*SceneIntelligenceResultPrimarykey)(nil),          // 5: service.SceneIntelligenceResultPrimarykey
	(*SceneIntelligenceResultBatchDeleteRequest)(nil),  // 6: service.SceneIntelligenceResultBatchDeleteRequest
	(*timestamppb.Timestamp)(nil),                      // 7: google.protobuf.Timestamp
}
var file_scene_intelligence_result_model_proto_depIdxs = []int32{
	7, // 0: service.SceneIntelligenceResult.runTime:type_name -> google.protobuf.Timestamp
	7, // 1: service.SceneIntelligenceResultFilter.runTime:type_name -> google.protobuf.Timestamp
	0, // 2: service.SceneIntelligenceResultListRequest.query:type_name -> service.SceneIntelligenceResult
	0, // 3: service.SceneIntelligenceResultResponse.data:type_name -> service.SceneIntelligenceResult
	0, // 4: service.SceneIntelligenceResultUpdateFieldsRequest.data:type_name -> service.SceneIntelligenceResult
	5, // 5: service.SceneIntelligenceResultBatchDeleteRequest.keys:type_name -> service.SceneIntelligenceResultPrimarykey
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_scene_intelligence_result_model_proto_init() }
func file_scene_intelligence_result_model_proto_init() {
	if File_scene_intelligence_result_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scene_intelligence_result_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneIntelligenceResult); i {
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
		file_scene_intelligence_result_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneIntelligenceResultFilter); i {
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
		file_scene_intelligence_result_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneIntelligenceResultListRequest); i {
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
		file_scene_intelligence_result_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneIntelligenceResultResponse); i {
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
		file_scene_intelligence_result_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneIntelligenceResultUpdateFieldsRequest); i {
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
		file_scene_intelligence_result_model_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneIntelligenceResultPrimarykey); i {
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
		file_scene_intelligence_result_model_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneIntelligenceResultBatchDeleteRequest); i {
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
			RawDescriptor: file_scene_intelligence_result_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_scene_intelligence_result_model_proto_goTypes,
		DependencyIndexes: file_scene_intelligence_result_model_proto_depIdxs,
		MessageInfos:      file_scene_intelligence_result_model_proto_msgTypes,
	}.Build()
	File_scene_intelligence_result_model_proto = out.File
	file_scene_intelligence_result_model_proto_rawDesc = nil
	file_scene_intelligence_result_model_proto_goTypes = nil
	file_scene_intelligence_result_model_proto_depIdxs = nil
}
