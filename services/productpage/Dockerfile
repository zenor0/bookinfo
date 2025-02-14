# 构建阶段
FROM golang:1.22-alpine AS builder

WORKDIR /app

# 安装构建依赖
RUN apk add --no-cache git

# 复制 go.mod 和 go.sum
COPY go.mod ./
COPY go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -o productpage-service ./services/productpage/cmd/main.go

# 运行阶段
FROM alpine:3.19

WORKDIR /app

# 安装运行时依赖
RUN apk add --no-cache ca-certificates tzdata

# 复制配置文件和静态资源
COPY services/productpage/config/config.yaml /app/config/config.yaml
COPY services/productpage/static /app/static
COPY services/productpage/templates /app/templates

# 从构建阶段复制二进制文件
COPY --from=builder /app/productpage-service .

# 暴露端口
EXPOSE 9083

# 运行应用
CMD ["./productpage-service"] 