// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: file/service/v1/file_error.proto

package servicev1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FileErrorReason int32

const (
	// common error
	FileErrorReason_METHOD_NOT_ALLOWED FileErrorReason = 0 // 405
	// file upload/download errors
	FileErrorReason_FILE_NOT_FOUND        FileErrorReason = 1 // 404
	FileErrorReason_FILE_TOO_LARGE        FileErrorReason = 2 // 413
	FileErrorReason_UNSUPPORTED_FILE_TYPE FileErrorReason = 3 // 415
	FileErrorReason_UPLOAD_FAILED         FileErrorReason = 4 // 500
	FileErrorReason_DOWNLOAD_FAILED       FileErrorReason = 5 // 500
)

// Enum value maps for FileErrorReason.
var (
	FileErrorReason_name = map[int32]string{
		0: "METHOD_NOT_ALLOWED",
		1: "FILE_NOT_FOUND",
		2: "FILE_TOO_LARGE",
		3: "UNSUPPORTED_FILE_TYPE",
		4: "UPLOAD_FAILED",
		5: "DOWNLOAD_FAILED",
	}
	FileErrorReason_value = map[string]int32{
		"METHOD_NOT_ALLOWED":    0,
		"FILE_NOT_FOUND":        1,
		"FILE_TOO_LARGE":        2,
		"UNSUPPORTED_FILE_TYPE": 3,
		"UPLOAD_FAILED":         4,
		"DOWNLOAD_FAILED":       5,
	}
)

func (x FileErrorReason) Enum() *FileErrorReason {
	p := new(FileErrorReason)
	*p = x
	return p
}

func (x FileErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FileErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_file_service_v1_file_error_proto_enumTypes[0].Descriptor()
}

func (FileErrorReason) Type() protoreflect.EnumType {
	return &file_file_service_v1_file_error_proto_enumTypes[0]
}

func (x FileErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FileErrorReason.Descriptor instead.
func (FileErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_file_service_v1_file_error_proto_rawDescGZIP(), []int{0}
}

var File_file_service_v1_file_error_proto protoreflect.FileDescriptor

var file_file_service_v1_file_error_proto_rawDesc = string([]byte{
	0x0a, 0x20, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xbe, 0x01, 0x0a, 0x0f, 0x46, 0x69, 0x6c,
	0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x12,
	0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57,
	0x45, 0x44, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0x95, 0x03, 0x12, 0x18, 0x0a, 0x0e, 0x46, 0x49,
	0x4c, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x1a, 0x04,
	0xa8, 0x45, 0x94, 0x03, 0x12, 0x18, 0x0a, 0x0e, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x54, 0x4f, 0x4f,
	0x5f, 0x4c, 0x41, 0x52, 0x47, 0x45, 0x10, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0x9d, 0x03, 0x12, 0x1f,
	0x0a, 0x15, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x46, 0x49,
	0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0x9f, 0x03, 0x12,
	0x17, 0x0a, 0x0d, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44,
	0x10, 0x04, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12, 0x19, 0x0a, 0x0f, 0x44, 0x4f, 0x57, 0x4e,
	0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x05, 0x1a, 0x04, 0xa8,
	0x45, 0xf4, 0x03, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0xc4, 0x01, 0x0a, 0x13, 0x63, 0x6f,
	0x6d, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x42, 0x0e, 0x46, 0x69, 0x6c, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x3f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x2d, 0x66, 0x69, 0x6c, 0x65, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x46, 0x53, 0x58, 0xaa, 0x02, 0x0f, 0x46, 0x69, 0x6c,
	0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x46,
	0x69, 0x6c, 0x65, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x1b, 0x46, 0x69, 0x6c, 0x65, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x46,
	0x69, 0x6c, 0x65, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_file_service_v1_file_error_proto_rawDescOnce sync.Once
	file_file_service_v1_file_error_proto_rawDescData []byte
)

func file_file_service_v1_file_error_proto_rawDescGZIP() []byte {
	file_file_service_v1_file_error_proto_rawDescOnce.Do(func() {
		file_file_service_v1_file_error_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_file_service_v1_file_error_proto_rawDesc), len(file_file_service_v1_file_error_proto_rawDesc)))
	})
	return file_file_service_v1_file_error_proto_rawDescData
}

var file_file_service_v1_file_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_file_service_v1_file_error_proto_goTypes = []any{
	(FileErrorReason)(0), // 0: file.service.v1.FileErrorReason
}
var file_file_service_v1_file_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_file_service_v1_file_error_proto_init() }
func file_file_service_v1_file_error_proto_init() {
	if File_file_service_v1_file_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_file_service_v1_file_error_proto_rawDesc), len(file_file_service_v1_file_error_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_file_service_v1_file_error_proto_goTypes,
		DependencyIndexes: file_file_service_v1_file_error_proto_depIdxs,
		EnumInfos:         file_file_service_v1_file_error_proto_enumTypes,
	}.Build()
	File_file_service_v1_file_error_proto = out.File
	file_file_service_v1_file_error_proto_goTypes = nil
	file_file_service_v1_file_error_proto_depIdxs = nil
}
