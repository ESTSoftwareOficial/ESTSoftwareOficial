package infrastructure

import (
	"estsoftwareoficial/src/core"
	coursesDomain "estsoftwareoficial/src/courses/domain"
	coursesAdapters "estsoftwareoficial/src/courses/infrastructure/adapters"
	"estsoftwareoficial/src/modules/application"
	"estsoftwareoficial/src/modules/infrastructure/adapters"
	"estsoftwareoficial/src/modules/infrastructure/controllers"
)

type DependenciesModules struct {
	CreateModuleController       *controllers.CreateModuleController
	GetAllModulesController      *controllers.GetAllModulesController
	GetModuleByIdController      *controllers.GetModuleByIdController
	GetModulesByCourseController *controllers.GetModulesByCourseController
	UpdateModuleController       *controllers.UpdateModuleController
	DeleteModuleController       *controllers.DeleteModuleController
	ReorderModulesController     *controllers.ReorderModulesController
}

func InitModules() *DependenciesModules {
	conn := core.GetDBPool()
	moduleRepo := adapters.NewPostgreSQL(conn.DB)

	var courseRepo coursesDomain.CourseRepository = coursesAdapters.NewPostgreSQL(conn.DB)

	createModule := application.NewCreateModule(moduleRepo, courseRepo)
	getAllModules := application.NewGetAllModules(moduleRepo)
	getModuleById := application.NewGetModuleById(moduleRepo)
	getModulesByCourse := application.NewGetModulesByCourse(moduleRepo)
	updateModule := application.NewUpdateModule(moduleRepo)
	deleteModule := application.NewDeleteModule(moduleRepo, courseRepo)
	reorderModules := application.NewReorderModules(moduleRepo)

	return &DependenciesModules{
		CreateModuleController:       controllers.NewCreateModuleController(createModule),
		GetAllModulesController:      controllers.NewGetAllModulesController(getAllModules),
		GetModuleByIdController:      controllers.NewGetModuleByIdController(getModuleById),
		GetModulesByCourseController: controllers.NewGetModulesByCourseController(getModulesByCourse),
		UpdateModuleController:       controllers.NewUpdateModuleController(updateModule),
		DeleteModuleController:       controllers.NewDeleteModuleController(deleteModule),
		ReorderModulesController:     controllers.NewReorderModulesController(reorderModules),
	}
}
