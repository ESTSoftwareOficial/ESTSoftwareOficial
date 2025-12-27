package controllers

import (
	"net/http"

	"estsoftwareoficial/src/modules/application"
	"estsoftwareoficial/src/modules/domain/dto"

	"github.com/gin-gonic/gin"
)

type GetAllModulesController struct {
	getAllModules *application.GetAllModules
}

func NewGetAllModulesController(getAllModules *application.GetAllModules) *GetAllModulesController {
	return &GetAllModulesController{getAllModules: getAllModules}
}

func (gm *GetAllModulesController) Execute(c *gin.Context) {
	modules, err := gm.getAllModules.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var moduleResponses []dto.ModuleResponse
	for _, module := range modules {
		moduleResponses = append(moduleResponses, dto.ModuleResponse{
			ID:          module.ID,
			CourseID:    module.CourseID,
			Title:       module.Title,
			Description: module.Description,
			OrderIndex:  module.OrderIndex,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"modules": moduleResponses,
	})
}
