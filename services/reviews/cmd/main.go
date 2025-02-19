package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/zenor0/bookinfo/pkg/config"
	pb "github.com/zenor0/bookinfo/proto/reviews"
	"github.com/zenor0/bookinfo/services/reviews/internal/client"
	"github.com/zenor0/bookinfo/services/reviews/internal/repository"
	"github.com/zenor0/bookinfo/services/reviews/internal/server"
	"github.com/zenor0/bookinfo/services/reviews/internal/service"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig("config")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 连接数据库
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
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
	repo := repository.NewReviewRepository(db)
	svc := service.NewReviewService(repo, ratingsClient, formatter)
	srv := server.NewReviewServer(svc)

	// 启动 gRPC 服务器
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterReviewServiceServer(s, srv)

	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
