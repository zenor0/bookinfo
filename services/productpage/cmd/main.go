package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/zenor0/bookinfo/pkg/config"
	"github.com/zenor0/bookinfo/pkg/middleware"
	"github.com/zenor0/bookinfo/services/productpage/internal/client"
	"github.com/zenor0/bookinfo/services/productpage/internal/handler"
	"github.com/zenor0/bookinfo/services/productpage/internal/service"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

var (
	httpRequestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	}, []string{"method", "endpoint", "status"})

	httpRequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Duration of HTTP requests",
		Buckets: prometheus.DefBuckets,
	}, []string{"method", "endpoint"})
)

func initTracer() (*tracesdk.TracerProvider, error) {
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://jaeger:14268/api/traces")))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exporter),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("productpage"),
		)),
	)
	otel.SetTracerProvider(tp)
	return tp, nil
}

func main() {
	// Initialize tracer
	tp, err := initTracer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	// 加载配置
	cfg, err := config.LoadConfig("config")
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

	// Add Prometheus metrics endpoint
	http.Handle("/metrics", promhttp.Handler())

	// 启动服务器
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("Starting productpage service at %s, version: %s", addr, os.Getenv("SERVICE_VERSION"))
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
