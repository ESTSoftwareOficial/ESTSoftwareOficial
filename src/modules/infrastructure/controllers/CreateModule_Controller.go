package controllers

import (
	"net/http"

	"estsoftwareoficial/src/modules/application"
	"estsoftwareoficial/src/modules/domain/dto"
	"estsoftwareoficial/src/modules/domain/entities"

	"github.com/gin-gonic/gin"
)

type CreateModuleController struct {
	createModule *application.CreateModule
}

func NewCreateModuleController(createModule *application.CreateModule) *CreateModuleController {
	return &CreateModuleController{createModule: createModule}
}

func (cm *CreateModuleController) Execute(c *gin.Context) {
	var req dto.ModuleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	module := &entities.Module{
		CourseID:    req.CourseID,
		Title:       req.Title,
		Description: req.Description,
		OrderIndex:  req.OrderIndex,
	}

	savedModule, err := cm.createModule.Execute(module)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Módulo creado exitosamente",
		"module": dto.ModuleResponse{
			ID:          savedModule.ID,
			CourseID:    savedModule.CourseID,
			Title:       savedModule.Title,
			Description: savedModule.Description,
			OrderIndex:  savedModule.OrderIndex,
		},
	})
}
