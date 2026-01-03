package controllers

import (
	"net/http"

	"estsoftwareoficial/src/courses/application"
	"estsoftwareoficial/src/courses/domain/dto"

	"github.com/gin-gonic/gin"
)

type GetAllCoursesController struct {
	getAllCourses *application.GetAllCourses
}

func NewGetAllCoursesController(getAllCourses *application.GetAllCourses) *GetAllCoursesController {
	return &GetAllCoursesController{getAllCourses: getAllCourses}
}

func (gc *GetAllCoursesController) Execute(c *gin.Context) {
	courses, err := gc.getAllCourses.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var courseResponses []dto.CourseResponse
	for _, course := range courses {
		courseResponses = append(courseResponses, dto.CourseResponse{
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
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"courses": courseResponses,
	})
}