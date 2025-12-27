package routes

import (
	"estsoftwareoficial/src/core/security"
	"estsoftwareoficial/src/resource_types/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigureResourceTypeRoutes(
	router *gin.Engine,
	createResourceTypeCtrl *controllers.CreateResourceTypeController,
	getAllResourceTypesCtrl *controllers.GetAllResourceTypesController,
	getResourceTypeByIdCtrl *controllers.GetResourceTypeByIdController,
	updateResourceTypeCtrl *controllers.UpdateResourceTypeController,
	deleteResourceTypeCtrl *controllers.DeleteResourceTypeController,
) {
	resourceTypeGroup := router.Group("/resource-types")
	{
		// Rutas públicas
		resourceTypeGroup.GET("", getAllResourceTypesCtrl.Execute)
		resourceTypeGroup.GET("/:id", getResourceTypeByIdCtrl.Execute)

		// Rutas protegidas (Admin o Instructor)
		resourceTypeGroup.POST("", security.JWTMiddleware(), security.RequireAnyRole(1, 2), createResourceTypeCtrl.Execute)
		resourceTypeGroup.PUT("/:id", security.JWTMiddleware(), security.RequireAnyRole(1, 2), updateResourceTypeCtrl.Execute)

		// Solo Admin puede eliminar
		resourceTypeGroup.DELETE("/:id", security.JWTMiddleware(), security.RequireRole(1), deleteResourceTypeCtrl.Execute)
	}
}
