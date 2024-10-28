package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Approvers struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Email         string    `gorm:"size:100;not null"`
	Approver_Type string    `gorm:"size:100;not null"`
}

func GetApproversByID(DB *gorm.DB, Id uuid.UUID) (*Approvers, error) {
	var approvers *Approvers
	result := DB.First(&approvers, "id = ?", Id)
	return approvers, result.Error
}

func GetApproversByType(DB *gorm.DB, approver_type string) (*Approvers, error) {
	var approvers *Approvers
	result := DB.First(&approvers, "approver_type = ?", approver_type)
	return approvers, result.Error
}

func GetApproverss(DB *gorm.DB) ([]Approvers, error) {
	var approverss []Approvers
	result := DB.Find(&approverss)
	return approverss, result.Error
}

func CreateApprover(DB *gorm.DB, approver *Approvers) error {
	DB.Create(&Approvers{ID: uuid.MustParse("44f75fd1-67b7-411c-8c9e-311afd5cf1eb"),
		Email: "aden.mi15@gmail.com"})
	return nil
}

func CheckUserApprover(DB *gorm.DB, email string) (Approvers, error) {
	var approver Approvers
	result := DB.Find(&approver, "email = ?", email)
	return approver, result.Error
}

func GetRandomApprover(DB *gorm.DB, ApproverType string) (*Approvers, error) {
	var approver *Approvers
	result := DB.First(&approver, "approver_type = ?", ApproverType)
	return approver, result.Error
}
