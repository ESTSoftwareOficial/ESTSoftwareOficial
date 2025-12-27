package controllers

import (
	"net/http"

	"estsoftwareoficial/src/lessons/application"
	"estsoftwareoficial/src/lessons/domain/dto"

	"github.com/gin-gonic/gin"
)

type GetAllLessonsController struct {
	getAllLessons *application.GetAllLessons
}

func NewGetAllLessonsController(getAllLessons *application.GetAllLessons) *GetAllLessonsController {
	return &GetAllLessonsController{getAllLessons: getAllLessons}
}

func (gl *GetAllLessonsController) Execute(c *gin.Context) {
	lessons, err := gl.getAllLessons.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var lessonResponses []dto.LessonResponse
	for _, lesson := range lessons {
		lessonResponses = append(lessonResponses, dto.LessonResponse{
			ID:              lesson.ID,
			ModuleID:        lesson.ModuleID,
			Title:           lesson.Title,
			ContentType:     lesson.ContentType,
			ContentURL:      lesson.ContentURL,
			BodyText:        lesson.BodyText,
			DurationMinutes: lesson.DurationMinutes,
			OrderIndex:      lesson.OrderIndex,
			IsPreview:       lesson.IsPreview,
			CreatedAt:       lesson.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"lessons": lessonResponses,
	})
}
