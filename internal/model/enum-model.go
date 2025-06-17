package model

import (
	"database/sql/driver"
)

type Enum string

func (e *Enum) Scan(value any) error {
	if value == nil {
		*e = ""
		return nil
	}
	switch v := value.(type) {
	case string:
		*e = Enum(v)
	case []byte:
		*e = Enum(v)
	}
	return nil
}

func (e Enum) Value() (driver.Value, error) {
	return string(e), nil
}

// PolicyType enum
type PolicyType Enum

const (
	Cookie  PolicyType = "cookie"
	Term    PolicyType = "terms"
	Privacy PolicyType = "privacy"
)

type PolicyStatus Enum

const (
	PolicyStatusInactive  PolicyStatus = "inactive"
	PolicyStatusDraft     PolicyStatus = "draft"
	PolicyStatusPublished PolicyStatus = "published"
)
