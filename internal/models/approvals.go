package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"time"
)

type Approvals struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Comments      string    `gorm:"size:100;unique;not null"`
	Approver_type string    `gorm:"size:100;unique;not null"`
	Approved      bool      `gorm:"bool"`
	Requestor_id  uuid.UUID `gorm:"type:uuid"`
	Assignee_id   uuid.UUID `gorm:"type:uuid"`
	Approval_Date time.Time `gorm:"type:date"`
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

//func CreateApproval(DB *gorm.DB, approval *Approvals) error {
//	_, err := pool.Exec(ctx, "INSERT INTO Approvals (comments, approved, requestor_id, approver_id) VALUES ($1, $2, $3, $4)",
//		approval.Comments, approval.Approved, approval.Requestor_id, approval.Assignee_id)
//	DB.Create(&Approvals{approval.Comments, approval.Approved, approval.Requestor_id, approval.Assignee_id})
//
//	if err != nil {
//		return err
//	}
//	return nil
//}
