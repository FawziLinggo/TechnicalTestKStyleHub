package models

import (
	"database/sql"
	"time"
)

type LikeReview struct {
	ID              string `gorm:"primaryKey;default:(UUID())"`
	ProductReviewID string `json:"product_review_id"`
	MemberID        string `json:"member_id"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       sql.NullTime
}

func (LikeReview) TableName() string {
	return "like_review"
}
