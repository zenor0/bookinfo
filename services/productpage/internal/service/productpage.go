package service

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/zenor0/bookinfo/services/productpage/internal/client"
	"github.com/zenor0/bookinfo/services/productpage/internal/model"
)

type ProductPageService struct {
	detailsClient *client.DetailsClient
	reviewsClient *client.ReviewsClient
	ratingsClient *client.RatingsClient
}

func NewProductPageService(detailsClient *client.DetailsClient, reviewsClient *client.ReviewsClient, ratingsClient *client.RatingsClient) *ProductPageService {
	return &ProductPageService{
		detailsClient: detailsClient,
		reviewsClient: reviewsClient,
		ratingsClient: ratingsClient,
	}
}

func (s *ProductPageService) GetProductPage(bookID uint, user string) (*model.ProductPage, error) {
	// 创建响应结构
	productPage := &model.ProductPage{
		User:    user,
		Version: os.Getenv("SERVICE_VERSION"),
	}

	// 获取图书详情
	details, err := s.detailsClient.GetBookDetails(bookID)
	if err != nil {
		productPage.Error = fmt.Sprintf("Error fetching book details: %v", err)
		return productPage, nil
	}
	productPage.Book = details

	// 获取评论信息
	reviews, err := s.reviewsClient.GetBookReviews(bookID, user)
	if err != nil {
		productPage.Error = fmt.Sprintf("Error fetching reviews: %v", err)
		return productPage, nil
	}
	productPage.Reviews = reviews

	// 获取评分信息
	rating, err := s.ratingsClient.GetBookRating(bookID)
	if err != nil {
		productPage.Error = fmt.Sprintf("Error fetching ratings: %v", err)
		return productPage, nil
	}
	productPage.Rating = rating

	return productPage, nil
}

func (s *ProductPageService) GetBookList(page, pageSize int) ([]model.BookDetails, int64, error) {
	// 调用 details 服务获取图书列表
	ctx := context.Background()
	resp, err := s.detailsClient.GetBookList(ctx, page, pageSize)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get book list: %w", err)
	}

	// 转换响应数据
	var books []model.BookDetails
	for _, book := range resp.Data {
		year, _ := strconv.Atoi(book.Year)
		books = append(books, model.BookDetails{
			ID:          uint(book.Id),
			Title:       book.Title,
			Author:      book.Author,
			ISBN:        book.Isbn,
			Publisher:   book.Publisher,
			Language:    book.Language,
			Pages:       int(book.Pages),
			Year:        year,
			Price:       float64(book.Price),
			Description: book.Description,
			CoverImage:  book.CoverImage,
		})
	}

	return books, int64(resp.Meta.Total), nil
}

func (s *ProductPageService) SearchBooks(query string) ([]model.BookDetails, error) {
	if query == "" {
		return nil, fmt.Errorf("search query cannot be empty")
	}

	// 调用 details 服务搜索图书
	ctx := context.Background()
	resp, err := s.detailsClient.SearchBooks(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to search books: %w", err)
	}

	// 转换响应数据
	var books []model.BookDetails
	for _, book := range resp.Books {
		year, _ := strconv.Atoi(book.Year)
		books = append(books, model.BookDetails{
			ID:          uint(book.Id),
			Title:       book.Title,
			Author:      book.Author,
			ISBN:        book.Isbn,
			Publisher:   book.Publisher,
			Language:    book.Language,
			Pages:       int(book.Pages),
			Year:        year,
			Price:       float64(book.Price),
			Description: book.Description,
			CoverImage:  book.CoverImage,
		})
	}

	return books, nil
}
