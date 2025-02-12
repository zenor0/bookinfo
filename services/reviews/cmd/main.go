package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zenor0/bookinfo/pkg/config"
	"github.com/zenor0/bookinfo/pkg/database"
	"github.com/zenor0/bookinfo/pkg/logger"
	"github.com/zenor0/bookinfo/pkg/middleware"
	"github.com/zenor0/bookinfo/services/reviews/internal/client"
	"github.com/zenor0/bookinfo/services/reviews/internal/handler"
	"github.com/zenor0/bookinfo/services/reviews/internal/repository"
	"github.com/zenor0/bookinfo/services/reviews/internal/service"
	"go.uber.org/zap"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		logger.Fatal("Failed to load config", zap.Error(err))
		os.Exit(1)
	}

	// 初始化数据库连接
	db, err := database.NewPostgresConnection(&cfg.Database)
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
		os.Exit(1)
	}

	// 获取服务版本
	version := os.Getenv("SERVICE_VERSION")
	if version == "" {
		version = "v1"
	}

	// 根据版本选择评分格式化器
	var formatter client.RatingFormatter
	switch version {
	case "v1":
		formatter = &client.V1Formatter{}
	case "v2":
		formatter = &client.V2Formatter{}
	case "v3":
		formatter = &client.V3Formatter{}
	default:
		formatter = &client.V1Formatter{}
	}

	// 初始化依赖
	ratingsClient := client.NewRatingsClient("http://ratings:9080")
	reviewRepo := repository.NewReviewRepository(db)
	reviewService := service.NewReviewService(reviewRepo, ratingsClient, formatter)
	reviewHandler := handler.NewReviewHandler(reviewService)

	// 设置 Gin 路由
	router := gin.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recovery())
	router.Use(middleware.CORS())

	// 注册路由
	handler.RegisterRoutes(router, reviewHandler)

	// 启动服务器
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	logger.Info("Starting reviews service",
		zap.String("addr", addr),
		zap.String("version", version))
	if err := router.Run(addr); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
		os.Exit(1)
	}
}
