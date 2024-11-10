# 构建阶段
FROM golang:1.23-alpine AS builder

# 设置工作目录
WORKDIR /app/github.com/wanglilind/qqq

# 安装构建依赖
RUN apk add --no-cache git make

# 设置 Go 环境变量
ENV GO111MODULE=on
ENV GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
ENV GOSUMDB=off
ENV GOFLAGS=-mod=mod

# 复制整个项目
COPY . .

# 下载依赖（离线模式）
RUN go mod download -x

# 接收构建参数
ARG SERVICE
ARG VERSION

# 构建服务
RUN make build SERVICE=$SERVICE VERSION=$VERSION

# 运行阶段
FROM alpine:3.14

# 安装运行时依赖
RUN apk add --no-cache ca-certificates tzdata

# 设置时区
ENV TZ=Asia/Shanghai

# 复制构建产物
COPY --from=builder /app/./bin/${SERVICE}-service /usr/local/bin/

# 设置入口点
ENTRYPOINT ["/usr/local/bin/${SERVICE}-service"] 