# 变量定义
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=gfc
DOCKER_COMPOSE=docker-compose
GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
GOSUMDB=sum.golang.google.cn

# 服务列表
SERVICES := identity transaction consensus monitor

# 构建目标目录
BUILD_DIR=bin
PROTO_DIR=api/proto
PROTO_GO_DIR=api/proto

# 版本信息
VERSION=$(shell git describe --tags --always --dirty)
COMMIT=$(shell git rev-parse --short HEAD)
BUILD_TIME=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

# 编译标记
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.Commit=$(COMMIT) -X main.BuildTime=$(BUILD_TIME)"

.PHONY: all build clean test coverage deps proto docker-build docker-push install help env-dev env-prod

# 默认目标
all: deps proto build test

# 帮助信息
help:
	@echo "Global Fair Currency System Make Commands:"
	@echo "  make all              - 执行完整构建流程(依赖、协议、构建、测试)"
	@echo "  make build            - 构建所有服务"
	@echo "  make clean            - 清理构建文件"
	@echo "  make test             - 运行测试"
	@echo "  make coverage         - 生成测试覆盖率报告"
	@echo "  make deps             - 安装依赖"
	@echo "  make proto            - 生成协议文件"
	@echo "  make docker-build     - 构建Docker镜像"
	@echo "  make docker-push      - 推送Docker镜像"
	@echo "  make install          - 安装到系统"
	@echo "  make dev              - 启动开发环境"
	@echo "  make prod             - 启动生产环境"
	@echo "  make migrate          - 执行数据库迁移"
	@echo "  make rollback         - 回滚数据库迁移"
	@echo "  make env-dev          - 切换到开发环境"
	@echo "  make env-prod         - 切换到生产环境"

# 构建所有服务
build: $(SERVICES)

$(SERVICES):
	@echo "Building $@ service..."
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$@-service ./cmd/$@-service

# 清理构建文件
clean:
	@echo "Cleaning build files..."
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)
	rm -rf $(PROTO_GO_DIR)/*.pb.go
	find . -name "*.test" -delete
	find . -name "*.out" -delete

# 运行测试
test: 
	@echo "Running tests..."
	$(GOTEST) -v -race -cover ./...
	@echo "Running integration tests..."
	$(GOTEST) -v -tags=integration ./test/integration/...
	@echo "Running e2e tests..."
	$(GOTEST) -v -tags=e2e ./test/e2e/...

# 生成测试覆盖率报告
coverage:
	@echo "Generating test coverage report..."
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# 安装依赖
deps:
	@echo "Installing dependencies..."
	go env -w GOPROXY=$(GOPROXY)
	go env -w GOSUMDB=off
	go env -w GO111MODULE=on
	rm -f go.sum
	$(GOMOD) tidy
	$(GOMOD) verify
	@echo "Installing tools..."
	GOPROXY=$(GOPROXY) GOSUMDB=off go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOPROXY=$(GOPROXY) GOSUMDB=off go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	GOPROXY=$(GOPROXY) GOSUMDB=off go install github.com/golangci/golangint/cmd/golangci-lint@latest

# 生成协议文件
proto:
	@echo "Generating protocol buffers..."
	@mkdir -p api/proto
	protoc --go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		api/proto/*.proto

# 代码格式化
fmt:
	@echo "Formatting code..."
	$(GOCMD) fmt ./...
	gofmt -s -w .

# 代码检查
lint:
	@echo "Running linter..."
	golangci-lint run

# Docker相关命令
docker-build:
	@echo "Building Docker images..."
	@for service in $(SERVICES); do \
		docker build -t gfc-$$service:$(VERSION) \
			--build-arg SERVICE=$$service \
			--build-arg VERSION=$(VERSION) \
			-f deployments/docker/Dockerfile .; \
	done

docker-push:
	@echo "Pushing Docker images..."
	@for service in $(SERVICES); do \
		docker push gfc-$$service:$(VERSION); \
	done

# 安装到系统
install: build
	@echo "Installing binaries..."
	@for service in $(SERVICES); do \
		cp $(BUILD_DIR)/$$service-service /usr/local/bin/; \
	done

# 开发环境
dev:
	@echo "Starting development environment..."
	$(DOCKER_COMPOSE) -f deployments/docker/docker-compose.dev.yml up -d

# 生产环境
prod:
	@echo "Starting production environment..."
	$(DOCKER_COMPOSE) -f deployments/docker/docker-compose.prod.yml up -d

# 数据库迁移
migrate:
	@echo "Running database migrations..."
	migrate -path db/migrations -database "$(DB_URL)" up

rollback:
	@echo "Rolling back database migration..."
	migrate -path db/migrations -database "$(DB_URL)" down 1

# 版本信息
version:
	@echo "Version: $(VERSION)"
	@echo "Commit: $(COMMIT)"
	@echo "Build Time: $(BUILD_TIME)"

# 检查环境
check-env:
	@echo "Checking environment..."
	@command -v go >/dev/null 2>&1 || { echo "go is required but not installed"; exit 1; }
	@command -v docker >/dev/null 2>&1 || { echo "docker is required but not installed"; exit 1; }
	@command -v protoc >/dev/null 2>&1 || { echo "protoc is required but not installed"; exit 1; }
	@echo "Environment check passed"

# 生成文档
docs:
	@echo "Generating documentation..."
	swag init -g cmd/api/main.go -o api/swagger
	@echo "Documentation generated in api/swagger"

# 环境切换
env-dev:
	@echo "Switching to development environment..."
	@./scripts/init/init.sh development

env-prod:
	@echo "Switching to production environment..."
	@./scripts/init/init.sh production