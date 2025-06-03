package config

import (
	"log"
	"os"
	"time"

	"github.com/m-shahjalal/onepolicy-api/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	if dsn := os.Getenv("DB_URL"); dsn == "" {
		log.Fatal("DB_URL environment variable is not set")
	}

	// Configure GORM
	config := &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Silent), // Reduce logging overhead
		PrepareStmt: true,                                  // Enable prepared statement cache
	}

	// Configure connection pool
	dbConfig := postgres.Config{
		DSN:                  os.Getenv("DB_URL"),
		PreferSimpleProtocol: true, // Use simple protocol for better performance
	}

	if DB, err = gorm.Open(postgres.New(dbConfig), config); err != nil {
		log.Fatal("Failed to connect database: ", err)
	}

	// Configure connection pool settings
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get database instance: ", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(time.Minute * 5)

	if err = EnableExtensions(DB); err != nil {
		log.Fatal("Failed to enable extentions")
	}

	if err = InitMigration(DB); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	println("Database is successfully initalized 🎉")
}

func EnableExtensions(db *gorm.DB) error {
	return db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error
}

func InitMigration(DB *gorm.DB) error {
	// Create enum type if not exists
	DB.Exec(`DO $$ 
	BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'policy_type') THEN
			CREATE TYPE policy_type AS ENUM ('cookie', 'terms', 'privacy');
		END IF;
	END $$;`)

	// Initialize migrator with optimized settings
	migrator := NewSmartMigrator(DB)

	// Batch migrate all models at once
	err := migrator.MigrateIfChanged(&model.User{}, &model.Policy{})
	if err != nil {
		panic(err)
	}

	return err
}
