package controllers

import (
	"net/http"
	"strconv"

	"estsoftwareoficial/src/lesson_resources/application"

	"github.com/gin-gonic/gin"
)

type DeleteLessonResourceController struct {
	deleteLessonResource *application.DeleteLessonResource
}

func NewDeleteLessonResourceController(deleteLessonResource *application.DeleteLessonResource) *DeleteLessonResourceController {
	return &DeleteLessonResourceController{deleteLessonResource: deleteLessonResource}
}

func (dlr *DeleteLessonResourceController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = dlr.deleteLessonResource.Execute(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Recurso de lección eliminado exitosamente",
	})
}
