package db

import (
	"github.com/google/uuid"
	"github.com/palladiumkenya/individual-data-request-backend/internal/config"
	"github.com/palladiumkenya/individual-data-request-backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

// var DB *gorm.DB
var DB *gorm.DB

func Connect() (*gorm.DB, error) {
	cfg := config.LoadConfig()

	DB, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	//return db, nil
	return DB, nil
}

func MigrateDB() (*gorm.DB, error) {
	cfg := config.LoadConfig()

	//=========== migrate db
	DB, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Migrate the schema
	err = DB.AutoMigrate(&models.Requesters{}, &models.Requests{}, &models.Assignees{}, &models.Approvals{}, &models.Approvers{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Create
	DB.Create(&models.Requests{Summery: "need this moh request sorted asap",
		Status: "pending", Date_Due: time.Date(2024, 10, 10, 0, 0, 0, 0, time.UTC),
		Priority_level: "high", Requestor_id: uuid.MustParse("88f75fd1-67b7-411c-8c9e-311afd5cf1eb"),
		Assignee_id: uuid.MustParse("00000000-0000-0000-0000-000000000000"), Created_Date: time.Date(2024, 8, 10, 0, 0, 0, 0, time.UTC)})
	// ==========migrate db

	//return db, nil
	return DB, nil
}
