package controllers

import (
	"net/http"
	"strconv"

	"estsoftwareoficial/src/course_ratings/application"
	"estsoftwareoficial/src/course_ratings/domain/dto"
	"estsoftwareoficial/src/course_ratings/domain/entities"

	"github.com/gin-gonic/gin"
)

type UpdateCourseRatingController struct {
	updateCourseRating *application.UpdateCourseRating
}

func NewUpdateCourseRatingController(updateCourseRating *application.UpdateCourseRating) *UpdateCourseRatingController {
	return &UpdateCourseRatingController{updateCourseRating: updateCourseRating}
}

func (ucr *UpdateCourseRatingController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req dto.UpdateCourseRatingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	courseRating := &entities.CourseRating{
		ID:     id,
		Rating: req.Rating,
		Review: req.Review,
	}

	err = ucr.updateCourseRating.Execute(courseRating)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Calificación actualizada exitosamente",
	})
}
