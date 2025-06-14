package model

type User struct {
	ModelHeader
	Name     string
	Email    string
	Password string
	ModelFooter
}
