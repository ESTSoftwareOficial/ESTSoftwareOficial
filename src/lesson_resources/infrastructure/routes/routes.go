package routes

import (
	"estsoftwareoficial/src/core/security"
	"estsoftwareoficial/src/lesson_resources/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigureLessonResourceRoutes(
	router *gin.Engine,
	createLessonResourceCtrl *controllers.CreateLessonResourceController,
	getAllLessonResourcesCtrl *controllers.GetAllLessonResourcesController,
	getLessonResourceByIdCtrl *controllers.GetLessonResourceByIdController,
	getLessonResourcesByLessonCtrl *controllers.GetLessonResourcesByLessonController,
	updateLessonResourceCtrl *controllers.UpdateLessonResourceController,
	deleteLessonResourceCtrl *controllers.DeleteLessonResourceController,
) {
	lessonResourceGroup := router.Group("/lesson-resources")
	{
		// Rutas públicas
		lessonResourceGroup.GET("", getAllLessonResourcesCtrl.Execute)
		lessonResourceGroup.GET("/:id", getLessonResourceByIdCtrl.Execute)
		lessonResourceGroup.GET("/lesson/:lessonId", getLessonResourcesByLessonCtrl.Execute)

		// Rutas protegidas (Admin o Instructor)
		lessonResourceGroup.POST("", security.JWTMiddleware(), security.RequireAnyRole(1, 2), createLessonResourceCtrl.Execute)
		lessonResourceGroup.PUT("/:id", security.JWTMiddleware(), security.RequireAnyRole(1, 2), updateLessonResourceCtrl.Execute)

		// Solo Admin puede eliminar
		lessonResourceGroup.DELETE("/:id", security.JWTMiddleware(), security.RequireRole(1), deleteLessonResourceCtrl.Execute)
	}
}
