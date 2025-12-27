package controllers

import (
	"net/http"
	"strconv"

	"estsoftwareoficial/src/modules/application"
	"estsoftwareoficial/src/modules/domain/dto"

	"github.com/gin-gonic/gin"
)

type GetModulesByCourseController struct {
	getModulesByCourse *application.GetModulesByCourse
}

func NewGetModulesByCourseController(getModulesByCourse *application.GetModulesByCourse) *GetModulesByCourseController {
	return &GetModulesByCourseController{getModulesByCourse: getModulesByCourse}
}

func (gm *GetModulesByCourseController) Execute(c *gin.Context) {
	courseIDStr := c.Param("courseId")
	courseID, err := strconv.Atoi(courseIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de curso inválido"})
		return
	}

	modules, err := gm.getModulesByCourse.Execute(courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var moduleResponses []dto.ModuleResponse
	for _, module := range modules {
		moduleResponses = append(moduleResponses, dto.ModuleResponse{
			ID:          module.ID,
			CourseID:    module.CourseID,
			Title:       module.Title,
			Description: module.Description,
			OrderIndex:  module.OrderIndex,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"modules": moduleResponses,
	})
}
