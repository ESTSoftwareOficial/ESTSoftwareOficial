package controllers

import (
	"net/http"
	"strconv"

	"estsoftwareoficial/src/resource_types/application"
	"estsoftwareoficial/src/resource_types/domain/dto"

	"github.com/gin-gonic/gin"
)

type GetResourceTypeByIdController struct {
	getResourceTypeById *application.GetResourceTypeById
}

func NewGetResourceTypeByIdController(getResourceTypeById *application.GetResourceTypeById) *GetResourceTypeByIdController {
	return &GetResourceTypeByIdController{getResourceTypeById: getResourceTypeById}
}

func (grt *GetResourceTypeByIdController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	resourceType, err := grt.getResourceTypeById.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if resourceType == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tipo de recurso no encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"resourceType": dto.ResourceTypeResponse{
			ID:      resourceType.ID,
			Name:    resourceType.Name,
			IconURL: resourceType.IconURL,
		},
	})
}
