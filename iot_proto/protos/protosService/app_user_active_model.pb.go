// Code generated by protoc,2022-08-10 22:27:08. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: app_user_active_model.proto

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

type AppUserStatisticsFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppKey string `protobuf:"bytes,101,opt,name=appKey,proto3" json:"appKey,omitempty"`
}

func (x *AppUserStatisticsFilter) Reset() {
	*x = AppUserStatisticsFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_user_active_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppUserStatisticsFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppUserStatisticsFilter) ProtoMessage() {}

func (x *AppUserStatisticsFilter) ProtoReflect() protoreflect.Message {
	mi := &file_app_user_active_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppUserStatisticsFilter.ProtoReflect.Descriptor instead.
func (*AppUserStatisticsFilter) Descriptor() ([]byte, []int) {
	return file_app_user_active_model_proto_rawDescGZIP(), []int{0}
}

func (x *AppUserStatisticsFilter) GetAppKey() string {
	if x != nil {
		return x.AppKey
	}
	return ""
}

type AppUserEntitys struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppUser            *Data `protobuf:"bytes,101,opt,name=appUser,proto3" json:"appUser,omitempty"`
	ActiveUser         *Data `protobuf:"bytes,102,opt,name=activeUser,proto3" json:"activeUser,omitempty"`
	AppUserTodayActive int32 `protobuf:"varint,103,opt,name=appUserTodayActive,proto3" json:"appUserTodayActive,omitempty"`
	AppUserToday       int32 `protobuf:"varint,104,opt,name=appUserToday,proto3" json:"appUserToday,omitempty"`
	AppUserAll         int32 `protobuf:"varint,105,opt,name=appUserAll,proto3" json:"appUserAll,omitempty"`
}

func (x *AppUserEntitys) Reset() {
	*x = AppUserEntitys{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_user_active_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppUserEntitys) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppUserEntitys) ProtoMessage() {}

func (x *AppUserEntitys) ProtoReflect() protoreflect.Message {
	mi := &file_app_user_active_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppUserEntitys.ProtoReflect.Descriptor instead.
func (*AppUserEntitys) Descriptor() ([]byte, []int) {
	return file_app_user_active_model_proto_rawDescGZIP(), []int{1}
}

func (x *AppUserEntitys) GetAppUser() *Data {
	if x != nil {
		return x.AppUser
	}
	return nil
}

func (x *AppUserEntitys) GetActiveUser() *Data {
	if x != nil {
		return x.ActiveUser
	}
	return nil
}

func (x *AppUserEntitys) GetAppUserTodayActive() int32 {
	if x != nil {
		return x.AppUserTodayActive
	}
	return 0
}

func (x *AppUserEntitys) GetAppUserToday() int32 {
	if x != nil {
		return x.AppUserToday
	}
	return 0
}

func (x *AppUserEntitys) GetAppUserAll() int32 {
	if x != nil {
		return x.AppUserAll
	}
	return 0
}

type AppUserStatisticsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32           `protobuf:"varint,101,opt,name=code,proto3" json:"code,omitempty"`
	Message string          `protobuf:"bytes,102,opt,name=message,proto3" json:"message,omitempty"`
	Data    *AppUserEntitys `protobuf:"bytes,104,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AppUserStatisticsResponse) Reset() {
	*x = AppUserStatisticsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_user_active_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppUserStatisticsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppUserStatisticsResponse) ProtoMessage() {}

func (x *AppUserStatisticsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_user_active_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppUserStatisticsResponse.ProtoReflect.Descriptor instead.
func (*AppUserStatisticsResponse) Descriptor() ([]byte, []int) {
	return file_app_user_active_model_proto_rawDescGZIP(), []int{2}
}

func (x *AppUserStatisticsResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *AppUserStatisticsResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AppUserStatisticsResponse) GetData() *AppUserEntitys {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_app_user_active_model_proto protoreflect.FileDescriptor

var file_app_user_active_model_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x70, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x19, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x31, 0x0a, 0x17, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x70, 0x70, 0x4b, 0x65, 0x79, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70,
	0x70, 0x4b, 0x65, 0x79, 0x22, 0xdc, 0x01, 0x0a, 0x0e, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x73, 0x12, 0x27, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x55, 0x73,
	0x65, 0x72, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x61, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x2d, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x18, 0x66,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x2e, 0x0a, 0x12, 0x61, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x64, 0x61, 0x79, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x61, 0x70, 0x70,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x64, 0x61, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x61, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x64, 0x61, 0x79, 0x18,
	0x68, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x61, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f,
	0x64, 0x61, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6c,
	0x6c, 0x18, 0x69, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x6c, 0x6c, 0x22, 0x76, 0x0a, 0x19, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x68, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x73, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x11, 0x5a, 0x0f, 0x2e,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_user_active_model_proto_rawDescOnce sync.Once
	file_app_user_active_model_proto_rawDescData = file_app_user_active_model_proto_rawDesc
)

func file_app_user_active_model_proto_rawDescGZIP() []byte {
	file_app_user_active_model_proto_rawDescOnce.Do(func() {
		file_app_user_active_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_user_active_model_proto_rawDescData)
	})
	return file_app_user_active_model_proto_rawDescData
}

var file_app_user_active_model_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_app_user_active_model_proto_goTypes = []interface{}{
	(*AppUserStatisticsFilter)(nil),   // 0: service.AppUserStatisticsFilter
	(*AppUserEntitys)(nil),            // 1: service.AppUserEntitys
	(*AppUserStatisticsResponse)(nil), // 2: service.AppUserStatisticsResponse
	(*Data)(nil),                      // 3: service.Data
}
var file_app_user_active_model_proto_depIdxs = []int32{
	3, // 0: service.AppUserEntitys.appUser:type_name -> service.Data
	3, // 1: service.AppUserEntitys.activeUser:type_name -> service.Data
	1, // 2: service.AppUserStatisticsResponse.data:type_name -> service.AppUserEntitys
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_app_user_active_model_proto_init() }
func file_app_user_active_model_proto_init() {
	if File_app_user_active_model_proto != nil {
		return
	}
	file_device_active_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_app_user_active_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppUserStatisticsFilter); i {
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
		file_app_user_active_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppUserEntitys); i {
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
		file_app_user_active_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppUserStatisticsResponse); i {
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
			RawDescriptor: file_app_user_active_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_user_active_model_proto_goTypes,
		DependencyIndexes: file_app_user_active_model_proto_depIdxs,
		MessageInfos:      file_app_user_active_model_proto_msgTypes,
	}.Build()
	File_app_user_active_model_proto = out.File
	file_app_user_active_model_proto_rawDesc = nil
	file_app_user_active_model_proto_goTypes = nil
	file_app_user_active_model_proto_depIdxs = nil
}
