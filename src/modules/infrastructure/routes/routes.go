package routes

import (
	"estsoftwareoficial/src/core/security"
	"estsoftwareoficial/src/modules/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigureModuleRoutes(
	router *gin.Engine,
	createModuleCtrl *controllers.CreateModuleController,
	getAllModulesCtrl *controllers.GetAllModulesController,
	getModuleByIdCtrl *controllers.GetModuleByIdController,
	getModulesByCourseCtrl *controllers.GetModulesByCourseController,
	updateModuleCtrl *controllers.UpdateModuleController,
	deleteModuleCtrl *controllers.DeleteModuleController,
	reorderModulesCtrl *controllers.ReorderModulesController,
) {
	moduleGroup := router.Group("/modules")
	{
		// Rutas públicas
		moduleGroup.GET("", getAllModulesCtrl.Execute)
		moduleGroup.GET("/:id", getModuleByIdCtrl.Execute)
		moduleGroup.GET("/course/:courseId", getModulesByCourseCtrl.Execute)

		// Rutas protegidas (Admin o Instructor)
		moduleGroup.POST("", security.JWTMiddleware(), security.RequireAnyRole(1, 2), createModuleCtrl.Execute)
		moduleGroup.PUT("/:id", security.JWTMiddleware(), security.RequireAnyRole(1, 2), updateModuleCtrl.Execute)
		moduleGroup.PUT("/reorder", security.JWTMiddleware(), security.RequireAnyRole(1, 2), reorderModulesCtrl.Execute)

		// Solo Admin puede eliminar
		moduleGroup.DELETE("/:id", security.JWTMiddleware(), security.RequireRole(1), deleteModuleCtrl.Execute)
	}
}
