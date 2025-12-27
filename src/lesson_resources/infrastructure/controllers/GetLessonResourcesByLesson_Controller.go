package controllers

import (
	"net/http"
	"strconv"

	"estsoftwareoficial/src/lesson_resources/application"
	"estsoftwareoficial/src/lesson_resources/domain/dto"

	"github.com/gin-gonic/gin"
)

type GetLessonResourcesByLessonController struct {
	getLessonResourcesByLesson *application.GetLessonResourcesByLesson
}

func NewGetLessonResourcesByLessonController(getLessonResourcesByLesson *application.GetLessonResourcesByLesson) *GetLessonResourcesByLessonController {
	return &GetLessonResourcesByLessonController{getLessonResourcesByLesson: getLessonResourcesByLesson}
}

func (glr *GetLessonResourcesByLessonController) Execute(c *gin.Context) {
	lessonIDStr := c.Param("lessonId")
	lessonID, err := strconv.Atoi(lessonIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de lección inválido"})
		return
	}

	lessonResources, err := glr.getLessonResourcesByLesson.Execute(lessonID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var lessonResourceResponses []dto.LessonResourceResponse
	for _, lessonResource := range lessonResources {
		lessonResourceResponses = append(lessonResourceResponses, dto.LessonResourceResponse{
			ID:             lessonResource.ID,
			LessonID:       lessonResource.LessonID,
			ResourceTypeID: lessonResource.ResourceTypeID,
			URL:            lessonResource.URL,
			Title:          lessonResource.Title,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"lessonResources": lessonResourceResponses,
	})
}
