package controllers

import (
	"estsoftwareoficial/src/users/application"
	"estsoftwareoficial/src/users/domain/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetUserByIdController struct {
	getUserById *application.GetUserById
}

func NewGetUserByIdController(getUserById *application.GetUserById) *GetUserByIdController {
	return &GetUserByIdController{getUserById: getUserById}
}

func (gc *GetUserByIdController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	user, err := gc.getUserById.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": dto.UserResponse{
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
		},
	})
}
