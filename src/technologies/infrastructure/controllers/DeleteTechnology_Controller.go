package controllers

import (
	"net/http"
	"strconv"

	"estsoftwareoficial/src/technologies/application"

	"github.com/gin-gonic/gin"
)

type DeleteTechnologyController struct {
	deleteTechnology *application.DeleteTechnology
}

func NewDeleteTechnologyController(deleteTechnology *application.DeleteTechnology) *DeleteTechnologyController {
	return &DeleteTechnologyController{deleteTechnology: deleteTechnology}
}

func (dt *DeleteTechnologyController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = dt.deleteTechnology.Execute(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tecnología eliminada exitosamente",
	})
}
