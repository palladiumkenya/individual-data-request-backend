package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Requesters struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Email        string    `gorm:"size:100;not null"`
	Name         string    `gorm:"size:100"`
	Organization string    `gorm:"size:100"`
}

func GetRequesterByID(DB *gorm.DB, Id uuid.UUID) (*Requesters, error) {
	var requester *Requesters
	result := DB.First(&requester, "id = ?", Id)
	return requester, result.Error
}

func GetRequesters(DB *gorm.DB) ([]Requesters, error) {
	var requesters []Requesters
	result := DB.Find(&requesters)
	return requesters, result.Error
}

func CreateRequester(DB *gorm.DB, requester Requesters) (uuid.UUID, error) {
	if err := DB.Create(&requester).Error; err != nil {
		return uuid.UUID{}, err // Return the error and a zero UUID
	}
	return requester.ID, nil

}

func CheckUserRequester(DB *gorm.DB, emailStr string) (Requesters, error) {
	var requester Requesters
	result := DB.Find(&requester, "email = ?", emailStr)
	return requester, result.Error
}
