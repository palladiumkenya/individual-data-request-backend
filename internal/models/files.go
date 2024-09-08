package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Files struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CreatedBy uuid.UUID  `gorm:"type:uuid;null"`
	RequestId *uuid.UUID `gorm:"type:uuid;null"`
	Request   *Requests  `gorm:"foreignKey:RequestId"`
	FileName  string     `gorm:"size:100;not null"`
	FileURL   string     `gorm:"size:500;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func UploadFiles(DB *gorm.DB, files *Files) error {
	DB.Create(&files)
	return nil
}


func FetchFiles(DB *gorm.DB, FileType string, RequestId uuid.UUID) (*Files, error) {
	var requestFile *Files
	result := DB.First(&requestFile, "request = ? and file_name = ?", RequestId, FileType)
	return requestFile, result.Error
}
