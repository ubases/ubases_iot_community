// Code generated by protoc,2022-05-11 16:38:45. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: opm_thing_model_service.proto

package protosService

import (
	reflect "reflect"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_opm_thing_model_service_proto protoreflect.FileDescriptor

var file_opm_thing_model_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6f, 0x70, 0x6d, 0x5f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x6f, 0x70, 0x6d, 0x5f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0x82, 0x0f, 0x0a, 0x14, 0x4f, 0x70, 0x6d, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f,
	0x70, 0x6d, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x11, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d,
	0x54, 0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x3a, 0x01, 0x2a, 0x12, 0x58, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x16,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x54, 0x68, 0x69, 0x6e,
	0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1d, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x60,
	0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x16, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22,
	0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x3a, 0x01, 0x2a,
	0x12, 0x74, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x73, 0x12,
	0x28, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x54, 0x68, 0x69,
	0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x22, 0x22, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x54, 0x68, 0x69,
	0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79,
	0x49, 0x64, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x58, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x54, 0x68,
	0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x54, 0x68, 0x69, 0x6e,
	0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a,
	0x12, 0x5e, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x12, 0x16, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x54, 0x68, 0x69, 0x6e, 0x67,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20,
	0x22, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x3a, 0x01, 0x2a,
	0x12, 0x77, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x12, 0x29, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x54, 0x68,
	0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x22, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x54,
	0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x6f, 0x0a, 0x08, 0x46, 0x69, 0x6e,
	0x64, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x4f, 0x70, 0x6d, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70,
	0x6d, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x22, 0x1a, 0x2f, 0x76, 0x31,
	0x2f, 0x6f, 0x70, 0x6d, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x66,
	0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x67, 0x0a, 0x04, 0x46, 0x69,
	0x6e, 0x64, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d,
	0x54, 0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x54, 0x68,
	0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70,
	0x6d, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x66, 0x69, 0x6e, 0x64,
	0x3a, 0x01, 0x2a, 0x12, 0x6e, 0x0a, 0x05, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x21, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x54, 0x68, 0x69,
	0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d,
	0x54, 0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x73,
	0x3a, 0x01, 0x2a, 0x12, 0x94, 0x01, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54,
	0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x26, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x54,
	0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x28, 0x22, 0x23, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x54, 0x68, 0x69, 0x6e, 0x67,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x68, 0x69,
	0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x3a, 0x01, 0x2a, 0x12, 0x96, 0x01, 0x0a, 0x12, 0x53,
	0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x12, 0x26, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x54,
	0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x22, 0x24, 0x2f, 0x76, 0x31, 0x2f,
	0x6f, 0x70, 0x6d, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74,
	0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x3a, 0x01, 0x2a, 0x12, 0x79, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x68, 0x69,
	0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x22, 0x22, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d,
	0x54, 0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x3a, 0x01, 0x2a, 0x12, 0x7a,
	0x0a, 0x17, 0x53, 0x65, 0x74, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x53, 0x63, 0x65, 0x6e, 0x65, 0x46, 0x75, 0x6e, 0x63, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x34, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x22, 0x29, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x73, 0x65, 0x74, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x63,
	0x65, 0x6e, 0x65, 0x46, 0x75, 0x6e, 0x63, 0x3a, 0x01, 0x2a, 0x12, 0x70, 0x0a, 0x12, 0x53, 0x65,
	0x74, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x75, 0x6e, 0x63,
	0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x54, 0x68,
	0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x29, 0x22, 0x24, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x54, 0x68, 0x69, 0x6e,
	0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x3a, 0x01, 0x2a, 0x12, 0x64, 0x0a, 0x0c,
	0x53, 0x65, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x16, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x22,
	0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x73, 0x65, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x3a,
	0x01, 0x2a, 0x12, 0x62, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x53, 0x6f, 0x72,
	0x74, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x54,
	0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x22, 0x22, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x54, 0x68, 0x69,
	0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x65, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x53,
	0x6f, 0x72, 0x74, 0x3a, 0x01, 0x2a, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_opm_thing_model_service_proto_goTypes = []interface{}{
	(*OpmThingModel)(nil),                    // 0: service.OpmThingModel
	(*OpmThingModelBatchDeleteRequest)(nil),  // 1: service.OpmThingModelBatchDeleteRequest
	(*OpmThingModelUpdateFieldsRequest)(nil), // 2: service.OpmThingModelUpdateFieldsRequest
	(*OpmThingModelFilter)(nil),              // 3: service.OpmThingModelFilter
	(*OpmThingModelListRequest)(nil),         // 4: service.OpmThingModelListRequest
	(*OpmThingModelByProductRequest)(nil),    // 5: service.OpmThingModelByProductRequest
	(*OpmThingModelDeleteRequest)(nil),       // 6: service.OpmThingModelDeleteRequest
	(*Response)(nil),                         // 7: service.Response
	(*OpmThingModelResponse)(nil),            // 8: service.OpmThingModelResponse
	(*OpmThingModelByProductResponse)(nil),   // 9: service.OpmThingModelByProductResponse
}
var file_opm_thing_model_service_proto_depIdxs = []int32{
	0,  // 0: service.OpmThingModelService.Create:input_type -> service.OpmThingModel
	0,  // 1: service.OpmThingModelService.Delete:input_type -> service.OpmThingModel
	0,  // 2: service.OpmThingModelService.DeleteById:input_type -> service.OpmThingModel
	1,  // 3: service.OpmThingModelService.DeleteByIds:input_type -> service.OpmThingModelBatchDeleteRequest
	0,  // 4: service.OpmThingModelService.Update:input_type -> service.OpmThingModel
	0,  // 5: service.OpmThingModelService.UpdateAll:input_type -> service.OpmThingModel
	2,  // 6: service.OpmThingModelService.UpdateFields:input_type -> service.OpmThingModelUpdateFieldsRequest
	3,  // 7: service.OpmThingModelService.FindById:input_type -> service.OpmThingModelFilter
	3,  // 8: service.OpmThingModelService.Find:input_type -> service.OpmThingModelFilter
	4,  // 9: service.OpmThingModelService.Lists:input_type -> service.OpmThingModelListRequest
	5,  // 10: service.OpmThingModelService.ProductThingModel:input_type -> service.OpmThingModelByProductRequest
	5,  // 11: service.OpmThingModelService.StandardThingModel:input_type -> service.OpmThingModelByProductRequest
	6,  // 12: service.OpmThingModelService.DeleteThingModel:input_type -> service.OpmThingModelDeleteRequest
	0,  // 13: service.OpmThingModelService.SetThingsModelSceneFunc:input_type -> service.OpmThingModel
	0,  // 14: service.OpmThingModelService.SetAppointmentFunc:input_type -> service.OpmThingModel
	0,  // 15: service.OpmThingModelService.SetFuncLevel:input_type -> service.OpmThingModel
	0,  // 16: service.OpmThingModelService.SetFuncSort:input_type -> service.OpmThingModel
	7,  // 17: service.OpmThingModelService.Create:output_type -> service.Response
	7,  // 18: service.OpmThingModelService.Delete:output_type -> service.Response
	7,  // 19: service.OpmThingModelService.DeleteById:output_type -> service.Response
	7,  // 20: service.OpmThingModelService.DeleteByIds:output_type -> service.Response
	7,  // 21: service.OpmThingModelService.Update:output_type -> service.Response
	7,  // 22: service.OpmThingModelService.UpdateAll:output_type -> service.Response
	7,  // 23: service.OpmThingModelService.UpdateFields:output_type -> service.Response
	8,  // 24: service.OpmThingModelService.FindById:output_type -> service.OpmThingModelResponse
	8,  // 25: service.OpmThingModelService.Find:output_type -> service.OpmThingModelResponse
	8,  // 26: service.OpmThingModelService.Lists:output_type -> service.OpmThingModelResponse
	9,  // 27: service.OpmThingModelService.ProductThingModel:output_type -> service.OpmThingModelByProductResponse
	9,  // 28: service.OpmThingModelService.StandardThingModel:output_type -> service.OpmThingModelByProductResponse
	7,  // 29: service.OpmThingModelService.DeleteThingModel:output_type -> service.Response
	7,  // 30: service.OpmThingModelService.SetThingsModelSceneFunc:output_type -> service.Response
	7,  // 31: service.OpmThingModelService.SetAppointmentFunc:output_type -> service.Response
	7,  // 32: service.OpmThingModelService.SetFuncLevel:output_type -> service.Response
	7,  // 33: service.OpmThingModelService.SetFuncSort:output_type -> service.Response
	17, // [17:34] is the sub-list for method output_type
	0,  // [0:17] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_opm_thing_model_service_proto_init() }
func file_opm_thing_model_service_proto_init() {
	if File_opm_thing_model_service_proto != nil {
		return
	}
	file_common_proto_init()
	file_opm_thing_model_model_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_opm_thing_model_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_opm_thing_model_service_proto_goTypes,
		DependencyIndexes: file_opm_thing_model_service_proto_depIdxs,
	}.Build()
	File_opm_thing_model_service_proto = out.File
	file_opm_thing_model_service_proto_rawDesc = nil
	file_opm_thing_model_service_proto_goTypes = nil
	file_opm_thing_model_service_proto_depIdxs = nil
}
