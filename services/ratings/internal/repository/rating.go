package repository

import (
	"github.com/zenor0/bookinfo/services/ratings/internal/model"
	"gorm.io/gorm"
)

type RatingRepository struct {
	db *gorm.DB
}

func NewRatingRepository(db *gorm.DB) *RatingRepository {
	return &RatingRepository{db: db}
}

func (r *RatingRepository) Create(rating *model.Rating) error {
	return r.db.Create(rating).Error
}

func (r *RatingRepository) GetByID(id uint) (*model.Rating, error) {
	var rating model.Rating
	err := r.db.First(&rating, id).Error
	if err != nil {
		return nil, err
	}
	return &rating, nil
}

func (r *RatingRepository) GetByProductID(productID string) ([]model.Rating, error) {
	var ratings []model.Rating
	err := r.db.Where("product_id = ?", productID).Find(&ratings).Error
	if err != nil {
		return nil, err
	}
	return ratings, nil
}

func (r *RatingRepository) Update(rating *model.Rating) error {
	return r.db.Save(rating).Error
}

func (r *RatingRepository) Delete(id uint) error {
	return r.db.Delete(&model.Rating{}, id).Error
}

func (r *RatingRepository) GetAverageRating(productID string) (float64, error) {
	var avgRating float64
	err := r.db.Model(&model.Rating{}).
		Select("COALESCE(AVG(rating), 0)").
		Where("product_id = ?", productID).
		Scan(&avgRating).Error
	if err != nil {
		return 0, err
	}
	return avgRating, nil
}
