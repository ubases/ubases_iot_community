// Code generated by protoc,2022-05-07 16:41:11. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: opm_product_service.proto

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

var File_opm_product_service_proto protoreflect.FileDescriptor

var file_opm_product_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6f, 0x70, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x6f, 0x70, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x6f, 0x70, 0x6d, 0x5f, 0x74,
	0x68, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xf9, 0x10, 0x0a, 0x11, 0x4f, 0x70, 0x6d, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a,
	0x12, 0x52, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a,
	0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f,
	0x6f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x3a, 0x01, 0x2a, 0x12, 0x5a, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79,
	0x49, 0x64, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1e, 0x22, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x3a, 0x01, 0x2a,
	0x12, 0x6e, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x73, 0x12,
	0x25, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1f, 0x22, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x73, 0x3a, 0x01, 0x2a,
	0x12, 0x52, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a,
	0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f,
	0x6f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x3a, 0x01, 0x2a, 0x12, 0x58, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c,
	0x6c, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1d, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x3a, 0x01, 0x2a, 0x12, 0x71,
	0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x26,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x20, 0x22, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x3a, 0x01,
	0x2a, 0x12, 0x66, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x12, 0x19, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f,
	0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x66, 0x69,
	0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x5e, 0x0a, 0x04, 0x46, 0x69, 0x6e,
	0x64, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x1b, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x18, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x65, 0x0a, 0x05, 0x4c, 0x69, 0x73,
	0x74, 0x73, 0x12, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x3a, 0x01, 0x2a,
	0x12, 0x86, 0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x61, 0x6e, 0x65,
	0x6c, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x49, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x50, 0x6d, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x61, 0x6e, 0x65,
	0x6c, 0x73, 0x56, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x25, 0x22, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x61, 0x6e, 0x65,
	0x6c, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x6c, 0x0a, 0x0b, 0x4d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6d,
	0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x56, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x22, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70,
	0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x6e, 0x0a, 0x08, 0x41, 0x70, 0x70, 0x4c, 0x69,
	0x73, 0x74, 0x73, 0x12, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x70,
	0x70, 0x4f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x76, 0x31,
	0x2f, 0x6f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x41, 0x70, 0x70, 0x4c,
	0x69, 0x73, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x74, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x42,
	0x79, 0x41, 0x6c, 0x6c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1d, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41,
	0x6c, 0x6c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1c, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x9b, 0x01,
	0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x73, 0x55,
	0x72, 0x6c, 0x41, 0x6e, 0x64, 0x4d, 0x64, 0x35, 0x12, 0x26, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x73,
	0x55, 0x72, 0x6c, 0x41, 0x6e, 0x64, 0x4d, 0x64, 0x35, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x27, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x73, 0x55, 0x72, 0x6c, 0x41, 0x6e, 0x64, 0x4d, 0x64,
	0x35, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x2a, 0x22, 0x25, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x73, 0x55,
	0x72, 0x6c, 0x41, 0x6e, 0x64, 0x4d, 0x64, 0x35, 0x3a, 0x01, 0x2a, 0x12, 0x80, 0x01, 0x0a, 0x11,
	0x4c, 0x69, 0x73, 0x74, 0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x73, 0x12, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f,
	0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x22, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x42,
	0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x7a,
	0x0a, 0x1a, 0x52, 0x65, 0x73, 0x65, 0x74, 0x4f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x13, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x34, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x22, 0x29, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x72, 0x65, 0x73,
	0x65, 0x74, 0x4f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x68, 0x69, 0x6e,
	0x67, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x3a, 0x01, 0x2a, 0x12, 0x8a, 0x01, 0x0a, 0x16, 0x50,
	0x61, 0x6e, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x64, 0x73, 0x12, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x22, 0x25, 0x2f,
	0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x70, 0x61,
	0x6e, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x64, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x98, 0x01, 0x0a, 0x17, 0x4d, 0x65, 0x72, 0x67,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x12, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4f, 0x70, 0x6d, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x79,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x31, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x22, 0x26, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x3a,
	0x01, 0x2a, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_opm_product_service_proto_goTypes = []interface{}{
	(*OpmProduct)(nil),                     // 0: service.OpmProduct
	(*OpmProductBatchDeleteRequest)(nil),   // 1: service.OpmProductBatchDeleteRequest
	(*OpmProductUpdateFieldsRequest)(nil),  // 2: service.OpmProductUpdateFieldsRequest
	(*OpmProductFilter)(nil),               // 3: service.OpmProductFilter
	(*OpmProductListRequest)(nil),          // 4: service.OpmProductListRequest
	(*ControlPanelIdsRequest)(nil),         // 5: service.ControlPanelIdsRequest
	(*ModuleIdsRequest)(nil),               // 6: service.ModuleIdsRequest
	(*AppOpmProductListRequest)(nil),       // 7: service.AppOpmProductListRequest
	(*OpmProductPrimarykey)(nil),           // 8: service.OpmProductPrimarykey
	(*ControlPanelsUrlAndMd5Request)(nil),  // 9: service.ControlPanelsUrlAndMd5Request
	(*ListsByProductIdsRequest)(nil),       // 10: service.ListsByProductIdsRequest
	(*Response)(nil),                       // 11: service.Response
	(*OpmProductResponse)(nil),             // 12: service.OpmProductResponse
	(*PmControlPanelsVoResponse)(nil),      // 13: service.PmControlPanelsVoResponse
	(*PmModuleVoResponse)(nil),             // 14: service.PmModuleVoResponse
	(*OpmProductAllDetails)(nil),           // 15: service.OpmProductAllDetails
	(*ControlPanelsUrlAndMd5Response)(nil), // 16: service.ControlPanelsUrlAndMd5Response
	(*OpmThingModelByProductResponse)(nil), // 17: service.OpmThingModelByProductResponse
}
var file_opm_product_service_proto_depIdxs = []int32{
	0,  // 0: service.OpmProductService.Create:input_type -> service.OpmProduct
	0,  // 1: service.OpmProductService.Delete:input_type -> service.OpmProduct
	0,  // 2: service.OpmProductService.DeleteById:input_type -> service.OpmProduct
	1,  // 3: service.OpmProductService.DeleteByIds:input_type -> service.OpmProductBatchDeleteRequest
	0,  // 4: service.OpmProductService.Update:input_type -> service.OpmProduct
	0,  // 5: service.OpmProductService.UpdateAll:input_type -> service.OpmProduct
	2,  // 6: service.OpmProductService.UpdateFields:input_type -> service.OpmProductUpdateFieldsRequest
	3,  // 7: service.OpmProductService.FindById:input_type -> service.OpmProductFilter
	3,  // 8: service.OpmProductService.Find:input_type -> service.OpmProductFilter
	4,  // 9: service.OpmProductService.Lists:input_type -> service.OpmProductListRequest
	5,  // 10: service.OpmProductService.ControlPanelsLists:input_type -> service.ControlPanelIdsRequest
	6,  // 11: service.OpmProductService.ModuleLists:input_type -> service.ModuleIdsRequest
	7,  // 12: service.OpmProductService.AppLists:input_type -> service.AppOpmProductListRequest
	8,  // 13: service.OpmProductService.FindByAllDetails:input_type -> service.OpmProductPrimarykey
	9,  // 14: service.OpmProductService.ControlPanelsUrlAndMd5:input_type -> service.ControlPanelsUrlAndMd5Request
	10, // 15: service.OpmProductService.ListsByProductIds:input_type -> service.ListsByProductIdsRequest
	0,  // 16: service.OpmProductService.ResetOpmProductThingsModel:input_type -> service.OpmProduct
	10, // 17: service.OpmProductService.PanelListsByProductIds:input_type -> service.ListsByProductIdsRequest
	10, // 18: service.OpmProductService.MergeProductThingsModel:input_type -> service.ListsByProductIdsRequest
	11, // 19: service.OpmProductService.Create:output_type -> service.Response
	11, // 20: service.OpmProductService.Delete:output_type -> service.Response
	11, // 21: service.OpmProductService.DeleteById:output_type -> service.Response
	11, // 22: service.OpmProductService.DeleteByIds:output_type -> service.Response
	11, // 23: service.OpmProductService.Update:output_type -> service.Response
	11, // 24: service.OpmProductService.UpdateAll:output_type -> service.Response
	11, // 25: service.OpmProductService.UpdateFields:output_type -> service.Response
	12, // 26: service.OpmProductService.FindById:output_type -> service.OpmProductResponse
	12, // 27: service.OpmProductService.Find:output_type -> service.OpmProductResponse
	12, // 28: service.OpmProductService.Lists:output_type -> service.OpmProductResponse
	13, // 29: service.OpmProductService.ControlPanelsLists:output_type -> service.PmControlPanelsVoResponse
	14, // 30: service.OpmProductService.ModuleLists:output_type -> service.PmModuleVoResponse
	12, // 31: service.OpmProductService.AppLists:output_type -> service.OpmProductResponse
	15, // 32: service.OpmProductService.FindByAllDetails:output_type -> service.OpmProductAllDetails
	16, // 33: service.OpmProductService.ControlPanelsUrlAndMd5:output_type -> service.ControlPanelsUrlAndMd5Response
	12, // 34: service.OpmProductService.ListsByProductIds:output_type -> service.OpmProductResponse
	11, // 35: service.OpmProductService.ResetOpmProductThingsModel:output_type -> service.Response
	12, // 36: service.OpmProductService.PanelListsByProductIds:output_type -> service.OpmProductResponse
	17, // 37: service.OpmProductService.MergeProductThingsModel:output_type -> service.OpmThingModelByProductResponse
	19, // [19:38] is the sub-list for method output_type
	0,  // [0:19] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_opm_product_service_proto_init() }
func file_opm_product_service_proto_init() {
	if File_opm_product_service_proto != nil {
		return
	}
	file_common_proto_init()
	file_opm_product_model_proto_init()
	file_opm_thing_model_model_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_opm_product_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_opm_product_service_proto_goTypes,
		DependencyIndexes: file_opm_product_service_proto_depIdxs,
	}.Build()
	File_opm_product_service_proto = out.File
	file_opm_product_service_proto_rawDesc = nil
	file_opm_product_service_proto_goTypes = nil
	file_opm_product_service_proto_depIdxs = nil
}
