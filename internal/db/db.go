package db

import (
    "github.com/palladiumkenya/individual-data-request-backend/internal/config"
    "github.com/palladiumkenya/individual-data-request-backend/internal/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
)

var DB *gorm.DB

func Connect() (*gorm.DB, error) {
    cfg := config.LoadConfig()

    log.Printf("Connecting to database with URL: %s", cfg.DatabaseURL)

    DB, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }

    return DB, nil
}

func MigrateDB() (*gorm.DB, error) {
    cfg := config.LoadConfig()

    log.Printf("Migrating database with URL: %s", cfg.DatabaseURL)

    DB, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }

    err = DB.AutoMigrate(&models.Requesters{}, &models.Requests{}, &models.Assignees{}, &models.Approvals{}, &models.Approvers{})
    if err != nil {
        log.Fatalf("failed to migrate database: %v", err)
    }

    return DB, nil
}
