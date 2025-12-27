package controllers

import (
	"net/http"
	"strconv"

	"estsoftwareoficial/src/modules/application"
	"estsoftwareoficial/src/modules/domain/dto"
	"estsoftwareoficial/src/modules/domain/entities"

	"github.com/gin-gonic/gin"
)

type UpdateModuleController struct {
	updateModule *application.UpdateModule
}

func NewUpdateModuleController(updateModule *application.UpdateModule) *UpdateModuleController {
	return &UpdateModuleController{updateModule: updateModule}
}

func (um *UpdateModuleController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req dto.UpdateModuleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	module := &entities.Module{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		OrderIndex:  req.OrderIndex,
	}

	err = um.updateModule.Execute(module)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Módulo actualizado exitosamente",
	})
}
