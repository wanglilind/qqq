#!/bin/bash

# 测试环境变量
export GO_ENV=testing
export TEST_DB_HOST=localhost
export TEST_DB_PORT=5432

# 颜色输出
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m'

# 运行单元测试
run_unit_tests() {
    echo "Running unit tests..."
    go test -v -race -cover ./... -tags=unit
    if [ $? -ne 0 ]; then
        echo -e "${RED}Unit tests failed${NC}"
        exit 1
    fi
    echo -e "${GREEN}Unit tests passed${NC}"
}

# 运行集成测试
run_integration_tests() {
    echo "Running integration tests..."
    
    # 启动测试数据库
    docker-compose -f deployments/docker/docker-compose.test.yml up -d

    # 等待数据库就绪
    sleep 5

    # 运行测试
    go test -v -race -cover ./... -tags=integration
    TEST_EXIT_CODE=$?

    # 清理测试环境
    docker-compose -f deployments/docker/docker-compose.test.yml down

    if [ $TEST_EXIT_CODE -ne 0 ]; then
        echo -e "${RED}Integration tests failed${NC}"
        exit 1
    fi
    echo -e "${GREEN}Integration tests passed${NC}"
}

# 运行端到端测试
run_e2e_tests() {
    echo "Running end-to-end tests..."
    
    # 启动测试环境
    docker-compose -f deployments/docker/docker-compose.test.yml up -d

    # 等待服务就绪
    sleep 10

    # 运行E2E测试
    go test -v ./test/e2e/...
    TEST_EXIT_CODE=$?

    # 清理测试环境
    docker-compose -f deployments/docker/docker-compose.test.yml down

    if [ $TEST_EXIT_CODE -ne 0 ]; then
        echo -e "${RED}E2E tests failed${NC}"
        exit 1
    fi
    echo -e "${GREEN}E2E tests passed${NC}"
}

# 生成测试覆盖率报告
generate_coverage_report() {
    echo "Generating coverage report..."
    go test -coverprofile=coverage.out ./...
    go tool cover -html=coverage.out -o coverage.html
    echo "Coverage report generated: coverage.html"
}

# 主测试流程
main() {
    # 检查Go环境
    if ! command -v go &> /dev/null; then
        echo -e "${RED}Go is not installed${NC}"
        exit 1
    }

    # 安装依赖
    go mod download

    # 运行所有测试
    run_unit_tests
    run_integration_tests
    run_e2e_tests
    generate_coverage_report

    echo -e "${GREEN}All tests passed successfully!${NC}"
}

main 