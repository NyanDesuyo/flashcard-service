package config

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var MainDB *gorm.DB

func ConnectMainPostgres() {
	var err error

	dsn := os.Getenv("DB_CONNECTION")
	if dsn == "" {
		log.Fatal("DB_CONNECTION environment variable not set")
	}

	logLevel := logger.Info

	if os.Getenv("GIN_MODE") == "release" {
		logLevel = logger.Error
	}

	config := &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	}

	MainDB, err = gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, err := MainDB.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Successfully connected to PostgresSQL Database")
}
