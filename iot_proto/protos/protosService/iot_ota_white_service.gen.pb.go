// Code generated by protoc,2022-04-21 14:54:15. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: iot_ota_white_service.gen.proto

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

var File_iot_ota_white_service_gen_proto protoreflect.FileDescriptor

var file_iot_ota_white_service_gen_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x69, 0x6f, 0x74, 0x5f, 0x6f, 0x74, 0x61, 0x5f, 0x77, 0x68, 0x69, 0x74, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x69, 0x6f, 0x74, 0x5f, 0x6f, 0x74, 0x61, 0x5f,
	0x77, 0x68, 0x69, 0x74, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x67, 0x65, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xef, 0x07, 0x0a, 0x12, 0x49, 0x6f, 0x74, 0x4f, 0x74, 0x61,
	0x57, 0x68, 0x69, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x49, 0x6f, 0x74, 0x4f, 0x74, 0x61, 0x57, 0x68, 0x69, 0x74, 0x65, 0x1a, 0x11, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6f, 0x74,
	0x4f, 0x74, 0x61, 0x57, 0x68, 0x69, 0x74, 0x65, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a,
	0x01, 0x2a, 0x12, 0x54, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6f, 0x74, 0x4f, 0x74, 0x61, 0x57, 0x68, 0x69,
	0x74, 0x65, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f,
	0x76, 0x31, 0x2f, 0x69, 0x6f, 0x74, 0x4f, 0x74, 0x61, 0x57, 0x68, 0x69, 0x74, 0x65, 0x2f, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x5c, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x49, 0x6f, 0x74, 0x4f, 0x74, 0x61, 0x57, 0x68, 0x69, 0x74, 0x65, 0x1a, 0x11, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x22, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6f, 0x74,
	0x4f, 0x74, 0x61, 0x57, 0x68, 0x69, 0x74, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x79, 0x49, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x70, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x42, 0x79, 0x49, 0x64, 0x73, 0x12, 0x26, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x49, 0x6f, 0x74, 0x4f, 0x74, 0x61, 0x57, 0x68, 0x69, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6f,
	0x74, 0x4f, 0x74, 0x61, 0x57, 0x68, 0x69, 0x74, 0x65, 0x2f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x42, 0x79, 0x49, 0x64, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x54, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6f, 0x74,
	0x4f, 0x74, 0x61, 0x57, 0x68, 0x69, 0x74, 0x65, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6f, 0x74, 0x4f, 0x74, 0x61, 0x57,
	0x68, 0x69, 0x74, 0x65, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x5a,
	0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x12, 0x14, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6f, 0x74, 0x4f, 0x74, 0x61, 0x57, 0x68, 0x69, 0x74,
	0x65, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x19, 0x2f, 0x76,
	0x31, 0x2f, 0x69, 0x6f, 0x74, 0x4f, 0x74, 0x61, 0x57, 0x68, 0x69, 0x74, 0x65, 0x2f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x3a, 0x01, 0x2a, 0x12, 0x73, 0x0a, 0x0c, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x27, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6f, 0x74, 0x4f, 0x74, 0x61, 0x57, 0x68, 0x69, 0x74, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22, 0x1c,
	0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6f, 0x74, 0x4f, 0x74, 0x61, 0x57, 0x68, 0x69, 0x74, 0x65, 0x2f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x3a, 0x01, 0x2a, 0x12,
	0x69, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6f, 0x74, 0x4f, 0x74, 0x61, 0x57, 0x68, 0x69, 0x74,
	0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x49, 0x6f, 0x74, 0x4f, 0x74, 0x61, 0x57, 0x68, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f,
	0x76, 0x31, 0x2f, 0x69, 0x6f, 0x74, 0x4f, 0x74, 0x61, 0x57, 0x68, 0x69, 0x74, 0x65, 0x2f, 0x66,
	0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x61, 0x0a, 0x04, 0x46, 0x69,
	0x6e, 0x64, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6f, 0x74,
	0x4f, 0x74, 0x61, 0x57, 0x68, 0x69, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x1c,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6f, 0x74, 0x4f, 0x74, 0x61, 0x57,
	0x68, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6f, 0x74, 0x4f, 0x74, 0x61,
	0x57, 0x68, 0x69, 0x74, 0x65, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x68, 0x0a,
	0x05, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x49, 0x6f, 0x74, 0x4f, 0x74, 0x61, 0x57, 0x68, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x49, 0x6f, 0x74, 0x4f, 0x74, 0x61, 0x57, 0x68, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f,
	0x76, 0x31, 0x2f, 0x69, 0x6f, 0x74, 0x4f, 0x74, 0x61, 0x57, 0x68, 0x69, 0x74, 0x65, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_iot_ota_white_service_gen_proto_goTypes = []interface{}{
	(*IotOtaWhite)(nil),                    // 0: service.IotOtaWhite
	(*IotOtaWhiteBatchDeleteRequest)(nil),  // 1: service.IotOtaWhiteBatchDeleteRequest
	(*IotOtaWhiteUpdateFieldsRequest)(nil), // 2: service.IotOtaWhiteUpdateFieldsRequest
	(*IotOtaWhiteFilter)(nil),              // 3: service.IotOtaWhiteFilter
	(*IotOtaWhiteListRequest)(nil),         // 4: service.IotOtaWhiteListRequest
	(*Response)(nil),                       // 5: service.Response
	(*IotOtaWhiteResponse)(nil),            // 6: service.IotOtaWhiteResponse
}
var file_iot_ota_white_service_gen_proto_depIdxs = []int32{
	0,  // 0: service.IotOtaWhiteService.Create:input_type -> service.IotOtaWhite
	0,  // 1: service.IotOtaWhiteService.Delete:input_type -> service.IotOtaWhite
	0,  // 2: service.IotOtaWhiteService.DeleteById:input_type -> service.IotOtaWhite
	1,  // 3: service.IotOtaWhiteService.DeleteByIds:input_type -> service.IotOtaWhiteBatchDeleteRequest
	0,  // 4: service.IotOtaWhiteService.Update:input_type -> service.IotOtaWhite
	0,  // 5: service.IotOtaWhiteService.UpdateAll:input_type -> service.IotOtaWhite
	2,  // 6: service.IotOtaWhiteService.UpdateFields:input_type -> service.IotOtaWhiteUpdateFieldsRequest
	3,  // 7: service.IotOtaWhiteService.FindById:input_type -> service.IotOtaWhiteFilter
	3,  // 8: service.IotOtaWhiteService.Find:input_type -> service.IotOtaWhiteFilter
	4,  // 9: service.IotOtaWhiteService.Lists:input_type -> service.IotOtaWhiteListRequest
	5,  // 10: service.IotOtaWhiteService.Create:output_type -> service.Response
	5,  // 11: service.IotOtaWhiteService.Delete:output_type -> service.Response
	5,  // 12: service.IotOtaWhiteService.DeleteById:output_type -> service.Response
	5,  // 13: service.IotOtaWhiteService.DeleteByIds:output_type -> service.Response
	5,  // 14: service.IotOtaWhiteService.Update:output_type -> service.Response
	5,  // 15: service.IotOtaWhiteService.UpdateAll:output_type -> service.Response
	5,  // 16: service.IotOtaWhiteService.UpdateFields:output_type -> service.Response
	6,  // 17: service.IotOtaWhiteService.FindById:output_type -> service.IotOtaWhiteResponse
	6,  // 18: service.IotOtaWhiteService.Find:output_type -> service.IotOtaWhiteResponse
	6,  // 19: service.IotOtaWhiteService.Lists:output_type -> service.IotOtaWhiteResponse
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_iot_ota_white_service_gen_proto_init() }
func file_iot_ota_white_service_gen_proto_init() {
	if File_iot_ota_white_service_gen_proto != nil {
		return
	}
	file_common_proto_init()
	file_iot_ota_white_model_gen_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_iot_ota_white_service_gen_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_iot_ota_white_service_gen_proto_goTypes,
		DependencyIndexes: file_iot_ota_white_service_gen_proto_depIdxs,
	}.Build()
	File_iot_ota_white_service_gen_proto = out.File
	file_iot_ota_white_service_gen_proto_rawDesc = nil
	file_iot_ota_white_service_gen_proto_goTypes = nil
	file_iot_ota_white_service_gen_proto_depIdxs = nil
}
