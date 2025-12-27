package routes

import (
	"estsoftwareoficial/src/core/security"
	"estsoftwareoficial/src/courses/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigureCourseRoutes(
	router *gin.Engine,
	createCourseCtrl *controllers.CreateCourseController,
	getAllCoursesCtrl *controllers.GetAllCoursesController,
	getCourseByIdCtrl *controllers.GetCourseByIdController,
	getCoursesByInstructorCtrl *controllers.GetCoursesByInstructorController,
	getCoursesByCategoryCtrl *controllers.GetCoursesByCategoryController,
	getCoursesByTechnologyCtrl *controllers.GetCoursesByTechnologyController,
	updateCourseCtrl *controllers.UpdateCourseController,
	deleteCourseCtrl *controllers.DeleteCourseController,
	searchCoursesCtrl *controllers.SearchCoursesController,
) {
	courseGroup := router.Group("/courses")
	{
		// Rutas públicas
		courseGroup.GET("", getAllCoursesCtrl.Execute)
		courseGroup.GET("/search", searchCoursesCtrl.Execute)
		courseGroup.GET("/:id", getCourseByIdCtrl.Execute)
		courseGroup.GET("/instructor/:instructorId", getCoursesByInstructorCtrl.Execute)
		courseGroup.GET("/category/:categoryId", getCoursesByCategoryCtrl.Execute)
		courseGroup.GET("/technology/:technologyId", getCoursesByTechnologyCtrl.Execute)

		// Rutas protegidas (Admin o Instructor)
		courseGroup.POST("", security.JWTMiddleware(), security.RequireAnyRole(1, 2), createCourseCtrl.Execute)
		courseGroup.PUT("/:id", security.JWTMiddleware(), security.RequireAnyRole(1, 2), updateCourseCtrl.Execute)

		// Solo Admin puede eliminar
		courseGroup.DELETE("/:id", security.JWTMiddleware(), security.RequireRole(1), deleteCourseCtrl.Execute)
	}
}
