package controllers

import (
	"net/http"
	"strconv"

	"estsoftwareoficial/src/resource_types/application"
	"estsoftwareoficial/src/resource_types/domain/dto"
	"estsoftwareoficial/src/resource_types/domain/entities"

	"github.com/gin-gonic/gin"
)

type UpdateResourceTypeController struct {
	updateResourceType *application.UpdateResourceType
}

func NewUpdateResourceTypeController(updateResourceType *application.UpdateResourceType) *UpdateResourceTypeController {
	return &UpdateResourceTypeController{updateResourceType: updateResourceType}
}

func (urt *UpdateResourceTypeController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req dto.ResourceTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resourceType := &entities.ResourceType{
		ID:      id,
		Name:    req.Name,
		IconURL: req.IconURL,
	}

	err = urt.updateResourceType.Execute(resourceType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tipo de recurso actualizado exitosamente",
	})
}
