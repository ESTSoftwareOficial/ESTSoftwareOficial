package controllers

import (
	"net/http"

	"estsoftwareoficial/src/resource_types/application"
	"estsoftwareoficial/src/resource_types/domain/dto"

	"github.com/gin-gonic/gin"
)

type GetAllResourceTypesController struct {
	getAllResourceTypes *application.GetAllResourceTypes
}

func NewGetAllResourceTypesController(getAllResourceTypes *application.GetAllResourceTypes) *GetAllResourceTypesController {
	return &GetAllResourceTypesController{getAllResourceTypes: getAllResourceTypes}
}

func (grt *GetAllResourceTypesController) Execute(c *gin.Context) {
	resourceTypes, err := grt.getAllResourceTypes.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resourceTypeResponses []dto.ResourceTypeResponse
	for _, resourceType := range resourceTypes {
		resourceTypeResponses = append(resourceTypeResponses, dto.ResourceTypeResponse{
			ID:      resourceType.ID,
			Name:    resourceType.Name,
			IconURL: resourceType.IconURL,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"resourceTypes": resourceTypeResponses,
	})
}
