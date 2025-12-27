package infrastructure

import (
	"estsoftwareoficial/src/core"
	"estsoftwareoficial/src/technologies/application"
	"estsoftwareoficial/src/technologies/infrastructure/adapters"
	"estsoftwareoficial/src/technologies/infrastructure/controllers"
)

type DependenciesTechnologies struct {
	CreateTechnologyController   *controllers.CreateTechnologyController
	GetAllTechnologiesController *controllers.GetAllTechnologiesController
	GetTechnologyByIdController  *controllers.GetTechnologyByIdController
	UpdateTechnologyController   *controllers.UpdateTechnologyController
	DeleteTechnologyController   *controllers.DeleteTechnologyController
}

func InitTechnologies() *DependenciesTechnologies {
	conn := core.GetDBPool()
	technologyRepo := adapters.NewPostgreSQL(conn.DB)

	createTechnology := application.NewCreateTechnology(technologyRepo)
	getAllTechnologies := application.NewGetAllTechnologies(technologyRepo)
	getTechnologyById := application.NewGetTechnologyById(technologyRepo)
	updateTechnology := application.NewUpdateTechnology(technologyRepo)
	deleteTechnology := application.NewDeleteTechnology(technologyRepo)

	return &DependenciesTechnologies{
		CreateTechnologyController:   controllers.NewCreateTechnologyController(createTechnology),
		GetAllTechnologiesController: controllers.NewGetAllTechnologiesController(getAllTechnologies),
		GetTechnologyByIdController:  controllers.NewGetTechnologyByIdController(getTechnologyById),
		UpdateTechnologyController:   controllers.NewUpdateTechnologyController(updateTechnology),
		DeleteTechnologyController:   controllers.NewDeleteTechnologyController(deleteTechnology),
	}
}
