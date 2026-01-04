package controllers

import (
	"estsoftwareoficial/src/comments/application"
	"estsoftwareoficial/src/comments/domain/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateCommentController struct {
	createComment *application.CreateComment
}

func NewCreateCommentController(createComment *application.CreateComment) *CreateCommentController {
	return &CreateCommentController{createComment: createComment}
}

func (ccc *CreateCommentController) Execute(c *gin.Context) {
	lessonIDStr := c.Param("id")
	lessonID, err := strconv.Atoi(lessonIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de lección inválido"})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	var req dto.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment, err := ccc.createComment.Execute(lessonID, userID.(int), req.Comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":           comment.ID,
		"lessonId":     comment.LessonID,
		"userId":       comment.UserID,
		"comment":      comment.Comment,
		"createdAt":    comment.CreatedAt,
		"isEdited":     comment.IsEdited,
		"likesCount":   0,
		"userHasLiked": false,
	})
}
