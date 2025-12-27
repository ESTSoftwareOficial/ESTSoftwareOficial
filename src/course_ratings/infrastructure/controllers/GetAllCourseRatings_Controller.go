package controllers

import (
	"net/http"

	"estsoftwareoficial/src/course_ratings/application"
	"estsoftwareoficial/src/course_ratings/domain/dto"

	"github.com/gin-gonic/gin"
)

type GetAllCourseRatingsController struct {
	getAllCourseRatings *application.GetAllCourseRatings
}

func NewGetAllCourseRatingsController(getAllCourseRatings *application.GetAllCourseRatings) *GetAllCourseRatingsController {
	return &GetAllCourseRatingsController{getAllCourseRatings: getAllCourseRatings}
}

func (gcr *GetAllCourseRatingsController) Execute(c *gin.Context) {
	courseRatings, err := gcr.getAllCourseRatings.Execute()
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
