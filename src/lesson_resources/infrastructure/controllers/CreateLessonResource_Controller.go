package controllers

import (
	"net/http"

	"estsoftwareoficial/src/lesson_resources/application"
	"estsoftwareoficial/src/lesson_resources/domain/dto"
	"estsoftwareoficial/src/lesson_resources/domain/entities"

	"github.com/gin-gonic/gin"
)

type CreateLessonResourceController struct {
	createLessonResource *application.CreateLessonResource
}

func NewCreateLessonResourceController(createLessonResource *application.CreateLessonResource) *CreateLessonResourceController {
	return &CreateLessonResourceController{createLessonResource: createLessonResource}
}

func (clr *CreateLessonResourceController) Execute(c *gin.Context) {
	var req dto.LessonResourceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lessonResource := &entities.LessonResource{
		LessonID:       req.LessonID,
		ResourceTypeID: req.ResourceTypeID,
		URL:            req.URL,
		Title:          req.Title,
	}

	savedLessonResource, err := clr.createLessonResource.Execute(lessonResource)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Recurso de lección creado exitosamente",
		"lessonResource": dto.LessonResourceResponse{
			ID:             savedLessonResource.ID,
			LessonID:       savedLessonResource.LessonID,
			ResourceTypeID: savedLessonResource.ResourceTypeID,
			URL:            savedLessonResource.URL,
			Title:          savedLessonResource.Title,
		},
	})
}
