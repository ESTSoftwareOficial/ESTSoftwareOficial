package application

import (
	"errors"

	"estsoftwareoficial/src/resource_types/domain"
	"estsoftwareoficial/src/resource_types/domain/entities"
)

type UpdateResourceType struct {
	resourceTypeRepo domain.ResourceTypeRepository
}

func NewUpdateResourceType(resourceTypeRepo domain.ResourceTypeRepository) *UpdateResourceType {
	return &UpdateResourceType{resourceTypeRepo: resourceTypeRepo}
}

func (urt *UpdateResourceType) Execute(resourceType *entities.ResourceType) error {
	existingResourceType, err := urt.resourceTypeRepo.GetByID(resourceType.ID)
	if err != nil {
		return err
	}
	if existingResourceType == nil {
		return errors.New("tipo de recurso no encontrado")
	}

	resourceTypeWithSameName, _ := urt.resourceTypeRepo.GetByName(resourceType.Name)
	if resourceTypeWithSameName != nil && resourceTypeWithSameName.ID != resourceType.ID {
		return errors.New("ya existe otro tipo de recurso con ese nombre")
	}

	return urt.resourceTypeRepo.Update(resourceType)
}
