package controllers

import (
	"net/http"
	"strconv"

	"estsoftwareoficial/src/lessons/application"
	"estsoftwareoficial/src/lessons/domain/dto"

	"github.com/gin-gonic/gin"
)

type GetLessonsByModuleController struct {
	getLessonsByModule *application.GetLessonsByModule
}

func NewGetLessonsByModuleController(getLessonsByModule *application.GetLessonsByModule) *GetLessonsByModuleController {
	return &GetLessonsByModuleController{getLessonsByModule: getLessonsByModule}
}

func (gl *GetLessonsByModuleController) Execute(c *gin.Context) {
	moduleIDStr := c.Param("moduleId")
	moduleID, err := strconv.Atoi(moduleIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de módulo inválido"})
		return
	}

	lessons, err := gl.getLessonsByModule.Execute(moduleID)
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
			VideoURL:        lesson.GetVideoURL(),
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
