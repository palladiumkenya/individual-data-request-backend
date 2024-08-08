package models

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
	"gorm.io/gorm"

	//"github.com/palladiumkenya/individual-data-request-backend/internal/db"
	"time"
)

type Requests struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Summery        string    `gorm:"size:100;unique;not null"`
	Status         string    `gorm:"size:100;unique;not null"`
	Date_Due       time.Time `gorm:"type:date"`
	Priority_level string    `gorm:"size:100;unique;not null"`
	Requestor_id   uuid.UUID `gorm:"type:uuid"`
	Assignee_id    uuid.UUID `gorm:"type:uuid"`
	Created_Date   time.Time `gorm:"type:date"`
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

func GetRequests(DB *gorm.DB) ([]Requests, error) {
	var requests []Requests
	result := DB.Find(&requests)
	return requests, result.Error
}

//func GetRequests(ctx context.Context, pool *pgxpool.Pool) ([]Requests, error) {
//	rows, err := pool.Query(ctx, "SELECT \"Id\", summery, status, date_due, priority_level,requestor_id, created_date FROM requests")
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	var requests []Requests
//	for rows.Next() {
//		var request Requests
//		err := rows.Scan(&request.ID, &request.Summery, &request.Status, &request.Date_Due, &request.Priority_level, &request.Requestor_id, &request.Created_Date)
//		if err != nil {
//			return nil, err
//		}
//
//		requests = append(requests, request)
//	}
//
//	if rows.Err() != nil {
//		return nil, rows.Err()
//	}
//
//	return requests, nil
//}

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
