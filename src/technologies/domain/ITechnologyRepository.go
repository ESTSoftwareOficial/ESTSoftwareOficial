package domain

import "estsoftwareoficial/src/technologies/domain/entities"

type TechnologyRepository interface {
	Save(technology *entities.Technology) (*entities.Technology, error)
	GetByID(id int) (*entities.Technology, error)
	GetAll() ([]*entities.Technology, error)
	Update(technology *entities.Technology) error
	Delete(id int) error
	GetByName(name string) (*entities.Technology, error)
}
