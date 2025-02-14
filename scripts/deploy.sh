#!/bin/bash

# 设置变量
REGISTRY="bookinfo"
VERSION="latest"
NAMESPACE="bookinfo"

# 构建 Docker 镜像
build_images() {
    echo "Building Docker images..."
    
    # 构建 Details 服务
    docker build -t $REGISTRY/details:$VERSION -f services/details/Dockerfile .
    
    # 构建 Reviews 服务的三个版本
    docker build -t $REGISTRY/reviews:v1 --build-arg SERVICE_VERSION=v1 -f services/reviews/Dockerfile .
    docker build -t $REGISTRY/reviews:v2 --build-arg SERVICE_VERSION=v2 -f services/reviews/Dockerfile .
    docker build -t $REGISTRY/reviews:v3 --build-arg SERVICE_VERSION=v3 -f services/reviews/Dockerfile .
    
    # 构建 Ratings 服务
    docker build -t $REGISTRY/ratings:$VERSION -f services/ratings/Dockerfile .
    
    # 构建 ProductPage 服务
    docker build -t $REGISTRY/productpage:$VERSION -f services/productpage/Dockerfile .
}

# 创建命名空间
create_namespace() {
    echo "Creating namespace..."
    kubectl create namespace $NAMESPACE --dry-run=client -o yaml | kubectl apply -f -
}

# 部署数据库
deploy_database() {
    echo "Deploying database..."
    kubectl apply -f deployments/kubernetes/postgres.yaml
    
    # 等待数据库就绪
    echo "Waiting for database to be ready..."
    kubectl wait --for=condition=ready pod -l app=postgres -n $NAMESPACE --timeout=300s
    
    # 初始化数据库
    echo "Initializing database..."
    kubectl create configmap db-init-scripts --from-file=scripts/db/init.sql -n $NAMESPACE --dry-run=client -o yaml | kubectl apply -f -
}

# 部署服务
deploy_services() {
    echo "Deploying services..."
    
    # 部署 ConfigMaps
    kubectl apply -f deployments/kubernetes/details.yaml
    kubectl apply -f deployments/kubernetes/reviews.yaml
    kubectl apply -f deployments/kubernetes/ratings.yaml
    kubectl apply -f deployments/kubernetes/productpage.yaml
}

# 部署 Istio 配置
deploy_istio() {
    echo "Deploying Istio configurations..."
    
    kubectl apply -f deployments/istio/details-vs.yaml
    kubectl apply -f deployments/istio/reviews-vs.yaml
    kubectl apply -f deployments/istio/ratings-vs.yaml
    kubectl apply -f deployments/istio/productpage-vs.yaml
}

# 部署 APISIX 配置
deploy_apisix() {
    echo "Deploying APISIX configurations..."
    
    kubectl apply -f deployments/apisix/details-route.yaml
    kubectl apply -f deployments/apisix/reviews-route.yaml
    kubectl apply -f deployments/apisix/ratings-route.yaml
    kubectl apply -f deployments/apisix/productpage-route.yaml
}

# 主函数
main() {
    echo "Starting deployment..."
    
    # 检查 kubectl 是否可用
    if ! command -v kubectl &> /dev/null; then

        echo "kubectl not found"
        exit 1
    fi
    
    # 检查 docker 是否可用
    if ! command -v docker &> /dev/null; then
        echo "docker not found"
        exit 1
    fi
    
    # 执行部署步骤
    build_images
    create_namespace
    deploy_database
    deploy_services
    deploy_istio
    deploy_apisix
    
    echo "Deployment completed successfully!"
}

# 执行主函数
main 