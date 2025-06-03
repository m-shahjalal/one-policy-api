package model

import (
	"database/sql/driver"
	"time"
)

type Policy struct {
	ModelHeader
	Inputs      string
	Effect_date time.Time `gorm:"default:NOW()"`
	Markdown    string
	Policy_type PolicyType `gorm:"type:policy_type;check:policy_type IN ('cookie', 'terms', 'privacy')"`
	ModelFooter
}

type PolicyType string

const (
	Cookie  PolicyType = "cookie"
	Term    PolicyType = "terms"
	Privacy PolicyType = "privacy"
)

func (p *PolicyType) Scan(value any) error {
	if value == nil {
		*p = ""
		return nil
	}
	switch v := value.(type) {
	case string:
		*p = PolicyType(v)
	case []byte:
		*p = PolicyType(v)
	}
	return nil
}

func (p PolicyType) Value() (driver.Value, error) {
	return string(p), nil
}
