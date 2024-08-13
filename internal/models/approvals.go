package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Approvals struct {
	ID            uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Comments      string     `gorm:"size:500;not null"`
	Approver_type string     `gorm:"size:100;not null"`
	Approved      bool       `gorm:"bool"`
	Requestor_id  uuid.UUID  `gorm:"type:uuid"`
	Requester     Requesters `gorm:"foreignKey:Requestor_id"`
	Approval_Date time.Time  `gorm:"type:date"`
}

func GetApprovalByID(DB *gorm.DB, Id uuid.UUID) (*Approvals, error) {
	var approval *Approvals
	result := DB.First(&approval, "id = ?", Id)
	return approval, result.Error
}

func GetApprovals(DB *gorm.DB) ([]Approvals, error) {
	var approvals []Approvals
	result := DB.Find(&approvals)
	return approvals, result.Error
}

func CreateApproval(DB *gorm.DB, request *Approvals) (*Approvals, error) {
	var approval *Approvals

	result := DB.Create(&Approvals{Comments: request.Comments, Approver_type: request.Approver_type, Approved: request.Approved,
		Requestor_id: request.Requestor_id, Approval_Date: request.Approval_Date})
	return approval, result.Error
}
