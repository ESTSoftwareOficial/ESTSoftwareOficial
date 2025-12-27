package controllers

import (
	"net/http"
	"strconv"

	"estsoftwareoficial/src/course_ratings/application"
	"estsoftwareoficial/src/course_ratings/domain/dto"

	"github.com/gin-gonic/gin"
)

type GetCourseRatingByUserAndCourseController struct {
	getCourseRatingByUserAndCourse *application.GetCourseRatingByUserAndCourse
}

func NewGetCourseRatingByUserAndCourseController(getCourseRatingByUserAndCourse *application.GetCourseRatingByUserAndCourse) *GetCourseRatingByUserAndCourseController {
	return &GetCourseRatingByUserAndCourseController{getCourseRatingByUserAndCourse: getCourseRatingByUserAndCourse}
}

func (gcr *GetCourseRatingByUserAndCourseController) Execute(c *gin.Context) {
	userIDStr := c.Param("userId")
	courseIDStr := c.Param("courseId")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	courseID, err := strconv.Atoi(courseIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de curso inválido"})
		return
	}

	courseRating, err := gcr.getCourseRatingByUserAndCourse.Execute(userID, courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if courseRating == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Calificación no encontrada"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"courseRating": dto.CourseRatingResponse{
			ID:        courseRating.ID,
			CourseID:  courseRating.CourseID,
			UserID:    courseRating.UserID,
			Rating:    courseRating.Rating,
			Review:    courseRating.Review,
			CreatedAt: courseRating.CreatedAt,
			UpdatedAt: courseRating.UpdatedAt,
		},
	})
}
