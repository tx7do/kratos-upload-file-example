// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: admin/service/v1/admin_error.proto

package servicev1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AdminErrorReason int32

const (
	// common error
	AdminErrorReason_METHOD_NOT_ALLOWED    AdminErrorReason = 0 // 405
	AdminErrorReason_REQUEST_TIMEOUT       AdminErrorReason = 1 // 408
	AdminErrorReason_INTERNAL_SERVER_ERROR AdminErrorReason = 2 // 500
	AdminErrorReason_NOT_IMPLEMENTED       AdminErrorReason = 3 // 501
	AdminErrorReason_NETWORK_ERROR         AdminErrorReason = 4 // 502
	AdminErrorReason_SERVICE_UNAVAILABLE   AdminErrorReason = 5 // 503
	AdminErrorReason_NETWORK_TIMEOUT       AdminErrorReason = 6 // 504
	AdminErrorReason_REQUEST_NOT_SUPPORT   AdminErrorReason = 7 // 505
	// 400
	AdminErrorReason_BAD_REQUEST        AdminErrorReason = 100 // 400
	AdminErrorReason_INVALID_GRANT_TYPE AdminErrorReason = 101 // 400
	AdminErrorReason_INVALID_USERID     AdminErrorReason = 102 // 用户ID无效
	AdminErrorReason_INVALID_TOKEN      AdminErrorReason = 103 // token无效
	AdminErrorReason_INVALID_PASSWORD   AdminErrorReason = 104 // 密码无效
	// 404
	AdminErrorReason_RESOURCE_NOT_FOUND AdminErrorReason = 200 // 404
	AdminErrorReason_USER_NOT_FOUND     AdminErrorReason = 201 // 用户不存在
	// 401
	AdminErrorReason_NOT_LOGGED_IN           AdminErrorReason = 300
	AdminErrorReason_USER_FREEZE             AdminErrorReason = 301 // 用户被冻结
	AdminErrorReason_INCORRECT_PASSWORD      AdminErrorReason = 302 // 密码错误
	AdminErrorReason_INCORRECT_APP_SECRET    AdminErrorReason = 303 // 密钥错误
	AdminErrorReason_INCORRECT_ACCESS_TOKEN  AdminErrorReason = 304 // 访问令牌错误
	AdminErrorReason_INCORRECT_REFRESH_TOKEN AdminErrorReason = 305 // 刷新令牌错误
	AdminErrorReason_TOKEN_EXPIRED           AdminErrorReason = 306 // token过期
	AdminErrorReason_TOKEN_NOT_EXIST         AdminErrorReason = 307 // token不存在
	// 403
	AdminErrorReason_ACCESS_FORBIDDEN AdminErrorReason = 400
)

// Enum value maps for AdminErrorReason.
var (
	AdminErrorReason_name = map[int32]string{
		0:   "METHOD_NOT_ALLOWED",
		1:   "REQUEST_TIMEOUT",
		2:   "INTERNAL_SERVER_ERROR",
		3:   "NOT_IMPLEMENTED",
		4:   "NETWORK_ERROR",
		5:   "SERVICE_UNAVAILABLE",
		6:   "NETWORK_TIMEOUT",
		7:   "REQUEST_NOT_SUPPORT",
		100: "BAD_REQUEST",
		101: "INVALID_GRANT_TYPE",
		102: "INVALID_USERID",
		103: "INVALID_TOKEN",
		104: "INVALID_PASSWORD",
		200: "RESOURCE_NOT_FOUND",
		201: "USER_NOT_FOUND",
		300: "NOT_LOGGED_IN",
		301: "USER_FREEZE",
		302: "INCORRECT_PASSWORD",
		303: "INCORRECT_APP_SECRET",
		304: "INCORRECT_ACCESS_TOKEN",
		305: "INCORRECT_REFRESH_TOKEN",
		306: "TOKEN_EXPIRED",
		307: "TOKEN_NOT_EXIST",
		400: "ACCESS_FORBIDDEN",
	}
	AdminErrorReason_value = map[string]int32{
		"METHOD_NOT_ALLOWED":      0,
		"REQUEST_TIMEOUT":         1,
		"INTERNAL_SERVER_ERROR":   2,
		"NOT_IMPLEMENTED":         3,
		"NETWORK_ERROR":           4,
		"SERVICE_UNAVAILABLE":     5,
		"NETWORK_TIMEOUT":         6,
		"REQUEST_NOT_SUPPORT":     7,
		"BAD_REQUEST":             100,
		"INVALID_GRANT_TYPE":      101,
		"INVALID_USERID":          102,
		"INVALID_TOKEN":           103,
		"INVALID_PASSWORD":        104,
		"RESOURCE_NOT_FOUND":      200,
		"USER_NOT_FOUND":          201,
		"NOT_LOGGED_IN":           300,
		"USER_FREEZE":             301,
		"INCORRECT_PASSWORD":      302,
		"INCORRECT_APP_SECRET":    303,
		"INCORRECT_ACCESS_TOKEN":  304,
		"INCORRECT_REFRESH_TOKEN": 305,
		"TOKEN_EXPIRED":           306,
		"TOKEN_NOT_EXIST":         307,
		"ACCESS_FORBIDDEN":        400,
	}
)

func (x AdminErrorReason) Enum() *AdminErrorReason {
	p := new(AdminErrorReason)
	*p = x
	return p
}

func (x AdminErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AdminErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_admin_service_v1_admin_error_proto_enumTypes[0].Descriptor()
}

func (AdminErrorReason) Type() protoreflect.EnumType {
	return &file_admin_service_v1_admin_error_proto_enumTypes[0]
}

func (x AdminErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AdminErrorReason.Descriptor instead.
func (AdminErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_admin_service_v1_admin_error_proto_rawDescGZIP(), []int{0}
}

var File_admin_service_v1_admin_error_proto protoreflect.FileDescriptor

var file_admin_service_v1_admin_error_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xc9, 0x05, 0x0a, 0x10,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x12, 0x1c, 0x0a, 0x12, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41,
	0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0x95, 0x03, 0x12, 0x19,
	0x0a, 0x0f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55,
	0x54, 0x10, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0x98, 0x03, 0x12, 0x1f, 0x0a, 0x15, 0x49, 0x4e, 0x54,
	0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12, 0x19, 0x0a, 0x0f, 0x4e, 0x4f,
	0x54, 0x5f, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x45, 0x44, 0x10, 0x03, 0x1a,
	0x04, 0xa8, 0x45, 0xf5, 0x03, 0x12, 0x17, 0x0a, 0x0d, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x1a, 0x04, 0xa8, 0x45, 0xf6, 0x03, 0x12, 0x1d,
	0x0a, 0x13, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x41, 0x56, 0x41, 0x49,
	0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x05, 0x1a, 0x04, 0xa8, 0x45, 0xf7, 0x03, 0x12, 0x19, 0x0a,
	0x0f, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54,
	0x10, 0x06, 0x1a, 0x04, 0xa8, 0x45, 0xf8, 0x03, 0x12, 0x1d, 0x0a, 0x13, 0x52, 0x45, 0x51, 0x55,
	0x45, 0x53, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x10,
	0x07, 0x1a, 0x04, 0xa8, 0x45, 0xf9, 0x03, 0x12, 0x15, 0x0a, 0x0b, 0x42, 0x41, 0x44, 0x5f, 0x52,
	0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x64, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x1c,
	0x0a, 0x12, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x47, 0x52, 0x41, 0x4e, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x10, 0x65, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x18, 0x0a, 0x0e,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x49, 0x44, 0x10, 0x66,
	0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x17, 0x0a, 0x0d, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x67, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12,
	0x1a, 0x0a, 0x10, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x57,
	0x4f, 0x52, 0x44, 0x10, 0x68, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x1d, 0x0a, 0x12, 0x52,
	0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e,
	0x44, 0x10, 0xc8, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x19, 0x0a, 0x0e, 0x55, 0x53,
	0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0xc9, 0x01, 0x1a,
	0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x18, 0x0a, 0x0d, 0x4e, 0x4f, 0x54, 0x5f, 0x4c, 0x4f, 0x47,
	0x47, 0x45, 0x44, 0x5f, 0x49, 0x4e, 0x10, 0xac, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03, 0x12,
	0x16, 0x0a, 0x0b, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x46, 0x52, 0x45, 0x45, 0x5a, 0x45, 0x10, 0xad,
	0x02, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03, 0x12, 0x1d, 0x0a, 0x12, 0x49, 0x4e, 0x43, 0x4f, 0x52,
	0x52, 0x45, 0x43, 0x54, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x10, 0xae, 0x02,
	0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03, 0x12, 0x1f, 0x0a, 0x14, 0x49, 0x4e, 0x43, 0x4f, 0x52, 0x52,
	0x45, 0x43, 0x54, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54, 0x10, 0xaf,
	0x02, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03, 0x12, 0x21, 0x0a, 0x16, 0x49, 0x4e, 0x43, 0x4f, 0x52,
	0x52, 0x45, 0x43, 0x54, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x4f, 0x4b, 0x45,
	0x4e, 0x10, 0xb0, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03, 0x12, 0x22, 0x0a, 0x17, 0x49, 0x4e,
	0x43, 0x4f, 0x52, 0x52, 0x45, 0x43, 0x54, 0x5f, 0x52, 0x45, 0x46, 0x52, 0x45, 0x53, 0x48, 0x5f,
	0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0xb1, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03, 0x12, 0x18,
	0x0a, 0x0d, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10,
	0xb2, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03, 0x12, 0x1a, 0x0a, 0x0f, 0x54, 0x4f, 0x4b, 0x45,
	0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10, 0xb3, 0x02, 0x1a, 0x04,
	0xa8, 0x45, 0x91, 0x03, 0x12, 0x1b, 0x0a, 0x10, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x46,
	0x4f, 0x52, 0x42, 0x49, 0x44, 0x44, 0x45, 0x4e, 0x10, 0x90, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0x93,
	0x03, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0xcb, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x42, 0x0f, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x40, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x2d, 0x66, 0x69, 0x6c, 0x65, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x53, 0x58, 0xaa, 0x02, 0x10, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x10, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x1c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x12, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_service_v1_admin_error_proto_rawDescOnce sync.Once
	file_admin_service_v1_admin_error_proto_rawDescData = file_admin_service_v1_admin_error_proto_rawDesc
)

func file_admin_service_v1_admin_error_proto_rawDescGZIP() []byte {
	file_admin_service_v1_admin_error_proto_rawDescOnce.Do(func() {
		file_admin_service_v1_admin_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_service_v1_admin_error_proto_rawDescData)
	})
	return file_admin_service_v1_admin_error_proto_rawDescData
}

var file_admin_service_v1_admin_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_admin_service_v1_admin_error_proto_goTypes = []any{
	(AdminErrorReason)(0), // 0: admin.service.v1.AdminErrorReason
}
var file_admin_service_v1_admin_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_admin_service_v1_admin_error_proto_init() }
func file_admin_service_v1_admin_error_proto_init() {
	if File_admin_service_v1_admin_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_admin_service_v1_admin_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_admin_service_v1_admin_error_proto_goTypes,
		DependencyIndexes: file_admin_service_v1_admin_error_proto_depIdxs,
		EnumInfos:         file_admin_service_v1_admin_error_proto_enumTypes,
	}.Build()
	File_admin_service_v1_admin_error_proto = out.File
	file_admin_service_v1_admin_error_proto_rawDesc = nil
	file_admin_service_v1_admin_error_proto_goTypes = nil
	file_admin_service_v1_admin_error_proto_depIdxs = nil
}
