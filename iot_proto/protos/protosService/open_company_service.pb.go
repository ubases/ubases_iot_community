// Code generated by protoc,2022-05-06 10:24:50. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: open_company_service.proto

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

var File_open_company_service_proto protoreflect.FileDescriptor

var file_open_company_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x18, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xde, 0x09, 0x0a, 0x12,
	0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x54, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f,
	0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x54, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65,
	0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x5c,
	0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x14, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x22, 0x1a, 0x2f,
	0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2f, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x70, 0x0a, 0x0b,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x73, 0x12, 0x26, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22, 0x1b,
	0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2f,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x54,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x1a, 0x11,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x65, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x3a, 0x01, 0x2a, 0x12, 0x5a, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c,
	0x6c, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x6e,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1e, 0x22, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x3a, 0x01, 0x2a,
	0x12, 0x73, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x12, 0x27, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x21, 0x22, 0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x69, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49,
	0x64, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x6e,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x1c, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x3a, 0x01, 0x2a,
	0x12, 0x61, 0x0a, 0x04, 0x46, 0x69, 0x6e, 0x64, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f,
	0x70, 0x65, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f,
	0x6f, 0x70, 0x65, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2f, 0x66, 0x69, 0x6e, 0x64,
	0x3a, 0x01, 0x2a, 0x12, 0x68, 0x0a, 0x05, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x1f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x71, 0x0a,
	0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x27, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20,
	0x22, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x61, 0x75, 0x74, 0x68, 0x3a, 0x01, 0x2a,
	0x12, 0x7a, 0x0a, 0x12, 0x4f, 0x70, 0x65, 0x6e, 0x44, 0x65, 0x76, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x44, 0x65, 0x76, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x27, 0x22, 0x22, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x76, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x61, 0x75, 0x74, 0x68, 0x3a, 0x01, 0x2a, 0x42, 0x11, 0x5a, 0x0f,
	0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_open_company_service_proto_goTypes = []interface{}{
	(*OpenCompany)(nil),                    // 0: service.OpenCompany
	(*OpenCompanyBatchDeleteRequest)(nil),  // 1: service.OpenCompanyBatchDeleteRequest
	(*OpenCompanyUpdateFieldsRequest)(nil), // 2: service.OpenCompanyUpdateFieldsRequest
	(*OpenCompanyFilter)(nil),              // 3: service.OpenCompanyFilter
	(*OpenCompanyListRequest)(nil),         // 4: service.OpenCompanyListRequest
	(*OpenDevCompanyAuthRequest)(nil),      // 5: service.OpenDevCompanyAuthRequest
	(*Response)(nil),                       // 6: service.Response
	(*OpenCompanyResponse)(nil),            // 7: service.OpenCompanyResponse
}
var file_open_company_service_proto_depIdxs = []int32{
	0,  // 0: service.OpenCompanyService.Create:input_type -> service.OpenCompany
	0,  // 1: service.OpenCompanyService.Delete:input_type -> service.OpenCompany
	0,  // 2: service.OpenCompanyService.DeleteById:input_type -> service.OpenCompany
	1,  // 3: service.OpenCompanyService.DeleteByIds:input_type -> service.OpenCompanyBatchDeleteRequest
	0,  // 4: service.OpenCompanyService.Update:input_type -> service.OpenCompany
	0,  // 5: service.OpenCompanyService.UpdateAll:input_type -> service.OpenCompany
	2,  // 6: service.OpenCompanyService.UpdateFields:input_type -> service.OpenCompanyUpdateFieldsRequest
	3,  // 7: service.OpenCompanyService.FindById:input_type -> service.OpenCompanyFilter
	3,  // 8: service.OpenCompanyService.Find:input_type -> service.OpenCompanyFilter
	4,  // 9: service.OpenCompanyService.Lists:input_type -> service.OpenCompanyListRequest
	2,  // 10: service.OpenCompanyService.CompanyAuth:input_type -> service.OpenCompanyUpdateFieldsRequest
	5,  // 11: service.OpenCompanyService.OpenDevCompanyAuth:input_type -> service.OpenDevCompanyAuthRequest
	6,  // 12: service.OpenCompanyService.Create:output_type -> service.Response
	6,  // 13: service.OpenCompanyService.Delete:output_type -> service.Response
	6,  // 14: service.OpenCompanyService.DeleteById:output_type -> service.Response
	6,  // 15: service.OpenCompanyService.DeleteByIds:output_type -> service.Response
	6,  // 16: service.OpenCompanyService.Update:output_type -> service.Response
	6,  // 17: service.OpenCompanyService.UpdateAll:output_type -> service.Response
	6,  // 18: service.OpenCompanyService.UpdateFields:output_type -> service.Response
	7,  // 19: service.OpenCompanyService.FindById:output_type -> service.OpenCompanyResponse
	7,  // 20: service.OpenCompanyService.Find:output_type -> service.OpenCompanyResponse
	7,  // 21: service.OpenCompanyService.Lists:output_type -> service.OpenCompanyResponse
	6,  // 22: service.OpenCompanyService.CompanyAuth:output_type -> service.Response
	6,  // 23: service.OpenCompanyService.OpenDevCompanyAuth:output_type -> service.Response
	12, // [12:24] is the sub-list for method output_type
	0,  // [0:12] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_open_company_service_proto_init() }
func file_open_company_service_proto_init() {
	if File_open_company_service_proto != nil {
		return
	}
	file_common_proto_init()
	file_open_company_model_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_open_company_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_open_company_service_proto_goTypes,
		DependencyIndexes: file_open_company_service_proto_depIdxs,
	}.Build()
	File_open_company_service_proto = out.File
	file_open_company_service_proto_rawDesc = nil
	file_open_company_service_proto_goTypes = nil
	file_open_company_service_proto_depIdxs = nil
}
