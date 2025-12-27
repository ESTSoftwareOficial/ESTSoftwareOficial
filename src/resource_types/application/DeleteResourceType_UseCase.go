package application

import (
	"errors"

	"estsoftwareoficial/src/resource_types/domain"
)

type DeleteResourceType struct {
	resourceTypeRepo domain.ResourceTypeRepository
}

func NewDeleteResourceType(resourceTypeRepo domain.ResourceTypeRepository) *DeleteResourceType {
	return &DeleteResourceType{resourceTypeRepo: resourceTypeRepo}
}

func (drt *DeleteResourceType) Execute(id int) error {
	existingResourceType, err := drt.resourceTypeRepo.GetByID(id)
	if err != nil {
		return err
	}
	if existingResourceType == nil {
		return errors.New("tipo de recurso no encontrado")
	}

	return drt.resourceTypeRepo.Delete(id)
}
