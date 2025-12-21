package controllers

import (
	"estsoftwareoficial/src/users/application"
	"estsoftwareoficial/src/users/domain/dto"
	"estsoftwareoficial/src/users/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserController struct {
	authService *application.AuthService
}

func NewCreateUserController(authService *application.AuthService) *CreateUserController {
	return &CreateUserController{authService: authService}
}

func (uc *CreateUserController) Execute(c *gin.Context) {
	var req dto.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &entities.User{
		FirstName:      req.FirstName,
		SecondName:     req.SecondName,
		LastName:       req.LastName,
		SecondLastName: req.SecondLastName,
		Email:          req.Email,
		SecondaryEmail: req.SecondaryEmail,
		Password:       &req.Password,
		RoleID:         req.RoleID,
	}

	savedUser, err := uc.authService.Register(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario creado exitosamente",
		"user": dto.UserResponse{
			ID:               savedUser.ID,
			FirstName:        savedUser.FirstName,
			SecondName:       savedUser.SecondName,
			LastName:         savedUser.LastName,
			SecondLastName:   savedUser.SecondLastName,
			Email:            savedUser.Email,
			SecondaryEmail:   savedUser.SecondaryEmail,
			RegistrationDate: savedUser.RegistrationDate,
			RoleID:           savedUser.RoleID,
		},
	})
}
