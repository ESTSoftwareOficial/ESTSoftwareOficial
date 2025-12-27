package controllers

import (
	"net/http"
	"strconv"

	"estsoftwareoficial/src/course_ratings/application"
	"estsoftwareoficial/src/course_ratings/domain/dto"

	"github.com/gin-gonic/gin"
)

type GetCourseRatingsByCourseController struct {
	getCourseRatingsByCourse *application.GetCourseRatingsByCourse
}

func NewGetCourseRatingsByCourseController(getCourseRatingsByCourse *application.GetCourseRatingsByCourse) *GetCourseRatingsByCourseController {
	return &GetCourseRatingsByCourseController{getCourseRatingsByCourse: getCourseRatingsByCourse}
}

func (gcr *GetCourseRatingsByCourseController) Execute(c *gin.Context) {
	courseIDStr := c.Param("courseId")
	courseID, err := strconv.Atoi(courseIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de curso inválido"})
		return
	}

	courseRatings, err := gcr.getCourseRatingsByCourse.Execute(courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var courseRatingResponses []dto.CourseRatingResponse
	for _, courseRating := range courseRatings {
		courseRatingResponses = append(courseRatingResponses, dto.CourseRatingResponse{
			ID:        courseRating.ID,
			CourseID:  courseRating.CourseID,
			UserID:    courseRating.UserID,
			Rating:    courseRating.Rating,
			Review:    courseRating.Review,
			CreatedAt: courseRating.CreatedAt,
			UpdatedAt: courseRating.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"courseRatings": courseRatingResponses,
	})
}
