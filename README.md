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
├── services/              # 微服务组件
│   ├── productpage/      # 产品页面服务
│   ├── details/          # 商品详情服务
│   ├── reviews/          # 评论服务
│   └── ratings/          # 评分服务
└── tools/                 # 工具脚本
```

## 服务说明

- **ProductPage**: 提供 Web UI，聚合其他服务数据
- **Details**: 提供商品详细信息
- **Reviews**: 提供商品评论信息
- **Ratings**: 提供商品评分信息

## 开发环境要求

- Go 1.21+
- Docker
- Kubernetes 1.28+
- Istio 1.20+
- Apache APISIX
- PostgreSQL

## 快速开始

1. 克隆项目
```bash
git clone [项目地址]
cd microservice-bookinfo
```

2. 安装依赖
```bash
go mod download
```

3. 本地开发
```bash
# 运行特定服务
cd services/[service-name]
go run main.go
```

## 部署指南

详细的部署指南请参考 `docs/deployment.md`

## 贡献指南

1. Fork 本仓库
2. 创建您的特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交您的更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启一个 Pull Request

## 许可证

MIT License - 详见 LICENSE 文件
