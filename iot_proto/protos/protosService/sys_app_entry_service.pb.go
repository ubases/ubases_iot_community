// Code generated by protoc,2022-07-25 09:29:22. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: sys_app_entry_service.proto

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

var File_sys_app_entry_service_proto protoreflect.FileDescriptor

var file_sys_app_entry_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x79, 0x73, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x73, 0x79, 0x73, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd1, 0x0a,
	0x0a, 0x12, 0x53, 0x79, 0x73, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x14,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x79, 0x73, 0x41, 0x70, 0x70, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22,
	0x16, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x73, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x54, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53,
	0x79, 0x73, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x73, 0x41, 0x70,
	0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x3a, 0x01, 0x2a,
	0x12, 0x5c, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x14,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x79, 0x73, 0x41, 0x70, 0x70, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x22,
	0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x73, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x70,
	0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x73, 0x12, 0x26, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x79, 0x73, 0x41, 0x70, 0x70, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20,
	0x22, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x73, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x2f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x73, 0x3a, 0x01, 0x2a,
	0x12, 0x54, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x79, 0x73, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x79, 0x73, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x5a, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x6c, 0x6c, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x79,
	0x73, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1e, 0x22, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x73, 0x41, 0x70, 0x70,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x3a,
	0x01, 0x2a, 0x12, 0x73, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x12, 0x27, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x79, 0x73,
	0x41, 0x70, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22, 0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x73, 0x41,
	0x70, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x69, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x79,
	0x73, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a,
	0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x79, 0x73, 0x41, 0x70, 0x70,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x73, 0x41, 0x70,
	0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x3a,
	0x01, 0x2a, 0x12, 0x61, 0x0a, 0x04, 0x46, 0x69, 0x6e, 0x64, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x79, 0x73, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x53, 0x79, 0x73, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x79, 0x73, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x66, 0x69,
	0x6e, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x68, 0x0a, 0x05, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x1f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x79, 0x73, 0x41, 0x70, 0x70, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x79, 0x73, 0x41, 0x70, 0x70,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x73, 0x41, 0x70,
	0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x12,
	0x77, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x79, 0x12, 0x23, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x79, 0x73, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x79, 0x52, 0x65, 0x71, 0x75, 0x71, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x79, 0x73, 0x41, 0x70, 0x70,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x79, 0x73, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x6c, 0x69,
	0x73, 0x74, 0x44, 0x69, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x7b, 0x0a, 0x0d, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x4c, 0x61, 0x6e, 0x67, 0x73, 0x44, 0x69, 0x79, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x53, 0x79, 0x73, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x53, 0x79, 0x73, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4c, 0x61, 0x6e, 0x67, 0x73,
	0x44, 0x69, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x22, 0x22, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x73, 0x41, 0x70, 0x70, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x2f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x4c, 0x61, 0x6e, 0x67, 0x73, 0x44,
	0x69, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x6a, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x12, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53,
	0x79, 0x73, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x20, 0x22, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x73, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x3a, 0x01,
	0x2a, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_sys_app_entry_service_proto_goTypes = []interface{}{
	(*SysAppEntry)(nil),                    // 0: service.SysAppEntry
	(*SysAppEntryBatchDeleteRequest)(nil),  // 1: service.SysAppEntryBatchDeleteRequest
	(*SysAppEntryUpdateFieldsRequest)(nil), // 2: service.SysAppEntryUpdateFieldsRequest
	(*SysAppEntryFilter)(nil),              // 3: service.SysAppEntryFilter
	(*SysAppEntryListRequest)(nil),         // 4: service.SysAppEntryListRequest
	(*SysAppEntryListDiyRequqest)(nil),     // 5: service.SysAppEntryListDiyRequqest
	(*SysAppEntryBatchRequest)(nil),        // 6: service.SysAppEntryBatchRequest
	(*Response)(nil),                       // 7: service.Response
	(*SysAppEntryResponse)(nil),            // 8: service.SysAppEntryResponse
	(*SysAppEntryListDiyResponse)(nil),     // 9: service.SysAppEntryListDiyResponse
	(*SysAppEntryLangsDiyResponse)(nil),    // 10: service.SysAppEntryLangsDiyResponse
}
var file_sys_app_entry_service_proto_depIdxs = []int32{
	0,  // 0: service.SysAppEntryService.Create:input_type -> service.SysAppEntry
	0,  // 1: service.SysAppEntryService.Delete:input_type -> service.SysAppEntry
	0,  // 2: service.SysAppEntryService.DeleteById:input_type -> service.SysAppEntry
	1,  // 3: service.SysAppEntryService.DeleteByIds:input_type -> service.SysAppEntryBatchDeleteRequest
	0,  // 4: service.SysAppEntryService.Update:input_type -> service.SysAppEntry
	0,  // 5: service.SysAppEntryService.UpdateAll:input_type -> service.SysAppEntry
	2,  // 6: service.SysAppEntryService.UpdateFields:input_type -> service.SysAppEntryUpdateFieldsRequest
	3,  // 7: service.SysAppEntryService.FindById:input_type -> service.SysAppEntryFilter
	3,  // 8: service.SysAppEntryService.Find:input_type -> service.SysAppEntryFilter
	4,  // 9: service.SysAppEntryService.Lists:input_type -> service.SysAppEntryListRequest
	5,  // 10: service.SysAppEntryService.ListDiy:input_type -> service.SysAppEntryListDiyRequqest
	3,  // 11: service.SysAppEntryService.EntryLangsDiy:input_type -> service.SysAppEntryFilter
	6,  // 12: service.SysAppEntryService.CreateBatch:input_type -> service.SysAppEntryBatchRequest
	7,  // 13: service.SysAppEntryService.Create:output_type -> service.Response
	7,  // 14: service.SysAppEntryService.Delete:output_type -> service.Response
	7,  // 15: service.SysAppEntryService.DeleteById:output_type -> service.Response
	7,  // 16: service.SysAppEntryService.DeleteByIds:output_type -> service.Response
	7,  // 17: service.SysAppEntryService.Update:output_type -> service.Response
	7,  // 18: service.SysAppEntryService.UpdateAll:output_type -> service.Response
	7,  // 19: service.SysAppEntryService.UpdateFields:output_type -> service.Response
	8,  // 20: service.SysAppEntryService.FindById:output_type -> service.SysAppEntryResponse
	8,  // 21: service.SysAppEntryService.Find:output_type -> service.SysAppEntryResponse
	8,  // 22: service.SysAppEntryService.Lists:output_type -> service.SysAppEntryResponse
	9,  // 23: service.SysAppEntryService.ListDiy:output_type -> service.SysAppEntryListDiyResponse
	10, // 24: service.SysAppEntryService.EntryLangsDiy:output_type -> service.SysAppEntryLangsDiyResponse
	7,  // 25: service.SysAppEntryService.CreateBatch:output_type -> service.Response
	13, // [13:26] is the sub-list for method output_type
	0,  // [0:13] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_sys_app_entry_service_proto_init() }
func file_sys_app_entry_service_proto_init() {
	if File_sys_app_entry_service_proto != nil {
		return
	}
	file_common_proto_init()
	file_sys_app_entry_model_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sys_app_entry_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sys_app_entry_service_proto_goTypes,
		DependencyIndexes: file_sys_app_entry_service_proto_depIdxs,
	}.Build()
	File_sys_app_entry_service_proto = out.File
	file_sys_app_entry_service_proto_rawDesc = nil
	file_sys_app_entry_service_proto_goTypes = nil
	file_sys_app_entry_service_proto_depIdxs = nil
}
