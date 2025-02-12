package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zenor0/bookinfo/pkg/config"
	"github.com/zenor0/bookinfo/pkg/database"
	"github.com/zenor0/bookinfo/pkg/logger"
	"github.com/zenor0/bookinfo/pkg/middleware"
	"github.com/zenor0/bookinfo/services/details/internal/handler"
	"github.com/zenor0/bookinfo/services/details/internal/repository"
	"github.com/zenor0/bookinfo/services/details/internal/service"
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

	// 初始化依赖
	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)

	// 设置 Gin 路由
	router := gin.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recovery())
	router.Use(middleware.CORS())

	// 注册路由
	handler.RegisterRoutes(router, bookHandler)

	// 启动服务器
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	logger.Info("Starting details service", zap.String("addr", addr))
	if err := router.Run(addr); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
		os.Exit(1)
	}
}
