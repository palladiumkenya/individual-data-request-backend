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

//func CreateRequester(ctx context.Context, pool *pgxpool.Pool, requester *Requesters) error {
//	_, err := pool.Exec(ctx, "INSERT INTO Requesters (email) VALUES ($1)",
//		requester.Email)
//	if err != nil {
//		return err
//	}
//	return nil
//}

func CheckUserRequester(DB *gorm.DB, emailStr string) (Requesters, error) {
	var requester Requesters
	result := DB.Find(&requester, "email = ?", emailStr)
	return requester, result.Error
}
