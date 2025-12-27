package application

import (
	"errors"

	"estsoftwareoficial/src/technologies/domain"
	"estsoftwareoficial/src/technologies/domain/entities"
)

type CreateTechnology struct {
	technologyRepo domain.TechnologyRepository
}

func NewCreateTechnology(technologyRepo domain.TechnologyRepository) *CreateTechnology {
	return &CreateTechnology{technologyRepo: technologyRepo}
}

func (ct *CreateTechnology) Execute(technology *entities.Technology) (*entities.Technology, error) {
	if technology.Name == "" {
		return nil, errors.New("el nombre es obligatorio")
	}

	if technology.Icon == "" {
		return nil, errors.New("el icono es obligatorio")
	}

	existingTechnology, _ := ct.technologyRepo.GetByName(technology.Name)
	if existingTechnology != nil {
		return nil, errors.New("ya existe una tecnología con ese nombre")
	}

	return ct.technologyRepo.Save(technology)
}
