package domain

import "estsoftwareoficial/src/resource_types/domain/entities"

type ResourceTypeRepository interface {
	Save(resourceType *entities.ResourceType) (*entities.ResourceType, error)
	GetByID(id int) (*entities.ResourceType, error)
	GetAll() ([]*entities.ResourceType, error)
	Update(resourceType *entities.ResourceType) error
	Delete(id int) error
	GetByName(name string) (*entities.ResourceType, error)
}
