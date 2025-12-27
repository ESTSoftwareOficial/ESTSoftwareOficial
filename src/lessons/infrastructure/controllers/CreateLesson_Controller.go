package controllers

import (
	"net/http"

	"estsoftwareoficial/src/lessons/application"
	"estsoftwareoficial/src/lessons/domain/dto"
	"estsoftwareoficial/src/lessons/domain/entities"

	"github.com/gin-gonic/gin"
)

type CreateLessonController struct {
	createLesson *application.CreateLesson
}

func NewCreateLessonController(createLesson *application.CreateLesson) *CreateLessonController {
	return &CreateLessonController{createLesson: createLesson}
}

func (cl *CreateLessonController) Execute(c *gin.Context) {
	var req dto.LessonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lesson := &entities.Lesson{
		ModuleID:        req.ModuleID,
		Title:           req.Title,
		ContentType:     req.ContentType,
		ContentURL:      req.ContentURL,
		BodyText:        req.BodyText,
		DurationMinutes: req.DurationMinutes,
		OrderIndex:      req.OrderIndex,
		IsPreview:       req.IsPreview,
	}

	savedLesson, err := cl.createLesson.Execute(lesson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Lección creada exitosamente",
		"lesson": dto.LessonResponse{
			ID:              savedLesson.ID,
			ModuleID:        savedLesson.ModuleID,
			Title:           savedLesson.Title,
			ContentType:     savedLesson.ContentType,
			ContentURL:      savedLesson.ContentURL,
			BodyText:        savedLesson.BodyText,
			DurationMinutes: savedLesson.DurationMinutes,
			OrderIndex:      savedLesson.OrderIndex,
			IsPreview:       savedLesson.IsPreview,
			CreatedAt:       savedLesson.CreatedAt,
		},
	})
}
