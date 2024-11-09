#!/bin/bash

# 安装 Go 环境
install_golang() {
    echo "Installing Go..."
    
    # 使用 Go 1.20 版本和国内镜像
    GO_VERSION="1.20.11"
    wget https://golang.google.cn/dl/go${GO_VERSION}.linux-amd64.tar.gz || \
    wget https://studygolang.com/dl/golang/go${GO_VERSION}.linux-amd64.tar.gz || \
    wget https://gomirrors.org/dl/go/go${GO_VERSION}.linux-amd64.tar.gz
    
    if [ ! -f "go${GO_VERSION}.linux-amd64.tar.gz" ]; then
        echo "Failed to download Go"
        exit 1
    fi

    rm -rf /usr/local/go
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz
    rm go${GO_VERSION}.linux-amd64.tar.gz

    # 设置环境变量
    echo 'export PATH=$PATH:/usr/local/go/bin' >> /etc/profile
    echo 'export GOPATH=$HOME/go' >> /etc/profile
    echo 'export GOBIN=$GOPATH/bin' >> /etc/profile
    echo 'export GO111MODULE=on' >> /etc/profile
    echo 'export GOPROXY=https://goproxy.cn,direct' >> /etc/profile
    echo 'export GOSUMDB=off' >> /etc/profile
    
    source /etc/profile
    
    # 验证安装
    go version
}

# 安装 protoc
install_protoc() {
    echo "Installing protoc..."
    
    # 使用国内镜像下载 protoc
    PROTOC_VERSION="3.19.4"
    PROTOC_MIRROR="https://gitee.com/mirrors/protobuf/releases/download/v${PROTOC_VERSION}"
    PROTOC_ZIP="protoc-${PROTOC_VERSION}-linux-x86_64.zip"
    
    wget ${PROTOC_MIRROR}/${PROTOC_ZIP} || \
    wget https://ghproxy.com/https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/${PROTOC_ZIP}
    
    if [ ! -f "${PROTOC_ZIP}" ]; then
        echo "Failed to download protoc"
        exit 1
    fi

    # 解压到 /usr/local
    unzip -o ${PROTOC_ZIP} -d /usr/local
    rm ${PROTOC_ZIP}
    
    # 设置权限
    chmod 755 /usr/local/bin/protoc
    chmod -R 755 /usr/local/include/google

    # 安装 Go protoc 插件
    echo "Installing protoc-gen-go..."
    GOBIN=/usr/local/bin go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
    GOBIN=/usr/local/bin go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0

    # 验证安装
    echo "Verifying installations..."
    protoc --version
    which protoc
    which protoc-gen-go
    which protoc-gen-go-grpc

    # 添加到环境变量
    echo 'export PATH=$PATH:/usr/local/bin' >> /etc/profile
    source /etc/profile
}

# 设置环境
set_environment() {
    local env=${1:-development}
    echo "Setting up ${env} environment..."
    
    # 加载对应的环境变量文件
    if [ -f ".env.${env}" ]; then
        export $(cat .env.${env} | grep -v '^#' | xargs)
        cp .env.${env} .env
    else
        echo "Environment file .env.${env} not found!"
        exit 1
    fi
}

# 初始化环境
init_environment() {
    # 创建必要的目录
    mkdir -p /var/log/gfc
    mkdir -p /var/lib/gfc
    mkdir -p /etc/gfc/ssl

    # 设置权限
    chown -R www:www /var/log/gfc
    chown -R www:www /var/lib/gfc
}

# 主函数
main() {
    local env=${1:-development}
    
    # 安装必要的环境和工具
    install_golang
    install_protoc
    
    # 设置环境和初始化
    set_environment $env
    init_environment
    
    echo "Initialization completed successfully!"
}

main "$@"