// Code generated by protoc,2022-04-21 14:24:40. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: iot_device_triad_service.proto

package protosService

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_iot_device_triad_service_proto protoreflect.FileDescriptor

var file_iot_device_triad_service_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x69, 0x6f, 0x74, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x72, 0x69,
	0x61, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x69, 0x6f, 0x74, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x74, 0x72, 0x69, 0x61, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0xa5, 0x0e, 0x0a, 0x15, 0x49, 0x6f, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x54, 0x72, 0x69, 0x61, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x49, 0x6f, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x72, 0x69, 0x61, 0x64,
	0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x19, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x6f, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x72, 0x69, 0x61, 0x64, 0x2f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x5a, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6f,
	0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x72, 0x69, 0x61, 0x64, 0x1a, 0x11, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6f, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x72, 0x69, 0x61, 0x64, 0x2f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x62, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6f,
	0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x72, 0x69, 0x61, 0x64, 0x1a, 0x11, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x22, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6f, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x72, 0x69, 0x61, 0x64, 0x2f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x76, 0x0a, 0x0b, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x73, 0x12, 0x29, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x49, 0x6f, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x72, 0x69, 0x61,
	0x64, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x22, 0x1e,
	0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6f, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x72, 0x69,
	0x61, 0x64, 0x2f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x73, 0x3a, 0x01,
	0x2a, 0x12, 0x5a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6f, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54,
	0x72, 0x69, 0x61, 0x64, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22,
	0x19, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6f, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x72,
	0x69, 0x61, 0x64, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x60, 0x0a,
	0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6f, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x72,
	0x69, 0x61, 0x64, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22, 0x1c,
	0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6f, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x72, 0x69,
	0x61, 0x64, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x3a, 0x01, 0x2a, 0x12,
	0x79, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12,
	0x2a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6f, 0x74, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x72, 0x69, 0x61, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x22, 0x1f, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6f, 0x74, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x72, 0x69, 0x61, 0x64, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x72, 0x0a, 0x08, 0x46, 0x69,
	0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x49, 0x6f, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x72, 0x69, 0x61, 0x64, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x49, 0x6f, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x72, 0x69, 0x61, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22, 0x1b,
	0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6f, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x72, 0x69,
	0x61, 0x64, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x6a,
	0x0a, 0x04, 0x46, 0x69, 0x6e, 0x64, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x49, 0x6f, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x72, 0x69, 0x61, 0x64, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x49, 0x6f, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x72, 0x69, 0x61, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17,
	0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6f, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x72, 0x69,
	0x61, 0x64, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x71, 0x0a, 0x05, 0x4c, 0x69,
	0x73, 0x74, 0x73, 0x12, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6f,
	0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x72, 0x69, 0x61, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x49, 0x6f, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x72, 0x69, 0x61, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d,
	0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6f, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54,
	0x72, 0x69, 0x61, 0x64, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x7a, 0x0a,
	0x14, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x54, 0x72, 0x69, 0x61, 0x64, 0x12, 0x26, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x49, 0x6f, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x72, 0x69, 0x61, 0x64, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22, 0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6f,
	0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x72, 0x69, 0x61, 0x64, 0x2f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x76, 0x0a, 0x1d, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x72, 0x69, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x79, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6f, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x72,
	0x69, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6f, 0x74, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x72, 0x69, 0x61, 0x64, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x01,
	0x2a, 0x12, 0x77, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x54, 0x72, 0x69, 0x61, 0x64, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x54, 0x72, 0x69, 0x61, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a,
	0x22, 0x25, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6f, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54,
	0x72, 0x69, 0x61, 0x64, 0x2f, 0x73, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x54, 0x72, 0x69, 0x61, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x74, 0x0a, 0x0f, 0x42, 0x69,
	0x6e, 0x64, 0x54, 0x65, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x69, 0x6e, 0x64, 0x54, 0x65, 0x73, 0x74,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x22, 0x22, 0x2f, 0x76, 0x31, 0x2f, 0x69,
	0x6f, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x72, 0x69, 0x61, 0x64, 0x2f, 0x62, 0x69,
	0x6e, 0x64, 0x54, 0x65, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x01, 0x2a,
	0x12, 0x82, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x42, 0x69,
	0x6e, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x72, 0x69, 0x61, 0x64, 0x12, 0x26, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6f, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x54, 0x72, 0x69, 0x61, 0x64, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25,
	0x22, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6f, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54,
	0x72, 0x69, 0x61, 0x64, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x42, 0x69,
	0x6e, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x83, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x45, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x49, 0x6f, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x72, 0x69, 0x61,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6f, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x54, 0x72, 0x69, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x26, 0x22, 0x21, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6f, 0x74, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x54, 0x72, 0x69, 0x61, 0x64, 0x2f, 0x73, 0x65, 0x74, 0x45, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x01, 0x2a, 0x42, 0x11, 0x5a, 0x0f, 0x2e,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_iot_device_triad_service_proto_goTypes = []interface{}{
	(*IotDeviceTriad)(nil),                    // 0: service.IotDeviceTriad
	(*IotDeviceTriadBatchDeleteRequest)(nil),  // 1: service.IotDeviceTriadBatchDeleteRequest
	(*IotDeviceTriadUpdateFieldsRequest)(nil), // 2: service.IotDeviceTriadUpdateFieldsRequest
	(*IotDeviceTriadFilter)(nil),              // 3: service.IotDeviceTriadFilter
	(*IotDeviceTriadListRequest)(nil),         // 4: service.IotDeviceTriadListRequest
	(*IotDeviceTriadGenerateRequest)(nil),     // 5: service.IotDeviceTriadGenerateRequest
	(*SetTestTriadRequest)(nil),               // 6: service.SetTestTriadRequest
	(*BindTestAccountRequest)(nil),            // 7: service.BindTestAccountRequest
	(*Response)(nil),                          // 8: service.Response
	(*IotDeviceTriadResponse)(nil),            // 9: service.IotDeviceTriadResponse
}
var file_iot_device_triad_service_proto_depIdxs = []int32{
	0,  // 0: service.IotDeviceTriadService.Create:input_type -> service.IotDeviceTriad
	0,  // 1: service.IotDeviceTriadService.Delete:input_type -> service.IotDeviceTriad
	0,  // 2: service.IotDeviceTriadService.DeleteById:input_type -> service.IotDeviceTriad
	1,  // 3: service.IotDeviceTriadService.DeleteByIds:input_type -> service.IotDeviceTriadBatchDeleteRequest
	0,  // 4: service.IotDeviceTriadService.Update:input_type -> service.IotDeviceTriad
	0,  // 5: service.IotDeviceTriadService.UpdateAll:input_type -> service.IotDeviceTriad
	2,  // 6: service.IotDeviceTriadService.UpdateFields:input_type -> service.IotDeviceTriadUpdateFieldsRequest
	3,  // 7: service.IotDeviceTriadService.FindById:input_type -> service.IotDeviceTriadFilter
	3,  // 8: service.IotDeviceTriadService.Find:input_type -> service.IotDeviceTriadFilter
	4,  // 9: service.IotDeviceTriadService.Lists:input_type -> service.IotDeviceTriadListRequest
	5,  // 10: service.IotDeviceTriadService.GeneratorDeviceTriad:input_type -> service.IotDeviceTriadGenerateRequest
	3,  // 11: service.IotDeviceTriadService.GetDeviceTriadCountByTenantId:input_type -> service.IotDeviceTriadFilter
	6,  // 12: service.IotDeviceTriadService.SetTestDeviceTriad:input_type -> service.SetTestTriadRequest
	7,  // 13: service.IotDeviceTriadService.BindTestAccount:input_type -> service.BindTestAccountRequest
	5,  // 14: service.IotDeviceTriadService.CreateAndBindDeviceTriad:input_type -> service.IotDeviceTriadGenerateRequest
	4,  // 15: service.IotDeviceTriadService.SetExportCount:input_type -> service.IotDeviceTriadListRequest
	8,  // 16: service.IotDeviceTriadService.Create:output_type -> service.Response
	8,  // 17: service.IotDeviceTriadService.Delete:output_type -> service.Response
	8,  // 18: service.IotDeviceTriadService.DeleteById:output_type -> service.Response
	8,  // 19: service.IotDeviceTriadService.DeleteByIds:output_type -> service.Response
	8,  // 20: service.IotDeviceTriadService.Update:output_type -> service.Response
	8,  // 21: service.IotDeviceTriadService.UpdateAll:output_type -> service.Response
	8,  // 22: service.IotDeviceTriadService.UpdateFields:output_type -> service.Response
	9,  // 23: service.IotDeviceTriadService.FindById:output_type -> service.IotDeviceTriadResponse
	9,  // 24: service.IotDeviceTriadService.Find:output_type -> service.IotDeviceTriadResponse
	9,  // 25: service.IotDeviceTriadService.Lists:output_type -> service.IotDeviceTriadResponse
	8,  // 26: service.IotDeviceTriadService.GeneratorDeviceTriad:output_type -> service.Response
	8,  // 27: service.IotDeviceTriadService.GetDeviceTriadCountByTenantId:output_type -> service.Response
	8,  // 28: service.IotDeviceTriadService.SetTestDeviceTriad:output_type -> service.Response
	8,  // 29: service.IotDeviceTriadService.BindTestAccount:output_type -> service.Response
	8,  // 30: service.IotDeviceTriadService.CreateAndBindDeviceTriad:output_type -> service.Response
	9,  // 31: service.IotDeviceTriadService.SetExportCount:output_type -> service.IotDeviceTriadResponse
	16, // [16:32] is the sub-list for method output_type
	0,  // [0:16] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_iot_device_triad_service_proto_init() }
func file_iot_device_triad_service_proto_init() {
	if File_iot_device_triad_service_proto != nil {
		return
	}
	file_common_proto_init()
	file_iot_device_triad_model_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_iot_device_triad_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_iot_device_triad_service_proto_goTypes,
		DependencyIndexes: file_iot_device_triad_service_proto_depIdxs,
	}.Build()
	File_iot_device_triad_service_proto = out.File
	file_iot_device_triad_service_proto_rawDesc = nil
	file_iot_device_triad_service_proto_goTypes = nil
	file_iot_device_triad_service_proto_depIdxs = nil
}
