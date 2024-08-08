package db

import (
	"github.com/google/uuid"
	"github.com/palladiumkenya/individual-data-request-backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

//var DB *gorm.DB

func Connect(connString string) (*gorm.DB, error) {
	//config, err := pgxpool.ParseConfig(connString)
	//if err != nil {
	//	log.Fatalf("Unable to parse connection string: %v\n", err)
	//	return nil, err
	//}

	//pool, err := pgxpool.ConnectConfig(context.Background(), config)
	//if err != nil {
	//	log.Fatalf("Unable to connect to database: %v\n", err)
	//	return nil, err
	//}

	//=========== migrate db
	//dsn := "host=localhost user=yourusername password=yourpassword dbname=yourdbname port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Migrate the schema
	err = DB.AutoMigrate(&models.Requests{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Create
	DB.Create(&models.Requests{Summery: "need this moh request sorted asap",
		Status: "pending", Date_Due: time.Date(1992, 10, 10, 0, 0, 0, 0, time.UTC),
		Priority_level: "high", Requestor_id: uuid.MustParse("afd58c35-1658-4566-b62f-cce87fc850cb"),
		Assignee_id: uuid.MustParse("00000000-0000-0000-0000-000000000000"), Created_Date: time.Date(1992, 10, 10, 0, 0, 0, 0, time.UTC)})
	// ==========migrate db

	//return db, nil
	return DB, nil
}
