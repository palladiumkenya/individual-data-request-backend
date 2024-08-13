package models

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"gorm.io/gorm"
)

type Approvers struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Email         string    `gorm:"size:100;unique;not null"`
	Approver_Type string    `gorm:"size:100;unique;not null"`
}

func GetApproversByID(DB *gorm.DB, Id uuid.UUID) (*Approvers, error) {
	var approvers *Approvers
	result := DB.First(&approvers, "id = ?", Id)
	return approvers, result.Error
}

func GetApproverss(DB *gorm.DB) ([]Approvers, error) {
	var approverss []Approvers
	result := DB.Find(&approverss)
	return approverss, result.Error
}

func CreateApprover(ctx context.Context, pool *pgxpool.Pool, approver *Approvers) error {
	_, err := pool.Exec(ctx, "INSERT INTO Approvers (email, approver_type) VALUES ($1, $2)",
		approver.Email, approver.Approver_Type)
	if err != nil {
		return err
	}
	return nil
}
