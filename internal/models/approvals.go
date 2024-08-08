package models

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Approvals struct {
	ID            uuid.UUID   `json:"Id"`
	Comments      string      `json:"comments"`
	Approver_type string      `json:"approver_type"`
	Approved      bool        `json:"approved"`
	Requestor_id  uuid.UUID   `json:"requestor_id"`
	Assignee_id   uuid.UUID   `json:"assignee_id"`
	Approval_Date pgtype.Date `json:"approval_date"`
}

func GetApprovalByID(ctx context.Context, pool *pgxpool.Pool, ID int) (*Approvals, error) {
	var approval Approvals
	row := pool.QueryRow(ctx, "SELECT \"Id\", comments, approved, requestor_id,assignee_id FROM approvals WHERE id=$1", ID)
	err := row.Scan(&approval.ID, &approval.Comments, &approval.Approved, &approval.Requestor_id, &approval.Assignee_id)
	if err != nil {
		return nil, err
	}
	return &approval, nil
}

func GetApprovals(ctx context.Context, pool *pgxpool.Pool) ([]Approvals, error) {
	rows, err := pool.Query(ctx, "SELECT \"Id\", comments, approved, requestor_id,assignee_id FROM approvals")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var approvals []Approvals
	for rows.Next() {
		var approval Approvals
		err := rows.Scan(&approval.ID, &approval.Comments, &approval.Approved, &approval.Requestor_id, &approval.Assignee_id)
		if err != nil {
			return nil, err
		}
		approvals = append(approvals, approval)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return approvals, nil
}

func CreateApproval(ctx context.Context, pool *pgxpool.Pool, approval *Approvals) error {
	_, err := pool.Exec(ctx, "INSERT INTO Approvals (comments, approved, requestor_id, approver_id) VALUES ($1, $2, $3, $4)",
		approval.Comments, approval.Approved, approval.Requestor_id, approval.Assignee_id)
	if err != nil {
		return err
	}
	return nil
}
