package controllers

import (
	"net/http"

	"estsoftwareoficial/src/course_ratings/application"
	"estsoftwareoficial/src/course_ratings/domain/dto"
	"estsoftwareoficial/src/course_ratings/domain/entities"

	"github.com/gin-gonic/gin"
)

type CreateCourseRatingController struct {
	createCourseRating *application.CreateCourseRating
}

func NewCreateCourseRatingController(createCourseRating *application.CreateCourseRating) *CreateCourseRatingController {
	return &CreateCourseRatingController{createCourseRating: createCourseRating}
}

func (ccr *CreateCourseRatingController) Execute(c *gin.Context) {
	var req dto.CourseRatingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	courseRating := &entities.CourseRating{
		CourseID: req.CourseID,
		UserID:   req.UserID,
		Rating:   req.Rating,
		Review:   req.Review,
	}

	savedCourseRating, err := ccr.createCourseRating.Execute(courseRating)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Calificación creada exitosamente",
		"courseRating": dto.CourseRatingResponse{
			ID:        savedCourseRating.ID,
			CourseID:  savedCourseRating.CourseID,
			UserID:    savedCourseRating.UserID,
			Rating:    savedCourseRating.Rating,
			Review:    savedCourseRating.Review,
			CreatedAt: savedCourseRating.CreatedAt,
			UpdatedAt: savedCourseRating.UpdatedAt,
		},
	})
}
