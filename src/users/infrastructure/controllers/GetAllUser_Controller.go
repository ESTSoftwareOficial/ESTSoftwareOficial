package controllers

import (
	"estsoftwareoficial/src/users/application"
	"estsoftwareoficial/src/users/domain/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllUsersController struct {
	getAllUsers *application.GetAllUsers
}

func NewGetAllUsersController(getAllUsers *application.GetAllUsers) *GetAllUsersController {
	return &GetAllUsersController{getAllUsers: getAllUsers}
}

func (gc *GetAllUsersController) Execute(c *gin.Context) {
	users, err := gc.getAllUsers.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var userResponses []dto.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, dto.UserResponse{
			ID:               user.ID,
			FirstName:        user.FirstName,
			SecondName:       user.SecondName,
			LastName:         user.LastName,
			SecondLastName:   user.SecondLastName,
			Email:            user.Email,
			SecondaryEmail:   user.SecondaryEmail,
			ProfilePhoto:     user.ProfilePhoto,
			RegistrationDate: user.RegistrationDate,
			RoleID:           user.RoleID,
			OAuthProvider:    user.OAuthProvider,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"users": userResponses,
	})
}
