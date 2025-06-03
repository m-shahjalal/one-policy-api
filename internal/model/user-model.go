package model

type User struct {
	ModelHeader
	First_name    string
	Last_name     string
	Email         string
	Password      string
	Token         string
	Refresh_token string
	ModelFooter
}
