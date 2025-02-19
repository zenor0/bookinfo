.PHONY: help dev-db dev-db-stop dev-db-clean dev-skaffold build test lint proto clean dev prod

# 显示帮助信息
help:
	@echo "可用的 make 命令："
	@echo "  make dev-db          - 启动开发环境数据库"
	@echo "  make dev-db-stop     - 停止开发环境数据库"
	@echo "  make dev-db-clean    - 清理开发环境数据库（包括数据）"
	@echo "  make dev-skaffold    - 启动 skaffold 开发环境"
	@echo "  make build           - 构建所有服务"
	@echo "  make test            - 运行所有测试"
	@echo "  make lint            - 运行代码检查"
	@echo "  make proto           - 生成 protobuf 代码"
	@echo "  make clean           - 清理构建产物"
	@echo "  make dev             - 启动完整开发环境（数据库 + skaffold）"
	@echo "  make prod            - 部署到生产环境"

# 开发环境数据库管理
dev-db:
	docker-compose -f docker-compose.dev.yaml up -d
	@echo "开发数据库已启动"
	@echo "PostgreSQL: localhost:5432"
	@echo "pgAdmin: http://localhost:5050"
	@echo "pgAdmin 登录信息："
	@echo "  邮箱: admin@bookinfo.local"
	@echo "  密码: admin"

dev-db-stop:
	docker-compose -f docker-compose.dev.yaml down

dev-db-clean:
	docker-compose -f docker-compose.dev.yaml down -v

# Skaffold 开发
dev-skaffold:
	skaffold dev --profile dev

# 构建
build:
	go mod tidy
	go build -v ./...

# 测试
test:
	go test -v -race -cover ./...

# 代码检查
lint:
	golangci-lint run ./...
	go vet ./...

# 生成 protobuf 代码
proto:
	@echo "生成 protobuf 代码..."
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/**/*.proto

# 清理
clean:
	go clean
	rm -rf bin/
	find . -type f -name '*.pb.go' -delete

# 完整开发环境
dev: dev-db dev-skaffold

# 生产环境部署
prod:
	@echo "正在部署到生产环境..."
	kubectl apply -f deployments/kubernetes/db-credentials.yaml
	skaffold run --profile prod

# 创建本地 k8s 集群（使用 kind）
k8s-local:
	kind create cluster --name bookinfo --config deployments/kind/config.yaml
	kubectl create namespace bookinfo
	kubectl config set-context --current --namespace=bookinfo

# 删除本地 k8s 集群
k8s-local-clean:
	kind delete cluster --name bookinfo

# 安装开发依赖
install-dev-tools:
	@echo "安装开发工具..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@echo "请确保已安装以下工具："
	@echo "- Docker"
	@echo "- kubectl"
	@echo "- skaffold"
	@echo "- kind (用于本地 k8s 开发)"
	@echo "- protoc (protobuf 编译器)"

# 生成 swagger 文档
swagger:
	swag init -g services/productpage/cmd/main.go -o services/productpage/docs
	swag init -g services/ratings/cmd/main.go -o services/ratings/docs

# 检查服务健康状态
check-health:
	@echo "检查服务健康状态..."
	@curl -s http://localhost:9083/health || echo "ProductPage 服务未运行"
	@curl -s http://localhost:9080/health || echo "Ratings 服务未运行"
	@curl -s http://localhost:9081/health || echo "Details 服务未运行"
	@curl -s http://localhost:9082/health || echo "Reviews 服务未运行"

# 查看服务日志
logs:
	@echo "请选择要查看的服务日志："
	@echo "1) make logs-productpage"
	@echo "2) make logs-details"
	@echo "3) make logs-reviews"
	@echo "4) make logs-ratings"

logs-productpage:
	kubectl logs -f -l app=productpage

logs-details:
	kubectl logs -f -l app=details

logs-reviews:
	kubectl logs -f -l app=reviews

logs-ratings:
	kubectl logs -f -l app=ratings 