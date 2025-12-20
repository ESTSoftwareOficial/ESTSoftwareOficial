package entities

import "time"

type User struct {
	ID               int       `json:"id"`
	FirstName        string    `json:"firstName"`
	SecondName       string    `json:"secondName"`
	LastName         string    `json:"lastName"`
	SecondLastName   string    `json:"secondLastName"`
	Email            string    `json:"email"`
	Password         string    `json:"password"`
	RegistrationDate time.Time `json:"registrationDate"`
	RoleID           int       `json:"roleId"`
	OAuthProvider    string    `json:"oauthProvider"`
	OAuthID          string    `json:"oauthId"`
}
