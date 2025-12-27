package routes

import (
	"estsoftwareoficial/src/core/security"
	"estsoftwareoficial/src/technologies/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigureTechnologyRoutes(
	router *gin.Engine,
	createTechnologyCtrl *controllers.CreateTechnologyController,
	getAllTechnologiesCtrl *controllers.GetAllTechnologiesController,
	getTechnologyByIdCtrl *controllers.GetTechnologyByIdController,
	updateTechnologyCtrl *controllers.UpdateTechnologyController,
	deleteTechnologyCtrl *controllers.DeleteTechnologyController,
) {
	technologyGroup := router.Group("/technologies")
	{
		// Rutas públicas
		technologyGroup.GET("", getAllTechnologiesCtrl.Execute)
		technologyGroup.GET("/:id", getTechnologyByIdCtrl.Execute)

		// Rutas protegidas (Admin o Instructor)
		technologyGroup.POST("", security.JWTMiddleware(), security.RequireAnyRole(1, 2), createTechnologyCtrl.Execute)
		technologyGroup.PUT("/:id", security.JWTMiddleware(), security.RequireAnyRole(1, 2), updateTechnologyCtrl.Execute)

		// Solo Admin puede eliminar
		technologyGroup.DELETE("/:id", security.JWTMiddleware(), security.RequireRole(1), deleteTechnologyCtrl.Execute)
	}
}
