package controllers

import (
	"net/http"
	"strconv"

	"estsoftwareoficial/src/lessons/application"
	"estsoftwareoficial/src/lessons/domain/dto"
	"estsoftwareoficial/src/lessons/domain/entities"

	"github.com/gin-gonic/gin"
)

type UpdateLessonController struct {
	updateLesson *application.UpdateLesson
}

func NewUpdateLessonController(updateLesson *application.UpdateLesson) *UpdateLessonController {
	return &UpdateLessonController{updateLesson: updateLesson}
}

func (ul *UpdateLessonController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req dto.UpdateLessonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lesson := &entities.Lesson{
		ID:              id,
		Title:           req.Title,
		ContentType:     req.ContentType,
		ContentURL:      req.ContentURL,
		BodyText:        req.BodyText,
		DurationMinutes: req.DurationMinutes,
		OrderIndex:      req.OrderIndex,
		IsPreview:       req.IsPreview,
	}

	err = ul.updateLesson.Execute(lesson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Lección actualizada exitosamente",
	})
}
