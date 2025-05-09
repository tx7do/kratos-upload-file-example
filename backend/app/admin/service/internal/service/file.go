package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/trans"

	"kratos-upload-file-example/app/admin/service/internal/data"

	adminV1 "kratos-upload-file-example/api/gen/go/admin/service/v1"
	fileV1 "kratos-upload-file-example/api/gen/go/file/service/v1"
)

type FileService struct {
	adminV1.FileServiceHTTPServer

	log *log.Helper

	mc *data.MinIOClient
}

func NewFileService(logger log.Logger, mc *data.MinIOClient) *FileService {
	l := log.NewHelper(log.With(logger, "module", "file/service/admin-service"))
	return &FileService{
		log: l,
		mc:  mc,
	}
}

func (s *FileService) OssUploadUrl(ctx context.Context, req *fileV1.OssUploadUrlRequest) (*fileV1.OssUploadUrlResponse, error) {
	return s.mc.OssUploadUrl(ctx, req)
}

func (s *FileService) PostUploadFile(ctx context.Context, req *fileV1.UploadFileRequest) (*fileV1.UploadFileResponse, error) {
	if req.File == nil {
		return nil, fileV1.ErrorUploadFailed("unknown file")
	}

	if req.BucketName == nil {
		req.BucketName = trans.Ptr(s.mc.ContentTypeToBucketName(req.GetMime()))
	}
	if req.ObjectName == nil {
		req.ObjectName = trans.Ptr(req.GetSourceFileName())
	}

	s.log.Infof("UPLOAD %d\n", len(req.File))

	downloadUrl, err := s.mc.UploadFile(ctx, req.GetBucketName(), req.GetObjectName(), req.GetFile())
	return &fileV1.UploadFileResponse{
		Url: downloadUrl,
	}, err
}

func (s *FileService) PutUploadFile(ctx context.Context, req *fileV1.UploadFileRequest) (*fileV1.UploadFileResponse, error) {
	if req.File == nil {
		return nil, fileV1.ErrorUploadFailed("unknown file")
	}

	if req.BucketName == nil {
		req.BucketName = trans.Ptr(s.mc.ContentTypeToBucketName(req.GetMime()))
	}
	if req.ObjectName == nil {
		req.ObjectName = trans.Ptr(req.GetSourceFileName())
	}

	downloadUrl, err := s.mc.UploadFile(ctx, req.GetBucketName(), req.GetObjectName(), req.GetFile())
	return &fileV1.UploadFileResponse{
		Url: downloadUrl,
	}, err
}
