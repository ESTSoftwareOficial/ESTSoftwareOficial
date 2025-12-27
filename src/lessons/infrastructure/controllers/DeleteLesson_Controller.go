package controllers

import (
	"net/http"
	"strconv"

	"estsoftwareoficial/src/lessons/application"

	"github.com/gin-gonic/gin"
)

type DeleteLessonController struct {
	deleteLesson *application.DeleteLesson
}

func NewDeleteLessonController(deleteLesson *application.DeleteLesson) *DeleteLessonController {
	return &DeleteLessonController{deleteLesson: deleteLesson}
}

func (dl *DeleteLessonController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = dl.deleteLesson.Execute(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Lección eliminada exitosamente",
	})
}
