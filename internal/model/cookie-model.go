package model

import (
	"database/sql/driver"
	"time"

	"gorm.io/gorm"
)

type Cookie struct {
	*gorm.Model
	Inputs      string
	Effect_date *time.Time
	Markdown    string
	Policy_type PolicyType `gorm:"type:policy_type`
}

type PolicyType string

const (
	CookieType  PolicyType = "cookie"
	TermType    PolicyType = "terms"
	PrivacyType PolicyType = "privacy"
)

func (p *PolicyType) Scan(value any) error {
	*p = PolicyType(value.([]byte))
	return nil
}
func (p PolicyType) Value() (driver.Value, error) {
	return string(p), nil
}
