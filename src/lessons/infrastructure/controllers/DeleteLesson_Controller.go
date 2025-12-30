package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"estsoftwareoficial/src/core/bunny"
	"estsoftwareoficial/src/lessons/application"

	"github.com/gin-gonic/gin"
)

type DeleteLessonController struct {
	deleteLesson  *application.DeleteLesson
	getLessonById *application.GetLessonById
}

func NewDeleteLessonController(deleteLesson *application.DeleteLesson, getLessonById *application.GetLessonById) *DeleteLessonController {
	return &DeleteLessonController{
		deleteLesson:  deleteLesson,
		getLessonById: getLessonById,
	}
}

func (dl *DeleteLessonController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	lesson, err := dl.getLessonById.Execute(id)
	if err != nil || lesson == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lección no encontrada"})
		return
	}

	if lesson.BunnyVideoID != nil && *lesson.BunnyVideoID != "" {
		fmt.Printf("Eliminando video de Bunny.net: %s\n", *lesson.BunnyVideoID)
		err := bunny.DeleteVideo(*lesson.BunnyVideoID)
		if err != nil {
			fmt.Printf("WARNING: No se pudo eliminar el video de Bunny.net: %v\n", err)
		}
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
