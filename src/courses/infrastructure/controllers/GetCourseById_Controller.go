package controllers

import (
	"net/http"
	"strconv"

	"estsoftwareoficial/src/courses/application"
	"estsoftwareoficial/src/courses/domain/dto"

	"github.com/gin-gonic/gin"
)

type GetCourseByIdController struct {
	getCourseById *application.GetCourseById
}

func NewGetCourseByIdController(getCourseById *application.GetCourseById) *GetCourseByIdController {
	return &GetCourseByIdController{getCourseById: getCourseById}
}

func (gc *GetCourseByIdController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	course, err := gc.getCourseById.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if course == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Curso no encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"course": dto.CourseResponse{
			ID:          course.ID,
			NameCourse:  course.NameCourse,
			Description: course.Description,
			Technology: dto.TechnologyDTO{
				TechnologyName:  course.TechnologyName,
				TechnologyImage: course.TechnologyImage,
			},
			Instructor: dto.InstructorDTO{
				InstructorName:  course.InstructorName,
				InstructorImage: course.InstructorImage,
			},
			CategoryName:  course.CategoryName,
			Level:         course.Level,
			ImageURL:      course.ImageURL,
			TotalModules:  course.TotalModules,
			AverageRating: course.AverageRating,
			TotalRatings:  course.TotalRatings,
			DurationHours: course.DurationHours,
			CreatedAt:     course.CreatedAt,
			UpdatedAt:     course.UpdatedAt,
			IsActive:      course.IsActive,
		},
	})
}