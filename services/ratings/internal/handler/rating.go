package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zenor0/bookinfo/services/ratings/internal/model"
	"github.com/zenor0/bookinfo/services/ratings/internal/service"
)

type RatingHandler struct {
	service *service.RatingService
}

func NewRatingHandler(service *service.RatingService) *RatingHandler {
	return &RatingHandler{service: service}
}

func (h *RatingHandler) CreateRating(c *gin.Context) {
	var rating model.Rating
	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateRating(&rating); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, rating)
}

func (h *RatingHandler) GetRating(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	rating, err := h.service.GetRating(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "rating not found"})
		return
	}

	c.JSON(http.StatusOK, rating)
}

func (h *RatingHandler) GetProductRatings(c *gin.Context) {
	productID := c.Param("productId")
	ratings, err := h.service.GetProductRatings(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ratings)
}

func (h *RatingHandler) UpdateRating(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var rating model.Rating
	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rating.ID = uint(id)
	if err := h.service.UpdateRating(&rating); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rating)
}

func (h *RatingHandler) DeleteRating(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.service.DeleteRating(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *RatingHandler) GetAverageRating(c *gin.Context) {
	productID := c.Param("productId")
	avgRating, err := h.service.GetAverageRating(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"average_rating": avgRating})
}

func RegisterRoutes(router *gin.Engine, handler *RatingHandler) {
	v1 := router.Group("/api/v1")
	{
		v1.POST("/ratings", handler.CreateRating)
		v1.GET("/ratings/:id", handler.GetRating)
		v1.GET("/products/:productId/ratings", handler.GetProductRatings)
		v1.PUT("/ratings/:id", handler.UpdateRating)
		v1.DELETE("/ratings/:id", handler.DeleteRating)
		v1.GET("/products/:productId/average-rating", handler.GetAverageRating)
	}
}
