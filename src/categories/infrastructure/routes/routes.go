package routes

import (
	"estsoftwareoficial/src/categories/infrastructure/controllers"
	"estsoftwareoficial/src/core/security"

	"github.com/gin-gonic/gin"
)

func ConfigureCategoryRoutes(
	router *gin.Engine,
	createCategoryCtrl *controllers.CreateCategoryController,
	getAllCategoriesCtrl *controllers.GetAllCategoriesController,
	getCategoryByIdCtrl *controllers.GetCategoryByIdController,
	updateCategoryCtrl *controllers.UpdateCategoryController,
	deleteCategoryCtrl *controllers.DeleteCategoryController,
) {
	categoryGroup := router.Group("/categories")
	{
		categoryGroup.GET("", getAllCategoriesCtrl.Execute)
		categoryGroup.GET("/:id", getCategoryByIdCtrl.Execute)

		categoryGroup.POST("", security.JWTMiddleware(), security.RequireAnyRole(1, 2), createCategoryCtrl.Execute)
		categoryGroup.PUT("/:id", security.JWTMiddleware(), security.RequireAnyRole(1, 2), updateCategoryCtrl.Execute)
		categoryGroup.DELETE("/:id", security.JWTMiddleware(), security.RequireRole(1), deleteCategoryCtrl.Execute)
	}
}
