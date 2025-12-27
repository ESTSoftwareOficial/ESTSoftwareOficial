package controllers

import (
	"net/http"
	"strconv"

	"estsoftwareoficial/src/course_ratings/application"
	"estsoftwareoficial/src/course_ratings/domain/dto"

	"github.com/gin-gonic/gin"
)

type GetCourseRatingByIdController struct {
	getCourseRatingById *application.GetCourseRatingById
}

func NewGetCourseRatingByIdController(getCourseRatingById *application.GetCourseRatingById) *GetCourseRatingByIdController {
	return &GetCourseRatingByIdController{getCourseRatingById: getCourseRatingById}
}

func (gcr *GetCourseRatingByIdController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	courseRating, err := gcr.getCourseRatingById.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
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
