package models

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type Requests struct {
	ID             uuid.UUID  `json:"Id"`
	Summery        string     `json:"summery"`
	Status         string     `json:"status"`
	Date_Due       *time.Time `json:"due_date"`
	Priority_level string     `json:"priority_level"`
	Requestor_id   uuid.UUID  `json:"requestor_id"`
	Assignee_id    uuid.UUID  `json:"assignee_id"`
	Created_Date   *time.Time `json:"created_date"`
}

func GetRequestByID(ctx context.Context, pool *pgxpool.Pool, ID int) (*Requests, error) {
	var request Requests
	row := pool.QueryRow(ctx, "SELECT Id, summery, status, due_date, priority_level,requestor_id FROM requests WHERE id=$1", ID)
	err := row.Scan(&request.ID, &request.Summery, &request.Status, &request.Date_Due, &request.Priority_level, &request.Requestor_id)
	if err != nil {
		return nil, err
	}
	return &request, nil
}

//func GetRequests(ctx context.Context, pool *pgxpool.Pool) (*Requests, error) {
//	var request Requests
//	row := pool.QueryRow(ctx, "SELECT Id, summery, status, due_date, priority_level,requestor_id FROM requests ")
//	err := row.Scan(&request.ID, &request.Summery, &request.Status, &request.Date_Due, &request.Priority_level, &request.Requestor_id)
//	if err != nil {
//		return nil, err
//	}
//	return &request, nil
//}

func GetRequests(ctx context.Context, pool *pgxpool.Pool) ([]Requests, error) {
	rows, err := pool.Query(ctx, "SELECT \"Id\", summery, status, date_due, priority_level,requestor_id, created_date FROM requests")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requests []Requests
	for rows.Next() {
		var request Requests
		err := rows.Scan(&request.ID, &request.Summery, &request.Status, &request.Date_Due, &request.Priority_level, &request.Requestor_id, &request.Created_Date)
		if err != nil {
			return nil, err
		}

		requests = append(requests, request)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return requests, nil
}

func CreateRequest(ctx context.Context, pool *pgxpool.Pool, request *Requests) error {
	_, err := pool.Exec(ctx, "INSERT INTO Requests (summery, status, due_date, priority_level, requestor_id) VALUES ($1, $2, $3)",
		request.Summery, request.Status, request.Date_Due, request.Priority_level)
	if err != nil {
		return err
	}
	return nil
}

func convertPgDateToTime(pgDate pgtype.Date) *time.Time {
	if pgDate.Status == pgtype.Present {
		t := pgDate.Time
		return &t
	}
	return nil
}
