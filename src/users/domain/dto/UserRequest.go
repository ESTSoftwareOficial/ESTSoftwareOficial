package dto

type UserRequest struct {
	FirstName      string  `json:"firstName" binding:"required"`
	SecondName     string  `json:"secondName"`
	LastName       string  `json:"lastName" binding:"required"`
	SecondLastName string  `json:"secondLastName"`
	Email          string  `json:"email" binding:"required,email"`
	SecondaryEmail *string `json:"secondaryEmail,omitempty"`
	Password       string  `json:"password" binding:"required,min=6"`
	RoleID         int     `json:"roleId" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
