// Code generated by protoc,2022-08-31 14:21:29. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: opm_voice_service.gen.proto

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

var File_opm_voice_service_gen_proto protoreflect.FileDescriptor

var file_opm_voice_service_gen_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6f, 0x70, 0x6d, 0x5f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x6f, 0x70, 0x6d, 0x5f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa7, 0x07,
	0x0a, 0x0f, 0x4f, 0x70, 0x6d, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x1a, 0x11,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x6d, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01,
	0x2a, 0x12, 0x4e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x1a, 0x11,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x6d, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x3a, 0x01,
	0x2a, 0x12, 0x56, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12,
	0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x56, 0x6f, 0x69,
	0x63, 0x65, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f,
	0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x6a, 0x0a, 0x0b, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x73, 0x12, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70,
	0x6d, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x2f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49,
	0x64, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x4e, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x56, 0x6f, 0x69,
	0x63, 0x65, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f,
	0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x2f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x54, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x6c, 0x6c, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d,
	0x56, 0x6f, 0x69, 0x63, 0x65, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b,
	0x22, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x2f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x3a, 0x01, 0x2a, 0x12, 0x6d, 0x0a, 0x0c, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x24, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x19, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x60, 0x0a, 0x08, 0x46, 0x69,
	0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4f, 0x70, 0x6d, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a,
	0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x56, 0x6f, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x56, 0x6f, 0x69, 0x63, 0x65,
	0x2f, 0x66, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x58, 0x0a, 0x04,
	0x46, 0x69, 0x6e, 0x64, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f,
	0x70, 0x6d, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x19, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x56, 0x6f, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16,
	0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x2f, 0x66,
	0x69, 0x6e, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x5f, 0x0a, 0x05, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x12,
	0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x56, 0x6f, 0x69,
	0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x6d, 0x56, 0x6f, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17,
	0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6d, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_opm_voice_service_gen_proto_goTypes = []interface{}{
	(*OpmVoice)(nil),                    // 0: service.OpmVoice
	(*OpmVoiceBatchDeleteRequest)(nil),  // 1: service.OpmVoiceBatchDeleteRequest
	(*OpmVoiceUpdateFieldsRequest)(nil), // 2: service.OpmVoiceUpdateFieldsRequest
	(*OpmVoiceFilter)(nil),              // 3: service.OpmVoiceFilter
	(*OpmVoiceListRequest)(nil),         // 4: service.OpmVoiceListRequest
	(*Response)(nil),                    // 5: service.Response
	(*OpmVoiceResponse)(nil),            // 6: service.OpmVoiceResponse
}
var file_opm_voice_service_gen_proto_depIdxs = []int32{
	0,  // 0: service.OpmVoiceService.Create:input_type -> service.OpmVoice
	0,  // 1: service.OpmVoiceService.Delete:input_type -> service.OpmVoice
	0,  // 2: service.OpmVoiceService.DeleteById:input_type -> service.OpmVoice
	1,  // 3: service.OpmVoiceService.DeleteByIds:input_type -> service.OpmVoiceBatchDeleteRequest
	0,  // 4: service.OpmVoiceService.Update:input_type -> service.OpmVoice
	0,  // 5: service.OpmVoiceService.UpdateAll:input_type -> service.OpmVoice
	2,  // 6: service.OpmVoiceService.UpdateFields:input_type -> service.OpmVoiceUpdateFieldsRequest
	3,  // 7: service.OpmVoiceService.FindById:input_type -> service.OpmVoiceFilter
	3,  // 8: service.OpmVoiceService.Find:input_type -> service.OpmVoiceFilter
	4,  // 9: service.OpmVoiceService.Lists:input_type -> service.OpmVoiceListRequest
	5,  // 10: service.OpmVoiceService.Create:output_type -> service.Response
	5,  // 11: service.OpmVoiceService.Delete:output_type -> service.Response
	5,  // 12: service.OpmVoiceService.DeleteById:output_type -> service.Response
	5,  // 13: service.OpmVoiceService.DeleteByIds:output_type -> service.Response
	5,  // 14: service.OpmVoiceService.Update:output_type -> service.Response
	5,  // 15: service.OpmVoiceService.UpdateAll:output_type -> service.Response
	5,  // 16: service.OpmVoiceService.UpdateFields:output_type -> service.Response
	6,  // 17: service.OpmVoiceService.FindById:output_type -> service.OpmVoiceResponse
	6,  // 18: service.OpmVoiceService.Find:output_type -> service.OpmVoiceResponse
	6,  // 19: service.OpmVoiceService.Lists:output_type -> service.OpmVoiceResponse
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_opm_voice_service_gen_proto_init() }
func file_opm_voice_service_gen_proto_init() {
	if File_opm_voice_service_gen_proto != nil {
		return
	}
	file_common_proto_init()
	file_opm_voice_model_gen_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_opm_voice_service_gen_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_opm_voice_service_gen_proto_goTypes,
		DependencyIndexes: file_opm_voice_service_gen_proto_depIdxs,
	}.Build()
	File_opm_voice_service_gen_proto = out.File
	file_opm_voice_service_gen_proto_rawDesc = nil
	file_opm_voice_service_gen_proto_goTypes = nil
	file_opm_voice_service_gen_proto_depIdxs = nil
}
