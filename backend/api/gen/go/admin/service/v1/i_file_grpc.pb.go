// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: admin/service/v1/i_file.proto

package servicev1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	v1 "kratos-upload-file-example/api/gen/go/file/service/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	FileService_OssUploadUrl_FullMethodName   = "/admin.service.v1.FileService/OssUploadUrl"
	FileService_PostUploadFile_FullMethodName = "/admin.service.v1.FileService/PostUploadFile"
	FileService_PutUploadFile_FullMethodName  = "/admin.service.v1.FileService/PutUploadFile"
)

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 文件服务
type FileServiceClient interface {
	// 获取对象存储（OSS）上传用的预签名链接
	OssUploadUrl(ctx context.Context, in *v1.OssUploadUrlRequest, opts ...grpc.CallOption) (*v1.OssUploadUrlResponse, error)
	// POST方法上传文件
	PostUploadFile(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[v1.UploadFileRequest, v1.UploadFileResponse], error)
	// PUT方法上传文件
	PutUploadFile(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[v1.UploadFileRequest, v1.UploadFileResponse], error)
}

type fileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServiceClient(cc grpc.ClientConnInterface) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) OssUploadUrl(ctx context.Context, in *v1.OssUploadUrlRequest, opts ...grpc.CallOption) (*v1.OssUploadUrlResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.OssUploadUrlResponse)
	err := c.cc.Invoke(ctx, FileService_OssUploadUrl_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) PostUploadFile(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[v1.UploadFileRequest, v1.UploadFileResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &FileService_ServiceDesc.Streams[0], FileService_PostUploadFile_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[v1.UploadFileRequest, v1.UploadFileResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FileService_PostUploadFileClient = grpc.ClientStreamingClient[v1.UploadFileRequest, v1.UploadFileResponse]

func (c *fileServiceClient) PutUploadFile(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[v1.UploadFileRequest, v1.UploadFileResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &FileService_ServiceDesc.Streams[1], FileService_PutUploadFile_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[v1.UploadFileRequest, v1.UploadFileResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FileService_PutUploadFileClient = grpc.ClientStreamingClient[v1.UploadFileRequest, v1.UploadFileResponse]

// FileServiceServer is the server API for FileService service.
// All implementations must embed UnimplementedFileServiceServer
// for forward compatibility.
//
// 文件服务
type FileServiceServer interface {
	// 获取对象存储（OSS）上传用的预签名链接
	OssUploadUrl(context.Context, *v1.OssUploadUrlRequest) (*v1.OssUploadUrlResponse, error)
	// POST方法上传文件
	PostUploadFile(grpc.ClientStreamingServer[v1.UploadFileRequest, v1.UploadFileResponse]) error
	// PUT方法上传文件
	PutUploadFile(grpc.ClientStreamingServer[v1.UploadFileRequest, v1.UploadFileResponse]) error
	mustEmbedUnimplementedFileServiceServer()
}

// UnimplementedFileServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFileServiceServer struct{}

func (UnimplementedFileServiceServer) OssUploadUrl(context.Context, *v1.OssUploadUrlRequest) (*v1.OssUploadUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OssUploadUrl not implemented")
}
func (UnimplementedFileServiceServer) PostUploadFile(grpc.ClientStreamingServer[v1.UploadFileRequest, v1.UploadFileResponse]) error {
	return status.Errorf(codes.Unimplemented, "method PostUploadFile not implemented")
}
func (UnimplementedFileServiceServer) PutUploadFile(grpc.ClientStreamingServer[v1.UploadFileRequest, v1.UploadFileResponse]) error {
	return status.Errorf(codes.Unimplemented, "method PutUploadFile not implemented")
}
func (UnimplementedFileServiceServer) mustEmbedUnimplementedFileServiceServer() {}
func (UnimplementedFileServiceServer) testEmbeddedByValue()                     {}

// UnsafeFileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServiceServer will
// result in compilation errors.
type UnsafeFileServiceServer interface {
	mustEmbedUnimplementedFileServiceServer()
}

func RegisterFileServiceServer(s grpc.ServiceRegistrar, srv FileServiceServer) {
	// If the following call pancis, it indicates UnimplementedFileServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FileService_ServiceDesc, srv)
}

func _FileService_OssUploadUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.OssUploadUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).OssUploadUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileService_OssUploadUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).OssUploadUrl(ctx, req.(*v1.OssUploadUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_PostUploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileServiceServer).PostUploadFile(&grpc.GenericServerStream[v1.UploadFileRequest, v1.UploadFileResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FileService_PostUploadFileServer = grpc.ClientStreamingServer[v1.UploadFileRequest, v1.UploadFileResponse]

func _FileService_PutUploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileServiceServer).PutUploadFile(&grpc.GenericServerStream[v1.UploadFileRequest, v1.UploadFileResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type FileService_PutUploadFileServer = grpc.ClientStreamingServer[v1.UploadFileRequest, v1.UploadFileResponse]

// FileService_ServiceDesc is the grpc.ServiceDesc for FileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.service.v1.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OssUploadUrl",
			Handler:    _FileService_OssUploadUrl_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PostUploadFile",
			Handler:       _FileService_PostUploadFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "PutUploadFile",
			Handler:       _FileService_PutUploadFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "admin/service/v1/i_file.proto",
}
