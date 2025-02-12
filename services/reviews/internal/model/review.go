package model

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	BookID    uint      `json:"book_id" gorm:"index"`
	UserID    uint      `json:"user_id"`
	Username  string    `json:"username"`
	Content   string    `json:"content" gorm:"type:text"`
	Rating    float64   `json:"rating,omitempty"` // 从 ratings 服务获取
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (r *Review) BeforeCreate(tx *gorm.DB) error {
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
	return nil
}

func (r *Review) BeforeUpdate(tx *gorm.DB) error {
	r.UpdatedAt = time.Now()
	return nil
}

// ReviewResponse 包含评论和评分信息的响应结构
type ReviewResponse struct {
	Review
	Rating     float64 `json:"rating,omitempty"`
	RatingHTML string  `json:"rating_html,omitempty"` // 用于不同版本的评分展示
	Color      string  `json:"color,omitempty"`       // 用于不同版本的颜色展示
}
