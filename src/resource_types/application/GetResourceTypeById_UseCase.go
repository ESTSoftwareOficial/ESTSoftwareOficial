package application

import (
	"estsoftwareoficial/src/resource_types/domain"
	"estsoftwareoficial/src/resource_types/domain/entities"
)

type GetResourceTypeById struct {
	resourceTypeRepo domain.ResourceTypeRepository
}

func NewGetResourceTypeById(resourceTypeRepo domain.ResourceTypeRepository) *GetResourceTypeById {
	return &GetResourceTypeById{resourceTypeRepo: resourceTypeRepo}
}

func (grt *GetResourceTypeById) Execute(id int) (*entities.ResourceType, error) {
	return grt.resourceTypeRepo.GetByID(id)
}
