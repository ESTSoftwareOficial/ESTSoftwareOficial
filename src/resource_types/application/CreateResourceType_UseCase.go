package application

import (
	"errors"

	"estsoftwareoficial/src/resource_types/domain"
	"estsoftwareoficial/src/resource_types/domain/entities"
)

type CreateResourceType struct {
	resourceTypeRepo domain.ResourceTypeRepository
}

func NewCreateResourceType(resourceTypeRepo domain.ResourceTypeRepository) *CreateResourceType {
	return &CreateResourceType{resourceTypeRepo: resourceTypeRepo}
}

func (crt *CreateResourceType) Execute(resourceType *entities.ResourceType) (*entities.ResourceType, error) {
	if resourceType.Name == "" {
		return nil, errors.New("el nombre es obligatorio")
	}

	existingResourceType, _ := crt.resourceTypeRepo.GetByName(resourceType.Name)
	if existingResourceType != nil {
		return nil, errors.New("ya existe un tipo de recurso con ese nombre")
	}

	return crt.resourceTypeRepo.Save(resourceType)
}
