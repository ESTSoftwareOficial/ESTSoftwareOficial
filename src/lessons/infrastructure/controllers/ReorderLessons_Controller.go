package controllers

import (
	"net/http"

	"estsoftwareoficial/src/lessons/application"
	"estsoftwareoficial/src/lessons/domain/dto"

	"github.com/gin-gonic/gin"
)

type ReorderLessonsController struct {
	reorderLessons *application.ReorderLessons
}

func NewReorderLessonsController(reorderLessons *application.ReorderLessons) *ReorderLessonsController {
	return &ReorderLessonsController{reorderLessons: reorderLessons}
}

func (rl *ReorderLessonsController) Execute(c *gin.Context) {
	var req dto.ReorderLessonsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var lessons []application.LessonOrder
	for _, l := range req.Lessons {
		lessons = append(lessons, application.LessonOrder{
			ID:         l.ID,
			OrderIndex: l.OrderIndex,
		})
	}

	err := rl.reorderLessons.Execute(lessons)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Lecciones reordenadas exitosamente",
	})
}
