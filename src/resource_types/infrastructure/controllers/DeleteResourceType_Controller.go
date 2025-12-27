package controllers

import (
	"net/http"
	"strconv"

	"estsoftwareoficial/src/resource_types/application"

	"github.com/gin-gonic/gin"
)

type DeleteResourceTypeController struct {
	deleteResourceType *application.DeleteResourceType
}

func NewDeleteResourceTypeController(deleteResourceType *application.DeleteResourceType) *DeleteResourceTypeController {
	return &DeleteResourceTypeController{deleteResourceType: deleteResourceType}
}

func (drt *DeleteResourceTypeController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = drt.deleteResourceType.Execute(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tipo de recurso eliminado exitosamente",
	})
}
