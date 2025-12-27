package routes

import (
	"estsoftwareoficial/src/core/security"
	"estsoftwareoficial/src/course_ratings/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigureCourseRatingRoutes(
	router *gin.Engine,
	createCourseRatingCtrl *controllers.CreateCourseRatingController,
	getAllCourseRatingsCtrl *controllers.GetAllCourseRatingsController,
	getCourseRatingByIdCtrl *controllers.GetCourseRatingByIdController,
	getCourseRatingsByCourseCtrl *controllers.GetCourseRatingsByCourseController,
	getCourseRatingByUserAndCourseCtrl *controllers.GetCourseRatingByUserAndCourseController,
	updateCourseRatingCtrl *controllers.UpdateCourseRatingController,
	deleteCourseRatingCtrl *controllers.DeleteCourseRatingController,
) {
	courseRatingGroup := router.Group("/course-ratings")
	{
		// Rutas públicas
		courseRatingGroup.GET("", getAllCourseRatingsCtrl.Execute)
		courseRatingGroup.GET("/:id", getCourseRatingByIdCtrl.Execute)
		courseRatingGroup.GET("/course/:courseId", getCourseRatingsByCourseCtrl.Execute)
		courseRatingGroup.GET("/user/:userId/course/:courseId", getCourseRatingByUserAndCourseCtrl.Execute)

		// Rutas protegidas (requieren autenticación)
		courseRatingGroup.POST("", security.JWTMiddleware(), createCourseRatingCtrl.Execute)
		courseRatingGroup.PUT("/:id", security.JWTMiddleware(), updateCourseRatingCtrl.Execute)
		courseRatingGroup.DELETE("/:id", security.JWTMiddleware(), deleteCourseRatingCtrl.Execute)
	}
}
