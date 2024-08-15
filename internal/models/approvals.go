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
	Request_id    uuid.UUID  `gorm:"type:uuid"`
	Request       Requests   `gorm:"foreignKey:Request_id"`
	Approver_id   uuid.UUID  `gorm:"type:uuid"`
	Approver      Approvers  `gorm:"foreignKey:Approver_id"`
	Approval_Date time.Time  `gorm:"type:date"`
}

func GetApprovalByID(DB *gorm.DB, Id uuid.UUID) (*Approvals, error) {
	var approval *Approvals
	result := DB.Preload("Requester").Preload("Request").Preload("Approver").First(&approval, "request_id = ?", Id)
	return approval, result.Error
}

func GetApprovals(DB *gorm.DB) ([]Approvals, error) {
	var approvals []Approvals
	result := DB.Find(&approvals)
	return approvals, result.Error
}

func CreateApproval(DB *gorm.DB, approvalData *Approvals) (*Approvals, error) {
	var approval *Approvals

	result := DB.Create(&Approvals{Comments: approvalData.Comments, Approver_type: approvalData.Approver_type, Approved: approvalData.Approved,
		Requestor_id: approvalData.Requestor_id, Request_id: approvalData.Request_id, Approval_Date: approvalData.Approval_Date})

	var request *Requests
	DB.First(&request, "id = ?", approvalData.Request_id)

	if approvalData.Approved == true {
		request.Status = "approved"
	} else if approvalData.Approved == false {
		request.Status = "rejected"
	}
	DB.Save(&request)
	return approval, result.Error
}
