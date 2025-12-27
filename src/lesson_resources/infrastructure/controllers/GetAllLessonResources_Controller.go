package controllers

import (
	"net/http"

	"estsoftwareoficial/src/lesson_resources/application"
	"estsoftwareoficial/src/lesson_resources/domain/dto"

	"github.com/gin-gonic/gin"
)

type GetAllLessonResourcesController struct {
	getAllLessonResources *application.GetAllLessonResources
}

func NewGetAllLessonResourcesController(getAllLessonResources *application.GetAllLessonResources) *GetAllLessonResourcesController {
	return &GetAllLessonResourcesController{getAllLessonResources: getAllLessonResources}
}

func (glr *GetAllLessonResourcesController) Execute(c *gin.Context) {
	lessonResources, err := glr.getAllLessonResources.Execute()
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
