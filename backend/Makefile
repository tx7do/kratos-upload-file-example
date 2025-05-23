# Not Support Windows

.PHONY: help wire gen ent build api openapi init all

ifeq ($(OS),Windows_NT)
    IS_WINDOWS:=1
endif

CURRENT_DIR := $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST)))))
ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

SRCS_MK := $(foreach dir, app, $(wildcard $(dir)/*/*/Makefile))

# initialize develop environment
init: plugin cli

# install protoc plugin
plugin:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	@go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	@go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	@go install github.com/envoyproxy/protoc-gen-validate@latest

# install cli tools
cli:
	@go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	@go install github.com/google/gnostic@latest
	@go install github.com/bufbuild/buf/cmd/buf@latest
	@go install entgo.io/ent/cmd/ent@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# generate wire code
wire:
	$(foreach dir, $(dir $(realpath $(SRCS_MK))),\
      cd $(dir);\
      make wire;\
    )

# generate code by go:generate
gen:
	$(foreach dir, $(dir $(realpath $(SRCS_MK))),\
      cd $(dir);\
      make gen;\
    )

# generate ent code
ent:
	$(foreach dir, $(dir $(realpath $(SRCS_MK))),\
      cd $(dir);\
      make ent;\
    )

# generate protobuf api go code
api:
	@cd api && \
	buf generate

# generate OpenAPI v3 docs.
openapi:
	@cd api && \
	buf generate --template buf.admin.openapi.gen.yaml

# generate typescript.
ts:
	@cd api && \
	buf generate --template buf.vue.typescript.gen.yaml && \
	buf generate --template buf.react.typescript.gen.yaml

# build all service applications
build:
	$(foreach dir, $(dir $(realpath $(SRCS_MK))),\
      cd $(dir);\
      make build;\
    )

# generate & build all service applications
all:
	$(foreach dir, $(dir $(realpath $(SRCS_MK))),\
      cd $(dir);\
      make app;\
    )

# show help
help:
	@echo ""
	@echo "Usage:"
	@echo " make [target]"
	@echo ""
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
