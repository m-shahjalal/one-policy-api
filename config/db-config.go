package config

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func DatabaseConnection() {
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

	if err = enableExtensions(DB); err != nil {
		log.Fatal("Failed to enable extensions")
	}

	if err = InitMigration(DB); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	println("Database is successfully initialized 🎉")

}

func enableExtensions(db *gorm.DB) error {
	// Create UUID extension
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error; err != nil {
		return err
	}

	// Create status_type enum if it doesn't exist
	if err := db.Exec(`DO $$ 
		BEGIN 
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status_type') THEN
				CREATE TYPE status_type AS ENUM ('draft', 'published', 'archived');
			END IF;
		END $$;`).Error; err != nil {
		return err
	}

	// Create policy_type enum if it doesn't exist
	if err := db.Exec(`DO $$ 
		BEGIN 
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'policy_type') THEN
				CREATE TYPE policy_type AS ENUM ('privacy', 'cookie', 'terms');
			END IF;
		END $$;`).Error; err != nil {
		return err
	}

	return nil
}
