// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: admin/service/v1/admin_doc.proto

package servicev1

import (
	_ "github.com/google/gnostic/openapiv3"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_admin_service_v1_admin_doc_proto protoreflect.FileDescriptor

var file_admin_service_v1_admin_doc_proto_rawDesc = string([]byte{
	0x0a, 0x20, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x64, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x1a, 0x24, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0xbd, 0x05, 0xba, 0x47, 0xf0,
	0x03, 0x12, 0x39, 0x0a, 0x18, 0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0xe6, 0x96, 0x87, 0xe4, 0xbb,
	0xb6, 0xe4, 0xb8, 0x8a, 0xe4, 0xbc, 0xa0, 0xe5, 0xae, 0x9e, 0xe4, 0xbe, 0x8b, 0x12, 0x18, 0x4b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0xe6, 0x96, 0x87, 0xe4, 0xbb, 0xb6, 0xe4, 0xb8, 0x8a, 0xe4, 0xbc,
	0xa0, 0xe5, 0xae, 0x9e, 0xe4, 0xbe, 0x8b, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x96, 0x03, 0x0a,
	0xd4, 0x01, 0x0a, 0xd1, 0x01, 0x0a, 0x0c, 0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0xc0, 0x01, 0x0a, 0xbd, 0x01, 0xca, 0x01, 0x06, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0xfa, 0x01, 0x9b, 0x01, 0x0a, 0x27, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1f,
	0x0a, 0x1d, 0xca, 0x01, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x92, 0x02, 0x09, 0xe9, 0x94,
	0x99, 0xe8, 0xaf, 0xaf, 0xe7, 0xa0, 0x81, 0x9a, 0x02, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x0a,
	0x25, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x18, 0xca, 0x01,
	0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x92, 0x02, 0x0c, 0xe9, 0x94, 0x99, 0xe8, 0xaf, 0xaf,
	0xe6, 0xb6, 0x88, 0xe6, 0x81, 0xaf, 0x0a, 0x24, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x18, 0xca, 0x01, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x92, 0x02, 0x0c,
	0xe9, 0x94, 0x99, 0xe8, 0xaf, 0xaf, 0xe5, 0x8e, 0x9f, 0xe5, 0x9b, 0xa0, 0x0a, 0x23, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x15, 0xca, 0x01, 0x06, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x92, 0x02, 0x09, 0xe5, 0x85, 0x83, 0xe6, 0x95, 0xb0, 0xe6, 0x8d,
	0xae, 0x92, 0x02, 0x12, 0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0xe9, 0x94, 0x99, 0xe8, 0xaf, 0xaf,
	0xe8, 0xbf, 0x94, 0xe5, 0x9b, 0x9e, 0x12, 0x67, 0x0a, 0x65, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x12, 0x5a, 0x0a, 0x58, 0x0a, 0x17, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x20, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x20, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x1a, 0x3d, 0x0a, 0x3b, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x25, 0x12, 0x23, 0x0a, 0x21, 0x23, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x73, 0x2f, 0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a,
	0x54, 0x0a, 0x52, 0x0a, 0x14, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x38, 0x0a, 0x06, 0x6f,
	0x61, 0x75, 0x74, 0x68, 0x32, 0x3a, 0x2e, 0x12, 0x2c, 0x12, 0x0f, 0x2f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x1a, 0x17, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x00, 0x32, 0x1a, 0x0a, 0x18, 0x0a, 0x14, 0x4f, 0x41, 0x75, 0x74, 0x68,
	0x32, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12,
	0x00, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x44, 0x6f,
	0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73,
	0x2d, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2d, 0x66, 0x69, 0x6c, 0x65, 0x2d, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31,
	0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x53, 0x58,
	0xaa, 0x02, 0x10, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x3a, 0x3a, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var file_admin_service_v1_admin_doc_proto_goTypes = []any{}
var file_admin_service_v1_admin_doc_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_admin_service_v1_admin_doc_proto_init() }
func file_admin_service_v1_admin_doc_proto_init() {
	if File_admin_service_v1_admin_doc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_admin_service_v1_admin_doc_proto_rawDesc), len(file_admin_service_v1_admin_doc_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_admin_service_v1_admin_doc_proto_goTypes,
		DependencyIndexes: file_admin_service_v1_admin_doc_proto_depIdxs,
	}.Build()
	File_admin_service_v1_admin_doc_proto = out.File
	file_admin_service_v1_admin_doc_proto_goTypes = nil
	file_admin_service_v1_admin_doc_proto_depIdxs = nil
}
