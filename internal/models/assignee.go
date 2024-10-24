package models

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"gorm.io/gorm"
)

type Assignees struct {
	ID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Email string    `gorm:"size:100;not null"`
}

func GetAssigneeByID(DB *gorm.DB, Id uuid.UUID) (*Assignees, error) {
	var assignee *Assignees
	result := DB.First(&assignee, "id = ?", Id)
	return assignee, result.Error
}

func GetAssignees(DB *gorm.DB) ([]Assignees, error) {
	var assignees []Assignees
	result := DB.Find(&assignees)
	return assignees, result.Error
}

func CreateAssignee(ctx context.Context, pool *pgxpool.Pool, assignee *Assignees) error {
	_, err := pool.Exec(ctx, "INSERT INTO Assignees (email) VALUES ($1)",
		assignee.Email)
	if err != nil {
		return err
	}
	return nil
}

func CheckUserAnalyst(DB *gorm.DB, email string) (Assignees, error) {
	var analyst Assignees
	result := DB.Find(&analyst, "email = ?", email)
	if result.RowsAffected == 0 {
		return analyst, nil
	}
	return analyst, result.Error
}

func GetAnalysts(DB *gorm.DB) ([]Assignees, error) {
	var assigneesAvailable []Assignees
	result := DB.Find(&assigneesAvailable)
	return assigneesAvailable, result.Error
}
