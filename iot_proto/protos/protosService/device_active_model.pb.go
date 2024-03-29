// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: device_active_model.proto

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

type DeviceActiveListFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantId   string `protobuf:"bytes,101,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	ProductKey string `protobuf:"bytes,102,opt,name=productKey,proto3" json:"productKey,omitempty"`
}

func (x *DeviceActiveListFilter) Reset() {
	*x = DeviceActiveListFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_active_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceActiveListFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceActiveListFilter) ProtoMessage() {}

func (x *DeviceActiveListFilter) ProtoReflect() protoreflect.Message {
	mi := &file_device_active_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceActiveListFilter.ProtoReflect.Descriptor instead.
func (*DeviceActiveListFilter) Descriptor() ([]byte, []int) {
	return file_device_active_model_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceActiveListFilter) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *DeviceActiveListFilter) GetProductKey() string {
	if x != nil {
		return x.ProductKey
	}
	return ""
}

type OpenActiveEntitys struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceMonActive   *Data `protobuf:"bytes,101,opt,name=deviceMonActive,proto3" json:"deviceMonActive,omitempty"`
	DeviceDayActive   *Data `protobuf:"bytes,102,opt,name=deviceDayActive,proto3" json:"deviceDayActive,omitempty"`
	DeviceTodayActive int32 `protobuf:"varint,103,opt,name=deviceTodayActive,proto3" json:"deviceTodayActive,omitempty"`
	Device7DayActive  int32 `protobuf:"varint,104,opt,name=device7DayActive,proto3" json:"device7DayActive,omitempty"`
	DeviceActiveAll   int32 `protobuf:"varint,105,opt,name=deviceActiveAll,proto3" json:"deviceActiveAll,omitempty"`
}

func (x *OpenActiveEntitys) Reset() {
	*x = OpenActiveEntitys{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_active_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenActiveEntitys) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenActiveEntitys) ProtoMessage() {}

func (x *OpenActiveEntitys) ProtoReflect() protoreflect.Message {
	mi := &file_device_active_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenActiveEntitys.ProtoReflect.Descriptor instead.
func (*OpenActiveEntitys) Descriptor() ([]byte, []int) {
	return file_device_active_model_proto_rawDescGZIP(), []int{1}
}

func (x *OpenActiveEntitys) GetDeviceMonActive() *Data {
	if x != nil {
		return x.DeviceMonActive
	}
	return nil
}

func (x *OpenActiveEntitys) GetDeviceDayActive() *Data {
	if x != nil {
		return x.DeviceDayActive
	}
	return nil
}

func (x *OpenActiveEntitys) GetDeviceTodayActive() int32 {
	if x != nil {
		return x.DeviceTodayActive
	}
	return 0
}

func (x *OpenActiveEntitys) GetDevice7DayActive() int32 {
	if x != nil {
		return x.Device7DayActive
	}
	return 0
}

func (x *OpenActiveEntitys) GetDeviceActiveAll() int32 {
	if x != nil {
		return x.DeviceActiveAll
	}
	return 0
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64       `protobuf:"varint,101,opt,name=total,proto3" json:"total,omitempty"`
	Data  []*TimeData `protobuf:"bytes,102,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_active_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_device_active_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_device_active_model_proto_rawDescGZIP(), []int{2}
}

func (x *Data) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Data) GetData() []*TimeData {
	if x != nil {
		return x.Data
	}
	return nil
}

type TimeData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time  string `protobuf:"bytes,101,opt,name=time,proto3" json:"time,omitempty"`
	Total int64  `protobuf:"varint,102,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *TimeData) Reset() {
	*x = TimeData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_active_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeData) ProtoMessage() {}

func (x *TimeData) ProtoReflect() protoreflect.Message {
	mi := &file_device_active_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeData.ProtoReflect.Descriptor instead.
func (*TimeData) Descriptor() ([]byte, []int) {
	return file_device_active_model_proto_rawDescGZIP(), []int{3}
}

func (x *TimeData) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *TimeData) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type DeviceActiveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32              `protobuf:"varint,101,opt,name=code,proto3" json:"code,omitempty"`
	Message string             `protobuf:"bytes,102,opt,name=message,proto3" json:"message,omitempty"`
	Data    *OpenActiveEntitys `protobuf:"bytes,104,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DeviceActiveResponse) Reset() {
	*x = DeviceActiveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_active_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceActiveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceActiveResponse) ProtoMessage() {}

func (x *DeviceActiveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_device_active_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceActiveResponse.ProtoReflect.Descriptor instead.
func (*DeviceActiveResponse) Descriptor() ([]byte, []int) {
	return file_device_active_model_proto_rawDescGZIP(), []int{4}
}

func (x *DeviceActiveResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DeviceActiveResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *DeviceActiveResponse) GetData() *OpenActiveEntitys {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_device_active_model_proto protoreflect.FileDescriptor

var file_device_active_model_proto_rawDesc = []byte{
	0x0a, 0x19, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x22, 0x54, 0x0a, 0x16, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b, 0x65, 0x79, 0x22, 0x89, 0x02, 0x0a, 0x11, 0x4f,
	0x70, 0x65, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x73,
	0x12, 0x37, 0x0a, 0x0f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x6e, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x4d, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x37, 0x0a, 0x0f, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x44, 0x61, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x66, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x0f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x79, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x64, 0x61,
	0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x64, 0x61, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x2a, 0x0a, 0x10, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x37, 0x44, 0x61, 0x79, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x18, 0x68, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x37, 0x44, 0x61, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x28, 0x0a, 0x0f,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x41, 0x6c, 0x6c, 0x18,
	0x69, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x41, 0x6c, 0x6c, 0x22, 0x43, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x65, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x66, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x34, 0x0a, 0x08, 0x54,
	0x69, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x66, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x22, 0x74, 0x0a, 0x14, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x68, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x4f, 0x70, 0x65, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x73, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_device_active_model_proto_rawDescOnce sync.Once
	file_device_active_model_proto_rawDescData = file_device_active_model_proto_rawDesc
)

func file_device_active_model_proto_rawDescGZIP() []byte {
	file_device_active_model_proto_rawDescOnce.Do(func() {
		file_device_active_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_device_active_model_proto_rawDescData)
	})
	return file_device_active_model_proto_rawDescData
}

var file_device_active_model_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_device_active_model_proto_goTypes = []interface{}{
	(*DeviceActiveListFilter)(nil), // 0: service.DeviceActiveListFilter
	(*OpenActiveEntitys)(nil),      // 1: service.OpenActiveEntitys
	(*Data)(nil),                   // 2: service.Data
	(*TimeData)(nil),               // 3: service.TimeData
	(*DeviceActiveResponse)(nil),   // 4: service.DeviceActiveResponse
}
var file_device_active_model_proto_depIdxs = []int32{
	2, // 0: service.OpenActiveEntitys.deviceMonActive:type_name -> service.Data
	2, // 1: service.OpenActiveEntitys.deviceDayActive:type_name -> service.Data
	3, // 2: service.Data.data:type_name -> service.TimeData
	1, // 3: service.DeviceActiveResponse.data:type_name -> service.OpenActiveEntitys
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_device_active_model_proto_init() }
func file_device_active_model_proto_init() {
	if File_device_active_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_device_active_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceActiveListFilter); i {
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
		file_device_active_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenActiveEntitys); i {
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
		file_device_active_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_device_active_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeData); i {
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
		file_device_active_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceActiveResponse); i {
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
			RawDescriptor: file_device_active_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_device_active_model_proto_goTypes,
		DependencyIndexes: file_device_active_model_proto_depIdxs,
		MessageInfos:      file_device_active_model_proto_msgTypes,
	}.Build()
	File_device_active_model_proto = out.File
	file_device_active_model_proto_rawDesc = nil
	file_device_active_model_proto_goTypes = nil
	file_device_active_model_proto_depIdxs = nil
}
