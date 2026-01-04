package controllers

import (
	"estsoftwareoficial/src/comments/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetLessonCommentsController struct {
	getLessonComments *application.GetLessonComments
}

func NewGetLessonCommentsController(getLessonComments *application.GetLessonComments) *GetLessonCommentsController {
	return &GetLessonCommentsController{getLessonComments: getLessonComments}
}

func (glcc *GetLessonCommentsController) Execute(c *gin.Context) {
	lessonIDStr := c.Param("id")
	lessonID, err := strconv.Atoi(lessonIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de lección inválido"})
		return
	}

	userID := 0
	if uid, exists := c.Get("userID"); exists {
		userID = uid.(int)
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	response, err := glcc.getLessonComments.Execute(lessonID, userID, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
