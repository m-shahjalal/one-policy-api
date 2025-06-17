package seed

import (
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/m-shahjalal/onepolicy-api/internal/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserSeeder struct {
	db *gorm.DB
}

func NewUserSeeder(db *gorm.DB) *UserSeeder {
	return &UserSeeder{db: db}
}

func (s *UserSeeder) SeedUsers() error {
	var count int64
	s.db.Model(&model.User{}).Count(&count)

	if count > 0 {
		log.Println("Sample users already exist, skipping user seeding")
		return nil
	}

	users := s.generateSampleUsers(12)
	println("Creating sample users...", len(users), "users")
	if err := s.db.Create(&users).Error; err != nil {
		println("Failed to create user")
		return err
	}

	println("Created users successfully")
	return nil
}

func (s *UserSeeder) generateSampleUsers(count int) []model.User {
	passBytes, err := bcrypt.GenerateFromPassword([]byte("Pass@123"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	pass := string(passBytes)

	firstNames := []string{
		"Alex", "Sarah", "Michael", "Emma", "David", "Lisa", "James", "Maria",
		"Robert", "Jennifer", "William", "Elizabeth", "John", "Anna", "Daniel",
		"Jessica", "Matthew", "Ashley", "Christopher", "Amanda", "Andrew", "Nicole",
		"Joshua", "Samantha", "Ryan", "Stephanie", "Brandon", "Rachel", "Kevin",
		"Lauren", "Tyler", "Hannah", "Jordan", "Megan", "Justin", "Brittany",
	}

	lastNames := []string{
		"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller",
		"Davis", "Rodriguez", "Martinez", "Hernandez", "Lopez", "Gonzalez",
		"Wilson", "Anderson", "Thomas", "Taylor", "Moore", "Jackson", "Martin",
		"Lee", "Perez", "Thompson", "White", "Harris", "Sanchez", "Clark",
		"Price", "Alvarez", "Castillo", "Sanders", "Patel", "Myers", "Long",
	}

	users := []model.User{
		{Name: "System Administrator", Email: "system@onepolicy.com", Password: pass},
		{Name: "System Manager", Email: "admin@onepolicy.com", Password: pass},
		{Name: "Developer", Email: "dev@mail.com", Password: pass},
	}

	for range count {
		firstName := firstNames[rand.Intn(len(firstNames))]
		lastName := lastNames[rand.Intn(len(lastNames))]
		domain := "mail.com"
		var email string

		emailFormats := []string{"%s.%s@%s", "%s_%s@%s", "%s%s@%s", "%s.%s%d@%s"}
		format := emailFormats[rand.Intn(len(emailFormats))]

		if format == "%s.%s%d@%s" {
			email = fmt.Sprintf(format,
				strings.ToLower(firstName),
				strings.ToLower(lastName),
				rand.Intn(99)+1,
				domain,
			)
		} else {
			email = fmt.Sprintf(format,
				strings.ToLower(firstName),
				strings.ToLower(lastName),
				domain,
			)
		}

		user := model.User{
			Name:     firstName + " " + lastName,
			Email:    email,
			Password: pass,
		}

		users = append(users, user)
	}

	return users
}
