package controllers

import (
	"net/http"
	"strconv"

	"estsoftwareoficial/src/technologies/application"
	"estsoftwareoficial/src/technologies/domain/dto"
	"estsoftwareoficial/src/technologies/domain/entities"

	"github.com/gin-gonic/gin"
)

type UpdateTechnologyController struct {
	updateTechnology *application.UpdateTechnology
}

func NewUpdateTechnologyController(updateTechnology *application.UpdateTechnology) *UpdateTechnologyController {
	return &UpdateTechnologyController{updateTechnology: updateTechnology}
}

func (ut *UpdateTechnologyController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req dto.TechnologyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	technology := &entities.Technology{
		ID:   id,
		Name: req.Name,
		Icon: req.Icon,
	}

	err = ut.updateTechnology.Execute(technology)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tecnología actualizada exitosamente",
	})
}
