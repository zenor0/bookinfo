package service

import (
	"errors"
	"strconv"

	"github.com/zenor0/bookinfo/services/reviews/internal/client"
	"github.com/zenor0/bookinfo/services/reviews/internal/model"
	"github.com/zenor0/bookinfo/services/reviews/internal/repository"
)

type ReviewService struct {
	repo          *repository.ReviewRepository
	ratingsClient *client.RatingsClient
	formatter     client.RatingFormatter
}

func NewReviewService(repo *repository.ReviewRepository, ratingsClient *client.RatingsClient, formatter client.RatingFormatter) *ReviewService {
	return &ReviewService{
		repo:          repo,
		ratingsClient: ratingsClient,
		formatter:     formatter,
	}
}

func (s *ReviewService) CreateReview(review *model.Review) error {
	if review.BookID == 0 {
		return errors.New("book id is required")
	}
	if review.UserID == 0 {
		return errors.New("user id is required")
	}
	if review.Content == "" {
		return errors.New("content is required")
	}
	return s.repo.Create(review)
}

func (s *ReviewService) GetReview(id uint) (*model.ReviewResponse, error) {
	review, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return s.enrichReviewWithRating(review)
}

func (s *ReviewService) GetBookReviews(bookID uint, page, pageSize int) ([]model.ReviewResponse, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	reviews, total, err := s.repo.GetByBookID(bookID, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	var enrichedReviews []model.ReviewResponse
	for _, review := range reviews {
		enrichedReview, err := s.enrichReviewWithRating(&review)
		if err != nil {
			return nil, 0, err
		}
		enrichedReviews = append(enrichedReviews, *enrichedReview)
	}

	return enrichedReviews, total, nil
}

func (s *ReviewService) GetUserReviews(userID uint) ([]model.ReviewResponse, error) {
	reviews, err := s.repo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	var enrichedReviews []model.ReviewResponse
	for _, review := range reviews {
		enrichedReview, err := s.enrichReviewWithRating(&review)
		if err != nil {
			return nil, err
		}
		enrichedReviews = append(enrichedReviews, *enrichedReview)
	}

	return enrichedReviews, nil
}

func (s *ReviewService) UpdateReview(review *model.Review) error {
	if review.ID == 0 {
		return errors.New("review id is required")
	}
	return s.repo.Update(review)
}

func (s *ReviewService) DeleteReview(id uint) error {
	return s.repo.Delete(id)
}

func (s *ReviewService) enrichReviewWithRating(review *model.Review) (*model.ReviewResponse, error) {
	rating, err := s.ratingsClient.GetProductRating(strconv.FormatUint(uint64(review.BookID), 10))
	if err != nil {
		// 如果获取评分失败，不影响评论的展示
		rating = 0
	}

	ratingHTML, color := s.formatter.FormatRating(rating)

	return &model.ReviewResponse{
		Review:     *review,
		Rating:     rating,
		RatingHTML: ratingHTML,
		Color:      color,
	}, nil
}
