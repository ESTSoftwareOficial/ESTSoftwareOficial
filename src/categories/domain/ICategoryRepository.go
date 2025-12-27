package domain

import "estsoftwareoficial/src/categories/domain/entities"

type CategoryRepository interface {
	Save(category *entities.Category) (*entities.Category, error)
	GetByID(id int) (*entities.Category, error)
	GetAll() ([]*entities.Category, error)
	Update(category *entities.Category) error
	Delete(id int) error
	GetByName(name string) (*entities.Category, error)
}
