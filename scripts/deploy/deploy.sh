#!/bin/bash

# 部署环境变量
ENVIRONMENT=${1:-staging}  # 默认部署到staging环境
NAMESPACE="gfc-${ENVIRONMENT}"
SERVICES=("identity" "transaction" "consensus" "monitor")
REGISTRY="registry.example.com"

# 检查必要的工具
check_requirements() {
    command -v kubectl >/dev/null 2>&1 || { echo "kubectl is required but not installed."; exit 1; }
    command -v docker >/dev/null 2>&1 || { echo "docker is required but not installed."; exit 1; }
}

# 构建和推送Docker镜像
build_and_push() {
    local service=$1
    local tag=$2

    echo "Building ${service} service..."
    docker build -t ${REGISTRY}/gfc-${service}:${tag} \
        --build-arg SERVICE=${service} \
        -f deployments/docker/Dockerfile .

    echo "Pushing ${service} image..."
    docker push ${REGISTRY}/gfc-${service}:${tag}
}

# 部署服务到Kubernetes
deploy_service() {
    local service=$1
    local tag=$2

    echo "Deploying ${service} service to ${ENVIRONMENT}..."
    
    # 应用配置
    kubectl apply -f deployments/kubernetes/${service}-service.yaml -n ${NAMESPACE}
    
    # 更新镜像
    kubectl set image deployment/${service}-service \
        ${service}=${REGISTRY}/gfc-${service}:${tag} \
        -n ${NAMESPACE}

    # 等待部署完成
    kubectl rollout status deployment/${service}-service -n ${NAMESPACE}
}

# 主部署流程
main() {
    check_requirements

    # 获取当前Git commit hash作为标签
    TAG=$(git rev-parse --short HEAD)

    # 创建namespace（如果不存在）
    kubectl create namespace ${NAMESPACE} --dry-run=client -o yaml | kubectl apply -f -

    # 部署配置和密钥
    kubectl apply -f deployments/kubernetes/config.yaml -n ${NAMESPACE}

    # 部署每个服务
    for service in "${SERVICES[@]}"; do
        build_and_push ${service} ${TAG}
        deploy_service ${service} ${TAG}
    done

    echo "Deployment completed successfully!"
}

main 