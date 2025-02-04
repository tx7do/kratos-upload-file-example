package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/transport/http"

	swaggerUI "github.com/tx7do/kratos-swagger-ui"

	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
	"github.com/tx7do/kratos-bootstrap/rpc"

	"kratos-upload-file-example/app/admin/service/cmd/server/assets"
	"kratos-upload-file-example/app/admin/service/internal/service"

	adminV1 "kratos-upload-file-example/api/gen/go/admin/service/v1"
)

// NewMiddleware 创建中间件
func newRestMiddleware(logger log.Logger) []middleware.Middleware {
	var ms []middleware.Middleware
	ms = append(ms, logging.Server(logger))
	return ms
}

// NewRESTServer new an HTTP server.
func NewRESTServer(
	cfg *conf.Bootstrap,
	logger log.Logger,

	fileSvc *service.FileService,
	ueditorSvc *service.UEditorService,
) *http.Server {
	srv := rpc.CreateRestServer(
		cfg,
		newRestMiddleware(logger)...,
	)

	adminV1.RegisterFileServiceHTTPServer(srv, fileSvc)
	adminV1.RegisterUEditorServiceHTTPServer(srv, ueditorSvc)

	registerFileUploadHandler(srv, fileSvc)
	registerUEditorUploadHandler(srv, ueditorSvc)

	if cfg.GetServer().GetRest().GetEnableSwagger() {
		swaggerUI.RegisterSwaggerUIServerWithOption(
			srv,
			swaggerUI.WithTitle("Kratos文件上传实例"),
			swaggerUI.WithMemoryData(assets.OpenApiData, "yaml"),
		)
	}

	return srv
}
