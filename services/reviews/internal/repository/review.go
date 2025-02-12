package repository

import (
	"github.com/zenor0/bookinfo/services/reviews/internal/model"
	"gorm.io/gorm"
)

type ReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) *ReviewRepository {
	return &ReviewRepository{db: db}
}

func (r *ReviewRepository) Create(review *model.Review) error {
	return r.db.Create(review).Error
}

func (r *ReviewRepository) GetByID(id uint) (*model.Review, error) {
	var review model.Review
	err := r.db.First(&review, id).Error
	if err != nil {
		return nil, err
	}
	return &review, nil
}

func (r *ReviewRepository) GetByBookID(bookID uint, page, pageSize int) ([]model.Review, int64, error) {
	var reviews []model.Review
	var total int64

	err := r.db.Model(&model.Review{}).Where("book_id = ?", bookID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = r.db.Where("book_id = ?", bookID).
		Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&reviews).Error
	if err != nil {
		return nil, 0, err
	}

	return reviews, total, nil
}

func (r *ReviewRepository) GetByUserID(userID uint) ([]model.Review, error) {
	var reviews []model.Review
	err := r.db.Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&reviews).Error
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

func (r *ReviewRepository) Update(review *model.Review) error {
	return r.db.Save(review).Error
}

func (r *ReviewRepository) Delete(id uint) error {
	return r.db.Delete(&model.Review{}, id).Error
}
