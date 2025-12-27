package controllers

import (
	"net/http"

	"estsoftwareoficial/src/modules/application"
	"estsoftwareoficial/src/modules/domain/dto"

	"github.com/gin-gonic/gin"
)

type ReorderModulesController struct {
	reorderModules *application.ReorderModules
}

func NewReorderModulesController(reorderModules *application.ReorderModules) *ReorderModulesController {
	return &ReorderModulesController{reorderModules: reorderModules}
}

func (rm *ReorderModulesController) Execute(c *gin.Context) {
	var req dto.ReorderModulesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var modules []application.ModuleOrder
	for _, m := range req.Modules {
		modules = append(modules, application.ModuleOrder{
			ID:         m.ID,
			OrderIndex: m.OrderIndex,
		})
	}

	err := rm.reorderModules.Execute(modules)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Módulos reordenados exitosamente",
	})
}
