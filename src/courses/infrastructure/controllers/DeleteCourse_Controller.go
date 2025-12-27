package controllers

import (
	"net/http"
	"strconv"

	"estsoftwareoficial/src/courses/application"

	"github.com/gin-gonic/gin"
)

type DeleteCourseController struct {
	deleteCourse *application.DeleteCourse
}

func NewDeleteCourseController(deleteCourse *application.DeleteCourse) *DeleteCourseController {
	return &DeleteCourseController{deleteCourse: deleteCourse}
}

func (dc *DeleteCourseController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = dc.deleteCourse.Execute(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Curso eliminado exitosamente",
	})
}
