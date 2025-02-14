package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zenor0/bookinfo/services/productpage/internal/service"
)

type ProductPageHandler struct {
	service *service.ProductPageService
}

func NewProductPageHandler(service *service.ProductPageService) *ProductPageHandler {
	return &ProductPageHandler{service: service}
}

func (h *ProductPageHandler) GetProductPage(c *gin.Context) {
	bookID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}

	// 从请求头或查询参数中获取用户信息
	user := c.GetHeader("end-user")
	if user == "" {
		user = c.Query("user")
	}

	productPage, err := h.service.GetProductPage(uint(bookID), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, productPage)
}

func (h *ProductPageHandler) GetBookList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	books, total, err := h.service.GetBookList(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": books,
		"meta": gin.H{
			"page":      page,
			"page_size": pageSize,
			"total":     total,
		},
	})
}

func (h *ProductPageHandler) SearchBooks(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "search query is required"})
		return
	}

	books, err := h.service.SearchBooks(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

// 健康检查接口
func (h *ProductPageHandler) HealthCheck(c *gin.Context) {
	c.Status(http.StatusOK)
}

func RegisterRoutes(router *gin.Engine, handler *ProductPageHandler) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", handler.HealthCheck)
		v1.GET("/books", handler.GetBookList)
		v1.GET("/books/search", handler.SearchBooks)
		v1.GET("/books/:id", handler.GetProductPage)
	}

	// 静态文件服务
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}
