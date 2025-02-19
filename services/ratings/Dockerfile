# 构建阶段
FROM golang:1.23-alpine AS builder

WORKDIR /app

# 安装构建依赖
RUN apk add --no-cache git

# 复制 go.mod 和 go.sum
COPY go.mod ./
COPY go.sum ./

# 下载依赖
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

# 复制源代码
COPY . .

# 构建应用
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build -o ratings-service ./services/ratings/cmd/main.go

# 运行阶段
FROM alpine:3.19

WORKDIR /app

# 安装运行时依赖
RUN apk add --no-cache ca-certificates tzdata

# 复制配置文件
COPY services/ratings/config /app/config

# 从构建阶段复制二进制文件
COPY --from=builder /app/ratings-service .

# 暴露端口
EXPOSE 9080

# 运行应用
CMD ["./ratings-service"]