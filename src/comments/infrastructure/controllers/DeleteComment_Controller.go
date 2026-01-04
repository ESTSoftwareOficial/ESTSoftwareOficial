package controllers

import (
	"estsoftwareoficial/src/comments/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteCommentController struct {
	deleteComment *application.DeleteComment
}

func NewDeleteCommentController(deleteComment *application.DeleteComment) *DeleteCommentController {
	return &DeleteCommentController{deleteComment: deleteComment}
}

func (dcc *DeleteCommentController) Execute(c *gin.Context) {
	commentIDStr := c.Param("commentId")
	commentID, err := strconv.Atoi(commentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de comentario inválido"})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	err = dcc.deleteComment.Execute(commentID, userID.(int))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Comentario eliminado exitosamente",
		"commentId": commentID,
	})
}
