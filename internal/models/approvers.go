package models

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Approvers struct {
	ID            uuid.UUID `json:"Id"`
	Email         string    `json:"email"`
	Approver_Type string    `json:"approver_type"`
}

func GetApproverByID(ctx context.Context, pool *pgxpool.Pool, ID int) (*Approvers, error) {
	var approver Approvers
	row := pool.QueryRow(ctx, "SELECT Id, email,approver_type FROM approvers WHERE id=$1", ID)
	err := row.Scan(&approver.ID, &approver.Email, &approver.Approver_Type)
	if err != nil {
		return nil, err
	}
	return &approver, nil
}

func GetApprovers(ctx context.Context, pool *pgxpool.Pool) ([]Approvers, error) {
	rows, err := pool.Query(ctx, "SELECT \"Id\", email,approver_type FROM approvers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var approvers []Approvers
	for rows.Next() {
		var approver Approvers
		err := rows.Scan(&approver.ID, &approver.Email, &approver.Approver_Type)
		if err != nil {
			return nil, err
		}
		approvers = append(approvers, approver)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return approvers, nil
}

func CreateApprover(ctx context.Context, pool *pgxpool.Pool, approver *Approvers) error {
	_, err := pool.Exec(ctx, "INSERT INTO Approvers (email, approver_type) VALUES ($1, $2)",
		approver.Email, approver.Approver_Type)
	if err != nil {
		return err
	}
	return nil
}
