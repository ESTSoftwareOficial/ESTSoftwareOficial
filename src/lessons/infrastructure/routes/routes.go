package routes

import (
	"estsoftwareoficial/src/core/security"
	"estsoftwareoficial/src/lessons/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigureLessonRoutes(
	router *gin.Engine,
	createLessonCtrl *controllers.CreateLessonController,
	getAllLessonsCtrl *controllers.GetAllLessonsController,
	getLessonByIdCtrl *controllers.GetLessonByIdController,
	getLessonsByModuleCtrl *controllers.GetLessonsByModuleController,
	updateLessonCtrl *controllers.UpdateLessonController,
	deleteLessonCtrl *controllers.DeleteLessonController,
	reorderLessonsCtrl *controllers.ReorderLessonsController,
) {
	lessonGroup := router.Group("/lessons")
	{
		// Rutas públicas
		lessonGroup.GET("", getAllLessonsCtrl.Execute)
		lessonGroup.GET("/:id", getLessonByIdCtrl.Execute)
		lessonGroup.GET("/module/:moduleId", getLessonsByModuleCtrl.Execute)

		// Rutas protegidas (Admin o Instructor)
		lessonGroup.POST("", security.JWTMiddleware(), security.RequireAnyRole(1, 2), createLessonCtrl.Execute)
		lessonGroup.PUT("/:id", security.JWTMiddleware(), security.RequireAnyRole(1, 2), updateLessonCtrl.Execute)
		lessonGroup.PUT("/reorder", security.JWTMiddleware(), security.RequireAnyRole(1, 2), reorderLessonsCtrl.Execute)

		// Solo Admin puede eliminar
		lessonGroup.DELETE("/:id", security.JWTMiddleware(), security.RequireRole(1), deleteLessonCtrl.Execute)
	}
}
