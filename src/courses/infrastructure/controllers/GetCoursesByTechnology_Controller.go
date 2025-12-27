package controllers

import (
	"net/http"
	"strconv"

	"estsoftwareoficial/src/courses/application"
	"estsoftwareoficial/src/courses/domain/dto"

	"github.com/gin-gonic/gin"
)

type GetCoursesByTechnologyController struct {
	getCoursesByTechnology *application.GetCoursesByTechnology
}

func NewGetCoursesByTechnologyController(getCoursesByTechnology *application.GetCoursesByTechnology) *GetCoursesByTechnologyController {
	return &GetCoursesByTechnologyController{getCoursesByTechnology: getCoursesByTechnology}
}

func (gc *GetCoursesByTechnologyController) Execute(c *gin.Context) {
	technologyIDStr := c.Param("technologyId")
	technologyID, err := strconv.Atoi(technologyIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de tecnología inválido"})
		return
	}

	courses, err := gc.getCoursesByTechnology.Execute(technologyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var courseResponses []dto.CourseResponse
	for _, course := range courses {
		courseResponses = append(courseResponses, dto.CourseResponse{
			ID:            course.ID,
			NameCourse:    course.NameCourse,
			Description:   course.Description,
			TechnologyID:  course.TechnologyID,
			InstructorID:  course.InstructorID,
			CategoryID:    course.CategoryID,
			Level:         course.Level,
			ImageURL:      course.ImageURL,
			TotalModules:  course.TotalModules,
			AverageRating: course.AverageRating,
			TotalRatings:  course.TotalRatings,
			DurationHours: course.DurationHours,
			CreatedAt:     course.CreatedAt,
			UpdatedAt:     course.UpdatedAt,
			IsActive:      course.IsActive,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"courses": courseResponses,
	})
}
