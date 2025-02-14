package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zenor0/bookinfo/pkg/config"
	"github.com/zenor0/bookinfo/pkg/middleware"
	"github.com/zenor0/bookinfo/services/productpage/internal/client"
	"github.com/zenor0/bookinfo/services/productpage/internal/handler"
	"github.com/zenor0/bookinfo/services/productpage/internal/service"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 初始化服务客户端
	detailsClient, err := client.NewDetailsClient("details:9081")
	if err != nil {
		log.Fatalf("Failed to create details client: %v", err)
	}
	defer detailsClient.Close()

	reviewsClient, err := client.NewReviewsClient("reviews:9082")
	if err != nil {
		log.Fatalf("Failed to create reviews client: %v", err)
	}
	defer reviewsClient.Close()

	ratingsClient := client.NewRatingsClient("http://ratings:9080")

	// 初始化依赖
	productPageService := service.NewProductPageService(detailsClient, reviewsClient, ratingsClient)
	productPageHandler := handler.NewProductPageHandler(productPageService)

	// 设置 Gin 路由
	router := gin.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recovery())
	router.Use(middleware.CORS())

	// 注册路由
	handler.RegisterRoutes(router, productPageHandler)

	// 启动服务器
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("Starting productpage service at %s, version: %s", addr, os.Getenv("SERVICE_VERSION"))
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
