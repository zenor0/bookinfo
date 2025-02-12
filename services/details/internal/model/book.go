package model

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	ISBN        string    `json:"isbn" gorm:"uniqueIndex"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Publisher   string    `json:"publisher"`
	Language    string    `json:"language"`
	Pages       int       `json:"pages"`
	Year        int       `json:"year"`
	Price       float64   `json:"price"`
	Description string    `json:"description" gorm:"type:text"`
	CoverImage  string    `json:"cover_image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (b *Book) BeforeCreate(tx *gorm.DB) error {
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
	return nil
}

func (b *Book) BeforeUpdate(tx *gorm.DB) error {
	b.UpdatedAt = time.Now()
	return nil
}
