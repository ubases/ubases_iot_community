// Code generated by protoc,2022-11-11 10:46:48. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: opm_product_materials_service.proto

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

var File_opm_product_materials_service_proto protoreflect.FileDescriptor

var file_opm_product_materials_service_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6f, 0x70, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6d, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0c,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x6f, 0x70, 0x6d, 0x5f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x97, 0x0a,
	0x0a, 0x1a, 0x4f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x73, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x22,
	0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a,
	0x01, 0x2a, 0x12, 0x64, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x23, 0x22, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2f, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x6c, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x73, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x22,
	0x22, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x79, 0x49, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x80, 0x01, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x79, 0x49, 0x64, 0x73, 0x12, 0x2e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x73, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x28, 0x22, 0x23, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2f, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x79, 0x49, 0x64, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x64, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70,
	0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x73, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x22, 0x1e, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12,
	0x6a, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x12, 0x1c, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x26, 0x22, 0x21, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x3a, 0x01, 0x2a, 0x12, 0x83, 0x01, 0x0a, 0x0c,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x2f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x22, 0x24, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70,
	0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x73, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x3a, 0x01,
	0x2a, 0x12, 0x81, 0x01, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x12, 0x22,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x1a, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25,
	0x22, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x42, 0x79,
	0x49, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x79, 0x0a, 0x04, 0x46, 0x69, 0x6e, 0x64, 0x12, 0x22, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x1a, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22,
	0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x3a, 0x01, 0x2a,
	0x12, 0x80, 0x01, 0x0a, 0x05, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x27, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70,
	0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x22, 0x22, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x73,
	0x3a, 0x01, 0x2a, 0x12, 0x62, 0x0a, 0x05, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x12, 0x1c, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x22, 0x22, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2f, 0x63,
	0x6c, 0x69, 0x63, 0x6b, 0x3a, 0x01, 0x2a, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_opm_product_materials_service_proto_goTypes = []interface{}{
	(*OpmProductMaterials)(nil),                    // 0: service.OpmProductMaterials
	(*OpmProductMaterialsBatchDeleteRequest)(nil),  // 1: service.OpmProductMaterialsBatchDeleteRequest
	(*OpmProductMaterialsUpdateFieldsRequest)(nil), // 2: service.OpmProductMaterialsUpdateFieldsRequest
	(*OpmProductMaterialsFilter)(nil),              // 3: service.OpmProductMaterialsFilter
	(*OpmProductMaterialsListRequest)(nil),         // 4: service.OpmProductMaterialsListRequest
	(*Response)(nil),                               // 5: service.Response
	(*OpmProductMaterialsResponse)(nil),            // 6: service.OpmProductMaterialsResponse
}
var file_opm_product_materials_service_proto_depIdxs = []int32{
	0,  // 0: service.OpmProductMaterialsService.Create:input_type -> service.OpmProductMaterials
	0,  // 1: service.OpmProductMaterialsService.Delete:input_type -> service.OpmProductMaterials
	0,  // 2: service.OpmProductMaterialsService.DeleteById:input_type -> service.OpmProductMaterials
	1,  // 3: service.OpmProductMaterialsService.DeleteByIds:input_type -> service.OpmProductMaterialsBatchDeleteRequest
	0,  // 4: service.OpmProductMaterialsService.Update:input_type -> service.OpmProductMaterials
	0,  // 5: service.OpmProductMaterialsService.UpdateAll:input_type -> service.OpmProductMaterials
	2,  // 6: service.OpmProductMaterialsService.UpdateFields:input_type -> service.OpmProductMaterialsUpdateFieldsRequest
	3,  // 7: service.OpmProductMaterialsService.FindById:input_type -> service.OpmProductMaterialsFilter
	3,  // 8: service.OpmProductMaterialsService.Find:input_type -> service.OpmProductMaterialsFilter
	4,  // 9: service.OpmProductMaterialsService.Lists:input_type -> service.OpmProductMaterialsListRequest
	0,  // 10: service.OpmProductMaterialsService.Click:input_type -> service.OpmProductMaterials
	5,  // 11: service.OpmProductMaterialsService.Create:output_type -> service.Response
	5,  // 12: service.OpmProductMaterialsService.Delete:output_type -> service.Response
	5,  // 13: service.OpmProductMaterialsService.DeleteById:output_type -> service.Response
	5,  // 14: service.OpmProductMaterialsService.DeleteByIds:output_type -> service.Response
	5,  // 15: service.OpmProductMaterialsService.Update:output_type -> service.Response
	5,  // 16: service.OpmProductMaterialsService.UpdateAll:output_type -> service.Response
	5,  // 17: service.OpmProductMaterialsService.UpdateFields:output_type -> service.Response
	6,  // 18: service.OpmProductMaterialsService.FindById:output_type -> service.OpmProductMaterialsResponse
	6,  // 19: service.OpmProductMaterialsService.Find:output_type -> service.OpmProductMaterialsResponse
	6,  // 20: service.OpmProductMaterialsService.Lists:output_type -> service.OpmProductMaterialsResponse
	5,  // 21: service.OpmProductMaterialsService.Click:output_type -> service.Response
	11, // [11:22] is the sub-list for method output_type
	0,  // [0:11] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_opm_product_materials_service_proto_init() }
func file_opm_product_materials_service_proto_init() {
	if File_opm_product_materials_service_proto != nil {
		return
	}
	file_common_proto_init()
	file_opm_product_materials_model_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_opm_product_materials_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_opm_product_materials_service_proto_goTypes,
		DependencyIndexes: file_opm_product_materials_service_proto_depIdxs,
	}.Build()
	File_opm_product_materials_service_proto = out.File
	file_opm_product_materials_service_proto_rawDesc = nil
	file_opm_product_materials_service_proto_goTypes = nil
	file_opm_product_materials_service_proto_depIdxs = nil
}
