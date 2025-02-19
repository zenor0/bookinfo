FROM golang:1.23-alpine AS builder

WORKDIR /app

# 安装构建依赖
RUN apk add --no-cache git

# 先只复制依赖文件
COPY go.mod go.sum ./

# 使用 BuildKit 缓存挂载
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

# 复制源代码
COPY . .

# 构建应用，同样使用缓存
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build -o productpage-service ./services/productpage/cmd/main.go

# 运行阶段
FROM alpine:3.19

WORKDIR /app

# 安装运行时依赖
RUN apk add --no-cache ca-certificates tzdata

# 复制配置文件和静态资源
COPY services/productpage/config /app/config
COPY services/productpage/static /app/static
COPY services/productpage/templates /app/templates

# 从构建阶段复制二进制文件
COPY --from=builder /app/productpage-service .

# 暴露端口
EXPOSE 9083

# 运行应用
CMD ["./productpage-service"] 