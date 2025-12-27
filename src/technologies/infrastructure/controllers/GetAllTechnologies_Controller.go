package controllers

import (
	"net/http"

	"estsoftwareoficial/src/technologies/application"
	"estsoftwareoficial/src/technologies/domain/dto"

	"github.com/gin-gonic/gin"
)

type GetAllTechnologiesController struct {
	getAllTechnologies *application.GetAllTechnologies
}

func NewGetAllTechnologiesController(getAllTechnologies *application.GetAllTechnologies) *GetAllTechnologiesController {
	return &GetAllTechnologiesController{getAllTechnologies: getAllTechnologies}
}

func (gt *GetAllTechnologiesController) Execute(c *gin.Context) {
	technologies, err := gt.getAllTechnologies.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var technologyResponses []dto.TechnologyResponse
	for _, technology := range technologies {
		technologyResponses = append(technologyResponses, dto.TechnologyResponse{
			ID:   technology.ID,
			Name: technology.Name,
			Icon: technology.Icon,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"technologies": technologyResponses,
	})
}
