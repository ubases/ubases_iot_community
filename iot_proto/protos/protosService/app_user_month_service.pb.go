// Code generated by protoc,2022-08-10 22:26:35. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: app_user_month_service.proto

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

var File_app_user_month_service_proto protoreflect.FileDescriptor

var file_app_user_month_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x70, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x61, 0x70, 0x70, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x6f,
	0x6e, 0x74, 0x68, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0x87, 0x08, 0x0a, 0x13, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x55,
	0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72,
	0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12,
	0x56, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68,
	0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x2f, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x5e, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x1a, 0x11, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70,
	0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x42, 0x79, 0x49, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x72, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x79, 0x49, 0x64, 0x73, 0x12, 0x27, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22, 0x1c, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x2f, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x56, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x1a, 0x11, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70,
	0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x3a, 0x01, 0x2a, 0x12, 0x5c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c,
	0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x55, 0x73,
	0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1f, 0x22, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x4d,
	0x6f, 0x6e, 0x74, 0x68, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x3a, 0x01,
	0x2a, 0x12, 0x75, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x12, 0x28, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x55,
	0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x22, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x55,
	0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x6c, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64,
	0x42, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41,
	0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x55,
	0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70,
	0x70, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x42,
	0x79, 0x49, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x64, 0x0a, 0x04, 0x46, 0x69, 0x6e, 0x64, 0x12, 0x1b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72,
	0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x1d, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e,
	0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x4d,
	0x6f, 0x6e, 0x74, 0x68, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x6b, 0x0a, 0x05,
	0x4c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68,
	0x2f, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_app_user_month_service_proto_goTypes = []interface{}{
	(*AppUserMonth)(nil),                    // 0: service.AppUserMonth
	(*AppUserMonthBatchDeleteRequest)(nil),  // 1: service.AppUserMonthBatchDeleteRequest
	(*AppUserMonthUpdateFieldsRequest)(nil), // 2: service.AppUserMonthUpdateFieldsRequest
	(*AppUserMonthFilter)(nil),              // 3: service.AppUserMonthFilter
	(*AppUserMonthListRequest)(nil),         // 4: service.AppUserMonthListRequest
	(*Response)(nil),                        // 5: service.Response
	(*AppUserMonthResponse)(nil),            // 6: service.AppUserMonthResponse
}
var file_app_user_month_service_proto_depIdxs = []int32{
	0,  // 0: service.AppUserMonthService.Create:input_type -> service.AppUserMonth
	0,  // 1: service.AppUserMonthService.Delete:input_type -> service.AppUserMonth
	0,  // 2: service.AppUserMonthService.DeleteById:input_type -> service.AppUserMonth
	1,  // 3: service.AppUserMonthService.DeleteByIds:input_type -> service.AppUserMonthBatchDeleteRequest
	0,  // 4: service.AppUserMonthService.Update:input_type -> service.AppUserMonth
	0,  // 5: service.AppUserMonthService.UpdateAll:input_type -> service.AppUserMonth
	2,  // 6: service.AppUserMonthService.UpdateFields:input_type -> service.AppUserMonthUpdateFieldsRequest
	3,  // 7: service.AppUserMonthService.FindById:input_type -> service.AppUserMonthFilter
	3,  // 8: service.AppUserMonthService.Find:input_type -> service.AppUserMonthFilter
	4,  // 9: service.AppUserMonthService.Lists:input_type -> service.AppUserMonthListRequest
	5,  // 10: service.AppUserMonthService.Create:output_type -> service.Response
	5,  // 11: service.AppUserMonthService.Delete:output_type -> service.Response
	5,  // 12: service.AppUserMonthService.DeleteById:output_type -> service.Response
	5,  // 13: service.AppUserMonthService.DeleteByIds:output_type -> service.Response
	5,  // 14: service.AppUserMonthService.Update:output_type -> service.Response
	5,  // 15: service.AppUserMonthService.UpdateAll:output_type -> service.Response
	5,  // 16: service.AppUserMonthService.UpdateFields:output_type -> service.Response
	6,  // 17: service.AppUserMonthService.FindById:output_type -> service.AppUserMonthResponse
	6,  // 18: service.AppUserMonthService.Find:output_type -> service.AppUserMonthResponse
	6,  // 19: service.AppUserMonthService.Lists:output_type -> service.AppUserMonthResponse
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_app_user_month_service_proto_init() }
func file_app_user_month_service_proto_init() {
	if File_app_user_month_service_proto != nil {
		return
	}
	file_common_proto_init()
	file_app_user_month_model_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_app_user_month_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_user_month_service_proto_goTypes,
		DependencyIndexes: file_app_user_month_service_proto_depIdxs,
	}.Build()
	File_app_user_month_service_proto = out.File
	file_app_user_month_service_proto_rawDesc = nil
	file_app_user_month_service_proto_goTypes = nil
	file_app_user_month_service_proto_depIdxs = nil
}
