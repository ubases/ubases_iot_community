// Code generated by protoc,2022-05-20 13:36:04. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: scene_intelligence_task_model.proto

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
type SceneIntelligenceTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:唯一主键;comment:唯一主键;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"intelligenceId" gorm:"default:智能场景Id;comment:智能场景Id;size:19;"
	IntelligenceId int64 `protobuf:"varint,101,opt,name=intelligenceId,proto3" json:"intelligenceId,omitempty"`
	// @inject_tag: json:"taskImg" gorm:"default:任务图片(产品图片、智能图片、功能图标）;comment:任务图片(产品图片、智能图片、功能图标）;size:255;"
	TaskImg string `protobuf:"bytes,102,opt,name=taskImg,proto3" json:"taskImg,omitempty"`
	// @inject_tag: json:"taskDesc" gorm:"default:任务描述;comment:任务描述;size:50;"
	TaskDesc string `protobuf:"bytes,103,opt,name=taskDesc,proto3" json:"taskDesc,omitempty"`
	// @inject_tag: json:"taskType" gorm:"default:任务类型（=1 延时 =2 设备执行 =3 场景开关）;comment:任务类型（=1 延时 =2 设备执行 =3 场景开关）;size:10;"
	TaskType int32 `protobuf:"varint,104,opt,name=taskType,proto3" json:"taskType,omitempty"`
	// @inject_tag: json:"objectId" gorm:"default:对象ID（设备Id、场景Id）;comment:对象ID（设备Id、场景Id）;size:50;"
	ObjectId string `protobuf:"bytes,105,opt,name=objectId,proto3" json:"objectId,omitempty"`
	// @inject_tag: json:"objectDesc" gorm:"default:对象的标题或者描述（设备名称、场景名称）;comment:对象的标题或者描述（设备名称、场景名称）;size:50;"
	ObjectDesc string `protobuf:"bytes,106,opt,name=objectDesc,proto3" json:"objectDesc,omitempty"`
	// @inject_tag: json:"funcKey" gorm:"default:执行功能Key;comment:执行功能Key;size:50;"
	FuncKey string `protobuf:"bytes,107,opt,name=funcKey,proto3" json:"funcKey,omitempty"`
	// @inject_tag: json:"funcDesc" gorm:"default:冗余：功能描述;comment:冗余：功能描述;size:50;"
	FuncDesc string `protobuf:"bytes,108,opt,name=funcDesc,proto3" json:"funcDesc,omitempty"`
	// @inject_tag: json:"funcValue" gorm:"default:;comment:;size:50;"
	FuncValue      string `protobuf:"bytes,109,opt,name=funcValue,proto3" json:"funcValue,omitempty"`
	ProductKey     string `protobuf:"bytes,110,opt,name=productKey,proto3" json:"productKey,omitempty"`
	ProductId      int64  `protobuf:"varint,111,opt,name=productId,proto3" json:"productId,omitempty"`
	Functions      string `protobuf:"bytes,112,opt,name=functions,proto3" json:"functions,omitempty"`
	FuncIdentifier string `protobuf:"bytes,113,opt,name=funcIdentifier,proto3" json:"funcIdentifier,omitempty"` //功能标识符
}

func (x *SceneIntelligenceTask) Reset() {
	*x = SceneIntelligenceTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scene_intelligence_task_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneIntelligenceTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneIntelligenceTask) ProtoMessage() {}

func (x *SceneIntelligenceTask) ProtoReflect() protoreflect.Message {
	mi := &file_scene_intelligence_task_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneIntelligenceTask.ProtoReflect.Descriptor instead.
func (*SceneIntelligenceTask) Descriptor() ([]byte, []int) {
	return file_scene_intelligence_task_model_proto_rawDescGZIP(), []int{0}
}

func (x *SceneIntelligenceTask) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SceneIntelligenceTask) GetIntelligenceId() int64 {
	if x != nil {
		return x.IntelligenceId
	}
	return 0
}

func (x *SceneIntelligenceTask) GetTaskImg() string {
	if x != nil {
		return x.TaskImg
	}
	return ""
}

func (x *SceneIntelligenceTask) GetTaskDesc() string {
	if x != nil {
		return x.TaskDesc
	}
	return ""
}

func (x *SceneIntelligenceTask) GetTaskType() int32 {
	if x != nil {
		return x.TaskType
	}
	return 0
}

func (x *SceneIntelligenceTask) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *SceneIntelligenceTask) GetObjectDesc() string {
	if x != nil {
		return x.ObjectDesc
	}
	return ""
}

func (x *SceneIntelligenceTask) GetFuncKey() string {
	if x != nil {
		return x.FuncKey
	}
	return ""
}

func (x *SceneIntelligenceTask) GetFuncDesc() string {
	if x != nil {
		return x.FuncDesc
	}
	return ""
}

func (x *SceneIntelligenceTask) GetFuncValue() string {
	if x != nil {
		return x.FuncValue
	}
	return ""
}

func (x *SceneIntelligenceTask) GetProductKey() string {
	if x != nil {
		return x.ProductKey
	}
	return ""
}

func (x *SceneIntelligenceTask) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *SceneIntelligenceTask) GetFunctions() string {
	if x != nil {
		return x.Functions
	}
	return ""
}

func (x *SceneIntelligenceTask) GetFuncIdentifier() string {
	if x != nil {
		return x.FuncIdentifier
	}
	return ""
}

type SceneIntelligenceTaskFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:唯一主键;comment:唯一主键;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"intelligenceId" gorm:"default:智能场景Id;comment:智能场景Id;size:19;"
	IntelligenceId int64 `protobuf:"varint,101,opt,name=intelligenceId,proto3" json:"intelligenceId,omitempty"`
	// @inject_tag: json:"taskImg" gorm:"default:任务图片(产品图片、智能图片、功能图标）;comment:任务图片(产品图片、智能图片、功能图标）;size:255;"
	TaskImg string `protobuf:"bytes,102,opt,name=taskImg,proto3" json:"taskImg,omitempty"`
	// @inject_tag: json:"taskDesc" gorm:"default:任务描述;comment:任务描述;size:50;"
	TaskDesc string `protobuf:"bytes,103,opt,name=taskDesc,proto3" json:"taskDesc,omitempty"`
	// @inject_tag: json:"taskType" gorm:"default:任务类型（=1 延时 =2 设备执行 =3 场景开关）;comment:任务类型（=1 延时 =2 设备执行 =3 场景开关）;size:10;"
	TaskType int32 `protobuf:"varint,104,opt,name=taskType,proto3" json:"taskType,omitempty"`
	// @inject_tag: json:"objectId" gorm:"default:对象ID（设备Id、场景Id）;comment:对象ID（设备Id、场景Id）;size:50;"
	ObjectId string `protobuf:"bytes,105,opt,name=objectId,proto3" json:"objectId,omitempty"`
	// @inject_tag: json:"objectDesc" gorm:"default:对象的标题或者描述（设备名称、场景名称）;comment:对象的标题或者描述（设备名称、场景名称）;size:50;"
	ObjectDesc string `protobuf:"bytes,106,opt,name=objectDesc,proto3" json:"objectDesc,omitempty"`
	// @inject_tag: json:"funcKey" gorm:"default:执行功能Key;comment:执行功能Key;size:50;"
	FuncKey string `protobuf:"bytes,107,opt,name=funcKey,proto3" json:"funcKey,omitempty"`
	// @inject_tag: json:"funcDesc" gorm:"default:冗余：功能描述;comment:冗余：功能描述;size:50;"
	FuncDesc string `protobuf:"bytes,108,opt,name=funcDesc,proto3" json:"funcDesc,omitempty"`
	// @inject_tag: json:"funcValue" gorm:"default:;comment:;size:50;"
	FuncValue  string `protobuf:"bytes,109,opt,name=funcValue,proto3" json:"funcValue,omitempty"`
	ProductKey string `protobuf:"bytes,110,opt,name=productKey,proto3" json:"productKey,omitempty"`
	ProductId  int64  `protobuf:"varint,111,opt,name=productId,proto3" json:"productId,omitempty"`
}

func (x *SceneIntelligenceTaskFilter) Reset() {
	*x = SceneIntelligenceTaskFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scene_intelligence_task_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneIntelligenceTaskFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneIntelligenceTaskFilter) ProtoMessage() {}

func (x *SceneIntelligenceTaskFilter) ProtoReflect() protoreflect.Message {
	mi := &file_scene_intelligence_task_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneIntelligenceTaskFilter.ProtoReflect.Descriptor instead.
func (*SceneIntelligenceTaskFilter) Descriptor() ([]byte, []int) {
	return file_scene_intelligence_task_model_proto_rawDescGZIP(), []int{1}
}

func (x *SceneIntelligenceTaskFilter) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SceneIntelligenceTaskFilter) GetIntelligenceId() int64 {
	if x != nil {
		return x.IntelligenceId
	}
	return 0
}

func (x *SceneIntelligenceTaskFilter) GetTaskImg() string {
	if x != nil {
		return x.TaskImg
	}
	return ""
}

func (x *SceneIntelligenceTaskFilter) GetTaskDesc() string {
	if x != nil {
		return x.TaskDesc
	}
	return ""
}

func (x *SceneIntelligenceTaskFilter) GetTaskType() int32 {
	if x != nil {
		return x.TaskType
	}
	return 0
}

func (x *SceneIntelligenceTaskFilter) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *SceneIntelligenceTaskFilter) GetObjectDesc() string {
	if x != nil {
		return x.ObjectDesc
	}
	return ""
}

func (x *SceneIntelligenceTaskFilter) GetFuncKey() string {
	if x != nil {
		return x.FuncKey
	}
	return ""
}

func (x *SceneIntelligenceTaskFilter) GetFuncDesc() string {
	if x != nil {
		return x.FuncDesc
	}
	return ""
}

func (x *SceneIntelligenceTaskFilter) GetFuncValue() string {
	if x != nil {
		return x.FuncValue
	}
	return ""
}

func (x *SceneIntelligenceTaskFilter) GetProductKey() string {
	if x != nil {
		return x.ProductKey
	}
	return ""
}

func (x *SceneIntelligenceTaskFilter) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

type SceneIntelligenceTaskListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query     *SceneIntelligenceTask `protobuf:"bytes,101,opt,name=query,proto3" json:"query,omitempty"`
	Page      int64                  `protobuf:"varint,102,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int64                  `protobuf:"varint,103,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	OrderKey  string                 `protobuf:"bytes,104,opt,name=orderKey,proto3" json:"orderKey,omitempty"`
	OrderDesc string                 `protobuf:"bytes,105,opt,name=orderDesc,proto3" json:"orderDesc,omitempty"`
	SearchKey string                 `protobuf:"bytes,106,opt,name=searchKey,proto3" json:"searchKey,omitempty"`
}

func (x *SceneIntelligenceTaskListRequest) Reset() {
	*x = SceneIntelligenceTaskListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scene_intelligence_task_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneIntelligenceTaskListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneIntelligenceTaskListRequest) ProtoMessage() {}

func (x *SceneIntelligenceTaskListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scene_intelligence_task_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneIntelligenceTaskListRequest.ProtoReflect.Descriptor instead.
func (*SceneIntelligenceTaskListRequest) Descriptor() ([]byte, []int) {
	return file_scene_intelligence_task_model_proto_rawDescGZIP(), []int{2}
}

func (x *SceneIntelligenceTaskListRequest) GetQuery() *SceneIntelligenceTask {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *SceneIntelligenceTaskListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SceneIntelligenceTaskListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SceneIntelligenceTaskListRequest) GetOrderKey() string {
	if x != nil {
		return x.OrderKey
	}
	return ""
}

func (x *SceneIntelligenceTaskListRequest) GetOrderDesc() string {
	if x != nil {
		return x.OrderDesc
	}
	return ""
}

func (x *SceneIntelligenceTaskListRequest) GetSearchKey() string {
	if x != nil {
		return x.SearchKey
	}
	return ""
}

type SceneIntelligenceTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32                    `protobuf:"varint,101,opt,name=code,proto3" json:"code,omitempty"`
	Message string                   `protobuf:"bytes,102,opt,name=message,proto3" json:"message,omitempty"`
	Total   int64                    `protobuf:"varint,103,opt,name=total,proto3" json:"total,omitempty"`
	Data    []*SceneIntelligenceTask `protobuf:"bytes,104,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *SceneIntelligenceTaskResponse) Reset() {
	*x = SceneIntelligenceTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scene_intelligence_task_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneIntelligenceTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneIntelligenceTaskResponse) ProtoMessage() {}

func (x *SceneIntelligenceTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scene_intelligence_task_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneIntelligenceTaskResponse.ProtoReflect.Descriptor instead.
func (*SceneIntelligenceTaskResponse) Descriptor() ([]byte, []int) {
	return file_scene_intelligence_task_model_proto_rawDescGZIP(), []int{3}
}

func (x *SceneIntelligenceTaskResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SceneIntelligenceTaskResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SceneIntelligenceTaskResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *SceneIntelligenceTaskResponse) GetData() []*SceneIntelligenceTask {
	if x != nil {
		return x.Data
	}
	return nil
}

type SceneIntelligenceTaskUpdateFieldsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []string               `protobuf:"bytes,101,rep,name=fields,proto3" json:"fields,omitempty"`
	Data   *SceneIntelligenceTask `protobuf:"bytes,102,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SceneIntelligenceTaskUpdateFieldsRequest) Reset() {
	*x = SceneIntelligenceTaskUpdateFieldsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scene_intelligence_task_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneIntelligenceTaskUpdateFieldsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneIntelligenceTaskUpdateFieldsRequest) ProtoMessage() {}

func (x *SceneIntelligenceTaskUpdateFieldsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scene_intelligence_task_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneIntelligenceTaskUpdateFieldsRequest.ProtoReflect.Descriptor instead.
func (*SceneIntelligenceTaskUpdateFieldsRequest) Descriptor() ([]byte, []int) {
	return file_scene_intelligence_task_model_proto_rawDescGZIP(), []int{4}
}

func (x *SceneIntelligenceTaskUpdateFieldsRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *SceneIntelligenceTaskUpdateFieldsRequest) GetData() *SceneIntelligenceTask {
	if x != nil {
		return x.Data
	}
	return nil
}

// 表模型主键，只能是整数或字符串，若生成错误，请检查数据库表主键是否正确。支持组合主键
type SceneIntelligenceTaskPrimarykey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SceneIntelligenceTaskPrimarykey) Reset() {
	*x = SceneIntelligenceTaskPrimarykey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scene_intelligence_task_model_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneIntelligenceTaskPrimarykey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneIntelligenceTaskPrimarykey) ProtoMessage() {}

func (x *SceneIntelligenceTaskPrimarykey) ProtoReflect() protoreflect.Message {
	mi := &file_scene_intelligence_task_model_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneIntelligenceTaskPrimarykey.ProtoReflect.Descriptor instead.
func (*SceneIntelligenceTaskPrimarykey) Descriptor() ([]byte, []int) {
	return file_scene_intelligence_task_model_proto_rawDescGZIP(), []int{5}
}

func (x *SceneIntelligenceTaskPrimarykey) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SceneIntelligenceTaskBatchDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []*SceneIntelligenceTaskPrimarykey `protobuf:"bytes,101,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *SceneIntelligenceTaskBatchDeleteRequest) Reset() {
	*x = SceneIntelligenceTaskBatchDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scene_intelligence_task_model_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneIntelligenceTaskBatchDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneIntelligenceTaskBatchDeleteRequest) ProtoMessage() {}

func (x *SceneIntelligenceTaskBatchDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scene_intelligence_task_model_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneIntelligenceTaskBatchDeleteRequest.ProtoReflect.Descriptor instead.
func (*SceneIntelligenceTaskBatchDeleteRequest) Descriptor() ([]byte, []int) {
	return file_scene_intelligence_task_model_proto_rawDescGZIP(), []int{6}
}

func (x *SceneIntelligenceTaskBatchDeleteRequest) GetKeys() []*SceneIntelligenceTaskPrimarykey {
	if x != nil {
		return x.Keys
	}
	return nil
}

var File_scene_intelligence_task_model_proto protoreflect.FileDescriptor

var file_scene_intelligence_task_model_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x67,
	0x65, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0xb5,
	0x03, 0x0a, 0x15, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x67,
	0x65, 0x6e, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65,
	0x6c, 0x6c, 0x69, 0x67, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x65, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x67, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x6d, 0x67, 0x18, 0x66, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x6d, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61,
	0x73, 0x6b, 0x44, 0x65, 0x73, 0x63, 0x18, 0x67, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61,
	0x73, 0x6b, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x68, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x69,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x73, 0x63, 0x18, 0x6a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x73, 0x63, 0x12, 0x18,
	0x0a, 0x07, 0x66, 0x75, 0x6e, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x6b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x66, 0x75, 0x6e, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63,
	0x44, 0x65, 0x73, 0x63, 0x18, 0x6c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x63,
	0x44, 0x65, 0x73, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x75, 0x6e, 0x63, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x6d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x75, 0x6e, 0x63, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b, 0x65, 0x79,
	0x18, 0x6e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b,
	0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18,
	0x6f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x70, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26,
	0x0a, 0x0e, 0x66, 0x75, 0x6e, 0x63, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x18, 0x71, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x66, 0x75, 0x6e, 0x63, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0xf5, 0x02, 0x0a, 0x1b, 0x53, 0x63, 0x65, 0x6e, 0x65,
	0x49, 0x6e, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x67, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x6c, 0x6c,
	0x69, 0x67, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x65, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e,
	0x69, 0x6e, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x67, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x6d, 0x67, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x6d, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b,
	0x44, 0x65, 0x73, 0x63, 0x18, 0x67, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b,
	0x44, 0x65, 0x73, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x68, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x69, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x73, 0x63, 0x18, 0x6a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x73, 0x63, 0x12, 0x18, 0x0a, 0x07,
	0x66, 0x75, 0x6e, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x6b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66,
	0x75, 0x6e, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x44, 0x65,
	0x73, 0x63, 0x18, 0x6c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x44, 0x65,
	0x73, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x75, 0x6e, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x6d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x75, 0x6e, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x6e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b, 0x65, 0x79,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x6f, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0xe0,
	0x01, 0x0a, 0x20, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x67,
	0x65, 0x6e, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x65, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x63, 0x65,
	0x6e, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x67, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x4b, 0x65, 0x79, 0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65,
	0x73, 0x63, 0x18, 0x69, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44,
	0x65, 0x73, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79,
	0x18, 0x6a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65,
	0x79, 0x22, 0x97, 0x01, 0x0a, 0x1d, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x6c,
	0x6c, 0x69, 0x67, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x32, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x68, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x67, 0x65, 0x6e, 0x63,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x76, 0x0a, 0x28, 0x53,
	0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x67, 0x65, 0x6e, 0x63, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12,
	0x32, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x74,
	0x65, 0x6c, 0x6c, 0x69, 0x67, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x31, 0x0a, 0x1f, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x74, 0x65,
	0x6c, 0x6c, 0x69, 0x67, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x67, 0x0a, 0x27, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49,
	0x6e, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x67, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3c, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49,
	0x6e, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x67, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x50,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x42,
	0x11, 0x5a, 0x0f, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scene_intelligence_task_model_proto_rawDescOnce sync.Once
	file_scene_intelligence_task_model_proto_rawDescData = file_scene_intelligence_task_model_proto_rawDesc
)

func file_scene_intelligence_task_model_proto_rawDescGZIP() []byte {
	file_scene_intelligence_task_model_proto_rawDescOnce.Do(func() {
		file_scene_intelligence_task_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_scene_intelligence_task_model_proto_rawDescData)
	})
	return file_scene_intelligence_task_model_proto_rawDescData
}

var file_scene_intelligence_task_model_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_scene_intelligence_task_model_proto_goTypes = []interface{}{
	(*SceneIntelligenceTask)(nil),                    // 0: service.SceneIntelligenceTask
	(*SceneIntelligenceTaskFilter)(nil),              // 1: service.SceneIntelligenceTaskFilter
	(*SceneIntelligenceTaskListRequest)(nil),         // 2: service.SceneIntelligenceTaskListRequest
	(*SceneIntelligenceTaskResponse)(nil),            // 3: service.SceneIntelligenceTaskResponse
	(*SceneIntelligenceTaskUpdateFieldsRequest)(nil), // 4: service.SceneIntelligenceTaskUpdateFieldsRequest
	(*SceneIntelligenceTaskPrimarykey)(nil),          // 5: service.SceneIntelligenceTaskPrimarykey
	(*SceneIntelligenceTaskBatchDeleteRequest)(nil),  // 6: service.SceneIntelligenceTaskBatchDeleteRequest
}
var file_scene_intelligence_task_model_proto_depIdxs = []int32{
	0, // 0: service.SceneIntelligenceTaskListRequest.query:type_name -> service.SceneIntelligenceTask
	0, // 1: service.SceneIntelligenceTaskResponse.data:type_name -> service.SceneIntelligenceTask
	0, // 2: service.SceneIntelligenceTaskUpdateFieldsRequest.data:type_name -> service.SceneIntelligenceTask
	5, // 3: service.SceneIntelligenceTaskBatchDeleteRequest.keys:type_name -> service.SceneIntelligenceTaskPrimarykey
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_scene_intelligence_task_model_proto_init() }
func file_scene_intelligence_task_model_proto_init() {
	if File_scene_intelligence_task_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scene_intelligence_task_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneIntelligenceTask); i {
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
		file_scene_intelligence_task_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneIntelligenceTaskFilter); i {
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
		file_scene_intelligence_task_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneIntelligenceTaskListRequest); i {
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
		file_scene_intelligence_task_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneIntelligenceTaskResponse); i {
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
		file_scene_intelligence_task_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneIntelligenceTaskUpdateFieldsRequest); i {
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
		file_scene_intelligence_task_model_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneIntelligenceTaskPrimarykey); i {
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
		file_scene_intelligence_task_model_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneIntelligenceTaskBatchDeleteRequest); i {
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
			RawDescriptor: file_scene_intelligence_task_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_scene_intelligence_task_model_proto_goTypes,
		DependencyIndexes: file_scene_intelligence_task_model_proto_depIdxs,
		MessageInfos:      file_scene_intelligence_task_model_proto_msgTypes,
	}.Build()
	File_scene_intelligence_task_model_proto = out.File
	file_scene_intelligence_task_model_proto_rawDesc = nil
	file_scene_intelligence_task_model_proto_goTypes = nil
	file_scene_intelligence_task_model_proto_depIdxs = nil
}
