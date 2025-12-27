package controllers

import (
	"net/http"
	"strconv"

	"estsoftwareoficial/src/technologies/application"
	"estsoftwareoficial/src/technologies/domain/dto"

	"github.com/gin-gonic/gin"
)

type GetTechnologyByIdController struct {
	getTechnologyById *application.GetTechnologyById
}

func NewGetTechnologyByIdController(getTechnologyById *application.GetTechnologyById) *GetTechnologyByIdController {
	return &GetTechnologyByIdController{getTechnologyById: getTechnologyById}
}

func (gt *GetTechnologyByIdController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	technology, err := gt.getTechnologyById.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if technology == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tecnología no encontrada"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"technology": dto.TechnologyResponse{
			ID:   technology.ID,
			Name: technology.Name,
			Icon: technology.Icon,
		},
	})
}
