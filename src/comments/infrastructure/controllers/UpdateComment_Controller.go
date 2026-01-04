package controllers

import (
	"estsoftwareoficial/src/comments/application"
	"estsoftwareoficial/src/comments/domain/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateCommentController struct {
	updateComment *application.UpdateComment
}

func NewUpdateCommentController(updateComment *application.UpdateComment) *UpdateCommentController {
	return &UpdateCommentController{updateComment: updateComment}
}

func (ucc *UpdateCommentController) Execute(c *gin.Context) {
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

	var req dto.UpdateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ucc.updateComment.Execute(commentID, userID.(int), req.Comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Comentario actualizado exitosamente",
		"commentId": commentID,
	})
}
