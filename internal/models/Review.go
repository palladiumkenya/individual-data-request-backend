package models

import (
	"time"
)

type Review struct {
	ID             uint   `gorm:"primaryKey"`
	Content        string `gorm:"type:text;not null"`
	UserID         uint   `gorm:"not null"`
	ReviewerID     uint   `gorm:"not null"`
	ReviewThreadID uint   `gorm:"not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type ReviewThread struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	UserID    uint   `gorm:"not null"`
	RequestID uint   `gorm:"not null"`
	CreatedAt time.Time
	Reviews   []Review `gorm:"foreignKey:ReviewThreadID"`
}
