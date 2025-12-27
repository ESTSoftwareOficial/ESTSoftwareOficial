package controllers

import (
	"net/http"
	"strconv"

	"estsoftwareoficial/src/lesson_resources/application"
	"estsoftwareoficial/src/lesson_resources/domain/dto"
	"estsoftwareoficial/src/lesson_resources/domain/entities"

	"github.com/gin-gonic/gin"
)

type UpdateLessonResourceController struct {
	updateLessonResource *application.UpdateLessonResource
}

func NewUpdateLessonResourceController(updateLessonResource *application.UpdateLessonResource) *UpdateLessonResourceController {
	return &UpdateLessonResourceController{updateLessonResource: updateLessonResource}
}

func (ulr *UpdateLessonResourceController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req dto.UpdateLessonResourceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lessonResource := &entities.LessonResource{
		ID:             id,
		ResourceTypeID: req.ResourceTypeID,
		URL:            req.URL,
		Title:          req.Title,
	}

	err = ulr.updateLessonResource.Execute(lessonResource)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Recurso de lección actualizado exitosamente",
	})
}
