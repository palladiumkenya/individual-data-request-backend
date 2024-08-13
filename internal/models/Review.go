package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Review struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Content        string    `gorm:"type:text;not null"`
	UserID         uint      `gorm:"not null"`
	ReviewerID     uuid.UUID `gorm:"type:uuid"`
	ReviewThreadID uuid.UUID `gorm:"type:uuid"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type ReviewThread struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title     string    `gorm:"not null"`
	UserID    uint      `gorm:"not null"`
	RequestID uuid.UUID `gorm:"type:uuid"`
	CreatedAt time.Time
	Reviews   []Review `gorm:"foreignKey:ReviewThreadID"`
}
