package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	First_name    string
	Last_name     string
	Email         string
	Password      string
	Token         string
	Refresh_token string
}
