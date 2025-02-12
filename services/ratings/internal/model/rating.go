package model

import (
	"time"

	"gorm.io/gorm"
)

type Rating struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	ProductID string    `json:"product_id" gorm:"index"`
	Rating    int       `json:"rating" gorm:"check:rating >= 1 AND rating <= 5"`
	Reviewer  string    `json:"reviewer"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (r *Rating) BeforeCreate(tx *gorm.DB) error {
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
	return nil
}

func (r *Rating) BeforeUpdate(tx *gorm.DB) error {
	r.UpdatedAt = time.Now()
	return nil
}
