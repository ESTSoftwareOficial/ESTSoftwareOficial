package infrastructure

import (
	"estsoftwareoficial/src/core"
	"estsoftwareoficial/src/resource_types/application"
	"estsoftwareoficial/src/resource_types/infrastructure/adapters"
	"estsoftwareoficial/src/resource_types/infrastructure/controllers"
)

type DependenciesResourceTypes struct {
	CreateResourceTypeController  *controllers.CreateResourceTypeController
	GetAllResourceTypesController *controllers.GetAllResourceTypesController
	GetResourceTypeByIdController *controllers.GetResourceTypeByIdController
	UpdateResourceTypeController  *controllers.UpdateResourceTypeController
	DeleteResourceTypeController  *controllers.DeleteResourceTypeController
}

func InitResourceTypes() *DependenciesResourceTypes {
	conn := core.GetDBPool()
	resourceTypeRepo := adapters.NewPostgreSQL(conn.DB)

	createResourceType := application.NewCreateResourceType(resourceTypeRepo)
	getAllResourceTypes := application.NewGetAllResourceTypes(resourceTypeRepo)
	getResourceTypeById := application.NewGetResourceTypeById(resourceTypeRepo)
	updateResourceType := application.NewUpdateResourceType(resourceTypeRepo)
	deleteResourceType := application.NewDeleteResourceType(resourceTypeRepo)

	return &DependenciesResourceTypes{
		CreateResourceTypeController:  controllers.NewCreateResourceTypeController(createResourceType),
		GetAllResourceTypesController: controllers.NewGetAllResourceTypesController(getAllResourceTypes),
		GetResourceTypeByIdController: controllers.NewGetResourceTypeByIdController(getResourceTypeById),
		UpdateResourceTypeController:  controllers.NewUpdateResourceTypeController(updateResourceType),
		DeleteResourceTypeController:  controllers.NewDeleteResourceTypeController(deleteResourceType),
	}
}
