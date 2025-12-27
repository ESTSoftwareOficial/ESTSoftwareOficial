package controllers

import (
	"net/http"
	"strconv"

	"estsoftwareoficial/src/modules/application"
	"estsoftwareoficial/src/modules/domain/dto"

	"github.com/gin-gonic/gin"
)

type GetModuleByIdController struct {
	getModuleById *application.GetModuleById
}

func NewGetModuleByIdController(getModuleById *application.GetModuleById) *GetModuleByIdController {
	return &GetModuleByIdController{getModuleById: getModuleById}
}

func (gm *GetModuleByIdController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	module, err := gm.getModuleById.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if module == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Módulo no encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"module": dto.ModuleResponse{
			ID:          module.ID,
			CourseID:    module.CourseID,
			Title:       module.Title,
			Description: module.Description,
			OrderIndex:  module.OrderIndex,
		},
	})
}
