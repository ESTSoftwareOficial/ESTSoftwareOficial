package infrastructure

import (
	"estsoftwareoficial/src/categories/application"
	"estsoftwareoficial/src/categories/infrastructure/adapters"
	"estsoftwareoficial/src/categories/infrastructure/controllers"
	"estsoftwareoficial/src/core"
)

type DependenciesCategories struct {
	CreateCategoryController   *controllers.CreateCategoryController
	GetAllCategoriesController *controllers.GetAllCategoriesController
	GetCategoryByIdController  *controllers.GetCategoryByIdController
	UpdateCategoryController   *controllers.UpdateCategoryController
	DeleteCategoryController   *controllers.DeleteCategoryController
}

func InitCategories() *DependenciesCategories {
	conn := core.GetDBPool()
	categoryRepo := adapters.NewPostgreSQL(conn.DB)

	createCategory := application.NewCreateCategory(categoryRepo)
	getAllCategories := application.NewGetAllCategories(categoryRepo)
	getCategoryById := application.NewGetCategoryById(categoryRepo)
	updateCategory := application.NewUpdateCategory(categoryRepo)
	deleteCategory := application.NewDeleteCategory(categoryRepo)

	return &DependenciesCategories{
		CreateCategoryController:   controllers.NewCreateCategoryController(createCategory),
		GetAllCategoriesController: controllers.NewGetAllCategoriesController(getAllCategories),
		GetCategoryByIdController:  controllers.NewGetCategoryByIdController(getCategoryById),
		UpdateCategoryController:   controllers.NewUpdateCategoryController(updateCategory),
		DeleteCategoryController:   controllers.NewDeleteCategoryController(deleteCategory),
	}
}
