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

func GetApprovers(DB *gorm.DB) ([]Approvers, error) {
	var approverss []Approvers
	result := DB.Find(&approverss)
	return approverss, result.Error
}

func GetAllExternalApprovers(DB *gorm.DB) ([]string, error) {
	var approvers []string
	//result := DB.Find(&approverss, "approver_type = ?", "external")
	result := DB.Model(&Approvers{}).Select("email").Where("approver_type = ?", "external").Scan(&approvers)
	return approvers, result.Error
}

func GetApproversByEmail(DB *gorm.DB, email string) ([]Approvers, error) {
	var approvers []Approvers
	result := DB.Find(&approvers, "email = ?", email)
	//result := DB.Model(&Approvers{}).Select("email").Scan(&approvers)
	return approvers, result.Error
}

func CreateApprover(DB *gorm.DB, approver Approvers) (uuid.UUID, error) {
	if err := DB.Create(&approver).Error; err != nil {
		return uuid.UUID{}, err // Return the error and a zero UUID
	}
	return approver.ID, nil
}

func DeleteApprover(DB *gorm.DB, id uuid.UUID) error {
	result := DB.Delete(&Approvers{}, "id = ?", id)
	return result.Error
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
