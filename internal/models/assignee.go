package models

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Assignees struct {
	ID    uuid.UUID `json:"Id"`
	Email string    `json:"email"`
}

func GetAssigneeByID(ctx context.Context, pool *pgxpool.Pool, ID int) (*Assignees, error) {
	var assignee Assignees
	row := pool.QueryRow(ctx, "SELECT Id, email FROM assignees WHERE id=$1", ID)
	err := row.Scan(&assignee.ID, &assignee.Email)
	if err != nil {
		return nil, err
	}
	return &assignee, nil
}

func GetAssignees(ctx context.Context, pool *pgxpool.Pool) ([]Assignees, error) {
	rows, err := pool.Query(ctx, "SELECT \"Id\", email FROM assignees")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var assignees []Assignees
	for rows.Next() {
		var assignee Assignees
		err := rows.Scan(&assignee.ID, &assignee.Email)
		if err != nil {
			return nil, err
		}
		assignees = append(assignees, assignee)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return assignees, nil
}

func CreateAssignee(ctx context.Context, pool *pgxpool.Pool, assignee *Assignees) error {
	_, err := pool.Exec(ctx, "INSERT INTO Assignees (email) VALUES ($1)",
		assignee.Email)
	if err != nil {
		return err
	}
	return nil
}
