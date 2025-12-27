package controllers

import (
	"net/http"
	"strconv"

	"estsoftwareoficial/src/modules/application"

	"github.com/gin-gonic/gin"
)

type DeleteModuleController struct {
	deleteModule *application.DeleteModule
}

func NewDeleteModuleController(deleteModule *application.DeleteModule) *DeleteModuleController {
	return &DeleteModuleController{deleteModule: deleteModule}
}

func (dm *DeleteModuleController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = dm.deleteModule.Execute(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Módulo eliminado exitosamente",
	})
}
