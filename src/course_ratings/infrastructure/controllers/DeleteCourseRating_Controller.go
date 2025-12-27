package controllers

import (
	"net/http"
	"strconv"

	"estsoftwareoficial/src/course_ratings/application"

	"github.com/gin-gonic/gin"
)

type DeleteCourseRatingController struct {
	deleteCourseRating *application.DeleteCourseRating
}

func NewDeleteCourseRatingController(deleteCourseRating *application.DeleteCourseRating) *DeleteCourseRatingController {
	return &DeleteCourseRatingController{deleteCourseRating: deleteCourseRating}
}

func (dcr *DeleteCourseRatingController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = dcr.deleteCourseRating.Execute(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Calificación eliminada exitosamente",
	})
}
