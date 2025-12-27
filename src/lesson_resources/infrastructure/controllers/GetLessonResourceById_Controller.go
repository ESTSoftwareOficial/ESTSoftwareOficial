package controllers

import (
	"net/http"
	"strconv"

	"estsoftwareoficial/src/lesson_resources/application"
	"estsoftwareoficial/src/lesson_resources/domain/dto"

	"github.com/gin-gonic/gin"
)

type GetLessonResourceByIdController struct {
	getLessonResourceById *application.GetLessonResourceById
}

func NewGetLessonResourceByIdController(getLessonResourceById *application.GetLessonResourceById) *GetLessonResourceByIdController {
	return &GetLessonResourceByIdController{getLessonResourceById: getLessonResourceById}
}

func (glr *GetLessonResourceByIdController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	lessonResource, err := glr.getLessonResourceById.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if lessonResource == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recurso de lección no encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"lessonResource": dto.LessonResourceResponse{
			ID:             lessonResource.ID,
			LessonID:       lessonResource.LessonID,
			ResourceTypeID: lessonResource.ResourceTypeID,
			URL:            lessonResource.URL,
			Title:          lessonResource.Title,
		},
	})
}
