package seed

import (
	"gorm.io/gorm"
	"log"
	"os"
)

func SeedData(db *gorm.DB) error {
	if os.Getenv("SEED_DATA") != "true" {
		log.Println("Skipping seed data (SEED_DATA != true)")
		return nil
	}

	// Seed users first
	userSeeder := NewUserSeeder(db)
	if err := userSeeder.SeedUsers(); err != nil {
		log.Printf("Failed to seed users: %v", err)
		return err
	}

	// Then seed policies
	policySeeder := NewPolicySeeder(db)
	if err := policySeeder.SeedPolicies(); err != nil {
		log.Printf("Failed to seed policies: %v", err)
		return err
	}

	println("Database is successfully seeded 🎉")
	return nil
}
