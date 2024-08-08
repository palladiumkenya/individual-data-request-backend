package models

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Requesters struct {
	ID    uuid.UUID `json:"Id"`
	Email string    `json:"email"`
}

func GetRequesterByID(ctx context.Context, pool *pgxpool.Pool, ID int) (*Requesters, error) {
	var requester Requesters
	row := pool.QueryRow(ctx, "SELECT Id, email FROM requesters WHERE id=$1", ID)
	err := row.Scan(&requester.ID, &requester.Email)
	if err != nil {
		return nil, err
	}
	return &requester, nil
}

func GetRequesters(ctx context.Context, pool *pgxpool.Pool) ([]Requesters, error) {
	rows, err := pool.Query(ctx, "SELECT \"Id\", email FROM requesters")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requesters []Requesters
	for rows.Next() {
		var requester Requesters
		err := rows.Scan(&requester.ID, &requester.Email)
		if err != nil {
			return nil, err
		}
		requesters = append(requesters, requester)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return requesters, nil
}

func CreateRequester(ctx context.Context, pool *pgxpool.Pool, requester *Requesters) error {
	_, err := pool.Exec(ctx, "INSERT INTO Requesters (email) VALUES ($1)",
		requester.Email)
	if err != nil {
		return err
	}
	return nil
}
