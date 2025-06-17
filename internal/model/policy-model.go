package model

import (
	"time"

	"github.com/google/uuid"
)

type Policy struct {
	ModelHeader
	Title       string    `gorm:"type:text;not null;default:'Untitled Policy'" json:"title"`
	UserID      uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	User        User      `gorm:"foreignKey:UserID" json:"user"`
	Status      string    `gorm:"type:status_type;default:'draft'" json:"status"`
	Inputs      string    `gorm:"type:text" json:"inputs"`
	Effect_date time.Time `gorm:"default:NOW()" json:"effect_date"`
	Markdown    string    `gorm:"type:text" json:"markdown"`
	Policy_type string    `gorm:"type:policy_type;not null" json:"policy_type"`
	View_count  int       `gorm:"default:0" json:"view_count"`
	ModelFooter
}
