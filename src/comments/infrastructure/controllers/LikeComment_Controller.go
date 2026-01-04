package controllers

import (
	"estsoftwareoficial/src/comments/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LikeCommentController struct {
	likeComment *application.LikeComment
}

func NewLikeCommentController(likeComment *application.LikeComment) *LikeCommentController {
	return &LikeCommentController{likeComment: likeComment}
}

func (lcc *LikeCommentController) Execute(c *gin.Context) {
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

	likesCount, err := lcc.likeComment.Execute(commentID, userID.(int))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"commentId":  commentID,
		"liked":      true,
		"likesCount": likesCount,
	})
}
