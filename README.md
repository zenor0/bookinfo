# BookInfo 微服务示例项目

这是一个使用 Go 语言实现的 BookInfo 微服务示例项目，展示了如何使用现代云原生技术栈构建微服务应用。

## 技术栈

- 开发语言：Go 1.21+
- Web 框架：Gin
- 数据库：PostgreSQL
- 服务网格：Istio 1.20
- API 网关：Apache APISIX
- 容器编排：Kubernetes 1.28+

## 项目结构

```
.
├── api/                    # API 定义
├── build/                  # 构建脚本
├── deployments/            # K8s 和 Istio 配置
│   ├── kubernetes/        # Kubernetes 资源配置
│   ├── istio/            # Istio 配置
│   └── apisix/           # APISIX 网关配置
├── docs/                  # 文档
├── pkg/                   # 公共包
├── scripts/               # 脚本文件
│   ├── db/               # 数据库脚本
│   └── deploy.sh         # 部署脚本
├── services/              # 微服务组件
│   ├── productpage/      # 产品页面服务
│   ├── details/          # 商品详情服务
│   ├── reviews/          # 评论服务
│   └── ratings/          # 评分服务
└── docker-compose.yml     # 本地开发配置
```

## 服务说明

- **ProductPage**: 提供 Web UI，聚合其他服务数据
- **Details**: 提供商品详细信息
- **Reviews**: 提供商品评论信息（有三个版本）
  - v1: 不显示评分
  - v2: 显示黑色星星评分
  - v3: 显示红色星星评分
- **Ratings**: 提供商品评分信息

## 开发环境要求

- Go 1.21+
- Docker & Docker Compose
- Kubernetes 1.28+
- Istio 1.20+
- Apache APISIX
- PostgreSQL

## 本地开发

1. 克隆项目
```bash
git clone [项目地址]
cd microservice-bookinfo
```

2. 安装依赖
```bash
go mod download
```

3. 启动本地开发环境
```bash
# 启动所有服务
docker-compose up -d

# 查看服务状态
docker-compose ps

# 查看服务日志
docker-compose logs -f [service-name]
```

4. 访问服务
- ProductPage: http://localhost:9083
- Details: http://localhost:9081
- Reviews v1: http://localhost:9082
- Reviews v2: http://localhost:9084
- Reviews v3: http://localhost:9085
- Ratings: http://localhost:9080

## Kubernetes 部署

1. 确保已安装必要的工具
- kubectl
- docker
- istioctl
- helm (用于安装 APISIX)

2. 执行部署脚本
```bash
# 添加执行权限
chmod +x scripts/deploy.sh

# 执行部署
./scripts/deploy.sh
```

3. 验证部署
```bash
# 查看所有 Pod 状态
kubectl get pods -n bookinfo

# 查看服务状态
kubectl get svc -n bookinfo

# 查看 Istio 配置
kubectl get virtualservice,destinationrule -n bookinfo
```

## 开发指南

1. 添加新功能
- 创建功能分支
- 实现功能
- 编写测试
- 提交 PR

2. 运行测试
```bash
go test ./...
```

3. 构建服务
```bash
# 构建单个服务
docker build -t bookinfo/[service-name]:[version] -f services/[service-name]/Dockerfile .

# 构建所有服务
./scripts/deploy.sh build_images
```

## 故障排除

1. 检查服务日志
```bash
# Docker Compose 环境
docker-compose logs [service-name]

# Kubernetes 环境
kubectl logs -n bookinfo [pod-name]
```

2. 检查服务健康状态
```bash
# 健康检查接口
curl http://localhost:[port]/api/v1/health
```

3. 常见问题
- 数据库连接失败：检查数据库配置和连接信息
- 服务间调用失败：检查服务发现配置和网络策略
- 评分不显示：检查用户权限和路由规则

## 贡献指南

1. Fork 本仓库
2. 创建您的特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交您的更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启一个 Pull Request

## 许可证

MIT License - 详见 LICENSE 文件
