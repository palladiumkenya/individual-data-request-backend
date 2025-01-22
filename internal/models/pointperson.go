package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PointPersons struct {
	ID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Email string    `gorm:"size:100;not null"`
}

func GetPointPersonByID(DB *gorm.DB, Id uuid.UUID) (*PointPersons, error) {
	var PointPersons *PointPersons
	result := DB.First(&PointPersons, "id = ?", Id)
	return PointPersons, result.Error
}

func GetPointPersonByEmail(DB *gorm.DB, email string) ([]PointPersons, error) {
	var pointpersons []PointPersons
	result := DB.Find(&pointpersons, "email = ?", email)
	//result := DB.Model(&Approvers{}).Select("email").Scan(&approvers)
	return pointpersons, result.Error
}

func GetPointPersonsEmails(DB *gorm.DB) ([]string, error) {
	var pointpersons []string
	result := DB.Model(&PointPersons{}).Select("email").Scan(&pointpersons)
	return pointpersons, result.Error
}

func GetPointPerson(DB *gorm.DB) ([]PointPersons, error) {
	var PointPersons []PointPersons
	result := DB.Find(&PointPersons)
	return PointPersons, result.Error
}

func CreatePointPerson(DB *gorm.DB, PointPersons PointPersons) (uuid.UUID, error) {
	if err := DB.Create(&PointPersons).Error; err != nil {
		return uuid.UUID{}, err // Return the error and a zero UUID
	}
	return PointPersons.ID, nil
}

func DeletePointPerson(DB *gorm.DB, id uuid.UUID) error {
	result := DB.Delete(&PointPersons{}, "id = ?", id)
	return result.Error
}
