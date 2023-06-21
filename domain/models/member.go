package models

import (
	"database/sql"
	"time"
)

type Member struct {
	ID        string `gorm:"primaryKey;default:(UUID())"`
	Username  string
	Gender    string
	Skintype  string
	Skincolor string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

func (Member) TableName() string {
	return "member"
}
