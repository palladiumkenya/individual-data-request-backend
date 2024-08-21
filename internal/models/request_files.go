package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type RequestFiles struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Assignee uuid.UUID `gorm:"type:uuid"`
	Request  uuid.UUID `gorm:"foreignKey:Requestor_id"`
	FileName string    `gorm:"size:100;not null"`
	FileURL  string    `gorm:"size:500;not null"`
}

func UploadFiles(DB *gorm.DB, files *RequestFiles) error {
	DB.Create(&files)
	return nil
}
