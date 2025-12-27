package controllers

import (
	"net/http"

	"estsoftwareoficial/src/resource_types/application"
	"estsoftwareoficial/src/resource_types/domain/dto"
	"estsoftwareoficial/src/resource_types/domain/entities"

	"github.com/gin-gonic/gin"
)

type CreateResourceTypeController struct {
	createResourceType *application.CreateResourceType
}

func NewCreateResourceTypeController(createResourceType *application.CreateResourceType) *CreateResourceTypeController {
	return &CreateResourceTypeController{createResourceType: createResourceType}
}

func (crt *CreateResourceTypeController) Execute(c *gin.Context) {
	var req dto.ResourceTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resourceType := &entities.ResourceType{
		Name:    req.Name,
		IconURL: req.IconURL,
	}

	savedResourceType, err := crt.createResourceType.Execute(resourceType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Tipo de recurso creado exitosamente",
		"resourceType": dto.ResourceTypeResponse{
			ID:      savedResourceType.ID,
			Name:    savedResourceType.Name,
			IconURL: savedResourceType.IconURL,
		},
	})
}
