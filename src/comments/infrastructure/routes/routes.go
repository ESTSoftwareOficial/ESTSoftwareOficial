package routes

import (
	"estsoftwareoficial/src/comments/infrastructure/controllers"
	"estsoftwareoficial/src/core/security"

	"github.com/gin-gonic/gin"
)

func ConfigureCommentRoutes(
	router *gin.Engine,
	createCommentCtrl *controllers.CreateCommentController,
	getLessonCommentsCtrl *controllers.GetLessonCommentsController,
	updateCommentCtrl *controllers.UpdateCommentController,
	deleteCommentCtrl *controllers.DeleteCommentController,
	likeCommentCtrl *controllers.LikeCommentController,
	unlikeCommentCtrl *controllers.UnlikeCommentController,
) {
	// Comentarios en lección
	router.POST("/lessons/:id/comments", security.JWTMiddleware(), createCommentCtrl.Execute)
	router.GET("/lessons/:id/comments", getLessonCommentsCtrl.Execute)

	// Operaciones sobre comentario individual
	commentGroup := router.Group("/comments")
	commentGroup.Use(security.JWTMiddleware())
	{
		commentGroup.PUT("/:commentId", updateCommentCtrl.Execute)
		commentGroup.DELETE("/:commentId", deleteCommentCtrl.Execute)
		commentGroup.POST("/:commentId/like", likeCommentCtrl.Execute)
		commentGroup.DELETE("/:commentId/like", unlikeCommentCtrl.Execute)
	}
}
