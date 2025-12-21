package dto

import "time"

type UserResponse struct {
	ID               int       `json:"id"`
	FirstName        string    `json:"firstName"`
	SecondName       string    `json:"secondName"`
	LastName         string    `json:"lastName"`
	SecondLastName   string    `json:"secondLastName"`
	Email            string    `json:"email"`
	SecondaryEmail   *string   `json:"secondaryEmail,omitempty"`
	RegistrationDate time.Time `json:"registrationDate"`
	RoleID           int       `json:"roleId"`
	OAuthProvider    *string   `json:"oauthProvider,omitempty"`
}

type LoginResponse struct {
	Message string       `json:"message"`
	User    UserResponse `json:"user"`
}
