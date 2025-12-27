package controllers

import (
	"net/http"

	"estsoftwareoficial/src/technologies/application"
	"estsoftwareoficial/src/technologies/domain/dto"
	"estsoftwareoficial/src/technologies/domain/entities"

	"github.com/gin-gonic/gin"
)

type CreateTechnologyController struct {
	createTechnology *application.CreateTechnology
}

func NewCreateTechnologyController(createTechnology *application.CreateTechnology) *CreateTechnologyController {
	return &CreateTechnologyController{createTechnology: createTechnology}
}

func (ct *CreateTechnologyController) Execute(c *gin.Context) {
	var req dto.TechnologyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	technology := &entities.Technology{
		Name: req.Name,
		Icon: req.Icon,
	}

	savedTechnology, err := ct.createTechnology.Execute(technology)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Tecnología creada exitosamente",
		"technology": dto.TechnologyResponse{
			ID:   savedTechnology.ID,
			Name: savedTechnology.Name,
			Icon: savedTechnology.Icon,
		},
	})
}
