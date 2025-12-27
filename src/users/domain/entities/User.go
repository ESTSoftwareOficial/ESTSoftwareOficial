package entities

import "time"

type User struct {
	ID               int       `json:"id"`
	FirstName        string    `json:"firstName"`
	SecondName       string    `json:"secondName"`
	LastName         string    `json:"lastName"`
	SecondLastName   string    `json:"secondLastName"`
	Email            string    `json:"email"`
	SecondaryEmail   *string   `json:"secondaryEmail,omitempty"`
	Password         *string   `json:"password,omitempty"`
	ProfilePhoto     *string   `json:"profilePhoto,omitempty"`
	RegistrationDate time.Time `json:"registrationDate"`
	RoleID           int       `json:"roleId"`
	OAuthProvider    *string   `json:"oauthProvider,omitempty"`
	OAuthID          *string   `json:"oauthId,omitempty"`
}
