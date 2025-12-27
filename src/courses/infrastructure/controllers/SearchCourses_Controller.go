package controllers

import (
	"net/http"
	"strconv"

	"estsoftwareoficial/src/courses/application"
	"estsoftwareoficial/src/courses/domain/dto"

	"github.com/gin-gonic/gin"
)

type SearchCoursesController struct {
	searchCourses *application.SearchCourses
}

func NewSearchCoursesController(searchCourses *application.SearchCourses) *SearchCoursesController {
	return &SearchCoursesController{searchCourses: searchCourses}
}

func (sc *SearchCoursesController) Execute(c *gin.Context) {
	keyword := c.Query("keyword")
	categoryIDStr := c.Query("categoryId")
	technologyIDStr := c.Query("technologyId")
	level := c.Query("level")
	minRatingStr := c.Query("minRating")

	var categoryID *int
	if categoryIDStr != "" {
		id, err := strconv.Atoi(categoryIDStr)
		if err == nil {
			categoryID = &id
		}
	}

	var technologyID *int
	if technologyIDStr != "" {
		id, err := strconv.Atoi(technologyIDStr)
		if err == nil {
			technologyID = &id
		}
	}

	var levelPtr *string
	if level != "" {
		levelPtr = &level
	}

	var minRating *float64
	if minRatingStr != "" {
		rating, err := strconv.ParseFloat(minRatingStr, 64)
		if err == nil {
			minRating = &rating
		}
	}

	courses, err := sc.searchCourses.Execute(keyword, categoryID, technologyID, levelPtr, minRating)
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
		"total":   len(courseResponses),
	})
}
