package service

import (
	"errors"

	"github.com/zenor0/bookinfo/services/ratings/internal/model"
	"github.com/zenor0/bookinfo/services/ratings/internal/repository"
)

type RatingService struct {
	repo *repository.RatingRepository
}

func NewRatingService(repo *repository.RatingRepository) *RatingService {
	return &RatingService{repo: repo}
}

func (s *RatingService) CreateRating(rating *model.Rating) error {
	if rating.Rating < 1 || rating.Rating > 5 {
		return errors.New("rating must be between 1 and 5")
	}
	return s.repo.Create(rating)
}

func (s *RatingService) GetRating(id uint) (*model.Rating, error) {
	return s.repo.GetByID(id)
}

func (s *RatingService) GetProductRatings(productID string) ([]model.Rating, error) {
	return s.repo.GetByProductID(productID)
}

func (s *RatingService) UpdateRating(rating *model.Rating) error {
	if rating.Rating < 1 || rating.Rating > 5 {
		return errors.New("rating must be between 1 and 5")
	}
	return s.repo.Update(rating)
}

func (s *RatingService) DeleteRating(id uint) error {
	return s.repo.Delete(id)
}

func (s *RatingService) GetAverageRating(productID string) (float64, error) {
	return s.repo.GetAverageRating(productID)
}
