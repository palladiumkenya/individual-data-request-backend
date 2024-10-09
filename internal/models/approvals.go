package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"time"
)

type Approvals struct {
	ID            uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Comments      string     `gorm:"size:500;null"`
	Approver_type string     `gorm:"size:100;not null"`
	Approved      *bool      `gorm:"bool;default:null"`
	Requestor_id  uuid.UUID  `gorm:"type:uuid"`
	Requester     Requesters `gorm:"foreignKey:Requestor_id"`
	Request_id    uuid.UUID  `gorm:"type:uuid"`
	Request       Requests   `gorm:"foreignKey:Request_id"`
	Approver_id   uuid.UUID  `gorm:"type:uuid"`
	Approver      Approvers  `gorm:"foreignKey:Approver_id"`
	Approval_Date time.Time  `gorm:"type:date"`
}

func GetApprovalsByType(DB *gorm.DB, ApproveType string) ([]Approvals, error) {
	var approvals []Approvals
	result := DB.Preload("Requester").Preload("Request").Preload("Approver").Find(&approvals, "Approver_type = ?", ApproveType)
	return approvals, result.Error
}

func GetApprovalByID(DB *gorm.DB, Id uuid.UUID) (*Approvals, error) {
	var approval *Approvals
	result := DB.Preload("Requester").Preload("Request").Preload("Approver").First(&approval, "request_id = ?", Id)
	return approval, result.Error
}

func GetApprovalByIDAndType(DB *gorm.DB, Id uuid.UUID, approvalType string) (*Approvals, error) {
	var approval *Approvals
	result := DB.Preload("Requester").Preload("Request").Preload("Approver").First(&approval, "request_id = ? and approver_type=?", Id, approvalType)
	return approval, result.Error
}

type Result struct {
	PriorityLevel string
	Count         int64
}

func GetApprovalsCounts(DB *gorm.DB, approvalType string) ([]Result, error) {
	var results []Result
	result := DB.Model(&Approvals{}).
		Select("requests.priority_level, COUNT(*) as count").
		Joins("LEFT JOIN requests ON approvals.request_id = requests.id").
		Where("approvals.approver_type = ?", approvalType).
		Group("requests.priority_level").
		Scan(&results)

	return results, result.Error
}

func GetApprovals(DB *gorm.DB) ([]Approvals, error) {
	var approvals []Approvals
	result := DB.Find(&approvals)
	return approvals, result.Error
}

func CreateApproval(DB *gorm.DB, approvalData *Approvals) (*Approvals, error) {
	var approver *Approvers
	DB.First(&approver, "approver_type = ?", approvalData.Approver_type)

	var approval *Approvals

	result := DB.Create(&Approvals{Comments: approvalData.Comments, Approver_type: approvalData.Approver_type,
		Approved: approvalData.Approved, Requestor_id: approvalData.Requestor_id,
		Request_id: approvalData.Request_id, Approval_Date: time.Now(), Approver_id: approver.ID})

	var request Requests
	DB.First(&request, "id = ?", approvalData.Request_id)

	if isApproved(approvalData.Approved) && approvalData.Approver_type == "internal" {
		// Update the status
		request.Status = "review stage"
		if err := DB.Save(&request).Error; err != nil {
			log.Fatalf("Error updating request status: %v\n", err)
		}

	} else if isApproved(approvalData.Approved) && approvalData.Approver_type == "external" {
		// Update the status
		request.Status = "approved"
		if err := DB.Save(&request).Error; err != nil {
			log.Fatalf("Error updating request status: %v\n", err)
		}
	} else {
		//"rejected
		if err := DB.Model(&request).Update("status", "rejected").Error; err != nil {
		}
	}
	//DB.Save(&request)
	return approval, result.Error
}

func isApproved(approved *bool) bool {
	return approved != nil && *approved
}
