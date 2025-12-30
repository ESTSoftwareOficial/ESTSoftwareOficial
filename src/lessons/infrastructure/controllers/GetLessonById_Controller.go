package controllers

import (
	"net/http"
	"strconv"

	"estsoftwareoficial/src/lessons/application"
	"estsoftwareoficial/src/lessons/domain/dto"

	"github.com/gin-gonic/gin"
)

type GetLessonByIdController struct {
	getLessonById *application.GetLessonById
}

func NewGetLessonByIdController(getLessonById *application.GetLessonById) *GetLessonByIdController {
	return &GetLessonByIdController{getLessonById: getLessonById}
}

func (gl *GetLessonByIdController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	lesson, err := gl.getLessonById.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if lesson == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lección no encontrada"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"lesson": dto.LessonResponse{
			ID:              lesson.ID,
			ModuleID:        lesson.ModuleID,
			Title:           lesson.Title,
			ContentType:     lesson.ContentType,
			VideoURL:        lesson.GetVideoURL(),
			BodyText:        lesson.BodyText,
			DurationMinutes: lesson.DurationMinutes,
			OrderIndex:      lesson.OrderIndex,
			IsPreview:       lesson.IsPreview,
			CreatedAt:       lesson.CreatedAt,
		},
	})
}
