package application

import (
	"errors"

	"estsoftwareoficial/src/technologies/domain"
	"estsoftwareoficial/src/technologies/domain/entities"
)

type UpdateTechnology struct {
	technologyRepo domain.TechnologyRepository
}

func NewUpdateTechnology(technologyRepo domain.TechnologyRepository) *UpdateTechnology {
	return &UpdateTechnology{technologyRepo: technologyRepo}
}

func (ut *UpdateTechnology) Execute(technology *entities.Technology) error {
	existingTechnology, err := ut.technologyRepo.GetByID(technology.ID)
	if err != nil {
		return err
	}
	if existingTechnology == nil {
		return errors.New("tecnología no encontrada")
	}

	technologyWithSameName, _ := ut.technologyRepo.GetByName(technology.Name)
	if technologyWithSameName != nil && technologyWithSameName.ID != technology.ID {
		return errors.New("ya existe otra tecnología con ese nombre")
	}

	return ut.technologyRepo.Update(technology)
}
