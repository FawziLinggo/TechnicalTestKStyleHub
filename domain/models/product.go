package models

import (
	"database/sql"
	"time"
)

type Product struct {
	ID            string `gorm:"primaryKey;default:(UUID())"`
	Name          string
	Price         float64
	ProductReview []ProductReview `gorm:"foreignKey:ProductID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     sql.NullTime
}

func (Product) TableName() string {
	return "product"
}

type ProductReview struct {
	ID         string       `gorm:"primaryKey;default:(UUID())"`
	ProductID  string       `json:"product_id"`
	MemberID   string       `json:"member_id"`
	DescReview string       `json:"desc_review"`
	LikeReview []LikeReview `gorm:"foreignKey:ProductReviewID"`
	Member     Member       `gorm:"foreignKey:MemberID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  sql.NullTime
}

func (ProductReview) TableName() string {
	return "product_review"
}
