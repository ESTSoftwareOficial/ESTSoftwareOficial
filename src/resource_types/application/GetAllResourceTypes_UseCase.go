package application

import (
	"estsoftwareoficial/src/resource_types/domain"
	"estsoftwareoficial/src/resource_types/domain/entities"
)

type GetAllResourceTypes struct {
	resourceTypeRepo domain.ResourceTypeRepository
}

func NewGetAllResourceTypes(resourceTypeRepo domain.ResourceTypeRepository) *GetAllResourceTypes {
	return &GetAllResourceTypes{resourceTypeRepo: resourceTypeRepo}
}

func (grt *GetAllResourceTypes) Execute() ([]*entities.ResourceType, error) {
	return grt.resourceTypeRepo.GetAll()
}
