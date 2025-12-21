package controllers

import (
	"estsoftwareoficial/src/users/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetTotalUsersController struct {
	getTotalUsers *application.GetTotalUsers
}

func NewGetTotalUsersController(getTotalUsers *application.GetTotalUsers) *GetTotalUsersController {
	return &GetTotalUsersController{getTotalUsers: getTotalUsers}
}

func (gc *GetTotalUsersController) Execute(c *gin.Context) {
	total, err := gc.getTotalUsers.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total": total,
	})
}
