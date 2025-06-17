package seed

import (
	"fmt"
	"github.com/m-shahjalal/onepolicy-api/internal/model"
	"gorm.io/gorm"
	"log"
	"time"
)

type PolicySeeder struct {
	db *gorm.DB
}

func NewPolicySeeder(db *gorm.DB) *PolicySeeder {
	return &PolicySeeder{db: db}
}

func (s *PolicySeeder) SeedPolicies() error {
	var count int64
	s.db.Model(&model.Policy{}).Count(&count)
	if count > 0 {
		log.Println("Policies already exist, skipping seeding")
		return nil
	}

	var systemUser model.User
	if err := s.db.Where("email = ?", "system@onepolicy.com").First(&systemUser).Error; err != nil {
		return fmt.Errorf("system user not found: %v", err)
	}

	companies := []string{"TechFlow", "DataVault", "CloudSync", "SecureLink", "InnovateNow", "DigitalBridge", "SmartConnect"}
	policyTypes := []string{string(model.Privacy), string(model.Cookie), string(model.Term)}
	statuses := []string{string(model.PolicyStatusPublished), string(model.PolicyStatusDraft)}

	policyList := []model.Policy{}

	for i := range 20 {
		company := companies[i%len(companies)]
		policyType := policyTypes[i%len(policyTypes)]
		status := statuses[i%len(statuses)]

		policy := model.Policy{
			Title:       fmt.Sprintf("%s %s Policy v%d", company, policyType, (i/3)+1),
			UserID:      systemUser.ID,
			Status:      status,
			Policy_type: policyType,
			View_count:  i * 50,
			Effect_date: time.Now().AddDate(0, 0, -i*7),
			Markdown:    s.generateContent(company, policyType),
		}

		policyList = append(policyList, policy)
	}

	println("Seeding policies...", len(policyList), policyList)

	if err := s.db.Create(&policyList).Error; err != nil {
		return fmt.Errorf("failed to seed policies: %v", err)
	}

	log.Println("Successfully seeded 20 policies for system user")
	return nil
}

func (s *PolicySeeder) generateContent(company, policyType string) string {
	switch policyType {
	case string(model.Privacy):
		return fmt.Sprintf(`# %s Privacy Policy

## Data Collection
We collect personal information when you:
- Create an account
- Use our services
- Contact support

## Data Usage
Your information is used to:
- Provide our services
- Improve user experience
- Send important updates

## Your Rights
- Access your data
- Request deletion
- Update information

Contact: privacy@%s.com`, company, company)

	case string(model.Cookie):
		return fmt.Sprintf(`# %s Cookie Policy

## What Are Cookies
Cookies are small files stored on your device.

## Types We Use
- **Essential**: Required for site functionality
- **Analytics**: Help improve our service
- **Preferences**: Remember your settings

## Managing Cookies
You can control cookies in your browser settings.

Contact: cookies@%s.com`, company, company)

	case string(model.Term):
		return fmt.Sprintf(`# %s Terms of Service

## Acceptance
By using our service, you agree to these terms.

## User Responsibilities
- Provide accurate information
- Use service lawfully
- Respect other users

## Service Availability
We strive for 99.9%% uptime but cannot guarantee uninterrupted service.

## Termination
We may terminate accounts for violations of these terms.

Contact: legal@%s.com`, company, company)

	default:
		return fmt.Sprintf("# %s Policy\n\nPolicy content for %s", company, policyType)
	}
}
