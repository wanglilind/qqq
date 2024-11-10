# 项目变量
PROJECT := ./
VERSION ?= $(shell git describe --tags --always --dirty)
COMMIT := $(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date +%Y-%m-%d_%H:%M:%S)

# Go 相关变量
GO := go
GOFLAGS := -v
LDFLAGS := -X main.Version=${VERSION} -X main.Commit=${COMMIT} -X main.BuildTime=${BUILD_TIME}

# Docker 相关变量
DOCKER_REGISTRY := registry.example.com/wanglilind/qqq
DOCKER_TAG ?= latest

# Proto 相关变量
PROTO_DIR := api/proto
PROTO_GO_OUT := .
PROTO_GO_OPT := paths=source_relative
PROTO_GRPC_OPT := paths=source_relative,require_unimplemented_servers=false

# 服务列表
SERVICES := identity transaction consensus monitor

.PHONY: all
all: deps proto build test

# 依赖管理
.PHONY: deps
deps:
	@echo "Installing dependencies..."
	@echo "Cleaning go mod cache..."
	$(GO) clean -modcache
	@echo "Downloading dependencies..."
	$(GO) mod download || (sleep 2 && $(GO) mod download)
	$(GO) mod tidy
	$(GO) mod verify

# 生成 Proto 文件
.PHONY: proto
proto:
	@echo "Generating protocol buffers..."
	@mkdir -p $(PROTO_GO_OUT)
	protoc --proto_path=$(PROTO_DIR) \
		--go_out=$(PROTO_GO_OUT) \
		--go_opt=$(PROTO_GO_OPT) \
		--go-grpc_out=$(PROTO_GO_OUT) \
		--go-grpc_opt=$(PROTO_GRPC_OPT) \
		$(PROTO_DIR)/*.proto

# 构建服务
.PHONY: build
build: $(SERVICES)

.PHONY: $(SERVICES)
$(SERVICES): deps proto
	@echo "Building $@ service..."
	$(GO) build -ldflags "$(LDFLAGS)" -o bin/$@-service ./cmd/$@-service

# 清理
.PHONY: clean
clean:
	@echo "Cleaning..."
	rm -rf bin/
	rm -f $(PROTO_DIR)/*.pb.go

# 测试
.PHONY: test
test:
	@echo "Running tests..."
	$(GO) test -v -race -cover ./...

# Docker 构建
.PHONY: docker
docker: $(addprefix docker-,$(SERVICES))

.PHONY: docker-%
docker-%:
	docker build \
		--build-arg SERVICE=$* \
		--build-arg VERSION=$(VERSION) \
		-t $(DOCKER_REGISTRY)/$*:$(DOCKER_TAG) \
		-f deployments/docker/Dockerfile .

# 开发环境
.PHONY: dev
dev:
	docker-compose -f deployments/docker/docker-compose.dev.yml up --build

# 生产环境
.PHONY: prod
prod:
	docker-compose -f deployments/docker/docker-compose.yml up -d

# 帮助信息
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  all        - Build everything"
	@echo "  deps       - Install dependencies"
	@echo "  proto      - Generate protobuf files"
	@echo "  build      - Build all services"
	@echo "  clean      - Clean build artifacts"
	@echo "  test       - Run tests"
	@echo "  docker     - Build Docker images"
	@echo "  dev        - Start development environment"
	@echo "  prod       - Start production environment"